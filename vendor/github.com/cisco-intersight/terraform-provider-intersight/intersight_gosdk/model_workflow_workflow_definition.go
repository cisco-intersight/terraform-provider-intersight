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

// WorkflowWorkflowDefinition Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
type WorkflowWorkflowDefinition struct {
	MoBaseMo
	// When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.
	DefaultVersion *bool `json:"DefaultVersion,omitempty"`
	// The description for this workflow.
	Description     *string                 `json:"Description,omitempty"`
	InputDefinition *[]WorkflowBaseDataType `json:"InputDefinition,omitempty"`
	// A user friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
	Label *string `json:"Label,omitempty"`
	// License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.
	LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
	// The maximum number of tasks that can be executed on this workflow.
	MaxTaskCount *int64 `json:"MaxTaskCount,omitempty"`
	// The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
	Name             *string                 `json:"Name,omitempty"`
	OutputDefinition *[]WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
	// The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output.
	OutputParameters interface{}                 `json:"OutputParameters,omitempty"`
	Properties       *WorkflowWorkflowProperties `json:"Properties,omitempty"`
	Tasks            *[]WorkflowWorkflowTask     `json:"Tasks,omitempty"`
	// This will hold the data needed for workflow to be rendered in the user interface.
	UiRenderingData       interface{}                    `json:"UiRenderingData,omitempty"`
	ValidationInformation *WorkflowValidationInformation `json:"ValidationInformation,omitempty"`
	// The version of the workflow to support multiple versions.
	Version              *int64                       `json:"Version,omitempty"`
	Catalog              *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWorkflowDefinition WorkflowWorkflowDefinition

// NewWorkflowWorkflowDefinition instantiates a new WorkflowWorkflowDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWorkflowDefinition() *WorkflowWorkflowDefinition {
	this := WorkflowWorkflowDefinition{}
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	return &this
}

// NewWorkflowWorkflowDefinitionWithDefaults instantiates a new WorkflowWorkflowDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWorkflowDefinitionWithDefaults() *WorkflowWorkflowDefinition {
	this := WorkflowWorkflowDefinition{}
	var licenseEntitlement string = "Base"
	this.LicenseEntitlement = &licenseEntitlement
	return &this
}

// GetDefaultVersion returns the DefaultVersion field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetDefaultVersion() bool {
	if o == nil || o.DefaultVersion == nil {
		var ret bool
		return ret
	}
	return *o.DefaultVersion
}

// GetDefaultVersionOk returns a tuple with the DefaultVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetDefaultVersionOk() (*bool, bool) {
	if o == nil || o.DefaultVersion == nil {
		return nil, false
	}
	return o.DefaultVersion, true
}

// HasDefaultVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasDefaultVersion() bool {
	if o != nil && o.DefaultVersion != nil {
		return true
	}

	return false
}

// SetDefaultVersion gets a reference to the given bool and assigns it to the DefaultVersion field.
func (o *WorkflowWorkflowDefinition) SetDefaultVersion(v bool) {
	o.DefaultVersion = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkflowWorkflowDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetInputDefinition returns the InputDefinition field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetInputDefinition() []WorkflowBaseDataType {
	if o == nil || o.InputDefinition == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return *o.InputDefinition
}

// GetInputDefinitionOk returns a tuple with the InputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetInputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.InputDefinition == nil {
		return nil, false
	}
	return o.InputDefinition, true
}

// HasInputDefinition returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasInputDefinition() bool {
	if o != nil && o.InputDefinition != nil {
		return true
	}

	return false
}

// SetInputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the InputDefinition field.
func (o *WorkflowWorkflowDefinition) SetInputDefinition(v []WorkflowBaseDataType) {
	o.InputDefinition = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowWorkflowDefinition) SetLabel(v string) {
	o.Label = &v
}

// GetLicenseEntitlement returns the LicenseEntitlement field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetLicenseEntitlement() string {
	if o == nil || o.LicenseEntitlement == nil {
		var ret string
		return ret
	}
	return *o.LicenseEntitlement
}

// GetLicenseEntitlementOk returns a tuple with the LicenseEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetLicenseEntitlementOk() (*string, bool) {
	if o == nil || o.LicenseEntitlement == nil {
		return nil, false
	}
	return o.LicenseEntitlement, true
}

// HasLicenseEntitlement returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasLicenseEntitlement() bool {
	if o != nil && o.LicenseEntitlement != nil {
		return true
	}

	return false
}

