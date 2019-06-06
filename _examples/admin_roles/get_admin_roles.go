package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	params := methods.GetAdminRolesParams{WithEntries:true, IncludedAdminUserId:"22", ShowingAdminUserId:11, Count:2}
	res, verr, err := client.AdminRoles.GetAdminRoles(params)
	fmt.Println(res, verr, err)
}