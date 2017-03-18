package spider

import (
	"github.com/zhenwusw/logan/app/downloader/request"
	"github.com/zhenwusw/logan/app/pipeline/collector/data"
	"github.com/zhenwusw/logan/common/goquery"
	"github.com/zhenwusw/logan/logs"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"strings"
	"sync"
	"fmt"
)

type Context struct {
	spider   *Spider           // 规则
	Request  *request.Request  // 原始请求
	Response *http.Response    //
	text     []byte            // 下载内容 Body 的字节流格式
	dom      *goquery.Document // 下载内容 Body 为html时，可转换为 Dom 对象
	items    []data.DataCell   // 存放以文本格式输出的结果数据
	files    []data.FileCell   // 存放直接输出的文件("Name": string; "Body": io.ReadCloser)
	err      error
	sync.Mutex
}

var (
	contextPool = &sync.Pool{
		New: func() interface{} {
			return &Context{
				items: []data.DataCell{},
				files: []data.FileCell{},
			}
		},
	}
)

//**************************************** 初始化 *******************************************\\

func GetContext(sp *Spider, req *request.Request) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.spider = sp
	ctx.Request = req
	return ctx
}

func PutContext(ctx *Context) {
	ctx.items = ctx.items[:0] // clear items
	ctx.files = ctx.files[:0] // clear files
	ctx.spider = nil
	ctx.Request = nil
	ctx.Response = nil
	ctx.text = nil
	ctx.dom = nil
	ctx.err = nil
	contextPool.Put(ctx)
}

func (self *Context) SetResponse(resp *http.Response) *Context {
	self.Response = resp
	return self
}

// 标记下载错误
func (self *Context) SetError(err error) {
	self.err = err
}

// 从原始请求获取Url，从而保证请求前后的Url完全相等，且中文未被编码。
func (self *Context) GetUrl() string {
	return self.Request.Url
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
	fmt.Printf("...... context#AddQueue %v \n", req.Url)
	// 若已主动终止任务，则崩溃爬虫协程
	// self.spider.TryPanic()

	err := req.
		SetSpiderName(self.spider.GetName()).
		SetEnableCookie(self.spider.GetEnableCookie()).
		Prepare()

	if err != nil {
		// logs.Log.Error(err.Error())
		fmt.Printf(err.Error())
		return self
	}

	// 自动设置 Referer
	if req.GetReferer() == "" && self.Response != nil {
		req.SetReferer(self.GetUrl())
	}

	self.spider.RequestPush(req)
	return self
}

// 用于动态规则添加请求。
func (self *Context) JsAddQueue(jreq map[string]interface{}) *Context {
	// do the job
	return self
}

// 获取下载错误
func (self *Context) GetError() error {
	// 若已主动终止任务，则奔溃爬虫协程
	self.spider.tryPanic()
	return self.err
}

// 解析响应流
// 用 ruleName 指定匹配的 ParseFunc 字段，为空时默认调用 Root()
func (self *Context) Parse(ruleName ...string) *Context {
	// 若已主动终止任务，则奔溃爬虫协程
	self.spider.tryPanic()

	_rulename, rule, found := self.getRule(ruleName...)
	if self.Response != nil {
		self.Request.SetRuleName(_rulename)
	}
	if !found {
		self.spider.RuleTree.Root(self)
		return self
	}
	if rule.ParseFunc == nil {
		// logs.Log.Error("蜘蛛 %s 的规则 %s 未定义ParseFunc", self.spider.GetName(), ruleName[0])
		return self
	}
	rule.ParseFunc(self)
	return self
}

/*

// 输出文本结果。
// item类型为map[int]interface{}时，根据ruleName现有的ItemFields字段进行输出，
// item类型为map[string]interface{}时，ruleName不存在的ItemFields字段将被自动添加，
// ruleName为空时默认当前规则。
func (self *Context) Output(item interface{}, ruleName ...string) {
	// _ruleName, rule, found := self.getRule(_ruleName)
}

// 输出文件
func (self *Context) FileOutput(name ...string) {
}

// 生成文本结果
func (self *Context) CreateItem() map[string]interface{} {
}

// 在请求中保存临时数据。
func (self *Context) SetTemp() *Context {
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

// 获取日志接口实例
func (self *Context) Log() logs.Logs {
}

// 获取蜘蛛名称。
func (self *Context) GetSpider() *Spider {
}

// 获取规则树
func (self *Context) GetRule(ruleName string) (*Rule, bool) {
	return self.spider.GetRule(ruleName)
}*/

