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

	params := methods.UpdateKeyParams{KeyId: "ab98c70e-573e-4446-9af9-105269dfafca", Description: "test_desc"}
	res, verr, err := client.RoleSystem.UpdateKey(params)
	fmt.Println(res, verr, err)
}
