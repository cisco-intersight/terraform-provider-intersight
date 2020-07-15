# FabricQosPolicyListAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The total number of &#39;fabric.QosPolicy&#39; resources matching the request, accross all pages. The &#39;Count&#39; attribute is included when the HTTP GET request includes the &#39;$inlinecount&#39; parameter. | [optional] 
**Results** | Pointer to [**[]FabricQosPolicy**](fabric.QosPolicy.md) | The array of &#39;fabric.QosPolicy&#39; resources matching the request. | [optional] 

## Methods

### NewFabricQosPolicyListAllOf

`func NewFabricQosPolicyListAllOf() *FabricQosPolicyListAllOf`

NewFabricQosPolicyListAllOf instantiates a new FabricQosPolicyListAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricQosPolicyListAllOfWithDefaults

`func NewFabricQosPolicyListAllOfWithDefaults() *FabricQosPolicyListAllOf`

NewFabricQosPolicyListAllOfWithDefaults instantiates a new FabricQosPolicyListAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *FabricQosPolicyListAllOf) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FabricQosPolicyListAllOf) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FabricQosPolicyListAllOf) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *FabricQosPolicyListAllOf) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetResults

`func (o *FabricQosPolicyListAllOf) GetResults() []FabricQosPolicy`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *FabricQosPolicyListAllOf) GetResultsOk() (*[]FabricQosPolicy, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *FabricQosPolicyListAllOf) SetResults(v []FabricQosPolicy)`

SetResults sets Results field to given value.

### HasResults

`func (o *FabricQosPolicyListAllOf) HasResults() bool`

HasResults returns a boolean if a field has been set.

### SetResultsNil

`func (o *FabricQosPolicyListAllOf) SetResultsNil(b bool)`

 SetResultsNil sets the value for Results to be an explicit nil

### UnsetResults
`func (o *FabricQosPolicyListAllOf) UnsetResults()`

UnsetResults ensures that no value is present for Results, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


