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
	params := methods.GetAdminUsersParams{WithAccessEntries:true, Count:2}
	res, verr, err := client.AdminUsers.GetAdminUsers(params)
	fmt.Println(res, verr, err)
}