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
	BindRaw := true
	Bind := &BindRaw

	params := methods.BindUserToQueueParams{Bind: Bind, ApplicationId: 1, UserId: "12;987;456", AcdQueueName: "myqueue"}
	res, verr, err := client.Queues.BindUserToQueue(params)
	fmt.Println(res, verr, err)
}
