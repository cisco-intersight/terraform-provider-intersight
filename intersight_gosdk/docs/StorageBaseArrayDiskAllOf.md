# StorageBaseArrayDiskAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Disk name available in storage array. | [optional] [readonly] 
**PartNumber** | Pointer to **string** | Storage disk part number. | [optional] [readonly] 
**Protocol** | Pointer to **string** | Storage protocol used in disk for communication. Possible values are SAS, SATA and NVMe. | [optional] [readonly] [default to "Unknown"]
**Speed** | Pointer to **int64** | Disk speed for read or write operation measured in rpm. | [optional] [readonly] 
**Status** | Pointer to **string** | Storage disk health status. | [optional] [readonly] [default to "Unknown"]
**StorageUtilization** | Pointer to [**StorageBaseCapacity**](storage.BaseCapacity.md) |  | [optional] 
**Type** | Pointer to **string** | Storage disk type - it can be SSD, HDD, NVRAM. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Storage disk version number. | [optional] [readonly] 

## Methods

### NewStorageBaseArrayDiskAllOf

`func NewStorageBaseArrayDiskAllOf() *StorageBaseArrayDiskAllOf`

NewStorageBaseArrayDiskAllOf instantiates a new StorageBaseArrayDiskAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayDiskAllOfWithDefaults

`func NewStorageBaseArrayDiskAllOfWithDefaults() *StorageBaseArrayDiskAllOf`

NewStorageBaseArrayDiskAllOfWithDefaults instantiates a new StorageBaseArrayDiskAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageBaseArrayDiskAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseArrayDiskAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseArrayDiskAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseArrayDiskAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartNumber

`func (o *StorageBaseArrayDiskAllOf) GetPartNumber() string`

GetPartNumber returns the PartNumber field if non-nil, zero value otherwise.

### GetPartNumberOk

`func (o *StorageBaseArrayDiskAllOf) GetPartNumberOk() (*string, bool)`

GetPartNumberOk returns a tuple with the PartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartNumber

`func (o *StorageBaseArrayDiskAllOf) SetPartNumber(v string)`

SetPartNumber sets PartNumber field to given value.

### HasPartNumber

`func (o *StorageBaseArrayDiskAllOf) HasPartNumber() bool`

HasPartNumber returns a boolean if a field has been set.

### GetProtocol

`func (o *StorageBaseArrayDiskAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *StorageBaseArrayDiskAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *StorageBaseArrayDiskAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *StorageBaseArrayDiskAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSpeed

`func (o *StorageBaseArrayDiskAllOf) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StorageBaseArrayDiskAllOf) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StorageBaseArrayDiskAllOf) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StorageBaseArrayDiskAllOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBaseArrayDiskAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBaseArrayDiskAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBaseArrayDiskAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBaseArrayDiskAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) GetStorageUtilization() StorageBaseCapacity`

GetStorageUtilization returns the StorageUtilization field if non-nil, zero value otherwise.

### GetStorageUtilizationOk

`func (o *StorageBaseArrayDiskAllOf) GetStorageUtilizationOk() (*StorageBaseCapacity, bool)`

GetStorageUtilizationOk returns a tuple with the StorageUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) SetStorageUtilization(v StorageBaseCapacity)`

SetStorageUtilization sets StorageUtilization field to given value.

### HasStorageUtilization

`func (o *StorageBaseArrayDiskAllOf) HasStorageUtilization() bool`

HasStorageUtilization returns a boolean if a field has been set.

### GetType

`func (o *StorageBaseArrayDiskAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StorageBaseArrayDiskAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StorageBaseArrayDiskAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *StorageBaseArrayDiskAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *StorageBaseArrayDiskAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageBaseArrayDiskAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageBaseArrayDiskAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageBaseArrayDiskAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


