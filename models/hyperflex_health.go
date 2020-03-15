// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HyperflexHealth Hyperflex:Health
//
// swagger:model hyperflexHealth
type HyperflexHealth struct {
	MoBaseMo

	// arbitration service state
	// Read Only: true
	// Enum: [NOT_AVAILABLE UNKNOWN ONLINE OFFLINE]
	ArbitrationServiceState string `json:"ArbitrationServiceState,omitempty"`

	// A collection of references to the [hyperflex.Cluster](mo://hyperflex.Cluster) Managed Object.
	// When this managed object is deleted, the referenced [hyperflex.Cluster](mo://hyperflex.Cluster) MO unsets its reference to this deleted MO.
	// Read Only: true
	Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

	// data replication compliance
	// Read Only: true
	// Enum: [UNKNOWN COMPLIANT NON_COMPLIANT]
	DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

	// resiliency details
	// Read Only: true
	ResiliencyDetails *HyperflexHxResiliencyInfoDt `json:"ResiliencyDetails,omitempty"`

	// state
	// Read Only: true
	// Enum: [UNKNOWN ONLINE OFFLINE ENOSPACE READONLY]
	State string `json:"State,omitempty"`

	// Uuid
	// Read Only: true
	UUID string `json:"Uuid,omitempty"`

	// zk health
	// Read Only: true
	// Enum: [NOT_AVAILABLE UNKNOWN ONLINE OFFLINE]
	ZkHealth string `json:"ZkHealth,omitempty"`

	// zone resiliency list
	// Read Only: true
	ZoneResiliencyList []*HyperflexHxZoneResiliencyInfoDt `json:"ZoneResiliencyList"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HyperflexHealth) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 MoBaseMo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.MoBaseMo = aO0

	// AO1
	var dataAO1 struct {
		ArbitrationServiceState string `json:"ArbitrationServiceState,omitempty"`

		Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

		DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

		ResiliencyDetails *HyperflexHxResiliencyInfoDt `json:"ResiliencyDetails,omitempty"`

		State string `json:"State,omitempty"`

		UUID string `json:"Uuid,omitempty"`

		ZkHealth string `json:"ZkHealth,omitempty"`

		ZoneResiliencyList []*HyperflexHxZoneResiliencyInfoDt `json:"ZoneResiliencyList"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.ArbitrationServiceState = dataAO1.ArbitrationServiceState

	m.Cluster = dataAO1.Cluster

	m.DataReplicationCompliance = dataAO1.DataReplicationCompliance

	m.ResiliencyDetails = dataAO1.ResiliencyDetails

	m.State = dataAO1.State

	m.UUID = dataAO1.UUID

	m.ZkHealth = dataAO1.ZkHealth

	m.ZoneResiliencyList = dataAO1.ZoneResiliencyList

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HyperflexHealth) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.MoBaseMo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		ArbitrationServiceState string `json:"ArbitrationServiceState,omitempty"`

		Cluster *HyperflexClusterRef `json:"Cluster,omitempty"`

		DataReplicationCompliance string `json:"DataReplicationCompliance,omitempty"`

		ResiliencyDetails *HyperflexHxResiliencyInfoDt `json:"ResiliencyDetails,omitempty"`

		State string `json:"State,omitempty"`

		UUID string `json:"Uuid,omitempty"`

		ZkHealth string `json:"ZkHealth,omitempty"`

		ZoneResiliencyList []*HyperflexHxZoneResiliencyInfoDt `json:"ZoneResiliencyList"`
	}

	dataAO1.ArbitrationServiceState = m.ArbitrationServiceState

	dataAO1.Cluster = m.Cluster

	dataAO1.DataReplicationCompliance = m.DataReplicationCompliance

	dataAO1.ResiliencyDetails = m.ResiliencyDetails

	dataAO1.State = m.State

	dataAO1.UUID = m.UUID

	dataAO1.ZkHealth = m.ZkHealth

	dataAO1.ZoneResiliencyList = m.ZoneResiliencyList

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this hyperflex health
func (m *HyperflexHealth) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with MoBaseMo
	if err := m.MoBaseMo.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateArbitrationServiceState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataReplicationCompliance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResiliencyDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZkHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneResiliencyList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var hyperflexHealthTypeArbitrationServiceStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_AVAILABLE","UNKNOWN","ONLINE","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHealthTypeArbitrationServiceStatePropEnum = append(hyperflexHealthTypeArbitrationServiceStatePropEnum, v)
	}
}

// property enum
func (m *HyperflexHealth) validateArbitrationServiceStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHealthTypeArbitrationServiceStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHealth) validateArbitrationServiceState(formats strfmt.Registry) error {

	if swag.IsZero(m.ArbitrationServiceState) { // not required
		return nil
	}

	// value enum
	if err := m.validateArbitrationServiceStateEnum("ArbitrationServiceState", "body", m.ArbitrationServiceState); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexHealth) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

var hyperflexHealthTypeDataReplicationCompliancePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","COMPLIANT","NON_COMPLIANT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHealthTypeDataReplicationCompliancePropEnum = append(hyperflexHealthTypeDataReplicationCompliancePropEnum, v)
	}
}

// property enum
func (m *HyperflexHealth) validateDataReplicationComplianceEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHealthTypeDataReplicationCompliancePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHealth) validateDataReplicationCompliance(formats strfmt.Registry) error {

	if swag.IsZero(m.DataReplicationCompliance) { // not required
		return nil
	}

	// value enum
	if err := m.validateDataReplicationComplianceEnum("DataReplicationCompliance", "body", m.DataReplicationCompliance); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexHealth) validateResiliencyDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.ResiliencyDetails) { // not required
		return nil
	}

	if m.ResiliencyDetails != nil {
		if err := m.ResiliencyDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ResiliencyDetails")
			}
			return err
		}
	}

	return nil
}

var hyperflexHealthTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","ONLINE","OFFLINE","ENOSPACE","READONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHealthTypeStatePropEnum = append(hyperflexHealthTypeStatePropEnum, v)
	}
}

// property enum
func (m *HyperflexHealth) validateStateEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHealthTypeStatePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHealth) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("State", "body", m.State); err != nil {
		return err
	}

	return nil
}

var hyperflexHealthTypeZkHealthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NOT_AVAILABLE","UNKNOWN","ONLINE","OFFLINE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		hyperflexHealthTypeZkHealthPropEnum = append(hyperflexHealthTypeZkHealthPropEnum, v)
	}
}

// property enum
func (m *HyperflexHealth) validateZkHealthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, hyperflexHealthTypeZkHealthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *HyperflexHealth) validateZkHealth(formats strfmt.Registry) error {

	if swag.IsZero(m.ZkHealth) { // not required
		return nil
	}

	// value enum
	if err := m.validateZkHealthEnum("ZkHealth", "body", m.ZkHealth); err != nil {
		return err
	}

	return nil
}

func (m *HyperflexHealth) validateZoneResiliencyList(formats strfmt.Registry) error {

	if swag.IsZero(m.ZoneResiliencyList) { // not required
		return nil
	}

	for i := 0; i < len(m.ZoneResiliencyList); i++ {
		if swag.IsZero(m.ZoneResiliencyList[i]) { // not required
			continue
		}

		if m.ZoneResiliencyList[i] != nil {
			if err := m.ZoneResiliencyList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ZoneResiliencyList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HyperflexHealth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HyperflexHealth) UnmarshalBinary(b []byte) error {
	var res HyperflexHealth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
