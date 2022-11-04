package log

import (
	"time"

	"github.com/sirupsen/logrus"
)

type Fields logrus.Fields
type Level logrus.Level
type Formatter logrus.Formatter

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

var (
	std = New()
)

func StandardLogger() *Logger {
	return std
}

//基于配置文件自动初始化
func Init(cfg LogSimpleConfig) error {
	return std.Init(cfg)
}

//单一日志文件
func FileOutput(fileName string) {
	std.FileOutput(fileName)
}

//按天分割的日志文件
func FileDailyRotateOutput(fileName string, maxBackupDays int, compress bool) {
	std.FileDailyRotateOutput(fileName, maxBackupDays, compress)
}

//按小时分割的日志文件
func FileHourlyRotateOutput(fileName string, maxBackupHours int, compress bool) {
	std.FileHourlyRotateOutput(fileName, maxBackupHours, compress)
}

//按文件大小分割的日志文件
func FileSizeRotateOutput(fileName string, maxFileSize int64, maxBackups int, compress bool) {
	std.FileSizeRotateOutput(fileName, maxFileSize, maxBackups, compress)
}

//设置输出日志前缀
func SetOutputPrefix(prefix string) {
	std.SetOutputPrefix(prefix)
}

//控制台输出
func ConsoleOutput() {
	std.ConsoleOutput()
}

func SetFormatter(formatter Formatter) {
	std.SetFormatter(formatter)
}

func SetLevel(level Level) {
	std.SetLevel(level)
}

// todo
// func NewEntry(logger *Logger) *Entry {
// 	return (*Entry)(logrus.NewEntry(logger.Logger))
// }

func WithTime(t time.Time) *Entry {
	return std.WithTime(t)
}

func WithError(err error) *Entry {
	return std.WithError(err)
}

func WithField(key string, value interface{}) *Entry {
	return std.WithField(key, value)
}

func WithFields(fields Fields) *Entry {
	return std.WithFields(fields)
}

func WithMaps(fields map[string]interface{}) *Entry {
	return std.WithMaps(fields)
}

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Print(args ...interface{}) {
	std.Print(args...)
}

func Log(level Level, args ...interface{}) {
	std.Log(logrus.Level(level), args...)
}

func Logf(level Level, format string, args ...interface{}) {
	std.Logf(logrus.Level(level), format, args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Warning(args ...interface{}) {
	std.Warning(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Tracef(format string, args ...interface{}) {
	std.Tracef(format, args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Printf(format string, args ...interface{}) {
	std.Printf(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Traceln(args ...interface{}) {
	std.Traceln(args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Println(args ...interface{}) {
	std.Println(args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}
