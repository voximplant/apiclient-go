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

	params := methods.SetKeyValueItemParams{Key: "key1", Value: "value1", ApplicationId: 1, Ttl: 864000}
	res, verr, err := client.KeyValueStorage.SetKeyValueItem(params)
	fmt.Println(res, verr, err)
}
