package main

import (
	"context"
	"fmt"
	"log"
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
			info, resp, err := client.AttachmentsApi.GetAttachment(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.AttachmentsApi.GetAllAttachments(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "community":
		if len(opts.Id) > 0 {
			info, resp, err := client.CommunitiesApi.GetCommunity(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.CommunitiesApi.GetAllCommunities(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "folder":
		if len(opts.Id) > 0 {
			info, resp, err := client.FoldersApi.GetFolder(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.FoldersApi.GetAllFolders(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "identity":
		if len(opts.Id) > 0 {
			info, resp, err := client.IdentitiesApi.GetIdentity(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.IdentitiesApi.GetAllIdentities(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "locale":
		info, resp, err := client.LocalesApi.GetAllLocales(context.Background())
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "source":
		if len(opts.Id) > 0 {
			info, resp, err := client.SourcesApi.GetSource(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.SourcesApi.GetAllSources(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "role":
		if len(opts.Id) > 0 {
			info, resp, err := client.RolesApi.GetRole(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.RolesApi.GetAllRoles(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "team":
		if len(opts.Id) > 0 {
			info, resp, err := client.TeamsApi.GetTeam(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.TeamsApi.GetAllTeams(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	case "user":
		if len(opts.Id) > 0 {
			info, resp, err := client.UsersApi.GetUser(context.Background(), opts.Id)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		} else {
			info, resp, err := client.UsersApi.GetAllUsers(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			} else if resp.StatusCode != 200 {
				log.Fatal(resp.StatusCode)
			}
			fmtutil.PrintJSON(info)
		}
	}

	fmt.Println("DONE")
}
