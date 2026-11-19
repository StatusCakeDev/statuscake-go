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

	res, err := client.CreatePagespeedTest(ctx).
		Name("statuscake.com").
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.PagespeedTestCheckRateOneDay).
		AlertSmaller(10).
		AlertBigger(100).
		AlertSlower(1000).
		ContactGroups([]string{groupID}).
		Paused(true).
		Region(statuscake.PagespeedTestRegionUnitedKingdom).
		Execute()
	if err != nil {
		return err
	}

	testID := res.Data.NewID
	fmt.Printf("PAGESPEED CHECK ID: %s\n", testID)

	defer func() {
		if err := client.DeletePagespeedTest(ctx, testID).Execute(); err != nil {
			fmt.Println("could not delete the pagespeed check:", err)
		}
	}()

	test, err := client.GetPagespeedTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("PAGESPEED CHECK: %+v\n", test.Data)

	err = client.UpdatePagespeedTest(ctx, testID).
		CheckRate(statuscake.PagespeedTestCheckRateOneHour).
		ContactGroups([]string{}). // Remove all contact groups.
		Paused(false).
		Execute()
	if err != nil {
		return err
	}

	test, err = client.GetPagespeedTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED PAGESPEED CHECK: %+v\n", test.Data)

	tests, err := client.ListPagespeedTests(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("PAGESPEED CHECKS: %+v\n", tests.Data)

	results, err := client.ListPagespeedTestHistory(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("PAGESPEED CHECK HISTORY: %+v\n", results.Data)

	return nil
}
