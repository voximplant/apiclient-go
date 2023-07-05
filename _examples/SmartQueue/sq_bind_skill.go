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
	params := methods.SQBindSkillParams{ApplicationId:1, UserId:"all", SqSkills:[{"sq_skill_id":1,"sq_skill_level":1},{"sq_skill_id":2,"sq_skill_level":5}]}
	res, verr, err := client.SmartQueue.SQBindSkill(params)
	fmt.Println(res, verr, err)
}