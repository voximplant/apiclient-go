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
	WithDetailsRaw := true
	WithDetails := &WithDetailsRaw
	params := methods.GetAccountDocumentsParams{WithDetails: WithDetails}
	res, verr, err := client.Accounts.GetAccountDocuments(params)
	fmt.Println(res, verr, err)
}
