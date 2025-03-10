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
	FromDateRaw := time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2014, 1, 1, 0, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	params := methods.GetTransactionHistoryParams{FromDate: FromDate, ToDate: ToDate, Timezone: "Etc/GMT", TransactionType: "gift;money_distribution", Count: 3}
	res, verr, err := client.History.GetTransactionHistory(params)
	fmt.Println(res, verr, err)
}
