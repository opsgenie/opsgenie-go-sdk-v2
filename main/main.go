package main

import (
	"context"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert/savedsearches"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"time"
)

func main() {

	alertTest := client.NewAlertClient(client.Config{
		ApiKey:         "5d2891dc-8e22-403c-a124-0becc4e4c460", //8f8e3be1-9684-4bb9-9ebb-46015d0c9952
		OpsGenieAPIURL: "http://localhost:9002",
	})

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	//list alert
	/*req, err := alert.NewListAlertRequest(&alert.ListAlertInput{
			Limit:15,
			Offset:               0,
			SearchIdentifierType: alert.Name,
		})

	response, err := alertTest.List(ctx, req)



	if err != nil {
		fmt.Println(err.Error())
	} else {
		for i, alert := range response.Alerts {
			fmt.Println(strconv.Itoa(i) + ". " + alert.Message)
		}
	}*/

	//create alert
	/*req2, err := alert.NewCreateAlertRequest(&alert.CreateAlertInput{
		Message:     "message99",
		Alias:       "alias99",
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
		User:     "busra+test@opsgenie.com",
		Note:     "alert note2",
	})

	response2, err := alertTest.Create(ctx, req2)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Create request ID: " + response2.RequestID)
	}*/

	// delete alert
	/*identifierInput := alert.Identifier{Alias:"alias4"}

	//identifier , _ := alert.NewIdentifierRequest(&identifierInput)


	req3, err := alert.NewDeleteAlertRequest(
		&alert.DeleteAlertInput{
		Identifier: &identifierInput,

	})

	response3, err := alertTest.Delete(ctx,req3)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Create request ID: " + response3.RequestID)
	}*/

	/*request := alertsv2.DeleteAlertRequest{
		Identifier: &alertsv2.Identifier{
			TinyID: "2",
		},
		Source: "source",
		User:   "user@opsgenie.com",
	}

	response, err := alertCli.Delete(request)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("RequestID" + response.RequestID)
	}*/

	/*req4, err := savedsearches.NewCreateSavedSearchRequest(&savedsearches.CreateSavedSearchInput{
		Name: "test-busra",
		Owner: alert.User{
			Username: "busra+test@opsgenie.com",
		},
		Teams: []alert.Team{
			{Name: "Marketing"},
		},
		Description: "description",
		Query:       "status: Open",
	})

	response4, err := alertTest.CreateSavedSearch(ctx, req4)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response4.SavedSearch

		fmt.Println("ID: " + savedSearch.ID)
		fmt.Println("Name: " + savedSearch.Name)
	}*/

	/*savedSearchIdentifier := savedsearches.SavedSearchIdentifier{Name:"test3"}

	req5, err := savedsearches.NewUpdateSavedSearchRequest(&savedsearches.UpdateSavedSearchInput{
		Name:  "test",
		Owner: alert.User{Username: "busra+test@opsgenie.com", },
		Query:       "status: Open",

	}, savedSearchIdentifier)

	response5, err := alertTest.UpdateSavedSearch(ctx, req5)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		savedSearch := response5.SavedSearch

		fmt.Println("ID: " + savedSearch.ID)
		fmt.Println("Name: " + savedSearch.Name)
		fmt.Println("RequestId: " + response5.RequestID)
		fmt.Println("RateLimitState: " + response5.RateLimitState)
		fmt.Printf("ResponseTime: %f " , response5.ResponseTime)

	}*/

	req6, err := savedsearches.NewListSavedSearchRequest()

	response6, err := alertTest.ListSavedSearches(nil, req6)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, search := range response6.SavedSearches {
			fmt.Println("ID: " + search.ID)
			fmt.Println("Name: " + search.Name)
		}
	}

	/*req7, err := savedsearches.NewDeleteSavedSearchRequest(&savedsearches.DeleteSavedSearchInput{ Name: "test3"})


	response7, err := alertTest.DeleteSavedSearch(ctx, req7)



	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Deleted")
		fmt.Println(response7.RequestID)
	}*/

	//cli.SendGetRequest()

	//alertCli.SendPostRequest()

	/*req, _ :=cli.NewRequest("GET","http://golang.org/",nil)

	println(req.Response)

	resp, err := cli.Do(req,nil)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	//printing the response
	log.Println(string(data))

	println(resp.Status)
	println(err)*/

}
