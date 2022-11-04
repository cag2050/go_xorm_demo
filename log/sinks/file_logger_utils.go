package sinks

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func (thiz *FileLogger) oldLogFiles() ([]logInfo, error) {
	files, err := ioutil.ReadDir(thiz.fileDir)
	if err != nil {
		return nil, fmt.Errorf("can't read log file directory: %s", err)
	}

	logFiles := []logInfo{}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		if strings.HasPrefix(f.Name(), thiz.filePrefix) {
			if strings.HasSuffix(f.Name(), thiz.fileExt) ||
				strings.HasSuffix(f.Name(), thiz.fileExt+compressSuffix) {
				logFiles = append(logFiles, logInfo{
					timestamp: f.ModTime(),
					FileInfo:  f,
				})
			}
		}
	}

	sort.Sort(byFormatTime(logFiles))
	return logFiles, nil
}

func compressLogFile(src, dst string) (err error) {
	f, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open log file: %v", err)
	}
	defer f.Close()

	fi, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat log file: %v", err)
	}

	gzf, err := os.OpenFile(dst, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, fi.Mode())
	if err != nil {
		return fmt.Errorf("failed to open compressed log file: %v", err)
	}
	defer gzf.Close()

	gz := gzip.NewWriter(gzf)

	defer func() {
		if err != nil {
			os.Remove(dst)
			err = fmt.Errorf("failed to compress log file: %v", err)
		}
	}()

	if _, err := io.Copy(gz, f); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	if err := gzf.Close(); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	if err := os.Remove(src); err != nil {
		return err
	}

	return nil
}
