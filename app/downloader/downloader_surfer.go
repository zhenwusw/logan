package downloader

import (
	"errors"
	"github.com/zhenwusw/logan/app/downloader/request"
	"github.com/zhenwusw/logan/app/downloader/surfer"
	"github.com/zhenwusw/logan/app/spider"
	"net/http"
)

type Surfer struct {
	surf surfer.Surfer
	// phantom surfer.Surfer
}

var SurferDownloader = &Surfer{
	surf: surfer.New(),
	// phantom: surfer.NewPhantom()
}

func (self *Surfer) Download(sp *spider.Spider, cReq *request.Request) *spider.Context {
	ctx := spider.GetContext(sp, cReq)

	var resp *http.Response
	var err error

	switch cReq.GetDownloaderID() {
	case request.SURF_ID:
		resp, err = self.surf.Download(cReq)
	case request.PHANTOM_ID:
		// resp, err = self.phantom.Download(cReq)
	}

	if resp.StatusCode >= 400 {
		err = errors.New("响应状态 " + resp.Status)
	}

	ctx.SetResponse(resp).SetError(err)
	return ctx
}
