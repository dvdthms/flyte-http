# flyte-http
A HTTP integration pack for Flyte

## Command

This pack provides the `DoRequest` command, which allows the user to make arbitrary HTTP requests.

The command takes the following arguments:

**Method**
The HTTP method used for the request. E.g: "POST", "GET", "PUT", etc.

**URL**
The URL you desire to send the request to. E.g: "http://example.com"

**Headers**
Any headers you wish to provide to the request, listed in key-value pairs. E.g: "Key":"value"

**Body**
The data you wish to use as the request body. E.g: "TestBody"

**Timeout**
The timeout value you wish to specify the http client to have, measured in Milliseconds.
E.g: a 10ms value would result in a 10ms timeout, or 0.01 seconds.
If not specified, will default to 0, which is equivalent to no timeout.

<p>
### Command Output

In the event of success, the command will return the following information:

**Status Code**
The HTTP status code of the response. E.g: 200, 404, 500, etc.

**Headers**
Any headers that were part of the response. E.g: "Content-Type":"text/plain"

**Body**
Any data that formed the response body. For example, the contents of a GET request may be ``{"ExampleBody"}`

*NB: If the content-type of the response is JSON, the body will return a JSON object.
For any other content-type, the body will consist of a base-64 encoded string.*

<p>

In the event of an error, the command will return:

**Error**
Information pertaining to the error which occurred.

The Method, URL, Headers & Body that were provided as part of the input will also be returned, so that they may be anaylsed
for any potential mistakes.



## Environment Configuration


| env. variable             | default                      |description                      |
|---------------------------|------------------------------|---------------------------------|
| FLYTE_API                 |                              | The FLYTE API endpoint to use.  |


## Build and run

### GO
- must have [dep](https://github.com/golang/dep) installed

```
    dep ensure
    go test ./...
    go build && ./flyte-http
```

### Docker

```
docker build -t <image>:latest .
docker run --rm <image>:latest
```