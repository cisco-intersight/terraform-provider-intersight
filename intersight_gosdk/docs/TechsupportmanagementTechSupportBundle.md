# TechsupportmanagementTechSupportBundle

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

### NewTechsupportmanagementTechSupportBundle

`func NewTechsupportmanagementTechSupportBundle() *TechsupportmanagementTechSupportBundle`

NewTechsupportmanagementTechSupportBundle instantiates a new TechsupportmanagementTechSupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTechsupportmanagementTechSupportBundleWithDefaults

`func NewTechsupportmanagementTechSupportBundleWithDefaults() *TechsupportmanagementTechSupportBundle`

NewTechsupportmanagementTechSupportBundleWithDefaults instantiates a new TechsupportmanagementTechSupportBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetDeviceType

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetPid

`func (o *TechsupportmanagementTechSupportBundle) GetPid() string`

GetPid returns the Pid field if non-nil, zero value otherwise.

### GetPidOk

`func (o *TechsupportmanagementTechSupportBundle) GetPidOk() (*string, bool)`

GetPidOk returns a tuple with the Pid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPid

`func (o *TechsupportmanagementTechSupportBundle) SetPid(v string)`

SetPid sets Pid field to given value.

### HasPid

`func (o *TechsupportmanagementTechSupportBundle) HasPid() bool`

HasPid returns a boolean if a field has been set.

### GetPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformParam() ConnectorPlatformParamBase`

GetPlatformParam returns the PlatformParam field if non-nil, zero value otherwise.

### GetPlatformParamOk

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformParamOk() (*ConnectorPlatformParamBase, bool)`

GetPlatformParamOk returns a tuple with the PlatformParam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) SetPlatformParam(v ConnectorPlatformParamBase)`

SetPlatformParam sets PlatformParam field to given value.

### HasPlatformParam

`func (o *TechsupportmanagementTechSupportBundle) HasPlatformParam() bool`

HasPlatformParam returns a boolean if a field has been set.

### GetPlatformType

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *TechsupportmanagementTechSupportBundle) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *TechsupportmanagementTechSupportBundle) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *TechsupportmanagementTechSupportBundle) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetSerial

`func (o *TechsupportmanagementTechSupportBundle) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *TechsupportmanagementTechSupportBundle) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *TechsupportmanagementTechSupportBundle) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *TechsupportmanagementTechSupportBundle) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistration() AssetDeviceRegistrationRelationship`

GetDeviceRegistration returns the DeviceRegistration field if non-nil, zero value otherwise.

### GetDeviceRegistrationOk

`func (o *TechsupportmanagementTechSupportBundle) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool)`

GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) SetDeviceRegistration(v AssetDeviceRegistrationRelationship)`

SetDeviceRegistration sets DeviceRegistration field to given value.

### HasDeviceRegistration

`func (o *TechsupportmanagementTechSupportBundle) HasDeviceRegistration() bool`

HasDeviceRegistration returns a boolean if a field has been set.

### GetTargetResource

`func (o *TechsupportmanagementTechSupportBundle) GetTargetResource() MoBaseMoRelationship`

GetTargetResource returns the TargetResource field if non-nil, zero value otherwise.

### GetTargetResourceOk

`func (o *TechsupportmanagementTechSupportBundle) GetTargetResourceOk() (*MoBaseMoRelationship, bool)`

GetTargetResourceOk returns a tuple with the TargetResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResource

`func (o *TechsupportmanagementTechSupportBundle) SetTargetResource(v MoBaseMoRelationship)`

SetTargetResource sets TargetResource field to given value.

### HasTargetResource

`func (o *TechsupportmanagementTechSupportBundle) HasTargetResource() bool`

HasTargetResource returns a boolean if a field has been set.

### GetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatus() TechsupportmanagementTechSupportStatusRelationship`

GetTechSupportStatus returns the TechSupportStatus field if non-nil, zero value otherwise.

### GetTechSupportStatusOk

`func (o *TechsupportmanagementTechSupportBundle) GetTechSupportStatusOk() (*TechsupportmanagementTechSupportStatusRelationship, bool)`

GetTechSupportStatusOk returns a tuple with the TechSupportStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) SetTechSupportStatus(v TechsupportmanagementTechSupportStatusRelationship)`

SetTechSupportStatus sets TechSupportStatus field to given value.

### HasTechSupportStatus

`func (o *TechsupportmanagementTechSupportBundle) HasTechSupportStatus() bool`

HasTechSupportStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


