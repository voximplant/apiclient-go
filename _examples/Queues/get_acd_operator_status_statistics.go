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
	params := methods.GetACDOperatorStatusStatisticsParams{FromDate:2019-05-20 11:00:00, ToDate:2019-05-20 13:00:00, AcdStatus:"READY;ONLINE", UserId:"all", Aggregation:"hour", Group:"user"}
	res, verr, err := client.Queues.GetACDOperatorStatusStatistics(params)
	fmt.Println(res, verr, err)
}