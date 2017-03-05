// 基本业务执行顺序依次为: New() --> [SetLog(io.Writer)-->]Init()-->SpiderPrepare()-->Run()
package app

import (
	"io"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/runtime/cache"
	"github.com/henrylee2cn/pholcus/app/distribute"
)

type (
	App interface {
		SetLog(io.Writer) App
		LogGoOn() App
		LogReset() App
		Init() App
		ReInit() App
		GetAppConf()
		SetAppConf() App
		SpiderPrepare() App
		Run()
		Stop()
		IsRunning() bool
		IsPause() bool
		isStopped() bool
		PauseRecover()
		Status() int
		GetSpiderLib() []*spider.Spider
		// GetTaskJar() *distribute.TaskJar
		// distribute.Distributer
	}

	Logic struct {
		*cache.AppConf
		*spider.SpiderS
	}
)

// 全局唯一的核心接口实例
var LogicApp = New()

func New() App {
	return newLogic()
}

func newLogic() *Logic {
	return &Logic{
		AppConf:
		SpiderSpe
	}
}

// 设置全局log实时显示终端
func (self *Logic) SetLog(w io.Writer) App {
}

// 暂停log打印
func (self *Logic) LogReset() App {
}

// 继续log打印
func (self *Logic) LogGoOn() App {
}

// 获取全局参数
func (self *Logic) GetAppConf(k ...string) interface{} {
}

// 设置全局参数
func (self *Logic) SetAppConf(k string, v interface{}) App {
}

// 使用App前必须进行Init初始化
func (self *Logic) Init() App {
}

// 切换运行模式时使用
func (self *Logic) ReInit() App {
}

// SpiderPrepare() 必须在设置全局运行参数后，Run() 的前一刻执行
// original为spider包中未有过复制操作的原始蜘蛛种类
// 已被显示赋值过的spider将不再重新分配Keyin
// client模式下不调用该方法
func (self *Logic) SpiderPrepare() App {
}

// 获取全部输出方式
func (self *Logic) GetOutputLib() []string {
}

// 获取全部蜘蛛种类
func (self *Logic) GetSpiderLib() []*spider.Spider {
}

// 通过名字获取蜘蛛
func (self *Logic) GetSpiderByName(name string) *spider.Spider {
}

// 返回当前运行模式
func (self *Logic) GetMode() int {
}

// 获取任务库
func (self *Logic) GetTaskJar() *distribute.TaskJar {
}

// 服务器客户端模式下返回节点数
func (self *Logic) CountNodes() int {
}

// 获取蜘蛛队列接口实例
func (self *Logic) GetSpiderQueue() crawler.S














