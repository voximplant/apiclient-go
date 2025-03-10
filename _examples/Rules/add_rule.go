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

	params := methods.AddRuleParams{ApplicationId: 1, RuleName: "allowall", RulePattern: ".*"}
	res, verr, err := client.Rules.AddRule(params)
	fmt.Println(res, verr, err)
}
