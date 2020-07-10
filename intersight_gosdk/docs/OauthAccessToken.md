# OauthAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiType** | Pointer to **string** | Oauth api type | [optional] [default to "Unknown"]
**ExpiryTime** | Pointer to [**time.Time**](time.Time.md) | Access token expiry time | [optional] [readonly] 
**Status** | Pointer to **string** | Access token status | [optional] [default to "Inactive"]
**Token** | Pointer to **string** | Access token | [optional] [readonly] 

## Methods

### NewOauthAccessToken

`func NewOauthAccessToken() *OauthAccessToken`

NewOauthAccessToken instantiates a new OauthAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthAccessTokenWithDefaults

`func NewOauthAccessTokenWithDefaults() *OauthAccessToken`

NewOauthAccessTokenWithDefaults instantiates a new OauthAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiType

`func (o *OauthAccessToken) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OauthAccessToken) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OauthAccessToken) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OauthAccessToken) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetExpiryTime

`func (o *OauthAccessToken) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *OauthAccessToken) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *OauthAccessToken) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *OauthAccessToken) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetStatus

`func (o *OauthAccessToken) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OauthAccessToken) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OauthAccessToken) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OauthAccessToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *OauthAccessToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OauthAccessToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OauthAccessToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OauthAccessToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


