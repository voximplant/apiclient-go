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

	params := methods.EditCallListParams{ListId: 1000, IntervalSeconds: 1, NumAttempts: 2, MaxSimultaneous: 3, IpAddress: "127.0.0.1", Name: "awesome_list", Priority: 0, StartAt: "start_at", ServerLocation: "ru"}
	res, verr, err := client.CallLists.EditCallList(params)
	fmt.Println(res, verr, err)
}
