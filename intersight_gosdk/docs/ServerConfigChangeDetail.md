# ServerConfigChangeDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Profile** | Pointer to [**ServerProfileRelationship**](server.Profile.Relationship.md) |  | [optional] 

## Methods

### NewServerConfigChangeDetail

`func NewServerConfigChangeDetail() *ServerConfigChangeDetail`

NewServerConfigChangeDetail instantiates a new ServerConfigChangeDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerConfigChangeDetailWithDefaults

`func NewServerConfigChangeDetailWithDefaults() *ServerConfigChangeDetail`

NewServerConfigChangeDetailWithDefaults instantiates a new ServerConfigChangeDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfile

`func (o *ServerConfigChangeDetail) GetProfile() ServerProfileRelationship`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ServerConfigChangeDetail) GetProfileOk() (*ServerProfileRelationship, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ServerConfigChangeDetail) SetProfile(v ServerProfileRelationship)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ServerConfigChangeDetail) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


