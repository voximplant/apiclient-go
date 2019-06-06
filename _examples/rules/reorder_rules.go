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
	params := methods.ReorderRulesParams{RuleId:"1;7;3"}
	res, verr, err := client.Rules.ReorderRules(params)
	fmt.Println(res, verr, err)
}