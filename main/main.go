package main

import (
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"github.com/sirupsen/logrus"
)

func main() {
	sc, _ := schedule.NewClient(&client.Config{ApiKey: "a871eb83-2d00-4b09-9fb9-7c134a369082", LogLevel: logrus.ErrorLevel})
	req := schedule.GetRequest{IdentifierType: schedule.Name, IdentifierValue: "ScheduleName"}

	res, err := sc.Get(nil, req)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
