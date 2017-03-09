package collector

import (
	"fmt"
	"github.com/zhenwusw/logan/app/pipeline/collector/data"
)

// 文件输出
func (self *Collector) outputFile(file data.FileCell) {
	fmt.Printf("... %v", "collector#outputFile()")
	/*
		// 复用 FileCell
		defer func() {
			data.PutFileCell(file)
			self.wait.Done()
		}()

		// 路径: file/"RuleName"/"time"/"Name"
		p, n := filepath.Split(filepath.Clean(file["Name"].(string)))
		// dir := filepath.Join(config.FILE_DIR, util.FileNameReplace(self.namespace())+"__"+cache.StartTime.Format("2006年01月02日 15时04分05秒"), p)
		dir := filepath.Join(config.FILE)
	*/
}
