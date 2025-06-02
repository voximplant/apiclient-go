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

	params := methods.SQUnbindSkillParams{ApplicationId: 1, UserId: "1", SqSkillId: "1"}
	res, verr, err := client.SmartQueue.SQUnbindSkill(params)
	fmt.Println(res, verr, err)
}
