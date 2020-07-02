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

// SoftwareHclMeta A JSON file wth HCL metadata uploaded for consumption by the HCL service.
type SoftwareHclMeta struct {
	FirmwareBaseDistributable
	// The type of content that the Json file holds (Incremental or full dump).
	ContentType *string                                `json:"ContentType,omitempty"`
	Catalog     *SoftwarerepositoryCatalogRelationship `json:"Catalog,omitempty"`
}

// NewSoftwareHclMeta instantiates a new SoftwareHclMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoftwareHclMeta() *SoftwareHclMeta {
	this := SoftwareHclMeta{}
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// NewSoftwareHclMetaWithDefaults instantiates a new SoftwareHclMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoftwareHclMetaWithDefaults() *SoftwareHclMeta {
	this := SoftwareHclMeta{}
	var contentType string = "Full"
	this.ContentType = &contentType
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *SoftwareHclMeta) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *SoftwareHclMeta) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *SoftwareHclMeta) SetContentType(v string) {
	o.ContentType = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *SoftwareHclMeta) GetCatalog() SoftwarerepositoryCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret SoftwarerepositoryCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoftwareHclMeta) GetCatalogOk() (*SoftwarerepositoryCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *SoftwareHclMeta) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given SoftwarerepositoryCatalogRelationship and assigns it to the Catalog field.
func (o *SoftwareHclMeta) SetCatalog(v SoftwarerepositoryCatalogRelationship) {
	o.Catalog = &v
}

func (o SoftwareHclMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFirmwareBaseDistributable, errFirmwareBaseDistributable := json.Marshal(o.FirmwareBaseDistributable)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	errFirmwareBaseDistributable = json.Unmarshal([]byte(serializedFirmwareBaseDistributable), &toSerialize)
	if errFirmwareBaseDistributable != nil {
		return []byte{}, errFirmwareBaseDistributable
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}
	return json.Marshal(toSerialize)
}

type NullableSoftwareHclMeta struct {
	value *SoftwareHclMeta
	isSet bool
}

func (v NullableSoftwareHclMeta) Get() *SoftwareHclMeta {
	return v.value
}

func (v *NullableSoftwareHclMeta) Set(val *SoftwareHclMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableSoftwareHclMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableSoftwareHclMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoftwareHclMeta(val *SoftwareHclMeta) *NullableSoftwareHclMeta {
	return &NullableSoftwareHclMeta{value: val, isSet: true}
}

func (v NullableSoftwareHclMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoftwareHclMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}