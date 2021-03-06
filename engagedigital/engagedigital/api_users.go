/*
 * Engage Digital API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package engagedigital

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type UsersApiService service

/*
UsersApiService Creating a user
This method creates a new user. In case of success it renders the created user, otherwise, it renders an error (422 HTTP code).
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param email User email (mandatory).
 * @param firstname User firstname (mandatory).
 * @param lastname User lastname (mandatory).
 * @param password User plain password (mandatory).
 * @param roleId User role id (mandatory).
 * @param optional nil or *CreateUserOpts - Optional Parameters:
 * @param "CategoryIds" (optional.Interface of []string) -  User list of category ids (multiple).
 * @param "Enabled" (optional.Bool) -  Whether the user is enabled or not (boolean).
 * @param "ExternalId" (optional.String) -  User external id, used for SSO.
 * @param "Gender" (optional.String) -  User gender (\"man\" or \"woman\").
 * @param "IdentityIds" (optional.Interface of []string) -  User list of identity ids (multiple).
 * @param "Locale" (optional.String) -  Language for the user interface.
 * @param "Nickname" (optional.String) -  User nickname.
 * @param "TeamIds" (optional.Interface of []string) -  User list of team ids (multiple).
 * @param "Timezone" (optional.String) -  Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone.
 * @param "SpokenLanguages" (optional.Interface of []string) -  List of locales corresponding to the languages spoken by the user (multiple).
@return User
*/

type CreateUserOpts struct {
	CategoryIds     optional.Interface
	Enabled         optional.Bool
	ExternalId      optional.String
	Gender          optional.String
	IdentityIds     optional.Interface
	Locale          optional.String
	Nickname        optional.String
	TeamIds         optional.Interface
	Timezone        optional.String
	SpokenLanguages optional.Interface
}

