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
	params := methods.SetChildAccountInfoParams{ChildAccountId:"1321", Active:false}
	res, verr, err := client.Accounts.SetChildAccountInfo(params)
	fmt.Println(res, verr, err)
}