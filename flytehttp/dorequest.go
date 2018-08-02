/*
Copyright (C) 2018 Expedia Group.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flytehttp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var validMethods = []string{http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodConnect, http.MethodTrace,
	http.MethodPatch, http.MethodPut, http.MethodHead, http.MethodOptions}

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

func DoRequest(client HttpClient, input DoRequestInput) (*doRequestOutputPayload, error) {
	if err := validateMethod(input.Method); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(input.Method, input.URL, bytes.NewReader([]byte(input.Body)))
	if err != nil {
		return nil, err
	}

	req.Header = input.Headers

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return &doRequestOutputPayload{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       contentTypeHandler(resp.Header.Get("Content-Type"), bodyContent),
	}, nil
}

func validateMethod(method string) error {
	if method == "" {
		return fmt.Errorf("no method provided")
	} else {
		if !isValidHttpMethod(method) {
			return fmt.Errorf("invalid method provided")
		}
		return nil
	}
}

func isValidHttpMethod(input string) bool {
	for _, method := range validMethods {
		if method == strings.ToUpper(input) {
			return true
		}
	}
	return false
}

func contentTypeHandler(contentType string, body []byte) interface{} {
	switch contentType {
	case "application/json":
		return body
	default:
		return base64.URLEncoding.EncodeToString(body)
	}
}
