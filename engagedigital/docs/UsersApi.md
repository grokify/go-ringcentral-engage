# \UsersApi

All URIs are relative to *https://DOMAIN.api.engagement.dimelo.com/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /users/{userId} | Deleting a user
[**GetAllUsers**](UsersApi.md#GetAllUsers) | **Get** /users | Getting all users
[**GetUser**](UsersApi.md#GetUser) | **Get** /users/{userId} | Getting a user from its id



## DeleteUser

> User DeleteUser(ctx, userId)
Deleting a user

This method deletes the given user. In case of success it renders the deleted user, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can update users. The user affiliated to the token must have at least all the permissions of the other user. If the user affiliated to the token has the manage_users_of_my_teams permission, the deleted user will need to belong to at least one of the teams he’s the leader of.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllUsers

> GetAllTeamsResponse GetAllUsers(ctx, optional)
Getting all users

This method renders users ordered by creation date (descending).  Authorization​: only users that can view users. If the user affiliated to the token has the manage_users_of_my_teams permission, only the users belonging to at least one of the teams he’s the leader of will be returned.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **optional.String**| To filter users on given email. | 
 **categoryId** | **optional.String**| To filter users on given category id. | 
 **identityId** | **optional.String**| To filter users on given identity id. | 
 **externalId** | **optional.String**| To filter users on given external id. | 
 **roleId** | **optional.String**| To filter users on given role id. | 
 **teamId** | **optional.String**| To filter users on given team id. | 
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


## GetUser

> User GetUser(ctx, userId)
Getting a user from its id

This method renders a user from given id.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**|  | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
