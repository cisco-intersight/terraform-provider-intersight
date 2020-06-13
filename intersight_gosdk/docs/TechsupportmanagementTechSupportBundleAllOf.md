# TechsupportmanagementTechSupportBundleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIdentifier** | Pointer to **string** | The device identifier used to uniquely identify an individual device. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The device type obtained from the inventory. | [optional] [readonly] 
**Pid** | Pointer to **string** | Product identification of the device. | [optional] 
**PlatformParam** | Pointer to [**ConnectorPlatformParamBase**](connector.PlatformParamBase.md) |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type of the device. | [optional] [default to ""]
**Serial** | Pointer to **string** | Serial number of the device. | [optional] 
**DeviceRegistration** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**TargetResource** | Pointer to [**MoBaseMoRelationship**](mo.BaseMo.Relationship.md) |  | [optional] 
**TechSupportStatus** | Pointer to [**TechsupportmanagementTechSupportStatusRelationship**](techsupportmanagement.TechSupportStatus.Relationship.md) |  | [optional] 

## Methods

### NewTechsupportmanagementTechSupportBundleAllOf

`func NewTechsupportmanagementTechSupportBundleAllOf() *TechsupportmanagementTechSupportBundleAllOf`

NewTechsupportmanagementTechSupportBundleAllOf instantiates a new TechsupportmanagementTechSupportBundleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportBundleAllOfWithDefaults

`func NewTechsupportmanagementTechSupportBundleAllOfWithDefaults() *TechsupportmanagementTechSupportBundleAllOf`

NewTechsupportmanagementTechSupportBundleAllOfWithDefaults instantiates a new TechsupportmanagementTechSupportBundleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParam() ConnectorPlatformParamBase`

GetPlatformParam returns the PlatformParam field if non-nil, zero value otherwise.

### GetPlatformParamOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool)`

GetPlatformParamOk returns a tuple with the PlatformParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformParam(v ConnectorPlatformParamBase)`

SetPlatformParam sets PlatformParam field to given value.

### HasPlatformParam

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformParam() bool`

HasPlatformParam returns a boolean if a field has been set.

### GetPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResource() MoBaseMoRelationship`

GetTargetResource returns the TargetResource field if non-nil, zero value otherwise.

### GetTargetResourceOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTargetResourceOk() (*MoBaseMoRelationship, bool)`

GetTargetResourceOk returns a tuple with the TargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetTargetResource(v MoBaseMoRelationship)`

SetTargetResource sets TargetResource field to given value.

### HasTargetResource

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasTargetResource() bool`

HasTargetResource returns a boolean if a field has been set.

### GetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship`

GetTechSupportStatus returns the TechSupportStatus field if non-nil, zero value otherwise.

### GetTechSupportStatusOk

`func (o *TechsupportmanagementTechSupportBundleAllOf) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool)`

GetTechSupportStatusOk returns a tuple with the TechSupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship)`

SetTechSupportStatus sets TechSupportStatus field to given value.

### HasTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundleAllOf) HasTechSupportStatus() bool`

HasTechSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


