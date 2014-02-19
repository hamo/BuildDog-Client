package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	taskAPI = "/task/"
)

type taskStatus struct {
	Creator string

	Repo string
	Rev  string
	PPA  string

	Status  string
	Process string
	Error   string

	WorkingDir string
}

/*
 * @args[0]: task id
 */
func cmdTask(args []string) error {
	taskId := args[0]

	resp, _ := request("GET", servAddr+taskAPI+taskId, nil, "")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ts taskStatus
	if resp.StatusCode == 200 {
		dec := json.NewDecoder(strings.NewReader(string(body)))
		dec.Decode(&ts)
		fmt.Printf("%+v", ts)
	}

	return nil
}
