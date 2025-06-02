package main

import (
	"fmt"
	"github.com/voximplant/apiclient-go/v2/config"
	"github.com/voximplant/apiclient-go/v2/methods"
	"github.com/voximplant/apiclient-go/v2/structure"
	"time"
)

func main() {
	voxConfig := config.NewConfig().WithEndpoint("https://api.voximplant.com/platform_api/").WithKeyPath("vox_key_jwt.json")
	client, err := methods.NewClient(voxConfig)
	if err != nil {
		panic(err)
	}

	StartAtRaw := time.Date(2023, 11, 13, 18, 0, 0, 0, time.UTC)
	StartAt := (*structure.Timestamp)(&StartAtRaw)

	params := methods.EditCallListTaskParams{ListId: 1, TaskId: 1, StartAt: StartAt, AttemptsLeft: 2, CustomData: "{\"phone\":\"555111222333\",\"name\":\"Mr.Fate\"}"}
	res, verr, err := client.CallLists.EditCallListTask(params)
	fmt.Println(res, verr, err)
}
