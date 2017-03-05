package cmd

import (
	"flag"
	"strconv"
	"github.com/zhenwusw/logan/runtime/status"
	"github.com/zhenwusw/logan/app"
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
		"CMD-EXAMPLE: $ logan -_ui=cmd -a_mode=" + strconv.Itoa(status.OFFLINE) + " -c_spider=3,8 -a_outtype=csv -a_thread=20 -a_pause=300 -a_proxyminute=0 -a_keyins=\"<logan><golang>\" -a_limit=10 -a_success=true -a_failure=true\n",
	)
}
