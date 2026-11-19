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

	locations, err := client.ListPagespeedMonitoringLocations(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("PAGESPEED LOCATIONS: %+v\n", locations.Data)

	// The region codes are the values that an uptime check accepts as regions.
	locations, err = client.ListUptimeMonitoringLocations(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPTIME LOCATIONS: %+v\n", locations.Data)

	return nil
}
