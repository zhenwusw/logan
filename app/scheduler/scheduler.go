package scheduler

import (
	"github.com/henrylee2cn/pholcus/app/aid/proxy"
	"github.com/zhenwusw/logan/runtime/cache"
	"github.com/zhenwusw/logan/runtime/status"
	"sync"
)

// 调度器
type scheduler struct {
	status       int          // 运行状态
	count        chan bool    // 总并发量计数
	useProxy     bool         // 标记是否使用代理 IP
	proxy        *proxy.Proxy // 全局代理IP
	matrices     []*Matrix    // Spider 实例的请求矩阵列表
	sync.RWMutex              // 全局读写锁
}

// 定义全局调度
var sdl = &scheduler{
	status: status.RUN,
	count:  make(chan bool, cache.Task.ThreadNum),
	proxy:  proxy.New(),
}

func Init() {
	for sdl.proxy == nil {
	}
	sdl.matrices = []*Matrix{}
	sdl.count = make(chan bool, cache.Task.ThreadNum)

	//if cache.Task.ProxyMinute > 0 {
	//
	//} else {
	//	sdl.useProxy = false
	//	logs.Log.Informational()
	//}
	sdl.status = status.RUN
}

// 注册资源队列
func AddMatrix(spiderName, spiderSubName string, maxPage int64) *Matrix {
	matrix := newMatrix(spiderName, spiderSubName, maxPage)
	sdl.RLock()
	defer sdl.RUnlock()
	sdl.matrices = append(sdl.matrices, matrix)
	return matrix
}

// 暂停/恢复所有爬行任务
func PauseRecover() {
}

// 终止任务
func Stop() {
}

// 每个 spider 实例分配到的平均资源量
func (self *scheduler) avgRes() int32 {
	avg := int32(cap(sdl.count) / len(sdl.matrices))
	if avg == 0 {
		avg = 1
	}
	return avg
}

func (self *scheduler) checkStatus(s int) bool {
	self.RLock()
	b := self.status == s
	self.RUnlock()
	return b
}
