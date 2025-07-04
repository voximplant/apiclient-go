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

	params := methods.BindSkillParams{SkillId: "1;6", UserId: "all", AcdQueueId: "11;12"}
	res, verr, err := client.Skills.BindSkill(params)
	fmt.Println(res, verr, err)
}
