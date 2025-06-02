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

	params := methods.GetNewPhoneNumbersParams{CountryCode: "RU", PhoneCategoryName: "GEOGRAPHIC", PhoneRegionId: 1, Count: 2}
	res, verr, err := client.PhoneNumbers.GetNewPhoneNumbers(params)
	fmt.Println(res, verr, err)
}
