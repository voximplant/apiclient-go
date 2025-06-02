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

	FromDateRaw := time.Date(2019, 3, 1, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	params := methods.GetSmsHistoryParams{DestinationNumber: "12345678222", FromDate: FromDate}
	res, verr, err := client.SMS.GetSmsHistory(params)
	fmt.Println(res, verr, err)
}
