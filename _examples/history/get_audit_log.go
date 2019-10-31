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
	params := methods.GetAuditLogParams{FromDate:2018-02-01 00:00:00, ToDate:2018-03-01 00:00:00, FilteredCmd:"BindSkill;AddSkill;DelSkill", AdvancedFilters:"152", Count:3}
	res, verr, err := client.History.GetAuditLog(params)
	fmt.Println(res, verr, err)
}