func (a *UsersApiService) CreateUser(ctx context.Context, email string, firstname string, lastname string, password string, roleId string, localVarOptionals *CreateUserOpts) (User, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  User
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CategoryIds.IsSet() {
		t := localVarOptionals.CategoryIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("category_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("category_ids[]", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("email", parameterToString(email, ""))
	if localVarOptionals != nil && localVarOptionals.Enabled.IsSet() {
		localVarQueryParams.Add("enabled", parameterToString(localVarOptionals.Enabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExternalId.IsSet() {
		localVarQueryParams.Add("external_id", parameterToString(localVarOptionals.ExternalId.Value(), ""))
	}
	localVarQueryParams.Add("firstname", parameterToString(firstname, ""))
	if localVarOptionals != nil && localVarOptionals.Gender.IsSet() {
		localVarQueryParams.Add("gender", parameterToString(localVarOptionals.Gender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdentityIds.IsSet() {
		t := localVarOptionals.IdentityIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("identity_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("identity_ids[]", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("lastname", parameterToString(lastname, ""))
	if localVarOptionals != nil && localVarOptionals.Locale.IsSet() {
		localVarQueryParams.Add("locale", parameterToString(localVarOptionals.Locale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nickname.IsSet() {
		localVarQueryParams.Add("nickname", parameterToString(localVarOptionals.Nickname.Value(), ""))
	}
	localVarQueryParams.Add("password", parameterToString(password, ""))
	localVarQueryParams.Add("role_id", parameterToString(roleId, ""))
	if localVarOptionals != nil && localVarOptionals.TeamIds.IsSet() {
		t := localVarOptionals.TeamIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("team_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("team_ids[]", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Timezone.IsSet() {
		localVarQueryParams.Add("timezone", parameterToString(localVarOptionals.Timezone.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SpokenLanguages.IsSet() {
		t := localVarOptionals.SpokenLanguages.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("spoken_languages[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("spoken_languages[]", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v User
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
UsersApiService Deleting a user
This method deletes the given user. In case of success it renders the deleted user, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can update users. The user affiliated to the token must have at least all the permissions of the other user. If the user affiliated to the token has the manage_users_of_my_teams permission, the deleted user will need to belong to at least one of the teams he’s the leader of.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId
@return User
*/
func (a *UsersApiService) DeleteUser(ctx context.Context, userId string) (User, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  User
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v User
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
UsersApiService Getting all users
This method renders users ordered by creation date (descending).  Authorization​: only users that can view users. If the user affiliated to the token has the manage_users_of_my_teams permission, only the users belonging to at least one of the teams he’s the leader of will be returned.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *GetAllUsersOpts - Optional Parameters:
 * @param "Email" (optional.String) -  To filter users on given email.
 * @param "CategoryId" (optional.String) -  To filter users on given category id.
 * @param "IdentityId" (optional.String) -  To filter users on given identity id.
 * @param "ExternalId" (optional.String) -  To filter users on given external id.
 * @param "RoleId" (optional.String) -  To filter users on given role id.
 * @param "TeamId" (optional.String) -  To filter users on given team id.
 * @param "Offset" (optional.Int32) -  The record index to start. Default value is 0.
 * @param "Limit" (optional.Int32) -  The max number of records to return. Default value is 30, max value is 150.
@return GetAllUsersResponse
*/

type GetAllUsersOpts struct {
	Email      optional.String
	CategoryId optional.String
	IdentityId optional.String
	ExternalId optional.String
	RoleId     optional.String
	TeamId     optional.String
	Offset     optional.Int32
	Limit      optional.Int32
}

func (a *UsersApiService) GetAllUsers(ctx context.Context, localVarOptionals *GetAllUsersOpts) (GetAllUsersResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  GetAllUsersResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CategoryId.IsSet() {
		localVarQueryParams.Add("category_id", parameterToString(localVarOptionals.CategoryId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdentityId.IsSet() {
		localVarQueryParams.Add("identity_id", parameterToString(localVarOptionals.IdentityId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExternalId.IsSet() {
		localVarQueryParams.Add("external_id", parameterToString(localVarOptionals.ExternalId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoleId.IsSet() {
		localVarQueryParams.Add("role_id", parameterToString(localVarOptionals.RoleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TeamId.IsSet() {
		localVarQueryParams.Add("team_id", parameterToString(localVarOptionals.TeamId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v GetAllUsersResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
UsersApiService Getting a user from its id
This method renders a user from given id.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId
@return User
*/
func (a *UsersApiService) GetUser(ctx context.Context, userId string) (User, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  User
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v User
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
UsersApiService Inviting a user
This method invites a new user. In case of success it renders the created user, otherwise, it renders an error (422 HTTP code).  Authorization​: only users that can invite other users. If the user affiliated to the token has the manage_users_of_my_teams permission, the invited user will need to belong to at least one of the teams he’s the leader of. It will not be possible to assign the user to other teams.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param email User email (mandatory).
 * @param firstname User firstname (mandatory).
 * @param lastname User lastname (mandatory).
 * @param roleId User role id (mandatory).
 * @param optional nil or *InviteUserOpts - Optional Parameters:
 * @param "CategoryIds" (optional.Interface of []string) -  User list of category ids (multiple).
 * @param "Enabled" (optional.Bool) -  Whether the user is enabled or not (boolean).
 * @param "ExternalId" (optional.String) -  User external id.
 * @param "Gender" (optional.String) -  User gender (\"man\" or \"woman\").
 * @param "IdentityIds" (optional.Interface of []string) -  User list of identity ids (multiple).
 * @param "Locale" (optional.String) -  Language for the user interface.
 * @param "Nickname" (optional.String) -  User nickname.
 * @param "TeamIds" (optional.Interface of []string) -  User list of team ids (multiple).
 * @param "Timezone" (optional.String) -  Use the timezone endpoint to get the timezone name (String), default is empty for
 * @param "SpokenLanguages" (optional.Interface of []string) -  List of locales corresponding to the languages spoken by the user (multiple).
@return User
*/

type InviteUserOpts struct {
	CategoryIds     optional.Interface
	Enabled         optional.Bool
	ExternalId      optional.String
	Gender          optional.String
	IdentityIds     optional.Interface
	Locale          optional.String
	Nickname        optional.String
	TeamIds         optional.Interface
	Timezone        optional.String
	SpokenLanguages optional.Interface
}

func (a *UsersApiService) InviteUser(ctx context.Context, email string, firstname string, lastname string, roleId string, localVarOptionals *InviteUserOpts) (User, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  User
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/invite"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CategoryIds.IsSet() {
		t := localVarOptionals.CategoryIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("category_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("category_ids", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("email", parameterToString(email, ""))
	if localVarOptionals != nil && localVarOptionals.Enabled.IsSet() {
		localVarQueryParams.Add("enabled", parameterToString(localVarOptionals.Enabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExternalId.IsSet() {
		localVarQueryParams.Add("external_id", parameterToString(localVarOptionals.ExternalId.Value(), ""))
	}
	localVarQueryParams.Add("firstname", parameterToString(firstname, ""))
	if localVarOptionals != nil && localVarOptionals.Gender.IsSet() {
		localVarQueryParams.Add("gender", parameterToString(localVarOptionals.Gender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdentityIds.IsSet() {
		t := localVarOptionals.IdentityIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("identity_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("identity_ids", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("lastname", parameterToString(lastname, ""))
	if localVarOptionals != nil && localVarOptionals.Locale.IsSet() {
		localVarQueryParams.Add("locale", parameterToString(localVarOptionals.Locale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nickname.IsSet() {
		localVarQueryParams.Add("nickname", parameterToString(localVarOptionals.Nickname.Value(), ""))
	}
	localVarQueryParams.Add("role_id", parameterToString(roleId, ""))
	if localVarOptionals != nil && localVarOptionals.TeamIds.IsSet() {
		t := localVarOptionals.TeamIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("team_ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("team_ids", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Timezone.IsSet() {
		localVarQueryParams.Add("timezone", parameterToString(localVarOptionals.Timezone.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SpokenLanguages.IsSet() {
		t := localVarOptionals.SpokenLanguages.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("spoken_languages", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("spoken_languages", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v User
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
UsersApiService Updating a user
This method updates users from given attributes and renders it in case of success.  Authorization​: only users that can update users. If the user affiliated to the token has the &#x60;manage_users_of_my_teams&#x60; permission, the updated user will need to belong to at least one of the teams he’s the leader of. The teams the user affiliated to the token is the leader of will be the only ones which can be added or removed.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param userId
 * @param optional nil or *UpdateUserOpts - Optional Parameters:
 * @param "CategoryIds" (optional.Interface of []string) -  User list of category ids (multiple).
 * @param "Email" (optional.String) -  User email.
 * @param "Enabled" (optional.Bool) -  Whether the user is enabled or not (boolean).
 * @param "ExternalId" (optional.String) -  User external id, used for SSO.
 * @param "Firstname" (optional.String) -  User firstname.
 * @param "Gender" (optional.String) -  User gender (\"man\" or \"woman\").
 * @param "IdentityIds" (optional.Interface of []string) -  User list of identity ids (multiple).
 * @param "Lastname" (optional.String) -  User lastname.
 * @param "Locale" (optional.String) -  Language for the user interface.
 * @param "Nickname" (optional.String) -  User nickname.
 * @param "Password" (optional.String) -  User plain password.
 * @param "RoleId" (optional.String) -  User role id.
 * @param "TeamIds" (optional.Interface of []string) -  User list of team ids (multiple).
 * @param "Timezone" (optional.String) -  Use the timezone endpoint to get the timezone name (String), default is empty for domain timezone.
 * @param "SpokenLanguages" (optional.Interface of []string) -  List of locales corresponding to the languages spoken by the user (multiple).
@return User
*/

type UpdateUserOpts struct {
	CategoryIds     optional.Interface
	Email           optional.String
	Enabled         optional.Bool
	ExternalId      optional.String
	Firstname       optional.String
	Gender          optional.String
	IdentityIds     optional.Interface
	Lastname        optional.String
	Locale          optional.String
	Nickname        optional.String
	Password        optional.String
	RoleId          optional.String
	TeamIds         optional.Interface
	Timezone        optional.String
	SpokenLanguages optional.Interface
}

func (a *UsersApiService) UpdateUser(ctx context.Context, userId string, localVarOptionals *UpdateUserOpts) (User, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  User
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/users/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", fmt.Sprintf("%v", userId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.CategoryIds.IsSet() {
		t := localVarOptionals.CategoryIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("category_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("category_ids[]", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Email.IsSet() {
		localVarQueryParams.Add("email", parameterToString(localVarOptionals.Email.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Enabled.IsSet() {
		localVarQueryParams.Add("enabled", parameterToString(localVarOptionals.Enabled.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ExternalId.IsSet() {
		localVarQueryParams.Add("external_id", parameterToString(localVarOptionals.ExternalId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Firstname.IsSet() {
		localVarQueryParams.Add("firstname", parameterToString(localVarOptionals.Firstname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Gender.IsSet() {
		localVarQueryParams.Add("gender", parameterToString(localVarOptionals.Gender.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IdentityIds.IsSet() {
		t := localVarOptionals.IdentityIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("identity_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("identity_ids[]", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Lastname.IsSet() {
		localVarQueryParams.Add("lastname", parameterToString(localVarOptionals.Lastname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Locale.IsSet() {
		localVarQueryParams.Add("locale", parameterToString(localVarOptionals.Locale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nickname.IsSet() {
		localVarQueryParams.Add("nickname", parameterToString(localVarOptionals.Nickname.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
		localVarQueryParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RoleId.IsSet() {
		localVarQueryParams.Add("role_id", parameterToString(localVarOptionals.RoleId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TeamIds.IsSet() {
		t := localVarOptionals.TeamIds.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("team_ids[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("team_ids[]", parameterToString(t, "multi"))
		}
	}
	if localVarOptionals != nil && localVarOptionals.Timezone.IsSet() {
		localVarQueryParams.Add("timezone", parameterToString(localVarOptionals.Timezone.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SpokenLanguages.IsSet() {
		t := localVarOptionals.SpokenLanguages.Value()
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("spoken_languages[]", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("spoken_languages[]", parameterToString(t, "multi"))
		}
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v User
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
