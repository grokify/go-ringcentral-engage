package utils

import (
	"fmt"

	"github.com/grokify/go-ringcentral-engage/engagedigital/engagedigital"
	"github.com/grokify/goauth"
)

const (
	ApiUrlFormat = `https://%s.api.engagement.dimelo.com/1.0`
)

func NewApiClient(site, token string) *engagedigital.APIClient {
	cfg := engagedigital.NewConfiguration()
	cfg.HTTPClient = goauth.NewClientBearerTokenSimple(token)
	cfg.BasePath = fmt.Sprintf(ApiUrlFormat, site)
	return engagedigital.NewAPIClient(cfg)
}
