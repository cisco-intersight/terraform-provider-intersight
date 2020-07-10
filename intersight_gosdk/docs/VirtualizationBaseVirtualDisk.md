# VirtualizationBaseVirtualDisk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int64** | Disk capacity represented in bytes. | [optional] 
**Name** | Pointer to **string** | Name of the storage disk. Name must be unique within a Datastore. | [optional] 

## Methods

### NewVirtualizationBaseVirtualDisk

`func NewVirtualizationBaseVirtualDisk() *VirtualizationBaseVirtualDisk`

NewVirtualizationBaseVirtualDisk instantiates a new VirtualizationBaseVirtualDisk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationBaseVirtualDiskWithDefaults

`func NewVirtualizationBaseVirtualDiskWithDefaults() *VirtualizationBaseVirtualDisk`

NewVirtualizationBaseVirtualDiskWithDefaults instantiates a new VirtualizationBaseVirtualDisk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VirtualizationBaseVirtualDisk) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VirtualizationBaseVirtualDisk) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VirtualizationBaseVirtualDisk) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VirtualizationBaseVirtualDisk) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationBaseVirtualDisk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationBaseVirtualDisk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationBaseVirtualDisk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationBaseVirtualDisk) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


