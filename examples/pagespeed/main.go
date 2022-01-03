package main

import (
	"context"
	"fmt"
	"os"

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

	res, err := client.CreatePagespeedTest(context.Background()).
		Name("statuscake.com").
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.PagespeedTestCheckRateOneDay).
		AlertSmaller(10).
		AlertBigger(100).
		AlertSlower(1000).
		ContactGroups([]string{
			groupID,
		}).
		Paused(true).
		Region(statuscake.PagespeedTestRegionUnitedKingdom).
		Execute()
	if err != nil {
		printError(err)
		return
	}

	testID := res.Data.NewID
	fmt.Printf("PAGESPEED CHECK ID: %s\n", testID)

	test, err := client.GetPagespeedTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED CHECK: %+v\n", test.Data)

	err = client.UpdatePagespeedTest(context.Background(), testID).
		CheckRate(statuscake.PagespeedTestCheckRateOneHour).
		ContactGroups([]string{}). // Remove all contact groups.
		Paused(false).
		Execute()
	if err != nil {
		printError(err)
	}

	test, err = client.GetPagespeedTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED PAGESPEED CHECK: %+v\n", test.Data)

	tests, err := client.ListPagespeedTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED CHECKS: %+v\n", tests.Data)

	results, err := client.ListPagespeedTestHistory(context.Background(), testID).
		Days(4).
		Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED CHECK HISTORY: %+v\n", results.Data)

	err = client.DeletePagespeedTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
