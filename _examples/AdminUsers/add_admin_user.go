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
	params := methods.AddAdminUserParams{NewAdminUserName:"adm1", AdminUserDisplayName:"adm1", NewAdminUserPassword:"1234567", AdminRoleId:"1"}
	res, verr, err := client.AdminUsers.AddAdminUser(params)
	fmt.Println(res, verr, err)
}