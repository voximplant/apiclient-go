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
	params := methods.SetKeyRolesParams{KeyId:"ab81c76e-573e-4046-9af9-105269dfafca", RoleId:"1"}
	res, verr, err := client.RoleSystem.SetKeyRoles(params)
	fmt.Println(res, verr, err)
}