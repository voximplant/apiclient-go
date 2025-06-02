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

	params := methods.SQUnbindAgentParams{ApplicationId: 1, SqQueueId: "all", UserId: "1"}
	res, verr, err := client.SmartQueue.SQUnbindAgent(params)
	fmt.Println(res, verr, err)
}
