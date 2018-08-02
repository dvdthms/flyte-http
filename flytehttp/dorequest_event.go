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
	"github.com/HotelsDotCom/flyte-client/flyte"
	"net/http"
	"time"
)

type DoRequestInput struct {
	Method  string        `json:"method"`
	URL     string        `json:"url"`
	Header http.Header   `json:"header"`
	Body    string        `json:"body"`
	Timeout time.Duration `json:"timeout,string"`
}

type doRequestOutputPayload struct {
	StatusCode int         `json:"statusCode"`
	Header     http.Header `json:"header"`
	Body       interface{} `json:"body"`
}

type doRequestErrorOutputPayload struct {
	DoRequestInput
	Error string `json:"error"`
}

var (
	doRequestSuccessEventDef = flyte.EventDef{Name: "DoRequestSuccess"}
	doRequestErrorEventDef   = flyte.EventDef{Name: "DoRequestFailed"}
)

func newDoRequestSuccessEvent(payload *doRequestOutputPayload) flyte.Event {
	return flyte.Event{
		EventDef: doRequestSuccessEventDef,
		Payload:  payload,
	}
}

func newDoRequestErrorEvent(input DoRequestInput, err error) flyte.Event {
	return flyte.Event{
		EventDef: doRequestErrorEventDef,
		Payload: doRequestErrorOutputPayload{
			DoRequestInput: input,
			Error:          err.Error(),
		},
	}
}
