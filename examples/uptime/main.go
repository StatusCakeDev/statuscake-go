package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
)

func main() {
	var apiToken string

	if apiToken = os.Getenv("STATUSCAKE_API_TOKEN"); apiToken == "" {
		panic("STATUSCAKE_API_TOKEN not set in environment")
	}

	client := statuscake.NewAPIClient(apiToken)

	res, _ := client.CreateContactGroup(context.Background()).
		Name("Development Team").
		Execute()

	groupID := res.Data.NewID
	defer client.DeleteContactGroup(context.Background(), groupID).Execute()

	t := time.Now()
	res, _ = client.CreateMaintenanceWindow(context.Background()).
		Name("Saturday maintenance").
		Start(t).
		End(t.Add(time.Hour * 3)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
		Tags([]string{"testing"}).
		Timezone("UTC").
		Execute()

	windowID := res.Data.NewID
	defer client.DeleteMaintenanceWindow(context.Background(), windowID).Execute()

	res, err := client.CreateUptimeTest(context.Background()).
		Name("statuscake.com").
		TestType(statuscake.UptimeTestTypeHTTP).
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.UptimeTestCheckRateFifteenMinutes).
		ContactGroups([]string{
			groupID,
		}).
		EnableSSLAlert(true).
		FollowRedirects(true).
		Paused(true).
		Regions([]string{
			"london",
		}).
		Tags([]string{
			"testing",
		}).
		Execute()
	if err != nil {
		printError(err)
		return
	}

	testID := res.Data.NewID
	fmt.Printf("UPTIME TEST ID: %s\n", testID)

	test, err := client.GetUptimeTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME TEST: %+v\n", test.Data)

	err = client.UpdateUptimeTest(context.Background(), testID).
		CheckRate(statuscake.UptimeTestCheckRateThirtyMinutes).
		Paused(false).
		Regions([]string{
			"london",
			"paris",
		}).
		Execute()
	if err != nil {
		printError(err)
	}

	test, err = client.GetUptimeTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED UPTIME TEST: %+v\n", test.Data)

	tests, err := client.ListUptimeTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME TESTS: %+v\n", tests.Data)

	history, err := client.ListUptimeTestHistory(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME TEST HISTORY: %+v\n", history.Data)

	alerts, err := client.ListSentAlerts(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("SENT ALERTS: %+v\n", alerts.Data)

	err = client.DeleteUptimeTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
