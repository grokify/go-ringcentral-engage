package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/utils"
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"object" description:"An object" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
}

func handleRepsonse(info interface{}, resp *http.Response, err error) {
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
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
			handleRepsonse(client.AttachmentsApi.GetAttachment(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.AttachmentsApi.GetAllAttachments(context.Background(), nil))
		}
	case "category":
		if len(opts.Id) > 0 {
			handleRepsonse(client.CategoriesApi.GetCategory(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.CategoriesApi.GetAllCategories(context.Background(), nil))
		}
	case "community":
		if len(opts.Id) > 0 {
			handleRepsonse(client.CommunitiesApi.GetCommunity(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.CommunitiesApi.GetAllCommunities(context.Background(), nil))
		}
	case "folder":
		if len(opts.Id) > 0 {
			handleRepsonse(client.FoldersApi.GetFolder(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.FoldersApi.GetAllFolders(context.Background(), nil))
		}
	case "identity":
		if len(opts.Id) > 0 {
			handleRepsonse(client.IdentitiesApi.GetIdentity(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.IdentitiesApi.GetAllIdentities(context.Background(), nil))
		}
	case "locale":
		handleRepsonse(client.LocalesApi.GetAllLocales(context.Background()))
	case "source":
		if len(opts.Id) > 0 {
			handleRepsonse(client.SourcesApi.GetSource(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.SourcesApi.GetAllSources(context.Background(), nil))
		}
	case "role":
		if len(opts.Id) > 0 {
			handleRepsonse(client.RolesApi.GetRole(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.RolesApi.GetAllRoles(context.Background(), nil))
		}
	case "tag":
		if len(opts.Id) > 0 {
			handleRepsonse(client.TagsApi.GetTag(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.TagsApi.GetAllTags(context.Background(), nil))
		}
	case "team":
		if len(opts.Id) > 0 {
			handleRepsonse(client.TeamsApi.GetTeam(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.TeamsApi.GetAllTeams(context.Background(), nil))
		}
	case "user":
		if len(opts.Id) > 0 {
			handleRepsonse(client.UsersApi.GetUser(context.Background(), opts.Id))
		} else {
			handleRepsonse(client.UsersApi.GetAllUsers(context.Background(), nil))
		}
	}

	fmt.Println("DONE")
}
