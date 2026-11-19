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

	res, err := client.CreateContactGroup(ctx).
		Name("Operations Team").
		EmailAddresses([]string{"johnsmith@example.com", "janesmith@example.com"}).
		Integrations([]string{"1", "2", "3"}).
		MobileNumbers([]string{"447712345678", "447987462344"}).
		PingURL("https://ping.example.com").
		Execute()
	if err != nil {
		return err
	}

	groupID := res.Data.NewID
	fmt.Printf("CONTACT GROUP ID: %s\n", groupID)

	// The group is deleted however the rest of the example ends.
	defer func() {
		if err := client.DeleteContactGroup(ctx, groupID).Execute(); err != nil {
			fmt.Println("could not delete the contact group:", err)
		}
	}()

	group, err := client.GetContactGroup(ctx, groupID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("CONTACT GROUP: %+v\n", group.Data)

	err = client.UpdateContactGroup(ctx, groupID).
		Name("Development Team").
		EmailAddresses([]string{}). // Remove all email addresses.
		Integrations([]string{"4", "5", "6"}).
		MobileNumbers([]string{"447891998195"}).
		PingURL("https://ping.example.com/groups").
		Execute()
	if err != nil {
		return err
	}

	group, err = client.GetContactGroup(ctx, groupID).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("UPDATED CONTACT GROUP: %+v\n", group.Data)

	groups, err := client.ListContactGroups(ctx).Execute()
	if err != nil {
		return err
	}

	fmt.Printf("CONTACT GROUPS: %+v\n", groups.Data)

	return nil
}
