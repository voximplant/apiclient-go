package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/v2/config"
	"github.com/voximplant/apiclient-go/v2/methods"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}

	params := methods.GetSmartQueueRealtimeMetricsParams{ApplicationId: 1, ReportType: "sum_agents_dialing_time"}
	res, verr, err := client.SmartQueue.GetSmartQueueRealtimeMetrics(params)
	fmt.Println(res, verr, err)
}
