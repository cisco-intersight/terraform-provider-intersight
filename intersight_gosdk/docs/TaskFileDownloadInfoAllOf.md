# TaskFileDownloadInfoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contents** | Pointer to **string** | When the type of the download is inline, it will read the file from the contents here. | [optional] 
**Path** | Pointer to **string** | The path of the download from the specified source location for type S3, then this is the object key. | [optional] 
**Source** | Pointer to **string** | The source of the download location and if type is S3, then this is the bucket name. | [optional] 
**Type** | Pointer to **string** | The type of download location is captured in type. | [optional] [default to "S3"]

## Methods

### NewTaskFileDownloadInfoAllOf

`func NewTaskFileDownloadInfoAllOf() *TaskFileDownloadInfoAllOf`

NewTaskFileDownloadInfoAllOf instantiates a new TaskFileDownloadInfoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskFileDownloadInfoAllOfWithDefaults

`func NewTaskFileDownloadInfoAllOfWithDefaults() *TaskFileDownloadInfoAllOf`

NewTaskFileDownloadInfoAllOfWithDefaults instantiates a new TaskFileDownloadInfoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContents

`func (o *TaskFileDownloadInfoAllOf) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *TaskFileDownloadInfoAllOf) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *TaskFileDownloadInfoAllOf) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *TaskFileDownloadInfoAllOf) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetPath

`func (o *TaskFileDownloadInfoAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *TaskFileDownloadInfoAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *TaskFileDownloadInfoAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *TaskFileDownloadInfoAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetSource

`func (o *TaskFileDownloadInfoAllOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TaskFileDownloadInfoAllOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TaskFileDownloadInfoAllOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TaskFileDownloadInfoAllOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *TaskFileDownloadInfoAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskFileDownloadInfoAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskFileDownloadInfoAllOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskFileDownloadInfoAllOf) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


