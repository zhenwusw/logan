package logs

import (
	"github.com/zhenwusw/logan/config"
	"github.com/zhenwusw/logan/logs/logs"
	"io"
	"os"
	"path"
)

type (
	Logs interface {
		//SetOutput(show io.Writer) Logs
		//Reset()
		//GoOn()
		//StealOne()
		//Close()
		//Status()
		//DelLogger()
		//SetLogger()

		//Debug(format string, v ...interface{})
		//Informational(format string, v ...interface{})
		//App(format string, v ...interface{})
		//Notice(format string, v ...interface{})
		Warning(format string, v ...interface{})
		//Error(format string, v ...interface{})
		//Critical(format string, v ...interface{})
		//Alert(format string, v ...interface{})
		//Emergency(format string, v ...interface{})
	}

	mylog struct {
		*logs.BeeLogger
	}
)

var Log = func() Logs {
	p, _ := path.Split(config.LOG)
	// 不存在目录时创建目录
	d, err := os.Stat(p)
	if err != nil || !d.IsDir() {
		if err := os.MkdirAll(p, 0777); err != nil {
			// Log.Error("Error: %v\n", err)
		}
	}

	ml := &mylog{
		BeeLogger: logs.NewLogger(config.LOG_CAP, config.LOG_FEEDBACK_LEVEL),
	}

	// 是否打印行信息
	ml.BeeLogger.EnableFuncCallDepth(config.LOG_LINEINFO)
	// 全局日志打印级别（亦是日志文件输出级别）
	ml.BeeLogger.SetLevel(config.LOG_LEVEL)
	// 是否异步输出日志
	ml.BeeLogger.Async(config.LOG_ASYNC)
	// 设置日志显示位置
	ml.BeeLogger.SetLogger("console", map[string]interface{}{
		"level": config.LOG_CONSOLE_LEVEL,
	})
	return ml
}()

func (self *mylog) SetOutput(show io.Writer) Logs {
	// do the job
	return self
}
