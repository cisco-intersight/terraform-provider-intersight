# FabricFcUplinkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminSpeed** | Pointer to **string** | Admin configured speed for the port. | [optional] [default to "Auto"]
**FillPattern** | Pointer to **string** | Fill pattern to differentiate the configs in NPIV. | [optional] [default to "Idle"]
**VsanId** | Pointer to **int64** | Virtual San Identifier associated to the FC port. | [optional] 

## Methods

### NewFabricFcUplinkRole

`func NewFabricFcUplinkRole() *FabricFcUplinkRole`

NewFabricFcUplinkRole instantiates a new FabricFcUplinkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricFcUplinkRoleWithDefaults

`func NewFabricFcUplinkRoleWithDefaults() *FabricFcUplinkRole`

NewFabricFcUplinkRoleWithDefaults instantiates a new FabricFcUplinkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminSpeed

`func (o *FabricFcUplinkRole) GetAdminSpeed() string`

GetAdminSpeed returns the AdminSpeed field if non-nil, zero value otherwise.

### GetAdminSpeedOk

`func (o *FabricFcUplinkRole) GetAdminSpeedOk() (*string, bool)`

GetAdminSpeedOk returns a tuple with the AdminSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSpeed

`func (o *FabricFcUplinkRole) SetAdminSpeed(v string)`

SetAdminSpeed sets AdminSpeed field to given value.

### HasAdminSpeed

`func (o *FabricFcUplinkRole) HasAdminSpeed() bool`

HasAdminSpeed returns a boolean if a field has been set.

### GetFillPattern

`func (o *FabricFcUplinkRole) GetFillPattern() string`

GetFillPattern returns the FillPattern field if non-nil, zero value otherwise.

### GetFillPatternOk

`func (o *FabricFcUplinkRole) GetFillPatternOk() (*string, bool)`

GetFillPatternOk returns a tuple with the FillPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFillPattern

`func (o *FabricFcUplinkRole) SetFillPattern(v string)`

SetFillPattern sets FillPattern field to given value.

### HasFillPattern

`func (o *FabricFcUplinkRole) HasFillPattern() bool`

HasFillPattern returns a boolean if a field has been set.

### GetVsanId

`func (o *FabricFcUplinkRole) GetVsanId() int64`

GetVsanId returns the VsanId field if non-nil, zero value otherwise.

### GetVsanIdOk

`func (o *FabricFcUplinkRole) GetVsanIdOk() (*int64, bool)`

GetVsanIdOk returns a tuple with the VsanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanId

`func (o *FabricFcUplinkRole) SetVsanId(v int64)`

SetVsanId sets VsanId field to given value.

### HasVsanId

`func (o *FabricFcUplinkRole) HasVsanId() bool`

HasVsanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


