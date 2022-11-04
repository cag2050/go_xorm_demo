package sinks

import (
	"os"
	"syscall"
)

func (thiz *FileLogger) redirectStdErrToFile() {
	syscall.Dup2(int(thiz.file.Fd()), int(os.Stderr.Fd()))
}
