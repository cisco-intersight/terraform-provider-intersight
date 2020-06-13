# FirmwareUpgradeStatusAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownloadError** | Pointer to **string** | The error message from the endpoint during the download. | [optional] 
**DownloadMessage** | Pointer to **interface{}** | The message from the endpoint during the download.} type: string | [optional] 
**DownloadPercentage** | Pointer to **int64** | The percentage of the image downloaded in the endpoint. | [optional] 
**DownloadProgress** | Pointer to **int64** | The download progress of the file represented as a percentage between 0% and 100%. If progress reporting is not possible a value of -1 is sent. | [optional] 
**DownloadRetries** | Pointer to **int64** | The number of retries the plugin attempted before succeeding or failing the download. | [optional] 
**DownloadStage** | Pointer to **string** | The image download stages. Example:downloading, flashing. | [optional] 
**EpPowerStatus** | Pointer to **string** | The server power status after the upgrade request is submitted in the endpoint. | [optional] [default to "none"]
**OverallError** | Pointer to **string** | The reason for the operation failure. | [optional] 
**OverallPercentage** | Pointer to **int64** | The overall percentage of the operation. | [optional] 
**Overallstatus** | Pointer to **string** | The overall status of the operation. | [optional] [default to "none"]
**PendingType** | Pointer to **string** | Pending reason for the upgrade waiting. | [optional] [default to "none"]
**Sha256checksum** | Pointer to **string** | The sha256checksum of the downloaded file as calculated by the download plugin after successfully downloading a file. | [optional] 
**Upgrade** | Pointer to [**FirmwareUpgradeBaseRelationship**](firmware.UpgradeBase.Relationship.md) |  | [optional] 
**Workflow** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewFirmwareUpgradeStatusAllOf

`func NewFirmwareUpgradeStatusAllOf() *FirmwareUpgradeStatusAllOf`

NewFirmwareUpgradeStatusAllOf instantiates a new FirmwareUpgradeStatusAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpgradeStatusAllOfWithDefaults

`func NewFirmwareUpgradeStatusAllOfWithDefaults() *FirmwareUpgradeStatusAllOf`

NewFirmwareUpgradeStatusAllOfWithDefaults instantiates a new FirmwareUpgradeStatusAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownloadError

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadError() string`

GetDownloadError returns the DownloadError field if non-nil, zero value otherwise.

### GetDownloadErrorOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadErrorOk() (*string, bool)`

GetDownloadErrorOk returns a tuple with the DownloadError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadError

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadError(v string)`

SetDownloadError sets DownloadError field to given value.

### HasDownloadError

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadError() bool`

HasDownloadError returns a boolean if a field has been set.

### GetDownloadMessage

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessage() interface{}`

GetDownloadMessage returns the DownloadMessage field if non-nil, zero value otherwise.

### GetDownloadMessageOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadMessageOk() (*interface{}, bool)`

GetDownloadMessageOk returns a tuple with the DownloadMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadMessage

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadMessage(v interface{})`

SetDownloadMessage sets DownloadMessage field to given value.

### HasDownloadMessage

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadMessage() bool`

HasDownloadMessage returns a boolean if a field has been set.

### SetDownloadMessageNil

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadMessageNil(b bool)`

 SetDownloadMessageNil sets the value for DownloadMessage to be an explicit nil

### UnsetDownloadMessage
`func (o *FirmwareUpgradeStatusAllOf) UnsetDownloadMessage()`

UnsetDownloadMessage ensures that no value is present for DownloadMessage, not even an explicit nil
### GetDownloadPercentage

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentage() int64`

GetDownloadPercentage returns the DownloadPercentage field if non-nil, zero value otherwise.

### GetDownloadPercentageOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadPercentageOk() (*int64, bool)`

GetDownloadPercentageOk returns a tuple with the DownloadPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadPercentage

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadPercentage(v int64)`

SetDownloadPercentage sets DownloadPercentage field to given value.

### HasDownloadPercentage

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadPercentage() bool`

HasDownloadPercentage returns a boolean if a field has been set.

### GetDownloadProgress

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadProgress() int64`

GetDownloadProgress returns the DownloadProgress field if non-nil, zero value otherwise.

### GetDownloadProgressOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadProgressOk() (*int64, bool)`

GetDownloadProgressOk returns a tuple with the DownloadProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadProgress

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadProgress(v int64)`

SetDownloadProgress sets DownloadProgress field to given value.

### HasDownloadProgress

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadProgress() bool`

HasDownloadProgress returns a boolean if a field has been set.

### GetDownloadRetries

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadRetries() int64`

GetDownloadRetries returns the DownloadRetries field if non-nil, zero value otherwise.

### GetDownloadRetriesOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadRetriesOk() (*int64, bool)`

GetDownloadRetriesOk returns a tuple with the DownloadRetries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadRetries

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadRetries(v int64)`

SetDownloadRetries sets DownloadRetries field to given value.

### HasDownloadRetries

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadRetries() bool`

HasDownloadRetries returns a boolean if a field has been set.

