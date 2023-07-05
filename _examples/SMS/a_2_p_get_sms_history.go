package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.A2PGetSmsHistoryParams{DestinationNumber:"12345678222", FromDate:2019-03-01 00:00:00}
	res, verr, err := client.SMS.A2PGetSmsHistory(params)
	fmt.Println(res, verr, err)
}