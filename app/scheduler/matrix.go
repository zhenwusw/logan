package scheduler

import "github.com/zhenwusw/logan/app/downloader/request"

// 一个 Spider 实例的请求矩阵
type Matrix struct {
}

func newMatrix() *Matrix {
}

// 添加请求到队列，并发安全
func (self *Matrix) Push() {
}

// 从队列去除请求，不存在时返回 nil, 并发安全
func (self *Matrix) Pull() (req *request.Request) {
}

func (self *Matrix) Use() {
}

func (self *Matrix) Free() {
}

// 返回是否作为新的失败请求被添加至队列尾部
func (self *Matrix) DoHistory() bool {
}

func (self *Matrix) CanStop() bool {
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
}

func (self *Matrix) hasHistory() bool {
}

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
