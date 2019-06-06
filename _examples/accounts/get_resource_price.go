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
	params := methods.GetResourcePriceParams{ResourceType:"PSTNOUT", ResourceParam:"79263332211"}
	res, verr, err := client.Accounts.GetResourcePrice(params)
	fmt.Println(res, verr, err)
}