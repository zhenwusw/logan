package crawler

import (
	. "github.com/zhenwusw/logan/app/spider"
)

// 采集引擎中规则队列
type (
	SpiderQueue interface {
		Reset() // 重置清空队列
		Add(*Spider)
		// AddAll([]*Spider)
		GetByIndex(int) *Spider
		// GetByName(string) *Spider
		// GetAll() []*Spider
		Len() int // 返回队列长度
	}
	sq struct {
		list []*Spider
	}
)

func NewSpiderQueue() SpiderQueue {
	return &sq{
		list: []*Spider{},
	}
}

func (self *sq) Reset() {
	self.list = []*Spider{}
}

func (self *sq) Add(sp *Spider) {
	sp.SetId(self.Len())
	self.list = append(self.list, sp)
}

func (self *sq) GetByIndex(idx int) *Spider {
	return self.list[idx]
}

func (self *sq) Len() int {
	return len(self.list)
}