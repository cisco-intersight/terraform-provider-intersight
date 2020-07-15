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

// EquipmentPsu This represents power supply unit for chassis/server.
type EquipmentPsu struct {
	EquipmentBase
	// This field is to provide description for the power supply unit.
	Description *string `json:"Description,omitempty"`
	// This field identifies the psu operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this Power Supply Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the Product ID for the Power Supply.
	Pid *string `json:"Pid,omitempty"`
	// This field identifies the presence state of the psu.
	Presence *string `json:"Presence,omitempty"`
	// This field identifies the Firmware Version of the Power Supply.
	PsuFwVersion *string `json:"PsuFwVersion,omitempty"`
	// This represents power supply unit identifier in chassis/server.
	PsuId *int64 `json:"PsuId,omitempty"`
	// This field identifies the input source for the Power Supply.
	PsuInputSrc *string `json:"PsuInputSrc,omitempty"`
	// This field identifies the type of the Power Supply.
	PsuType *string `json:"PsuType,omitempty"`
	// This field identifies the Wattage of the Power Supply.
	PsuWattage *string `json:"PsuWattage,omitempty"`
	// This field identifies the Stockkeeping Unit for this Power Supply.
	Sku *string `json:"Sku,omitempty"`
	// This field identifies the Vendor ID for this Power Supply Unit.
	Vid *string `json:"Vid,omitempty"`
	// This field is used to indicate the Voltage for this Power Supply.
	Voltage                *string                              `json:"Voltage,omitempty"`
	ComputeRackUnit        *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
	EquipmentChassis       *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
	EquipmentFex           *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
	EquipmentRackEnclosure *EquipmentRackEnclosureRelationship  `json:"EquipmentRackEnclosure,omitempty"`
	InventoryDeviceInfo    *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	NetworkElement         *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties   map[string]interface{}
}

type _EquipmentPsu EquipmentPsu

// NewEquipmentPsu instantiates a new EquipmentPsu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentPsu() *EquipmentPsu {
	this := EquipmentPsu{}
	return &this
}

// NewEquipmentPsuWithDefaults instantiates a new EquipmentPsu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentPsuWithDefaults() *EquipmentPsu {
	this := EquipmentPsu{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EquipmentPsu) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EquipmentPsu) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EquipmentPsu) SetDescription(v string) {
	o.Description = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentPsu) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentPsu) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentPsu) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentPsu) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPid() string {
	if o == nil || o.Pid == nil {
		var ret string
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPidOk() (*string, bool) {
	if o == nil || o.Pid == nil {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPid() bool {
	if o != nil && o.Pid != nil {
		return true
	}

	return false
}

// SetPid gets a reference to the given string and assigns it to the Pid field.
func (o *EquipmentPsu) SetPid(v string) {
	o.Pid = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *EquipmentPsu) SetPresence(v string) {
	o.Presence = &v
}

// GetPsuFwVersion returns the PsuFwVersion field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPsuFwVersion() string {
	if o == nil || o.PsuFwVersion == nil {
		var ret string
		return ret
	}
	return *o.PsuFwVersion
}

// GetPsuFwVersionOk returns a tuple with the PsuFwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPsuFwVersionOk() (*string, bool) {
	if o == nil || o.PsuFwVersion == nil {
		return nil, false
	}
	return o.PsuFwVersion, true
}

// HasPsuFwVersion returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPsuFwVersion() bool {
	if o != nil && o.PsuFwVersion != nil {
		return true
	}

	return false
}

// SetPsuFwVersion gets a reference to the given string and assigns it to the PsuFwVersion field.
func (o *EquipmentPsu) SetPsuFwVersion(v string) {
	o.PsuFwVersion = &v
}

// GetPsuId returns the PsuId field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPsuId() int64 {
	if o == nil || o.PsuId == nil {
		var ret int64
		return ret
	}
	return *o.PsuId
}

// GetPsuIdOk returns a tuple with the PsuId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPsuIdOk() (*int64, bool) {
	if o == nil || o.PsuId == nil {
		return nil, false
	}
	return o.PsuId, true
}

