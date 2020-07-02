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
	"time"
)

// UcsdBackupInfoAllOf Definition of the list of properties defined in 'ucsd.BackupInfo', excluding properties defined in parent classes.
type UcsdBackupInfoAllOf struct {
	// Auto generated backup File Name with combination of file prefix given an user input and the timestamp.
	BackupFileName *string `json:"BackupFileName,omitempty"`
	// Backup location that contains the backup images for end device which can be used for restore operation.
	BackupLocation *string `json:"BackupLocation,omitempty"`
	// Backup server where backup images are maintained.
	BackupServerIp *string `json:"BackupServerIp,omitempty"`
	// Size of the backup image in bytes.
	BackupSize *int64               `json:"BackupSize,omitempty"`
	Connectors *[]UcsdConnectorPack `json:"Connectors,omitempty"`
	// Time taken for the backup to be completed.
	Duration *int64 `json:"Duration,omitempty"`
	// The key used for encrypting the backup file.
	EncryptionKey *string `json:"EncryptionKey,omitempty"`
	// Reason for backup failure.
	FailureReason *string `json:"FailureReason,omitempty"`
	// Backup image got purged or not. The backup images get purged based on the retention count set by the user in the backup config policy.
	IsPurged *bool `json:"IsPurged,omitempty"`
	// Last modified time when this backup record got updated.
	LastModified *time.Time `json:"LastModified,omitempty"`
	// Backup current precentage completion status information.
	PercentageCompletion *int64 `json:"PercentageCompletion,omitempty"`
	// The end device product version when the backup image was taken.
	ProductVersion *string `json:"ProductVersion,omitempty"`
	// Protocol used for the remote backup. possible values are FTP, SCP and SFTP. Not applicable for the localhost (127.0.0.1).
	Protocol *string `json:"Protocol,omitempty"`
	// Backup current status stage information.
	StageCompletion *string `json:"StageCompletion,omitempty"`
	// Start time of backup when it got initiated.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Current status of Backup current.
	Status *string `json:"Status,omitempty"`
}

// NewUcsdBackupInfoAllOf instantiates a new UcsdBackupInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUcsdBackupInfoAllOf() *UcsdBackupInfoAllOf {
	this := UcsdBackupInfoAllOf{}
	return &this
}

// NewUcsdBackupInfoAllOfWithDefaults instantiates a new UcsdBackupInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUcsdBackupInfoAllOfWithDefaults() *UcsdBackupInfoAllOf {
	this := UcsdBackupInfoAllOf{}
	return &this
}

// GetBackupFileName returns the BackupFileName field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetBackupFileName() string {
	if o == nil || o.BackupFileName == nil {
		var ret string
		return ret
	}
	return *o.BackupFileName
}

// GetBackupFileNameOk returns a tuple with the BackupFileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetBackupFileNameOk() (*string, bool) {
	if o == nil || o.BackupFileName == nil {
		return nil, false
	}
	return o.BackupFileName, true
}

// HasBackupFileName returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasBackupFileName() bool {
	if o != nil && o.BackupFileName != nil {
		return true
	}

	return false
}

// SetBackupFileName gets a reference to the given string and assigns it to the BackupFileName field.
func (o *UcsdBackupInfoAllOf) SetBackupFileName(v string) {
	o.BackupFileName = &v
}

// GetBackupLocation returns the BackupLocation field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetBackupLocation() string {
	if o == nil || o.BackupLocation == nil {
		var ret string
		return ret
	}
	return *o.BackupLocation
}

// GetBackupLocationOk returns a tuple with the BackupLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetBackupLocationOk() (*string, bool) {
	if o == nil || o.BackupLocation == nil {
		return nil, false
	}
	return o.BackupLocation, true
}

// HasBackupLocation returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasBackupLocation() bool {
	if o != nil && o.BackupLocation != nil {
		return true
	}

	return false
}

// SetBackupLocation gets a reference to the given string and assigns it to the BackupLocation field.
func (o *UcsdBackupInfoAllOf) SetBackupLocation(v string) {
	o.BackupLocation = &v
}

// GetBackupServerIp returns the BackupServerIp field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetBackupServerIp() string {
	if o == nil || o.BackupServerIp == nil {
		var ret string
		return ret
	}
	return *o.BackupServerIp
}

// GetBackupServerIpOk returns a tuple with the BackupServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetBackupServerIpOk() (*string, bool) {
	if o == nil || o.BackupServerIp == nil {
		return nil, false
	}
	return o.BackupServerIp, true
}

