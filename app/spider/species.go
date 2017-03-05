package spider

import "github.com/zhenwusw/logan/common/pinyin"

// 蜘蛛种类
type SpiderSpecies struct {
	list   []*Spider
	hash   map[string]*Spider
	sorted bool
}

// 全局蜘蛛种类实例
var Species = &SpiderSpecies{
	list: []*Spider{},
	hash: map[string]*Spider{},
}

// 获取全部蜘蛛种类
func (self *SpiderSpecies) Get() []*Spider {
	if !self.sorted {
		l := len(self.list)
		initials := make([]string, l)
		newlist := map[string]*Spider{}
		for i := 0; i < l; i++ {
			initials[i] = self.list[i].GetName()
			newlist[initials[i]] = self.list[i]
		}
		pinyin.SortInitials(initials)
		for i := 0; i < l; i++ {
			self.list[i] = newlist[initials[i]]
		}
		self.sorted = true
	}
	return self.list
}

/*
// 向蜘蛛种类清单添加新种类
func (self *SpiderSpecies) Add(sp *Spider) *Spider {
	name := sp.Name
}

func (self *SpiderSpecies) GetByName(name string) *Spider {
	return self.hash[name]
}*/
