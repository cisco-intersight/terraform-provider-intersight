/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-06-17T02:04:50-07:00.
 *
 * API version: 1.0.9-1867
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
)

// HclHardwareCompatibilityProfileAllOf Definition of the list of properties defined in 'hcl.HardwareCompatibilityProfile', excluding properties defined in parent classes.
type HclHardwareCompatibilityProfileAllOf struct {
	// Url for the ISO with the drivers supported for the server.
	DriverIsoUrl *string `json:"DriverIsoUrl,omitempty"`
	// Error code indicating the compatibility status.
	ErrorCode *string `json:"ErrorCode,omitempty"`
	// Identifier of the hardware compatibility profile.
	Id *string `json:"Id,omitempty"`
	// Vendor of the Operating System running on the server.
	OsVendor *string `json:"OsVendor,omitempty"`
	// Version of the Operating System running on the server.
	OsVersion *string `json:"OsVersion,omitempty"`
	// Model of the processor present in the server.
	ProcessorModel *string       `json:"ProcessorModel,omitempty"`
	Products       *[]HclProduct `json:"Products,omitempty"`
	// Model of the server as returned by UCSM/CIMC XML API.
	ServerModel *string `json:"ServerModel,omitempty"`
	// Revision of the server model.
	ServerRevision *string `json:"ServerRevision,omitempty"`
	// Version of the UCS software.
	UcsVersion *string `json:"UcsVersion,omitempty"`
	// Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
	VersionType          *string `json:"VersionType,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HclHardwareCompatibilityProfileAllOf HclHardwareCompatibilityProfileAllOf

// NewHclHardwareCompatibilityProfileAllOf instantiates a new HclHardwareCompatibilityProfileAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHclHardwareCompatibilityProfileAllOf() *HclHardwareCompatibilityProfileAllOf {
	this := HclHardwareCompatibilityProfileAllOf{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// NewHclHardwareCompatibilityProfileAllOfWithDefaults instantiates a new HclHardwareCompatibilityProfileAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHclHardwareCompatibilityProfileAllOfWithDefaults() *HclHardwareCompatibilityProfileAllOf {
	this := HclHardwareCompatibilityProfileAllOf{}
	var errorCode string = "Success"
	this.ErrorCode = &errorCode
	var versionType string = "UCSM"
	this.VersionType = &versionType
	return &this
}

// GetDriverIsoUrl returns the DriverIsoUrl field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetDriverIsoUrl() string {
	if o == nil || o.DriverIsoUrl == nil {
		var ret string
		return ret
	}
	return *o.DriverIsoUrl
}

// GetDriverIsoUrlOk returns a tuple with the DriverIsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetDriverIsoUrlOk() (*string, bool) {
	if o == nil || o.DriverIsoUrl == nil {
		return nil, false
	}
	return o.DriverIsoUrl, true
}

// HasDriverIsoUrl returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasDriverIsoUrl() bool {
	if o != nil && o.DriverIsoUrl != nil {
		return true
	}

	return false
}

// SetDriverIsoUrl gets a reference to the given string and assigns it to the DriverIsoUrl field.
func (o *HclHardwareCompatibilityProfileAllOf) SetDriverIsoUrl(v string) {
	o.DriverIsoUrl = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *HclHardwareCompatibilityProfileAllOf) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HclHardwareCompatibilityProfileAllOf) SetId(v string) {
	o.Id = &v
}

// GetOsVendor returns the OsVendor field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetOsVendor() string {
	if o == nil || o.OsVendor == nil {
		var ret string
		return ret
	}
	return *o.OsVendor
}

// GetOsVendorOk returns a tuple with the OsVendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetOsVendorOk() (*string, bool) {
	if o == nil || o.OsVendor == nil {
		return nil, false
	}
	return o.OsVendor, true
}

// HasOsVendor returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasOsVendor() bool {
	if o != nil && o.OsVendor != nil {
		return true
	}

	return false
}

// SetOsVendor gets a reference to the given string and assigns it to the OsVendor field.
func (o *HclHardwareCompatibilityProfileAllOf) SetOsVendor(v string) {
	o.OsVendor = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetOsVersion() string {
	if o == nil || o.OsVersion == nil {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetOsVersionOk() (*string, bool) {
	if o == nil || o.OsVersion == nil {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasOsVersion() bool {
	if o != nil && o.OsVersion != nil {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *HclHardwareCompatibilityProfileAllOf) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProcessorModel returns the ProcessorModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetProcessorModel() string {
	if o == nil || o.ProcessorModel == nil {
		var ret string
		return ret
	}
	return *o.ProcessorModel
}

// GetProcessorModelOk returns a tuple with the ProcessorModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetProcessorModelOk() (*string, bool) {
	if o == nil || o.ProcessorModel == nil {
		return nil, false
	}
	return o.ProcessorModel, true
}

// HasProcessorModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasProcessorModel() bool {
	if o != nil && o.ProcessorModel != nil {
		return true
	}

	return false
}

// SetProcessorModel gets a reference to the given string and assigns it to the ProcessorModel field.
func (o *HclHardwareCompatibilityProfileAllOf) SetProcessorModel(v string) {
	o.ProcessorModel = &v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetProducts() []HclProduct {
	if o == nil || o.Products == nil {
		var ret []HclProduct
		return ret
	}
	return *o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetProductsOk() (*[]HclProduct, bool) {
	if o == nil || o.Products == nil {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasProducts() bool {
	if o != nil && o.Products != nil {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []HclProduct and assigns it to the Products field.
func (o *HclHardwareCompatibilityProfileAllOf) SetProducts(v []HclProduct) {
	o.Products = &v
}

// GetServerModel returns the ServerModel field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetServerModel() string {
	if o == nil || o.ServerModel == nil {
		var ret string
		return ret
	}
	return *o.ServerModel
}

// GetServerModelOk returns a tuple with the ServerModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetServerModelOk() (*string, bool) {
	if o == nil || o.ServerModel == nil {
		return nil, false
	}
	return o.ServerModel, true
}

// HasServerModel returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasServerModel() bool {
	if o != nil && o.ServerModel != nil {
		return true
	}

	return false
}

// SetServerModel gets a reference to the given string and assigns it to the ServerModel field.
func (o *HclHardwareCompatibilityProfileAllOf) SetServerModel(v string) {
	o.ServerModel = &v
}

// GetServerRevision returns the ServerRevision field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetServerRevision() string {
	if o == nil || o.ServerRevision == nil {
		var ret string
		return ret
	}
	return *o.ServerRevision
}

// GetServerRevisionOk returns a tuple with the ServerRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetServerRevisionOk() (*string, bool) {
	if o == nil || o.ServerRevision == nil {
		return nil, false
	}
	return o.ServerRevision, true
}

// HasServerRevision returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasServerRevision() bool {
	if o != nil && o.ServerRevision != nil {
		return true
	}

	return false
}

// SetServerRevision gets a reference to the given string and assigns it to the ServerRevision field.
func (o *HclHardwareCompatibilityProfileAllOf) SetServerRevision(v string) {
	o.ServerRevision = &v
}

// GetUcsVersion returns the UcsVersion field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetUcsVersion() string {
	if o == nil || o.UcsVersion == nil {
		var ret string
		return ret
	}
	return *o.UcsVersion
}

// GetUcsVersionOk returns a tuple with the UcsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetUcsVersionOk() (*string, bool) {
	if o == nil || o.UcsVersion == nil {
		return nil, false
	}
	return o.UcsVersion, true
}

// HasUcsVersion returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasUcsVersion() bool {
	if o != nil && o.UcsVersion != nil {
		return true
	}

	return false
}

// SetUcsVersion gets a reference to the given string and assigns it to the UcsVersion field.
func (o *HclHardwareCompatibilityProfileAllOf) SetUcsVersion(v string) {
	o.UcsVersion = &v
}

// GetVersionType returns the VersionType field value if set, zero value otherwise.
func (o *HclHardwareCompatibilityProfileAllOf) GetVersionType() string {
	if o == nil || o.VersionType == nil {
		var ret string
		return ret
	}
	return *o.VersionType
}

// GetVersionTypeOk returns a tuple with the VersionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HclHardwareCompatibilityProfileAllOf) GetVersionTypeOk() (*string, bool) {
	if o == nil || o.VersionType == nil {
		return nil, false
	}
	return o.VersionType, true
}

// HasVersionType returns a boolean if a field has been set.
func (o *HclHardwareCompatibilityProfileAllOf) HasVersionType() bool {
	if o != nil && o.VersionType != nil {
		return true
	}

	return false
}

// SetVersionType gets a reference to the given string and assigns it to the VersionType field.
func (o *HclHardwareCompatibilityProfileAllOf) SetVersionType(v string) {
	o.VersionType = &v
}

func (o HclHardwareCompatibilityProfileAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DriverIsoUrl != nil {
		toSerialize["DriverIsoUrl"] = o.DriverIsoUrl
	}
	if o.ErrorCode != nil {
		toSerialize["ErrorCode"] = o.ErrorCode
	}
	if o.Id != nil {
		toSerialize["Id"] = o.Id
	}
	if o.OsVendor != nil {
		toSerialize["OsVendor"] = o.OsVendor
	}
	if o.OsVersion != nil {
		toSerialize["OsVersion"] = o.OsVersion
	}
	if o.ProcessorModel != nil {
		toSerialize["ProcessorModel"] = o.ProcessorModel
	}
	if o.Products != nil {
		toSerialize["Products"] = o.Products
	}
	if o.ServerModel != nil {
		toSerialize["ServerModel"] = o.ServerModel
	}
	if o.ServerRevision != nil {
		toSerialize["ServerRevision"] = o.ServerRevision
	}
	if o.UcsVersion != nil {
		toSerialize["UcsVersion"] = o.UcsVersion
	}
	if o.VersionType != nil {
		toSerialize["VersionType"] = o.VersionType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HclHardwareCompatibilityProfileAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHclHardwareCompatibilityProfileAllOf := _HclHardwareCompatibilityProfileAllOf{}

	if err = json.Unmarshal(bytes, &varHclHardwareCompatibilityProfileAllOf); err == nil {
		*o = HclHardwareCompatibilityProfileAllOf(varHclHardwareCompatibilityProfileAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DriverIsoUrl")
		delete(additionalProperties, "ErrorCode")
		delete(additionalProperties, "Id")
		delete(additionalProperties, "OsVendor")
		delete(additionalProperties, "OsVersion")
		delete(additionalProperties, "ProcessorModel")
		delete(additionalProperties, "Products")
		delete(additionalProperties, "ServerModel")
		delete(additionalProperties, "ServerRevision")
		delete(additionalProperties, "UcsVersion")
		delete(additionalProperties, "VersionType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHclHardwareCompatibilityProfileAllOf struct {
	value *HclHardwareCompatibilityProfileAllOf
	isSet bool
}

func (v NullableHclHardwareCompatibilityProfileAllOf) Get() *HclHardwareCompatibilityProfileAllOf {
	return v.value
}

func (v *NullableHclHardwareCompatibilityProfileAllOf) Set(val *HclHardwareCompatibilityProfileAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHclHardwareCompatibilityProfileAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHclHardwareCompatibilityProfileAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHclHardwareCompatibilityProfileAllOf(val *HclHardwareCompatibilityProfileAllOf) *NullableHclHardwareCompatibilityProfileAllOf {
	return &NullableHclHardwareCompatibilityProfileAllOf{value: val, isSet: true}
}

func (v NullableHclHardwareCompatibilityProfileAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHclHardwareCompatibilityProfileAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
