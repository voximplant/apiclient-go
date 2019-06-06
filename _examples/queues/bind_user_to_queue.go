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
	params := methods.BindUserToQueueParams{Bind:true, ApplicationId:1, UserId:"12;987;456", AcdQueueName:"myqueue"}
	res, verr, err := client.Queues.BindUserToQueue(params)
	fmt.Println(res, verr, err)
}