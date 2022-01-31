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

	res, err := client.CreateContactGroup(context.Background()).
		Name("Operations Team").
		EmailAddresses([]string{
			"johnsmith@example.com",
			"janesmith@example.com",
		}).
		Integrations([]string{
			"1",
			"2",
			"3",
		}).
		MobileNumbers([]string{
			"447712345678",
			"447987462344",
		}).
		PingURL("https://ping.example.com").
		Execute()
	if err != nil {
		printError(err)
		return
	}

	groupID := res.Data.NewID
	fmt.Printf("CONTACT GROUP ID: %s\n", groupID)

	group, err := client.GetContactGroup(context.Background(), groupID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("CONTACT GROUP: %+v\n", group.Data)

	err = client.UpdateContactGroup(context.Background(), groupID).
		Name("Development Team").
		EmailAddresses([]string{}). // Remove all email addresses.
		Integrations([]string{
			"4",
			"5",
			"6",
		}).
		MobileNumbers([]string{
			"447891998195",
		}).
		PingURL("https://ping.example.com/groups").
		Execute()
	if err != nil {
		printError(err)
	}

	group, err = client.GetContactGroup(context.Background(), groupID).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("UPDATED CONTACT GROUP: %+v\n", group.Data)

	groups, err := client.ListContactGroups(context.Background()).Execute()
	if err != nil {
		printError(err)
	}

	fmt.Printf("CONTACT GROUPS: %+v\n", groups.Data)

	err = client.DeleteContactGroup(context.Background(), groupID).Execute()
	if err != nil {
		printError(err)
	}
}

func printError(err error) {
	fmt.Println(err)
	fmt.Printf("%+v\n", statuscake.Errors(err))
}
