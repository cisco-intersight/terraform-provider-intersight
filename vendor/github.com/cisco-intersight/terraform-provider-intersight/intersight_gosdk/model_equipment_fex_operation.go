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

// EquipmentFexOperation Models the configuration states of a FEX in Intersight.
type EquipmentFexOperation struct {
	MoBaseMo
	// Action performed on the locator LED for a FEX.
	AdminLocatorLedAction *string `json:"AdminLocatorLedAction,omitempty"`
	// Defines status of action performed on AdminLocatorLedState.
	AdminLocatorLedActionState *string                              `json:"AdminLocatorLedActionState,omitempty"`
	DeviceRegistration         *AssetDeviceRegistrationRelationship `json:"DeviceRegistration,omitempty"`
	Fex                        *EquipmentFexRelationship            `json:"Fex,omitempty"`
}

// NewEquipmentFexOperation instantiates a new EquipmentFexOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentFexOperation() *EquipmentFexOperation {
	this := EquipmentFexOperation{}
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	var adminLocatorLedActionState string = "None"
	this.AdminLocatorLedActionState = &adminLocatorLedActionState
	return &this
}

// NewEquipmentFexOperationWithDefaults instantiates a new EquipmentFexOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentFexOperationWithDefaults() *EquipmentFexOperation {
	this := EquipmentFexOperation{}
	var adminLocatorLedAction string = "None"
	this.AdminLocatorLedAction = &adminLocatorLedAction
	var adminLocatorLedActionState string = "None"
	this.AdminLocatorLedActionState = &adminLocatorLedActionState
	return &this
}

// GetAdminLocatorLedAction returns the AdminLocatorLedAction field value if set, zero value otherwise.
func (o *EquipmentFexOperation) GetAdminLocatorLedAction() string {
	if o == nil || o.AdminLocatorLedAction == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedAction
}

// GetAdminLocatorLedActionOk returns a tuple with the AdminLocatorLedAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexOperation) GetAdminLocatorLedActionOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedAction == nil {
		return nil, false
	}
	return o.AdminLocatorLedAction, true
}

// HasAdminLocatorLedAction returns a boolean if a field has been set.
func (o *EquipmentFexOperation) HasAdminLocatorLedAction() bool {
	if o != nil && o.AdminLocatorLedAction != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedAction gets a reference to the given string and assigns it to the AdminLocatorLedAction field.
func (o *EquipmentFexOperation) SetAdminLocatorLedAction(v string) {
	o.AdminLocatorLedAction = &v
}

// GetAdminLocatorLedActionState returns the AdminLocatorLedActionState field value if set, zero value otherwise.
func (o *EquipmentFexOperation) GetAdminLocatorLedActionState() string {
	if o == nil || o.AdminLocatorLedActionState == nil {
		var ret string
		return ret
	}
	return *o.AdminLocatorLedActionState
}

// GetAdminLocatorLedActionStateOk returns a tuple with the AdminLocatorLedActionState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexOperation) GetAdminLocatorLedActionStateOk() (*string, bool) {
	if o == nil || o.AdminLocatorLedActionState == nil {
		return nil, false
	}
	return o.AdminLocatorLedActionState, true
}

// HasAdminLocatorLedActionState returns a boolean if a field has been set.
func (o *EquipmentFexOperation) HasAdminLocatorLedActionState() bool {
	if o != nil && o.AdminLocatorLedActionState != nil {
		return true
	}

	return false
}

// SetAdminLocatorLedActionState gets a reference to the given string and assigns it to the AdminLocatorLedActionState field.
func (o *EquipmentFexOperation) SetAdminLocatorLedActionState(v string) {
	o.AdminLocatorLedActionState = &v
}

// GetDeviceRegistration returns the DeviceRegistration field value if set, zero value otherwise.
func (o *EquipmentFexOperation) GetDeviceRegistration() AssetDeviceRegistrationRelationship {
	if o == nil || o.DeviceRegistration == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.DeviceRegistration
}

// GetDeviceRegistrationOk returns a tuple with the DeviceRegistration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexOperation) GetDeviceRegistrationOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.DeviceRegistration == nil {
		return nil, false
	}
	return o.DeviceRegistration, true
}

// HasDeviceRegistration returns a boolean if a field has been set.
func (o *EquipmentFexOperation) HasDeviceRegistration() bool {
	if o != nil && o.DeviceRegistration != nil {
		return true
	}

	return false
}

// SetDeviceRegistration gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the DeviceRegistration field.
func (o *EquipmentFexOperation) SetDeviceRegistration(v AssetDeviceRegistrationRelationship) {
	o.DeviceRegistration = &v
}

// GetFex returns the Fex field value if set, zero value otherwise.
func (o *EquipmentFexOperation) GetFex() EquipmentFexRelationship {
	if o == nil || o.Fex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.Fex
}

// GetFexOk returns a tuple with the Fex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentFexOperation) GetFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.Fex == nil {
		return nil, false
	}
	return o.Fex, true
}

// HasFex returns a boolean if a field has been set.
func (o *EquipmentFexOperation) HasFex() bool {
	if o != nil && o.Fex != nil {
		return true
	}

	return false
}

// SetFex gets a reference to the given EquipmentFexRelationship and assigns it to the Fex field.
func (o *EquipmentFexOperation) SetFex(v EquipmentFexRelationship) {
	o.Fex = &v
}

func (o EquipmentFexOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.AdminLocatorLedAction != nil {
		toSerialize["AdminLocatorLedAction"] = o.AdminLocatorLedAction
	}
	if o.AdminLocatorLedActionState != nil {
		toSerialize["AdminLocatorLedActionState"] = o.AdminLocatorLedActionState
	}
	if o.DeviceRegistration != nil {
		toSerialize["DeviceRegistration"] = o.DeviceRegistration
	}
	if o.Fex != nil {
		toSerialize["Fex"] = o.Fex
	}
	return json.Marshal(toSerialize)
}

type NullableEquipmentFexOperation struct {
	value *EquipmentFexOperation
	isSet bool
}

func (v NullableEquipmentFexOperation) Get() *EquipmentFexOperation {
	return v.value
}

func (v *NullableEquipmentFexOperation) Set(val *EquipmentFexOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentFexOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentFexOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentFexOperation(val *EquipmentFexOperation) *NullableEquipmentFexOperation {
	return &NullableEquipmentFexOperation{value: val, isSet: true}
}

func (v NullableEquipmentFexOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentFexOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}