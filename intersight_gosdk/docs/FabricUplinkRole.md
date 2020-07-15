# FabricUplinkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. | [optional] [default to "Auto"]
**Fec** | Pointer to **string** | Forward error correction configuration for the port. | [optional] [default to "Auto"]
**UdldAdminState** | Pointer to **string** | Admin configured state for UDLD for this port. | [optional] [default to "Disabled"]

## Methods

### NewFabricUplinkRole

`func NewFabricUplinkRole() *FabricUplinkRole`

NewFabricUplinkRole instantiates a new FabricUplinkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUplinkRoleWithDefaults

`func NewFabricUplinkRoleWithDefaults() *FabricUplinkRole`

NewFabricUplinkRoleWithDefaults instantiates a new FabricUplinkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *FabricUplinkRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricUplinkRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricUplinkRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricUplinkRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFec

`func (o *FabricUplinkRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricUplinkRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricUplinkRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricUplinkRole) HasFec() bool`

HasFec returns a boolean if a field has been set.

### GetUdldAdminState

`func (o *FabricUplinkRole) GetUdldAdminState() string`

GetUdldAdminState returns the UdldAdminState field if non-nil, zero value otherwise.

### GetUdldAdminStateOk

`func (o *FabricUplinkRole) GetUdldAdminStateOk() (*string, bool)`

GetUdldAdminStateOk returns a tuple with the UdldAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldAdminState

`func (o *FabricUplinkRole) SetUdldAdminState(v string)`

SetUdldAdminState sets UdldAdminState field to given value.

### HasUdldAdminState

`func (o *FabricUplinkRole) HasUdldAdminState() bool`

HasUdldAdminState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