// HasBackupServerIp returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasBackupServerIp() bool {
	if o != nil && o.BackupServerIp != nil {
		return true
	}

	return false
}

// SetBackupServerIp gets a reference to the given string and assigns it to the BackupServerIp field.
func (o *UcsdBackupInfoAllOf) SetBackupServerIp(v string) {
	o.BackupServerIp = &v
}

// GetBackupSize returns the BackupSize field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetBackupSize() int64 {
	if o == nil || o.BackupSize == nil {
		var ret int64
		return ret
	}
	return *o.BackupSize
}

// GetBackupSizeOk returns a tuple with the BackupSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetBackupSizeOk() (*int64, bool) {
	if o == nil || o.BackupSize == nil {
		return nil, false
	}
	return o.BackupSize, true
}

// HasBackupSize returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasBackupSize() bool {
	if o != nil && o.BackupSize != nil {
		return true
	}

	return false
}

// SetBackupSize gets a reference to the given int64 and assigns it to the BackupSize field.
func (o *UcsdBackupInfoAllOf) SetBackupSize(v int64) {
	o.BackupSize = &v
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetConnectors() []UcsdConnectorPack {
	if o == nil || o.Connectors == nil {
		var ret []UcsdConnectorPack
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetConnectorsOk() (*[]UcsdConnectorPack, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []UcsdConnectorPack and assigns it to the Connectors field.
func (o *UcsdBackupInfoAllOf) SetConnectors(v []UcsdConnectorPack) {
	o.Connectors = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetDuration() int64 {
	if o == nil || o.Duration == nil {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetDurationOk() (*int64, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *UcsdBackupInfoAllOf) SetDuration(v int64) {
	o.Duration = &v
}

// GetEncryptionKey returns the EncryptionKey field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetEncryptionKey() string {
	if o == nil || o.EncryptionKey == nil {
		var ret string
		return ret
	}
	return *o.EncryptionKey
}

// GetEncryptionKeyOk returns a tuple with the EncryptionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetEncryptionKeyOk() (*string, bool) {
	if o == nil || o.EncryptionKey == nil {
		return nil, false
	}
	return o.EncryptionKey, true
}

// HasEncryptionKey returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasEncryptionKey() bool {
	if o != nil && o.EncryptionKey != nil {
		return true
	}

	return false
}

// SetEncryptionKey gets a reference to the given string and assigns it to the EncryptionKey field.
func (o *UcsdBackupInfoAllOf) SetEncryptionKey(v string) {
	o.EncryptionKey = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *UcsdBackupInfoAllOf) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetIsPurged returns the IsPurged field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetIsPurged() bool {
	if o == nil || o.IsPurged == nil {
		var ret bool
		return ret
	}
	return *o.IsPurged
}

// GetIsPurgedOk returns a tuple with the IsPurged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetIsPurgedOk() (*bool, bool) {
	if o == nil || o.IsPurged == nil {
		return nil, false
	}
	return o.IsPurged, true
}

// HasIsPurged returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasIsPurged() bool {
	if o != nil && o.IsPurged != nil {
		return true
	}

	return false
}

// SetIsPurged gets a reference to the given bool and assigns it to the IsPurged field.
func (o *UcsdBackupInfoAllOf) SetIsPurged(v bool) {
	o.IsPurged = &v
}

// GetLastModified returns the LastModified field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetLastModified() time.Time {
	if o == nil || o.LastModified == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil || o.LastModified == nil {
		return nil, false
	}
	return o.LastModified, true
}

// HasLastModified returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasLastModified() bool {
	if o != nil && o.LastModified != nil {
		return true
	}

	return false
}

// SetLastModified gets a reference to the given time.Time and assigns it to the LastModified field.
func (o *UcsdBackupInfoAllOf) SetLastModified(v time.Time) {
	o.LastModified = &v
}

// GetPercentageCompletion returns the PercentageCompletion field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetPercentageCompletion() int64 {
	if o == nil || o.PercentageCompletion == nil {
		var ret int64
		return ret
	}
	return *o.PercentageCompletion
}

// GetPercentageCompletionOk returns a tuple with the PercentageCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetPercentageCompletionOk() (*int64, bool) {
	if o == nil || o.PercentageCompletion == nil {
		return nil, false
	}
	return o.PercentageCompletion, true
}

