package client

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type FlyteHttpClient struct {
	client *http.Client
}

func NewFlyteHttpClient() FlyteHttpClient {
	return FlyteHttpClient{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c FlyteHttpClient) DoRequest(method string, url string, body string, headers map[string][]string) (*http.Response, error) {
	if err := validateArgs(method, url); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(body)))
	if err != nil {
		return nil, err
	}
	addHeaders(req, headers)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func validateArgs(method string, url string) error {
	if method == "" {
		return fmt.Errorf("error: method not set")
	}
	if !isValidHttpMethod(method) {
		return fmt.Errorf("error: invalid method supplied")
	}
	if url == "" {
		return fmt.Errorf("error: url not set")
	}
	return nil
}

func isValidHttpMethod(input string) bool {
	var validMethods = []string{"GET", "OPTIONS", "HEAD", "POST", "PUT", "DELETE", "TRACE", "CONNECT", "PATCH"}

	for _, method := range validMethods {
		if method == strings.ToUpper(input) {
			return true
		}
	}
	return false
}

func addHeaders(req *http.Request, headers map[string][]string) {
	if headers != nil {
		for k, v := range headers {
			for _, e := range v {
				req.Header.Add(k, e)
			}
		}
	}
}
