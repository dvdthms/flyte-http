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
	"encoding/json"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//Add tests for all individual Methods? or "supported" ones? - i.e: POST, PATCH, GET

func TestDoRequestCommandShouldReturnSuccessEventIfDoRequestIsSuccessful(t *testing.T) { //Repurpose this to be specific to POST?

	expectedHeader := http.Header{"Content-Type": {"text/plain"}}
	var expectedResponse = doRequestOutputPayload{
		StatusCode: http.StatusOK,
		Body:       "",
		Header:     expectedHeader,
	}

	command := DoRequestCommand()
	outputEvent := command.Handler(json.RawMessage(`{"method":"POST", "url":"http://example.com", "body":"test", "headers":{"Testheader":["Test"]}}, "timeout":"10"`))
	assert.Equal(t, doRequestSuccessEventDef, outputEvent.EventDef)

	var payload = outputEvent.Payload.(doRequestOutputPayload)
	assert.Equal(t, expectedResponse.StatusCode, payload.StatusCode)
	assert.Equal(t, expectedResponse.Body, payload.Body)
	assert.Contains(t, payload.Header, expectedResponse.Header)
}

func TestDoRequestCommandShouldReturnFatalFlyteEventIfJsonCannotBeUnmarshalled(t *testing.T) {
	command := DoRequestCommand()
	outputEvent := command.Handler(json.RawMessage(`invalidjson`))

	assert.Equal(t, flyte.NewFatalEvent(t.Errorf).EventDef, outputEvent.EventDef)
	assert.Contains(t, outputEvent.Payload, "could not unmarshall 'doRequest' rawInput into json")
}

func TestDoRequestCommandShouldReturnErrorEventIfMethodIsEmpty(t *testing.T) {
	command := DoRequestCommand()
	outputEvent := command.Handler(json.RawMessage(`{"method":"", "url":"http://example.com", "body":"test", "headers":{"Testheader":["Test"]}, "timeout":"10"}`))

	var payload = outputEvent.Payload.(doRequestErrorOutputPayload)
	assert.Equal(t, doRequestErrorEventDef, outputEvent.EventDef)
	assert.Equal(t, "no method provided", payload.Error)
}

func TestDoRequestCommandShouldReturnErrorEventIfUrlIsEmpty(t *testing.T) {
	command := DoRequestCommand()
	outputEvent := command.Handler(json.RawMessage(`{"method":"POST", "url":"", "body":"test", "headers":{"Testheader":["Test"]}, "timeout":"10"}`))

	var payload = outputEvent.Payload.(doRequestErrorOutputPayload)
	assert.Equal(t, doRequestErrorEventDef, outputEvent.EventDef)
	assert.Contains(t, payload.Error, "unsupported protocol scheme \"\"")
}

func TestDoRequestCommandShouldReturnErrorEventIfDoRequestReturnsError(t *testing.T) {
	command := DoRequestCommand()
	outputEvent := command.Handler(json.RawMessage(`{"method":"GET", "url":"://", "body":"test", "headers":{"Testheader":["Test"]}, "timeout":"10"}`))

	var payload = outputEvent.Payload.(doRequestErrorOutputPayload)
	assert.Equal(t, doRequestErrorEventDef, outputEvent.EventDef)
	assert.Equal(t, "", payload.Error)
}
