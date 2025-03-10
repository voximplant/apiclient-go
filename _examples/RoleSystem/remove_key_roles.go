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

	params := methods.RemoveKeyRolesParams{KeyId: "ab81c90e-543e-4446-9af9-105269dfafca"}
	res, verr, err := client.RoleSystem.RemoveKeyRoles(params)
	fmt.Println(res, verr, err)
}
