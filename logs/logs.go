package logs

import (
	"github.com/zhenwusw/logan/logs/logs"
	"io"
)

type (
	Logs interface {
		SetOutput(show io.Writer) Logs
		Reset()
		GoOn()
		StealOne()
		Close()
		Status()
		DelLogger()
		SetLogger()

		Debug(format string, v ...interface{})
		Informational(format string, v ...interface{})
		App(format string, v ...interface{})
		Notice(format string, v ...interface{})
		Warning(format string, v ...interface{})
		Error(format string, v ...interface{})
		Critical(format string, v ...interface{})
		Alert(format string, v ...interface{})
		Emergency(format string, v ...interface{})
	}
	mylog struct {
		*logs.BeeLogger
	}
)

/*
var Log = func() Logs {
	p, _ := path.Split(config.LOG)
	// 不存在目录时创建目录
}()

func (self *mylog) SetOutput(show io.Writer) Logs {
}
*/
