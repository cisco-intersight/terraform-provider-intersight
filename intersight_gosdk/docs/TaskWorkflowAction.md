# TaskWorkflowAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action for test workflow. | [optional] [default to "start"]
**InputParams** | Pointer to **string** | Json formatted string input parameters to workflow. | [optional] 
**IsDynamic** | Pointer to **bool** | When true, this workflow type will be triggered as a dynamic workflow, if not it will be treated as a static workflow. | [optional] 
**WaitOnDuplicate** | Pointer to **bool** | When true, the workflow will wait for previous one to complete before starting a new one. | [optional] 
**WorkflowContext** | Pointer to **string** | Json formatted string that has the contents of the workflow context used when starting a workflow. | [optional] 
**WorkflowFile** | Pointer to [**TaskFileDownloadInfo**](task.FileDownloadInfo.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTaskWorkflowAction

`func NewTaskWorkflowAction() *TaskWorkflowAction`

NewTaskWorkflowAction instantiates a new TaskWorkflowAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWorkflowActionWithDefaults

`func NewTaskWorkflowActionWithDefaults() *TaskWorkflowAction`

NewTaskWorkflowActionWithDefaults instantiates a new TaskWorkflowAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *TaskWorkflowAction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *TaskWorkflowAction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *TaskWorkflowAction) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *TaskWorkflowAction) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetInputParams

`func (o *TaskWorkflowAction) GetInputParams() string`

GetInputParams returns the InputParams field if non-nil, zero value otherwise.

### GetInputParamsOk

`func (o *TaskWorkflowAction) GetInputParamsOk() (*string, bool)`

GetInputParamsOk returns a tuple with the InputParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParams

`func (o *TaskWorkflowAction) SetInputParams(v string)`

SetInputParams sets InputParams field to given value.

### HasInputParams

`func (o *TaskWorkflowAction) HasInputParams() bool`

HasInputParams returns a boolean if a field has been set.

### GetIsDynamic

`func (o *TaskWorkflowAction) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *TaskWorkflowAction) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *TaskWorkflowAction) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.

### HasIsDynamic

`func (o *TaskWorkflowAction) HasIsDynamic() bool`

HasIsDynamic returns a boolean if a field has been set.

### GetWaitOnDuplicate

`func (o *TaskWorkflowAction) GetWaitOnDuplicate() bool`

GetWaitOnDuplicate returns the WaitOnDuplicate field if non-nil, zero value otherwise.

### GetWaitOnDuplicateOk

`func (o *TaskWorkflowAction) GetWaitOnDuplicateOk() (*bool, bool)`

GetWaitOnDuplicateOk returns a tuple with the WaitOnDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitOnDuplicate

`func (o *TaskWorkflowAction) SetWaitOnDuplicate(v bool)`

SetWaitOnDuplicate sets WaitOnDuplicate field to given value.

### HasWaitOnDuplicate

`func (o *TaskWorkflowAction) HasWaitOnDuplicate() bool`

HasWaitOnDuplicate returns a boolean if a field has been set.

### GetWorkflowContext

`func (o *TaskWorkflowAction) GetWorkflowContext() string`

GetWorkflowContext returns the WorkflowContext field if non-nil, zero value otherwise.

### GetWorkflowContextOk

`func (o *TaskWorkflowAction) GetWorkflowContextOk() (*string, bool)`

GetWorkflowContextOk returns a tuple with the WorkflowContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowContext

`func (o *TaskWorkflowAction) SetWorkflowContext(v string)`

SetWorkflowContext sets WorkflowContext field to given value.

### HasWorkflowContext

`func (o *TaskWorkflowAction) HasWorkflowContext() bool`

HasWorkflowContext returns a boolean if a field has been set.

### GetWorkflowFile

`func (o *TaskWorkflowAction) GetWorkflowFile() TaskFileDownloadInfo`

GetWorkflowFile returns the WorkflowFile field if non-nil, zero value otherwise.

### GetWorkflowFileOk

`func (o *TaskWorkflowAction) GetWorkflowFileOk() (*TaskFileDownloadInfo, bool)`

GetWorkflowFileOk returns a tuple with the WorkflowFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowFile

`func (o *TaskWorkflowAction) SetWorkflowFile(v TaskFileDownloadInfo)`

SetWorkflowFile sets WorkflowFile field to given value.

### HasWorkflowFile

`func (o *TaskWorkflowAction) HasWorkflowFile() bool`

HasWorkflowFile returns a boolean if a field has been set.

### GetAccount

`func (o *TaskWorkflowAction) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TaskWorkflowAction) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TaskWorkflowAction) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TaskWorkflowAction) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


