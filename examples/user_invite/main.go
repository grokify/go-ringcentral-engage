package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/smtputil"
	"github.com/grokify/oauth2more"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/errors"

	"github.com/grokify/go-ringcentral-engage/engagedigital"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

type options struct {
	Site  string `short:"s" long:"site" description:"A site" required:"true"`
	Token string `short:"t" long:"token" description:"A token" required:"true"`
}

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = oauth2more.NewClientBearerTokenSimple(token)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}

func GetRole(client *engagedigital.APIClient, roleName string) (engagedigital.Role, error) {
	info, resp, err := client.RolesApi.GetAllRoles(context.Background(), nil)
	if err != nil {
		err = errors.Wrap(err, "E_GET_ALL_ROLES")
		return engagedigital.Role{}, err
	} else if resp.StatusCode != 200 {
		return engagedigital.Role{}, fmt.Errorf("E_STATUS_CODE [%v]", resp.StatusCode)
	}
	for _, role := range info.Records {
		if role.Label == roleName {
			return role, nil
		}
	}
	return engagedigital.Role{}, fmt.Errorf("E_ROLE_NOT_FOUND [%v]", roleName)
}

func main() {
	opts := options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	client := NewApiClient(opts.Site, opts.Token)

	roleName := "Agent"
	userEmail, err := smtputil.GmailAddressPlusTime("example")
	if err != nil {
		log.Fatal(err)
	}
	userEmail = "joe@example.com"
	userFirstName := "Joe"
	userLastName := "Example"

	role, err := GetRole(client, roleName)
	if err != nil {
		log.Fatal(err)
	}

	info, resp, err := client.UsersApi.InviteUser(
		context.Background(),
		userEmail,
		userFirstName,
		userLastName,
		role.Id, nil)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode != 200 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)

	fmt.Println("DONE")
}
