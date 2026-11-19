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

	res, err := client.CreateHeartbeatTest(ctx).
		Name("statuscake.com").
		Period(1800).
		ContactGroups([]string{groupID}).
		Paused(true).
		Tags([]string{"testing"}).
		Execute()
	if err != nil {
		return err
	}

	testID := res.Data.NewID
	fmt.Printf("HEARTBEAT CHECK ID: %s\n", testID)

	defer func() {
		if err := client.DeleteHeartbeatTest(ctx, testID).Execute(); err != nil {
			fmt.Println("could not delete the heartbeat check:", err)
		}
	}()

	test, err := client.GetHeartbeatTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("HEARTBEAT CHECK: %+v\n", test.Data)

	err = client.UpdateHeartbeatTest(ctx, testID).
		Period(3600).
		Paused(false).
		Execute()
	if err != nil {
		return err
	}

	test, err = client.GetHeartbeatTest(ctx, testID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED HEARTBEAT CHECK: %+v\n", test.Data)

	tests, err := client.ListHeartbeatTests(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("HEARTBEAT CHECKS: %+v\n", tests.Data)

	return nil
}
