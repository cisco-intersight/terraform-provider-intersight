# HyperflexVmDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootOrder** | Pointer to **int64** | Boot order for this disk. | [optional] [readonly] 
**Bus** | Pointer to **string** | Virtual machine brige name of interface. | [optional] [readonly] [default to "virtio"]
**Name** | Pointer to **string** | Disk name associated with virtual machine. | [optional] [readonly] 
**Type** | Pointer to **string** | Type of the Disk (hdd, cdrom). | [optional] [readonly] [default to "hdd"]
**VirtualDisk** | Pointer to [**HyperflexVdiskConfig**](hyperflex.VdiskConfig.md) |  | [optional] 
**VirtualDiskReference** | Pointer to **string** | Virtual disk reference name. | [optional] [readonly] 

## Methods

### NewHyperflexVmDisk

`func NewHyperflexVmDisk() *HyperflexVmDisk`

NewHyperflexVmDisk instantiates a new HyperflexVmDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmDiskWithDefaults

`func NewHyperflexVmDiskWithDefaults() *HyperflexVmDisk`

NewHyperflexVmDiskWithDefaults instantiates a new HyperflexVmDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootOrder

`func (o *HyperflexVmDisk) GetBootOrder() int64`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *HyperflexVmDisk) GetBootOrderOk() (*int64, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *HyperflexVmDisk) SetBootOrder(v int64)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *HyperflexVmDisk) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.

### GetBus

`func (o *HyperflexVmDisk) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *HyperflexVmDisk) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *HyperflexVmDisk) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *HyperflexVmDisk) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *HyperflexVmDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexVmDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexVmDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexVmDisk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *HyperflexVmDisk) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexVmDisk) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexVmDisk) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexVmDisk) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *HyperflexVmDisk) GetVirtualDisk() HyperflexVdiskConfig`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *HyperflexVmDisk) GetVirtualDiskOk() (*HyperflexVdiskConfig, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *HyperflexVmDisk) SetVirtualDisk(v HyperflexVdiskConfig)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *HyperflexVmDisk) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### GetVirtualDiskReference

`func (o *HyperflexVmDisk) GetVirtualDiskReference() string`

GetVirtualDiskReference returns the VirtualDiskReference field if non-nil, zero value otherwise.

### GetVirtualDiskReferenceOk

`func (o *HyperflexVmDisk) GetVirtualDiskReferenceOk() (*string, bool)`

GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskReference

`func (o *HyperflexVmDisk) SetVirtualDiskReference(v string)`

SetVirtualDiskReference sets VirtualDiskReference field to given value.

### HasVirtualDiskReference

`func (o *HyperflexVmDisk) HasVirtualDiskReference() bool`

HasVirtualDiskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


