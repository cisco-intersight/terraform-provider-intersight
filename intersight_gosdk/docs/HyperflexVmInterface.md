# HyperflexVmInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bridge** | Pointer to **string** | Virtual machine brige name of interface. | [optional] [readonly] 
**IpAddress** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **string** | Virtual machine interface mac address. | [optional] [readonly] 
**Model** | Pointer to **string** | Virtual machine interface model. | [optional] [readonly] 

## Methods

### NewHyperflexVmInterface

`func NewHyperflexVmInterface() *HyperflexVmInterface`

NewHyperflexVmInterface instantiates a new HyperflexVmInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexVmInterfaceWithDefaults

`func NewHyperflexVmInterfaceWithDefaults() *HyperflexVmInterface`

NewHyperflexVmInterfaceWithDefaults instantiates a new HyperflexVmInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBridge

`func (o *HyperflexVmInterface) GetBridge() string`

GetBridge returns the Bridge field if non-nil, zero value otherwise.

### GetBridgeOk

`func (o *HyperflexVmInterface) GetBridgeOk() (*string, bool)`

GetBridgeOk returns a tuple with the Bridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridge

`func (o *HyperflexVmInterface) SetBridge(v string)`

SetBridge sets Bridge field to given value.

### HasBridge

`func (o *HyperflexVmInterface) HasBridge() bool`

HasBridge returns a boolean if a field has been set.

### GetIpAddress

`func (o *HyperflexVmInterface) GetIpAddress() []string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *HyperflexVmInterface) GetIpAddressOk() (*[]string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *HyperflexVmInterface) SetIpAddress(v []string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *HyperflexVmInterface) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetMacAddress

`func (o *HyperflexVmInterface) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *HyperflexVmInterface) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *HyperflexVmInterface) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *HyperflexVmInterface) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetModel

`func (o *HyperflexVmInterface) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *HyperflexVmInterface) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *HyperflexVmInterface) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *HyperflexVmInterface) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


