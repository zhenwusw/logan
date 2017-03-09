package crawler

import (
	"github.com/zhenwusw/logan/app/downloader"
	"github.com/zhenwusw/logan/app/pipeline"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/app/downloader/request"
	"time"
	"fmt"
	"math/rand"
)

// 采集引擎
type (
	Crawler interface {
		Init(*spider.Spider) Crawler // 初始化采集引擎
		Run()                        // 运行任务
		Stop()         // 主动终止
		CanStop() bool // 是否终止
		// GetId() int                  // 获取引擎ID
	}

	crawler struct {
		*spider.Spider                 // 执行的采集规则
		downloader.Downloader          //全局公用的下载器
		pipeline.Pipeline              //结果收集与输出管道
		id                    int      //引擎ID
		pause                 [2]int64 //[请求间隔的最短时长,请求间隔的增幅时长]
	}
)

func New(id int) Crawler {
	return &crawler{
		id:         id,
		Downloader: downloader.SurferDownloader,
	}
}

func (self *crawler) Init(sp *spider.Spider) Crawler {
	self.Spider = sp.ReqmatrixInit()
	self.Pipeline = pipeline.New(sp)
	self.pause[0] = sp.Pausetime / 2
	if self.pause[0] > 0 {
		self.pause[1] = self.pause[0] * 3
	} else {
		self.pause[1] = 1
	}
	return self
}

// 任务执行入口
func (self *crawler) Run() {
	fmt.Printf("...... %v \n", "crawler#Run()")

	// 预先启动数据收集/输出管道
	self.Pipeline.Start()

	// 运行处理协程
	c := make(chan bool)
	go func() {
		self.run()
		close(c)
	}()

	// 启动任务
	self.Spider.Start()

	<-c // 等待处理协程退出

	// 停止数据收集/输出管道
	self.Pipeline.Stop()
}

func (self *crawler) run() {
	fmt.Printf("...... %v\n", "crawler#run()")
	for {
		fmt.Printf("...... %v\n", "hello")
		// 队列中取出一条请求并处理
		req := self.GetOne()
		if req == nil {
			// 停止任务
			if self.Spider.CanStop() {
				break
			}
			time.Sleep(20 * time.Millisecond)
			continue
		}

		// 执行请求
		self.UseOne()
		go func(req *request.Request) {
			defer func() {
				self.FreeOne()
			}()
			// logs.Log.Debug()
			fmt.Printf(" *     Start: %v", req.GetUrl())
			self.Process(req)
		}(req)

		// 随机等待
		self.sleep()
	}

	// 等待处理中的任务完成
	self.Spider.Defer()
}


// 主动终止
func (self *crawler) Stop() {
	self.Spider.Stop()
	self.Pipeline.Stop()
}

// 从调度读取一个请求
func (self *crawler) GetOne() *request.Request {
	return self.Spider.RequestPull()
}

//从调度使用一个资源空位
func (self *crawler) UseOne() {
	self.Spider.RequestUse()
}

//从调度释放一个资源空位
func (self *crawler) FreeOne() {
	self.Spider.RequestFree()
}

// 常用基础方法
func (self *crawler) sleep() {
	sleeptime := self.pause[0] + rand.Int63n(self.pause[1])
	time.Sleep(time.Duration(sleeptime) * time.Millisecond)
}

// core processer
func (self *crawler) Process(req *request.Request) {
	fmt.Printf("... crawler#Process()")
}


/*
func (self *crawler) SetId(id int) {
}

func (self *crawler) GetId() int {
}
*/
