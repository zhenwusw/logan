package downloader

import (
	"github.com/zhenwusw/logan/app/spider"
	"github.com/zhenwusw/logan/app/downloader/request"
)

// The Downloader interface
// You can implement the interface by implement function Download
// Function Download need to return Page instance pointer that has request result downloaded from Request

type Downloader interface {
	Download(spider *spider.Spider, request *request.Request) *spider.Context
}