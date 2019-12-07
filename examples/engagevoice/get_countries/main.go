package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grokify/go-ringcentral-engage/engagevoiceutil"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath   string `short:"e" long:"envPath" description:"Environment File Path"`
	Token     string `short:"t" long:"token" description:"Token"`
	AccountID string `short:"a" long:"accountID" description:"AccountID"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	err = config.LoadDotEnvFirst(opts.EnvPath, os.Getenv("ENV_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Token) == 0 {
		opts.Token = os.Getenv("ENGAGE_VOICE_API_TOKEN")
	}
	if len(opts.AccountID) == 0 {
		opts.AccountID = os.Getenv("ENGAGE_VOICE_ACCOUNT_ID")
	}

	clientApis := engagevoiceutil.NewClientAPIs(opts.Token)
	apiClient := clientApis.APIClient

	info, resp, err := apiClient.CountriesApi.GetAvailableCountries(
		context.Background(), opts.AccountID,
	)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode >= 300 {
		log.Fatal("E_STATUS_CODE_GTE_300")
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
