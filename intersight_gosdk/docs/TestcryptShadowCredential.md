# TestcryptShadowCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Password associated with username. | [optional] [readonly] 
**Username** | Pointer to **string** | The username of user, whose password marked as &#39;secure&#39;. | [optional] [readonly] 
**Account** | Pointer to [**IamAccountRelationship**](iam.Account.Relationship.md) |  | [optional] 

## Methods

### NewTestcryptShadowCredential

`func NewTestcryptShadowCredential() *TestcryptShadowCredential`

NewTestcryptShadowCredential instantiates a new TestcryptShadowCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestcryptShadowCredentialWithDefaults

`func NewTestcryptShadowCredentialWithDefaults() *TestcryptShadowCredential`

NewTestcryptShadowCredentialWithDefaults instantiates a new TestcryptShadowCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *TestcryptShadowCredential) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TestcryptShadowCredential) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TestcryptShadowCredential) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TestcryptShadowCredential) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUsername

`func (o *TestcryptShadowCredential) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TestcryptShadowCredential) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TestcryptShadowCredential) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TestcryptShadowCredential) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccount

`func (o *TestcryptShadowCredential) GetAccount() IamAccountRelationship`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *TestcryptShadowCredential) GetAccountOk() (*IamAccountRelationship, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *TestcryptShadowCredential) SetAccount(v IamAccountRelationship)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *TestcryptShadowCredential) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


