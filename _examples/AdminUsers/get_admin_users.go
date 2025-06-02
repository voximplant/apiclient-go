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
	WithAccessEntriesRaw := true
	WithAccessEntries := &WithAccessEntriesRaw

	params := methods.GetAdminUsersParams{WithAccessEntries: WithAccessEntries, Count: 2}
	res, verr, err := client.AdminUsers.GetAdminUsers(params)
	fmt.Println(res, verr, err)
}
