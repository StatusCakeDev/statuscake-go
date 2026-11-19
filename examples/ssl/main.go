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

	res, err := client.CreateSslTest(ctx).
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.SSLTestCheckRateFiveMinutes).
		AlertAt([]int32{1, 7, 30}).
		AlertBroken(true).
		AlertExpiry(true).
		AlertMixed(true).
		AlertReminder(true).
		ContactGroups([]string{groupID}).
		FollowRedirects(true).
		Paused(true).
		Execute()
	if err != nil {
		return err
	}

	testID := res.Data.NewID
	fmt.Printf("SSL CHECK ID: %s\n", testID)

	defer func() {
		if err := client.DeleteSslTest(ctx, testID).Execute(); err != nil {
			fmt.Println("could not delete the SSL check:", err)
		}
	}()

	test, err := client.GetSslTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("SSL CHECK: %+v\n", test.Data)

	err = client.UpdateSslTest(ctx, testID).
		CheckRate(statuscake.SSLTestCheckRateOneHour).
		ContactGroups([]string{}). // Remove all contact groups.
		Paused(false).
		Execute()
	if err != nil {
		return err
	}

	test, err = client.GetSslTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED SSL CHECK: %+v\n", test.Data)

	tests, err := client.ListSslTests(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("SSL CHECKS: %+v\n", tests.Data)

	return nil
}
