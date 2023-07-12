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

	res, err := client.CreateHeartbeatTest(context.Background()).
		Name("statuscake.com").
		Period(1800).
		ContactGroups([]string{
			groupID,
		}).
		Paused(true).
		Tags([]string{
			"testing",
		}).
		Execute()
	if err != nil {
		printError(err)
		return
	}

	testID := res.Data.NewID
	fmt.Printf("HEARTBEAT CHECK ID: %s\n", testID)

	test, err := client.GetHeartbeatTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("HEARTBEAT CHECK: %+v\n", test.Data)

	err = client.UpdateHeartbeatTest(context.Background(), testID).
		Period(3600).
		Paused(false).
		Execute()
	if err != nil {
		printError(err)
	}

	test, err = client.GetHeartbeatTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED HEARTBEAT CHECK: %+v\n", test.Data)

	tests, err := client.ListHeartbeatTests(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("HEARTBEAT CHECKS: %+v\n", tests.Data)

	err = client.DeleteHeartbeatTest(context.Background(), testID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
