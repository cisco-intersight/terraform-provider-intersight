# RecoveryAbstractBackupConfigAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileNamePrefix** | Pointer to **string** | The file name for the backup image. This name is added as a prefix in the name for the backup image. A unique file name for the backup image is created along with a timestamp. For example: prefix-1572431305418. | [optional] 
**IsPasswordSet** | Pointer to **bool** | Indicates whether the value of the &#39;password&#39; property has been set. | [optional] [readonly] 
**LocationType** | Pointer to **string** | Specifies whether the backup will be stored locally or remotely. | [optional] [default to "Network Share"]
**Password** | Pointer to **string** | Backup server password. | [optional] 
**Path** | Pointer to **string** | The file system path where the backup images must be stored. Include the IP address/hostname of the network share location and the complete file system path. For example: 172.29.109.234/var/backups/. | [optional] 
**Protocol** | Pointer to **string** | Protocol for transferring the backup image to the network share location. | [optional] [default to "SCP"]
**RetentionCount** | Pointer to **int64** | Number of backup copies maintained on the local or remote server. When the created backup files exceed this number, the initial backup files are overwritten in a sequential manner. | [optional] 
**UserName** | Pointer to **string** | Backup server user name. | [optional] 

## Methods

### NewRecoveryAbstractBackupConfigAllOf

`func NewRecoveryAbstractBackupConfigAllOf() *RecoveryAbstractBackupConfigAllOf`

NewRecoveryAbstractBackupConfigAllOf instantiates a new RecoveryAbstractBackupConfigAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecoveryAbstractBackupConfigAllOfWithDefaults

`func NewRecoveryAbstractBackupConfigAllOfWithDefaults() *RecoveryAbstractBackupConfigAllOf`

NewRecoveryAbstractBackupConfigAllOfWithDefaults instantiates a new RecoveryAbstractBackupConfigAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileNamePrefix

`func (o *RecoveryAbstractBackupConfigAllOf) GetFileNamePrefix() string`

GetFileNamePrefix returns the FileNamePrefix field if non-nil, zero value otherwise.

### GetFileNamePrefixOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetFileNamePrefixOk() (*string, bool)`

GetFileNamePrefixOk returns a tuple with the FileNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePrefix

`func (o *RecoveryAbstractBackupConfigAllOf) SetFileNamePrefix(v string)`

SetFileNamePrefix sets FileNamePrefix field to given value.

### HasFileNamePrefix

`func (o *RecoveryAbstractBackupConfigAllOf) HasFileNamePrefix() bool`

HasFileNamePrefix returns a boolean if a field has been set.

### GetIsPasswordSet

`func (o *RecoveryAbstractBackupConfigAllOf) GetIsPasswordSet() bool`

GetIsPasswordSet returns the IsPasswordSet field if non-nil, zero value otherwise.

### GetIsPasswordSetOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetIsPasswordSetOk() (*bool, bool)`

GetIsPasswordSetOk returns a tuple with the IsPasswordSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPasswordSet

`func (o *RecoveryAbstractBackupConfigAllOf) SetIsPasswordSet(v bool)`

SetIsPasswordSet sets IsPasswordSet field to given value.

### HasIsPasswordSet

`func (o *RecoveryAbstractBackupConfigAllOf) HasIsPasswordSet() bool`

HasIsPasswordSet returns a boolean if a field has been set.

### GetLocationType

`func (o *RecoveryAbstractBackupConfigAllOf) GetLocationType() string`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetLocationTypeOk() (*string, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *RecoveryAbstractBackupConfigAllOf) SetLocationType(v string)`

SetLocationType sets LocationType field to given value.

### HasLocationType

`func (o *RecoveryAbstractBackupConfigAllOf) HasLocationType() bool`

HasLocationType returns a boolean if a field has been set.

### GetPassword

`func (o *RecoveryAbstractBackupConfigAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RecoveryAbstractBackupConfigAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RecoveryAbstractBackupConfigAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPath

`func (o *RecoveryAbstractBackupConfigAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *RecoveryAbstractBackupConfigAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *RecoveryAbstractBackupConfigAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProtocol

`func (o *RecoveryAbstractBackupConfigAllOf) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *RecoveryAbstractBackupConfigAllOf) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *RecoveryAbstractBackupConfigAllOf) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRetentionCount

`func (o *RecoveryAbstractBackupConfigAllOf) GetRetentionCount() int64`

GetRetentionCount returns the RetentionCount field if non-nil, zero value otherwise.

### GetRetentionCountOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetRetentionCountOk() (*int64, bool)`

GetRetentionCountOk returns a tuple with the RetentionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCount

`func (o *RecoveryAbstractBackupConfigAllOf) SetRetentionCount(v int64)`

SetRetentionCount sets RetentionCount field to given value.

### HasRetentionCount

`func (o *RecoveryAbstractBackupConfigAllOf) HasRetentionCount() bool`

HasRetentionCount returns a boolean if a field has been set.

### GetUserName

`func (o *RecoveryAbstractBackupConfigAllOf) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *RecoveryAbstractBackupConfigAllOf) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *RecoveryAbstractBackupConfigAllOf) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *RecoveryAbstractBackupConfigAllOf) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


