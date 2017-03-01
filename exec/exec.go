package exec

import (
	"runtime"
	"github.com/zhenwusw/logan/common/gc"
	"fmt"
)

func init() {
	// Enable max gorotines according to CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())
	gc.ManualGC() // ???
}

func DefaultRun() {
	fmt.Printf()
}

func flagCommon() {
}

func writeFlag() {
}