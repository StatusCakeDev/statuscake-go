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

	res, err := client.CreateSslTest(context.Background()).
		WebsiteURL("https://www.statuscake.com").
		CheckRate(statuscake.SSLTestCheckRateFiveMinutes).
		AlertAt([]int32{1, 7, 30}).
		AlertBroken(true).
		AlertExpiry(true).
		AlertMixed(true).
		AlertReminder(true).
		ContactGroups([]string{
			groupID,
		}).
		FollowRedirects(true).
		Paused(true).
		Execute()
	if err != nil {
		printError(err)
		return
	}

	testID := res.Data.NewID
	fmt.Printf("SSL CHECK ID: %s\n", testID)

	test, err := client.GetSslTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("SSL CHECK: %+v\n", test.Data)

	err = client.UpdateSslTest(context.Background(), testID).
		CheckRate(statuscake.SSLTestCheckRateOneHour).
		ContactGroups([]string{}). // Remove all contact groups.
		Paused(false).
		Execute()
	if err != nil {
		printError(err)
	}

	test, err = client.GetSslTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED SSL CHECK: %+v\n", test.Data)

	tests, err := client.ListSslTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("SSL CHECKS: %+v\n", tests.Data)

	err = client.DeleteSslTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
