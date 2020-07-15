# SoftwarerepositoryCachedImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to be performed on the imported file. If &#39;PreCache&#39; is set, the image will be cached in endpoint. If &#39;Evict&#39; is set, the cached file will be removed. Applicable in Intersight appliance deployment. If &#39;Cancel&#39; is set, the ImportState is marked as &#39;Failed&#39;. Applicable for local machine source. | [optional] [default to "None"]
**CacheState** | Pointer to **string** | The state  of this file in the endpoint The importState is updated during the cache operation and as part of the cache monitoring process. | [optional] [readonly] [default to "ReadyForImport"]
**CachedTime** | Pointer to [**time.Time**](time.Time.md) | The time when the image or file was cached into the FI storage. | [optional] [readonly] 
**LastAccessTime** | Pointer to [**time.Time**](time.Time.md) | Used by the cache monitoring process to determine the files that are to be evicted from the cache. | [optional] [readonly] 
**Md5sum** | Pointer to **string** | The MD5 sum of the firmware image that will be used by the endpoint to validate the integrity of the image. | [optional] [readonly] 
**OriginalSha512sum** | Pointer to **string** | The actual sha512sum of the cached image. | [optional] [readonly] 
**Path** | Pointer to **string** | The absolute path of the imported file in the endpoint. | [optional] [readonly] 
**RegisteredWorkflows** | Pointer to **[]string** |  | [optional] 
**UsedCount** | Pointer to **int64** | The number of times this file has been used to copy or upgrade or install actions. Used by the cache monitoring process to determine the files to be evicted from the cache. | [optional] [readonly] 
**File** | Pointer to [**SoftwarerepositoryFileRelationship**](softwarerepository.File.Relationship.md) |  | [optional] 
**NetworkElement** | Pointer to [**NetworkElementRelationship**](network.Element.Relationship.md) |  | [optional] 

## Methods

### NewSoftwarerepositoryCachedImage

`func NewSoftwarerepositoryCachedImage() *SoftwarerepositoryCachedImage`

NewSoftwarerepositoryCachedImage instantiates a new SoftwarerepositoryCachedImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwarerepositoryCachedImageWithDefaults

`func NewSoftwarerepositoryCachedImageWithDefaults() *SoftwarerepositoryCachedImage`

NewSoftwarerepositoryCachedImageWithDefaults instantiates a new SoftwarerepositoryCachedImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SoftwarerepositoryCachedImage) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SoftwarerepositoryCachedImage) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SoftwarerepositoryCachedImage) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *SoftwarerepositoryCachedImage) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCacheState

`func (o *SoftwarerepositoryCachedImage) GetCacheState() string`

GetCacheState returns the CacheState field if non-nil, zero value otherwise.

### GetCacheStateOk

`func (o *SoftwarerepositoryCachedImage) GetCacheStateOk() (*string, bool)`

GetCacheStateOk returns a tuple with the CacheState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheState

`func (o *SoftwarerepositoryCachedImage) SetCacheState(v string)`

SetCacheState sets CacheState field to given value.

### HasCacheState

`func (o *SoftwarerepositoryCachedImage) HasCacheState() bool`

HasCacheState returns a boolean if a field has been set.

### GetCachedTime

`func (o *SoftwarerepositoryCachedImage) GetCachedTime() time.Time`

GetCachedTime returns the CachedTime field if non-nil, zero value otherwise.

### GetCachedTimeOk

`func (o *SoftwarerepositoryCachedImage) GetCachedTimeOk() (*time.Time, bool)`

GetCachedTimeOk returns a tuple with the CachedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTime

`func (o *SoftwarerepositoryCachedImage) SetCachedTime(v time.Time)`

SetCachedTime sets CachedTime field to given value.

### HasCachedTime

`func (o *SoftwarerepositoryCachedImage) HasCachedTime() bool`

HasCachedTime returns a boolean if a field has been set.

### GetLastAccessTime

`func (o *SoftwarerepositoryCachedImage) GetLastAccessTime() time.Time`

GetLastAccessTime returns the LastAccessTime field if non-nil, zero value otherwise.

### GetLastAccessTimeOk

