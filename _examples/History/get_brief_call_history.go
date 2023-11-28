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
	params := methods.GetBriefCallHistoryParams{FromDate:2020-02-25 00:00:00, ToDate:2020-02-26 00:00:00, Timezone:"Etc/GMT", Output:"cvs", IsAsync:true}
	res, verr, err := client.History.GetBriefCallHistory(params)
	fmt.Println(res, verr, err)
}