// 基本业务执行顺序依次为:
// New() --> [SetLog(io.Writer)-->]Init()-->SpiderPrepare()-->Run()
package app

import (
	"fmt"
	"github.com/zhenwusw/logan/app/crawler"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/runtime/cache"
	"github.com/zhenwusw/logan/runtime/status"
	"io"
	"time"
	"strings"
	"reflect"
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
		GetAppConf(k ...string) interface{} // 获取全局参数
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
		crawler.SpiderQueue // 当前任务的蜘蛛队列
		// *distribute.TaskJar
		crawler.CrawlerPool // 爬行回收池
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
		SpiderQueue:   crawler.NewSpiderQueue(),
		CrawlerPool:   crawler.NewCrawlerPool(),
	}
}

// 使用App前必须进行Init初始化
func (self *Logic) Init(mode int, port int, master string, w ...io.Writer) App {
	// 配置 AppConf{ Mode, Port, Master}
	self.AppConf.Mode = mode
	self.SpiderQueue = crawler.NewSpiderQueue()
	self.CrawlerPool = crawler.NewCrawlerPool()

	switch self.AppConf.Mode {
	case status.OFFLINE:
		// log something
		return self
	default:
		// log something
		return self
	}
	return self
}

// 运行任务
func (self *Logic) Run() {
	fmt.Printf("...... %v \n", "app#Run()")

	// 开启报告
	// 客户端模式
	self.finish = make(chan bool)
	// self.finishOnce = sync.Once{}
	// self.sum[0]
	// self.setStatus(status.RUN)
	// defer self.setStatus(status.STOPPED)

	// 任务执行
	switch self.AppConf.Mode {
	case status.OFFLINE:
		self.offline()
	case status.SERVER:
		self.server()
	case status.CLIENT:
		self.client()
	default:
		return
	}
	<-self.finish
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
	self.SpiderQueue.Reset()
	// 遍历任务
	for _, sp := range original {
		spcopy := sp.Copy()
		spcopy.SetPausetime(self.AppConf.Pausetime)
		if spcopy.GetLimit() == spider.LIMIT {
			spcopy.SetLimit(self.AppConf.Limit)
		} else {
			spcopy.SetLimit(-1 * self.AppConf.Limit)
		}
		self.SpiderQueue.Add(spcopy)
	}

	// 遍历自定义配置
	// self.SpiderQueue.AddKeyins(self.AppConf.Keyins)
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

// 获取全局参数
func (self *Logic) GetAppConf(k ...string) interface{} {
	defer func() {
		if err := recover(); err != nil {
			// logs.Log.Error()
			fmt.Printf("%v", err)
		}
	}()
	if len(k) == 0 {
		return self.AppConf
	}
	key := strings.Title(k[0])
	acv := reflect.ValueOf(self.AppConf).Elem()
	return acv.FieldByName(key).Interface()
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

// 离线模式运行
func (self *Logic) offline() {
	self.exec()
}

func (self *Logic) server() {
}

func (self *Logic) client() {
}

// 开始执行任务
func (self *Logic) exec() {
	fmt.Printf("...... %v \n", "app#exec()")

	count := 1
	// count := self.SpiderQueue.Len()
	// cache.ResetPageCount()
	// 刷新输出方式的状态
	// pipeline.RefreshOutput()

	// 初始化资源队列
	// scheduler.Init()

	// 设置爬虫队列
	crawlerCap := self.CrawlerPool.Reset(count)
	fmt.Printf(" *     采集引擎池容量为 %v\n", crawlerCap)

	// 开始计时
	// cache.StartTime = time.Now()

	// 根据模式选择合理的并发
	if self.AppConf.Mode == status.OFFLINE {
		// 可控制执行状态
		go self.goRun(count)
	} else {
		// 保证接收服务端任务的同步
		// self.goRun(count)
	}
}

// 任务执行
func (self *Logic) goRun(count int) {
	fmt.Printf("...... %v \n", "app#goRun()")

	// 执行任务
	var i int
	for i = 0; i < count && self.status != status.STOP; i++ {
	pause:
		if self.IsPause() {
			time.Sleep(time.Second)
			goto pause
		}
		// 从爬行队列取出空闲蜘蛛，并发执行
		c := self.CrawlerPool.Use()
		if c != nil {
			go func(i int, c crawler.Crawler) {
				// 执行并返回结果消息
				c.Init(self.SpiderQueue.GetByIndex(i)).Run()
				// self.RWMutex.RLock()
				if self.status != status.STOP {
					self.CrawlerPool.Free(c)
				}
				// self.RWMutext.RUnlock()
				// 任务结束后回收改蜘蛛
			}(i, c)
		}
	}
	// 监控结束任务
	for ii := 0; ii < i; ii++ {
		s := <-cache.ReportChan
		if (s.DataNum == 0) && (s.FileNum == 0) {
			// ...
		}
		// self.sum[0] += s.DataNum
		// self.sum[1] += s.FileNum
	}

	// 单机模式并发运行，需要标记任务结束
	if self.AppConf.Mode == status.OFFLINE {
		// self.LogRest()
		// self.finishOnce.Do(func() { close(self.finish) })
	}
}
