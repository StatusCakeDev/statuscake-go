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

	locations, err := client.ListPagespeedMonitoringLocations(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("PAGESPEED LOCATIONS: %+v\n", locations.Data)

	locations, err = client.ListUptimeMonitoringLocations(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPTIME LOCATIONS: %+v\n", locations.Data)
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
