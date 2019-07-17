package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/oauth2more"
	"github.com/jessevdk/go-flags"

	"github.com/grokify/go-ringcentral-engage/engagedigital"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

type options struct {
	Site   string `short:"s" long:"site" description:"A site" required:"true"`
	Token  string `short:"t" long:"token" description:"A token" required:"true"`
	Object string `short:"o" long:"object" description:"An object" required:"true"`
	Id     string `short:"i" long:"id" description:"An object id" required:"false"`
}

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = oauth2more.NewClientBearerTokenSimple(token)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := NewApiClient(opts.Site, opts.Token)

	opts.Id = strings.TrimSpace(opts.Id)
	opts.Object = strings.ToLower(strings.TrimSpace(opts.Object))

	switch opts.Object {
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
	}

	fmt.Println("DONE")
}
