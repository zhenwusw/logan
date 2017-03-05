package crawler

import (
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/app/downloader"
	//"github.com/zhenwusw/logan/app/pipeline"
)

// 采集引擎
type (
	Crawler interface {
		Init(*spider.Spider) Crawler // 初始化采集引擎
		Run()                        // 运行任务
		Stop()                       // 主动终止
		CanStop() bool               // 是否终止
		GetId() int                  // 获取引擎ID
	}

	crawler struct {
		*spider.Spider        // 执行的采集规则
		downloader.Downloader //全局公用的下载器
		// pipeline.Pipeline     //结果收集与输出管道
		id    int      //引擎ID
		pause [2]int64 //[请求间隔的最短时长,请求间隔的增幅时长]
	}
)

/*
func New() Crawler {
}

func (self *crawler) Init(sp *spider.Spider) Crawler {
}

// 任务执行入口
func (self *crawler) Run() {
}

// 主动终止
func (self *crawler) Stop() {
}

func (self *crawler) run() {
}

// core processer
func (self *crawler) Process(req *request.Request) {
}

// 常用基础方法
func (self *crawler) sleep() {
}

// 从调度读取一个请求
func (self *crawler) GetOne() *request.Request {
}

//从调度使用一个资源空位
func (self *crawler) UseOne() {
}

//从调度释放一个资源空位
func (self *crawler) FreeOne() {
}

func (self *crawler) SetId(id int) {
}

func (self *crawler) GetId() int {
}
*/
