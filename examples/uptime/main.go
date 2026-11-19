package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/StatusCakeDev/statuscake-go"
	"github.com/StatusCakeDev/statuscake-go/credentials"
)

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println(err)
		fmt.Printf("%+v\n", statuscake.Errors(err))
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	token := os.Getenv("STATUSCAKE_API_TOKEN")
	if token == "" {
		return errors.New("STATUSCAKE_API_TOKEN not set in environment")
	}

	client := statuscake.NewClient(statuscake.WithRequestCredentials(credentials.NewBearerWithStaticToken(token)))

	// A contact group for the check to alert, deleted however the example ends.
	group, err := client.CreateContactGroup(ctx).Name("Development Team").Execute()
	if err != nil {
		return err
	}

	groupID := group.Data.NewID
	defer func() {
		if err := client.DeleteContactGroup(ctx, groupID).Execute(); err != nil {
			fmt.Println("could not delete the contact group:", err)
		}
	}()

	res, err := client.CreateUptimeTest(ctx).
		Name("statuscake.com").
		TestType(statuscake.UptimeTestTypeHTTP).
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.UptimeTestCheckRateFifteenMinutes).
		ContactGroups([]string{groupID}).
		EnableSSLAlert(true).
		FollowRedirects(true).
		Paused(true).
		Regions([]string{"london"}).
		Tags([]string{"testing"}).
		Execute()
	if err != nil {
		return err
	}

	testID := res.Data.NewID
	fmt.Printf("UPTIME CHECK ID: %s\n", testID)

	defer func() {
		if err := client.DeleteUptimeTest(ctx, testID).Execute(); err != nil {
			fmt.Println("could not delete the uptime check:", err)
		}
	}()

	test, err := client.GetUptimeTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPTIME CHECK: %+v\n", test.Data)

	err = client.UpdateUptimeTest(ctx, testID).
		CheckRate(statuscake.UptimeTestCheckRateThirtyMinutes).
		Paused(false).
		Regions([]string{"london", "paris"}).
		Execute()
	if err != nil {
		return err
	}

	test, err = client.GetUptimeTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED UPTIME CHECK: %+v\n", test.Data)

	tests, err := client.ListUptimeTests(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPTIME CHECKS: %+v\n", tests.Data)

	results, err := client.ListUptimeTestHistory(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPTIME CHECK HISTORY: %+v\n", results.Data)

	periods, err := client.ListUptimeTestPeriods(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPTIME CHECK PERIODS: %+v\n", periods.Data)

	alerts, err := client.ListUptimeTestAlerts(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("SENT ALERTS: %+v\n", alerts.Data)

	return nil
}
