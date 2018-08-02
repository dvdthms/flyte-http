package main

import (
	api "github.com/HotelsDotCom/flyte-client/client"
	"github.com/HotelsDotCom/flyte-client/flyte"
	"github.com/HotelsDotCom/flyte-http/flytehttp"
	"net/url"
	"time"
)

const packDefHelpUrl = "https://github.com/HotelsDotCom/flyte-http/blob/master/README.md"

func main() {
	helpUrl, _ := url.Parse(packDefHelpUrl)
	packDef := flyte.PackDef{
		Name:    "Slack",
		HelpURL: helpUrl,
		Commands: []flyte.Command{
			flytehttp.DoRequestCommand(),
		},
	}

	pack := flyte.NewPack(packDef, api.NewClient(ApiHost(), 10*time.Second))
	pack.Start()
	select {}
}
