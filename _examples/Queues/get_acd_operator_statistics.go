package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
	"github.com/voximplant/apiclient-go/structure"
	"time"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	FromDateRaw := time.Date(2021, 4, 8, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2021, 4, 10, 0, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	params := methods.GetACDOperatorStatisticsParams{FromDate: FromDate, ToDate: ToDate, AcdQueueId: "54", Aggregation: "day"}
	res, verr, err := client.Queues.GetACDOperatorStatistics(params)
	fmt.Println(res, verr, err)
}
