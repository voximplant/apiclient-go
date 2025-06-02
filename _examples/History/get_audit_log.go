package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/v2/config"
	"github.com/voximplant/apiclient-go/v2/methods"
	"github.com/voximplant/apiclient-go/v2/structure"
	"time"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}
	FromDateRaw := time.Date(2018, 2, 1, 0, 0, 0, 0, time.UTC)
	FromDate := (*structure.Timestamp)(&FromDateRaw)
	ToDateRaw := time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)
	ToDate := (*structure.Timestamp)(&ToDateRaw)

	params := methods.GetAuditLogParams{FromDate: FromDate, ToDate: ToDate, FilteredCmd: "BindSkill;AddSkill;DelSkill", AdvancedFilters: "152", Count: 3}
	res, verr, err := client.History.GetAuditLog(params)
	fmt.Println(res, verr, err)
}
