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
	params := methods.GetChildrenAccountsParams{ChildAccountEmail:"mychild@gmail.com"}
	res, verr, err := client.Accounts.GetChildrenAccounts(params)
	fmt.Println(res, verr, err)
}