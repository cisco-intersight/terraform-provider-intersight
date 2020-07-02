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
	"fmt"
)

// StorageDiskGroupPolicyRelationship - A relationship to the 'storage.DiskGroupPolicy' resource, or the expanded 'storage.DiskGroupPolicy' resource, or the 'null' value.
type StorageDiskGroupPolicyRelationship struct {
	MoMoRef                *MoMoRef
	StorageDiskGroupPolicy *StorageDiskGroupPolicy
}

// MoMoRefAsStorageDiskGroupPolicyRelationship is a convenience function that returns MoMoRef wrapped in StorageDiskGroupPolicyRelationship
func MoMoRefAsStorageDiskGroupPolicyRelationship(v *MoMoRef) StorageDiskGroupPolicyRelationship {
	return StorageDiskGroupPolicyRelationship{MoMoRef: v}
}

// StorageDiskGroupPolicyAsStorageDiskGroupPolicyRelationship is a convenience function that returns StorageDiskGroupPolicy wrapped in StorageDiskGroupPolicyRelationship
func StorageDiskGroupPolicyAsStorageDiskGroupPolicyRelationship(v *StorageDiskGroupPolicy) StorageDiskGroupPolicyRelationship {
	return StorageDiskGroupPolicyRelationship{StorageDiskGroupPolicy: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StorageDiskGroupPolicyRelationship) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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

	// check if the discriminator value is 'mo.MoRef'
	if jsonDict["ClassId"] == "mo.MoRef" {
		// try to unmarshal JSON data into MoMoRef
		err = json.Unmarshal(data, &dst.MoMoRef)
		if err == nil {
			jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
			if string(jsonMoMoRef) == "{}" { // empty struct
				dst.MoMoRef = nil
			} else {
				return nil // data stored in dst.MoMoRef, return on the first match
			}
		} else {
			dst.MoMoRef = nil
		}
	}

	// check if the discriminator value is 'storage.DiskGroupPolicy'
	if jsonDict["ClassId"] == "storage.DiskGroupPolicy" {
		// try to unmarshal JSON data into StorageDiskGroupPolicy
		err = json.Unmarshal(data, &dst.StorageDiskGroupPolicy)
		if err == nil {
			jsonStorageDiskGroupPolicy, _ := json.Marshal(dst.StorageDiskGroupPolicy)
			if string(jsonStorageDiskGroupPolicy) == "{}" { // empty struct
				dst.StorageDiskGroupPolicy = nil
			} else {
				return nil // data stored in dst.StorageDiskGroupPolicy, return on the first match
			}
		} else {
			dst.StorageDiskGroupPolicy = nil
		}
	}

	// try to unmarshal data into MoMoRef
	err = json.Unmarshal(data, &dst.MoMoRef)
	if err == nil {
		jsonMoMoRef, _ := json.Marshal(dst.MoMoRef)
		if string(jsonMoMoRef) == "{}" { // empty struct
			dst.MoMoRef = nil
		} else {
			match++
		}
	} else {
		dst.MoMoRef = nil
	}

	// try to unmarshal data into StorageDiskGroupPolicy
	err = json.Unmarshal(data, &dst.StorageDiskGroupPolicy)
	if err == nil {
		jsonStorageDiskGroupPolicy, _ := json.Marshal(dst.StorageDiskGroupPolicy)
		if string(jsonStorageDiskGroupPolicy) == "{}" { // empty struct
			dst.StorageDiskGroupPolicy = nil
		} else {
			match++
		}
	} else {
		dst.StorageDiskGroupPolicy = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MoMoRef = nil
		dst.StorageDiskGroupPolicy = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(StorageDiskGroupPolicyRelationship)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(StorageDiskGroupPolicyRelationship)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StorageDiskGroupPolicyRelationship) MarshalJSON() ([]byte, error) {
	if src.MoMoRef != nil {
		return json.Marshal(&src.MoMoRef)
	}

	if src.StorageDiskGroupPolicy != nil {
		return json.Marshal(&src.StorageDiskGroupPolicy)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StorageDiskGroupPolicyRelationship) GetActualInstance() interface{} {
	if obj.MoMoRef != nil {
		return obj.MoMoRef
	}

	if obj.StorageDiskGroupPolicy != nil {
		return obj.StorageDiskGroupPolicy
	}

	// all schemas are nil
	return nil
}

type NullableStorageDiskGroupPolicyRelationship struct {
	value *StorageDiskGroupPolicyRelationship
	isSet bool
}

func (v NullableStorageDiskGroupPolicyRelationship) Get() *StorageDiskGroupPolicyRelationship {
	return v.value
}

func (v *NullableStorageDiskGroupPolicyRelationship) Set(val *StorageDiskGroupPolicyRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageDiskGroupPolicyRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageDiskGroupPolicyRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageDiskGroupPolicyRelationship(val *StorageDiskGroupPolicyRelationship) *NullableStorageDiskGroupPolicyRelationship {
	return &NullableStorageDiskGroupPolicyRelationship{value: val, isSet: true}
}

func (v NullableStorageDiskGroupPolicyRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageDiskGroupPolicyRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}