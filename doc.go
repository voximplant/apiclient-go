/*
Prerequisites

In order to use Voximplant SDK for the Go programming language, you need the following:

A developer account. If you don't have one, sign up here.
A private API key. There are 2 options to obtain it:
Either generate it in the Voximplant Control panel
Or call the CreateKey HTTP API method with the specified authentication parameters. You'll receive a response with the result field in it. Save the result value in a file (since we don't store the keys, save it securely on your side).

Getting started

The best way to start is to use go get to add the SDK to your Go Workspace or application using Go modules:

	go get github.com/voximplant/apiclient-go

Without Go Modules (or in a GOPATH with Go 1.11 or 1.12), use go get with the /... suffix to retrieve all of the SDK's dependencies:

	go get github.com/voximplant/apiclient-go/...

Example

This example shows how you can use the API client.

	package main

	import (
		"fmt"
		"github.com/voximplant/apiclient-go/config"
		"github.com/voximplant/apiclient-go/methods"
	)

	func main() {
		voxConfig := config.NewConfig().WithKeyPath("vox_key_jwt.json")
		// Create a Voximplant client using the Config value
		client, err := methods.NewClient(voxConfig)
		if err != nil {
			panic("failed to create client, " + err.Error())
		}
		// Build input parameters
		params := methods.GetSubscriptionPriceParams{SubscriptionTemplateType:"SIP_REGISTRATION"}
		// Send the request and get the response or error back
		res, verr, err := client.Accounts.GetSubscriptionPrice(params)

		fmt.Println(res, verr, err)
	}
*/
package sdk