// SetLicenseEntitlement gets a reference to the given string and assigns it to the LicenseEntitlement field.
func (o *WorkflowWorkflowDefinition) SetLicenseEntitlement(v string) {
	o.LicenseEntitlement = &v
}

// GetMaxTaskCount returns the MaxTaskCount field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetMaxTaskCount() int64 {
	if o == nil || o.MaxTaskCount == nil {
		var ret int64
		return ret
	}
	return *o.MaxTaskCount
}

// GetMaxTaskCountOk returns a tuple with the MaxTaskCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetMaxTaskCountOk() (*int64, bool) {
	if o == nil || o.MaxTaskCount == nil {
		return nil, false
	}
	return o.MaxTaskCount, true
}

// HasMaxTaskCount returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasMaxTaskCount() bool {
	if o != nil && o.MaxTaskCount != nil {
		return true
	}

	return false
}

// SetMaxTaskCount gets a reference to the given int64 and assigns it to the MaxTaskCount field.
func (o *WorkflowWorkflowDefinition) SetMaxTaskCount(v int64) {
	o.MaxTaskCount = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowWorkflowDefinition) SetName(v string) {
	o.Name = &v
}

// GetOutputDefinition returns the OutputDefinition field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetOutputDefinition() []WorkflowBaseDataType {
	if o == nil || o.OutputDefinition == nil {
		var ret []WorkflowBaseDataType
		return ret
	}
	return *o.OutputDefinition
}

// GetOutputDefinitionOk returns a tuple with the OutputDefinition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetOutputDefinitionOk() (*[]WorkflowBaseDataType, bool) {
	if o == nil || o.OutputDefinition == nil {
		return nil, false
	}
	return o.OutputDefinition, true
}

// HasOutputDefinition returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasOutputDefinition() bool {
	if o != nil && o.OutputDefinition != nil {
		return true
	}

	return false
}

// SetOutputDefinition gets a reference to the given []WorkflowBaseDataType and assigns it to the OutputDefinition field.
func (o *WorkflowWorkflowDefinition) SetOutputDefinition(v []WorkflowBaseDataType) {
	o.OutputDefinition = &v
}

// GetOutputParameters returns the OutputParameters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinition) GetOutputParameters() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OutputParameters
}

// GetOutputParametersOk returns a tuple with the OutputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinition) GetOutputParametersOk() (*interface{}, bool) {
	if o == nil || o.OutputParameters == nil {
		return nil, false
	}
	return &o.OutputParameters, true
}

// HasOutputParameters returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasOutputParameters() bool {
	if o != nil && o.OutputParameters != nil {
		return true
	}

	return false
}

