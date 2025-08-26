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

	params := methods.AddWABPhoneNumberParams{WabPhoneNumber: "12126367890", VoicePassword: "abc", ApplicationId: 1234, RuleId: 5678}
	res, verr, err := client.WABPhoneNumbers.AddWABPhoneNumber(params)
	fmt.Println(res, verr, err)
}
