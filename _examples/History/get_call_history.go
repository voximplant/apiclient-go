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
	FromDateRaw := time.Date(2020, 2, 25, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2020, 2, 26, 0, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	WithCallsRaw := true
	WithCalls := &WithCallsRaw
	WithRecordsRaw := true
	WithRecords := &WithRecordsRaw

	params := methods.GetCallHistoryParams{FromDate: FromDate, ToDate: ToDate, Timezone: "Etc/GMT", WithCalls: WithCalls, WithRecords: WithRecords, Count: 1}
	res, verr, err := client.History.GetCallHistory(params)
	fmt.Println(res, verr, err)
}