// HasPercentageCompletion returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasPercentageCompletion() bool {
	if o != nil && o.PercentageCompletion != nil {
		return true
	}

	return false
}

// SetPercentageCompletion gets a reference to the given int64 and assigns it to the PercentageCompletion field.
func (o *UcsdBackupInfoAllOf) SetPercentageCompletion(v int64) {
	o.PercentageCompletion = &v
}

// GetProductVersion returns the ProductVersion field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetProductVersion() string {
	if o == nil || o.ProductVersion == nil {
		var ret string
		return ret
	}
	return *o.ProductVersion
}

// GetProductVersionOk returns a tuple with the ProductVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetProductVersionOk() (*string, bool) {
	if o == nil || o.ProductVersion == nil {
		return nil, false
	}
	return o.ProductVersion, true
}

// HasProductVersion returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasProductVersion() bool {
	if o != nil && o.ProductVersion != nil {
		return true
	}

	return false
}

// SetProductVersion gets a reference to the given string and assigns it to the ProductVersion field.
func (o *UcsdBackupInfoAllOf) SetProductVersion(v string) {
	o.ProductVersion = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *UcsdBackupInfoAllOf) SetProtocol(v string) {
	o.Protocol = &v
}

// GetStageCompletion returns the StageCompletion field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetStageCompletion() string {
	if o == nil || o.StageCompletion == nil {
		var ret string
		return ret
	}
	return *o.StageCompletion
}

// GetStageCompletionOk returns a tuple with the StageCompletion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetStageCompletionOk() (*string, bool) {
	if o == nil || o.StageCompletion == nil {
		return nil, false
	}
	return o.StageCompletion, true
}

// HasStageCompletion returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasStageCompletion() bool {
	if o != nil && o.StageCompletion != nil {
		return true
	}

	return false
}

// SetStageCompletion gets a reference to the given string and assigns it to the StageCompletion field.
func (o *UcsdBackupInfoAllOf) SetStageCompletion(v string) {
	o.StageCompletion = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *UcsdBackupInfoAllOf) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UcsdBackupInfoAllOf) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UcsdBackupInfoAllOf) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UcsdBackupInfoAllOf) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UcsdBackupInfoAllOf) SetStatus(v string) {
	o.Status = &v
}

func (o UcsdBackupInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BackupFileName != nil {
		toSerialize["BackupFileName"] = o.BackupFileName
	}
	if o.BackupLocation != nil {
		toSerialize["BackupLocation"] = o.BackupLocation
	}
	if o.BackupServerIp != nil {
		toSerialize["BackupServerIp"] = o.BackupServerIp
	}
	if o.BackupSize != nil {
		toSerialize["BackupSize"] = o.BackupSize
	}
	if o.Connectors != nil {
		toSerialize["Connectors"] = o.Connectors
	}
	if o.Duration != nil {
		toSerialize["Duration"] = o.Duration
	}
	if o.EncryptionKey != nil {
		toSerialize["EncryptionKey"] = o.EncryptionKey
	}
	if o.FailureReason != nil {
		toSerialize["FailureReason"] = o.FailureReason
	}
	if o.IsPurged != nil {
		toSerialize["IsPurged"] = o.IsPurged
	}
	if o.LastModified != nil {
		toSerialize["LastModified"] = o.LastModified
	}
	if o.PercentageCompletion != nil {
		toSerialize["PercentageCompletion"] = o.PercentageCompletion
	}
	if o.ProductVersion != nil {
		toSerialize["ProductVersion"] = o.ProductVersion
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.StageCompletion != nil {
		toSerialize["StageCompletion"] = o.StageCompletion
	}
	if o.StartTime != nil {
		toSerialize["StartTime"] = o.StartTime
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableUcsdBackupInfoAllOf struct {
	value *UcsdBackupInfoAllOf
	isSet bool
}

func (v NullableUcsdBackupInfoAllOf) Get() *UcsdBackupInfoAllOf {
	return v.value
}

func (v *NullableUcsdBackupInfoAllOf) Set(val *UcsdBackupInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUcsdBackupInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUcsdBackupInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcsdBackupInfoAllOf(val *UcsdBackupInfoAllOf) *NullableUcsdBackupInfoAllOf {
	return &NullableUcsdBackupInfoAllOf{value: val, isSet: true}
}

func (v NullableUcsdBackupInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcsdBackupInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}