# InventoryUemConnectionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | Connections status of Uem endpoint. | [optional] [readonly] [default to "Unknown"]
**EpInfo** | Pointer to [**InventoryEpInfo**](inventory.EpInfo.md) |  | [optional] 
**EpType** | Pointer to **string** | Type of Uem endpoint (BMC, CMC or VIC). | [optional] [readonly] 
**Ip** | Pointer to **string** | The IP address of the Uem endpoint. | [optional] [readonly] 
**MemberIdentity** | Pointer to **string** | The member identity of the UEM connection being inventoried. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the Uem endpoint. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the Uem endpoint. | [optional] [readonly] 
**TargetMoId** | Pointer to **string** | The MoId address of the Uem endpoint. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor of the Uem endpoint. | [optional] [readonly] 
**ComputeBlade** | Pointer to [**ComputeBladeRelationship**](compute.Blade.Relationship.md) |  | [optional] 
**ComputeRackUnit** | Pointer to [**ComputeRackUnitRelationship**](compute.RackUnit.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewInventoryUemConnectionAllOf

`func NewInventoryUemConnectionAllOf() *InventoryUemConnectionAllOf`

NewInventoryUemConnectionAllOf instantiates a new InventoryUemConnectionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryUemConnectionAllOfWithDefaults

`func NewInventoryUemConnectionAllOfWithDefaults() *InventoryUemConnectionAllOf`

NewInventoryUemConnectionAllOfWithDefaults instantiates a new InventoryUemConnectionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *InventoryUemConnectionAllOf) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *InventoryUemConnectionAllOf) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *InventoryUemConnectionAllOf) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *InventoryUemConnectionAllOf) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetEpInfo

`func (o *InventoryUemConnectionAllOf) GetEpInfo() InventoryEpInfo`

GetEpInfo returns the EpInfo field if non-nil, zero value otherwise.

### GetEpInfoOk

`func (o *InventoryUemConnectionAllOf) GetEpInfoOk() (*InventoryEpInfo, bool)`

GetEpInfoOk returns a tuple with the EpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpInfo

`func (o *InventoryUemConnectionAllOf) SetEpInfo(v InventoryEpInfo)`

SetEpInfo sets EpInfo field to given value.

### HasEpInfo

`func (o *InventoryUemConnectionAllOf) HasEpInfo() bool`

HasEpInfo returns a boolean if a field has been set.

### GetEpType

`func (o *InventoryUemConnectionAllOf) GetEpType() string`

GetEpType returns the EpType field if non-nil, zero value otherwise.

### GetEpTypeOk

`func (o *InventoryUemConnectionAllOf) GetEpTypeOk() (*string, bool)`

GetEpTypeOk returns a tuple with the EpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpType

`func (o *InventoryUemConnectionAllOf) SetEpType(v string)`

SetEpType sets EpType field to given value.

### HasEpType

`func (o *InventoryUemConnectionAllOf) HasEpType() bool`

HasEpType returns a boolean if a field has been set.

### GetIp

`func (o *InventoryUemConnectionAllOf) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InventoryUemConnectionAllOf) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InventoryUemConnectionAllOf) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InventoryUemConnectionAllOf) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *InventoryUemConnectionAllOf) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *InventoryUemConnectionAllOf) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *InventoryUemConnectionAllOf) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *InventoryUemConnectionAllOf) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetModel

`func (o *InventoryUemConnectionAllOf) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InventoryUemConnectionAllOf) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InventoryUemConnectionAllOf) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InventoryUemConnectionAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *InventoryUemConnectionAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InventoryUemConnectionAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InventoryUemConnectionAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InventoryUemConnectionAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTargetMoId

`func (o *InventoryUemConnectionAllOf) GetTargetMoId() string`

GetTargetMoId returns the TargetMoId field if non-nil, zero value otherwise.

### GetTargetMoIdOk

`func (o *InventoryUemConnectionAllOf) GetTargetMoIdOk() (*string, bool)`

GetTargetMoIdOk returns a tuple with the TargetMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoId

`func (o *InventoryUemConnectionAllOf) SetTargetMoId(v string)`

SetTargetMoId sets TargetMoId field to given value.

### HasTargetMoId

`func (o *InventoryUemConnectionAllOf) HasTargetMoId() bool`

HasTargetMoId returns a boolean if a field has been set.

### GetVendor

`func (o *InventoryUemConnectionAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InventoryUemConnectionAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InventoryUemConnectionAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InventoryUemConnectionAllOf) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetComputeBlade

`func (o *InventoryUemConnectionAllOf) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *InventoryUemConnectionAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *InventoryUemConnectionAllOf) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *InventoryUemConnectionAllOf) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *InventoryUemConnectionAllOf) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *InventoryUemConnectionAllOf) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *InventoryUemConnectionAllOf) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *InventoryUemConnectionAllOf) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryUemConnectionAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryUemConnectionAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryUemConnectionAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryUemConnectionAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


