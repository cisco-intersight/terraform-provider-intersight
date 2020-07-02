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

// IamCertificateRequestAllOf Definition of the list of properties defined in 'iam.CertificateRequest', excluding properties defined in parent classes.
type IamCertificateRequestAllOf struct {
	// User input email address, an optional part of the subject of the certificate request.
	EmailAddress *string `json:"EmailAddress,omitempty"`
	// Name of the certificate request.
	Name *string `json:"Name,omitempty"`
	// Generated certificate signing request (CSR) in PEM format.
	Request *string `json:"Request,omitempty"`
	// Whether the user wants the generated CSR to be self-signed by the appliance.
	SelfSigned           *bool                          `json:"SelfSigned,omitempty"`
	Subject              *PkixDistinguishedName         `json:"Subject,omitempty"`
	SubjectAlternateName *PkixSubjectAlternateName      `json:"SubjectAlternateName,omitempty"`
	Account              *IamAccountRelationship        `json:"Account,omitempty"`
	Certificate          *IamCertificateRelationship    `json:"Certificate,omitempty"`
	PrivateKeySpec       *IamPrivateKeySpecRelationship `json:"PrivateKeySpec,omitempty"`
}

// NewIamCertificateRequestAllOf instantiates a new IamCertificateRequestAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamCertificateRequestAllOf() *IamCertificateRequestAllOf {
	this := IamCertificateRequestAllOf{}
	return &this
}

// NewIamCertificateRequestAllOfWithDefaults instantiates a new IamCertificateRequestAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamCertificateRequestAllOfWithDefaults() *IamCertificateRequestAllOf {
	this := IamCertificateRequestAllOf{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetEmailAddress() string {
	if o == nil || o.EmailAddress == nil {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetEmailAddressOk() (*string, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *IamCertificateRequestAllOf) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamCertificateRequestAllOf) SetName(v string) {
	o.Name = &v
}

// GetRequest returns the Request field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetRequest() string {
	if o == nil || o.Request == nil {
		var ret string
		return ret
	}
	return *o.Request
}

// GetRequestOk returns a tuple with the Request field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetRequestOk() (*string, bool) {
	if o == nil || o.Request == nil {
		return nil, false
	}
	return o.Request, true
}

// HasRequest returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasRequest() bool {
	if o != nil && o.Request != nil {
		return true
	}

	return false
}

// SetRequest gets a reference to the given string and assigns it to the Request field.
func (o *IamCertificateRequestAllOf) SetRequest(v string) {
	o.Request = &v
}

// GetSelfSigned returns the SelfSigned field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetSelfSigned() bool {
	if o == nil || o.SelfSigned == nil {
		var ret bool
		return ret
	}
	return *o.SelfSigned
}

// GetSelfSignedOk returns a tuple with the SelfSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetSelfSignedOk() (*bool, bool) {
	if o == nil || o.SelfSigned == nil {
		return nil, false
	}
	return o.SelfSigned, true
}

// HasSelfSigned returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasSelfSigned() bool {
	if o != nil && o.SelfSigned != nil {
		return true
	}

	return false
}

