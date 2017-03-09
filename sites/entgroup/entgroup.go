package sites

import (
	"fmt"
	"github.com/zhenwusw/logan/app/downloader/request"
	. "github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/common/goquery"
)

// site '艺恩', 'entgroup'
// self.encoding = "UTF-8"
// exec Context.new("http://www.entgroup.cn/news/Exclusive/", nil, category)

func init() {
	Entgroup.Register()
}

var Entgroup = &Spider{
	Name:        "艺恩娱乐",
	Description: "艺恩娱乐",
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			fmt.Printf("...... %v \n", "fetching Entgroup")
			ctx.AddQueue(&request.Request{Url: "http://www.entgroup.cn/news/Exclusive/", Rule: "排行榜主页"})
		},

		Trunk: map[string]*Rule{
			"xxx": {
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					query.Find(".subNav a").Each(func(i int, s *goquery.Selection) {
						if url, ok := s.Attr("href"); ok {
							ctx.AddQueue(&request.Request{Url: url, Rule: "原创"})
						}
					})
				},
			},

			"yyy": {
				ParseFunc: func(ctx *Context) {
				},
			},
		},
	},
}
