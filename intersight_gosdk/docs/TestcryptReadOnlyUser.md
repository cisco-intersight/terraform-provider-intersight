# TestcryptReadOnlyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rouser** | Pointer to [**[]TestcryptUser**](testcrypt.User.md) |  | [optional] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptReadOnlyUser

`func NewTestcryptReadOnlyUser() *TestcryptReadOnlyUser`

NewTestcryptReadOnlyUser instantiates a new TestcryptReadOnlyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptReadOnlyUserWithDefaults

`func NewTestcryptReadOnlyUserWithDefaults() *TestcryptReadOnlyUser`

NewTestcryptReadOnlyUserWithDefaults instantiates a new TestcryptReadOnlyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouser

`func (o *TestcryptReadOnlyUser) GetRouser() []TestcryptUser`

GetRouser returns the Rouser field if non-nil, zero value otherwise.

### GetRouserOk

`func (o *TestcryptReadOnlyUser) GetRouserOk() (*[]TestcryptUser, bool)`

GetRouserOk returns a tuple with the Rouser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouser

`func (o *TestcryptReadOnlyUser) SetRouser(v []TestcryptUser)`

SetRouser sets Rouser field to given value.

### HasRouser

`func (o *TestcryptReadOnlyUser) HasRouser() bool`

HasRouser returns a boolean if a field has been set.

### GetAccount

`func (o *TestcryptReadOnlyUser) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptReadOnlyUser) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptReadOnlyUser) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptReadOnlyUser) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


