# HyperflexDiskStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadPercentage** | Pointer to **string** | Percentage of download completed. | [optional] [readonly] 
**State** | Pointer to **string** | Current state of the virtual disk. | [optional] [readonly] [default to "Unknown"]
**VolumeHandle** | Pointer to **string** | Identity of the Volume associated with virtual machine disk. | [optional] [readonly] 

## Methods

### NewHyperflexDiskStatusAllOf

`func NewHyperflexDiskStatusAllOf() *HyperflexDiskStatusAllOf`

NewHyperflexDiskStatusAllOf instantiates a new HyperflexDiskStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexDiskStatusAllOfWithDefaults

`func NewHyperflexDiskStatusAllOfWithDefaults() *HyperflexDiskStatusAllOf`

NewHyperflexDiskStatusAllOfWithDefaults instantiates a new HyperflexDiskStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadPercentage

`func (o *HyperflexDiskStatusAllOf) GetDownloadPercentage() string`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *HyperflexDiskStatusAllOf) GetDownloadPercentageOk() (*string, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *HyperflexDiskStatusAllOf) SetDownloadPercentage(v string)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *HyperflexDiskStatusAllOf) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetState

`func (o *HyperflexDiskStatusAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HyperflexDiskStatusAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HyperflexDiskStatusAllOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *HyperflexDiskStatusAllOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVolumeHandle

`func (o *HyperflexDiskStatusAllOf) GetVolumeHandle() string`

GetVolumeHandle returns the VolumeHandle field if non-nil, zero value otherwise.

### GetVolumeHandleOk

`func (o *HyperflexDiskStatusAllOf) GetVolumeHandleOk() (*string, bool)`

GetVolumeHandleOk returns a tuple with the VolumeHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeHandle

`func (o *HyperflexDiskStatusAllOf) SetVolumeHandle(v string)`

SetVolumeHandle sets VolumeHandle field to given value.

### HasVolumeHandle

`func (o *HyperflexDiskStatusAllOf) HasVolumeHandle() bool`

HasVolumeHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


