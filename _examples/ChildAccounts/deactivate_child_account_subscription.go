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
	params := methods.DeactivateChildAccountSubscriptionParams{SubscriptionId:20, ChildAccountId:10, SubscriptionFinishDate:2019-09-29 00:00:00}
	res, verr, err := client.ChildAccounts.DeactivateChildAccountSubscription(params)
	fmt.Println(res, verr, err)
}