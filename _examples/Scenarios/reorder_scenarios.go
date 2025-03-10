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

	params := methods.ReorderScenariosParams{RuleId: 2, ScenarioId: "17;15;20"}
	res, verr, err := client.Scenarios.ReorderScenarios(params)
	fmt.Println(res, verr, err)
}
