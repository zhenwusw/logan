package spider

import (
	"sync"
	"github.com/zhenwusw/logan/app/downloader/request"
	"time"
	"github.com/zhenwusw/logan/logs"
)

type Context struct {
}

var (
	contextPool = &sync.Pool{
	}
)

//**************************************** 初始化 *******************************************\\

func GetContext(sp *Spider, req *request.Request) *Context {
}

func PutContext(ctx *Context) {
}

func (self *Context) SetResponse() *Context {
}

// 标记下载错误
func (self *Context) SetError(err error) {
}

//**************************************** Set与Exec类公开方法 *******************************************\\
// 生成并添加请求至队列
// Request.Url 与 Request.Rule
// Request.Spider
// Request.EnableCookie
// 一下字段有默认值：
// Request.Method
// Request.DialTimeout
// Request.ConnTimeout
// Request.TryTimes
// Request.RedirectTimes
// Request.RetryPause
// Request.DownloaderID
// 默认自动填补 Referer
func (self *Context) AddQueue(req *request.Request) *Context {
}

// 用于动态规则添加请求。
func (self *Context) JsAddQueue(jreq map[string]interface{}) *Context {
}

// 输出文本结果。
func (self *Context) Output(item interface{}, ruleName ...string) {
}

// 输出文件
func (self *Context) FileOutput(name ...string) {
}

// 生成文本结果
func (self *Context) CreateItem() map[string]interface{} {
}

// 在请求中保存临时数据。
func (self *Context)  SetTemp() *Context {
}

func (self *Context) SetUrl(url string) *Context {
}

func (self *Context) SetReferer(referer string) *Context {
	// self.Request.Header.Set("Referer", referer)
	// return self
}

// 为指定Rule动态追加结果字段名，并获取索引位置，
func (self *Context) UpsertItemField(field string, ruleName ...string) (index int) {
}

// 调用指定 Rule 下辅助函数 AidFunc()
// 用 ruleName指定匹配的 AidFunc, 为空时默认当前规则
func (self *Context) Aid() interface{} {
}

// 解析响应流
// 用 ruleName 指定匹配的 ParseFunc 字段，为空时默认调用 Root()
func (self *Context) Parse(ruleName ...string) *Context {

}

// 设定自定义配置
func (self *Context) SetKeyin(keyin string) *Context {
}

// 设置采集上限
func (self *Context) SetLimit(max int) *Context {
}

// 自定义暂停区间
func (self *Context) SetPausetime() *Context {
}

// 设置定时器
func (self *Context) SetTimer(id string, tol time.Duration, bell *Bell) bool {
}

// 启动定时器，并获取定时器是否可以继续使用
func (self *Context) RunTimer(id string) bool {
}

// 重置下载的文本内容
func (self *Context) ResetText(body string) *Context {
}

//**************************************** Get 类公开方法 *******************************************\\

// 获取下载错误
func (self *Context) GetError() error {
}

// 获取日志接口实例
func (self *Context) Log() logs.Logs {
}

// 获取蜘蛛名称。
func (self *Context) GetSpider() *Spider {
}

// 获取响应流。

// 获取响应状态码。

// 获取原始请求。

// 获得一个原始请求的副本。

// 获取结果字段名列表。

// 由索引下标获取结果字段名，不存在时获取空字符串，
// 若ruleName为空，默认为当前规则。


// 由结果字段名获取索引下标，不存在时索引为-1，
// 若ruleName为空，默认为当前规则。


func (self *Context) PullItems() (ds []data.DataCell) {
}


func (self *Context) PullFiles() (fs []data.FileCell) {
}











