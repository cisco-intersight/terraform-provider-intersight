# PoolAbstractPoolMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assigned** | Pointer to **bool** | Boolean to represent whether the ID is assigned or not. | [optional] 
**AssignedToMoid** | Pointer to **string** | Moid of the entity/server profile that owns this ID. | [optional] 
**AssignedToType** | Pointer to **string** | Type of the entity that owns this ID. | [optional] 

## Methods

### NewPoolAbstractPoolMember

`func NewPoolAbstractPoolMember() *PoolAbstractPoolMember`

NewPoolAbstractPoolMember instantiates a new PoolAbstractPoolMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolAbstractPoolMemberWithDefaults

`func NewPoolAbstractPoolMemberWithDefaults() *PoolAbstractPoolMember`

NewPoolAbstractPoolMemberWithDefaults instantiates a new PoolAbstractPoolMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssigned

`func (o *PoolAbstractPoolMember) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *PoolAbstractPoolMember) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *PoolAbstractPoolMember) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *PoolAbstractPoolMember) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetAssignedToMoid

`func (o *PoolAbstractPoolMember) GetAssignedToMoid() string`

GetAssignedToMoid returns the AssignedToMoid field if non-nil, zero value otherwise.

### GetAssignedToMoidOk

`func (o *PoolAbstractPoolMember) GetAssignedToMoidOk() (*string, bool)`

GetAssignedToMoidOk returns a tuple with the AssignedToMoid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToMoid

`func (o *PoolAbstractPoolMember) SetAssignedToMoid(v string)`

SetAssignedToMoid sets AssignedToMoid field to given value.

### HasAssignedToMoid

`func (o *PoolAbstractPoolMember) HasAssignedToMoid() bool`

HasAssignedToMoid returns a boolean if a field has been set.

### GetAssignedToType

`func (o *PoolAbstractPoolMember) GetAssignedToType() string`

GetAssignedToType returns the AssignedToType field if non-nil, zero value otherwise.

### GetAssignedToTypeOk

`func (o *PoolAbstractPoolMember) GetAssignedToTypeOk() (*string, bool)`

GetAssignedToTypeOk returns a tuple with the AssignedToType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedToType

`func (o *PoolAbstractPoolMember) SetAssignedToType(v string)`

SetAssignedToType sets AssignedToType field to given value.

### HasAssignedToType

`func (o *PoolAbstractPoolMember) HasAssignedToType() bool`

HasAssignedToType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


