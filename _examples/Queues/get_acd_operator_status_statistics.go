package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/v2/config"
	"github.com/voximplant/apiclient-go/v2/methods"
	"github.com/voximplant/apiclient-go/v2/structure"
	"time"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	FromDateRaw := time.Date(2019, 5, 20, 11, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2019, 5, 20, 13, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	params := methods.GetACDOperatorStatusStatisticsParams{FromDate: FromDate, ToDate: ToDate, AcdStatus: "READY;ONLINE", UserId: "all", Aggregation: "hour", Group: "user"}
	res, verr, err := client.Queues.GetACDOperatorStatusStatistics(params)
	fmt.Println(res, verr, err)
}
