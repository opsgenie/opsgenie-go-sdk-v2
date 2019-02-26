package main

import (
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/heartbeat"
)

func main() {
	sub := &client.MetricSubscriber{
		Process: Process,
	}
	config := &client.Config{
		ApiKey: "e567c791-3ff4-43b4-8878-9386d7a3bd68",
	}
	sub.Register(client.API)
	sub.Register(client.HTTP)
	sub.Register(client.SDK)

	hc, _ := heartbeat.NewClient(config)

	res, err := hc.Ping(nil, "ads")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func Process(metric client.Metric) interface{} {
	fmt.Printf("received metric with type: %s %+v", metric.Type(), metric)
	fmt.Println()
	return nil
}
