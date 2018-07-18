// +build acceptance

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

package acceptancetest

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"
	"bytes"
	"github.com/HotelsDotCom/flyte-http/client"
)

func TestDoRequestWithPost(t *testing.T) {
	c := client.NewFlyteHttpClient()
	require.NotNil(t, c)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
		} else if r.RequestURI != "/testuri" {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusOK)
			if r.Body != nil {
				body, _ := ioutil.ReadAll(r.Body)
				w.Write(body)
			}
		}
	}))
	defer ts.Close()

	testUri := ts.URL + "/test"
	resp, err := c.DoRequest("POST", testUri, `{"name":"test"}`, nil)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDoRequestUsingGetShouldReturnCorrectResource(t *testing.T) {
	c := client.NewFlyteHttpClient()

	payload := []byte(`{"name":"test"}`)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(payload)
	}))
	defer ts.Close()

	testUrl := ts.URL + "/testuri"
	resp, err := c.DoRequest("GET", testUrl, "", nil)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBodyContent, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	expectedBodyContent, err := ioutil.ReadAll(bytes.NewReader(payload))
	require.NoError(t, err)

	assert.Equal(t, expectedBodyContent, respBodyContent)
}

//func TestDoRequestShouldReturnResponseIfRequestIsSuccessful(t *testing.T) {
//	c := NewFlyteHttpClient()
//
//	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		defer r.Body.Close()
//		if r.Method != http.MethodPost {
//			w.WriteHeader(http.StatusBadRequest)
//		} else if r.RequestURI != "/testuri" {
//			w.WriteHeader(http.StatusBadRequest)
//		} else {
//			w.WriteHeader(http.StatusOK)
//			if r.Body != nil {
//				body, _ := ioutil.ReadAll(r.Body)
//				w.Write(body)
//			}
//		}
//	}))
//	defer ts.Close()
//
//	testUrl := ts.URL + "/testuri"
//	resp, err := c.DoRequest("POST", testUrl, `{"name":"test"}`, nil)
//	require.NoError(t, err)
//
//	assert.Equal(t, http.StatusOK, resp.StatusCode)
//}