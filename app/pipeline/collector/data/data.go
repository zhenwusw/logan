package data

import "sync"

/*
  DataCell:
    RuleName
    Data
    Url
    ParentUrl
    DownloadTime
 */

/* FileCell:
     RuleName
     Name
     Bytes
 */
type (
	// 数据存储单元
	DataCell map[string]interface{}
	// 文件存储单元
	// FileCell 存储的完整文件名为：file/"Dir"/"RuleName"/"time"/"Name"
	FileCell map[string]interface{}
)

var (
	dataCellPool = &sync.Pool{
		New: func() interface{} {
			return DataCell{}
		},
	}
	fileCellPool = &sync.Pool{
		New: func() interface{} {
			return FileCell{}
		},
	}
)

func GetDataCell(ruleName string, data map[string]interface{}, url string, parentUrl string, downloadTime string) DataCell {
	cell := dataCellPool.Get().(DataCell)
	cell["RuleName"] = ruleName
	cell["Data"] = data    // 数据存储，key必须与Rule的Fields保持一致
	cell["Url"] = url
	cell["ParentUrl"] = parentUrl
	cell["DownloadTime"] = downloadTime
	return cell
}

func GetFileCell(ruleName, name string, bytes []byte) FileCell {
	cell := fileCellPool.Get().(FileCell)
	cell["RuleName"] = ruleName
	cell["Name"] = name
	cell["Bytes"] = bytes
	return cell
}

func PutDataCell(cell DataCell) {
	cell["RuleName"] = nil
	cell["Data"] = nil
	cell["Url"] = nil
	cell["ParentUrl"] = nil
	cell["DownloadTime"] = nil
	dataCellPool.Put(cell)
}

func PutFileCell(cell FileCell) {
	cell["RuleName"] = nil
	cell["Name"] = nil
	cell["Bytes"] = nil
	fileCellPool.Put(cell)
}
