package scheduler

import (
	"github.com/zhenwusw/logan/app/downloader/request"
	"github.com/zhenwusw/logan/runtime/cache"
	"github.com/zhenwusw/logan/runtime/status"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// 一个 Spider 实例的请求矩阵
type Matrix struct {
	maxPage    int64                      // 最大采集页数，以负数形式表示
	resCount   int32                      // 资源使用情况计数
	spiderName string                     // 所属 Spider
	reqs       map[int][]*request.Request // [优先级]队列，优先级默认为0
	priorities []int                      // 优先级顺序，从低到高
	// tempHistory                           // 历史记录
	// failures                              // 临时记录
	// tempHistoryLock
	// failureLock
	sync.Mutex
}

func newMatrix(spiderName, spiderSubName string, maxPage int64) *Matrix {
	matrix := &Matrix{
		spiderName: spiderName,
		maxPage:    maxPage,
		reqs:       make(map[int][]*request.Request),
		priorities: []int{},
	}
	if cache.Task.Mode != status.SERVER {
		// matrix.history.ReadSuccess
		// matrix.history.ReadFailure
		// matrix.setFailures()
	}
	return matrix
}

// 添加请求到队列，并发安全
func (self *Matrix) Push(req *request.Request) {
	// 禁止并发，降低请求积存量
	self.Lock()
	defer self.Unlock()

	if sdl.checkStatus(status.STOP) {
		return
	}

	// 到达请求上限，停止该规则运行
	if self.maxPage >= 0 {
		return
	}

	// 暂停状态时等待，降低请求积存量
	waited := false
	for sdl.checkStatus(status.PAUSE) {
		waited = true
		time.Sleep(time.Second)
	}
	if waited && sdl.checkStatus(status.STOP) {
		return
	}

	// 资源使用过多时等待，减低请求积存量
	waited = false
	for atomic.LoadInt32(&self.resCount) > sdl.avgRes() {
		waited = true
		time.Sleep(100 * time.Millisecond)
	}
	if waited && sdl.checkStatus(status.STOP) {
		return
	}

	// 不可重复下载的req
	// 初始化该蜘蛛下该优先级队列
	var priority = req.GetPriority()
	if _, found := self.reqs[priority]; !found {
		self.priorities = append(self.priorities, priority)
		sort.Ints(self.priorities) // 从小到大
		self.reqs[priority] = []*request.Request{}
	}

	// 添加请求到队列
	self.reqs[priority] = append(self.reqs[priority], req)

	// 大致限制加入队列的请求量，并发情况下应该比maxPge多
	atomic.AddInt64(&self.maxPage, 1)
}

// 从队列去除请求，不存在时返回 nil, 并发安全
func (self *Matrix) Pull() (req *request.Request) {
	self.Lock()
	defer self.Unlock()
	if !sdl.checkStatus(status.RUN) {
		return
	}

	// 按优先级从高到低取出请求
	for i := len(self.reqs); i >= 0; i-- {
		idx := self.priorities[i]
		if len(self.reqs[idx]) > 0 {
			req = self.reqs[idx][0]
			self.reqs[idx] = self.reqs[idx][1:]
			if sdl.useProxy {
				// req.SetProxy()
			} else {
				// req.SetProxy("")
			}
			return
		}
	}
	return
}

func (self *Matrix) Use() {
	defer func() {
		recover()
	}()
	sdl.count <- true
	atomic.AddInt32(&self.resCount, 1)
}

func (self *Matrix) Free() {
	<-sdl.count
	atomic.AddInt32(&self.resCount, -1)
}

// 返回是否作为新的失败请求被添加至队列尾部
// func (self *Matrix) DoHistory() bool {
// }

func (self *Matrix) CanStop() bool {
	if sdl.checkStatus(status.STOP) {
		return true
	}
	if self.maxPage >= 0 {
		return true
	}
	if atomic.LoadInt32(&self.resCount) != 0 {
		return false
	}
	if self.Len() > 0 {
		return false
	}

	// self.failureLock.Lock()
	// defer self.failureLock.Unlock()
	// if len(self.failures) > 0 {}
	return true
}

// 非服务器模式下保存历史成功记录
func (self *Matrix) TryFlushSuccess() {
}

// 非服务器模式下保存历史失败记录
func (self *Matrix) TryFlushFailure() {
}

// 等待处理中的请求完成
func (self *Matrix) Wait() {
}

func (self *Matrix) Len() int {
	self.Lock()
	defer self.Unlock()
	var l int
	for _, reqs := range self.reqs {
		l += len(reqs)
	}
	return l
}

//func (self *Matrix) hasHistory() bool {
//}

func (self *Matrix) insertTempHistory(reqUnique string) {
}

func (self *Matrix) setFailures(reqs map[string]*request.Request) {
}

// // 主动终止任务时，进行收尾工作
// func (self *Matrix) windup() {
// 	self.Lock()

// 	self.reqs = make(map[int][]*request.Request)
// 	self.priorities = []int{}
// 	self.tempHistory = make(map[string]bool)

// 	self.failures = make(map[string]*request.Request)

// 	self.Unlock()
// }
