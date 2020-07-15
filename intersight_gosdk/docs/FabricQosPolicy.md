# FabricQosPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classes** | Pointer to [**[]FabricQosClass**](fabric.QosClass.md) |  | [optional] 
**Organization** | Pointer to [**OrganizationOrganizationRelationship**](organization.Organization.Relationship.md) |  | [optional] 
**Profiles** | Pointer to [**[]FabricSwitchProfileRelationship**](fabric.SwitchProfile.Relationship.md) | An array of relationships to fabricSwitchProfile resources. | [optional] 

## Methods

### NewFabricQosPolicy

`func NewFabricQosPolicy() *FabricQosPolicy`

NewFabricQosPolicy instantiates a new FabricQosPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFabricQosPolicyWithDefaults

`func NewFabricQosPolicyWithDefaults() *FabricQosPolicy`

NewFabricQosPolicyWithDefaults instantiates a new FabricQosPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClasses

`func (o *FabricQosPolicy) GetClasses() []FabricQosClass`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *FabricQosPolicy) GetClassesOk() (*[]FabricQosClass, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *FabricQosPolicy) SetClasses(v []FabricQosClass)`

SetClasses sets Classes field to given value.

### HasClasses

`func (o *FabricQosPolicy) HasClasses() bool`

HasClasses returns a boolean if a field has been set.

### GetOrganization

`func (o *FabricQosPolicy) GetOrganization() OrganizationOrganizationRelationship`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *FabricQosPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *FabricQosPolicy) SetOrganization(v OrganizationOrganizationRelationship)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *FabricQosPolicy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetProfiles

`func (o *FabricQosPolicy) GetProfiles() []FabricSwitchProfileRelationship`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *FabricQosPolicy) GetProfilesOk() (*[]FabricSwitchProfileRelationship, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *FabricQosPolicy) SetProfiles(v []FabricSwitchProfileRelationship)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *FabricQosPolicy) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### SetProfilesNil

`func (o *FabricQosPolicy) SetProfilesNil(b bool)`

 SetProfilesNil sets the value for Profiles to be an explicit nil

### UnsetProfiles
`func (o *FabricQosPolicy) UnsetProfiles()`

UnsetProfiles ensures that no value is present for Profiles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


