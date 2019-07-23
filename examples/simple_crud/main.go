package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/antihax/optional"
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

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	switch opts.Object {
	case "customField":
		handleCustomField(client, opts)
	case "presenceStatus":
		handlePresenceStatus(client, opts)
	case "tag":
		handleTag(client, opts)
	default:
		log.Fatal(fmt.Sprintf("E_OBJECT_NOT_SUPPORTED [%v]", opts.Object))
	}

	fmt.Println("DONE")
}

func formatRespStatusCodeError(statusCode int) string {
	return fmt.Sprintf("E_API_ERROR [%v]", statusCode)
}

func handleCustomField(client *engagedigital.APIClient, opts options) {
	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_CUSTOM_FIELD")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_CUSTOM_FIELD_NO_NAME")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.CreateCustomField(
			context.Background(), "Intervention", opts.Name, nil))
	case "read":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CustomFieldsApi.GetCustomField(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CustomFieldsApi.GetAllCustomFields(context.Background(), nil))
		}
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		apiOpts := &engagedigital.UpdateCustomFieldOpts{
			Label: optional.NewString(opts.Name)}
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_CUSTOM_FIELD_NO_ID_OR_NAME")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.UpdateCustomField(
			context.Background(), opts.Id, apiOpts))
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_CUSTOM_FIELD_NO_ID")
		}
		ex.HandleApiResponse(client.CustomFieldsApi.DeleteCustomField(
			context.Background(), opts.Id))
	}
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

/*
func handleTopology(client *engagedigital.APIClient, opts options) {
	logSlug := "TOPOLOGY"
	switch opts.Action {
	case "create":
		all, resp, err := client.TopologiesApi.GetAllTopologies(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(formatRespStatusCodeError(resp.StatusCode))
		}
		defTop, err := utils.FindDefaultTopology(all.Records)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("I_CREATING_%s", logSlug)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal(fmt.Sprintf("E_CREATE_%s_NO_NAME", logSlug))
		}
		newTop := engagedigital.Topology{Config: defTop.Config}
		newTopString, err := json.Marshal(newTop)
		if err != nil {
			log.Fatal(err)
		}
		apiOpts := &engagedigital.CreateTopologyOpts{
			JsonConfig: optional.NewString(string(newTopString))}

		ex.HandleApiResponse(client.TopologiesApi.CreateTopology(
			context.Background(), opts.Name, apiOpts))
	case "read":

		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TopologiesApi.GetTopology(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TopologiesApi.GetAllTopologies(context.Background(), nil))
		}

			case "update":
				opts.Id = strings.TrimSpace(opts.Id)
				opts.Name = strings.TrimSpace(opts.Name)
				apiOpts := &engagedigital.UpdateCustomFieldOpts{
					Label: optional.NewString(opts.Name)}
				if len(opts.Name) == 0 || len(opts.Id) == 0 {
					log.Fatal(fmt.Sprintf("E_UPDATE_%s_NO_ID_OR_NAME", logSlug))
				}
				ex.HandleApiResponse(client.CustomFieldsApi.UpdateCustomField(
					context.Background(), opts.Id, apiOpts))
			case "delete":
				opts.Id = strings.TrimSpace(opts.Id)
				if len(opts.Id) == 0 {
					log.Fatal(fmt.Sprintf("E_DELETE_%s_NO_ID", logSlug))
				}
				ex.HandleApiResponse(client.CustomFieldsApi.DeleteCustomField(
					context.Background(), opts.Id))

	}
}
*/
