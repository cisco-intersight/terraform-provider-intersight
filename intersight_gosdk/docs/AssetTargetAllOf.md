# AssetTargetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]AssetConnection**](asset.Connection.md) |  | [optional] 
**Services** | Pointer to [**[]AssetService**](asset.Service.md) |  | [optional] 
**Status** | Pointer to **string** | Status indicates if Intersight can establish a connection and authenticate with the managed target. Status does not include information about the functional health of the target. | [optional] [readonly] [default to ""]
**StatusErrorReason** | Pointer to **string** | StatusErrorReason provides additional context for the Status. | [optional] [readonly] 
**TargetType** | Pointer to **string** | The type of the managed target. For example a UCS Server or Vmware Vcenter target. | [optional] [default to ""]
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 
**Assist** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 

## Methods

### NewAssetTargetAllOf

`func NewAssetTargetAllOf() *AssetTargetAllOf`

NewAssetTargetAllOf instantiates a new AssetTargetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetTargetAllOfWithDefaults

`func NewAssetTargetAllOfWithDefaults() *AssetTargetAllOf`

NewAssetTargetAllOfWithDefaults instantiates a new AssetTargetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *AssetTargetAllOf) GetConnections() []AssetConnection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *AssetTargetAllOf) GetConnectionsOk() (*[]AssetConnection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *AssetTargetAllOf) SetConnections(v []AssetConnection)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *AssetTargetAllOf) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetServices

`func (o *AssetTargetAllOf) GetServices() []AssetService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AssetTargetAllOf) GetServicesOk() (*[]AssetService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AssetTargetAllOf) SetServices(v []AssetService)`

SetServices sets Services field to given value.

### HasServices

`func (o *AssetTargetAllOf) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetStatus

`func (o *AssetTargetAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetTargetAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetTargetAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AssetTargetAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusErrorReason

`func (o *AssetTargetAllOf) GetStatusErrorReason() string`

GetStatusErrorReason returns the StatusErrorReason field if non-nil, zero value otherwise.

### GetStatusErrorReasonOk

`func (o *AssetTargetAllOf) GetStatusErrorReasonOk() (*string, bool)`

GetStatusErrorReasonOk returns a tuple with the StatusErrorReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusErrorReason

`func (o *AssetTargetAllOf) SetStatusErrorReason(v string)`

SetStatusErrorReason sets StatusErrorReason field to given value.

### HasStatusErrorReason

`func (o *AssetTargetAllOf) HasStatusErrorReason() bool`

HasStatusErrorReason returns a boolean if a field has been set.

### GetTargetType

`func (o *AssetTargetAllOf) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *AssetTargetAllOf) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *AssetTargetAllOf) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *AssetTargetAllOf) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### GetAccount

`func (o *AssetTargetAllOf) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AssetTargetAllOf) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AssetTargetAllOf) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AssetTargetAllOf) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAssist

`func (o *AssetTargetAllOf) GetAssist() AssetDeviceRegistrationRelationship`

GetAssist returns the Assist field if non-nil, zero value otherwise.

### GetAssistOk

`func (o *AssetTargetAllOf) GetAssistOk() (*AssetDeviceRegistrationRelationship, bool)`

GetAssistOk returns a tuple with the Assist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssist

`func (o *AssetTargetAllOf) SetAssist(v AssetDeviceRegistrationRelationship)`

SetAssist sets Assist field to given value.

### HasAssist

`func (o *AssetTargetAllOf) HasAssist() bool`

HasAssist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


