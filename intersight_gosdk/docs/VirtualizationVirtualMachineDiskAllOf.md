# VirtualizationVirtualMachineDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bus** | Pointer to **string** | Disk bus name given for a virtual machine. | [optional] [default to "virtio"]
**Name** | Pointer to **string** | Virtual machine network bridge name. | [optional] 
**Order** | Pointer to **int64** | Priority order of the disk. | [optional] 
**Type** | Pointer to **string** | Disk type hdd or cdrom for a virtual machine. | [optional] [default to "hdd"]
**VirtualDisk** | Pointer to [**VirtualizationVirtualDiskConfig**](virtualization.VirtualDiskConfig.md) |  | [optional] 
**VirtualDiskReference** | Pointer to **string** | Name of the existing virtual disk to be attached to the Virtual Machine. | [optional] 

## Methods

### NewVirtualizationVirtualMachineDiskAllOf

`func NewVirtualizationVirtualMachineDiskAllOf() *VirtualizationVirtualMachineDiskAllOf`

NewVirtualizationVirtualMachineDiskAllOf instantiates a new VirtualizationVirtualMachineDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachineDiskAllOfWithDefaults

`func NewVirtualizationVirtualMachineDiskAllOfWithDefaults() *VirtualizationVirtualMachineDiskAllOf`

NewVirtualizationVirtualMachineDiskAllOfWithDefaults instantiates a new VirtualizationVirtualMachineDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBus

`func (o *VirtualizationVirtualMachineDiskAllOf) GetBus() string`

GetBus returns the Bus field if non-nil, zero value otherwise.

### GetBusOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetBusOk() (*string, bool)`

GetBusOk returns a tuple with the Bus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBus

`func (o *VirtualizationVirtualMachineDiskAllOf) SetBus(v string)`

SetBus sets Bus field to given value.

### HasBus

`func (o *VirtualizationVirtualMachineDiskAllOf) HasBus() bool`

HasBus returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualMachineDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachineDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualMachineDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrder

`func (o *VirtualizationVirtualMachineDiskAllOf) GetOrder() int64`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetOrderOk() (*int64, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *VirtualizationVirtualMachineDiskAllOf) SetOrder(v int64)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *VirtualizationVirtualMachineDiskAllOf) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetType

`func (o *VirtualizationVirtualMachineDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VirtualizationVirtualMachineDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VirtualizationVirtualMachineDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVirtualDisk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDisk() VirtualizationVirtualDiskConfig`

GetVirtualDisk returns the VirtualDisk field if non-nil, zero value otherwise.

### GetVirtualDiskOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskOk() (*VirtualizationVirtualDiskConfig, bool)`

GetVirtualDiskOk returns a tuple with the VirtualDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisk

`func (o *VirtualizationVirtualMachineDiskAllOf) SetVirtualDisk(v VirtualizationVirtualDiskConfig)`

SetVirtualDisk sets VirtualDisk field to given value.

### HasVirtualDisk

`func (o *VirtualizationVirtualMachineDiskAllOf) HasVirtualDisk() bool`

HasVirtualDisk returns a boolean if a field has been set.

### GetVirtualDiskReference

`func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskReference() string`

GetVirtualDiskReference returns the VirtualDiskReference field if non-nil, zero value otherwise.

### GetVirtualDiskReferenceOk

`func (o *VirtualizationVirtualMachineDiskAllOf) GetVirtualDiskReferenceOk() (*string, bool)`

GetVirtualDiskReferenceOk returns a tuple with the VirtualDiskReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDiskReference

`func (o *VirtualizationVirtualMachineDiskAllOf) SetVirtualDiskReference(v string)`

SetVirtualDiskReference sets VirtualDiskReference field to given value.

### HasVirtualDiskReference

`func (o *VirtualizationVirtualMachineDiskAllOf) HasVirtualDiskReference() bool`

HasVirtualDiskReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


