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

	BindRaw := false
	Bind := &BindRaw
	params := methods.BindSipRegistrationParams{UserId: 123, Bind: Bind}
	res, verr, err := client.SIPRegistration.BindSipRegistration(params)
	fmt.Println(res, verr, err)
}
