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
)

// ComputePhysicalAllOf Definition of the list of properties defined in 'compute.Physical', excluding properties defined in parent classes.
type ComputePhysicalAllOf struct {
	// The desired power state of the server.
	AdminPowerState *string `json:"AdminPowerState,omitempty"`
	// The user defined asset tag assigned to the server.
	AssetTag *string `json:"AssetTag,omitempty"`
	// The amount of memory available on the server.
	AvailableMemory *int64 `json:"AvailableMemory,omitempty"`
	// The BIOS POST completion status of the server.
	BiosPostComplete *bool `json:"BiosPostComplete,omitempty"`
	// The fault summary for the server.
	FaultSummary   *int64              `json:"FaultSummary,omitempty"`
	KvmIpAddresses *[]ComputeIpAddress `json:"KvmIpAddresses,omitempty"`
	// The management mode of the server.
	ManagementMode *string `json:"ManagementMode,omitempty"`
	// The maximum memory speed in MHz available on the server.
	MemorySpeed *string `json:"MemorySpeed,omitempty"`
	// Management address of the server.
	MgmtIpAddress *string `json:"MgmtIpAddress,omitempty"`
	// The total number of network adapters present on the server.
	NumAdaptors *int64 `json:"NumAdaptors,omitempty"`
	// The total number of CPU cores present on the server.
	NumCpuCores *int64 `json:"NumCpuCores,omitempty"`
	// The total number of CPU cores enabled on the server.
	NumCpuCoresEnabled *int64 `json:"NumCpuCoresEnabled,omitempty"`
	// The total number of CPUs present on the server.
	NumCpus *int64 `json:"NumCpus,omitempty"`
	// The total number of vNICs which are visible to a host on the server.
	NumEthHostInterfaces *int64 `json:"NumEthHostInterfaces,omitempty"`
	// The total number of vHBAs which are visible to a host on the server.
	NumFcHostInterfaces *int64 `json:"NumFcHostInterfaces,omitempty"`
	// The total number of threads the server is capable of handling.
	NumThreads *int64 `json:"NumThreads,omitempty"`
	// The actual power state of the server.
	OperPowerState *string `json:"OperPowerState,omitempty"`
	// The operational state of the server.
	OperState *string `json:"OperState,omitempty"`
	// The operability of the server.
	Operability *string `json:"Operability,omitempty"`
	// The platform type of the registered device - whether managed by UCSM or operating in standalone mode.
	PlatformType *string `json:"PlatformType,omitempty"`
	// Indicates if a server is present in a slot and is applicable for blade servers.
	Presence *string `json:"Presence,omitempty"`
	// The distinguished name of the service profile to which the server is associated to. It is applicable only for servers which are managed via UCSM.
	ServiceProfile *string `json:"ServiceProfile,omitempty"`
	// The total memory available on the server.
	TotalMemory *int64 `json:"TotalMemory,omitempty"`
	// The user defined label assigned to the server.
	UserLabel *string `json:"UserLabel,omitempty"`
	// The universally unique identity of the server.
	Uuid                 *string                                `json:"Uuid,omitempty"`
	MgmtIdentity         *EquipmentPhysicalIdentityRelationship `json:"MgmtIdentity,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ComputePhysicalAllOf ComputePhysicalAllOf

// NewComputePhysicalAllOf instantiates a new ComputePhysicalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputePhysicalAllOf() *ComputePhysicalAllOf {
	this := ComputePhysicalAllOf{}
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// NewComputePhysicalAllOfWithDefaults instantiates a new ComputePhysicalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputePhysicalAllOfWithDefaults() *ComputePhysicalAllOf {
	this := ComputePhysicalAllOf{}
	var managementMode string = "IntersightStandalone"
	this.ManagementMode = &managementMode
	return &this
}

// GetAdminPowerState returns the AdminPowerState field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetAdminPowerState() string {
	if o == nil || o.AdminPowerState == nil {
		var ret string
		return ret
	}
	return *o.AdminPowerState
}

// GetAdminPowerStateOk returns a tuple with the AdminPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetAdminPowerStateOk() (*string, bool) {
	if o == nil || o.AdminPowerState == nil {
		return nil, false
	}
	return o.AdminPowerState, true
}

// HasAdminPowerState returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasAdminPowerState() bool {
	if o != nil && o.AdminPowerState != nil {
		return true
	}

	return false
}

// SetAdminPowerState gets a reference to the given string and assigns it to the AdminPowerState field.
func (o *ComputePhysicalAllOf) SetAdminPowerState(v string) {
	o.AdminPowerState = &v
}

// GetAssetTag returns the AssetTag field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetAssetTag() string {
	if o == nil || o.AssetTag == nil {
		var ret string
		return ret
	}
	return *o.AssetTag
}

// GetAssetTagOk returns a tuple with the AssetTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetAssetTagOk() (*string, bool) {
	if o == nil || o.AssetTag == nil {
		return nil, false
	}
	return o.AssetTag, true
}

// HasAssetTag returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasAssetTag() bool {
	if o != nil && o.AssetTag != nil {
		return true
	}

	return false
}

// SetAssetTag gets a reference to the given string and assigns it to the AssetTag field.
func (o *ComputePhysicalAllOf) SetAssetTag(v string) {
	o.AssetTag = &v
}

// GetAvailableMemory returns the AvailableMemory field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetAvailableMemory() int64 {
	if o == nil || o.AvailableMemory == nil {
		var ret int64
		return ret
	}
	return *o.AvailableMemory
}

// GetAvailableMemoryOk returns a tuple with the AvailableMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetAvailableMemoryOk() (*int64, bool) {
	if o == nil || o.AvailableMemory == nil {
		return nil, false
	}
	return o.AvailableMemory, true
}

// HasAvailableMemory returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasAvailableMemory() bool {
	if o != nil && o.AvailableMemory != nil {
		return true
	}

	return false
}

// SetAvailableMemory gets a reference to the given int64 and assigns it to the AvailableMemory field.
func (o *ComputePhysicalAllOf) SetAvailableMemory(v int64) {
	o.AvailableMemory = &v
}

// GetBiosPostComplete returns the BiosPostComplete field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetBiosPostComplete() bool {
	if o == nil || o.BiosPostComplete == nil {
		var ret bool
		return ret
	}
	return *o.BiosPostComplete
}

// GetBiosPostCompleteOk returns a tuple with the BiosPostComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetBiosPostCompleteOk() (*bool, bool) {
	if o == nil || o.BiosPostComplete == nil {
		return nil, false
	}
	return o.BiosPostComplete, true
}

// HasBiosPostComplete returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasBiosPostComplete() bool {
	if o != nil && o.BiosPostComplete != nil {
		return true
	}

	return false
}

// SetBiosPostComplete gets a reference to the given bool and assigns it to the BiosPostComplete field.
func (o *ComputePhysicalAllOf) SetBiosPostComplete(v bool) {
	o.BiosPostComplete = &v
}

// GetFaultSummary returns the FaultSummary field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetFaultSummary() int64 {
	if o == nil || o.FaultSummary == nil {
		var ret int64
		return ret
	}
	return *o.FaultSummary
}

// GetFaultSummaryOk returns a tuple with the FaultSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetFaultSummaryOk() (*int64, bool) {
	if o == nil || o.FaultSummary == nil {
		return nil, false
	}
	return o.FaultSummary, true
}

// HasFaultSummary returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasFaultSummary() bool {
	if o != nil && o.FaultSummary != nil {
		return true
	}

	return false
}

// SetFaultSummary gets a reference to the given int64 and assigns it to the FaultSummary field.
func (o *ComputePhysicalAllOf) SetFaultSummary(v int64) {
	o.FaultSummary = &v
}

// GetKvmIpAddresses returns the KvmIpAddresses field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetKvmIpAddresses() []ComputeIpAddress {
	if o == nil || o.KvmIpAddresses == nil {
		var ret []ComputeIpAddress
		return ret
	}
	return *o.KvmIpAddresses
}

// GetKvmIpAddressesOk returns a tuple with the KvmIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetKvmIpAddressesOk() (*[]ComputeIpAddress, bool) {
	if o == nil || o.KvmIpAddresses == nil {
		return nil, false
	}
	return o.KvmIpAddresses, true
}

// HasKvmIpAddresses returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasKvmIpAddresses() bool {
	if o != nil && o.KvmIpAddresses != nil {
		return true
	}

	return false
}

// SetKvmIpAddresses gets a reference to the given []ComputeIpAddress and assigns it to the KvmIpAddresses field.
func (o *ComputePhysicalAllOf) SetKvmIpAddresses(v []ComputeIpAddress) {
	o.KvmIpAddresses = &v
}

// GetManagementMode returns the ManagementMode field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetManagementMode() string {
	if o == nil || o.ManagementMode == nil {
		var ret string
		return ret
	}
	return *o.ManagementMode
}

// GetManagementModeOk returns a tuple with the ManagementMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetManagementModeOk() (*string, bool) {
	if o == nil || o.ManagementMode == nil {
		return nil, false
	}
	return o.ManagementMode, true
}

// HasManagementMode returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasManagementMode() bool {
	if o != nil && o.ManagementMode != nil {
		return true
	}

	return false
}

// SetManagementMode gets a reference to the given string and assigns it to the ManagementMode field.
func (o *ComputePhysicalAllOf) SetManagementMode(v string) {
	o.ManagementMode = &v
}

// GetMemorySpeed returns the MemorySpeed field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetMemorySpeed() string {
	if o == nil || o.MemorySpeed == nil {
		var ret string
		return ret
	}
	return *o.MemorySpeed
}

// GetMemorySpeedOk returns a tuple with the MemorySpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetMemorySpeedOk() (*string, bool) {
	if o == nil || o.MemorySpeed == nil {
		return nil, false
	}
	return o.MemorySpeed, true
}

// HasMemorySpeed returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasMemorySpeed() bool {
	if o != nil && o.MemorySpeed != nil {
		return true
	}

	return false
}

// SetMemorySpeed gets a reference to the given string and assigns it to the MemorySpeed field.
func (o *ComputePhysicalAllOf) SetMemorySpeed(v string) {
	o.MemorySpeed = &v
}

// GetMgmtIpAddress returns the MgmtIpAddress field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetMgmtIpAddress() string {
	if o == nil || o.MgmtIpAddress == nil {
		var ret string
		return ret
	}
	return *o.MgmtIpAddress
}

// GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetMgmtIpAddressOk() (*string, bool) {
	if o == nil || o.MgmtIpAddress == nil {
		return nil, false
	}
	return o.MgmtIpAddress, true
}

// HasMgmtIpAddress returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasMgmtIpAddress() bool {
	if o != nil && o.MgmtIpAddress != nil {
		return true
	}

	return false
}

// SetMgmtIpAddress gets a reference to the given string and assigns it to the MgmtIpAddress field.
func (o *ComputePhysicalAllOf) SetMgmtIpAddress(v string) {
	o.MgmtIpAddress = &v
}

// GetNumAdaptors returns the NumAdaptors field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumAdaptors() int64 {
	if o == nil || o.NumAdaptors == nil {
		var ret int64
		return ret
	}
	return *o.NumAdaptors
}

// GetNumAdaptorsOk returns a tuple with the NumAdaptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumAdaptorsOk() (*int64, bool) {
	if o == nil || o.NumAdaptors == nil {
		return nil, false
	}
	return o.NumAdaptors, true
}

// HasNumAdaptors returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumAdaptors() bool {
	if o != nil && o.NumAdaptors != nil {
		return true
	}

	return false
}

// SetNumAdaptors gets a reference to the given int64 and assigns it to the NumAdaptors field.
func (o *ComputePhysicalAllOf) SetNumAdaptors(v int64) {
	o.NumAdaptors = &v
}

// GetNumCpuCores returns the NumCpuCores field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumCpuCores() int64 {
	if o == nil || o.NumCpuCores == nil {
		var ret int64
		return ret
	}
	return *o.NumCpuCores
}

// GetNumCpuCoresOk returns a tuple with the NumCpuCores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumCpuCoresOk() (*int64, bool) {
	if o == nil || o.NumCpuCores == nil {
		return nil, false
	}
	return o.NumCpuCores, true
}

// HasNumCpuCores returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumCpuCores() bool {
	if o != nil && o.NumCpuCores != nil {
		return true
	}

	return false
}

// SetNumCpuCores gets a reference to the given int64 and assigns it to the NumCpuCores field.
func (o *ComputePhysicalAllOf) SetNumCpuCores(v int64) {
	o.NumCpuCores = &v
}

// GetNumCpuCoresEnabled returns the NumCpuCoresEnabled field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumCpuCoresEnabled() int64 {
	if o == nil || o.NumCpuCoresEnabled == nil {
		var ret int64
		return ret
	}
	return *o.NumCpuCoresEnabled
}

// GetNumCpuCoresEnabledOk returns a tuple with the NumCpuCoresEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumCpuCoresEnabledOk() (*int64, bool) {
	if o == nil || o.NumCpuCoresEnabled == nil {
		return nil, false
	}
	return o.NumCpuCoresEnabled, true
}

// HasNumCpuCoresEnabled returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumCpuCoresEnabled() bool {
	if o != nil && o.NumCpuCoresEnabled != nil {
		return true
	}

	return false
}

// SetNumCpuCoresEnabled gets a reference to the given int64 and assigns it to the NumCpuCoresEnabled field.
func (o *ComputePhysicalAllOf) SetNumCpuCoresEnabled(v int64) {
	o.NumCpuCoresEnabled = &v
}

// GetNumCpus returns the NumCpus field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumCpus() int64 {
	if o == nil || o.NumCpus == nil {
		var ret int64
		return ret
	}
	return *o.NumCpus
}

// GetNumCpusOk returns a tuple with the NumCpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumCpusOk() (*int64, bool) {
	if o == nil || o.NumCpus == nil {
		return nil, false
	}
	return o.NumCpus, true
}

// HasNumCpus returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumCpus() bool {
	if o != nil && o.NumCpus != nil {
		return true
	}

	return false
}

// SetNumCpus gets a reference to the given int64 and assigns it to the NumCpus field.
func (o *ComputePhysicalAllOf) SetNumCpus(v int64) {
	o.NumCpus = &v
}

// GetNumEthHostInterfaces returns the NumEthHostInterfaces field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumEthHostInterfaces() int64 {
	if o == nil || o.NumEthHostInterfaces == nil {
		var ret int64
		return ret
	}
	return *o.NumEthHostInterfaces
}

// GetNumEthHostInterfacesOk returns a tuple with the NumEthHostInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumEthHostInterfacesOk() (*int64, bool) {
	if o == nil || o.NumEthHostInterfaces == nil {
		return nil, false
	}
	return o.NumEthHostInterfaces, true
}

// HasNumEthHostInterfaces returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumEthHostInterfaces() bool {
	if o != nil && o.NumEthHostInterfaces != nil {
		return true
	}

	return false
}

// SetNumEthHostInterfaces gets a reference to the given int64 and assigns it to the NumEthHostInterfaces field.
func (o *ComputePhysicalAllOf) SetNumEthHostInterfaces(v int64) {
	o.NumEthHostInterfaces = &v
}

// GetNumFcHostInterfaces returns the NumFcHostInterfaces field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumFcHostInterfaces() int64 {
	if o == nil || o.NumFcHostInterfaces == nil {
		var ret int64
		return ret
	}
	return *o.NumFcHostInterfaces
}

// GetNumFcHostInterfacesOk returns a tuple with the NumFcHostInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumFcHostInterfacesOk() (*int64, bool) {
	if o == nil || o.NumFcHostInterfaces == nil {
		return nil, false
	}
	return o.NumFcHostInterfaces, true
}

// HasNumFcHostInterfaces returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumFcHostInterfaces() bool {
	if o != nil && o.NumFcHostInterfaces != nil {
		return true
	}

	return false
}

// SetNumFcHostInterfaces gets a reference to the given int64 and assigns it to the NumFcHostInterfaces field.
func (o *ComputePhysicalAllOf) SetNumFcHostInterfaces(v int64) {
	o.NumFcHostInterfaces = &v
}

// GetNumThreads returns the NumThreads field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetNumThreads() int64 {
	if o == nil || o.NumThreads == nil {
		var ret int64
		return ret
	}
	return *o.NumThreads
}

// GetNumThreadsOk returns a tuple with the NumThreads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetNumThreadsOk() (*int64, bool) {
	if o == nil || o.NumThreads == nil {
		return nil, false
	}
	return o.NumThreads, true
}

// HasNumThreads returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasNumThreads() bool {
	if o != nil && o.NumThreads != nil {
		return true
	}

	return false
}

// SetNumThreads gets a reference to the given int64 and assigns it to the NumThreads field.
func (o *ComputePhysicalAllOf) SetNumThreads(v int64) {
	o.NumThreads = &v
}

// GetOperPowerState returns the OperPowerState field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetOperPowerState() string {
	if o == nil || o.OperPowerState == nil {
		var ret string
		return ret
	}
	return *o.OperPowerState
}

// GetOperPowerStateOk returns a tuple with the OperPowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetOperPowerStateOk() (*string, bool) {
	if o == nil || o.OperPowerState == nil {
		return nil, false
	}
	return o.OperPowerState, true
}

// HasOperPowerState returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasOperPowerState() bool {
	if o != nil && o.OperPowerState != nil {
		return true
	}

	return false
}

// SetOperPowerState gets a reference to the given string and assigns it to the OperPowerState field.
func (o *ComputePhysicalAllOf) SetOperPowerState(v string) {
	o.OperPowerState = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *ComputePhysicalAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *ComputePhysicalAllOf) SetOperability(v string) {
	o.Operability = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *ComputePhysicalAllOf) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetPresence returns the Presence field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetPresence() string {
	if o == nil || o.Presence == nil {
		var ret string
		return ret
	}
	return *o.Presence
}

// GetPresenceOk returns a tuple with the Presence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetPresenceOk() (*string, bool) {
	if o == nil || o.Presence == nil {
		return nil, false
	}
	return o.Presence, true
}

// HasPresence returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasPresence() bool {
	if o != nil && o.Presence != nil {
		return true
	}

	return false
}

// SetPresence gets a reference to the given string and assigns it to the Presence field.
func (o *ComputePhysicalAllOf) SetPresence(v string) {
	o.Presence = &v
}

// GetServiceProfile returns the ServiceProfile field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetServiceProfile() string {
	if o == nil || o.ServiceProfile == nil {
		var ret string
		return ret
	}
	return *o.ServiceProfile
}

// GetServiceProfileOk returns a tuple with the ServiceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetServiceProfileOk() (*string, bool) {
	if o == nil || o.ServiceProfile == nil {
		return nil, false
	}
	return o.ServiceProfile, true
}

// HasServiceProfile returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasServiceProfile() bool {
	if o != nil && o.ServiceProfile != nil {
		return true
	}

	return false
}

// SetServiceProfile gets a reference to the given string and assigns it to the ServiceProfile field.
func (o *ComputePhysicalAllOf) SetServiceProfile(v string) {
	o.ServiceProfile = &v
}

// GetTotalMemory returns the TotalMemory field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetTotalMemory() int64 {
	if o == nil || o.TotalMemory == nil {
		var ret int64
		return ret
	}
	return *o.TotalMemory
}

// GetTotalMemoryOk returns a tuple with the TotalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetTotalMemoryOk() (*int64, bool) {
	if o == nil || o.TotalMemory == nil {
		return nil, false
	}
	return o.TotalMemory, true
}

// HasTotalMemory returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasTotalMemory() bool {
	if o != nil && o.TotalMemory != nil {
		return true
	}

	return false
}

// SetTotalMemory gets a reference to the given int64 and assigns it to the TotalMemory field.
func (o *ComputePhysicalAllOf) SetTotalMemory(v int64) {
	o.TotalMemory = &v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetUserLabel() string {
	if o == nil || o.UserLabel == nil {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetUserLabelOk() (*string, bool) {
	if o == nil || o.UserLabel == nil {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasUserLabel() bool {
	if o != nil && o.UserLabel != nil {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ComputePhysicalAllOf) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ComputePhysicalAllOf) SetUuid(v string) {
	o.Uuid = &v
}

// GetMgmtIdentity returns the MgmtIdentity field value if set, zero value otherwise.
func (o *ComputePhysicalAllOf) GetMgmtIdentity() EquipmentPhysicalIdentityRelationship {
	if o == nil || o.MgmtIdentity == nil {
		var ret EquipmentPhysicalIdentityRelationship
		return ret
	}
	return *o.MgmtIdentity
}

// GetMgmtIdentityOk returns a tuple with the MgmtIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputePhysicalAllOf) GetMgmtIdentityOk() (*EquipmentPhysicalIdentityRelationship, bool) {
	if o == nil || o.MgmtIdentity == nil {
		return nil, false
	}
	return o.MgmtIdentity, true
}

// HasMgmtIdentity returns a boolean if a field has been set.
func (o *ComputePhysicalAllOf) HasMgmtIdentity() bool {
	if o != nil && o.MgmtIdentity != nil {
		return true
	}

	return false
}

// SetMgmtIdentity gets a reference to the given EquipmentPhysicalIdentityRelationship and assigns it to the MgmtIdentity field.
func (o *ComputePhysicalAllOf) SetMgmtIdentity(v EquipmentPhysicalIdentityRelationship) {
	o.MgmtIdentity = &v
}

func (o ComputePhysicalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminPowerState != nil {
		toSerialize["AdminPowerState"] = o.AdminPowerState
	}
	if o.AssetTag != nil {
		toSerialize["AssetTag"] = o.AssetTag
	}
	if o.AvailableMemory != nil {
		toSerialize["AvailableMemory"] = o.AvailableMemory
	}
	if o.BiosPostComplete != nil {
		toSerialize["BiosPostComplete"] = o.BiosPostComplete
	}
	if o.FaultSummary != nil {
		toSerialize["FaultSummary"] = o.FaultSummary
	}
	if o.KvmIpAddresses != nil {
		toSerialize["KvmIpAddresses"] = o.KvmIpAddresses
	}
	if o.ManagementMode != nil {
		toSerialize["ManagementMode"] = o.ManagementMode
	}
	if o.MemorySpeed != nil {
		toSerialize["MemorySpeed"] = o.MemorySpeed
	}
	if o.MgmtIpAddress != nil {
		toSerialize["MgmtIpAddress"] = o.MgmtIpAddress
	}
	if o.NumAdaptors != nil {
		toSerialize["NumAdaptors"] = o.NumAdaptors
	}
	if o.NumCpuCores != nil {
		toSerialize["NumCpuCores"] = o.NumCpuCores
	}
	if o.NumCpuCoresEnabled != nil {
		toSerialize["NumCpuCoresEnabled"] = o.NumCpuCoresEnabled
	}
	if o.NumCpus != nil {
		toSerialize["NumCpus"] = o.NumCpus
	}
	if o.NumEthHostInterfaces != nil {
		toSerialize["NumEthHostInterfaces"] = o.NumEthHostInterfaces
	}
	if o.NumFcHostInterfaces != nil {
		toSerialize["NumFcHostInterfaces"] = o.NumFcHostInterfaces
	}
	if o.NumThreads != nil {
		toSerialize["NumThreads"] = o.NumThreads
	}
	if o.OperPowerState != nil {
		toSerialize["OperPowerState"] = o.OperPowerState
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.Presence != nil {
		toSerialize["Presence"] = o.Presence
	}
	if o.ServiceProfile != nil {
		toSerialize["ServiceProfile"] = o.ServiceProfile
	}
	if o.TotalMemory != nil {
		toSerialize["TotalMemory"] = o.TotalMemory
	}
	if o.UserLabel != nil {
		toSerialize["UserLabel"] = o.UserLabel
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.MgmtIdentity != nil {
		toSerialize["MgmtIdentity"] = o.MgmtIdentity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ComputePhysicalAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varComputePhysicalAllOf := _ComputePhysicalAllOf{}

	if err = json.Unmarshal(bytes, &varComputePhysicalAllOf); err == nil {
		*o = ComputePhysicalAllOf(varComputePhysicalAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "AdminPowerState")
		delete(additionalProperties, "AssetTag")
		delete(additionalProperties, "AvailableMemory")
		delete(additionalProperties, "BiosPostComplete")
		delete(additionalProperties, "FaultSummary")
		delete(additionalProperties, "KvmIpAddresses")
		delete(additionalProperties, "ManagementMode")
		delete(additionalProperties, "MemorySpeed")
		delete(additionalProperties, "MgmtIpAddress")
		delete(additionalProperties, "NumAdaptors")
		delete(additionalProperties, "NumCpuCores")
		delete(additionalProperties, "NumCpuCoresEnabled")
		delete(additionalProperties, "NumCpus")
		delete(additionalProperties, "NumEthHostInterfaces")
		delete(additionalProperties, "NumFcHostInterfaces")
		delete(additionalProperties, "NumThreads")
		delete(additionalProperties, "OperPowerState")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "Presence")
		delete(additionalProperties, "ServiceProfile")
		delete(additionalProperties, "TotalMemory")
		delete(additionalProperties, "UserLabel")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "MgmtIdentity")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableComputePhysicalAllOf struct {
	value *ComputePhysicalAllOf
	isSet bool
}

func (v NullableComputePhysicalAllOf) Get() *ComputePhysicalAllOf {
	return v.value
}

func (v *NullableComputePhysicalAllOf) Set(val *ComputePhysicalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableComputePhysicalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableComputePhysicalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComputePhysicalAllOf(val *ComputePhysicalAllOf) *NullableComputePhysicalAllOf {
	return &NullableComputePhysicalAllOf{value: val, isSet: true}
}

func (v NullableComputePhysicalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComputePhysicalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
