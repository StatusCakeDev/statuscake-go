package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/StatusCakeDev/statuscake-go/credentials"
)

func main() {
	var apiToken string

	if apiToken = os.Getenv("STATUSCAKE_API_TOKEN"); apiToken == "" {
		panic("STATUSCAKE_API_TOKEN not set in environment")
	}

	bearer := credentials.NewBearerWithStaticToken(apiToken)
	client := statuscake.NewClient(statuscake.WithRequestCredentials(bearer))

	res, _ := client.CreateContactGroup(context.Background()).
		Name("Development Team").
		Execute()

	groupID := res.Data.NewID
	defer client.DeleteContactGroup(context.Background(), groupID).Execute()

	t := time.Now()
	res, _ = client.CreateMaintenanceWindow(context.Background()).
		Name("Saturday maintenance").
		End(t.Add(time.Hour * 3)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
		Start(t).
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
	fmt.Printf("UPTIME CHECK ID: %s\n", testID)

	test, err := client.GetUptimeTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME CHECK: %+v\n", test.Data)

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

	fmt.Printf("UPDATED UPTIME CHECK: %+v\n", test.Data)

	tests, err := client.ListUptimeTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME CHECKS: %+v\n", tests.Data)

	results, err := client.ListUptimeTestHistory(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME CHECK HISTORY: %+v\n", results.Data)

	periods, err := client.ListUptimeTestPeriods(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME CHECK PERIODS: %+v\n", periods.Data)

	alerts, err := client.ListUptimeTestAlerts(context.Background(), testID).Execute()
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
