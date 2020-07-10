# VirtualizationActionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | Description of failure i.e. derived from the workflow failure message. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the Action performed on a resource like VM, Disk etc. | [optional] [readonly] 
**Status** | Pointer to **string** | Status of the Action like InProgress, Success, Failure etc. | [optional] [readonly] [default to "None"]

## Methods

### NewVirtualizationActionInfo

`func NewVirtualizationActionInfo() *VirtualizationActionInfo`

NewVirtualizationActionInfo instantiates a new VirtualizationActionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationActionInfoWithDefaults

`func NewVirtualizationActionInfoWithDefaults() *VirtualizationActionInfo`

NewVirtualizationActionInfoWithDefaults instantiates a new VirtualizationActionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *VirtualizationActionInfo) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *VirtualizationActionInfo) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *VirtualizationActionInfo) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *VirtualizationActionInfo) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationActionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationActionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationActionInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationActionInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualizationActionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualizationActionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualizationActionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualizationActionInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


