package client

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"net/http/httptest"
	"net/http"
)

var (
	testHeaders = map[string][]string{"Testheader": {"Test"}} //Have multiple? + slice len > 1
)

func TestDoRequestShouldReturnNoErrorIfRequestIsSuccessful(t *testing.T) {
	c := NewFlyteHttpClient()

	_, err := c.DoRequest("POST", "http://example.com", `{"name":"test"}`, nil)
	require.NoError(t, err)
}

func TestDoRequestShouldReturnErrorIfMethodIsMissing(t *testing.T) {
	c := NewFlyteHttpClient()

	_, err := c.DoRequest("", "http://example.com", `{"name":"test"}`, nil)
	require.Error(t, err)
	assert.Equal(t, "error: method not set", err.Error())
}

func TestDoRequestShouldReturnErrorIfHttpMethodIsNotValid(t *testing.T) {
	c := NewFlyteHttpClient()

	_, err := c.DoRequest("invalidmethod", "http://example.com", `{"name":"test"}`, nil)
	require.Error(t, err)
	assert.Equal(t, "error: invalid method supplied", err.Error())
}

func TestDoRequestShouldReturnErrorIfUrlIsMissing(t *testing.T) {
	c := NewFlyteHttpClient()
	_, err := c.DoRequest("POST", "", `{"name":"test"}`, nil)

	require.Error(t, err)
	assert.Equal(t, "error: url not set", err.Error())
}

func TestDoRequestShouldReturnErrorIfUrlIsMalformed(t *testing.T) {
	c := NewFlyteHttpClient()

	malformedTestUrl := "malformedurl"
	_, err := c.DoRequest("POST", malformedTestUrl, `{"name":"test"}`, nil)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported protocol scheme")
}

//Test to check if Req is correct method, url, body, etc?

func TestDoRequestShouldHaveRequestWithPopulatedHeadersIfHeadersArePassed(t *testing.T) {
	c := NewFlyteHttpClient()

	var req *http.Request
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req = r
	}))
	defer ts.Close()

	testUrl := ts.URL + "/testuri"
	_, err := c.DoRequest("POST", testUrl, `{"name":"test"}`, testHeaders)
	require.NoError(t, err)
	require.NotNil(t, req)

	_, isPresent := req.Header["Testheader"]
	assert.True(t, isPresent)
}
