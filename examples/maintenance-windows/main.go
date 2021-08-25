package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/StatusCakeDev/statuscake-go"
)

func main() {
	var apiToken string

	if apiToken = os.Getenv("STATUSCAKE_API_TOKEN"); apiToken == "" {
		panic("STATUSCAKE_API_TOKEN not set in environment")
	}

	client := statuscake.NewAPIClient(apiToken)

	t := time.Now()
	res, err := client.CreateMaintenanceWindow(context.Background()).
		Name("Weekly maintenance").
		Start(t).
		End(t.Add(time.Hour * 3)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalWeekly).
		Tags([]string{"testing"}).
		Timezone("UTC").
		Execute()
	if err != nil {
		printError(err)
		return
	}

	windowID := res.Data.NewID
	fmt.Printf("MAINTENANCE WINDOW ID: %s\n", windowID)

	window, err := client.GetMaintenanceWindow(context.Background(), windowID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("MAINTENANCE WINDOW: %+v\n", window.Data)

	err = client.UpdateMaintenanceWindow(context.Background(), windowID).
		Name("Monthly maintenance").
		Start(t).
		End(t.Add(time.Hour * 48)).
		RepeatInterval(statuscake.MaintenanceWindowRepeatIntervalMonthly).
		Execute()
	if err != nil {
		printError(err)
	}

	window, err = client.GetMaintenanceWindow(context.Background(), windowID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED MAINTENANCE WINDOW: %+v\n", window.Data)

	windows, err := client.ListMaintenanceWindows(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("MAINTENANCE WINDOWS: %+v\n", windows.Data)

	err = client.DeleteMaintenanceWindow(context.Background(), windowID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