### GetDownloadStage

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadStage() string`

GetDownloadStage returns the DownloadStage field if non-nil, zero value otherwise.

### GetDownloadStageOk

`func (o *FirmwareUpgradeStatusAllOf) GetDownloadStageOk() (*string, bool)`

GetDownloadStageOk returns a tuple with the DownloadStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadStage

`func (o *FirmwareUpgradeStatusAllOf) SetDownloadStage(v string)`

SetDownloadStage sets DownloadStage field to given value.

### HasDownloadStage

`func (o *FirmwareUpgradeStatusAllOf) HasDownloadStage() bool`

HasDownloadStage returns a boolean if a field has been set.

### GetEpPowerStatus

`func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatus() string`

GetEpPowerStatus returns the EpPowerStatus field if non-nil, zero value otherwise.

### GetEpPowerStatusOk

`func (o *FirmwareUpgradeStatusAllOf) GetEpPowerStatusOk() (*string, bool)`

GetEpPowerStatusOk returns a tuple with the EpPowerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpPowerStatus

`func (o *FirmwareUpgradeStatusAllOf) SetEpPowerStatus(v string)`

SetEpPowerStatus sets EpPowerStatus field to given value.

### HasEpPowerStatus

`func (o *FirmwareUpgradeStatusAllOf) HasEpPowerStatus() bool`

HasEpPowerStatus returns a boolean if a field has been set.

### GetOverallError

`func (o *FirmwareUpgradeStatusAllOf) GetOverallError() string`

GetOverallError returns the OverallError field if non-nil, zero value otherwise.

### GetOverallErrorOk

`func (o *FirmwareUpgradeStatusAllOf) GetOverallErrorOk() (*string, bool)`

GetOverallErrorOk returns a tuple with the OverallError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallError

`func (o *FirmwareUpgradeStatusAllOf) SetOverallError(v string)`

SetOverallError sets OverallError field to given value.

### HasOverallError

`func (o *FirmwareUpgradeStatusAllOf) HasOverallError() bool`

HasOverallError returns a boolean if a field has been set.

### GetOverallPercentage

`func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentage() int64`

GetOverallPercentage returns the OverallPercentage field if non-nil, zero value otherwise.

### GetOverallPercentageOk

`func (o *FirmwareUpgradeStatusAllOf) GetOverallPercentageOk() (*int64, bool)`

GetOverallPercentageOk returns a tuple with the OverallPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallPercentage

`func (o *FirmwareUpgradeStatusAllOf) SetOverallPercentage(v int64)`

SetOverallPercentage sets OverallPercentage field to given value.

### HasOverallPercentage

`func (o *FirmwareUpgradeStatusAllOf) HasOverallPercentage() bool`

HasOverallPercentage returns a boolean if a field has been set.

### GetOverallstatus

`func (o *FirmwareUpgradeStatusAllOf) GetOverallstatus() string`

GetOverallstatus returns the Overallstatus field if non-nil, zero value otherwise.

### GetOverallstatusOk

`func (o *FirmwareUpgradeStatusAllOf) GetOverallstatusOk() (*string, bool)`

GetOverallstatusOk returns a tuple with the Overallstatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallstatus

`func (o *FirmwareUpgradeStatusAllOf) SetOverallstatus(v string)`

SetOverallstatus sets Overallstatus field to given value.

### HasOverallstatus

`func (o *FirmwareUpgradeStatusAllOf) HasOverallstatus() bool`

HasOverallstatus returns a boolean if a field has been set.

### GetPendingType

`func (o *FirmwareUpgradeStatusAllOf) GetPendingType() string`

GetPendingType returns the PendingType field if non-nil, zero value otherwise.

### GetPendingTypeOk

`func (o *FirmwareUpgradeStatusAllOf) GetPendingTypeOk() (*string, bool)`

GetPendingTypeOk returns a tuple with the PendingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingType

`func (o *FirmwareUpgradeStatusAllOf) SetPendingType(v string)`

SetPendingType sets PendingType field to given value.

### HasPendingType

`func (o *FirmwareUpgradeStatusAllOf) HasPendingType() bool`

HasPendingType returns a boolean if a field has been set.

### GetSha256checksum

`func (o *FirmwareUpgradeStatusAllOf) GetSha256checksum() string`

GetSha256checksum returns the Sha256checksum field if non-nil, zero value otherwise.

### GetSha256checksumOk

`func (o *FirmwareUpgradeStatusAllOf) GetSha256checksumOk() (*string, bool)`

GetSha256checksumOk returns a tuple with the Sha256checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256checksum

`func (o *FirmwareUpgradeStatusAllOf) SetSha256checksum(v string)`

SetSha256checksum sets Sha256checksum field to given value.

### HasSha256checksum

`func (o *FirmwareUpgradeStatusAllOf) HasSha256checksum() bool`

HasSha256checksum returns a boolean if a field has been set.

### GetUpgrade

`func (o *FirmwareUpgradeStatusAllOf) GetUpgrade() FirmwareUpgradeBaseRelationship`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *FirmwareUpgradeStatusAllOf) GetUpgradeOk() (*FirmwareUpgradeBaseRelationship, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *FirmwareUpgradeStatusAllOf) SetUpgrade(v FirmwareUpgradeBaseRelationship)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *FirmwareUpgradeStatusAllOf) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.

### GetWorkflow

`func (o *FirmwareUpgradeStatusAllOf) GetWorkflow() WorkflowWorkflowInfoRelationship`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *FirmwareUpgradeStatusAllOf) GetWorkflowOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *FirmwareUpgradeStatusAllOf) SetWorkflow(v WorkflowWorkflowInfoRelationship)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *FirmwareUpgradeStatusAllOf) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


