# InventoryUemConnection

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

### NewInventoryUemConnection

`func NewInventoryUemConnection() *InventoryUemConnection`

NewInventoryUemConnection instantiates a new InventoryUemConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryUemConnectionWithDefaults

`func NewInventoryUemConnectionWithDefaults() *InventoryUemConnection`

NewInventoryUemConnectionWithDefaults instantiates a new InventoryUemConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *InventoryUemConnection) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *InventoryUemConnection) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *InventoryUemConnection) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *InventoryUemConnection) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetEpInfo

`func (o *InventoryUemConnection) GetEpInfo() InventoryEpInfo`

GetEpInfo returns the EpInfo field if non-nil, zero value otherwise.

### GetEpInfoOk

`func (o *InventoryUemConnection) GetEpInfoOk() (*InventoryEpInfo, bool)`

GetEpInfoOk returns a tuple with the EpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpInfo

`func (o *InventoryUemConnection) SetEpInfo(v InventoryEpInfo)`

SetEpInfo sets EpInfo field to given value.

### HasEpInfo

`func (o *InventoryUemConnection) HasEpInfo() bool`

HasEpInfo returns a boolean if a field has been set.

### GetEpType

`func (o *InventoryUemConnection) GetEpType() string`

GetEpType returns the EpType field if non-nil, zero value otherwise.

### GetEpTypeOk

`func (o *InventoryUemConnection) GetEpTypeOk() (*string, bool)`

GetEpTypeOk returns a tuple with the EpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpType

`func (o *InventoryUemConnection) SetEpType(v string)`

SetEpType sets EpType field to given value.

### HasEpType

`func (o *InventoryUemConnection) HasEpType() bool`

HasEpType returns a boolean if a field has been set.

### GetIp

`func (o *InventoryUemConnection) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InventoryUemConnection) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InventoryUemConnection) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InventoryUemConnection) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *InventoryUemConnection) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *InventoryUemConnection) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *InventoryUemConnection) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *InventoryUemConnection) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetModel

`func (o *InventoryUemConnection) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InventoryUemConnection) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InventoryUemConnection) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InventoryUemConnection) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *InventoryUemConnection) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InventoryUemConnection) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InventoryUemConnection) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InventoryUemConnection) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetTargetMoId

`func (o *InventoryUemConnection) GetTargetMoId() string`

GetTargetMoId returns the TargetMoId field if non-nil, zero value otherwise.

### GetTargetMoIdOk

`func (o *InventoryUemConnection) GetTargetMoIdOk() (*string, bool)`

GetTargetMoIdOk returns a tuple with the TargetMoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMoId

`func (o *InventoryUemConnection) SetTargetMoId(v string)`

SetTargetMoId sets TargetMoId field to given value.

### HasTargetMoId

`func (o *InventoryUemConnection) HasTargetMoId() bool`

HasTargetMoId returns a boolean if a field has been set.

### GetVendor

`func (o *InventoryUemConnection) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InventoryUemConnection) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InventoryUemConnection) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InventoryUemConnection) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetComputeBlade

`func (o *InventoryUemConnection) GetComputeBlade() ComputeBladeRelationship`

GetComputeBlade returns the ComputeBlade field if non-nil, zero value otherwise.

### GetComputeBladeOk

`func (o *InventoryUemConnection) GetComputeBladeOk() (*ComputeBladeRelationship, bool)`

GetComputeBladeOk returns a tuple with the ComputeBlade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeBlade

`func (o *InventoryUemConnection) SetComputeBlade(v ComputeBladeRelationship)`

SetComputeBlade sets ComputeBlade field to given value.

### HasComputeBlade

`func (o *InventoryUemConnection) HasComputeBlade() bool`

HasComputeBlade returns a boolean if a field has been set.

### GetComputeRackUnit

`func (o *InventoryUemConnection) GetComputeRackUnit() ComputeRackUnitRelationship`

GetComputeRackUnit returns the ComputeRackUnit field if non-nil, zero value otherwise.

### GetComputeRackUnitOk

`func (o *InventoryUemConnection) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool)`

GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeRackUnit

`func (o *InventoryUemConnection) SetComputeRackUnit(v ComputeRackUnitRelationship)`

SetComputeRackUnit sets ComputeRackUnit field to given value.

### HasComputeRackUnit

`func (o *InventoryUemConnection) HasComputeRackUnit() bool`

HasComputeRackUnit returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *InventoryUemConnection) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *InventoryUemConnection) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *InventoryUemConnection) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *InventoryUemConnection) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


