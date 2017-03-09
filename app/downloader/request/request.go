package request

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Request represents object waiting for being crawled.
type Request struct {
	Spider string      // 规则名，自动设置，禁止人为填写
	Url    string      // 目标URL，必须设置
	Rule   string      // 用于解析响应的规则节点名，必须设置
	Method string      // GET POST POST-M HEAD
	Header http.Header // 请求头信息
	EnableCookie  bool
	// PostData      string
	// DialTimeout   time.Duration
	// ConnTimeout   time.Duration
	// TryTimes      int
	// RetryPause    time.Duration
	// redirectTimes int
	// Temp       Temp
	// TempIsJson map[string]bool
	Priority int
	// Reloadable    bool

	// Surfer下载器内核ID
	// 0 高并发下载器，各种控制功能齐全
	// 1 为PhantomJS下载器,特点破防力强，速度慢，低并发
	DownloaderID int

	// proxy  string
	// unique string
	// lock   sync.RWMutex
}

/*
const (
	DefaultDialTimeout = 2 * time.Minute // 默认请求服务器超时
	DefaultConnTimeout = 2 * time.Minute // 默认下载超时
	DefaultTryTimes    = 3               // 默认最大下载数
	DefaultRetryPause  = 2 * time.Second // 默认重新下载前停顿时长
)*/

const (
	SURF_ID    = 0 // 默认的 surf下载内核 (Go原生), 此值不可改动
	PHANTOM_ID = 1 // 备用的 phantomjs 下载内核，一般不使用（效率差，头信息支持不完善）
)

/*
发送请求前的准备工作，设置一系列默认值:

	Request.Url与Request.Rule必须设置
	Request.Spider无需手动设置(由系统自动设置)

以下字段有默认值，可不设置:

	Request.Method默认为GET方法;
	Request.DownloaderID指定下载器ID，0为默认的Surf高并发下载器，功能完备，1为PhantomJS下载器，特点破防力强，速度慢，低并发。
*/

func (self *Request) Prepare() error {
	// 确保url正确，且和Response中Url字符串相等
	URL, err := url.Parse(self.Url)
	if err != nil {
		return err
	} else {
		self.Url = URL.String()
	}

	if self.Method == "" {
		self.Method = "GET"
	} else {
		self.Method = strings.ToUpper(self.Method)
	}

	if self.Header == nil {
		self.Header = make(http.Header)
	}
	// DialTimeout
	// ConnTimeout
	// TryTimes
	// RetryPause
	// Priority
	// DownloaderID
	// TempIsJson
	// Temp
	return nil
}

// 反序列化
//func UnSerialize(s string) (*Request, error) {
//	req := new(Request)
//	return req, json.Unmarshal([]byte(s), req)
//}

// 序列化
//func (self *Request) Serialize() string {
//	for k, v := range self.Temp {
//	}
//	b, _ := json.Marshal()
//}

// 请求的唯一识别码
// func (self *Request) Unique() string {
// }

// 获取副本
// func (self *Request) Copy() *Request {
// }

func (self *Request) SetUrl(url string) *Request {
	self.Url = url
	return self
}

// 获取 Url
func (self *Request) GetUrl() string {
	return self.Url
}

// 设定Http请求方法的类型
func (self *Request) SetMethod(method string) *Request {
	self.Method = strings.ToUpper(method)
	return self
}

// 获取Http请求的方法名称
func (self *Request) GetMethod() string {
	return self.Method
}

func (self *Request) GetHeader() http.Header {
	return self.Header
}

func (self *Request) SetHeader(key, value string) *Request {
	self.Header.Set(key, value)
	return self
}

func (self *Request) AddHeader(key, value string) *Request {
	self.Header.Add(key, value)
	return self
}

func (self *Request) GetReferer() string {
	return self.Header.Get("Referer")
}

func (self *Request) SetReferer(referer string) *Request {
	self.Header.Set("Referer", referer)
	return self
}

func (self *Request) GetRuleName() string {
	return self.Rule
}

func (self *Request) GetDownloaderID() int {
	return self.DownloaderID
}

func (self *Request) GetPostData() string {
	return "postdata"
	// return self.PostData
}

func (self *Request) GetEnableCookie() bool {
	return true
	// return self.EnableCookie
}

func (self *Request) GetDialTimeout() time.Duration {
	return 2 * time.Minute
	// return self.DialTimeout
}

func (self *Request) GetConnTimeout() time.Duration {
	return 2 * time.Minute
	// return self.ConnTimeout
}

func (self *Request) GetTryTimes() int {
	return 3
	// return self.TryTimes
}

func (self *Request) GetRetryPause() time.Duration {
	return 10 * time.Second
	// return self.RetryPause
}

func (self *Request) GetProxy() string {
	return ""
	// return self.proxy
}

func (self *Request) GetRedirectTimes() int {
	return 3
	// return self.RedirectTimes
}

func (self *Request) GetPriority() int {
	return self.Priority
}

func (self *Request) SetPriority(priority int) *Request {
	self.Priority = priority
	return self
}

func (self *Request) SetSpiderName(spiderName string) *Request {
	self.Spider = spiderName
	return self
}

func (self *Request) SetEnableCookie(enableCookie bool) *Request {
	self.EnableCookie = enableCookie
	return self
}

/*

func (self *Request) GetCookies() string {
	return self.Header.Get("Cookie")
}

func (self *Request) SetCookies(cookie string) *Request {
	self.Header.Set("Cookie", cookie)
	return self
}

func (self *Request) SetProxy(proxy string) *Request {
	self.proxy = proxy
	return self
}

func (self *Request) SetRuleName(ruleName string) *Request {
	self.Rule = ruleName
	return self
}

func (self *Request) GetSpiderName() string {
	return self.Spider
}
*/

/*
func (self *Request) IsReloadable() bool {
	return self.Reloadable
}

func (self *Request) SetReloadable(can bool) *Request {
	self.Reloadable = can
	return self
}

// 获取临时缓存数据
// defaultValue 不能为 interface{}(nil)
func (self *Request) GetTemp() interface{} {
}

func (self *Request) GetTemps() Temp {
}

func (self *Request) SetTemp() *Request {
}

func (self *Request) SetTemps() *Request {
}


func (self *Request) SetDownloaderID(id int) *Request {
	self.DownloaderID = id
	return self
}

func (self *Request) MarchalJSON() ([]byte, error) {
}
*/