// SetOutputParameters gets a reference to the given interface{} and assigns it to the OutputParameters field.
func (o *WorkflowWorkflowDefinition) SetOutputParameters(v interface{}) {
	o.OutputParameters = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetProperties() WorkflowWorkflowProperties {
	if o == nil || o.Properties == nil {
		var ret WorkflowWorkflowProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetPropertiesOk() (*WorkflowWorkflowProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given WorkflowWorkflowProperties and assigns it to the Properties field.
func (o *WorkflowWorkflowDefinition) SetProperties(v WorkflowWorkflowProperties) {
	o.Properties = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetTasks() []WorkflowWorkflowTask {
	if o == nil || o.Tasks == nil {
		var ret []WorkflowWorkflowTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetTasksOk() (*[]WorkflowWorkflowTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []WorkflowWorkflowTask and assigns it to the Tasks field.
func (o *WorkflowWorkflowDefinition) SetTasks(v []WorkflowWorkflowTask) {
	o.Tasks = &v
}

// GetUiRenderingData returns the UiRenderingData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWorkflowDefinition) GetUiRenderingData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.UiRenderingData
}

// GetUiRenderingDataOk returns a tuple with the UiRenderingData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWorkflowDefinition) GetUiRenderingDataOk() (*interface{}, bool) {
	if o == nil || o.UiRenderingData == nil {
		return nil, false
	}
	return &o.UiRenderingData, true
}

// HasUiRenderingData returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasUiRenderingData() bool {
	if o != nil && o.UiRenderingData != nil {
		return true
	}

	return false
}

// SetUiRenderingData gets a reference to the given interface{} and assigns it to the UiRenderingData field.
func (o *WorkflowWorkflowDefinition) SetUiRenderingData(v interface{}) {
	o.UiRenderingData = v
}

// GetValidationInformation returns the ValidationInformation field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetValidationInformation() WorkflowValidationInformation {
	if o == nil || o.ValidationInformation == nil {
		var ret WorkflowValidationInformation
		return ret
	}
	return *o.ValidationInformation
}

// GetValidationInformationOk returns a tuple with the ValidationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetValidationInformationOk() (*WorkflowValidationInformation, bool) {
	if o == nil || o.ValidationInformation == nil {
		return nil, false
	}
	return o.ValidationInformation, true
}

// HasValidationInformation returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasValidationInformation() bool {
	if o != nil && o.ValidationInformation != nil {
		return true
	}

	return false
}

// SetValidationInformation gets a reference to the given WorkflowValidationInformation and assigns it to the ValidationInformation field.
func (o *WorkflowWorkflowDefinition) SetValidationInformation(v WorkflowValidationInformation) {
	o.ValidationInformation = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetVersion() int64 {
	if o == nil || o.Version == nil {
		var ret int64
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetVersionOk() (*int64, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int64 and assigns it to the Version field.
func (o *WorkflowWorkflowDefinition) SetVersion(v int64) {
	o.Version = &v
}

// GetCatalog returns the Catalog field value if set, zero value otherwise.
func (o *WorkflowWorkflowDefinition) GetCatalog() WorkflowCatalogRelationship {
	if o == nil || o.Catalog == nil {
		var ret WorkflowCatalogRelationship
		return ret
	}
	return *o.Catalog
}

// GetCatalogOk returns a tuple with the Catalog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowWorkflowDefinition) GetCatalogOk() (*WorkflowCatalogRelationship, bool) {
	if o == nil || o.Catalog == nil {
		return nil, false
	}
	return o.Catalog, true
}

// HasCatalog returns a boolean if a field has been set.
func (o *WorkflowWorkflowDefinition) HasCatalog() bool {
	if o != nil && o.Catalog != nil {
		return true
	}

	return false
}

// SetCatalog gets a reference to the given WorkflowCatalogRelationship and assigns it to the Catalog field.
func (o *WorkflowWorkflowDefinition) SetCatalog(v WorkflowCatalogRelationship) {
	o.Catalog = &v
}

func (o WorkflowWorkflowDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if o.DefaultVersion != nil {
		toSerialize["DefaultVersion"] = o.DefaultVersion
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.InputDefinition != nil {
		toSerialize["InputDefinition"] = o.InputDefinition
	}
	if o.Label != nil {
		toSerialize["Label"] = o.Label
	}
	if o.LicenseEntitlement != nil {
		toSerialize["LicenseEntitlement"] = o.LicenseEntitlement
	}
	if o.MaxTaskCount != nil {
		toSerialize["MaxTaskCount"] = o.MaxTaskCount
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OutputDefinition != nil {
		toSerialize["OutputDefinition"] = o.OutputDefinition
	}
	if o.OutputParameters != nil {
		toSerialize["OutputParameters"] = o.OutputParameters
	}
	if o.Properties != nil {
		toSerialize["Properties"] = o.Properties
	}
	if o.Tasks != nil {
		toSerialize["Tasks"] = o.Tasks
	}
	if o.UiRenderingData != nil {
		toSerialize["UiRenderingData"] = o.UiRenderingData
	}
	if o.ValidationInformation != nil {
		toSerialize["ValidationInformation"] = o.ValidationInformation
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}
	if o.Catalog != nil {
		toSerialize["Catalog"] = o.Catalog
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWorkflowDefinition) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowWorkflowDefinitionWithoutEmbeddedStruct struct {
		// When true this will be the workflow version that is used when a specific workflow definition version is not specified. The default version is used when user executes a workflow without specifying a version or when workflow is included in another workflow without a specific version. The very first workflow definition created with a name will be set as the default version, after that user can explicitly set any version of the workflow definition as the default version.
		DefaultVersion *bool `json:"DefaultVersion,omitempty"`
		// The description for this workflow.
		Description     *string                 `json:"Description,omitempty"`
		InputDefinition *[]WorkflowBaseDataType `json:"InputDefinition,omitempty"`
		// A user friendly short name to identify the workflow. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ) or an underscore (_).
		Label *string `json:"Label,omitempty"`
		// License entitlement required to run this workflow. It is calculated based on the highest license requirement of all its tasks.
		LicenseEntitlement *string `json:"LicenseEntitlement,omitempty"`
		// The maximum number of tasks that can be executed on this workflow.
		MaxTaskCount *int64 `json:"MaxTaskCount,omitempty"`
		// The name for this workflow. You can have multiple versions of the workflow with the same name. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.) or an underscore (_).
		Name             *string                 `json:"Name,omitempty"`
		OutputDefinition *[]WorkflowBaseDataType `json:"OutputDefinition,omitempty"`
		// The output mappings for the workflow. The outputs for worflows will generally be task output variables that we want to export out at the end of the workflow. The format to specify the mapping is '${Source.output.JsonPath}'. 'Source' is the name of the task within the workflow. You can map any task output to a workflow output as long as the types are compatible. Following this is JSON path expression to extract JSON fragment from source's output.
		OutputParameters interface{}                 `json:"OutputParameters,omitempty"`
		Properties       *WorkflowWorkflowProperties `json:"Properties,omitempty"`
		Tasks            *[]WorkflowWorkflowTask     `json:"Tasks,omitempty"`
		// This will hold the data needed for workflow to be rendered in the user interface.
		UiRenderingData       interface{}                    `json:"UiRenderingData,omitempty"`
		ValidationInformation *WorkflowValidationInformation `json:"ValidationInformation,omitempty"`
		// The version of the workflow to support multiple versions.
		Version *int64                       `json:"Version,omitempty"`
		Catalog *WorkflowCatalogRelationship `json:"Catalog,omitempty"`
	}

	varWorkflowWorkflowDefinitionWithoutEmbeddedStruct := WorkflowWorkflowDefinitionWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowDefinitionWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowWorkflowDefinition := _WorkflowWorkflowDefinition{}
		varWorkflowWorkflowDefinition.DefaultVersion = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.DefaultVersion
		varWorkflowWorkflowDefinition.Description = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Description
		varWorkflowWorkflowDefinition.InputDefinition = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.InputDefinition
		varWorkflowWorkflowDefinition.Label = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Label
		varWorkflowWorkflowDefinition.LicenseEntitlement = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.LicenseEntitlement
		varWorkflowWorkflowDefinition.MaxTaskCount = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.MaxTaskCount
		varWorkflowWorkflowDefinition.Name = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Name
		varWorkflowWorkflowDefinition.OutputDefinition = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.OutputDefinition
		varWorkflowWorkflowDefinition.OutputParameters = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.OutputParameters
		varWorkflowWorkflowDefinition.Properties = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Properties
		varWorkflowWorkflowDefinition.Tasks = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Tasks
		varWorkflowWorkflowDefinition.UiRenderingData = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.UiRenderingData
		varWorkflowWorkflowDefinition.ValidationInformation = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.ValidationInformation
		varWorkflowWorkflowDefinition.Version = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Version
		varWorkflowWorkflowDefinition.Catalog = varWorkflowWorkflowDefinitionWithoutEmbeddedStruct.Catalog
		*o = WorkflowWorkflowDefinition(varWorkflowWorkflowDefinition)
	} else {
		return err
	}

	varWorkflowWorkflowDefinition := _WorkflowWorkflowDefinition{}

	err = json.Unmarshal(bytes, &varWorkflowWorkflowDefinition)
	if err == nil {
		o.MoBaseMo = varWorkflowWorkflowDefinition.MoBaseMo
	} else {
		return err
	}
	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "DefaultVersion")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "InputDefinition")
		delete(additionalProperties, "Label")
		delete(additionalProperties, "LicenseEntitlement")
		delete(additionalProperties, "MaxTaskCount")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OutputDefinition")
		delete(additionalProperties, "OutputParameters")
		delete(additionalProperties, "Properties")
		delete(additionalProperties, "Tasks")
		delete(additionalProperties, "UiRenderingData")
		delete(additionalProperties, "ValidationInformation")
		delete(additionalProperties, "Version")
		delete(additionalProperties, "Catalog")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableWorkflowWorkflowDefinition struct {
	value *WorkflowWorkflowDefinition
	isSet bool
}

func (v NullableWorkflowWorkflowDefinition) Get() *WorkflowWorkflowDefinition {
	return v.value
}

func (v *NullableWorkflowWorkflowDefinition) Set(val *WorkflowWorkflowDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWorkflowDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWorkflowDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWorkflowDefinition(val *WorkflowWorkflowDefinition) *NullableWorkflowWorkflowDefinition {
	return &NullableWorkflowWorkflowDefinition{value: val, isSet: true}
}

func (v NullableWorkflowWorkflowDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWorkflowDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
