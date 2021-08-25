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

	res, err := client.CreateContactGroup(context.Background()).
		Name("Operations Team").
		PingURL("https://ping.example.com").
		EmailAddresses([]string{
			"johnsmith@example.com",
			"janesmith@example.com",
		}).
		MobileNumbers([]string{
			"447712345678",
			"447987462344",
		}).
		Integrations([]string{
			"1",
			"2",
			"3",
		}).
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
		PingURL("https://ping.example.com/groups").
		EmailAddresses([]string{}). // Remove all email addresses.
		MobileNumbers([]string{
			"447891998195",
		}).
		Integrations([]string{
			"4",
			"5",
			"6",
		}).
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
