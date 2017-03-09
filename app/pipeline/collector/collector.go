// 结果收集与输出
package collector

import (
	"fmt"
	"github.com/zhenwusw/logan/app/pipeline/collector/data"
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/runtime/cache"
	"sync"
	"sync/atomic"
)

type Collector struct {
	*spider.Spider
	DataChan   chan data.DataCell
	FileChan   chan data.FileCell
	dataDocker []data.DataCell
	outType    string
	// size    [2]uint64
	dataBatch   uint64
	fileBatch   uint64
	wait        sync.WaitGroup
	sum         [4]uint64
	dataSumLock sync.RWMutex
	fileSumLock sync.RWMutex
}

func NewCollector(sp *spider.Spider) *Collector {
	var self = &Collector{}
	self.Spider = sp
	self.outType = cache.Task.OutType
	if cache.Task.DockerCap < 1 {
		cache.Task.DockerCap = 1
	}
	self.DataChan = make(chan data.DataCell, cache.Task.DockerCap)
	self.FileChan = make(chan data.FileCell, cache.Task.DockerCap)
	self.dataDocker = make([]data.DataCell, cache.Task.DockerCap)
	self.sum = [4]uint64{}
	// self.size = [2]uint64{
	self.dataBatch = 0
	self.fileBatch = 0
	return self
}

func (self *Collector) CollectData(dataCell data.DataCell) error {
	var err error
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("输出协程终止")
		}
	}()
	self.DataChan <- dataCell
	return err
}

func (self *Collector) CollectFile(fileCell data.FileCell) error {
	var err error
	defer func() {
		if recover() != nil {
			err = fmt.Errorf("输出协程终止")
		}
	}()
	self.FileChan <- fileCell
	return err
}

func (self *Collector) Stop() {
}

// 启动数据收集/输出管道
func (self *Collector) Start() {
	// 启动输出协程
	go func() {
		dataStop := make(chan bool)
		fileStop := make(chan bool)

		go func() {
			defer func() {
				recover()
			}()
			for data := range self.DataChan {
				// 缓存分批数据
				self.dataDocker = append(self.dataDocker, data)
				// 未达到设定的分批量时继续收集收据
				if len(self.dataDocker) < cache.Task.DockerCap {
					continue
				}
				// 执行输出
				self.dataBatch++
				self.outputData()
			}
			// 将剩余收集到但未输出的数据输出
			self.dataBatch++
			self.outputData()
			close(dataStop)
		}()

		go func() {
			defer func() {
				recover()
			}()
			// 只有当收到退出通知并且通道内无数据时，才退出循环
			for file := range self.FileChan {
				atomic.AddUint64(&self.fileBatch, 1)
				self.wait.Add(1)
				go self.outputFile(file)
			}
			close(fileStop)
		}()

		<-dataStop
		<-fileStop

		// 等待所有输出完成
		self.wait.Wait()
		// 返回报告
		self.Report()
	}()
}

// 返回报告
func (self *Collector) Report() {
	fmt.Printf("... %v", "collector#Report()")
}

/*
func (self *Collector) resetDataDocker() {
}

// 获取文本数据总量
func (self *Collector) dataSum() uint64 {
}

// 更新文本数据总量
func (self *Collector) addDataSum(add uint64) {
}

// 获取文件数据总量
func (self *Collector) fileSum() uint64 {
}

// 更新文件数据总量
func (self *Collector) addFileSum(add uint64) {
}
*/
