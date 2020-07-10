# OauthOauthUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;oauth.OauthUser&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]OauthOauthUser**](oauth.OauthUser.md) | The array of &#39;oauth.OauthUser&#39; resources matching the request. | [optional] 

## Methods

### NewOauthOauthUserList

`func NewOauthOauthUserList() *OauthOauthUserList`

NewOauthOauthUserList instantiates a new OauthOauthUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauthOauthUserListWithDefaults

`func NewOauthOauthUserListWithDefaults() *OauthOauthUserList`

NewOauthOauthUserListWithDefaults instantiates a new OauthOauthUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *OauthOauthUserList) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *OauthOauthUserList) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *OauthOauthUserList) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *OauthOauthUserList) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *OauthOauthUserList) GetResults() []OauthOauthUser`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *OauthOauthUserList) GetResultsOk() (*[]OauthOauthUser, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *OauthOauthUserList) SetResults(v []OauthOauthUser)`

SetResults sets Results field to given value.

### HasResults

`func (o *OauthOauthUserList) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *OauthOauthUserList) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *OauthOauthUserList) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


