package cmd

import (
	"flag"
	"fmt"
	"github.com/zhenwusw/logan/app"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/runtime/status"
	"strconv"
)

var (
	spiderflag *string
)

// 获取外部参数
func Flag() {
	// 分类说明
	flag.String("c ******************************************** only for cmd ******************************************** -c", "", "")

	// 蜘蛛列表
	spiderflag = flag.String(
		"c_spider",
		"",
		func() string {
			var spiderlist string
			for k, v := range app.LogicApp.GetSpiderLib() {
				spiderlist += " [" + strconv.Itoa(k) + "]" + v.GetName()
			}
			return "   <蜘蛛列表: 选择多蜘蛛以 \",\" 间隔>\r\n" + spiderlist
		}())

	// 备注说明
	flag.String(
		"c_z",
		"",
		"CMD-EXAMPLE: $ logan -_ui=cmd -a_mode="+strconv.Itoa(status.OFFLINE)+" -c_spider=3,8 -a_outtype=csv -a_thread=20 -a_pause=300 -a_proxyminute=0 -a_keyins=\"<logan><golang>\" -a_limit=10 -a_success=true -a_failure=true\n",
	)
}

// 执行入口
func Run() {
	// Init Logic App
	// app.LogicApp.Init(cache.Task.Mode, cache.Task.Port, cache.Task.Master)
	app.LogicApp.Init(status.OFFLINE, -1, "")
	//if cache.Task.Mode == status.UNSET {
	//	return
	//}

	// 获取全局参数
	switch app.LogicApp.GetAppConf("Mode").(int) {
	case status.SERVER:
	//
	case status.CLIENT:
	//
	default:
		run()
	}
}

// 运行
func run() {
	fmt.Printf("...... logan-cmd#run() \n")

	// 创建蜘蛛队列
	sps := []*spider.Spider{}
	/*
	*spiderflag = strings.TrimSpace(*spiderflag)
	if *spiderflag == "*" {
		// aps = app.LogicApp.GetSpiderLib()
	} else {
		for _, idx := range strings.Split(*spiderflag, ",") {
			idx = strings.TrimSpace(idx)
			if idx == "" {
				continue
			}
			i, _ := strconv.Atoi(idx)
			sps = append(sps, app.LogicApp.GetSpiderLib()[i])
		}
	}*/

	for _, sp := range app.LogicApp.GetSpiderLib() {
		sps = append(sps, sp)
		fmt.Printf("...... logan-cmd#run appends spider %v\n", sp.Name)
	}

	app.LogicApp.SpiderPrepare(sps).Run()
}