`func (o *SoftwarerepositoryCachedImage) GetLastAccessTimeOk() (*time.Time, bool)`

GetLastAccessTimeOk returns a tuple with the LastAccessTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessTime

`func (o *SoftwarerepositoryCachedImage) SetLastAccessTime(v time.Time)`

SetLastAccessTime sets LastAccessTime field to given value.

### HasLastAccessTime

`func (o *SoftwarerepositoryCachedImage) HasLastAccessTime() bool`

HasLastAccessTime returns a boolean if a field has been set.

### GetMd5sum

`func (o *SoftwarerepositoryCachedImage) GetMd5sum() string`

GetMd5sum returns the Md5sum field if non-nil, zero value otherwise.

### GetMd5sumOk

`func (o *SoftwarerepositoryCachedImage) GetMd5sumOk() (*string, bool)`

GetMd5sumOk returns a tuple with the Md5sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5sum

`func (o *SoftwarerepositoryCachedImage) SetMd5sum(v string)`

SetMd5sum sets Md5sum field to given value.

### HasMd5sum

`func (o *SoftwarerepositoryCachedImage) HasMd5sum() bool`

HasMd5sum returns a boolean if a field has been set.

### GetOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) GetOriginalSha512sum() string`

GetOriginalSha512sum returns the OriginalSha512sum field if non-nil, zero value otherwise.

### GetOriginalSha512sumOk

`func (o *SoftwarerepositoryCachedImage) GetOriginalSha512sumOk() (*string, bool)`

GetOriginalSha512sumOk returns a tuple with the OriginalSha512sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) SetOriginalSha512sum(v string)`

SetOriginalSha512sum sets OriginalSha512sum field to given value.

### HasOriginalSha512sum

`func (o *SoftwarerepositoryCachedImage) HasOriginalSha512sum() bool`

HasOriginalSha512sum returns a boolean if a field has been set.

### GetPath

`func (o *SoftwarerepositoryCachedImage) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SoftwarerepositoryCachedImage) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SoftwarerepositoryCachedImage) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *SoftwarerepositoryCachedImage) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) GetRegisteredWorkflows() []string`

GetRegisteredWorkflows returns the RegisteredWorkflows field if non-nil, zero value otherwise.

### GetRegisteredWorkflowsOk

`func (o *SoftwarerepositoryCachedImage) GetRegisteredWorkflowsOk() (*[]string, bool)`

GetRegisteredWorkflowsOk returns a tuple with the RegisteredWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) SetRegisteredWorkflows(v []string)`

SetRegisteredWorkflows sets RegisteredWorkflows field to given value.

### HasRegisteredWorkflows

`func (o *SoftwarerepositoryCachedImage) HasRegisteredWorkflows() bool`

HasRegisteredWorkflows returns a boolean if a field has been set.

### GetUsedCount

`func (o *SoftwarerepositoryCachedImage) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *SoftwarerepositoryCachedImage) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *SoftwarerepositoryCachedImage) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *SoftwarerepositoryCachedImage) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetFile

`func (o *SoftwarerepositoryCachedImage) GetFile() SoftwarerepositoryFileRelationship`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *SoftwarerepositoryCachedImage) GetFileOk() (*SoftwarerepositoryFileRelationship, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *SoftwarerepositoryCachedImage) SetFile(v SoftwarerepositoryFileRelationship)`

SetFile sets File field to given value.

### HasFile

`func (o *SoftwarerepositoryCachedImage) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetNetworkElement

`func (o *SoftwarerepositoryCachedImage) GetNetworkElement() NetworkElementRelationship`

GetNetworkElement returns the NetworkElement field if non-nil, zero value otherwise.

### GetNetworkElementOk

`func (o *SoftwarerepositoryCachedImage) GetNetworkElementOk() (*NetworkElementRelationship, bool)`

GetNetworkElementOk returns a tuple with the NetworkElement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkElement

`func (o *SoftwarerepositoryCachedImage) SetNetworkElement(v NetworkElementRelationship)`

SetNetworkElement sets NetworkElement field to given value.

### HasNetworkElement

`func (o *SoftwarerepositoryCachedImage) HasNetworkElement() bool`

HasNetworkElement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


