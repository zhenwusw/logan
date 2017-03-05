package exec

import (
	"os/exec"
	"github.com/henrylee2cn/pholcus/config"
	"github.com/zhenwusw/logan/cmd"  // cmd版
	"os"
	"os/signal"
	"github.com/henrylee2cn/pholcus/web"
)

func run(which string) {
	exec.Command("/bin/sh", "-c", "title", config.FULL_NAME).Start()

	// 选择运行界面
	switch which {
	case "cmd":
		cmd.Run()
	case "web":
		fallthrough
	default:
		ctrl := make(chan os.Signal, 1)
		signal.Notify(ctrl, os.Interrupt, os.Kill)
		go web.Run()
		<-ctrl
	}
}