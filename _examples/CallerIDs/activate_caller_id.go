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

	params := methods.ActivateCallerIDParams{CalleridId: 1, VerificationCode: "12345"}
	res, verr, err := client.CallerIDs.ActivateCallerID(params)
	fmt.Println(res, verr, err)
}
