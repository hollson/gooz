//-------------------------------------------------------------------------------------
// @ Copyright (C) free license,without warranty of any kind .
// @ Author: hollson <hollson@live.com>
// @ Date: 2020-04-30
// @ Version: 1.0.0
//
// 命令行参数，使程序启动为守护进程，如 ./appname -d=true
// thanks for "github.com/icattlecoder/godaemon"
//-------------------------------------------------------------------------------------

package app

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

var godaemon = flag.Bool("d", false, "run app as a daemon with -d=true")

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
	if *godaemon {
		args := os.Args[1:]
		i := 0
		for ; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		cmd.Start()
		fmt.Println("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}
}
