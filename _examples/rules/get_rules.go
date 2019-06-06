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
	params := methods.GetRulesParams{ApplicationId:1, Template:"74951234567", WithScenarios:true, Count:1}
	res, verr, err := client.Rules.GetRules(params)
	fmt.Println(res, verr, err)
}