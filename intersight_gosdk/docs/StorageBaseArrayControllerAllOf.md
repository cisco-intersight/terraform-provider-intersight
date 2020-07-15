# StorageBaseArrayControllerAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Storage array controller name. | [optional] [readonly] 
**OperationalMode** | Pointer to **string** | Controller running mode, Primary or Secondary. | [optional] [readonly] [default to "Unknown"]
**Status** | Pointer to **string** | Status of the storage controller. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Software version running on a storage controller. | [optional] [readonly] 

## Methods

### NewStorageBaseArrayControllerAllOf

`func NewStorageBaseArrayControllerAllOf() *StorageBaseArrayControllerAllOf`

NewStorageBaseArrayControllerAllOf instantiates a new StorageBaseArrayControllerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayControllerAllOfWithDefaults

`func NewStorageBaseArrayControllerAllOfWithDefaults() *StorageBaseArrayControllerAllOf`

NewStorageBaseArrayControllerAllOfWithDefaults instantiates a new StorageBaseArrayControllerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageBaseArrayControllerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseArrayControllerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseArrayControllerAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseArrayControllerAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalMode

`func (o *StorageBaseArrayControllerAllOf) GetOperationalMode() string`

GetOperationalMode returns the OperationalMode field if non-nil, zero value otherwise.

### GetOperationalModeOk

`func (o *StorageBaseArrayControllerAllOf) GetOperationalModeOk() (*string, bool)`

GetOperationalModeOk returns a tuple with the OperationalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalMode

`func (o *StorageBaseArrayControllerAllOf) SetOperationalMode(v string)`

SetOperationalMode sets OperationalMode field to given value.

### HasOperationalMode

`func (o *StorageBaseArrayControllerAllOf) HasOperationalMode() bool`

HasOperationalMode returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBaseArrayControllerAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBaseArrayControllerAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBaseArrayControllerAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBaseArrayControllerAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *StorageBaseArrayControllerAllOf) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageBaseArrayControllerAllOf) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageBaseArrayControllerAllOf) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageBaseArrayControllerAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


