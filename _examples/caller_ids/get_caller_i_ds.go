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
	params := methods.GetCallerIDsParams{CalleridNumber:"79997770044"}
	res, verr, err := client.CallerIds.GetCallerIDs(params)
	fmt.Println(res, verr, err)
}