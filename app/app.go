// 基本业务执行顺序依次为:
// New() --> [SetLog(io.Writer)-->]Init()-->SpiderPrepare()-->Run()
package app

import (
	"io"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/runtime/cache"
	"github.com/zhenwusw/logan/runtime/status"
)

type (
	App interface {
		Init(mode int, port int, master string, w ...io.Writer) App
		// ReInit(mode int, port int, master string, w ...io.Writer) App
		Run()
		Stop()

		IsRunning() bool
		IsPause() bool
		IsStopped() bool

		SpiderPrepare(original []*spider.Spider) App
		GetSpiderLib() []*spider.Spider

		GetOutputLib() []string // 获取全部输出方式

		// SetLog(io.Writer) App
		// LogGoOn() App
		// LogReset() App
		// GetAppConf()
		// SetAppConf() App
		//
		// PauseRecover()
		// Status() int
		// GetTaskJar() *distribute.TaskJar
		// distribute.Distributer
	}

	Logic struct {
		*cache.AppConf        // 全局配置
		*spider.SpiderSpecies // 全部蜘蛛种类
		status int
		finish chan bool
		// crawler.SpiderQueue   // 当前任务的蜘蛛队列
		// *distribute.TaskJar
		// crawler.CrawlerPool
		// teleport.Teleport
		// sum
		// takeTime

		// canSocketLog
		// sync.RWMutex
	}
)

// 全局唯一的核心接口实例
var LogicApp = New()

func New() App {
	return newLogic()
}

func newLogic() *Logic {
	return &Logic{
		AppConf:       cache.Task,
		SpiderSpecies: spider.Species,
		status:        status.STOPPED,
	}
}

// 使用App前必须进行Init初始化
func (self *Logic) Init(mode int, port int, master string, w ...io.Writer) App {
	// 初始化
	return self
}

// 运行任务
func (self *Logic) Run() {
	//
}

// Offline 模式下中途终止任务
func (self *Logic) Stop() {
}

// 检查任务是否正在运行
func (self *Logic) IsRunning() bool {
	return self.status == status.RUN
}

// 检查任务是否处于暂停状态
func (self *Logic) IsPause() bool {
	return self.status == status.PAUSE
}

// 检查任务是否已经终止
func (self *Logic) IsStopped() bool {
	return self.status == status.STOPPED
}

// SpiderPrepare()必须在设置全局运行参数之后，Run()的前一刻执行
// original为spider包中未有过赋值操作的原始蜘蛛种类
// 已被显式赋值过的spider将不再重新分配Keyin
// client模式下不调用该方法
func (self *Logic) SpiderPrepare(original []*spider.Spider) App {
	return self
}

// 获取全部蜘蛛种类
func (self *Logic) GetSpiderLib() []*spider.Spider {
	return self.SpiderSpecies.Get()
}

// 获取全部输出方式
func (self *Logic) GetOutputLib() []string {
	return make([]string, 10)
	// return collector.DataOutputLib
}

// 切换运行模式时使用
// func (self *Logic) ReInit(mode int, port int, master string, w ...io.Writer) App {
// }

/*
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
*/
