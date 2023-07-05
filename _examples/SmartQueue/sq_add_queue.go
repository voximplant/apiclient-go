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
	params := methods.SQAddQueueParams{ApplicationId:1, SqQueueName:"smartQueue1", CallAgentSelection:"MOST_QUALIFIED", CallTaskSelection:"MAX_WAITING_TIME"}
	res, verr, err := client.SmartQueue.SQAddQueue(params)
	fmt.Println(res, verr, err)
}