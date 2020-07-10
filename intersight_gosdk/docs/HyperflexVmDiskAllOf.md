# HyperflexVmDiskAllOf

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

### NewHyperflexVmDiskAllOf

`func NewHyperflexVmDiskAllOf() *HyperflexVmDiskAllOf`

NewHyperflexVmDiskAllOf instantiates a new HyperflexVmDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmDiskAllOfWithDefaults

`func NewHyperflexVmDiskAllOfWithDefaults() *HyperflexVmDiskAllOf`

NewHyperflexVmDiskAllOfWithDefaults instantiates a new HyperflexVmDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootOrder

`func (o *HyperflexVmDiskAllOf) GetBootOrder() int64`

GetBootOrder returns the BootOrder field if non-nil, zero value otherwise.

### GetBootOrderOk

`func (o *HyperflexVmDiskAllOf) GetBootOrderOk() (*int64, bool)`

GetBootOrderOk returns a tuple with the BootOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootOrder

`func (o *HyperflexVmDiskAllOf) SetBootOrder(v int64)`

SetBootOrder sets BootOrder field to given value.

### HasBootOrder

`func (o *HyperflexVmDiskAllOf) HasBootOrder() bool`

HasBootOrder returns a boolean if a field has been set.

### GetBus

`func (o *HyperflexVmDiskAllOf) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *HyperflexVmDiskAllOf) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *HyperflexVmDiskAllOf) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *HyperflexVmDiskAllOf) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *HyperflexVmDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HyperflexVmDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HyperflexVmDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HyperflexVmDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *HyperflexVmDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HyperflexVmDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HyperflexVmDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *HyperflexVmDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *HyperflexVmDiskAllOf) GetVirtualDisk() HyperflexVdiskConfig`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *HyperflexVmDiskAllOf) GetVirtualDiskOk() (*HyperflexVdiskConfig, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *HyperflexVmDiskAllOf) SetVirtualDisk(v HyperflexVdiskConfig)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *HyperflexVmDiskAllOf) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### GetVirtualDiskReference

`func (o *HyperflexVmDiskAllOf) GetVirtualDiskReference() string`

GetVirtualDiskReference returns the VirtualDiskReference field if non-nil, zero value otherwise.

### GetVirtualDiskReferenceOk

`func (o *HyperflexVmDiskAllOf) GetVirtualDiskReferenceOk() (*string, bool)`

GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskReference

`func (o *HyperflexVmDiskAllOf) SetVirtualDiskReference(v string)`

SetVirtualDiskReference sets VirtualDiskReference field to given value.

### HasVirtualDiskReference

`func (o *HyperflexVmDiskAllOf) HasVirtualDiskReference() bool`

HasVirtualDiskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


