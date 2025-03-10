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

	WithSqStatusesRaw := true
	WithSqStatuses := &WithSqStatusesRaw
	HandleCallsRaw := false
	HandleCalls := &HandleCallsRaw
	params := methods.SQGetAgentsParams{ApplicationId: 1, WithSqStatuses: WithSqStatuses, HandleCalls: HandleCalls}
	res, verr, err := client.SmartQueue.SQGetAgents(params)
	fmt.Println(res, verr, err)
}
