# Go API client for engagevoice

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./engagevoice"
```

## Documentation for API Endpoints

All URIs are relative to *https://portal.vacd.biz/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AgentsApi* | [**GetAgentGroups**](docs/AgentsApi.md#getagentgroups) | **Get** /admin/accounts/{accountId}/agentGroups | Returns a listing of agent groups for an account
*AgentsApi* | [**GetAgents**](docs/AgentsApi.md#getagents) | **Get** /admin/accounts/{accountId}/agentGroups/{agentGroupId}/agents | Returns a listing of agents for an agent group
*CampaignsApi* | [**GetDialGroupCampaigns**](docs/CampaignsApi.md#getdialgroupcampaigns) | **Get** /admin/accounts/{accountId}/dialGroups/{dialGroupId}/campaigns | Returns a listing of campaigns for a dial group
*CountriesApi* | [**GetAvailableCountries**](docs/CountriesApi.md#getavailablecountries) | **Get** /admin/accounts/{accountId}/countries/available | Get a list of available countries
*DialGroupsApi* | [**GetDialGroups**](docs/DialGroupsApi.md#getdialgroups) | **Get** /admin/accounts/{accountId}/dialGroups | Returns a listing of dial groups for an account
*LeadsApi* | [**UploadLeads**](docs/LeadsApi.md#uploadleads) | **Post** /admin/accounts/{accountId}/campaigns/{campaignId}/leadLoader/direct | Upload Leads


## Documentation For Models

 - [Agent](docs/Agent.md)
 - [AgentGroup](docs/AgentGroup.md)
 - [Campaign](docs/Campaign.md)
 - [Country](docs/Country.md)
 - [DialGroup](docs/DialGroup.md)
 - [Error](docs/Error.md)
 - [ExtendedLeadData](docs/ExtendedLeadData.md)
 - [Generic](docs/Generic.md)
 - [Lead](docs/Lead.md)
 - [UploadLeadsRequest](docs/UploadLeadsRequest.md)
 - [UploadLeadsResponse](docs/UploadLeadsResponse.md)


## Documentation For Authorization



## bearerAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Author


