package main

import (
	"context"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"time"
)

func main() {
	//var retries int32
	/*alertTest := alert.NewClient(client.Config{
		ApiKey:         "5d2891dc-8e22-403c-a124-0becc4e4c460", //8f8e3be1-9684-4bb9-9ebb-46015d0c9952
		OpsGenieAPIURL: "https://localhost:9002",
		/*Backoff:  func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		atomic.AddInt32(&retries, 1)
		return time.Millisecond * 1
	},*/
	/*RetryPolicy: func(ctx context.Context, resp *http.Response, err error) (b bool, e error) {
		if ctx.Err() != nil {
			return false, ctx.Err()
		}

		if err != nil {
			return true, err
		}
		// Check the response code. We retry on 500-range responses to allow
		// the server time to recover, as 500's are typically not permanent
		// errors and may relate to outages on the server side. This will catch
		// invalid response codes as well, like 0 and 999.
		if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
			return true, nil
		}
		if resp.StatusCode == 429  {
			return true, nil
		}

		return false, nil
	},*/

	/*})*/

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	/*heartbeatClient := heartbeat.NewClient(client.Config{
	ApiKey:         "a871eb83-2d00-4b09-9fb9-7c134a369082",
	OpsGenieAPIURL: "https://api.opsgenie.com/v2",
	}, ctx)

	pr ,err := heartbeatClient.Ping(heartbeat.PingRequest{HeartbeatName:"asd"})

	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(pr)
	}*/

	alertTest, err := alert.NewClient(client.Config{
		ApiKey: "8f8e3be1-9684-4bb9-9ebb-46015d0c9950", //"73fbb690-c912-41f0-a773-06a8719db6d6", //5d2891dc-8e22-403c-a124-0becc4e4c460
		//OpsGenieAPIURL: "https://localhost:9002",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	/*cr, err :=alertTest.Create(ctx,alert.CreateAlertRequest{
		Alias:       "aliasx",
		Description: "alert description2",
		Actions:     []string{"action12", "action22"},
		Tags:        []string{"tag12", "tag22"},
		Details: map[string]string{
			"key":  "value2",
			"key2": "value22",
		},
		Entity:   "entity2",
		Source:   "source2",
		Priority: alert.P1,
		User:     "mbtekinsen@gmail.com",
		Note:     "alert note2",})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Create request ID: " + cr.RequestID)
		fmt.Printf("Took: %f\n" , cr.ResponseTime)

	}*/

	/*response, err := alertTest.List(ctx, alert.ListAlertRequest{})

	fmt.Println(response)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i, alert := range response.Alerts {
			fmt.Println(strconv.Itoa(i) + ". " + alert.Message)
		}
	}*/

	/*response4, err := alertTest.CreateSavedSearch(ctx, alert.CreateSavedSearchRequest{
		Name: "test-busra",

		Teams: []alert.Team{
			{Name: "Marketing"},
		},
		Description: "description",
		Query:       "status: Open",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response4.SavedSearch

		fmt.Println("ID: " + savedSearch.ID)
		fmt.Println("Name: " + savedSearch.Name)
	}*/

	/*response6, err := alertTest.ListSavedSearches(ctx, alert.ListSavedSearchRequest{})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, search := range response6.SavedSearches {
			fmt.Println("ID: " + search.ID)
			fmt.Println("Name: " + search.Name)
		}
	}*/

	/*resp, err := alertTest.GetAsyncRequestStatus(ctx, alert.GetAsyncRequestStatusRequest{})

	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Printf("status: %s" , resp.Status.Status)
	}*/

	/*response5, err := alertTest.UpdateSavedSearch(ctx, alert.UpdateSavedSearchRequest{
		Name:  "testttttt2",
		Owner: alert.User{Username: "mbtekinsen@gmail.com", },
		Query:       "status: Open",
		NewName:"testttttt3",
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response5.SavedSearch

		fmt.Println("ID: " + savedSearch.ID)
		fmt.Println("Name: " + savedSearch.Name)
		fmt.Println("RequestId: " + response5.RequestID)

	}*/

	response5, err := alertTest.Delete(ctx, alert.DeleteAlertRequest{Identifier: &alert.Identifier{Alias: "alias3"}})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Create request ID: " + response5.RequestID)
	}

	//req7, err := savedsearches.NewDeleteSavedSearchRequest(&savedsearches.DeleteSavedSearchInput{ Name: "test3"})

	/*response7, err := alertTest.DeleteSavedSearch(ctx, alert.DeleteSavedSearchRequest{Name:"list-blue-team-alerts"})



	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted")
		fmt.Println(response7.RequestID)
	}*/

}
