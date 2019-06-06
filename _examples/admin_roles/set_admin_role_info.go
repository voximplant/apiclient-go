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
	params := methods.SetAdminRoleInfoParams{AdminRoleId:1, EntryModificationMode:"set", AllowedEntries:"all", DeniedEntries:"DelUser;DelApplication"}
	res, verr, err := client.AdminRoles.SetAdminRoleInfo(params)
	fmt.Println(res, verr, err)
}