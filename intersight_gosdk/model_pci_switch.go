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

// PciSwitch PCI Switch present in a server connected to two GPUs and one PCIe adapter.
type PciSwitch struct {
	EquipmentBase
	// The device id of the switch.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The composite health of the switch.
	Health *string `json:"Health,omitempty"`
	// The number of GPUs and PCI adapters connected the switch.
	NumOfAdaptors *string `json:"NumOfAdaptors,omitempty"`
	// The PCI address of the switch.
	PciAddress *string `json:"PciAddress,omitempty"`
	// The PCI slot name of the switch.
	PciSlot *string `json:"PciSlot,omitempty"`
	// The model information for the switch.
	ProductName *string `json:"ProductName,omitempty"`
	// The product revision of the switch.
	ProductRevision *string `json:"ProductRevision,omitempty"`
	// The sub device id of the switch.
	SubDeviceId *string `json:"SubDeviceId,omitempty"`
	// The sub vendor id of the switch.
	SubVendorId *string `json:"SubVendorId,omitempty"`
	// The current temperature of the switch.
	Temperature *string `json:"Temperature,omitempty"`
	// The type information of the switch.
	Type *string `json:"Type,omitempty"`
	// The vendor id of the switch.
	VendorId            *string                          `json:"VendorId,omitempty"`
	ComputeBoard        *ComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to pciLink resources.
	Links            []PciLinkRelationship                `json:"Links,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PciSwitch PciSwitch

// NewPciSwitch instantiates a new PciSwitch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPciSwitch() *PciSwitch {
	this := PciSwitch{}
	return &this
}

// NewPciSwitchWithDefaults instantiates a new PciSwitch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPciSwitchWithDefaults() *PciSwitch {
	this := PciSwitch{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *PciSwitch) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *PciSwitch) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *PciSwitch) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *PciSwitch) GetHealth() string {
	if o == nil || o.Health == nil {
		var ret string
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetHealthOk() (*string, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *PciSwitch) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given string and assigns it to the Health field.
func (o *PciSwitch) SetHealth(v string) {
	o.Health = &v
}

// GetNumOfAdaptors returns the NumOfAdaptors field value if set, zero value otherwise.
func (o *PciSwitch) GetNumOfAdaptors() string {
	if o == nil || o.NumOfAdaptors == nil {
		var ret string
		return ret
	}
	return *o.NumOfAdaptors
}

// GetNumOfAdaptorsOk returns a tuple with the NumOfAdaptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetNumOfAdaptorsOk() (*string, bool) {
	if o == nil || o.NumOfAdaptors == nil {
		return nil, false
	}
	return o.NumOfAdaptors, true
}

// HasNumOfAdaptors returns a boolean if a field has been set.
func (o *PciSwitch) HasNumOfAdaptors() bool {
	if o != nil && o.NumOfAdaptors != nil {
		return true
	}

	return false
}

// SetNumOfAdaptors gets a reference to the given string and assigns it to the NumOfAdaptors field.
func (o *PciSwitch) SetNumOfAdaptors(v string) {
	o.NumOfAdaptors = &v
}

// GetPciAddress returns the PciAddress field value if set, zero value otherwise.
func (o *PciSwitch) GetPciAddress() string {
	if o == nil || o.PciAddress == nil {
		var ret string
		return ret
	}
	return *o.PciAddress
}

// GetPciAddressOk returns a tuple with the PciAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetPciAddressOk() (*string, bool) {
	if o == nil || o.PciAddress == nil {
		return nil, false
	}
	return o.PciAddress, true
}

// HasPciAddress returns a boolean if a field has been set.
func (o *PciSwitch) HasPciAddress() bool {
	if o != nil && o.PciAddress != nil {
		return true
	}

	return false
}

// SetPciAddress gets a reference to the given string and assigns it to the PciAddress field.
func (o *PciSwitch) SetPciAddress(v string) {
	o.PciAddress = &v
}

// GetPciSlot returns the PciSlot field value if set, zero value otherwise.
func (o *PciSwitch) GetPciSlot() string {
	if o == nil || o.PciSlot == nil {
		var ret string
		return ret
	}
	return *o.PciSlot
}

// GetPciSlotOk returns a tuple with the PciSlot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetPciSlotOk() (*string, bool) {
	if o == nil || o.PciSlot == nil {
		return nil, false
	}
	return o.PciSlot, true
}

// HasPciSlot returns a boolean if a field has been set.
func (o *PciSwitch) HasPciSlot() bool {
	if o != nil && o.PciSlot != nil {
		return true
	}

	return false
}

// SetPciSlot gets a reference to the given string and assigns it to the PciSlot field.
func (o *PciSwitch) SetPciSlot(v string) {
	o.PciSlot = &v
}

// GetProductName returns the ProductName field value if set, zero value otherwise.
func (o *PciSwitch) GetProductName() string {
	if o == nil || o.ProductName == nil {
		var ret string
		return ret
	}
	return *o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetProductNameOk() (*string, bool) {
	if o == nil || o.ProductName == nil {
		return nil, false
	}
	return o.ProductName, true
}

// HasProductName returns a boolean if a field has been set.
func (o *PciSwitch) HasProductName() bool {
	if o != nil && o.ProductName != nil {
		return true
	}

	return false
}

// SetProductName gets a reference to the given string and assigns it to the ProductName field.
func (o *PciSwitch) SetProductName(v string) {
	o.ProductName = &v
}

// GetProductRevision returns the ProductRevision field value if set, zero value otherwise.
func (o *PciSwitch) GetProductRevision() string {
	if o == nil || o.ProductRevision == nil {
		var ret string
		return ret
	}
	return *o.ProductRevision
}

// GetProductRevisionOk returns a tuple with the ProductRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetProductRevisionOk() (*string, bool) {
	if o == nil || o.ProductRevision == nil {
		return nil, false
	}
	return o.ProductRevision, true
}

// HasProductRevision returns a boolean if a field has been set.
func (o *PciSwitch) HasProductRevision() bool {
	if o != nil && o.ProductRevision != nil {
		return true
	}

	return false
}

// SetProductRevision gets a reference to the given string and assigns it to the ProductRevision field.
func (o *PciSwitch) SetProductRevision(v string) {
	o.ProductRevision = &v
}

// GetSubDeviceId returns the SubDeviceId field value if set, zero value otherwise.
func (o *PciSwitch) GetSubDeviceId() string {
	if o == nil || o.SubDeviceId == nil {
		var ret string
		return ret
	}
	return *o.SubDeviceId
}

// GetSubDeviceIdOk returns a tuple with the SubDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetSubDeviceIdOk() (*string, bool) {
	if o == nil || o.SubDeviceId == nil {
		return nil, false
	}
	return o.SubDeviceId, true
}

// HasSubDeviceId returns a boolean if a field has been set.
func (o *PciSwitch) HasSubDeviceId() bool {
	if o != nil && o.SubDeviceId != nil {
		return true
	}

	return false
}

// SetSubDeviceId gets a reference to the given string and assigns it to the SubDeviceId field.
func (o *PciSwitch) SetSubDeviceId(v string) {
	o.SubDeviceId = &v
}

// GetSubVendorId returns the SubVendorId field value if set, zero value otherwise.
func (o *PciSwitch) GetSubVendorId() string {
	if o == nil || o.SubVendorId == nil {
		var ret string
		return ret
	}
	return *o.SubVendorId
}

// GetSubVendorIdOk returns a tuple with the SubVendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetSubVendorIdOk() (*string, bool) {
	if o == nil || o.SubVendorId == nil {
		return nil, false
	}
	return o.SubVendorId, true
}

// HasSubVendorId returns a boolean if a field has been set.
func (o *PciSwitch) HasSubVendorId() bool {
	if o != nil && o.SubVendorId != nil {
		return true
	}

	return false
}

// SetSubVendorId gets a reference to the given string and assigns it to the SubVendorId field.
func (o *PciSwitch) SetSubVendorId(v string) {
	o.SubVendorId = &v
}

// GetTemperature returns the Temperature field value if set, zero value otherwise.
func (o *PciSwitch) GetTemperature() string {
	if o == nil || o.Temperature == nil {
		var ret string
		return ret
	}
	return *o.Temperature
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetTemperatureOk() (*string, bool) {
	if o == nil || o.Temperature == nil {
		return nil, false
	}
	return o.Temperature, true
}

// HasTemperature returns a boolean if a field has been set.
func (o *PciSwitch) HasTemperature() bool {
	if o != nil && o.Temperature != nil {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given string and assigns it to the Temperature field.
func (o *PciSwitch) SetTemperature(v string) {
	o.Temperature = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PciSwitch) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PciSwitch) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PciSwitch) SetType(v string) {
	o.Type = &v
}

// GetVendorId returns the VendorId field value if set, zero value otherwise.
func (o *PciSwitch) GetVendorId() string {
	if o == nil || o.VendorId == nil {
		var ret string
		return ret
	}
	return *o.VendorId
}

// GetVendorIdOk returns a tuple with the VendorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetVendorIdOk() (*string, bool) {
	if o == nil || o.VendorId == nil {
		return nil, false
	}
	return o.VendorId, true
}

// HasVendorId returns a boolean if a field has been set.
func (o *PciSwitch) HasVendorId() bool {
	if o != nil && o.VendorId != nil {
		return true
	}

	return false
}

// SetVendorId gets a reference to the given string and assigns it to the VendorId field.
func (o *PciSwitch) SetVendorId(v string) {
	o.VendorId = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *PciSwitch) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *PciSwitch) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *PciSwitch) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *PciSwitch) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *PciSwitch) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *PciSwitch) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetLinks() []PciLinkRelationship {
	if o == nil {
		var ret []PciLinkRelationship
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetLinksOk() (*[]PciLinkRelationship, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return &o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PciSwitch) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []PciLinkRelationship and assigns it to the Links field.
func (o *PciSwitch) SetLinks(v []PciLinkRelationship) {
	o.Links = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *PciSwitch) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PciSwitch) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *PciSwitch) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *PciSwitch) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PciSwitch) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PciSwitch) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return &o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *PciSwitch) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *PciSwitch) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o PciSwitch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.NumOfAdaptors != nil {
		toSerialize["NumOfAdaptors"] = o.NumOfAdaptors
	}
	if o.PciAddress != nil {
		toSerialize["PciAddress"] = o.PciAddress
	}
	if o.PciSlot != nil {
		toSerialize["PciSlot"] = o.PciSlot
	}
	if o.ProductName != nil {
		toSerialize["ProductName"] = o.ProductName
	}
	if o.ProductRevision != nil {
		toSerialize["ProductRevision"] = o.ProductRevision
	}
	if o.SubDeviceId != nil {
		toSerialize["SubDeviceId"] = o.SubDeviceId
	}
	if o.SubVendorId != nil {
		toSerialize["SubVendorId"] = o.SubVendorId
	}
	if o.Temperature != nil {
		toSerialize["Temperature"] = o.Temperature
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	if o.VendorId != nil {
		toSerialize["VendorId"] = o.VendorId
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.Links != nil {
		toSerialize["Links"] = o.Links
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PciSwitch) UnmarshalJSON(bytes []byte) (err error) {
	type PciSwitchWithoutEmbeddedStruct struct {
		// The device id of the switch.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The composite health of the switch.
		Health *string `json:"Health,omitempty"`
		// The number of GPUs and PCI adapters connected the switch.
		NumOfAdaptors *string `json:"NumOfAdaptors,omitempty"`
		// The PCI address of the switch.
		PciAddress *string `json:"PciAddress,omitempty"`
		// The PCI slot name of the switch.
		PciSlot *string `json:"PciSlot,omitempty"`
		// The model information for the switch.
		ProductName *string `json:"ProductName,omitempty"`
		// The product revision of the switch.
		ProductRevision *string `json:"ProductRevision,omitempty"`
		// The sub device id of the switch.
		SubDeviceId *string `json:"SubDeviceId,omitempty"`
		// The sub vendor id of the switch.
		SubVendorId *string `json:"SubVendorId,omitempty"`
		// The current temperature of the switch.
		Temperature *string `json:"Temperature,omitempty"`
		// The type information of the switch.
		Type *string `json:"Type,omitempty"`
		// The vendor id of the switch.
		VendorId            *string                          `json:"VendorId,omitempty"`
		ComputeBoard        *ComputeBoardRelationship        `json:"ComputeBoard,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to pciLink resources.
		Links            []PciLinkRelationship                `json:"Links,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
		// An array of relationships to firmwareRunningFirmware resources.
		RunningFirmware []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	}

	varPciSwitchWithoutEmbeddedStruct := PciSwitchWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varPciSwitchWithoutEmbeddedStruct)
	if err == nil {
		varPciSwitch := _PciSwitch{}
		varPciSwitch.DeviceId = varPciSwitchWithoutEmbeddedStruct.DeviceId
		varPciSwitch.Health = varPciSwitchWithoutEmbeddedStruct.Health
		varPciSwitch.NumOfAdaptors = varPciSwitchWithoutEmbeddedStruct.NumOfAdaptors
		varPciSwitch.PciAddress = varPciSwitchWithoutEmbeddedStruct.PciAddress
		varPciSwitch.PciSlot = varPciSwitchWithoutEmbeddedStruct.PciSlot
		varPciSwitch.ProductName = varPciSwitchWithoutEmbeddedStruct.ProductName
		varPciSwitch.ProductRevision = varPciSwitchWithoutEmbeddedStruct.ProductRevision
		varPciSwitch.SubDeviceId = varPciSwitchWithoutEmbeddedStruct.SubDeviceId
		varPciSwitch.SubVendorId = varPciSwitchWithoutEmbeddedStruct.SubVendorId
		varPciSwitch.Temperature = varPciSwitchWithoutEmbeddedStruct.Temperature
		varPciSwitch.Type = varPciSwitchWithoutEmbeddedStruct.Type
		varPciSwitch.VendorId = varPciSwitchWithoutEmbeddedStruct.VendorId
		varPciSwitch.ComputeBoard = varPciSwitchWithoutEmbeddedStruct.ComputeBoard
		varPciSwitch.InventoryDeviceInfo = varPciSwitchWithoutEmbeddedStruct.InventoryDeviceInfo
		varPciSwitch.Links = varPciSwitchWithoutEmbeddedStruct.Links
		varPciSwitch.RegisteredDevice = varPciSwitchWithoutEmbeddedStruct.RegisteredDevice
		varPciSwitch.RunningFirmware = varPciSwitchWithoutEmbeddedStruct.RunningFirmware
		*o = PciSwitch(varPciSwitch)
	} else {
		return err
	}

	varPciSwitch := _PciSwitch{}

	err = json.Unmarshal(bytes, &varPciSwitch)
	if err == nil {
		o.EquipmentBase = varPciSwitch.EquipmentBase
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "NumOfAdaptors")
		delete(additionalProperties, "PciAddress")
		delete(additionalProperties, "PciSlot")
		delete(additionalProperties, "ProductName")
		delete(additionalProperties, "ProductRevision")
		delete(additionalProperties, "SubDeviceId")
		delete(additionalProperties, "SubVendorId")
		delete(additionalProperties, "Temperature")
		delete(additionalProperties, "Type")
		delete(additionalProperties, "VendorId")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Links")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")

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

type NullablePciSwitch struct {
	value *PciSwitch
	isSet bool
}

func (v NullablePciSwitch) Get() *PciSwitch {
	return v.value
}

func (v *NullablePciSwitch) Set(val *PciSwitch) {
	v.value = val
	v.isSet = true
}

func (v NullablePciSwitch) IsSet() bool {
	return v.isSet
}

func (v *NullablePciSwitch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePciSwitch(val *PciSwitch) *NullablePciSwitch {
	return &NullablePciSwitch{value: val, isSet: true}
}

func (v NullablePciSwitch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePciSwitch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
