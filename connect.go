package main

import (
	"net/http"
	"net/url"
	"strings"
	"io/ioutil"
)

const (
	userAgent = "BuildDog-Client"
	token     = "testing"
)

func request(method string, urlStr string, form map[string]string, body string) (*http.Response, error) {
	client := new(http.Client)

	req, err := http.NewRequest(method, urlStr, strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)
	req.Header.Add("Auth-Token", token)
	req.Header.Add("Username", config.username)

	if form != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		val := url.Values{}
		for k, v := range form {
			val.Add(k, v)
		}

		req.Body = ioutil.NopCloser(strings.NewReader(val.Encode()))
	}

	return client.Do(req)
}