// 获取指定规则
func (self *Context) GetRuleName() string {
	return self.Request.GetRuleName()
}

// 获取当前规则名

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
	self.Lock()
	ds = self.items
	self.items = []data.DataCell{}
	self.Unlock()
	return
}

func (self *Context) PullFiles() (fs []data.FileCell) {
	self.Lock()
	fs = self.files
	self.files = []data.FileCell{}
	self.Unlock()
	return
}

// GetHtmlParser returns goquery object binded to target crawl result
func (self *Context) GetDom() *goquery.Document {
	if self.dom == nil {
		self.initDom()
	}
	return self.dom
}

// ******************************** 私有方法 ********************************

// GetHtmlParser returns goquery object binded to target crawl result
func (self *Context) initDom() *goquery.Document {
	if self.text == nil {
		self.initText()
	}
	return self.dom
}

// GetBodyStr returns plain string crawled.
func (self *Context) initText() {
	var err error

	// 采用surf内核下载时，尝试自动转码
	if self.Request.DownloaderID == request.SURF_ID {
		var contentType, pageEncode string
		// 优先从响应头读取编码类型
		contentType = self.Response.Header.Get("Content-Type")
		if _, params, err := mime.ParseMediaType(contentType); err == nil {
			if cs, ok := params["charset"]; ok {
				pageEncode = strings.ToLower(strings.TrimSpace(cs))
			}
		}
		// 响应头未指定编码类型时，从请求头读取
		if len(pageEncode) == 0 {
			contentType = self.Request.Header.Get("Content-Type")
			if _, params, err := mime.ParseMediaType(contentType); err == nil {
				if cs, ok := params["charset"]; ok {
					pageEncode = strings.ToLower(strings.TrimSpace(cs))
				}
			}
		}

		switch pageEncode {
		// 不做转码处理
		case "utf8", "utf-8", "unicode-1-1-utf-8":
		default:
			// 指定了编码类型，但不是utf8时，自动转码为utf8
			// get converter to utf-8
			// Charset auto determine. Use golang.org/x/net/html/charset. Get response body and change it to utf-8
			var destReader io.Reader
			if len(pageEncode) == 0 {
				destReader, err = charset.NewReader(self.Response.Body, "")
			} else {
				destReader, err = charset.NewReaderLabel(pageEncode, self.Response.Body)
			}

			if err == nil {
				self.text, err = ioutil.ReadAll(destReader)
				if err == nil {
					self.Response.Body.Close()
					return
				} else {
					logs.Log.Warning(" *     [convert][%v]: %v (ignore transcoding)\n", self.GetUrl(), err)
				}
			} else {
				logs.Log.Warning(" *     [convert][%v]: %v (ignore transcoding)\n", self.GetUrl(), err)
			}
		}

	}
	// 不做转码处理
	self.text, err = ioutil.ReadAll(self.Response.Body)
	self.Response.Body.Close()
	if err != nil {
		panic(err.Error())
		return
	}
}

// 获取规则
func (self *Context) getRule(ruleName ...string) (name string, rule *Rule, found bool) {
	if len(ruleName) == 0 {
		if self.Response == nil {
			return
		}
		name = self.GetRuleName()
	} else {
		name = ruleName[0]
	}
	rule, found = self.spider.GetRule(name)
	return
}

