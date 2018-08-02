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
	"fmt"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"net/http"
)

func DoRequestCommand() flyte.Command {

	return flyte.Command{
		Name:    "DoRequest",
		Handler: doRequestHandler,
		OutputEvents: []flyte.EventDef{
			doRequestSuccessEventDef,
			doRequestErrorEventDef,
		},
	}
}

func doRequestHandler(rawInput json.RawMessage) flyte.Event {

	input := DoRequestInput{}
	if err := json.Unmarshal(rawInput, &input); err != nil {
		return flyte.NewFatalEvent(fmt.Sprintf("could not unmarshall 'doRequest' rawInput into json: %v", err))
	}

	client := &http.Client{Timeout: input.Timeout} //change to input.Timeout * time.MilliSecond?

	output, err := DoRequest(client, input)
	if err != nil {
		return newDoRequestErrorEvent(input, err)
	}
	return newDoRequestSuccessEvent(output)
}
