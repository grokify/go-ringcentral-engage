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
	Action string `short:"a" long:"action" description:"A action (create|update|delete)" required:"true"`
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

	switch opts.Action {
	case "create":
		fmt.Println("I_CREATING_TAG")
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 {
			log.Fatal("E_CREATE_NO_TAG_NAME")
		}
		info, resp, err := client.TagsApi.CreateTag(
			context.Background(),
			opts.Name)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "update":
		opts.Id = strings.TrimSpace(opts.Id)
		opts.Name = strings.TrimSpace(opts.Name)
		if len(opts.Name) == 0 || len(opts.Id) == 0 {
			log.Fatal("E_UPDATE_NO_TAG_ID_OR_NAME")
		}
		info, resp, err := client.TagsApi.UpdateTag(
			context.Background(), opts.Id, opts.Name)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	case "delete":
		opts.Id = strings.TrimSpace(opts.Id)
		if len(opts.Id) == 0 {
			log.Fatal("E_DELETE_NO_TAG_ID")
		}

		info, resp, err := client.TagsApi.DeleteTag(
			context.Background(), opts.Id)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode != 200 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(info)
	}

	fmt.Println("DONE")
}
