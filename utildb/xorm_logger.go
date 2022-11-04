package utildb

import (
	"fmt"
	"github.com/cag2050/go_xorm_demo/log"
	"github.com/cag2050/go_xorm_demo/utils"

	xormLog "xorm.io/xorm/log"
)

var _ xormLog.Logger = (*XormLogger)(nil)

type XormLogger struct {
	logger  utils.ILogger
	level   xormLog.LogLevel
	showSQL bool
	prefix  string
}

func NewXormLogger(prefix string, logger utils.ILogger) *XormLogger {
	if prefix == "" {
		prefix = "[XORM] "
	} else {
		prefix = prefix + " [XORM] "
	}
	if logger == nil {
		logger = log.StandardLogger()
	}
	l := &XormLogger{
		logger: logger,
		prefix: prefix,
		level:  xormLog.LOG_DEBUG,
	}
	return l
}

func (thiz *XormLogger) Error(v ...interface{}) {
	if thiz.level <= xormLog.LOG_ERR {
		thiz.logger.Error(thiz.prefix, fmt.Sprint(v...))
	}
}

func (thiz *XormLogger) Errorf(format string, v ...interface{}) {
	if thiz.level <= xormLog.LOG_ERR {
		thiz.logger.Error(thiz.prefix, fmt.Sprintf(format, v...))
	}
}

func (thiz *XormLogger) Debug(v ...interface{}) {
	if thiz.level <= xormLog.LOG_DEBUG {
		thiz.logger.Debug(thiz.prefix, fmt.Sprint(v...))
	}
}

func (thiz *XormLogger) Debugf(format string, v ...interface{}) {
	if thiz.level <= xormLog.LOG_DEBUG {
		thiz.logger.Debug(thiz.prefix, fmt.Sprintf(format, v...))
	}
}

func (thiz *XormLogger) Info(v ...interface{}) {
	if thiz.level <= xormLog.LOG_INFO {
		thiz.logger.Info(thiz.prefix, fmt.Sprint(v...))
	}
}

func (thiz *XormLogger) Infof(format string, v ...interface{}) {
	if thiz.level <= xormLog.LOG_INFO {
		thiz.logger.Info(thiz.prefix, fmt.Sprintf(format, v...))
	}
}

func (thiz *XormLogger) Warn(v ...interface{}) {
	if thiz.level <= xormLog.LOG_WARNING {
		thiz.logger.Warn(thiz.prefix, fmt.Sprint(v...))
	}
}

func (thiz *XormLogger) Warnf(format string, v ...interface{}) {
	if thiz.level <= xormLog.LOG_WARNING {
		thiz.logger.Warn(thiz.prefix, fmt.Sprintf(format, v...))
	}
}

func (thiz *XormLogger) Level() xormLog.LogLevel {
	return thiz.level
}

func (thiz *XormLogger) SetLevel(l xormLog.LogLevel) {
	thiz.level = l
}

func (thiz *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		thiz.showSQL = true
		return
	}
	thiz.showSQL = show[0]
}

func (thiz *XormLogger) IsShowSQL() bool {
	return thiz.showSQL
}