// HasPsuId returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPsuId() bool {
	if o != nil && o.PsuId != nil {
		return true
	}

	return false
}

// SetPsuId gets a reference to the given int64 and assigns it to the PsuId field.
func (o *EquipmentPsu) SetPsuId(v int64) {
	o.PsuId = &v
}

// GetPsuInputSrc returns the PsuInputSrc field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPsuInputSrc() string {
	if o == nil || o.PsuInputSrc == nil {
		var ret string
		return ret
	}
	return *o.PsuInputSrc
}

// GetPsuInputSrcOk returns a tuple with the PsuInputSrc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPsuInputSrcOk() (*string, bool) {
	if o == nil || o.PsuInputSrc == nil {
		return nil, false
	}
	return o.PsuInputSrc, true
}

// HasPsuInputSrc returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPsuInputSrc() bool {
	if o != nil && o.PsuInputSrc != nil {
		return true
	}

	return false
}

// SetPsuInputSrc gets a reference to the given string and assigns it to the PsuInputSrc field.
func (o *EquipmentPsu) SetPsuInputSrc(v string) {
	o.PsuInputSrc = &v
}

// GetPsuType returns the PsuType field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPsuType() string {
	if o == nil || o.PsuType == nil {
		var ret string
		return ret
	}
	return *o.PsuType
}

// GetPsuTypeOk returns a tuple with the PsuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPsuTypeOk() (*string, bool) {
	if o == nil || o.PsuType == nil {
		return nil, false
	}
	return o.PsuType, true
}

// HasPsuType returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPsuType() bool {
	if o != nil && o.PsuType != nil {
		return true
	}

	return false
}

// SetPsuType gets a reference to the given string and assigns it to the PsuType field.
func (o *EquipmentPsu) SetPsuType(v string) {
	o.PsuType = &v
}

// GetPsuWattage returns the PsuWattage field value if set, zero value otherwise.
func (o *EquipmentPsu) GetPsuWattage() string {
	if o == nil || o.PsuWattage == nil {
		var ret string
		return ret
	}
	return *o.PsuWattage
}

// GetPsuWattageOk returns a tuple with the PsuWattage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetPsuWattageOk() (*string, bool) {
	if o == nil || o.PsuWattage == nil {
		return nil, false
	}
	return o.PsuWattage, true
}

// HasPsuWattage returns a boolean if a field has been set.
func (o *EquipmentPsu) HasPsuWattage() bool {
	if o != nil && o.PsuWattage != nil {
		return true
	}

	return false
}

// SetPsuWattage gets a reference to the given string and assigns it to the PsuWattage field.
func (o *EquipmentPsu) SetPsuWattage(v string) {
	o.PsuWattage = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EquipmentPsu) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EquipmentPsu) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EquipmentPsu) SetSku(v string) {
	o.Sku = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentPsu) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentPsu) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentPsu) SetVid(v string) {
	o.Vid = &v
}

// GetVoltage returns the Voltage field value if set, zero value otherwise.
func (o *EquipmentPsu) GetVoltage() string {
	if o == nil || o.Voltage == nil {
		var ret string
		return ret
	}
	return *o.Voltage
}

// GetVoltageOk returns a tuple with the Voltage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetVoltageOk() (*string, bool) {
	if o == nil || o.Voltage == nil {
		return nil, false
	}
	return o.Voltage, true
}

// HasVoltage returns a boolean if a field has been set.
func (o *EquipmentPsu) HasVoltage() bool {
	if o != nil && o.Voltage != nil {
		return true
	}

	return false
}

// SetVoltage gets a reference to the given string and assigns it to the Voltage field.
func (o *EquipmentPsu) SetVoltage(v string) {
	o.Voltage = &v
}

// GetComputeRackUnit returns the ComputeRackUnit field value if set, zero value otherwise.
func (o *EquipmentPsu) GetComputeRackUnit() ComputeRackUnitRelationship {
	if o == nil || o.ComputeRackUnit == nil {
		var ret ComputeRackUnitRelationship
		return ret
	}
	return *o.ComputeRackUnit
}

