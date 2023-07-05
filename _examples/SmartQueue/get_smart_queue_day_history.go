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
	params := methods.GetSmartQueueDayHistoryParams{ApplicationId:1, ReportType:"sum_agents_online_time"}
	res, verr, err := client.SmartQueue.GetSmartQueueDayHistory(params)
	fmt.Println(res, verr, err)
}