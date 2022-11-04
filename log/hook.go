package log

import (
	"github.com/cag2050/go_xorm_demo/utils/gid"

	"github.com/sirupsen/logrus"
)

type hook struct {
	logger *Logger
}

func (thiz *hook) Fire(entry *logrus.Entry) error {
	showGID := thiz.logger.showGID
	if showGID {
		entry.Data["gid"] = gid.Get()

	}
	prefix := thiz.logger.prefix
	if prefix != "" {
		entry.Data["logPrefix"] = prefix
	}

	return nil
}

func (thisz *hook) Levels() []logrus.Level {
	return logrus.AllLevels
}