// GetComputeRackUnitOk returns a tuple with the ComputeRackUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetComputeRackUnitOk() (*ComputeRackUnitRelationship, bool) {
	if o == nil || o.ComputeRackUnit == nil {
		return nil, false
	}
	return o.ComputeRackUnit, true
}

// HasComputeRackUnit returns a boolean if a field has been set.
func (o *EquipmentPsu) HasComputeRackUnit() bool {
	if o != nil && o.ComputeRackUnit != nil {
		return true
	}

	return false
}

// SetComputeRackUnit gets a reference to the given ComputeRackUnitRelationship and assigns it to the ComputeRackUnit field.
func (o *EquipmentPsu) SetComputeRackUnit(v ComputeRackUnitRelationship) {
	o.ComputeRackUnit = &v
}

// GetEquipmentChassis returns the EquipmentChassis field value if set, zero value otherwise.
func (o *EquipmentPsu) GetEquipmentChassis() EquipmentChassisRelationship {
	if o == nil || o.EquipmentChassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.EquipmentChassis
}

// GetEquipmentChassisOk returns a tuple with the EquipmentChassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetEquipmentChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.EquipmentChassis == nil {
		return nil, false
	}
	return o.EquipmentChassis, true
}

// HasEquipmentChassis returns a boolean if a field has been set.
func (o *EquipmentPsu) HasEquipmentChassis() bool {
	if o != nil && o.EquipmentChassis != nil {
		return true
	}

	return false
}

// SetEquipmentChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the EquipmentChassis field.
func (o *EquipmentPsu) SetEquipmentChassis(v EquipmentChassisRelationship) {
	o.EquipmentChassis = &v
}

// GetEquipmentFex returns the EquipmentFex field value if set, zero value otherwise.
func (o *EquipmentPsu) GetEquipmentFex() EquipmentFexRelationship {
	if o == nil || o.EquipmentFex == nil {
		var ret EquipmentFexRelationship
		return ret
	}
	return *o.EquipmentFex
}

// GetEquipmentFexOk returns a tuple with the EquipmentFex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetEquipmentFexOk() (*EquipmentFexRelationship, bool) {
	if o == nil || o.EquipmentFex == nil {
		return nil, false
	}
	return o.EquipmentFex, true
}

// HasEquipmentFex returns a boolean if a field has been set.
func (o *EquipmentPsu) HasEquipmentFex() bool {
	if o != nil && o.EquipmentFex != nil {
		return true
	}

	return false
}

// SetEquipmentFex gets a reference to the given EquipmentFexRelationship and assigns it to the EquipmentFex field.
func (o *EquipmentPsu) SetEquipmentFex(v EquipmentFexRelationship) {
	o.EquipmentFex = &v
}

// GetEquipmentRackEnclosure returns the EquipmentRackEnclosure field value if set, zero value otherwise.
func (o *EquipmentPsu) GetEquipmentRackEnclosure() EquipmentRackEnclosureRelationship {
	if o == nil || o.EquipmentRackEnclosure == nil {
		var ret EquipmentRackEnclosureRelationship
		return ret
	}
	return *o.EquipmentRackEnclosure
}

// GetEquipmentRackEnclosureOk returns a tuple with the EquipmentRackEnclosure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetEquipmentRackEnclosureOk() (*EquipmentRackEnclosureRelationship, bool) {
	if o == nil || o.EquipmentRackEnclosure == nil {
		return nil, false
	}
	return o.EquipmentRackEnclosure, true
}

// HasEquipmentRackEnclosure returns a boolean if a field has been set.
func (o *EquipmentPsu) HasEquipmentRackEnclosure() bool {
	if o != nil && o.EquipmentRackEnclosure != nil {
		return true
	}

	return false
}

// SetEquipmentRackEnclosure gets a reference to the given EquipmentRackEnclosureRelationship and assigns it to the EquipmentRackEnclosure field.
func (o *EquipmentPsu) SetEquipmentRackEnclosure(v EquipmentRackEnclosureRelationship) {
	o.EquipmentRackEnclosure = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentPsu) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentPsu) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentPsu) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetNetworkElement returns the NetworkElement field value if set, zero value otherwise.
