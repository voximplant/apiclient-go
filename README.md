# Voximplant API client library

#### Version 1.6.0

## Prerequisites

In order to use Voximplant SDK for the Go programming language, you need the following:

1. A developer account. If you don't have one, [sign up here](https://voximplant.com/sign-up/).
2. A private API key. There are 2 options to obtain it:
    1. Either generate it in the [Voximplant Control panel](https://manage.voximplant.com/settings/service_accounts)
    1. Or call the [CreateKey](https://voximplant.com/docs/references/httpapi/managing_role_system#createkey) HTTP API
       method with the
       specified [authentication parameters](https://voximplant.com/docs/references/httpapi/auth_parameters). You'll
       receive a response with the __result__ field in it. Save the __result__ value in a file (since we don't store the
       keys, save it securely on your side).

## Getting started

* The best way to start is to use `go get` to add the SDK to your Go Workspace or application using Go modules:

  ```sh
  go get github.com/voximplant/apiclient-go
  ```

* Without Go Modules (or in a GOPATH with Go 1.11 or 1.12), use `go get` with the `/...` suffix to retrieve all of the
  SDK's dependencies:

  ```sh
  go get github.com/voximplant/apiclient-go/...
  ```

### Example

This example shows how you can use the API client.

```go
package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/config"
	"github.com/voximplant/apiclient-go/methods"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	// Create a Voximplant client using the Config value
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic("failed to create client, " + err.Error())
	}
	// Build input parameters
	params := methods.GetSubscriptionPriceParams{SubscriptionTemplateType: "SIP_REGISTRATION"}
	// Send the request and get the response or error back
	res, verr, err := client.Accounts.GetSubscriptionPrice(params)

	fmt.Println(res, verr, err)
}
```
