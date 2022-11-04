package sinks

import (
	"fmt"
	"time"
)

var _ IFileRotateRule = (*TimeIntervalRotateRule)(nil)
var _ IFileRotateRule = (*FileMaxSizeRotateRule)(nil)

type TimeIntervalRotateRule struct {
	TimeInterval time.Duration
}

func (thiz *TimeIntervalRotateRule) Check(logger *FileLogger) bool {
	interval := thiz.TimeInterval.Nanoseconds()
	rotateTime := logger.createTime.UnixNano()/interval*interval + interval
	if time.Now().UnixNano() >= rotateTime {
		return true
	}
	return false
}

func (thiz *TimeIntervalRotateRule) FormatFileName(fileTime time.Time, timeFormat string, filePrefix string, fileExt string) string {
	interval := thiz.TimeInterval.Nanoseconds()
	fileTime = time.Unix(0, fileTime.UnixNano()/interval*interval)

	return fmt.Sprintf("%s%s%s",
		filePrefix,
		fileTime.Format(timeFormat),
		fileExt)
}

type FileMaxSizeRotateRule struct {
	MaxSize int64
}

func (thiz *FileMaxSizeRotateRule) Check(logger *FileLogger) bool {
	return logger.fileSize >= thiz.MaxSize
}

func (thiz *FileMaxSizeRotateRule) FormatFileName(fileTime time.Time, timeFormat string, filePrefix string, fileExt string) string {
	return ""
}
