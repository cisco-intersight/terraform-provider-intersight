# InventoryEpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionStatus** | Pointer to **string** | Connections status of UEM endpoint. | [optional] [readonly] [default to "Unknown"]
**EpType** | Pointer to **string** | Type of UEM endpoint (BMC, CMC or VIC). | [optional] [readonly] 
**Ip** | Pointer to **string** | The IP address of the UEM endpoint. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address of the UEM endpoint. | [optional] [readonly] 
**MemberIdentity** | Pointer to **string** | The member identity of the UEM connection being inventoried. | [optional] [readonly] 
**Model** | Pointer to **string** | The model of the UEM endpoint. | [optional] [readonly] 
**Serial** | Pointer to **string** | The serial number of the UEM endpoint. | [optional] [readonly] 
**ServerRegistrationId** | Pointer to **string** | The device id of the server which this EP is a part of. | [optional] [readonly] 
**Vendor** | Pointer to **string** | The vendor of the UEM endpoint. | [optional] [readonly] 

## Methods

### NewInventoryEpInfo

`func NewInventoryEpInfo() *InventoryEpInfo`

NewInventoryEpInfo instantiates a new InventoryEpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryEpInfoWithDefaults

`func NewInventoryEpInfoWithDefaults() *InventoryEpInfo`

NewInventoryEpInfoWithDefaults instantiates a new InventoryEpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionStatus

`func (o *InventoryEpInfo) GetConnectionStatus() string`

GetConnectionStatus returns the ConnectionStatus field if non-nil, zero value otherwise.

### GetConnectionStatusOk

`func (o *InventoryEpInfo) GetConnectionStatusOk() (*string, bool)`

GetConnectionStatusOk returns a tuple with the ConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStatus

`func (o *InventoryEpInfo) SetConnectionStatus(v string)`

SetConnectionStatus sets ConnectionStatus field to given value.

### HasConnectionStatus

`func (o *InventoryEpInfo) HasConnectionStatus() bool`

HasConnectionStatus returns a boolean if a field has been set.

### GetEpType

`func (o *InventoryEpInfo) GetEpType() string`

GetEpType returns the EpType field if non-nil, zero value otherwise.

### GetEpTypeOk

`func (o *InventoryEpInfo) GetEpTypeOk() (*string, bool)`

GetEpTypeOk returns a tuple with the EpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpType

`func (o *InventoryEpInfo) SetEpType(v string)`

SetEpType sets EpType field to given value.

### HasEpType

`func (o *InventoryEpInfo) HasEpType() bool`

HasEpType returns a boolean if a field has been set.

### GetIp

`func (o *InventoryEpInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *InventoryEpInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *InventoryEpInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *InventoryEpInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMacAddress

`func (o *InventoryEpInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *InventoryEpInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *InventoryEpInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *InventoryEpInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMemberIdentity

`func (o *InventoryEpInfo) GetMemberIdentity() string`

GetMemberIdentity returns the MemberIdentity field if non-nil, zero value otherwise.

### GetMemberIdentityOk

`func (o *InventoryEpInfo) GetMemberIdentityOk() (*string, bool)`

GetMemberIdentityOk returns a tuple with the MemberIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberIdentity

`func (o *InventoryEpInfo) SetMemberIdentity(v string)`

SetMemberIdentity sets MemberIdentity field to given value.

### HasMemberIdentity

`func (o *InventoryEpInfo) HasMemberIdentity() bool`

HasMemberIdentity returns a boolean if a field has been set.

### GetModel

`func (o *InventoryEpInfo) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InventoryEpInfo) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InventoryEpInfo) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *InventoryEpInfo) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSerial

`func (o *InventoryEpInfo) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *InventoryEpInfo) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *InventoryEpInfo) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *InventoryEpInfo) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetServerRegistrationId

`func (o *InventoryEpInfo) GetServerRegistrationId() string`

GetServerRegistrationId returns the ServerRegistrationId field if non-nil, zero value otherwise.

### GetServerRegistrationIdOk

`func (o *InventoryEpInfo) GetServerRegistrationIdOk() (*string, bool)`

GetServerRegistrationIdOk returns a tuple with the ServerRegistrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistrationId

`func (o *InventoryEpInfo) SetServerRegistrationId(v string)`

SetServerRegistrationId sets ServerRegistrationId field to given value.

### HasServerRegistrationId

`func (o *InventoryEpInfo) HasServerRegistrationId() bool`

HasServerRegistrationId returns a boolean if a field has been set.

### GetVendor

`func (o *InventoryEpInfo) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *InventoryEpInfo) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *InventoryEpInfo) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *InventoryEpInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


