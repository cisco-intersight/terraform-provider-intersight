# HyperflexHxapEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstTime** | Pointer to **string** | First timestamp of the event occurrence. | [optional] 
**Identity** | Pointer to **string** | Internally generated identity (UUID) of this event. | [optional] 
**LastTime** | Pointer to **string** | Last timestamp of the event occurrence. | [optional] 
**Message** | Pointer to **string** | Full description of the event. | [optional] 
**OwnerName** | Pointer to **string** | Name of the owner with which event is associated. | [optional] 
**OwnerType** | Pointer to **string** | Type of the object with which event is associated (Host, Cluster, VM). | [optional] [default to "Unknown"]
**OwnerUuid** | Pointer to **string** | UUID of the owner with which event is associated. | [optional] 
**Severity** | Pointer to **string** | Severity level of the event (Info/Warning/Critical). | [optional] [default to "None"]
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapEvent

`func NewHyperflexHxapEvent() *HyperflexHxapEvent`

NewHyperflexHxapEvent instantiates a new HyperflexHxapEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapEventWithDefaults

`func NewHyperflexHxapEventWithDefaults() *HyperflexHxapEvent`

NewHyperflexHxapEventWithDefaults instantiates a new HyperflexHxapEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstTime

`func (o *HyperflexHxapEvent) GetFirstTime() string`

GetFirstTime returns the FirstTime field if non-nil, zero value otherwise.

### GetFirstTimeOk

`func (o *HyperflexHxapEvent) GetFirstTimeOk() (*string, bool)`

GetFirstTimeOk returns a tuple with the FirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstTime

`func (o *HyperflexHxapEvent) SetFirstTime(v string)`

SetFirstTime sets FirstTime field to given value.

### HasFirstTime

`func (o *HyperflexHxapEvent) HasFirstTime() bool`

HasFirstTime returns a boolean if a field has been set.

### GetIdentity

`func (o *HyperflexHxapEvent) GetIdentity() string`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *HyperflexHxapEvent) GetIdentityOk() (*string, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *HyperflexHxapEvent) SetIdentity(v string)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *HyperflexHxapEvent) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLastTime

`func (o *HyperflexHxapEvent) GetLastTime() string`

GetLastTime returns the LastTime field if non-nil, zero value otherwise.

### GetLastTimeOk

`func (o *HyperflexHxapEvent) GetLastTimeOk() (*string, bool)`

GetLastTimeOk returns a tuple with the LastTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTime

`func (o *HyperflexHxapEvent) SetLastTime(v string)`

SetLastTime sets LastTime field to given value.

### HasLastTime

`func (o *HyperflexHxapEvent) HasLastTime() bool`

HasLastTime returns a boolean if a field has been set.

### GetMessage

`func (o *HyperflexHxapEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HyperflexHxapEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HyperflexHxapEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HyperflexHxapEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOwnerName

`func (o *HyperflexHxapEvent) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *HyperflexHxapEvent) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *HyperflexHxapEvent) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *HyperflexHxapEvent) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *HyperflexHxapEvent) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *HyperflexHxapEvent) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *HyperflexHxapEvent) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *HyperflexHxapEvent) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetOwnerUuid

`func (o *HyperflexHxapEvent) GetOwnerUuid() string`

GetOwnerUuid returns the OwnerUuid field if non-nil, zero value otherwise.

### GetOwnerUuidOk

`func (o *HyperflexHxapEvent) GetOwnerUuidOk() (*string, bool)`

GetOwnerUuidOk returns a tuple with the OwnerUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerUuid

`func (o *HyperflexHxapEvent) SetOwnerUuid(v string)`

SetOwnerUuid sets OwnerUuid field to given value.

### HasOwnerUuid

`func (o *HyperflexHxapEvent) HasOwnerUuid() bool`

HasOwnerUuid returns a boolean if a field has been set.

### GetSeverity

`func (o *HyperflexHxapEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *HyperflexHxapEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *HyperflexHxapEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *HyperflexHxapEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapEvent) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapEvent) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapEvent) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapEvent) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


