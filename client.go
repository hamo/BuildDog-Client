package main

import (
	"flag"
	"os"
	"path/filepath"
)

const (
	confName = ".BuildDog"
	servAddr = "http://127.0.0.1:8888"
)

var (
	firstrun bool
	config   *confStruct
)

var (
	flConfDir  string
	flReConfig bool
)

func init() {
	flag.StringVar(&flConfDir, "confdir", os.Getenv("HOME"), "")
	flag.BoolVar(&flReConfig, "r", false, "")
}

func main() {
	if _, err := os.Stat(filepath.Join(flConfDir, confName)); err == nil {
		firstrun = false
	} else {
		firstrun = true
	}

	if flReConfig {
		firstrun = true
	}

	if !firstrun {
		var err error
		config, err = readConf(filepath.Join(flConfDir, confName))
		switch err {
		case errNeedReconfigure:
			firstrun = true
		case nil:
		default:
			panic(err)
		}
	}

	if firstrun {
		firstRun()
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		panic("arg error")
	}

	switch os.Args[1] {
	case "build":
		cmdBuild(os.Args[2:])
	case "task":
		cmdTask(os.Args[2:])
	default:
		panic("arg ERROR")
	}

}
