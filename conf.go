package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	confVersion = "1"
)

var (
	errNeedReconfigure = errors.New("Need Reconfigure")
)

type confStruct struct {
	username string
}

func firstRun() {
	var conf confStruct

	fmt.Printf("Please enter your username:")
	fmt.Scanf("%s\n", &conf.username)

	updateConf(conf, filepath.Join(flConfDir, confName))

}

func updateConf(c confStruct, f string) error {
	file, err := os.Create(f)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintf(file, "version=%d\n", confVersion)
	fmt.Fprintf(file, "username=%s\n", c.username)
	return nil
}

func readConf(f string) (*confStruct, error) {
	re := new(confStruct)

	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.SplitN(scanner.Text(), "=", 2)
		switch s[0] {
		case "username":
			re.username = s[1]
		case "version":
			if s[1] != confVersion {
				return nil, errNeedReconfigure
			}
		default:
			panic("ERROR")
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return re, nil
}
