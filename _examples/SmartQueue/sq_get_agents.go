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

	WithSqStatusesRaw := true
	WithSqStatuses := &WithSqStatusesRaw
	params := methods.SQGetAgentsParams{ApplicationId: 1, WithSqStatuses: WithSqStatuses}
	res, verr, err := client.SmartQueue.SQGetAgents(params)
	fmt.Println(res, verr, err)
}
