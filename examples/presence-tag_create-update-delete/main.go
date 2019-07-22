package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital"
	ex "github.com/grokify/go-ringcentral-engage/examples"
	"github.com/grokify/go-ringcentral-engage/utils"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"obect" description:"An object" required:"true"`
	Action string `short:"a" long:"action" description:"An action (create|update|delete)" required:"true"`
	Name   string `short:"n" long:"name" description:"A tag name" required:"false"`
	Id     string `short:"i" long:"id" description:"A tag id" required:"false"`
}

func handlePresenceStatus(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_PRESENCE_STATUS")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_PRESENCE_STATUS_NO_NAME")
		}
		ex.HandleApiResponse(client.PresenceStatusesApi.CreatePresenceStatus(
			context.Background(),
			opts.Name))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetPresenceStatus(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetAllPresenceStatuses(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_PRESENCE_STATUS_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.PresenceStatusesApi.UpdatePresenceStatus(
			context.Background(), opts.Id, opts.Name))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_PRESENCE_STATUS_NO_ID")
		}
		ex.HandleApiResponse(client.PresenceStatusesApi.DeletePresenceStatus(
			context.Background(), opts.Id))
	}
}

func handleTag(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_TAG")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_TAG_NO_NAME")
		}
		ex.HandleApiResponse(client.TagsApi.CreateTag(
			context.Background(),
			opts.Name))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TagsApi.GetTag(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TagsApi.GetAllTags(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_TAG_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.TagsApi.UpdateTag(
			context.Background(), opts.Id, opts.Name))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_TAG_NO_ID")
		}
		ex.HandleApiResponse(client.TagsApi.DeleteTag(
			context.Background(), opts.Id))
	}
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch opts.Object {
	case "presenceStatus":
		handlePresenceStatus(client, opts)
	case "tag":
		handleTag(client, opts)
	}

	fmt.Println("DONE")
}
