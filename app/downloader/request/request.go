package request

import (
	"sync"
	"net/http"
	"time"
)

// Request represents object waiting for being crawled.
type Request struct {
	Spider        string
	Url           string
	Rule          string
	Method        string
	Header        http.Header
	EnableCookie  bool
	PostData      string
	DialTimeout   time.Duration
	ConnTimeout   time.Duration
	CryTimes      int
	RetryPause    time.Duration
	redirectTimes int
	Temp          Temp
	TempIsJson    map[string]bool
	Priority      int
	Reloadable    bool

	// Surfer下载器内核ID
	// 0高并发下载器，各种控制功能齐全
	// 1为PhantomJS下载器,特点破防力强，速度慢，低并发
	DownloaderID int

	proxy  string
	unique string
	lock   sync.RWMutex
}

const (
	DefaultDialTimeout = 2 * time.Minute // 默认请求服务器超时
	DefaultConnTimeout = 2 * time.Minute // 默认下载超时
	DefaultTryTimes    = 3               // 默认最大下载数
	DefaultRetryPause  = 2 * time.Second // 默认重新下载前停顿时长
)

const (
	SURF_ID = 0    // 默认的 surf下载内核 (Go原生), 此值不可改动
	PHANTOM_ID = 1 // 备用的 phantomjs 下载内核，一般不使用（效率差，头信息支持不完善）
)


func (self *Request) Prepare() error {
}

// 反序列化
func UnSerialize(s string) (*Request, error) {
}

// 序列化
func (self *Request) Serialize() string {
}

// 请求的唯一识别码
func (self *Request) Unique() string {
}

// 获取副本
func (self *Request) Copy() *Request {
}

// 获取 Url
func (self *Request) GetUrl() string {
}

// 获取Http请求的方法名称
func (self *Request) GetMethod() string {
}

// 设定Http请求方法的类型
func (self *Request) SetMethod(method string) *Request {
}

func (self *Request) SetUrl(url string) *Request {
	self.Url = url
	return self
}

func (self *Request) GetReferer() string {
	return self.Header.Get("Referer")
}

func (self *Request) SetReferer(referer string) *Request {
	self.Header.Set("Referer", referer)
	return self
}

func (self *Request) GetPostData() string {
	return self.PostData
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

func (self *Request) GetEnableCookie() bool {
	return self.EnableCookie
}

func (self *Request) SetEnableCookie(enableCookie bool) *Request {
	self.EnableCookie = enableCookie
	return self
}

func (self *Request) GetCookies() string {
	return self.Header.Get("Cookie")
}

func (self *Request) SetCookies(cookie string) *Request {
	self.Header.Set("Cookie", cookie)
	return self
}

func (self *Request) GetDialTimeout() time.Duration {
	return self.DialTimeout
}

func (self *Request) GetConnTimeout() time.Duration {
	return self.ConnTimeout
}

func (self *Request) GetTryTimes() int {
	return self.TryTimes
}

func (self *Request) GetRetryPause() time.Duration {
	return self.RetryPause
}

func (self *Request) GetProxy() string {
	return self.proxy
}

func (self *Request) SetProxy(proxy string) *Request {
	self.proxy = proxy
	return self
}

func (self *Request) GetRedirectTimes() int {
	return self.RedirectTimes
}

func (self *Request) GetRuleName() string {
	return self.Rule
}

func (self *Request) SetRuleName(ruleName string) *Request {
	self.Rule = ruleName
	return self
}

func (self *Request) GetSpiderName() string {
	return self.Spider
}

func (self *Request) SetSpiderName(spiderName string) *Request {
	self.Spider = spiderName
	return self
}

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

func (self *Request) GetPriority() int {
	return self.Priority
}

func (self *Request) SetPriority(priority int) *Request {
	self.Priority = priority
	return self
}

func (self *Request) GetDownloaderID() int {
	return self.DownloaderID
}

func (self *Request) SetDownloaderID(id int) *Request {
	self.DownloaderID = id
	return self
}

func (self *Request) MarchalJSON() ([]byte, error) {
}















