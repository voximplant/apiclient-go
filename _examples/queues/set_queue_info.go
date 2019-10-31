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
	params := methods.SetQueueInfoParams{AcdQueueId:1, NewAcdQueueName:"support"}
	res, verr, err := client.Queues.SetQueueInfo(params)
	fmt.Println(res, verr, err)
}