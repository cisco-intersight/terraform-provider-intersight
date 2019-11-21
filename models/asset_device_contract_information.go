// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AssetDeviceContractInformation Asset:Device Contract Information
//
// Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.
//
// swagger:model assetDeviceContractInformation
type AssetDeviceContractInformation struct {
	MoBaseMo

	// Contract information for the Cisco support contract purchased for the Cisco device.
	//
	// Read Only: true
	Contract *AssetContractInformation `json:"Contract,omitempty"`

	// Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.
	//
	// Read Only: true
	// Enum: [Not Covered Active Expiring Soon]
	ContractStatus string `json:"ContractStatus,omitempty"`

	// End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.
	//
	// Read Only: true
	CoveredProductLineEndDate string `json:"CoveredProductLineEndDate,omitempty"`

	// Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.
	//
	// Read Only: true
	DeviceID string `json:"DeviceId,omitempty"`

	// Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.
	//
	// Read Only: true
	// Enum: [None CiscoUcsServer CiscoUcsFI CiscoUcsChassis]
	DeviceType string `json:"DeviceType,omitempty"`

	// End customer information for the Cisco support contract purchased for the Cisco device.
	//
	// Read Only: true
	EndCustomer *AssetCustomerInformation `json:"EndCustomer,omitempty"`

	// EndUserGlobalUltimate information listed in the contract.
	//
	// Read Only: true
	EndUserGlobalUltimate *AssetGlobalUltimate `json:"EndUserGlobalUltimate,omitempty"`

	// Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.
	//
	// Read Only: true
	IsValid *bool `json:"IsValid,omitempty"`

	// Item type of this specific Cisco device. example "Chassis".
	//
	// Read Only: true
	ItemType string `json:"ItemType,omitempty"`

	// Maintenance purchase order number for the Cisco device.
	//
	// Read Only: true
	MaintenancePurchaseOrderNumber string `json:"MaintenancePurchaseOrderNumber,omitempty"`

	// Maintenance sales order number for the Cisco device.
	//
	// Read Only: true
	MaintenanceSalesOrderNumber string `json:"MaintenanceSalesOrderNumber,omitempty"`

	// The platform type of the Cisco device.
	//
	// Read Only: true
	// Enum: [ APIC DCNM UCSFI IMC IMCM4 IMCM5 HX HXTriton UCSD IntersightAppliance PureStorage VMware ServiceEngine]
	PlatformType string `json:"PlatformType,omitempty"`

	// Product information of the offering under a contract.
	//
	// Read Only: true
	Product *AssetProductInformation `json:"Product,omitempty"`

	// Purchase order number for the Cisco device. It is a unique number assigned for every purchase.
	//
	// Read Only: true
	PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty"`

	// Reference to the device connector through which the device is connected.
	//
	// Read Only: true
	RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

	// ResellerGlobalUltimate information listed in the contract.
	//
	// Read Only: true
	ResellerGlobalUltimate *AssetGlobalUltimate `json:"ResellerGlobalUltimate,omitempty"`

	// Sales order number for the Cisco device. It is a unique number assigned for every sale.
	//
	// Read Only: true
	SalesOrderNumber string `json:"SalesOrderNumber,omitempty"`

	// The type of service contract that covers the Cisco device.
	//
	// Read Only: true
	ServiceDescription string `json:"ServiceDescription,omitempty"`

	// End date for the Cisco service contract that covers this Cisco device.
	//
	// Read Only: true
	// Format: date-time
	ServiceEndDate strfmt.DateTime `json:"ServiceEndDate,omitempty"`

	// The type of service contract that covers the Cisco device.
	//
	// Read Only: true
	ServiceLevel string `json:"ServiceLevel,omitempty"`

	// The SKU of the service contract that covers the Cisco device.
	//
	// Read Only: true
	ServiceSku string `json:"ServiceSku,omitempty"`

	// Start date for the Cisco service contract that covers this Cisco device.
	//
	// Read Only: true
	// Format: date-time
	ServiceStartDate strfmt.DateTime `json:"ServiceStartDate,omitempty"`

	// End date for the warranty that covers the Cisco device.
	//
	// Read Only: true
	WarrantyEndDate string `json:"WarrantyEndDate,omitempty"`

	// Type of warranty that covers the Cisco device.
	//
	// Read Only: true
	WarrantyType string `json:"WarrantyType,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *AssetDeviceContractInformation) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		Contract *AssetContractInformation `json:"Contract,omitempty"`

		ContractStatus string `json:"ContractStatus,omitempty"`

		CoveredProductLineEndDate string `json:"CoveredProductLineEndDate,omitempty"`

		DeviceID string `json:"DeviceId,omitempty"`

		DeviceType string `json:"DeviceType,omitempty"`

		EndCustomer *AssetCustomerInformation `json:"EndCustomer,omitempty"`

		EndUserGlobalUltimate *AssetGlobalUltimate `json:"EndUserGlobalUltimate,omitempty"`

		IsValid *bool `json:"IsValid,omitempty"`

		ItemType string `json:"ItemType,omitempty"`

		MaintenancePurchaseOrderNumber string `json:"MaintenancePurchaseOrderNumber,omitempty"`

		MaintenanceSalesOrderNumber string `json:"MaintenanceSalesOrderNumber,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		Product *AssetProductInformation `json:"Product,omitempty"`

		PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ResellerGlobalUltimate *AssetGlobalUltimate `json:"ResellerGlobalUltimate,omitempty"`

		SalesOrderNumber string `json:"SalesOrderNumber,omitempty"`

		ServiceDescription string `json:"ServiceDescription,omitempty"`

		ServiceEndDate strfmt.DateTime `json:"ServiceEndDate,omitempty"`

		ServiceLevel string `json:"ServiceLevel,omitempty"`

		ServiceSku string `json:"ServiceSku,omitempty"`

		ServiceStartDate strfmt.DateTime `json:"ServiceStartDate,omitempty"`

		WarrantyEndDate string `json:"WarrantyEndDate,omitempty"`

		WarrantyType string `json:"WarrantyType,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Contract = dataAO1.Contract

	m.ContractStatus = dataAO1.ContractStatus

	m.CoveredProductLineEndDate = dataAO1.CoveredProductLineEndDate

	m.DeviceID = dataAO1.DeviceID

	m.DeviceType = dataAO1.DeviceType

	m.EndCustomer = dataAO1.EndCustomer

	m.EndUserGlobalUltimate = dataAO1.EndUserGlobalUltimate

	m.IsValid = dataAO1.IsValid

	m.ItemType = dataAO1.ItemType

	m.MaintenancePurchaseOrderNumber = dataAO1.MaintenancePurchaseOrderNumber

	m.MaintenanceSalesOrderNumber = dataAO1.MaintenanceSalesOrderNumber

	m.PlatformType = dataAO1.PlatformType

	m.Product = dataAO1.Product

	m.PurchaseOrderNumber = dataAO1.PurchaseOrderNumber

	m.RegisteredDevice = dataAO1.RegisteredDevice

	m.ResellerGlobalUltimate = dataAO1.ResellerGlobalUltimate

	m.SalesOrderNumber = dataAO1.SalesOrderNumber

	m.ServiceDescription = dataAO1.ServiceDescription

	m.ServiceEndDate = dataAO1.ServiceEndDate

	m.ServiceLevel = dataAO1.ServiceLevel

	m.ServiceSku = dataAO1.ServiceSku

	m.ServiceStartDate = dataAO1.ServiceStartDate

	m.WarrantyEndDate = dataAO1.WarrantyEndDate

	m.WarrantyType = dataAO1.WarrantyType

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m AssetDeviceContractInformation) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Contract *AssetContractInformation `json:"Contract,omitempty"`

		ContractStatus string `json:"ContractStatus,omitempty"`

		CoveredProductLineEndDate string `json:"CoveredProductLineEndDate,omitempty"`

		DeviceID string `json:"DeviceId,omitempty"`

		DeviceType string `json:"DeviceType,omitempty"`

		EndCustomer *AssetCustomerInformation `json:"EndCustomer,omitempty"`

		EndUserGlobalUltimate *AssetGlobalUltimate `json:"EndUserGlobalUltimate,omitempty"`

		IsValid *bool `json:"IsValid,omitempty"`

		ItemType string `json:"ItemType,omitempty"`

		MaintenancePurchaseOrderNumber string `json:"MaintenancePurchaseOrderNumber,omitempty"`

		MaintenanceSalesOrderNumber string `json:"MaintenanceSalesOrderNumber,omitempty"`

		PlatformType string `json:"PlatformType,omitempty"`

		Product *AssetProductInformation `json:"Product,omitempty"`

		PurchaseOrderNumber string `json:"PurchaseOrderNumber,omitempty"`

		RegisteredDevice *AssetDeviceRegistrationRef `json:"RegisteredDevice,omitempty"`

		ResellerGlobalUltimate *AssetGlobalUltimate `json:"ResellerGlobalUltimate,omitempty"`

		SalesOrderNumber string `json:"SalesOrderNumber,omitempty"`

		ServiceDescription string `json:"ServiceDescription,omitempty"`

		ServiceEndDate strfmt.DateTime `json:"ServiceEndDate,omitempty"`

		ServiceLevel string `json:"ServiceLevel,omitempty"`

		ServiceSku string `json:"ServiceSku,omitempty"`

		ServiceStartDate strfmt.DateTime `json:"ServiceStartDate,omitempty"`

		WarrantyEndDate string `json:"WarrantyEndDate,omitempty"`

		WarrantyType string `json:"WarrantyType,omitempty"`
	}

	dataAO1.Contract = m.Contract

	dataAO1.ContractStatus = m.ContractStatus

	dataAO1.CoveredProductLineEndDate = m.CoveredProductLineEndDate

	dataAO1.DeviceID = m.DeviceID

	dataAO1.DeviceType = m.DeviceType

	dataAO1.EndCustomer = m.EndCustomer

	dataAO1.EndUserGlobalUltimate = m.EndUserGlobalUltimate

	dataAO1.IsValid = m.IsValid

	dataAO1.ItemType = m.ItemType

	dataAO1.MaintenancePurchaseOrderNumber = m.MaintenancePurchaseOrderNumber

	dataAO1.MaintenanceSalesOrderNumber = m.MaintenanceSalesOrderNumber

	dataAO1.PlatformType = m.PlatformType

	dataAO1.Product = m.Product

	dataAO1.PurchaseOrderNumber = m.PurchaseOrderNumber

	dataAO1.RegisteredDevice = m.RegisteredDevice

	dataAO1.ResellerGlobalUltimate = m.ResellerGlobalUltimate

	dataAO1.SalesOrderNumber = m.SalesOrderNumber

	dataAO1.ServiceDescription = m.ServiceDescription

	dataAO1.ServiceEndDate = m.ServiceEndDate

	dataAO1.ServiceLevel = m.ServiceLevel

	dataAO1.ServiceSku = m.ServiceSku

	dataAO1.ServiceStartDate = m.ServiceStartDate

	dataAO1.WarrantyEndDate = m.WarrantyEndDate

	dataAO1.WarrantyType = m.WarrantyType

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this asset device contract information
func (m *AssetDeviceContractInformation) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContract(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndUserGlobalUltimate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatformType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegisteredDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResellerGlobalUltimate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetDeviceContractInformation) validateContract(formats strfmt.Registry) error {

	if swag.IsZero(m.Contract) { // not required
		return nil
	}

	if m.Contract != nil {
		if err := m.Contract.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Contract")
			}
			return err
		}
	}

	return nil
}

var assetDeviceContractInformationTypeContractStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Not Covered","Active","Expiring Soon"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetDeviceContractInformationTypeContractStatusPropEnum = append(assetDeviceContractInformationTypeContractStatusPropEnum, v)
	}
}

// property enum
func (m *AssetDeviceContractInformation) validateContractStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetDeviceContractInformationTypeContractStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetDeviceContractInformation) validateContractStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ContractStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateContractStatusEnum("ContractStatus", "body", m.ContractStatus); err != nil {
		return err
	}

	return nil
}

var assetDeviceContractInformationTypeDeviceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["None","CiscoUcsServer","CiscoUcsFI","CiscoUcsChassis"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetDeviceContractInformationTypeDeviceTypePropEnum = append(assetDeviceContractInformationTypeDeviceTypePropEnum, v)
	}
}

// property enum
func (m *AssetDeviceContractInformation) validateDeviceTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetDeviceContractInformationTypeDeviceTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetDeviceContractInformation) validateDeviceType(formats strfmt.Registry) error {

	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateDeviceTypeEnum("DeviceType", "body", m.DeviceType); err != nil {
		return err
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateEndCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.EndCustomer) { // not required
		return nil
	}

	if m.EndCustomer != nil {
		if err := m.EndCustomer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndCustomer")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateEndUserGlobalUltimate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndUserGlobalUltimate) { // not required
		return nil
	}

	if m.EndUserGlobalUltimate != nil {
		if err := m.EndUserGlobalUltimate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("EndUserGlobalUltimate")
			}
			return err
		}
	}

	return nil
}

var assetDeviceContractInformationTypePlatformTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","APIC","DCNM","UCSFI","IMC","IMCM4","IMCM5","HX","HXTriton","UCSD","IntersightAppliance","PureStorage","VMware","ServiceEngine"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assetDeviceContractInformationTypePlatformTypePropEnum = append(assetDeviceContractInformationTypePlatformTypePropEnum, v)
	}
}

// property enum
func (m *AssetDeviceContractInformation) validatePlatformTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, assetDeviceContractInformationTypePlatformTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AssetDeviceContractInformation) validatePlatformType(formats strfmt.Registry) error {

	if swag.IsZero(m.PlatformType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformTypeEnum("PlatformType", "body", m.PlatformType); err != nil {
		return err
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateProduct(formats strfmt.Registry) error {

	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Product")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateRegisteredDevice(formats strfmt.Registry) error {

	if swag.IsZero(m.RegisteredDevice) { // not required
		return nil
	}

	if m.RegisteredDevice != nil {
		if err := m.RegisteredDevice.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegisteredDevice")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateResellerGlobalUltimate(formats strfmt.Registry) error {

	if swag.IsZero(m.ResellerGlobalUltimate) { // not required
		return nil
	}

	if m.ResellerGlobalUltimate != nil {
		if err := m.ResellerGlobalUltimate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResellerGlobalUltimate")
			}
			return err
		}
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateServiceEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ServiceEndDate", "body", "date-time", m.ServiceEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AssetDeviceContractInformation) validateServiceStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("ServiceStartDate", "body", "date-time", m.ServiceStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetDeviceContractInformation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetDeviceContractInformation) UnmarshalBinary(b []byte) error {
	var res AssetDeviceContractInformation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}