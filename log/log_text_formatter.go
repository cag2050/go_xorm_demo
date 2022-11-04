package log

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = time.RFC3339
)

type LogTextFormatter struct {
	TimestampFormat string
}

func (thiz *LogTextFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestampFormat := thiz.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}

	b.WriteString(fmt.Sprintf("time=\"%v\"", entry.Time.Format(timestampFormat)))
	b.WriteString(fmt.Sprintf(" level=%v", entry.Level.String()))

	for key, val := range entry.Data {
		thiz.appendKeyValue(b, key, val)
	}

	b.WriteString(fmt.Sprintf(" msg=%v", entry.Message))
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (thiz *LogTextFormatter) needsQuoting(text string) bool {
	if len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if !((ch >= 'a' && ch <= 'z') ||
			(ch >= 'A' && ch <= 'Z') ||
			(ch >= '0' && ch <= '9') ||
			ch == '-' || ch == '.' || ch == '_' || ch == '/' || ch == '@' || ch == '^' || ch == '+') {
			return true
		}
	}
	return false
}

func (thiz *LogTextFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteByte(' ')
	}
	b.WriteString(key)
	b.WriteByte('=')
	thiz.appendValue(b, value)
}

func (thiz *LogTextFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	if !thiz.needsQuoting(stringVal) {
		b.WriteString(stringVal)
	} else {
		b.WriteString(fmt.Sprintf("%q", stringVal))
	}
}
