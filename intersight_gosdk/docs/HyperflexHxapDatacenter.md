# HyperflexHxapDatacenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**HypervisorManager** | Pointer to [**HyperflexCiscoHypervisorManagerRelationship**](hyperflex.CiscoHypervisorManager.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapDatacenter

`func NewHyperflexHxapDatacenter() *HyperflexHxapDatacenter`

NewHyperflexHxapDatacenter instantiates a new HyperflexHxapDatacenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapDatacenterWithDefaults

`func NewHyperflexHxapDatacenterWithDefaults() *HyperflexHxapDatacenter`

NewHyperflexHxapDatacenterWithDefaults instantiates a new HyperflexHxapDatacenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *HyperflexHxapDatacenter) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *HyperflexHxapDatacenter) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *HyperflexHxapDatacenter) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *HyperflexHxapDatacenter) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetHypervisorManager

`func (o *HyperflexHxapDatacenter) GetHypervisorManager() HyperflexCiscoHypervisorManagerRelationship`

GetHypervisorManager returns the HypervisorManager field if non-nil, zero value otherwise.

### GetHypervisorManagerOk

`func (o *HyperflexHxapDatacenter) GetHypervisorManagerOk() (*HyperflexCiscoHypervisorManagerRelationship, bool)`

GetHypervisorManagerOk returns a tuple with the HypervisorManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisorManager

`func (o *HyperflexHxapDatacenter) SetHypervisorManager(v HyperflexCiscoHypervisorManagerRelationship)`

SetHypervisorManager sets HypervisorManager field to given value.

### HasHypervisorManager

`func (o *HyperflexHxapDatacenter) HasHypervisorManager() bool`

HasHypervisorManager returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


