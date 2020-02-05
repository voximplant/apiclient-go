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
	params := methods.TransferMoneyToChildAccountParams{ChildAccountId:"1", Amount:-10000000, Currency:"USD", StrictMode:false}
	res, verr, err := client.Accounts.TransferMoneyToChildAccount(params)
	fmt.Println(res, verr, err)
}