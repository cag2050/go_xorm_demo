package sinks

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	compressSuffix = ".gz"

	BackupTimeDayFormat  = "20060102"
	BackupTimeHourFormat = "20060102-15"
)

var _ io.WriteCloser = (*FileLogger)(nil)

type IFileRotateRule interface {
	Check(logger *FileLogger) bool
	FormatFileName(fileTime time.Time, timeFormat string, filePrefix string, fileExt string) string
}

type FileLogger struct {
	FileName         string
	Compress         bool
	RotateRule       IFileRotateRule
	MaxAge           int
	MaxBackups       int
	BackupTimeFormat string

	fileSize   int64
	createTime time.Time
	file       *os.File
	mu         sync.Mutex

	millCh    chan bool
	startMill sync.Once

	fileDir    string
	filePrefix string
	fileExt    string
}

func (thiz *FileLogger) init() {
	if thiz.FileName == "" {
		name := filepath.Base(os.Args[0]) + ".log"
		thiz.FileName = filepath.Join(os.TempDir(), name)
	}

	thiz.fileExt = filepath.Ext(thiz.FileName)
	if thiz.fileExt == "" {
		thiz.fileExt += ".log"
		thiz.FileName += ".log"
	} else if thiz.fileExt == "." {
		thiz.fileExt += "log"
		thiz.FileName += "log"
	}

	thiz.fileDir = filepath.Dir(thiz.FileName)

	name := filepath.Base(thiz.FileName)
	thiz.filePrefix = name[:len(name)-len(thiz.fileExt)] + "-"

	if thiz.BackupTimeFormat == "" {
		thiz.BackupTimeFormat = "20060102T150405"
	}
}

//implements io.Writer
func (thiz *FileLogger) Write(p []byte) (n int, err error) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()

	if thiz.file == nil {
		if err = thiz.openExistingOrNew(); err != nil {
			return 0, err
		}
	}

	if thiz.RotateRule != nil {
		if thiz.RotateRule.Check(thiz) {
			if err := thiz.rotate(); err != nil {
				return 0, err
			}
		}
	}

	n, err = thiz.file.Write(p)
	thiz.fileSize += int64(n)
	return n, err
}

//implements io.Closer
func (thiz *FileLogger) Close() error {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	return thiz.close()
}

func (thiz *FileLogger) close() error {
	if thiz.file == nil {
		return nil
	}
	err := thiz.file.Close()
	thiz.file = nil
	return err
}

func (thiz *FileLogger) openExistingOrNew() error {
	thiz.mill()
	info, err := os.Stat(thiz.FileName)
	if os.IsNotExist(err) {
		return thiz.openNew()
	}

	if err != nil {
		return fmt.Errorf("error getting log file info: %s", err)
	}

	thiz.fileSize = info.Size()
	thiz.createTime = info.ModTime()

	if thiz.RotateRule != nil {
		if thiz.RotateRule.Check(thiz) {
			return thiz.rotate()
		}
	}

	file, err := os.OpenFile(thiz.FileName, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return thiz.openNew()
	}

	thiz.file = file
	thiz.redirectStdErrToFile()
	return nil
}

func (thiz *FileLogger) openNew() error {
	err := os.MkdirAll(thiz.fileDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	_, err = os.Stat(thiz.FileName)
	if err == nil {
		t := thiz.createTime
		var newname string
		if thiz.RotateRule != nil {
			newname = thiz.RotateRule.FormatFileName(t, thiz.BackupTimeFormat, thiz.filePrefix, thiz.fileExt)
		}
		if newname == "" {
			newname = fmt.Sprintf("%s%s%s", thiz.filePrefix, t.Format(thiz.BackupTimeFormat), thiz.fileExt)
		}
		newname = filepath.Join(thiz.fileDir, newname)

		if err := os.Rename(thiz.FileName, newname); err != nil {
			return fmt.Errorf("can't rename log file: %s", err)
		}
	}

	f, err := os.OpenFile(thiz.FileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}
	thiz.file = f
	thiz.fileSize = 0
	thiz.createTime = time.Now()
	thiz.redirectStdErrToFile()
	return nil
}

func (thiz *FileLogger) rotate() error {
	if err := thiz.close(); err != nil {
		return err
	}
	if err := thiz.openNew(); err != nil {
		return err
	}
	thiz.mill()
	return nil
}

func (thiz *FileLogger) mill() {
	thiz.startMill.Do(func() {
		thiz.init()
		if thiz.RotateRule != nil {
			thiz.millCh = make(chan bool, 1)
			go thiz.millRun()
		}
	})
	if thiz.RotateRule != nil {
		select {
		case thiz.millCh <- true:
		default:
		}
	}
}

func (thiz *FileLogger) millRun() {
	for _ = range thiz.millCh {
		_ = thiz.millRunOnce()
	}
}

func (thiz *FileLogger) millRunOnce() error {
	if thiz.MaxBackups == 0 && thiz.MaxAge == 0 && !thiz.Compress {
		return nil
	}

	files, err := thiz.oldLogFiles()
	if err != nil {
		return err
	}

	var compress, remove []logInfo
	if thiz.MaxBackups > 0 && thiz.MaxBackups < len(files) {
		preserved := make(map[string]bool)
		var remaining []logInfo
		for _, f := range files {
			fn := f.Name()
			if strings.HasSuffix(fn, compressSuffix) {
				fn = fn[:len(fn)-len(compressSuffix)]
			}
			preserved[fn] = true

			if len(preserved) > thiz.MaxBackups {
				remove = append(remove, f)
			} else {
				remaining = append(remaining, f)
			}
		}
		files = remaining
	}

	if thiz.MaxAge > 0 {
		diff := time.Second * time.Duration(thiz.MaxAge)
		cutOff := time.Now().Add(-1 * time.Duration(diff))
		var remaining []logInfo
		for _, f := range files {
			if f.timestamp.Before(cutOff) {
				//fmt.Println("f:", f.Name(), "  now:", time.Now(), "  t:", f.timestamp, "  cutOff:", cutOff)
				remove = append(remove, f)
			} else {
				remaining = append(remaining, f)
			}
		}
		files = remaining
	}

	if thiz.Compress {
		for _, f := range files {
			if !strings.HasSuffix(f.Name(), compressSuffix) {
				compress = append(compress, f)
			}
		}
	}

	for _, f := range remove {
		errRemove := os.Remove(filepath.Join(thiz.fileDir, f.Name()))
		if err == nil && errRemove != nil {
			err = errRemove
		}
	}

	for _, f := range compress {
		fn := filepath.Join(thiz.fileDir, f.Name())
		errCompress := compressLogFile(fn, fn+compressSuffix)
		if err == nil && errCompress != nil {
			err = errCompress
		}
	}

	return err
}

type logInfo struct {
	timestamp time.Time
	os.FileInfo
}

type byFormatTime []logInfo

func (b byFormatTime) Less(i, j int) bool {
	return b[i].timestamp.After(b[j].timestamp)
}

func (b byFormatTime) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byFormatTime) Len() int {
	return len(b)
}
