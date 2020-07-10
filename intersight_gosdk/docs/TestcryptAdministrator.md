# TestcryptAdministrator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to [**TestcryptUser**](testcrypt.User.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptAdministrator

`func NewTestcryptAdministrator() *TestcryptAdministrator`

NewTestcryptAdministrator instantiates a new TestcryptAdministrator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptAdministratorWithDefaults

`func NewTestcryptAdministratorWithDefaults() *TestcryptAdministrator`

NewTestcryptAdministratorWithDefaults instantiates a new TestcryptAdministrator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *TestcryptAdministrator) GetAdmin() TestcryptUser`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *TestcryptAdministrator) GetAdminOk() (*TestcryptUser, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *TestcryptAdministrator) SetAdmin(v TestcryptUser)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *TestcryptAdministrator) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetAccount

`func (o *TestcryptAdministrator) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptAdministrator) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptAdministrator) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptAdministrator) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


