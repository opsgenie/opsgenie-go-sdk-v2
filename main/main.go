package main

import (
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"github.com/sirupsen/logrus"
)

func main() {
	config := &client.Config{ApiKey: "a871eb83-2d00-4b09-9fb9-7c134a369082"}
	config.LogLevel = logrus.DebugLevel
	scheduleClient, err := schedule.NewClient(config)
	if err != nil {
		return
	}

	/*req := schedule.CreateRequest{Name:"asd"};
	tr := og.TimeRestriction{Type:og.WeekdayAndTimeOfDay}
	tr.WithRestrictions(og.Restriction{StartDay:og.Friday, StartMin:5, StartHour: 5,  EndHour:5, EndDay:og.Saturday, EndMin:6})
	req.WithRotation(og.Rotation{Name:"rot", Type:og.Hourly, StartDate:"2017-02-06T05:00:00Z", EndDate:"2017-02-23T06:00:00Z"}.WithParticipants(og.Participant{Name:"Salesasd", Type: og.Team}).WithTimeRestriction(tr))
	res , err := scheduleClient.Create(nil, req)
	if err != nil {
		fmt.Println(res)
	}*/

	/*req := schedule.GetRequest{IdentifierType:schedule.Id, IdentifierValue:"05f06ce9-ca42-4aac-ae0e-7f989eca3f47"};

	res , err := scheduleClient.Get(nil, req)
	if err != nil {
		fmt.Println(res)
	}*/

	/*req := schedule.UpdateRequest{Name:"sdk-sche", IdentifierType:schedule.Name, IdentifierValue:"asd"};
	tr := og.TimeRestriction{Type:og.WeekdayAndTimeOfDay}
	tr.WithRestrictions(og.Restriction{StartDay:og.Friday, StartMin:15, StartHour: 15,  EndHour:15, EndDay:og.Saturday, EndMin:6})
	req.WithRotation(og.Rotation{Name:"rot2", Type:og.Hourly, StartDate:"2017-02-06T05:00:00Z", EndDate:"2018-02-23T06:00:00Z"}.WithParticipants(og.Participant{Name:"Marketing", Type: og.Team}).WithTimeRestriction(tr))

	res , err := scheduleClient.Update(nil, req)
	if err != nil {
		fmt.Println(res)
	}*/

	/*req := schedule.DeleteRequest{IdentifierType:schedule.Name, IdentifierValue:"sdk-sche"};

	res , err := scheduleClient.Delete(nil, req)
	if err != nil {
		fmt.Println(res)
	}*/

	/*req := schedule.ListRequest{Expand:false};

	res , err := scheduleClient.List(nil, req)
	if err != nil {
		fmt.Println(res)
	}*/

	req := &schedule.GetTimelineRequest{IdentifierValue: "a29ff1db-aa76-4aa5-84ef-848971fa1fcd", IdentifierType: schedule.Id}
	req.WithExpands(schedule.Base, schedule.Forwarding, schedule.Override)

	res, err := scheduleClient.GetTimeline(nil, *req)
	if err != nil {
		fmt.Println(res)
	}
}
