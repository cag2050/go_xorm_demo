package sinks

import (
	"os"
	"syscall"
)

func (thiz *FileLogger) redirectStdErrToFile() {
	syscall.Dup3(int(thiz.file.Fd()), int(os.Stderr.Fd()), 0)
}
