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
	case "content":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ContentsApi.GetContent(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ContentsApi.GetAllContents(context.Background(), nil))
		}
	case "customFields":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.CustomFieldsApi.GetCustomField(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.CustomFieldsApi.GetAllCustomFields(context.Background(), nil))
		}
	case "event":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.EventsApi.GetEvent(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.EventsApi.GetAllEvents(context.Background(), nil))
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
	case "identityGroup":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.IdentityGroupsApi.GetIdentityGroup(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.IdentityGroupsApi.GetAllIdentityGroups(context.Background(), nil))
		}
	case "locale":
		ex.HandleApiResponse(client.LocalesApi.GetAllLocales(context.Background()))
	case "presenceStatus":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetPresenceStatus(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.PresenceStatusesApi.GetAllPresenceStatuses(context.Background(), nil))
		}
	case "role":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.RolesApi.GetRole(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.RolesApi.GetAllRoles(context.Background(), nil))
		}
	case "settings":
		ex.HandleApiResponse(client.SettingsApi.GetAllSettings(context.Background()))
	case "source":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.SourcesApi.GetSource(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.SourcesApi.GetAllSources(context.Background(), nil))
		}
	case "tag":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TagsApi.GetTag(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TagsApi.GetAllTags(context.Background(), nil))
		}
	case "task":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TasksApi.GetTask(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TasksApi.GetAllTasks(context.Background(), nil))
		}
	case "team":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TeamsApi.GetTeam(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.TeamsApi.GetAllTeams(context.Background(), nil))
		}
	case "thread":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.ThreadsApi.GetThread(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.ThreadsApi.GetAllThreads(context.Background(), nil))
		}
	case "timesheet":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.TimeSheetsApi.GetTimeSheet(context.Background(), "me"))
		} else {
			ex.HandleApiResponse(client.TimeSheetsApi.GetAllTimeSheets(context.Background(), nil))
		}
	case "user":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.UsersApi.GetUser(context.Background(), opts.Id))
		} else {
			ex.HandleApiResponse(client.UsersApi.GetAllUsers(context.Background(), nil))
		}
	/*case "usersourcespermissions":
	info, resp, err := client.UsersApi.GetUser(context.Background(), "me")
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmt.Println(info.Id)
	ex.HandleApiResponse(client.UserSourcesPermissionsApi.GetUserSourcesPermission(
		context.Background(), info.Id))*/
	case "webhook":
		if len(opts.Id) > 0 {
			ex.HandleApiResponse(client.WebhooksApi.GetWebhook(context.Background(), opts.Id, opts.Token))
		} else {
			ex.HandleApiResponse(client.WebhooksApi.GetAllWebhooks(context.Background(), opts.Token, nil))
		}
	default:
		log.Fatal(fmt.Sprintf("E_ACTION_NOT_FOUND [%s]", opts.Object))
	}

	fmt.Println("DONE")
}
