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
	params := methods.AddPushCredentialParams{PushProviderName:"GOOGLE", SenderId:"704777431520", ServerKey:"AAAAjM-LQsc:APA91bGyCb5WhcGtaM-RaOI1GqWps1Uh9K-YoY75HIBy-En-4piH4c6_50gIEbSaCfuDrsLNfyZCvteiu6EjxA_rEBOvlc4xZ30uiGgbuM_jdT6y6Ku55OwnCyIxRNznvmx1jkkLexSU"}
	res, verr, err := client.PushCredentials.AddPushCredential(params)
	fmt.Println(res, verr, err)
}