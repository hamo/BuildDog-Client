package main

import (
	"fmt"
	"strings"
)

const (
	buildAPI = "/build"
)

/*
 * @args[0]: repo[#rev]
 * @args[1]: target PPA
 */
func cmdBuild(args []string) error {
	addr := strings.SplitN(args[0], "#", 2)
	repo := addr[0]
	var rev string
	if len(addr) == 2 {
		rev = addr[1]
	}

	ppa := args[1]

	form := make(map[string]string)
	form["repo"] = repo
	form["ppa"] = ppa
	form["rev"] = rev

	resp, _ := request("POST", servAddr+buildAPI, form, "")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}
	fmt.Print(body)

	return nil
}
