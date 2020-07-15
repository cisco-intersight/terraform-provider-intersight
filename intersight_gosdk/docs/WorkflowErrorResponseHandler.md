# WorkflowErrorResponseHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A detailed description about the error response handler. | [optional] 
**Name** | Pointer to **string** | Name for the error response handler. | [optional] 
**Parameters** | Pointer to [**[]ContentParameter**](content.Parameter.md) |  | [optional] 
**PlatformType** | Pointer to **string** | The platform type for which the error response handler is defined. | [optional] [default to ""]
**Types** | Pointer to [**[]ContentComplexType**](content.ComplexType.md) |  | [optional] 
**Catalog** | Pointer to [**WorkflowCatalogRelationship**](workflow.Catalog.Relationship.md) |  | [optional] 

## Methods

### NewWorkflowErrorResponseHandler

`func NewWorkflowErrorResponseHandler() *WorkflowErrorResponseHandler`

NewWorkflowErrorResponseHandler instantiates a new WorkflowErrorResponseHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowErrorResponseHandlerWithDefaults

`func NewWorkflowErrorResponseHandlerWithDefaults() *WorkflowErrorResponseHandler`

NewWorkflowErrorResponseHandlerWithDefaults instantiates a new WorkflowErrorResponseHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowErrorResponseHandler) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowErrorResponseHandler) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowErrorResponseHandler) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowErrorResponseHandler) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowErrorResponseHandler) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowErrorResponseHandler) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowErrorResponseHandler) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowErrorResponseHandler) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *WorkflowErrorResponseHandler) GetParameters() []ContentParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WorkflowErrorResponseHandler) GetParametersOk() (*[]ContentParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WorkflowErrorResponseHandler) SetParameters(v []ContentParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WorkflowErrorResponseHandler) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPlatformType

`func (o *WorkflowErrorResponseHandler) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *WorkflowErrorResponseHandler) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *WorkflowErrorResponseHandler) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *WorkflowErrorResponseHandler) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetTypes

`func (o *WorkflowErrorResponseHandler) GetTypes() []ContentComplexType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WorkflowErrorResponseHandler) GetTypesOk() (*[]ContentComplexType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WorkflowErrorResponseHandler) SetTypes(v []ContentComplexType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WorkflowErrorResponseHandler) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowErrorResponseHandler) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowErrorResponseHandler) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowErrorResponseHandler) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowErrorResponseHandler) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


