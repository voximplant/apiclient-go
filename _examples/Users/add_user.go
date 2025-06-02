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

	params := methods.AddUserParams{UserName: "GordonFreeman", UserDisplayName: "GordonFreeman", UserPassword: "1234567", ApplicationId: 1}
	res, verr, err := client.Users.AddUser(params)
	fmt.Println(res, verr, err)
}
