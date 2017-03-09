package spider

import (
	"github.com/zhenwusw/logan/app/scheduler"
	"github.com/zhenwusw/logan/common/util"
	"github.com/zhenwusw/logan/runtime/status"
	"math"
	"sync"
	"fmt"
	"github.com/zhenwusw/logan/app/downloader/request"
)

const (
	KEYIN       = util.USE_KEYIN
	LIMIT       = math.MaxInt64
	FORCED_STOP = "——主动终止Spider——"
)

type (
	// 蜘蛛规则
	Spider struct {
		// 由用户指定
		Name            string
		Description     string
		Pausetime       int64
		Limit           int64
		Keyin           string
		EnableCookie    bool
		NotDefaultField bool
		Namespace       func(self *Spider) string
		SubNamespace    func(self *Spider, dataCell map[string]interface{}) string
		RuleTree        *RuleTree

		// 有系统自动赋值
		id        int
		subName   string
		reqMatrix *scheduler.Matrix // 请求矩阵
		// timer *Timer
		status int // 执行状态
		lock   sync.RWMutex
		// once
	}

	RuleTree struct {
		Root  func(*Context)   // 根节点（执行入口）
		Trunk map[string]*Rule // 节点散列（执行采集过程）
	}

	Rule struct {
		ItemFields []string                                           // 结果字段列表（选填，写上可保证字段顺序）
		ParseFunc  func(*Context)                                     // 内容解析函数
		AidFunc    func(*Context, map[string]interface{}) interface{} // 通用辅助函数
	}
)

// 获取蜘蛛名称
func (self *Spider) GetName() string {
	return "default spider name"
}

// 添加自身到蜘蛛菜单
func (self Spider) Register() *Spider {
	self.status = status.STOPPED
	return Species.Add(&self)
}

// 安全返回指定规则
func (self *Spider) GetRule(ruleName string) (*Rule, bool) {
	rule, found := self.RuleTree.Trunk[ruleName]
	return rule, found
}

func (self *Spider) ReqmatrixInit() *Spider {
	if self.Limit < 0 {
		self.reqMatrix = scheduler.AddMatrix(self.GetName(), self.GetSubName(), self.Limit)
		self.SetLimit(0)
	} else {
		self.reqMatrix = scheduler.AddMatrix(self.GetName(), self.GetSubName(), math.MinInt64)
	}
	return self
}

// 获取蜘蛛二级标识名
func (self *Spider) GetSubName() string {
	return "Spider GetSubName()"
}

// 设置采集上限
// <0 表示采用限制请求数的方案
// >0 表示采用规则中的自定义限制方案
func (self *Spider) SetLimit(max int64) {
	self.Limit = max
}

// 开始执行蜘蛛
func (self *Spider) Start() {
	fmt.Printf("...... %v \n", "spider#Start()")

	defer func() {
		if p := recover(); p != nil {
			fmt.Printf(" *     Panic  [root]: %v\n", p)
			// logs.Log.Error
		}
		self.lock.Lock()
		self.status = status.RUN
		self.lock.Unlock()
	}()
	self.RuleTree.Root(GetContext(self, nil))
}

// 主动崩溃爬虫运行协程
func (self *Spider) Stop() {
}

func (self *Spider) CanStop() bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.status != status.STOPPED && self.reqMatrix.CanStop()
}

func (self *Spider) IsStopping() bool {
	self.lock.RLock()
	defer self.lock.RUnlock()
	return self.status == status.STOP
}

func (self *Spider) RequestPush(req *request.Request) {
	self.reqMatrix.Push(req)
}

func (self *Spider) RequestPull() *request.Request {
	return self.reqMatrix.Pull()
}

func (self *Spider) RequestUse() {
	self.reqMatrix.Use()
}

func (self *Spider) RequestFree() {
	self.reqMatrix.Free()
}

// 退出前收尾工作
func (self *Spider) Defer() {
	// 取消所有定时器
	//if self.timer != nil {
	//    self.timer.drop()
	//    self.timer = nil
	//}

	// 等待处理中的请求完成
	self.reqMatrix.Wait()
	// 更新失败记录
	self.reqMatrix.TryFlushFailure()
}

// 返回一个自身复制品
func (self *Spider) Copy() *Spider {
	ghost := &Spider{}
	ghost.Name = self.Name
	ghost.subName = self.subName

	ghost.RuleTree = &RuleTree{
		Root: self.RuleTree.Root,
		Trunk: make(map[string]*Rule, len(self.RuleTree.Trunk)),
	}
	for k, v := range self.RuleTree.Trunk {
		ghost.RuleTree.Trunk[k] = new(Rule)
		ghost.RuleTree.Trunk[k].ItemFields = make([]string, len(v.ItemFields))
		copy(ghost.RuleTree.Trunk[k].ItemFields, v.ItemFields)

		ghost.RuleTree.Trunk[k].ParseFunc = v.ParseFunc
		ghost.RuleTree.Trunk[k].AidFunc = v.AidFunc
	}

	ghost.Description = self.Description
	ghost.Pausetime = self.Pausetime
	ghost.EnableCookie = self.EnableCookie
	ghost.Limit = self.Limit
	ghost.Keyin = self.Keyin

	ghost.NotDefaultField = self.NotDefaultField
	ghost.Namespace = self.Namespace
	ghost.SubNamespace = self.SubNamespace

	// ghost.timer = self.timer
	ghost.status = self.status

	return ghost
}

// 自定义暂停时间 pause[0]~(pause[0]+pause[1])，优先级高于外部传参
// 当且仅当runtime[0]为true时可覆盖现有值
func (self *Spider) SetPausetime(pause int64, runtime ...bool) {
	if self.Pausetime == 0 || len(runtime) > 0 && runtime[0] {
		self.Pausetime = pause
	}
}

// 获取采集上限
// <0 表示采用限制请求书的方案
// >0 表示采用规则中的自定义限制方案
func (self *Spider) GetLimit() int64 {
	return self.Limit
}

// 获取蜘蛛 ID
func (self *Spider) GetId() int {
	return self.id
}

// 设置蜘蛛ID
func (self *Spider) SetId(id int) {
	self.id = id
}

// 控制所有请求是否使用 cookie
func (self *Spider) GetEnableCookie() bool {
	return self.EnableCookie
}


/*
// 指定规则的获取结果的字段名列表
func (self *Spider) GetItemFields(rule *Rule) []string {
}

// 返回结果字段名的值
// 不存在时返回空字符串
func (self *Spider) GetItemField() (fiedl string) {
}

// 返回结果字段名的索引
// 不存在时为 -1
func (self *Spider) GetItemFieldIndex() (index int) {
}

// 为指定 Rule 动态追加结果字段名, 并返回索引位置
// 已存在时返回原来索引位置
func (self *Spider) UpsertItemField() (index int) {
}

// 返回指定规则
func (self *Spider) MustGetRule() *Rule {
}

// 返回规则树
func (self *Spider) GetRules() map[string]*Rule {
}

// 获取蜘蛛描述
func (self *Spider) GetDescription() string {
}

// 获取自定义配置信息
func (self *Spider) GetKeyin() {
}

// 设置自定义配置信息
func (self *Spider) SetKeyin(keyword string) {
}

// 设置定时器
// @id为定时器唯一标识
func (self *Spider) SetTimer() bool {
}

// 启动定时器
func (self *Spider) RunTimer() bool {
}

// 返回是否作为新的失败请求被添加至队列尾部
func (self *Spider) DoHistory(req *request.Req) {
}
*/
