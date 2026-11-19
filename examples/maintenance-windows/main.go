package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

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

	t := time.Now()

	res, err := client.CreateMaintenanceWindow(ctx).
		Name("Weekly maintenance").
		Start(t).
		End(t.Add(3 * time.Hour)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
		Tags([]string{"testing"}).
		Timezone("UTC").
		Execute()
	if err != nil {
		return err
	}

	windowID := res.Data.NewID
	fmt.Printf("MAINTENANCE WINDOW ID: %s\n", windowID)

	defer func() {
		if err := client.DeleteMaintenanceWindow(ctx, windowID).Execute(); err != nil {
			fmt.Println("could not delete the maintenance window:", err)
		}
	}()

	window, err := client.GetMaintenanceWindow(ctx, windowID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("MAINTENANCE WINDOW: %+v\n", window.Data)

	err = client.UpdateMaintenanceWindow(ctx, windowID).
		Name("Monthly maintenance").
		Start(t).
		End(t.Add(48 * time.Hour)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalMonthly).
		Execute()
	if err != nil {
		return err
	}

	window, err = client.GetMaintenanceWindow(ctx, windowID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED MAINTENANCE WINDOW: %+v\n", window.Data)

	windows, err := client.ListMaintenanceWindows(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("MAINTENANCE WINDOWS: %+v\n", windows.Data)

	return nil
}
