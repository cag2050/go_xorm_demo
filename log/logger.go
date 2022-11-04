package log

import (
	"context"
	"fmt"
	"github.com/cag2050/go_xorm_demo/log/sinks"
	"github.com/cag2050/go_xorm_demo/utils"
	"io"
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger

	mu            sync.Mutex
	fileOutput    io.Writer
	consoleOutput bool
	colored       bool
	prefix        string
	showGID       bool
}

func New() *Logger {
	l := &Logger{
		Logger: &logrus.Logger{
			Out: os.Stdout,
			// Formatter: &logrus.TextFormatter{
			// 	FullTimestamp:   true,
			// 	TimestampFormat: "2006-01-02 15:04:05.000",
			// },
			Formatter: &LogTextFormatter{
				TimestampFormat: "2006-01-02 15:04:05.000",
			},
			Hooks: make(logrus.LevelHooks),
			Level: logrus.DebugLevel,
		},
		showGID: true,
	}
	l.Hooks.Add(&hook{logger: l})
	return l
}

func (thiz *Logger) Init(cfg LogSimpleConfig) error {
	var err error
	var level logrus.Level = logrus.DebugLevel
	var maxFileSize int64

	if cfg.Level != "" {
		if level, err = logrus.ParseLevel(cfg.Level); err != nil {
			return fmt.Errorf("ParseLevel error:%v", err)
		}
	}

	if cfg.Path != "" {
		switch cfg.RotateType {
		case "none", "null", "":
			thiz.FileOutput(cfg.Path)
		case "hourly", "hour":
			thiz.FileHourlyRotateOutput(cfg.Path, int(cfg.MaxStoreFiles), cfg.Compress)
		case "daily", "day":
			thiz.FileDailyRotateOutput(cfg.Path, int(cfg.MaxStoreFiles), cfg.Compress)
		case "size":
			if cfg.MaxFileSize != "" {
				if maxFileSize, err = utils.ParseByteSize(cfg.MaxFileSize); err != nil {
					return fmt.Errorf("ParseByteSize error:%v", err)
				} else if maxFileSize <= 0 {
					return fmt.Errorf("invalid maxFileSize:%v", maxFileSize)
				}
			}
			thiz.FileSizeRotateOutput(cfg.Path, maxFileSize, int(cfg.MaxStoreFiles), cfg.Compress)
		default:
			return fmt.Errorf("invalid rotateType:%v", cfg.RotateType)
		}
	}

	thiz.SetOutputPrefix(cfg.Prefix)
	thiz.SetLevel(Level(level))
	if cfg.Console {
		thiz.ConsoleOutput()
	}

	thiz.Colored(cfg.Colored)
	return nil
}

func (thiz *Logger) ShowGID(show bool) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.showGID = show
}

func (thiz *Logger) Colored(colored bool) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()

	thiz.colored = colored
	if thiz.Logger.Formatter != nil {
		if fmter, _ := thiz.Logger.Formatter.(*logrus.TextFormatter); fmter != nil {
			fmter.ForceColors = colored
		}
	}
}

func (thiz *Logger) updateOutput() {
	if thiz.fileOutput != nil && thiz.consoleOutput {
		thiz.SetOutput(io.MultiWriter(os.Stdout, thiz.fileOutput))
	} else if thiz.fileOutput != nil {
		thiz.SetOutput(thiz.fileOutput)
	} else {
		thiz.SetOutput(os.Stdout)
	}
}

//单一日志文件
func (thiz *Logger) FileOutput(fileName string) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.fileOutput = &sinks.FileLogger{
		FileName: fileName,
	}
	thiz.updateOutput()
}

//按天分割的日志文件
func (thiz *Logger) FileDailyRotateOutput(fileName string, maxBackupDays int, compress bool) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.fileOutput = &sinks.FileLogger{
		FileName:   fileName,
		MaxBackups: maxBackupDays,
		RotateRule: &sinks.TimeIntervalRotateRule{
			TimeInterval: time.Hour * 24,
		},
		BackupTimeFormat: "20060102",
		Compress:         compress,
	}
	thiz.updateOutput()
}

//按小时分割的日志文件
func (thiz *Logger) FileHourlyRotateOutput(fileName string, maxBackupHours int, compress bool) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.fileOutput = &sinks.FileLogger{
		FileName:   fileName,
		MaxBackups: maxBackupHours,
		RotateRule: &sinks.TimeIntervalRotateRule{
			TimeInterval: time.Hour,
		},
		BackupTimeFormat: "20060102-15",
		Compress:         compress,
	}
	thiz.updateOutput()
}

//按文件大小分割的日志文件
func (thiz *Logger) FileSizeRotateOutput(fileName string, maxFileSize int64, maxBackups int, compress bool) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.fileOutput = &sinks.FileLogger{
		FileName:   fileName,
		MaxBackups: maxBackups,
		RotateRule: &sinks.FileMaxSizeRotateRule{
			MaxSize: maxFileSize,
		},
		BackupTimeFormat: "20060102-150405-000",
		Compress:         compress,
	}
	thiz.updateOutput()
}

//控制台输出
func (thiz *Logger) ConsoleOutput() {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	if thiz.consoleOutput {
		return
	}
	thiz.consoleOutput = true
	thiz.updateOutput()
}

//设置输出日志前缀
func (thiz *Logger) SetOutputPrefix(prefix string) {
	thiz.mu.Lock()
	defer thiz.mu.Unlock()
	thiz.prefix = prefix
}

func (thiz *Logger) SetFormatter(formatter Formatter) {
	thiz.Logger.SetFormatter(logrus.Formatter(formatter))
}

func (thiz *Logger) SetLevel(level Level) {
	thiz.Logger.SetLevel(logrus.Level(level))
}

func (thiz *Logger) WithField(key string, value interface{}) *Entry {
	return (*Entry)(thiz.Logger.WithField(key, value))
}

func (thiz *Logger) WithFields(fields Fields) *Entry {
	return (*Entry)(thiz.Logger.WithFields(logrus.Fields(fields)))
}

func (thiz *Logger) WithMaps(fields map[string]interface{}) *Entry {
	return (*Entry)(thiz.Logger.WithFields(logrus.Fields(fields)))
}

func (thiz *Logger) WithError(err error) *Entry {
	return (*Entry)(thiz.Logger.WithError(err))
}

func (thiz *Logger) WithContext(ctx context.Context) *Entry {
	return (*Entry)(thiz.Logger.WithContext(ctx))
}

func (thiz *Logger) WithTime(t time.Time) *Entry {
	return (*Entry)(thiz.Logger.WithTime(t))
}