/**
 * 编码类型参考
 * 不区分大小写
 *var nameMap = map[string]htmlEncoding{
	"unicode-1-1-utf-8":   utf8,
	"utf-8":               utf8,
	"utf8":                utf8,
	"866":                 ibm866,
	"cp866":               ibm866,
	"csibm866":            ibm866,
	"ibm866":              ibm866,
	"csisolatin2":         iso8859_2,
	"iso-8859-2":          iso8859_2,
	"iso-ir-101":          iso8859_2,
	"iso8859-2":           iso8859_2,
	"iso88592":            iso8859_2,
	"iso_8859-2":          iso8859_2,
	"iso_8859-2:1987":     iso8859_2,
	"l2":                  iso8859_2,
	"latin2":              iso8859_2,
	"csisolatin3":         iso8859_3,
	"iso-8859-3":          iso8859_3,
	"iso-ir-109":          iso8859_3,
	"iso8859-3":           iso8859_3,
	"iso88593":            iso8859_3,
	"iso_8859-3":          iso8859_3,
	"iso_8859-3:1988":     iso8859_3,
	"l3":                  iso8859_3,
	"latin3":              iso8859_3,
	"csisolatin4":         iso8859_4,
	"iso-8859-4":          iso8859_4,
	"iso-ir-110":          iso8859_4,
	"iso8859-4":           iso8859_4,
	"iso88594":            iso8859_4,
	"iso_8859-4":          iso8859_4,
	"iso_8859-4:1988":     iso8859_4,
	"l4":                  iso8859_4,
	"latin4":              iso8859_4,
	"csisolatincyrillic":  iso8859_5,
	"cyrillic":            iso8859_5,
	"iso-8859-5":          iso8859_5,
	"iso-ir-144":          iso8859_5,
	"iso8859-5":           iso8859_5,
	"iso88595":            iso8859_5,
	"iso_8859-5":          iso8859_5,
	"iso_8859-5:1988":     iso8859_5,
	"arabic":              iso8859_6,
	"asmo-708":            iso8859_6,
	"csiso88596e":         iso8859_6,
	"csiso88596i":         iso8859_6,
	"csisolatinarabic":    iso8859_6,
	"ecma-114":            iso8859_6,
	"iso-8859-6":          iso8859_6,
	"iso-8859-6-e":        iso8859_6,
	"iso-8859-6-i":        iso8859_6,
	"iso-ir-127":          iso8859_6,
	"iso8859-6":           iso8859_6,
	"iso88596":            iso8859_6,
	"iso_8859-6":          iso8859_6,
	"iso_8859-6:1987":     iso8859_6,
	"csisolatingreek":     iso8859_7,
	"ecma-118":            iso8859_7,
	"elot_928":            iso8859_7,
	"greek":               iso8859_7,
	"greek8":              iso8859_7,
	"iso-8859-7":          iso8859_7,
	"iso-ir-126":          iso8859_7,
	"iso8859-7":           iso8859_7,
	"iso88597":            iso8859_7,
	"iso_8859-7":          iso8859_7,
	"iso_8859-7:1987":     iso8859_7,
	"sun_eu_greek":        iso8859_7,
	"csiso88598e":         iso8859_8,
	"csisolatinhebrew":    iso8859_8,
	"hebrew":              iso8859_8,
	"iso-8859-8":          iso8859_8,
	"iso-8859-8-e":        iso8859_8,
	"iso-ir-138":          iso8859_8,
	"iso8859-8":           iso8859_8,
	"iso88598":            iso8859_8,
	"iso_8859-8":          iso8859_8,
	"iso_8859-8:1988":     iso8859_8,
	"visual":              iso8859_8,
	"csiso88598i":         iso8859_8I,
	"iso-8859-8-i":        iso8859_8I,
	"logical":             iso8859_8I,
	"csisolatin6":         iso8859_10,
	"iso-8859-10":         iso8859_10,
	"iso-ir-157":          iso8859_10,
	"iso8859-10":          iso8859_10,
	"iso885910":           iso8859_10,
	"l6":                  iso8859_10,
	"latin6":              iso8859_10,
	"iso-8859-13":         iso8859_13,
	"iso8859-13":          iso8859_13,
	"iso885913":           iso8859_13,
	"iso-8859-14":         iso8859_14,
	"iso8859-14":          iso8859_14,
	"iso885914":           iso8859_14,
	"csisolatin9":         iso8859_15,
	"iso-8859-15":         iso8859_15,
	"iso8859-15":          iso8859_15,
	"iso885915":           iso8859_15,
	"iso_8859-15":         iso8859_15,
	"l9":                  iso8859_15,
	"iso-8859-16":         iso8859_16,
	"cskoi8r":             koi8r,
	"koi":                 koi8r,
	"koi8":                koi8r,
	"koi8-r":              koi8r,
	"koi8_r":              koi8r,
	"koi8-ru":             koi8u,
	"koi8-u":              koi8u,
	"csmacintosh":         macintosh,
	"mac":                 macintosh,
	"macintosh":           macintosh,
	"x-mac-roman":         macintosh,
	"dos-874":             windows874,
	"iso-8859-11":         windows874,
	"iso8859-11":          windows874,
	"iso885911":           windows874,
	"tis-620":             windows874,
	"windows-874":         windows874,
	"cp1250":              windows1250,
	"windows-1250":        windows1250,
	"x-cp1250":            windows1250,
	"cp1251":              windows1251,
	"windows-1251":        windows1251,
	"x-cp1251":            windows1251,
	"ansi_x3.4-1968":      windows1252,
	"ascii":               windows1252,
	"cp1252":              windows1252,
	"cp819":               windows1252,
	"csisolatin1":         windows1252,
	"ibm819":              windows1252,
	"iso-8859-1":          windows1252,
	"iso-ir-100":          windows1252,
	"iso8859-1":           windows1252,
	"iso88591":            windows1252,
	"iso_8859-1":          windows1252,
	"iso_8859-1:1987":     windows1252,
	"l1":                  windows1252,
	"latin1":              windows1252,
	"us-ascii":            windows1252,
	"windows-1252":        windows1252,
	"x-cp1252":            windows1252,
	"cp1253":              windows1253,
	"windows-1253":        windows1253,
	"x-cp1253":            windows1253,
	"cp1254":              windows1254,
	"csisolatin5":         windows1254,
	"iso-8859-9":          windows1254,
	"iso-ir-148":          windows1254,
	"iso8859-9":           windows1254,
	"iso88599":            windows1254,
	"iso_8859-9":          windows1254,
	"iso_8859-9:1989":     windows1254,
	"l5":                  windows1254,
	"latin5":              windows1254,
	"windows-1254":        windows1254,
	"x-cp1254":            windows1254,
	"cp1255":              windows1255,
	"windows-1255":        windows1255,
	"x-cp1255":            windows1255,
	"cp1256":              windows1256,
	"windows-1256":        windows1256,
	"x-cp1256":            windows1256,
	"cp1257":              windows1257,
	"windows-1257":        windows1257,
	"x-cp1257":            windows1257,
	"cp1258":              windows1258,
	"windows-1258":        windows1258,
	"x-cp1258":            windows1258,
	"x-mac-cyrillic":      macintoshCyrillic,
	"x-mac-ukrainian":     macintoshCyrillic,
	"chinese":             gbk,
	"csgb2312":            gbk,
	"csiso58gb231280":     gbk,
	"gb2312":              gbk,
	"gb_2312":             gbk,
	"gb_2312-80":          gbk,
	"gbk":                 gbk,
	"iso-ir-58":           gbk,
	"x-gbk":               gbk,
	"gb18030":             gb18030,
	"big5":                big5,
	"big5-hkscs":          big5,
	"cn-big5":             big5,
	"csbig5":              big5,
	"x-x-big5":            big5,
	"cseucpkdfmtjapanese": eucjp,
	"euc-jp":              eucjp,
	"x-euc-jp":            eucjp,
	"csiso2022jp":         iso2022jp,
	"iso-2022-jp":         iso2022jp,
	"csshiftjis":          shiftJIS,
	"ms932":               shiftJIS,
	"ms_kanji":            shiftJIS,
	"shift-jis":           shiftJIS,
	"shift_jis":           shiftJIS,
	"sjis":                shiftJIS,
	"windows-31j":         shiftJIS,
	"x-sjis":              shiftJIS,
	"cseuckr":             euckr,
	"csksc56011987":       euckr,
	"euc-kr":              euckr,
	"iso-ir-149":          euckr,
	"korean":              euckr,
	"ks_c_5601-1987":      euckr,
	"ks_c_5601-1989":      euckr,
	"ksc5601":             euckr,
	"ksc_5601":            euckr,
	"windows-949":         euckr,
	"csiso2022kr":         replacement,
	"hz-gb-2312":          replacement,
	"iso-2022-cn":         replacement,
	"iso-2022-cn-ext":     replacement,
	"iso-2022-kr":         replacement,
	"utf-16be":            utf16be,
	"utf-16":              utf16le,
	"utf-16le":            utf16le,
	"x-user-defined":      xUserDefined,
}*/
