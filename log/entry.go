package log

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type Entry logrus.Entry

func (thiz *Entry) WithError(err error) *Entry {
	return (*Entry)(((*logrus.Entry)(thiz)).WithError(err))
}

func (thiz *Entry) WithContext(ctx context.Context) *Entry {
	return (*Entry)(((*logrus.Entry)(thiz)).WithContext(ctx))
}

func (thiz *Entry) WithField(key string, value interface{}) *Entry {
	return (*Entry)(((*logrus.Entry)(thiz)).WithField(key, value))
}

func (thiz *Entry) WithFields(fields Fields) *Entry {
	return (*Entry)(((*logrus.Entry)(thiz)).WithFields(logrus.Fields(fields)))
}

func (thiz *Entry) WithTime(t time.Time) *Entry {
	return (*Entry)(((*logrus.Entry)(thiz)).WithTime(t))
}

func (thiz *Entry) Log(level Level, args ...interface{}) {
	((*logrus.Entry)(thiz)).Log(logrus.Level(level), args...)
}

func (thiz *Entry) Logf(level Level, format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Logf(logrus.Level(level), format, args...)
}

func (thiz *Entry) Trace(args ...interface{}) {
	((*logrus.Entry)(thiz)).Trace(args...)
}

func (thiz *Entry) Debug(args ...interface{}) {
	((*logrus.Entry)(thiz)).Debug(args...)
}

func (thiz *Entry) Info(args ...interface{}) {
	((*logrus.Entry)(thiz)).Info(args...)
}

func (thiz *Entry) Warn(args ...interface{}) {
	((*logrus.Entry)(thiz)).Warn(args...)
}

func (thiz *Entry) Error(args ...interface{}) {
	((*logrus.Entry)(thiz)).Error(args...)
}

func (thiz *Entry) Fatal(args ...interface{}) {
	((*logrus.Entry)(thiz)).Fatal(args...)
}

func (thiz *Entry) Panic(args ...interface{}) {
	((*logrus.Entry)(thiz)).Panic(args...)
}

func (thiz *Entry) Tracef(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Tracef(format, args...)
}

func (thiz *Entry) Debugf(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Debugf(format, args...)
}

func (thiz *Entry) Infof(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Infof(format, args...)
}

func (thiz *Entry) Warnf(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Warnf(format, args...)
}

func (thiz *Entry) Errorf(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Errorf(format, args...)
}

func (thiz *Entry) Fatalf(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Fatalf(format, args...)
}

func (thiz *Entry) Panicf(format string, args ...interface{}) {
	((*logrus.Entry)(thiz)).Panicf(format, args...)
}
