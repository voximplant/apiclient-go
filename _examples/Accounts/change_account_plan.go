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

	params := methods.ChangeAccountPlanParams{PlanType: "IM", PlanSubscriptionTemplateId: 3}
	res, verr, err := client.Accounts.ChangeAccountPlan(params)
	fmt.Println(res, verr, err)
}
