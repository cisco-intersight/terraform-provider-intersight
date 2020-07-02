# WorkflowErrorResponseHandlerAllOf

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

### NewWorkflowErrorResponseHandlerAllOf

`func NewWorkflowErrorResponseHandlerAllOf() *WorkflowErrorResponseHandlerAllOf`

NewWorkflowErrorResponseHandlerAllOf instantiates a new WorkflowErrorResponseHandlerAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowErrorResponseHandlerAllOfWithDefaults

`func NewWorkflowErrorResponseHandlerAllOfWithDefaults() *WorkflowErrorResponseHandlerAllOf`

NewWorkflowErrorResponseHandlerAllOfWithDefaults instantiates a new WorkflowErrorResponseHandlerAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkflowErrorResponseHandlerAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowErrorResponseHandlerAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowErrorResponseHandlerAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *WorkflowErrorResponseHandlerAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowErrorResponseHandlerAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowErrorResponseHandlerAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *WorkflowErrorResponseHandlerAllOf) GetParameters() []ContentParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetParametersOk() (*[]ContentParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WorkflowErrorResponseHandlerAllOf) SetParameters(v []ContentParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WorkflowErrorResponseHandlerAllOf) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPlatformType

`func (o *WorkflowErrorResponseHandlerAllOf) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *WorkflowErrorResponseHandlerAllOf) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *WorkflowErrorResponseHandlerAllOf) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetTypes

`func (o *WorkflowErrorResponseHandlerAllOf) GetTypes() []ContentComplexType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetTypesOk() (*[]ContentComplexType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *WorkflowErrorResponseHandlerAllOf) SetTypes(v []ContentComplexType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *WorkflowErrorResponseHandlerAllOf) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetCatalog

`func (o *WorkflowErrorResponseHandlerAllOf) GetCatalog() WorkflowCatalogRelationship`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *WorkflowErrorResponseHandlerAllOf) GetCatalogOk() (*WorkflowCatalogRelationship, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *WorkflowErrorResponseHandlerAllOf) SetCatalog(v WorkflowCatalogRelationship)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *WorkflowErrorResponseHandlerAllOf) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


