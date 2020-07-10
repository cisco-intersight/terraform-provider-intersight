# HyperflexHxapVirtualDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessMode** | Pointer to **string** | Access mode of the virtual disk. | [optional] [readonly] [default to "ReadWriteOnce"]
**Mode** | Pointer to **string** | File mode of the disk  example - Filesystem, Block. | [optional] [readonly] [default to "Block"]
**SourceFilePath** | Pointer to **string** | Source file path associated with virtual machine disk. | [optional] [readonly] 
**SourceVirtualDisk** | Pointer to **string** | Source disk name from where the clone is done. | [optional] [readonly] 
**Status** | Pointer to [**HyperflexDiskStatus**](hyperflex.DiskStatus.md) |  | [optional] 
**Uuid** | Pointer to **string** | UUID of the virtual disk. | [optional] [readonly] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 
**VirtualMachine** | Pointer to [**HyperflexHxapVirtualMachineRelationship**](hyperflex.HxapVirtualMachine.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapVirtualDiskAllOf

`func NewHyperflexHxapVirtualDiskAllOf() *HyperflexHxapVirtualDiskAllOf`

NewHyperflexHxapVirtualDiskAllOf instantiates a new HyperflexHxapVirtualDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapVirtualDiskAllOfWithDefaults

`func NewHyperflexHxapVirtualDiskAllOfWithDefaults() *HyperflexHxapVirtualDiskAllOf`

NewHyperflexHxapVirtualDiskAllOfWithDefaults instantiates a new HyperflexHxapVirtualDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessMode

`func (o *HyperflexHxapVirtualDiskAllOf) GetAccessMode() string`

GetAccessMode returns the AccessMode field if non-nil, zero value otherwise.

### GetAccessModeOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetAccessModeOk() (*string, bool)`

GetAccessModeOk returns a tuple with the AccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMode

`func (o *HyperflexHxapVirtualDiskAllOf) SetAccessMode(v string)`

SetAccessMode sets AccessMode field to given value.

### HasAccessMode

`func (o *HyperflexHxapVirtualDiskAllOf) HasAccessMode() bool`

HasAccessMode returns a boolean if a field has been set.

### GetMode

`func (o *HyperflexHxapVirtualDiskAllOf) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *HyperflexHxapVirtualDiskAllOf) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *HyperflexHxapVirtualDiskAllOf) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSourceFilePath

`func (o *HyperflexHxapVirtualDiskAllOf) GetSourceFilePath() string`

GetSourceFilePath returns the SourceFilePath field if non-nil, zero value otherwise.

### GetSourceFilePathOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetSourceFilePathOk() (*string, bool)`

GetSourceFilePathOk returns a tuple with the SourceFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceFilePath

`func (o *HyperflexHxapVirtualDiskAllOf) SetSourceFilePath(v string)`

SetSourceFilePath sets SourceFilePath field to given value.

### HasSourceFilePath

`func (o *HyperflexHxapVirtualDiskAllOf) HasSourceFilePath() bool`

HasSourceFilePath returns a boolean if a field has been set.

### GetSourceVirtualDisk

`func (o *HyperflexHxapVirtualDiskAllOf) GetSourceVirtualDisk() string`

GetSourceVirtualDisk returns the SourceVirtualDisk field if non-nil, zero value otherwise.

### GetSourceVirtualDiskOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetSourceVirtualDiskOk() (*string, bool)`

GetSourceVirtualDiskOk returns a tuple with the SourceVirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVirtualDisk

`func (o *HyperflexHxapVirtualDiskAllOf) SetSourceVirtualDisk(v string)`

SetSourceVirtualDisk sets SourceVirtualDisk field to given value.

### HasSourceVirtualDisk

`func (o *HyperflexHxapVirtualDiskAllOf) HasSourceVirtualDisk() bool`

HasSourceVirtualDisk returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxapVirtualDiskAllOf) GetStatus() HyperflexDiskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetStatusOk() (*HyperflexDiskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxapVirtualDiskAllOf) SetStatus(v HyperflexDiskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxapVirtualDiskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *HyperflexHxapVirtualDiskAllOf) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *HyperflexHxapVirtualDiskAllOf) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *HyperflexHxapVirtualDiskAllOf) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapVirtualDiskAllOf) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapVirtualDiskAllOf) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapVirtualDiskAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetVirtualMachine

`func (o *HyperflexHxapVirtualDiskAllOf) GetVirtualMachine() HyperflexHxapVirtualMachineRelationship`

GetVirtualMachine returns the VirtualMachine field if non-nil, zero value otherwise.

### GetVirtualMachineOk

`func (o *HyperflexHxapVirtualDiskAllOf) GetVirtualMachineOk() (*HyperflexHxapVirtualMachineRelationship, bool)`

GetVirtualMachineOk returns a tuple with the VirtualMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachine

`func (o *HyperflexHxapVirtualDiskAllOf) SetVirtualMachine(v HyperflexHxapVirtualMachineRelationship)`

SetVirtualMachine sets VirtualMachine field to given value.

### HasVirtualMachine

`func (o *HyperflexHxapVirtualDiskAllOf) HasVirtualMachine() bool`

HasVirtualMachine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


