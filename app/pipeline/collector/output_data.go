package collector

var (
	// 全局支持的输出方式
	DataOutput = make(map[string]func(self *Collector) error)
	// 全局支持的文本数据输出方式名称列表
	DataOutputLib []string
)

// 文本数据输出
func (self *Collector) outputData() {
}

