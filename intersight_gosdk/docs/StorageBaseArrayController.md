# StorageBaseArrayController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Storage array controller name. | [optional] [readonly] 
**OperationalMode** | Pointer to **string** | Controller running mode, Primary or Secondary. | [optional] [readonly] [default to "Unknown"]
**Status** | Pointer to **string** | Status of the storage controller. | [optional] [readonly] [default to "Unknown"]
**Version** | Pointer to **string** | Software version running on a storage controller. | [optional] [readonly] 

## Methods

### NewStorageBaseArrayController

`func NewStorageBaseArrayController() *StorageBaseArrayController`

NewStorageBaseArrayController instantiates a new StorageBaseArrayController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageBaseArrayControllerWithDefaults

`func NewStorageBaseArrayControllerWithDefaults() *StorageBaseArrayController`

NewStorageBaseArrayControllerWithDefaults instantiates a new StorageBaseArrayController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StorageBaseArrayController) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageBaseArrayController) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageBaseArrayController) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageBaseArrayController) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperationalMode

`func (o *StorageBaseArrayController) GetOperationalMode() string`

GetOperationalMode returns the OperationalMode field if non-nil, zero value otherwise.

### GetOperationalModeOk

`func (o *StorageBaseArrayController) GetOperationalModeOk() (*string, bool)`

GetOperationalModeOk returns a tuple with the OperationalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationalMode

`func (o *StorageBaseArrayController) SetOperationalMode(v string)`

SetOperationalMode sets OperationalMode field to given value.

### HasOperationalMode

`func (o *StorageBaseArrayController) HasOperationalMode() bool`

HasOperationalMode returns a boolean if a field has been set.

### GetStatus

`func (o *StorageBaseArrayController) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StorageBaseArrayController) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StorageBaseArrayController) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StorageBaseArrayController) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *StorageBaseArrayController) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *StorageBaseArrayController) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *StorageBaseArrayController) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *StorageBaseArrayController) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


