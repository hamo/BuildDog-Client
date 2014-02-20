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
 * @args[1](optional): sub-action
 */
func cmdTask(args []string) error {
	taskId := args[0]

	if len(args) == 1 {
		// Get task status

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
	}

	switch args[1] {
	case "output":
		return cmdTaskOutput(taskId)
	default:
		panic("ERR")
	}

	return nil
}

func cmdTaskOutput(taskId string) error {
	action := "/output"
	resp, _ := request("GET", servAddr+taskAPI+taskId+action, nil, "")
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		fmt.Print(body)
	}

	return nil
}
