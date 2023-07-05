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
	params := methods.RequestSmartQueueHistoryParams{ApplicationId:1, SqQueueId:"1", FromDate:2021-03-17 00:00:00, ToDate:2021-03-17 22:00:00, ReportType:"service_level", MaxWaitingSec:6}
	res, verr, err := client.SmartQueue.RequestSmartQueueHistory(params)
	fmt.Println(res, verr, err)
}