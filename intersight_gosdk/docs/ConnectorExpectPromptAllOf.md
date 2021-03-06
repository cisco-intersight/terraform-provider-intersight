# ConnectorExpectPromptAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expect** | Pointer to **string** | The regex of the expect prompt of the interactive command. | [optional] 
**ExpectTimeout** | Pointer to **int64** | The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used. | [optional] 
**Send** | Pointer to **string** | The answer string to the expect prompt. | [optional] 

## Methods

### NewConnectorExpectPromptAllOf

`func NewConnectorExpectPromptAllOf() *ConnectorExpectPromptAllOf`

NewConnectorExpectPromptAllOf instantiates a new ConnectorExpectPromptAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorExpectPromptAllOfWithDefaults

`func NewConnectorExpectPromptAllOfWithDefaults() *ConnectorExpectPromptAllOf`

NewConnectorExpectPromptAllOfWithDefaults instantiates a new ConnectorExpectPromptAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpect

`func (o *ConnectorExpectPromptAllOf) GetExpect() string`

GetExpect returns the Expect field if non-nil, zero value otherwise.

### GetExpectOk

`func (o *ConnectorExpectPromptAllOf) GetExpectOk() (*string, bool)`

GetExpectOk returns a tuple with the Expect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpect

`func (o *ConnectorExpectPromptAllOf) SetExpect(v string)`

SetExpect sets Expect field to given value.

### HasExpect

`func (o *ConnectorExpectPromptAllOf) HasExpect() bool`

HasExpect returns a boolean if a field has been set.

### GetExpectTimeout

`func (o *ConnectorExpectPromptAllOf) GetExpectTimeout() int64`

GetExpectTimeout returns the ExpectTimeout field if non-nil, zero value otherwise.

### GetExpectTimeoutOk

`func (o *ConnectorExpectPromptAllOf) GetExpectTimeoutOk() (*int64, bool)`

GetExpectTimeoutOk returns a tuple with the ExpectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectTimeout

`func (o *ConnectorExpectPromptAllOf) SetExpectTimeout(v int64)`

SetExpectTimeout sets ExpectTimeout field to given value.

### HasExpectTimeout

`func (o *ConnectorExpectPromptAllOf) HasExpectTimeout() bool`

HasExpectTimeout returns a boolean if a field has been set.

### GetSend

`func (o *ConnectorExpectPromptAllOf) GetSend() string`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *ConnectorExpectPromptAllOf) GetSendOk() (*string, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *ConnectorExpectPromptAllOf) SetSend(v string)`

SetSend sets Send field to given value.

### HasSend

`func (o *ConnectorExpectPromptAllOf) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


