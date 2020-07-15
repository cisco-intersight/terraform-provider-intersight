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
	"fmt"
)

// BootDeviceBootModeRelationship - A relationship to the 'boot.DeviceBootMode' resource, or the expanded 'boot.DeviceBootMode' resource, or the 'null' value.
type BootDeviceBootModeRelationship struct {
	BootDeviceBootMode *BootDeviceBootMode
	MoMoRef            *MoMoRef
}

// BootDeviceBootModeAsBootDeviceBootModeRelationship is a convenience function that returns BootDeviceBootMode wrapped in BootDeviceBootModeRelationship
func BootDeviceBootModeAsBootDeviceBootModeRelationship(v *BootDeviceBootMode) BootDeviceBootModeRelationship {
	return BootDeviceBootModeRelationship{BootDeviceBootMode: v}
}

// MoMoRefAsBootDeviceBootModeRelationship is a convenience function that returns MoMoRef wrapped in BootDeviceBootModeRelationship
func MoMoRefAsBootDeviceBootModeRelationship(v *MoMoRef) BootDeviceBootModeRelationship {
	return BootDeviceBootModeRelationship{MoMoRef: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BootDeviceBootModeRelationship) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'boot.DeviceBootMode'
	if jsonDict["ClassId"] == "boot.DeviceBootMode" {
		// try to unmarshal JSON data into BootDeviceBootMode
		err = json.Unmarshal(data, &dst.BootDeviceBootMode)
		if err == nil {
			return nil // data stored in dst.BootDeviceBootMode, return on the first match
		} else {
			dst.BootDeviceBootMode = nil
			return fmt.Errorf("Failed to unmarshal BootDeviceBootModeRelationship as BootDeviceBootMode: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			return nil // data stored in dst.MoMoRef, return on the first match
		} else {
			dst.MoMoRef = nil
			return fmt.Errorf("Failed to unmarshal BootDeviceBootModeRelationship as MoMoRef: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BootDeviceBootModeRelationship) MarshalJSON() ([]byte, error) {
	if src.BootDeviceBootMode != nil {
		return json.Marshal(&src.BootDeviceBootMode)
	}

	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BootDeviceBootModeRelationship) GetActualInstance() interface{} {
	if obj.BootDeviceBootMode != nil {
		return obj.BootDeviceBootMode
	}

	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	// all schemas are nil
	return nil
}

type NullableBootDeviceBootModeRelationship struct {
	value *BootDeviceBootModeRelationship
	isSet bool
}

func (v NullableBootDeviceBootModeRelationship) Get() *BootDeviceBootModeRelationship {
	return v.value
}

func (v *NullableBootDeviceBootModeRelationship) Set(val *BootDeviceBootModeRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableBootDeviceBootModeRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableBootDeviceBootModeRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBootDeviceBootModeRelationship(val *BootDeviceBootModeRelationship) *NullableBootDeviceBootModeRelationship {
	return &NullableBootDeviceBootModeRelationship{value: val, isSet: true}
}

func (v NullableBootDeviceBootModeRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBootDeviceBootModeRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
