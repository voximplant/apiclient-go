package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.GetCallHistoryParams{FromDate:2012-01-01 00:00:00, ToDate:2014-01-01 00:00:00, Timezone:"Etc/GMT", Count:1}
	res, verr, err := client.History.GetCallHistory(params)
	fmt.Println(res, verr, err)
}