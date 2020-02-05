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
	params := methods.AddSubUserParams{NewSubuserName:"test_login", NewSubuserPassword:"test_pass"}
	res, verr, err := client.RoleSystem.AddSubUser(params)
	fmt.Println(res, verr, err)
}