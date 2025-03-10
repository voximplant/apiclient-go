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

	FromDateRaw := time.Date(2021, 3, 17, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2021, 3, 17, 22, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	params := methods.RequestSmartQueueHistoryParams{ApplicationId: 1, SqQueueId: "1", FromDate: FromDate, ToDate: ToDate, ReportType: "service_level", MaxWaitingSec: 6}
	res, verr, err := client.SmartQueue.RequestSmartQueueHistory(params)
	fmt.Println(res, verr, err)
}
