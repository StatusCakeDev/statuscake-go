package main

import (
	"context"
	"fmt"
	"os"

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

	res, err := client.CreatePagespeedTest(context.Background()).
		Name("statuscake.com").
		WebsiteURL("https://www.statuscake.com").
		LocationISO(statuscake.PagespeedTestLocationISOUnitedKingdom).
		CheckRate(statuscake.PagespeedTestCheckRateOneDay).
		ContactGroups([]string{
			groupID,
		}).
		AlertSmaller(10).
		AlertBigger(100).
		AlertSlower(1000).
		Paused(true).
		Execute()
	if err != nil {
		printError(err)
		return
	}

	testID := res.Data.NewID
	fmt.Printf("PAGESPEED TEST ID: %s\n", testID)

	test, err := client.GetPagespeedTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED TEST: %+v\n", test.Data)

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

	fmt.Printf("UPDATED PAGESPEED TEST: %+v\n", test.Data)

	tests, err := client.ListPagespeedTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED TESTS: %+v\n", tests.Data)

	history, err := client.ListPagespeedTestHistory(context.Background(), testID).
		Days(4).
		Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED TEST HISTORY: %+v\n", history.Data)

	err = client.DeletePagespeedTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
