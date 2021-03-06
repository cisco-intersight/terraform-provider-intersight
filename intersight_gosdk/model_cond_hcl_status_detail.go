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
	"reflect"
	"strings"
)

// CondHclStatusDetail The HCL status detail for each component firmware and driver.
type CondHclStatusDetail struct {
	MoBaseMo
	// The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.
	HardwareStatus *string `json:"HardwareStatus,omitempty"`
	// The current CIMC version for the server normalized for querying HCL data.
	HclCimcVersion *string `json:"HclCimcVersion,omitempty"`
	// The current driver name of the component we are validating normalized for querying HCL data.
	HclDriverName *string `json:"HclDriverName,omitempty"`
	// The current driver version of the component we are validating normalized for querying HCL data.
	HclDriverVersion *string `json:"HclDriverVersion,omitempty"`
	// The current firmware version of the component model normalized for querying HCL data.
	HclFirmwareVersion *string `json:"HclFirmwareVersion,omitempty"`
	// The component model we are trying to validate normalized for querying HCL data.
	HclModel *string `json:"HclModel,omitempty"`
	// The current CIMC version for the server as received from inventory.
	InvCimcVersion *string `json:"InvCimcVersion,omitempty"`
	// The current driver name of the component we are validating as received from inventory.
	InvDriverName *string `json:"InvDriverName,omitempty"`
	// The current driver version of the component we are validating as received from inventory.
	InvDriverVersion *string `json:"InvDriverVersion,omitempty"`
	// The current firmware version of the component model as received from inventory.
	InvFirmwareVersion *string `json:"InvFirmwareVersion,omitempty"`
	// The component model we are trying to validate as received from inventory.
	InvModel *string `json:"InvModel,omitempty"`
	// The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.
	Reason *string `json:"Reason,omitempty"`
	// The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.
	SoftwareStatus *string `json:"SoftwareStatus,omitempty"`
	// The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.
	Status               *string                    `json:"Status,omitempty"`
	Component            *InventoryBaseRelationship `json:"Component,omitempty"`
	HclStatus            *CondHclStatusRelationship `json:"HclStatus,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CondHclStatusDetail CondHclStatusDetail

// NewCondHclStatusDetail instantiates a new CondHclStatusDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCondHclStatusDetail() *CondHclStatusDetail {
	this := CondHclStatusDetail{}
	var hardwareStatus string = "Missing-Os-Driver-Info"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Driver-Info"
	this.Reason = &reason
	var softwareStatus string = "Missing-Os-Driver-Info"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// NewCondHclStatusDetailWithDefaults instantiates a new CondHclStatusDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCondHclStatusDetailWithDefaults() *CondHclStatusDetail {
	this := CondHclStatusDetail{}
	var hardwareStatus string = "Missing-Os-Driver-Info"
	this.HardwareStatus = &hardwareStatus
	var reason string = "Missing-Os-Driver-Info"
	this.Reason = &reason
	var softwareStatus string = "Missing-Os-Driver-Info"
	this.SoftwareStatus = &softwareStatus
	var status string = "Incomplete"
	this.Status = &status
	return &this
}

// GetHardwareStatus returns the HardwareStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHardwareStatus() string {
	if o == nil || o.HardwareStatus == nil {
		var ret string
		return ret
	}
	return *o.HardwareStatus
}

// GetHardwareStatusOk returns a tuple with the HardwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHardwareStatusOk() (*string, bool) {
	if o == nil || o.HardwareStatus == nil {
		return nil, false
	}
	return o.HardwareStatus, true
}

// HasHardwareStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHardwareStatus() bool {
	if o != nil && o.HardwareStatus != nil {
		return true
	}

	return false
}

// SetHardwareStatus gets a reference to the given string and assigns it to the HardwareStatus field.
func (o *CondHclStatusDetail) SetHardwareStatus(v string) {
	o.HardwareStatus = &v
}

// GetHclCimcVersion returns the HclCimcVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclCimcVersion() string {
	if o == nil || o.HclCimcVersion == nil {
		var ret string
		return ret
	}
	return *o.HclCimcVersion
}

// GetHclCimcVersionOk returns a tuple with the HclCimcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclCimcVersionOk() (*string, bool) {
	if o == nil || o.HclCimcVersion == nil {
		return nil, false
	}
	return o.HclCimcVersion, true
}

// HasHclCimcVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclCimcVersion() bool {
	if o != nil && o.HclCimcVersion != nil {
		return true
	}

	return false
}

// SetHclCimcVersion gets a reference to the given string and assigns it to the HclCimcVersion field.
func (o *CondHclStatusDetail) SetHclCimcVersion(v string) {
	o.HclCimcVersion = &v
}

// GetHclDriverName returns the HclDriverName field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclDriverName() string {
	if o == nil || o.HclDriverName == nil {
		var ret string
		return ret
	}
	return *o.HclDriverName
}

// GetHclDriverNameOk returns a tuple with the HclDriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclDriverNameOk() (*string, bool) {
	if o == nil || o.HclDriverName == nil {
		return nil, false
	}
	return o.HclDriverName, true
}

// HasHclDriverName returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclDriverName() bool {
	if o != nil && o.HclDriverName != nil {
		return true
	}

	return false
}

// SetHclDriverName gets a reference to the given string and assigns it to the HclDriverName field.
func (o *CondHclStatusDetail) SetHclDriverName(v string) {
	o.HclDriverName = &v
}

// GetHclDriverVersion returns the HclDriverVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclDriverVersion() string {
	if o == nil || o.HclDriverVersion == nil {
		var ret string
		return ret
	}
	return *o.HclDriverVersion
}

// GetHclDriverVersionOk returns a tuple with the HclDriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclDriverVersionOk() (*string, bool) {
	if o == nil || o.HclDriverVersion == nil {
		return nil, false
	}
	return o.HclDriverVersion, true
}

// HasHclDriverVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclDriverVersion() bool {
	if o != nil && o.HclDriverVersion != nil {
		return true
	}

	return false
}

// SetHclDriverVersion gets a reference to the given string and assigns it to the HclDriverVersion field.
func (o *CondHclStatusDetail) SetHclDriverVersion(v string) {
	o.HclDriverVersion = &v
}

// GetHclFirmwareVersion returns the HclFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclFirmwareVersion() string {
	if o == nil || o.HclFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.HclFirmwareVersion
}

// GetHclFirmwareVersionOk returns a tuple with the HclFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclFirmwareVersionOk() (*string, bool) {
	if o == nil || o.HclFirmwareVersion == nil {
		return nil, false
	}
	return o.HclFirmwareVersion, true
}

// HasHclFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclFirmwareVersion() bool {
	if o != nil && o.HclFirmwareVersion != nil {
		return true
	}

	return false
}

// SetHclFirmwareVersion gets a reference to the given string and assigns it to the HclFirmwareVersion field.
func (o *CondHclStatusDetail) SetHclFirmwareVersion(v string) {
	o.HclFirmwareVersion = &v
}

// GetHclModel returns the HclModel field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclModel() string {
	if o == nil || o.HclModel == nil {
		var ret string
		return ret
	}
	return *o.HclModel
}

// GetHclModelOk returns a tuple with the HclModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclModelOk() (*string, bool) {
	if o == nil || o.HclModel == nil {
		return nil, false
	}
	return o.HclModel, true
}

// HasHclModel returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclModel() bool {
	if o != nil && o.HclModel != nil {
		return true
	}

	return false
}

// SetHclModel gets a reference to the given string and assigns it to the HclModel field.
func (o *CondHclStatusDetail) SetHclModel(v string) {
	o.HclModel = &v
}

// GetInvCimcVersion returns the InvCimcVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetInvCimcVersion() string {
	if o == nil || o.InvCimcVersion == nil {
		var ret string
		return ret
	}
	return *o.InvCimcVersion
}

// GetInvCimcVersionOk returns a tuple with the InvCimcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetInvCimcVersionOk() (*string, bool) {
	if o == nil || o.InvCimcVersion == nil {
		return nil, false
	}
	return o.InvCimcVersion, true
}

// HasInvCimcVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasInvCimcVersion() bool {
	if o != nil && o.InvCimcVersion != nil {
		return true
	}

	return false
}

// SetInvCimcVersion gets a reference to the given string and assigns it to the InvCimcVersion field.
func (o *CondHclStatusDetail) SetInvCimcVersion(v string) {
	o.InvCimcVersion = &v
}

// GetInvDriverName returns the InvDriverName field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetInvDriverName() string {
	if o == nil || o.InvDriverName == nil {
		var ret string
		return ret
	}
	return *o.InvDriverName
}

// GetInvDriverNameOk returns a tuple with the InvDriverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetInvDriverNameOk() (*string, bool) {
	if o == nil || o.InvDriverName == nil {
		return nil, false
	}
	return o.InvDriverName, true
}

// HasInvDriverName returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasInvDriverName() bool {
	if o != nil && o.InvDriverName != nil {
		return true
	}

	return false
}

// SetInvDriverName gets a reference to the given string and assigns it to the InvDriverName field.
func (o *CondHclStatusDetail) SetInvDriverName(v string) {
	o.InvDriverName = &v
}

// GetInvDriverVersion returns the InvDriverVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetInvDriverVersion() string {
	if o == nil || o.InvDriverVersion == nil {
		var ret string
		return ret
	}
	return *o.InvDriverVersion
}

// GetInvDriverVersionOk returns a tuple with the InvDriverVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetInvDriverVersionOk() (*string, bool) {
	if o == nil || o.InvDriverVersion == nil {
		return nil, false
	}
	return o.InvDriverVersion, true
}

// HasInvDriverVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasInvDriverVersion() bool {
	if o != nil && o.InvDriverVersion != nil {
		return true
	}

	return false
}

// SetInvDriverVersion gets a reference to the given string and assigns it to the InvDriverVersion field.
func (o *CondHclStatusDetail) SetInvDriverVersion(v string) {
	o.InvDriverVersion = &v
}

// GetInvFirmwareVersion returns the InvFirmwareVersion field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetInvFirmwareVersion() string {
	if o == nil || o.InvFirmwareVersion == nil {
		var ret string
		return ret
	}
	return *o.InvFirmwareVersion
}

// GetInvFirmwareVersionOk returns a tuple with the InvFirmwareVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetInvFirmwareVersionOk() (*string, bool) {
	if o == nil || o.InvFirmwareVersion == nil {
		return nil, false
	}
	return o.InvFirmwareVersion, true
}

// HasInvFirmwareVersion returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasInvFirmwareVersion() bool {
	if o != nil && o.InvFirmwareVersion != nil {
		return true
	}

	return false
}

// SetInvFirmwareVersion gets a reference to the given string and assigns it to the InvFirmwareVersion field.
func (o *CondHclStatusDetail) SetInvFirmwareVersion(v string) {
	o.InvFirmwareVersion = &v
}

// GetInvModel returns the InvModel field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetInvModel() string {
	if o == nil || o.InvModel == nil {
		var ret string
		return ret
	}
	return *o.InvModel
}

// GetInvModelOk returns a tuple with the InvModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetInvModelOk() (*string, bool) {
	if o == nil || o.InvModel == nil {
		return nil, false
	}
	return o.InvModel, true
}

// HasInvModel returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasInvModel() bool {
	if o != nil && o.InvModel != nil {
		return true
	}

	return false
}

// SetInvModel gets a reference to the given string and assigns it to the InvModel field.
func (o *CondHclStatusDetail) SetInvModel(v string) {
	o.InvModel = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *CondHclStatusDetail) SetReason(v string) {
	o.Reason = &v
}

// GetSoftwareStatus returns the SoftwareStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetSoftwareStatus() string {
	if o == nil || o.SoftwareStatus == nil {
		var ret string
		return ret
	}
	return *o.SoftwareStatus
}

// GetSoftwareStatusOk returns a tuple with the SoftwareStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetSoftwareStatusOk() (*string, bool) {
	if o == nil || o.SoftwareStatus == nil {
		return nil, false
	}
	return o.SoftwareStatus, true
}

// HasSoftwareStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasSoftwareStatus() bool {
	if o != nil && o.SoftwareStatus != nil {
		return true
	}

	return false
}

// SetSoftwareStatus gets a reference to the given string and assigns it to the SoftwareStatus field.
func (o *CondHclStatusDetail) SetSoftwareStatus(v string) {
	o.SoftwareStatus = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CondHclStatusDetail) SetStatus(v string) {
	o.Status = &v
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetComponent() InventoryBaseRelationship {
	if o == nil || o.Component == nil {
		var ret InventoryBaseRelationship
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetComponentOk() (*InventoryBaseRelationship, bool) {
	if o == nil || o.Component == nil {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasComponent() bool {
	if o != nil && o.Component != nil {
		return true
	}

	return false
}

// SetComponent gets a reference to the given InventoryBaseRelationship and assigns it to the Component field.
func (o *CondHclStatusDetail) SetComponent(v InventoryBaseRelationship) {
	o.Component = &v
}

// GetHclStatus returns the HclStatus field value if set, zero value otherwise.
func (o *CondHclStatusDetail) GetHclStatus() CondHclStatusRelationship {
	if o == nil || o.HclStatus == nil {
		var ret CondHclStatusRelationship
		return ret
	}
	return *o.HclStatus
}

// GetHclStatusOk returns a tuple with the HclStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CondHclStatusDetail) GetHclStatusOk() (*CondHclStatusRelationship, bool) {
	if o == nil || o.HclStatus == nil {
		return nil, false
	}
	return o.HclStatus, true
}

// HasHclStatus returns a boolean if a field has been set.
func (o *CondHclStatusDetail) HasHclStatus() bool {
	if o != nil && o.HclStatus != nil {
		return true
	}

	return false
}

// SetHclStatus gets a reference to the given CondHclStatusRelationship and assigns it to the HclStatus field.
func (o *CondHclStatusDetail) SetHclStatus(v CondHclStatusRelationship) {
	o.HclStatus = &v
}

func (o CondHclStatusDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.HardwareStatus != nil {
		toSerialize["HardwareStatus"] = o.HardwareStatus
	}
	if o.HclCimcVersion != nil {
		toSerialize["HclCimcVersion"] = o.HclCimcVersion
	}
	if o.HclDriverName != nil {
		toSerialize["HclDriverName"] = o.HclDriverName
	}
	if o.HclDriverVersion != nil {
		toSerialize["HclDriverVersion"] = o.HclDriverVersion
	}
	if o.HclFirmwareVersion != nil {
		toSerialize["HclFirmwareVersion"] = o.HclFirmwareVersion
	}
	if o.HclModel != nil {
		toSerialize["HclModel"] = o.HclModel
	}
	if o.InvCimcVersion != nil {
		toSerialize["InvCimcVersion"] = o.InvCimcVersion
	}
	if o.InvDriverName != nil {
		toSerialize["InvDriverName"] = o.InvDriverName
	}
	if o.InvDriverVersion != nil {
		toSerialize["InvDriverVersion"] = o.InvDriverVersion
	}
	if o.InvFirmwareVersion != nil {
		toSerialize["InvFirmwareVersion"] = o.InvFirmwareVersion
	}
	if o.InvModel != nil {
		toSerialize["InvModel"] = o.InvModel
	}
	if o.Reason != nil {
		toSerialize["Reason"] = o.Reason
	}
	if o.SoftwareStatus != nil {
		toSerialize["SoftwareStatus"] = o.SoftwareStatus
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Component != nil {
		toSerialize["Component"] = o.Component
	}
	if o.HclStatus != nil {
		toSerialize["HclStatus"] = o.HclStatus
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CondHclStatusDetail) UnmarshalJSON(bytes []byte) (err error) {
	type CondHclStatusDetailWithoutEmbeddedStruct struct {
		// The model is considered as part of the hardware profile for the component. This will provide the HCL validation status for the hardware profile. The reasons can be one of the following \"Incompatible-Server-With-Component\" - the server model and component combination is not listed in HCL \"Incompatible-Firmware\" - The server's firmware is not listed for this component's hardware profile \"Incompatible-Component\" - the component's model is not listed in the HCL \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the hardware profile was not evaulated for the component because the server's hw/sw status is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.
		HardwareStatus *string `json:"HardwareStatus,omitempty"`
		// The current CIMC version for the server normalized for querying HCL data.
		HclCimcVersion *string `json:"HclCimcVersion,omitempty"`
		// The current driver name of the component we are validating normalized for querying HCL data.
		HclDriverName *string `json:"HclDriverName,omitempty"`
		// The current driver version of the component we are validating normalized for querying HCL data.
		HclDriverVersion *string `json:"HclDriverVersion,omitempty"`
		// The current firmware version of the component model normalized for querying HCL data.
		HclFirmwareVersion *string `json:"HclFirmwareVersion,omitempty"`
		// The component model we are trying to validate normalized for querying HCL data.
		HclModel *string `json:"HclModel,omitempty"`
		// The current CIMC version for the server as received from inventory.
		InvCimcVersion *string `json:"InvCimcVersion,omitempty"`
		// The current driver name of the component we are validating as received from inventory.
		InvDriverName *string `json:"InvDriverName,omitempty"`
		// The current driver version of the component we are validating as received from inventory.
		InvDriverVersion *string `json:"InvDriverVersion,omitempty"`
		// The current firmware version of the component model as received from inventory.
		InvFirmwareVersion *string `json:"InvFirmwareVersion,omitempty"`
		// The component model we are trying to validate as received from inventory.
		InvModel *string `json:"InvModel,omitempty"`
		// The reason for the status. The reason can be one of \"Incompatible-Server-With-Component\" - HCL validation has failed because the server model is not validated with this component \"Incompatible-Processor\" - HCL validation has failed because the processor is not validated with this server \"Incompatible-Os-Info\" - HCL validation has failed because the os vendor and version is not validated with this server \"Incompatible-Component-Model\" - HCL validation has failed because the component model is not validated \"Incompatible-Firmware\" - HCL validation has failed because the component or server firmware version is not validated \"Incompatible-Driver\" - HCL validation has failed because the driver version is not validated \"Incompatible-Firmware-Driver\" - HCL validation has failed because the firmware version and driver version is not validated \"Missing-Os-Driver-Info\" - HCL validation was not performed because we are missing os driver information form the inventory \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Service-Error\" - HCL data service is available but an error occured when making the request or parsing the response \"Unrecognized-Protocol\" - This service does not recognize the reason code in the response from the HCL data service \"Compatible\" - this component's inventory data has \"Validated\" status with the HCL. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.
		Reason *string `json:"Reason,omitempty"`
		// The firmware, driver name and driver version are considered as part of the software profile for the component. This will provide the HCL validation status for the software profile. The reasons can be one of the following \"Incompatible-Firmware\" - the component's firmware is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Driver\" - the component's driver is not listed under the server's hardware and software profile and the component's hardware profile \"Incompatible-Firmware-Driver\" - the component's firmware and driver are not listed under the server's hardware and software profile and the component's hardware profile \"Service-Unavailable\" - HCL data service is unavailable at the moment (try again later). This could be due to HCL data updating \"Not-Evaluated\" - the component's hardware status was not evaluated because the server's hardware or software profile is not listed or server is exempted. \"Compatible\" - this component's hardware profile is listed in the HCL.
		SoftwareStatus *string `json:"SoftwareStatus,omitempty"`
		// The status for the component model, firmware version, driver name, and driver version after validating against the HCL. The status can be one of the following \"Unknown\" - we do not have enough information to evaluate against the HCL data \"Validated\" - we have validated this component against the HCL and it has \"Validated\" status \"Not-Validated\" - we have validated this component against the HCL and it has \"Not-Validated\" status. \"Not-Evaluated\" - The component is not evaluated against the HCL because the server is exempted.
		Status    *string                    `json:"Status,omitempty"`
		Component *InventoryBaseRelationship `json:"Component,omitempty"`
		HclStatus *CondHclStatusRelationship `json:"HclStatus,omitempty"`
	}

	varCondHclStatusDetailWithoutEmbeddedStruct := CondHclStatusDetailWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varCondHclStatusDetailWithoutEmbeddedStruct)
	if err == nil {
		varCondHclStatusDetail := _CondHclStatusDetail{}
		varCondHclStatusDetail.HardwareStatus = varCondHclStatusDetailWithoutEmbeddedStruct.HardwareStatus
		varCondHclStatusDetail.HclCimcVersion = varCondHclStatusDetailWithoutEmbeddedStruct.HclCimcVersion
		varCondHclStatusDetail.HclDriverName = varCondHclStatusDetailWithoutEmbeddedStruct.HclDriverName
		varCondHclStatusDetail.HclDriverVersion = varCondHclStatusDetailWithoutEmbeddedStruct.HclDriverVersion
		varCondHclStatusDetail.HclFirmwareVersion = varCondHclStatusDetailWithoutEmbeddedStruct.HclFirmwareVersion
		varCondHclStatusDetail.HclModel = varCondHclStatusDetailWithoutEmbeddedStruct.HclModel
		varCondHclStatusDetail.InvCimcVersion = varCondHclStatusDetailWithoutEmbeddedStruct.InvCimcVersion
		varCondHclStatusDetail.InvDriverName = varCondHclStatusDetailWithoutEmbeddedStruct.InvDriverName
		varCondHclStatusDetail.InvDriverVersion = varCondHclStatusDetailWithoutEmbeddedStruct.InvDriverVersion
		varCondHclStatusDetail.InvFirmwareVersion = varCondHclStatusDetailWithoutEmbeddedStruct.InvFirmwareVersion
		varCondHclStatusDetail.InvModel = varCondHclStatusDetailWithoutEmbeddedStruct.InvModel
		varCondHclStatusDetail.Reason = varCondHclStatusDetailWithoutEmbeddedStruct.Reason
		varCondHclStatusDetail.SoftwareStatus = varCondHclStatusDetailWithoutEmbeddedStruct.SoftwareStatus
		varCondHclStatusDetail.Status = varCondHclStatusDetailWithoutEmbeddedStruct.Status
		varCondHclStatusDetail.Component = varCondHclStatusDetailWithoutEmbeddedStruct.Component
		varCondHclStatusDetail.HclStatus = varCondHclStatusDetailWithoutEmbeddedStruct.HclStatus
		*o = CondHclStatusDetail(varCondHclStatusDetail)
	} else {
		return err
	}

	varCondHclStatusDetail := _CondHclStatusDetail{}

	err = json.Unmarshal(bytes, &varCondHclStatusDetail)
	if err == nil {
		o.MoBaseMo = varCondHclStatusDetail.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "HardwareStatus")
		delete(additionalProperties, "HclCimcVersion")
		delete(additionalProperties, "HclDriverName")
		delete(additionalProperties, "HclDriverVersion")
		delete(additionalProperties, "HclFirmwareVersion")
		delete(additionalProperties, "HclModel")
		delete(additionalProperties, "InvCimcVersion")
		delete(additionalProperties, "InvDriverName")
		delete(additionalProperties, "InvDriverVersion")
		delete(additionalProperties, "InvFirmwareVersion")
		delete(additionalProperties, "InvModel")
		delete(additionalProperties, "Reason")
		delete(additionalProperties, "SoftwareStatus")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Component")
		delete(additionalProperties, "HclStatus")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCondHclStatusDetail struct {
	value *CondHclStatusDetail
	isSet bool
}

func (v NullableCondHclStatusDetail) Get() *CondHclStatusDetail {
	return v.value
}

func (v *NullableCondHclStatusDetail) Set(val *CondHclStatusDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCondHclStatusDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCondHclStatusDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCondHclStatusDetail(val *CondHclStatusDetail) *NullableCondHclStatusDetail {
	return &NullableCondHclStatusDetail{value: val, isSet: true}
}

func (v NullableCondHclStatusDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCondHclStatusDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
