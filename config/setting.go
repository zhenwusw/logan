package config

import (
	"github.com/zhenwusw/logan/common/config"
	"path/filepath"
	"os"
	"strconv"
	"fmt"
	"github.com/zhenwusw/logan/runtime/status"
)

const (
	crawlcap int = 50 // 蜘蛛池最大容量
	// datachancap             int    = 2 << 14                     // 收集器容量(默认65536)
	logcap                  int64  = 10000
	loglevel                string = "debug"
	logconsolelevel         string = "info"
	logfeedbacklevel        string = "error"
	loglineinfo             bool   = false
	logsave                 bool   = true

	phantomjs               string = WORK_ROOT + "/phantomjs"    // phantomjs文件路径
	proxylib                string = WORK_ROOT + "/proxy.lib"    // 代理ip文件路径
	spiderdir               string = WORK_ROOT + "/spiders"      // 动态规则目录
	fileoutdir              string = WORK_ROOT + "/file_out"     // 文件（图片、HTML等）结果的输出目录
	textoutdir              string = WORK_ROOT + "/text_out"     // excel或csv输出方式下，文本结果的输出目录
	dbname                  string = TAG                         // 数据库名称
	mgoconnstring           string = "127.0.0.1:27017"           // mongodb连接字符串
	mgoconncap              int    = 1024                        // mongodb连接池容量
	mgoconngcsecond         int64  = 600                         // mongodb连接池GC时间，单位秒
	mysqlconnstring         string = "root:@tcp(127.0.0.1:3306)" // mysql连接字符串
	mysqlconncap            int    = 2048                        // mysql连接池容量
	mysqlmaxallowedpacketmb int    = 1                           //mysql通信缓冲区的最大长度，单位MB，默认1MB
	kafkabrokers            string = "127.0.0.1:9092"            //kafka broker字符串,逗号分割

	mode                    int    = status.UNSET
	port                    int    = 2015
	master                  string = "127.0.0.1"
	thread                  int    = 20
	pause                   int64  = 300
	outtype                 string = "csv"
	dockercap               int    = 10000
	limit                   int64  = 0
	proxyminute             int64  = 0
	success                 bool   = true
	failure                 bool   = true
)

var setting = func() config.Configer {
	os.MkdirAll(filepath.Clean(HISTORY_DIR), 0777)
	os.MkdirAll(filepath.Clean(CACHE_DIR), 0777)
	os.MkdirAll(filepath.Clean(PHANTOMJS_TEMP), 0777)

	iniconf, err := config.NewConfig("ini", CONFIG)
	if err != nil {
		file, err := os.Create(CONFIG)
		file.Close()
		iniconf, err = config.NewConfig("ini", CONFIG)
		if err != nil {
			panic(err)
		}
		defaultConfig(iniconf)
		iniconf.SaveConfigFile(CONFIG)
	} else {
		trySet(iniconf)
	}

	os.MkdirAll(filepath.Clean(iniconf.String("spiderdir")), 0777)
	os.MkdirAll(filepath.Clean(iniconf.String("fileoutdir")), 0777)
	os.MkdirAll(filepath.Clean(iniconf.String("textoutdir")), 0777)

	return iniconf
}()

func defaultConfig(iniconf config.Configer) {
	iniconf.Set("crawlcap", strconv.Itoa(crawlcap))
	// iniconf.Set("datachancap", strconv.Itoa(datachancap))
	iniconf.Set("log::cap", strconv.FormatInt(logcap, 10))
	iniconf.Set("log::level", loglevel)
	iniconf.Set("log::consolelevel", logconsolelevel)
	iniconf.Set("log::feedbacklevel", logfeedbacklevel)
	iniconf.Set("log::lineinfo", fmt.Sprint(loglineinfo))
	iniconf.Set("log::save", fmt.Sprint(logsave))
	// iniconf.Set("phantomjs", phantomjs)
	// iniconf.Set("proxylib", proxylib)
	// iniconf.Set("spiderdir", spiderdir)
	// iniconf.Set("fileoutdir", fileoutdir)
	// iniconf.Set("textoutdir", textoutdir)
	// iniconf.Set("dbname", dbname)
	// iniconf.Set("mgo::connstring", mgoconnstring)
	// iniconf.Set("mgo::conncap", strconv.Itoa(mgoconncap))
	// iniconf.Set("mgo::conngcsecond", strconv.FormatInt(mgoconngcsecond, 10))
	// iniconf.Set("mysql::connstring", mysqlconnstring)
	// iniconf.Set("mysql::conncap", strconv.Itoa(mysqlconncap))
	// iniconf.Set("mysql::maxallowedpacketmb", strconv.Itoa(mysqlmaxallowedpacketmb))
	// iniconf.Set("kafka::brokers", kafkabrokers)
	iniconf.Set("run::mode", strconv.Itoa(mode))
	iniconf.Set("run::port", strconv.Itoa(port))
	iniconf.Set("run::master", master)
	iniconf.Set("run::thread", strconv.Itoa(thread))
	iniconf.Set("run::pause", strconv.FormatInt(pause, 10))
	iniconf.Set("run::outtype", outtype)
	iniconf.Set("run::dockercap", strconv.Itoa(dockercap))
	iniconf.Set("run::limit", strconv.FormatInt(limit, 10))
	iniconf.Set("run::proxyminute", strconv.FormatInt(proxyminute, 10))
	iniconf.Set("run::success", fmt.Sprint(success))
	iniconf.Set("run::failure", fmt.Sprint(failure))
}

func trySet(iniconf config.Configer) {
}
