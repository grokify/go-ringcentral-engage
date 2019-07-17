# \TeamsApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteTeam**](TeamsApi.md#DeleteTeam) | **Delete** /teams/{teamId} | Deleting a team
[**GetAllTeams**](TeamsApi.md#GetAllTeams) | **Get** /teams | Getting all teams
[**GetTeam**](TeamsApi.md#GetTeam) | **Get** /teams/{teamId} | Getting a team from its id



## DeleteTeam

> Team DeleteTeam(ctx, teamId, optional)
Deleting a team

This method destroys an existing team. It renders the team itself. It renders a 404 if id is invalid.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 
 **optional** | ***DeleteTeamOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteTeamOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **takeOverCategoryId** | **optional.String**| ID of a category to recategorize (optional). | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTeams

> GetAllTeamsResponse GetAllTeams(ctx, optional)
Getting all teams

This method renders teams ordered by creation date (ascending).

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTeamsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllTeamsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **optional.Int32**| The record index to start. Default value is 0. | 
 **limit** | **optional.String**| The max number of records to return. Default value is 30, max value is 150. | 

### Return type

[**GetAllTeamsResponse**](GetAllTeamsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> Team GetTeam(ctx, teamId)
Getting a team from its id

This method renders a team from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string**|  | 

### Return type

[**Team**](Team.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