func (o *EquipmentPsu) GetNetworkElement() NetworkElementRelationship {
	if o == nil || o.NetworkElement == nil {
		var ret NetworkElementRelationship
		return ret
	}
	return *o.NetworkElement
}

// GetNetworkElementOk returns a tuple with the NetworkElement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetNetworkElementOk() (*NetworkElementRelationship, bool) {
	if o == nil || o.NetworkElement == nil {
		return nil, false
	}
	return o.NetworkElement, true
}

// HasNetworkElement returns a boolean if a field has been set.
func (o *EquipmentPsu) HasNetworkElement() bool {
	if o != nil && o.NetworkElement != nil {
		return true
	}

	return false
}

// SetNetworkElement gets a reference to the given NetworkElementRelationship and assigns it to the NetworkElement field.
func (o *EquipmentPsu) SetNetworkElement(v NetworkElementRelationship) {
	o.NetworkElement = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentPsu) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentPsu) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentPsu) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentPsu) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentPsu) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Pid != nil {
		toSerialize["Pid"] = o.Pid
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.PsuFwVersion != nil {
		toSerialize["PsuFwVersion"] = o.PsuFwVersion
	}
	if o.PsuId != nil {
		toSerialize["PsuId"] = o.PsuId
	}
	if o.PsuInputSrc != nil {
		toSerialize["PsuInputSrc"] = o.PsuInputSrc
	}
	if o.PsuType != nil {
		toSerialize["PsuType"] = o.PsuType
	}
	if o.PsuWattage != nil {
		toSerialize["PsuWattage"] = o.PsuWattage
	}
	if o.Sku != nil {
		toSerialize["Sku"] = o.Sku
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Voltage != nil {
		toSerialize["Voltage"] = o.Voltage
	}
	if o.ComputeRackUnit != nil {
		toSerialize["ComputeRackUnit"] = o.ComputeRackUnit
	}
	if o.EquipmentChassis != nil {
		toSerialize["EquipmentChassis"] = o.EquipmentChassis
	}
	if o.EquipmentFex != nil {
		toSerialize["EquipmentFex"] = o.EquipmentFex
	}
	if o.EquipmentRackEnclosure != nil {
		toSerialize["EquipmentRackEnclosure"] = o.EquipmentRackEnclosure
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.NetworkElement != nil {
		toSerialize["NetworkElement"] = o.NetworkElement
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentPsu) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentPsuWithoutEmbeddedStruct struct {
		// This field is to provide description for the power supply unit.
		Description *string `json:"Description,omitempty"`
		// This field identifies the psu operational state.
		OperState *string `json:"OperState,omitempty"`
		// This field identifies the Part Number for this Power Supply Unit.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the Product ID for the Power Supply.
		Pid *string `json:"Pid,omitempty"`
		// This field identifies the presence state of the psu.
		Presence *string `json:"Presence,omitempty"`
		// This field identifies the Firmware Version of the Power Supply.
		PsuFwVersion *string `json:"PsuFwVersion,omitempty"`
		// This represents power supply unit identifier in chassis/server.
		PsuId *int64 `json:"PsuId,omitempty"`
		// This field identifies the input source for the Power Supply.
		PsuInputSrc *string `json:"PsuInputSrc,omitempty"`
		// This field identifies the type of the Power Supply.
		PsuType *string `json:"PsuType,omitempty"`
		// This field identifies the Wattage of the Power Supply.
		PsuWattage *string `json:"PsuWattage,omitempty"`
		// This field identifies the Stockkeeping Unit for this Power Supply.
		Sku *string `json:"Sku,omitempty"`
		// This field identifies the Vendor ID for this Power Supply Unit.
		Vid *string `json:"Vid,omitempty"`
		// This field is used to indicate the Voltage for this Power Supply.
		Voltage                *string                              `json:"Voltage,omitempty"`
		ComputeRackUnit        *ComputeRackUnitRelationship         `json:"ComputeRackUnit,omitempty"`
		EquipmentChassis       *EquipmentChassisRelationship        `json:"EquipmentChassis,omitempty"`
		EquipmentFex           *EquipmentFexRelationship            `json:"EquipmentFex,omitempty"`
		EquipmentRackEnclosure *EquipmentRackEnclosureRelationship  `json:"EquipmentRackEnclosure,omitempty"`
		InventoryDeviceInfo    *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		NetworkElement         *NetworkElementRelationship          `json:"NetworkElement,omitempty"`
		RegisteredDevice       *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentPsuWithoutEmbeddedStruct := EquipmentPsuWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentPsuWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentPsu := _EquipmentPsu{}
		varEquipmentPsu.Description = varEquipmentPsuWithoutEmbeddedStruct.Description
		varEquipmentPsu.OperState = varEquipmentPsuWithoutEmbeddedStruct.OperState
		varEquipmentPsu.PartNumber = varEquipmentPsuWithoutEmbeddedStruct.PartNumber
		varEquipmentPsu.Pid = varEquipmentPsuWithoutEmbeddedStruct.Pid
		varEquipmentPsu.Presence = varEquipmentPsuWithoutEmbeddedStruct.Presence
		varEquipmentPsu.PsuFwVersion = varEquipmentPsuWithoutEmbeddedStruct.PsuFwVersion
		varEquipmentPsu.PsuId = varEquipmentPsuWithoutEmbeddedStruct.PsuId
		varEquipmentPsu.PsuInputSrc = varEquipmentPsuWithoutEmbeddedStruct.PsuInputSrc
		varEquipmentPsu.PsuType = varEquipmentPsuWithoutEmbeddedStruct.PsuType
		varEquipmentPsu.PsuWattage = varEquipmentPsuWithoutEmbeddedStruct.PsuWattage
		varEquipmentPsu.Sku = varEquipmentPsuWithoutEmbeddedStruct.Sku
		varEquipmentPsu.Vid = varEquipmentPsuWithoutEmbeddedStruct.Vid
		varEquipmentPsu.Voltage = varEquipmentPsuWithoutEmbeddedStruct.Voltage
		varEquipmentPsu.ComputeRackUnit = varEquipmentPsuWithoutEmbeddedStruct.ComputeRackUnit
		varEquipmentPsu.EquipmentChassis = varEquipmentPsuWithoutEmbeddedStruct.EquipmentChassis
		varEquipmentPsu.EquipmentFex = varEquipmentPsuWithoutEmbeddedStruct.EquipmentFex
		varEquipmentPsu.EquipmentRackEnclosure = varEquipmentPsuWithoutEmbeddedStruct.EquipmentRackEnclosure
		varEquipmentPsu.InventoryDeviceInfo = varEquipmentPsuWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentPsu.NetworkElement = varEquipmentPsuWithoutEmbeddedStruct.NetworkElement
		varEquipmentPsu.RegisteredDevice = varEquipmentPsuWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentPsu(varEquipmentPsu)
	} else {
		return err
	}

	varEquipmentPsu := _EquipmentPsu{}

	err = json.Unmarshal(bytes, &varEquipmentPsu)
	if err == nil {
		o.EquipmentBase = varEquipmentPsu.EquipmentBase
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Description")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Pid")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "PsuFwVersion")
		delete(additionalProperties, "PsuId")
		delete(additionalProperties, "PsuInputSrc")
		delete(additionalProperties, "PsuType")
		delete(additionalProperties, "PsuWattage")
		delete(additionalProperties, "Sku")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Voltage")
		delete(additionalProperties, "ComputeRackUnit")
		delete(additionalProperties, "EquipmentChassis")
		delete(additionalProperties, "EquipmentFex")
		delete(additionalProperties, "EquipmentRackEnclosure")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "NetworkElement")
		delete(additionalProperties, "RegisteredDevice")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableEquipmentPsu struct {
	value *EquipmentPsu
	isSet bool
}

func (v NullableEquipmentPsu) Get() *EquipmentPsu {
	return v.value
}

func (v *NullableEquipmentPsu) Set(val *EquipmentPsu) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentPsu) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentPsu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentPsu(val *EquipmentPsu) *NullableEquipmentPsu {
	return &NullableEquipmentPsu{value: val, isSet: true}
}

func (v NullableEquipmentPsu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentPsu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
