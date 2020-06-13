# FabricFcoeUplinkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. | [optional] [default to "Auto"]
**Fec** | Pointer to **string** | Forward error correction configuration for the port. | [optional] [default to "Auto"]
**UdldAdminState** | Pointer to **string** | Admin configured state for UDLD for this port. | [optional] [default to "Disabled"]

## Methods

### NewFabricFcoeUplinkRole

`func NewFabricFcoeUplinkRole() *FabricFcoeUplinkRole`

NewFabricFcoeUplinkRole instantiates a new FabricFcoeUplinkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcoeUplinkRoleWithDefaults

`func NewFabricFcoeUplinkRoleWithDefaults() *FabricFcoeUplinkRole`

NewFabricFcoeUplinkRoleWithDefaults instantiates a new FabricFcoeUplinkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *FabricFcoeUplinkRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcoeUplinkRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcoeUplinkRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcoeUplinkRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFec

`func (o *FabricFcoeUplinkRole) GetFec() string`

GetFec returns the Fec field if non-nil, zero value otherwise.

### GetFecOk

`func (o *FabricFcoeUplinkRole) GetFecOk() (*string, bool)`

GetFecOk returns a tuple with the Fec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFec

`func (o *FabricFcoeUplinkRole) SetFec(v string)`

SetFec sets Fec field to given value.

### HasFec

`func (o *FabricFcoeUplinkRole) HasFec() bool`

HasFec returns a boolean if a field has been set.

### GetUdldAdminState

`func (o *FabricFcoeUplinkRole) GetUdldAdminState() string`

GetUdldAdminState returns the UdldAdminState field if non-nil, zero value otherwise.

### GetUdldAdminStateOk

`func (o *FabricFcoeUplinkRole) GetUdldAdminStateOk() (*string, bool)`

GetUdldAdminStateOk returns a tuple with the UdldAdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdldAdminState

`func (o *FabricFcoeUplinkRole) SetUdldAdminState(v string)`

SetUdldAdminState sets UdldAdminState field to given value.

### HasUdldAdminState

`func (o *FabricFcoeUplinkRole) HasUdldAdminState() bool`

HasUdldAdminState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


