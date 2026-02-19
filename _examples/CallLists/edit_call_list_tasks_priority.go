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

	params := methods.EditCallListTasksPriorityParams{ListId: 558, Tasks: "[{\"task_id\":6034923,\"task_priority\":22},{\"task_id\":6034924,\"task_priority\":6},{\"task_uuid\":\"e074679c-c1ca-454e-b453-26611de42765\",\"task_priority\":11}]"}
	res, verr, err := client.CallLists.EditCallListTasksPriority(params)
	fmt.Println(res, verr, err)
}
