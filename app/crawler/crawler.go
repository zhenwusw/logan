package crawler

import "github.com/zhenwusw/logan/app/spider"

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
		*spider.Spider               // 执行的采集规则
		downl
	}
)


func New() Crawler {
}

func (self *crawler) Init(sp *spider.Spider) Crawler {
}