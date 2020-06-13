# FabricUplinkRoleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. | [optional] [default to "Auto"]
**Fec** | Pointer to **string** | Forward error correction configuration for the port. | [optional] [default to "Auto"]
**UdldAdminState** | Pointer to **string** | Admin configured state for UDLD for this port. | [optional] [default to "Disabled"]

## Methods

### NewFabricUplinkRoleAllOf

`func NewFabricUplinkRoleAllOf() *FabricUplinkRoleAllOf`

NewFabricUplinkRoleAllOf instantiates a new FabricUplinkRoleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricUplinkRoleAllOfWithDefaults

`func NewFabricUplinkRoleAllOfWithDefaults() *FabricUplinkRoleAllOf`

NewFabricUplinkRoleAllOfWithDefaults instantiates a new FabricUplinkRoleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *FabricUplinkRoleAllOf) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricUplinkRoleAllOf) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricUplinkRoleAllOf) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricUplinkRoleAllOf) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFec

`func (o *FabricUplinkRoleAllOf) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricUplinkRoleAllOf) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricUplinkRoleAllOf) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricUplinkRoleAllOf) HasFec() bool`

HasFec returns a boolean if a field has been set.

### GetUdldAdminState

`func (o *FabricUplinkRoleAllOf) GetUdldAdminState() string`

GetUdldAdminState returns the UdldAdminState field if non-nil, zero value otherwise.

### GetUdldAdminStateOk

`func (o *FabricUplinkRoleAllOf) GetUdldAdminStateOk() (*string, bool)`

GetUdldAdminStateOk returns a tuple with the UdldAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldAdminState

`func (o *FabricUplinkRoleAllOf) SetUdldAdminState(v string)`

SetUdldAdminState sets UdldAdminState field to given value.

### HasUdldAdminState

`func (o *FabricUplinkRoleAllOf) HasUdldAdminState() bool`

HasUdldAdminState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


