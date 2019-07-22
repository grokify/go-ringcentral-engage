package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/jessevdk/go-flags"

	ex "github.com/grokify/go-ringcentral-engage/examples"
	"github.com/grokify/go-ringcentral-engage/utils"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"object" description:"An object" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := utils.NewApiClient(opts.Site, opts.Token)

	opts.Id = strings.TrimSpace(opts.Id)
	opts.Object = strings.ToLower(strings.TrimSpace(opts.Object))

	switch opts.Object {
	case "attachment":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.AttachmentsApi.GetAttachment(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.AttachmentsApi.GetAllAttachments(context.Background(), nil))
		}
	case "category":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CategoriesApi.GetCategory(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CategoriesApi.GetAllCategories(context.Background(), nil))
		}
	case "channel":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ChannelsApi.GetChannel(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ChannelsApi.GetAllChannels(context.Background(), nil))
		}
	case "community":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CommunitiesApi.GetCommunity(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CommunitiesApi.GetAllCommunities(context.Background(), nil))
		}
	case "customFields":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CustomFieldsApi.GetCustomField(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CustomFieldsApi.GetAllCustomFields(context.Background(), nil))
		}
	case "folder":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.FoldersApi.GetFolder(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.FoldersApi.GetAllFolders(context.Background(), nil))
		}
	case "identity":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.IdentitiesApi.GetIdentity(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.IdentitiesApi.GetAllIdentities(context.Background(), nil))
		}
	case "locale":
		ex.HandleApiResponse(client.LocalesApi.GetAllLocales(context.Background()))
	case "source":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.SourcesApi.GetSource(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.SourcesApi.GetAllSources(context.Background(), nil))
		}
	case "role":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.RolesApi.GetRole(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.RolesApi.GetAllRoles(context.Background(), nil))
		}
	case "presenceStatus":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetPresenceStatus(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetAllPresenceStatuses(context.Background(), nil))
		}
	case "tag":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TagsApi.GetTag(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TagsApi.GetAllTags(context.Background(), nil))
		}
	case "team":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TeamsApi.GetTeam(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TeamsApi.GetAllTeams(context.Background(), nil))
		}
	case "user":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.UsersApi.GetUser(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.UsersApi.GetAllUsers(context.Background(), nil))
		}
	}

	fmt.Println("DONE")
}
