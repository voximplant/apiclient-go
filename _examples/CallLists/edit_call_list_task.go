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
	params := methods.EditCallListTaskParams{ListId:1, TaskId:1, StartAt:2023-11-13 18:00:00, AttemptsLeft:2, CustomData:"{"phone":"555111222333","name":"Mr.Fate"}"}
	res, verr, err := client.CallLists.EditCallListTask(params)
	fmt.Println(res, verr, err)
}