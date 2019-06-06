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
	params := methods.DelAdminUserParams{RequiredAdminUserId:"1"}
	res, verr, err := client.AdminUsers.DelAdminUser(params)
	fmt.Println(res, verr, err)
}