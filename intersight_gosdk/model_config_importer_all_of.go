/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-19T15:15:17Z.
 *
 * API version: 1.0.9-1903
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// ConfigImporterAllOf Definition of the list of properties defined in 'config.Importer', excluding properties defined in parent classes.
type ConfigImporterAllOf struct {
	// The path to the archive in Intersight storage that has all the MOs to be imported.
	ImportPath *string `json:"ImportPath,omitempty"`
	// The source of the archive in Intersight storage that has all the MOs to be imported.
	ImportSource *string `json:"ImportSource,omitempty"`
	// An identifier for the importer instance.
	Name *string `json:"Name,omitempty"`
	// Specifies whether integrity checks must be skipped.
	SkipIntegrityChecks *bool `json:"SkipIntegrityChecks,omitempty"`
	// Status of the import operation.
	Status *string `json:"Status,omitempty"`
	// Status message associated with failures or progress indication.
	StatusMessage *string `json:"StatusMessage,omitempty"`
	// An array of relationships to configImportedItem resources.
	ImportedItems []ConfigImportedItemRelationship      `json:"ImportedItems,omitempty"`
	Organization  *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
}

// NewConfigImporterAllOf instantiates a new ConfigImporterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigImporterAllOf() *ConfigImporterAllOf {
	this := ConfigImporterAllOf{}
	var importSource string = "ImageRepo"
	this.ImportSource = &importSource
	var status string = ""
	this.Status = &status
	return &this
}

// NewConfigImporterAllOfWithDefaults instantiates a new ConfigImporterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigImporterAllOfWithDefaults() *ConfigImporterAllOf {
	this := ConfigImporterAllOf{}
	var importSource string = "ImageRepo"
	this.ImportSource = &importSource
	var status string = ""
	this.Status = &status
	return &this
}

// GetImportPath returns the ImportPath field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetImportPath() string {
	if o == nil || o.ImportPath == nil {
		var ret string
		return ret
	}
	return *o.ImportPath
}

// GetImportPathOk returns a tuple with the ImportPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetImportPathOk() (*string, bool) {
	if o == nil || o.ImportPath == nil {
		return nil, false
	}
	return o.ImportPath, true
}

// HasImportPath returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasImportPath() bool {
	if o != nil && o.ImportPath != nil {
		return true
	}

	return false
}

// SetImportPath gets a reference to the given string and assigns it to the ImportPath field.
func (o *ConfigImporterAllOf) SetImportPath(v string) {
	o.ImportPath = &v
}

// GetImportSource returns the ImportSource field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetImportSource() string {
	if o == nil || o.ImportSource == nil {
		var ret string
		return ret
	}
	return *o.ImportSource
}

// GetImportSourceOk returns a tuple with the ImportSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetImportSourceOk() (*string, bool) {
	if o == nil || o.ImportSource == nil {
		return nil, false
	}
	return o.ImportSource, true
}

// HasImportSource returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasImportSource() bool {
	if o != nil && o.ImportSource != nil {
		return true
	}

	return false
}

// SetImportSource gets a reference to the given string and assigns it to the ImportSource field.
func (o *ConfigImporterAllOf) SetImportSource(v string) {
	o.ImportSource = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConfigImporterAllOf) SetName(v string) {
	o.Name = &v
}

// GetSkipIntegrityChecks returns the SkipIntegrityChecks field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetSkipIntegrityChecks() bool {
	if o == nil || o.SkipIntegrityChecks == nil {
		var ret bool
		return ret
	}
	return *o.SkipIntegrityChecks
}

// GetSkipIntegrityChecksOk returns a tuple with the SkipIntegrityChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetSkipIntegrityChecksOk() (*bool, bool) {
	if o == nil || o.SkipIntegrityChecks == nil {
		return nil, false
	}
	return o.SkipIntegrityChecks, true
}

// HasSkipIntegrityChecks returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasSkipIntegrityChecks() bool {
	if o != nil && o.SkipIntegrityChecks != nil {
		return true
	}

	return false
}

// SetSkipIntegrityChecks gets a reference to the given bool and assigns it to the SkipIntegrityChecks field.
func (o *ConfigImporterAllOf) SetSkipIntegrityChecks(v bool) {
	o.SkipIntegrityChecks = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConfigImporterAllOf) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ConfigImporterAllOf) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetImportedItems returns the ImportedItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConfigImporterAllOf) GetImportedItems() []ConfigImportedItemRelationship {
	if o == nil {
		var ret []ConfigImportedItemRelationship
		return ret
	}
	return o.ImportedItems
}

// GetImportedItemsOk returns a tuple with the ImportedItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConfigImporterAllOf) GetImportedItemsOk() (*[]ConfigImportedItemRelationship, bool) {
	if o == nil || o.ImportedItems == nil {
		return nil, false
	}
	return &o.ImportedItems, true
}

// HasImportedItems returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasImportedItems() bool {
	if o != nil && o.ImportedItems != nil {
		return true
	}

	return false
}

// SetImportedItems gets a reference to the given []ConfigImportedItemRelationship and assigns it to the ImportedItems field.
func (o *ConfigImporterAllOf) SetImportedItems(v []ConfigImportedItemRelationship) {
	o.ImportedItems = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ConfigImporterAllOf) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigImporterAllOf) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ConfigImporterAllOf) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ConfigImporterAllOf) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ConfigImporterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportPath != nil {
		toSerialize["ImportPath"] = o.ImportPath
	}
	if o.ImportSource != nil {
		toSerialize["ImportSource"] = o.ImportSource
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.SkipIntegrityChecks != nil {
		toSerialize["SkipIntegrityChecks"] = o.SkipIntegrityChecks
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["StatusMessage"] = o.StatusMessage
	}
	if o.ImportedItems != nil {
		toSerialize["ImportedItems"] = o.ImportedItems
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	return json.Marshal(toSerialize)
}

type NullableConfigImporterAllOf struct {
	value *ConfigImporterAllOf
	isSet bool
}

func (v NullableConfigImporterAllOf) Get() *ConfigImporterAllOf {
	return v.value
}

func (v *NullableConfigImporterAllOf) Set(val *ConfigImporterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigImporterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigImporterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigImporterAllOf(val *ConfigImporterAllOf) *NullableConfigImporterAllOf {
	return &NullableConfigImporterAllOf{value: val, isSet: true}
}

func (v NullableConfigImporterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigImporterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}