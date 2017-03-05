package spider

import (
	"github.com/zhenwusw/logan/common/util"
	"math"
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
		id      int
		subName string
		// reqMatrix *scheduler.Matrix // 请求矩阵
		// timer *Timer
		// status
		// lock
		// once
	}

	RuleTree struct {
	}

	Rule struct {
	}
)

// 获取蜘蛛名称
func (self *Spider) GetName() string {
	return "default spider name"
}

/*
// 添加自身到蜘蛛菜单
func (self Spider) Register() *Spider {
}

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

// 获取蜘蛛二级标识名
func (self *Spider) GetSubName() string {
}

// 安全返回指定规则
func (self *Spider) GetRule() (*Rule, bool) {
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

// 获取蜘蛛 ID
func (self *Spider) GetId() int {
}

// 设置蜘蛛ID
func (self *Spider) SetId(id int) {
}

// 获取自定义配置信息
func (self *Spider) GetKeyin() {
}

// 设置自定义配置信息
func (self *Spider) SetKeyin(keyword string) {
}

// 获取采集上限
// <0 表示采用限制请求书的方案
// >0 表示采用规则中的自定义限制方案
func (self *Spider) GetLimit() int64 {
}

// 设置采集上限
func (self *Spider) SetLimit(max int64) {
}

// 控制所有请求是否使用 cookie
func (self *Spider) GetEnableCookie() bool {
}

// 自定义暂停时间
func (self *Spider) SetPausetime(pause int64, runtime ...bool) {
}

// 设置定时器
// @id为定时器唯一标识
func (self *Spider) SetTimer() bool {
}

// 启动定时器
func (self *Spider) RunTimer() bool {
}

// 返回一个自身复制品
func (self *Spider) Copy() *Spider {
}

func (self *Spider) ReqmatrixInit() *Spider {
}

// 返回是否作为新的失败请求被添加至队列尾部
func (self *Spider) DoHistory(req *request.Req) {
}
*/
