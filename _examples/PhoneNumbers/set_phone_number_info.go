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

	AutoChargeRaw := true
	AutoCharge := &AutoChargeRaw
	params := methods.SetPhoneNumberInfoParams{PhoneId: "1", AutoCharge: AutoCharge}
	res, verr, err := client.PhoneNumbers.SetPhoneNumberInfo(params)
	fmt.Println(res, verr, err)
}
