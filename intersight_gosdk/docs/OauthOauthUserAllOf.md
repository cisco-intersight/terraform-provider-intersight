# OauthOauthUserAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessTokens** | Pointer to [**[]OauthAccessToken**](oauth.AccessToken.md) |  | [optional] 
**UserId** | Pointer to **string** | User name | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewOauthOauthUserAllOf

`func NewOauthOauthUserAllOf() *OauthOauthUserAllOf`

NewOauthOauthUserAllOf instantiates a new OauthOauthUserAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthOauthUserAllOfWithDefaults

`func NewOauthOauthUserAllOfWithDefaults() *OauthOauthUserAllOf`

NewOauthOauthUserAllOfWithDefaults instantiates a new OauthOauthUserAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessTokens

`func (o *OauthOauthUserAllOf) GetAccessTokens() []OauthAccessToken`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *OauthOauthUserAllOf) GetAccessTokensOk() (*[]OauthAccessToken, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *OauthOauthUserAllOf) SetAccessTokens(v []OauthAccessToken)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *OauthOauthUserAllOf) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.

### GetUserId

`func (o *OauthOauthUserAllOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OauthOauthUserAllOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OauthOauthUserAllOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *OauthOauthUserAllOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetAccount

`func (o *OauthOauthUserAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OauthOauthUserAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OauthOauthUserAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OauthOauthUserAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


