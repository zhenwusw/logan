package main

import (
	"github.com/zhenwusw/logan/exec"
	_ "github.com/zhenwusw/logan/sites"
)

func main() {
	exec.DefaultRun("cmd")
}
