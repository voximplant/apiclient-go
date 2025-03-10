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

	HandleCallsRaw := true
	HandleCalls := &HandleCallsRaw
	params := methods.SQSetAgentInfoParams{ApplicationId: 1, UserId: "2", HandleCalls: HandleCalls}
	res, verr, err := client.SmartQueue.SQSetAgentInfo(params)
	fmt.Println(res, verr, err)
}