// SetSelfSigned gets a reference to the given bool and assigns it to the SelfSigned field.
func (o *IamCertificateRequestAllOf) SetSelfSigned(v bool) {
	o.SelfSigned = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetSubject() PkixDistinguishedName {
	if o == nil || o.Subject == nil {
		var ret PkixDistinguishedName
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetSubjectOk() (*PkixDistinguishedName, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given PkixDistinguishedName and assigns it to the Subject field.
func (o *IamCertificateRequestAllOf) SetSubject(v PkixDistinguishedName) {
	o.Subject = &v
}

// GetSubjectAlternateName returns the SubjectAlternateName field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetSubjectAlternateName() PkixSubjectAlternateName {
	if o == nil || o.SubjectAlternateName == nil {
		var ret PkixSubjectAlternateName
		return ret
	}
	return *o.SubjectAlternateName
}

// GetSubjectAlternateNameOk returns a tuple with the SubjectAlternateName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetSubjectAlternateNameOk() (*PkixSubjectAlternateName, bool) {
	if o == nil || o.SubjectAlternateName == nil {
		return nil, false
	}
	return o.SubjectAlternateName, true
}

// HasSubjectAlternateName returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasSubjectAlternateName() bool {
	if o != nil && o.SubjectAlternateName != nil {
		return true
	}

	return false
}

// SetSubjectAlternateName gets a reference to the given PkixSubjectAlternateName and assigns it to the SubjectAlternateName field.
func (o *IamCertificateRequestAllOf) SetSubjectAlternateName(v PkixSubjectAlternateName) {
	o.SubjectAlternateName = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamCertificateRequestAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetCertificate() IamCertificateRelationship {
	if o == nil || o.Certificate == nil {
		var ret IamCertificateRelationship
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetCertificateOk() (*IamCertificateRelationship, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given IamCertificateRelationship and assigns it to the Certificate field.
func (o *IamCertificateRequestAllOf) SetCertificate(v IamCertificateRelationship) {
	o.Certificate = &v
}

// GetPrivateKeySpec returns the PrivateKeySpec field value if set, zero value otherwise.
func (o *IamCertificateRequestAllOf) GetPrivateKeySpec() IamPrivateKeySpecRelationship {
	if o == nil || o.PrivateKeySpec == nil {
		var ret IamPrivateKeySpecRelationship
		return ret
	}
	return *o.PrivateKeySpec
}

// GetPrivateKeySpecOk returns a tuple with the PrivateKeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamCertificateRequestAllOf) GetPrivateKeySpecOk() (*IamPrivateKeySpecRelationship, bool) {
	if o == nil || o.PrivateKeySpec == nil {
		return nil, false
	}
	return o.PrivateKeySpec, true
}

// HasPrivateKeySpec returns a boolean if a field has been set.
func (o *IamCertificateRequestAllOf) HasPrivateKeySpec() bool {
	if o != nil && o.PrivateKeySpec != nil {
		return true
	}

	return false
}

// SetPrivateKeySpec gets a reference to the given IamPrivateKeySpecRelationship and assigns it to the PrivateKeySpec field.
func (o *IamCertificateRequestAllOf) SetPrivateKeySpec(v IamPrivateKeySpecRelationship) {
	o.PrivateKeySpec = &v
}

func (o IamCertificateRequestAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EmailAddress != nil {
		toSerialize["EmailAddress"] = o.EmailAddress
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Request != nil {
		toSerialize["Request"] = o.Request
	}
	if o.SelfSigned != nil {
		toSerialize["SelfSigned"] = o.SelfSigned
	}
	if o.Subject != nil {
		toSerialize["Subject"] = o.Subject
	}
	if o.SubjectAlternateName != nil {
		toSerialize["SubjectAlternateName"] = o.SubjectAlternateName
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}
	if o.Certificate != nil {
		toSerialize["Certificate"] = o.Certificate
	}
	if o.PrivateKeySpec != nil {
		toSerialize["PrivateKeySpec"] = o.PrivateKeySpec
	}
	return json.Marshal(toSerialize)
}

type NullableIamCertificateRequestAllOf struct {
	value *IamCertificateRequestAllOf
	isSet bool
}

func (v NullableIamCertificateRequestAllOf) Get() *IamCertificateRequestAllOf {
	return v.value
}

func (v *NullableIamCertificateRequestAllOf) Set(val *IamCertificateRequestAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamCertificateRequestAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamCertificateRequestAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamCertificateRequestAllOf(val *IamCertificateRequestAllOf) *NullableIamCertificateRequestAllOf {
	return &NullableIamCertificateRequestAllOf{value: val, isSet: true}
}

func (v NullableIamCertificateRequestAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamCertificateRequestAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}