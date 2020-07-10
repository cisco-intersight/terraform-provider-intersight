# HyperflexHxapVirtualMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffinitySelectors** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**Age** | Pointer to **string** | Denotes age or life time of the VM in nano seconds. | [optional] 
**AntiAffinitySelectors** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**Disks** | Pointer to [**[]HyperflexVmDisk**](hyperflex.VmDisk.md) |  | [optional] 
**FailureReason** | Pointer to **string** | Reason of the failure when VM is in failed state. | [optional] 
**GraphicConsoleUrl** | Pointer to **string** | Graphical console URL of this VM. | [optional] 
**Interfaces** | Pointer to [**[]HyperflexVmInterface**](hyperflex.VmInterface.md) |  | [optional] 
**Labels** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**NetworkCount** | Pointer to **int64** | Number network interfaces associated with the virtual machine. | [optional] 
**StartTime** | Pointer to **string** | Denotes the VM start timestamp. | [optional] 
**Status** | Pointer to **string** | Status of virtual machine. | [optional] [default to "Unknown"]
**UpTime** | Pointer to **string** | Denotes how long this VM has been running in nano seconds. | [optional] 
**Cluster** | Pointer to [**HyperflexHxapClusterRelationship**](hyperflex.HxapCluster.Relationship.md) |  | [optional] 
**Host** | Pointer to [**HyperflexHxapHostRelationship**](hyperflex.HxapHost.Relationship.md) |  | [optional] 

## Methods

### NewHyperflexHxapVirtualMachine

`func NewHyperflexHxapVirtualMachine() *HyperflexHxapVirtualMachine`

NewHyperflexHxapVirtualMachine instantiates a new HyperflexHxapVirtualMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperflexHxapVirtualMachineWithDefaults

`func NewHyperflexHxapVirtualMachineWithDefaults() *HyperflexHxapVirtualMachine`

NewHyperflexHxapVirtualMachineWithDefaults instantiates a new HyperflexHxapVirtualMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachine) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### GetAge

`func (o *HyperflexHxapVirtualMachine) GetAge() string`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *HyperflexHxapVirtualMachine) GetAgeOk() (*string, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *HyperflexHxapVirtualMachine) SetAge(v string)`

SetAge sets Age field to given value.

### HasAge

`func (o *HyperflexHxapVirtualMachine) HasAge() bool`

HasAge returns a boolean if a field has been set.

### GetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *HyperflexHxapVirtualMachine) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *HyperflexHxapVirtualMachine) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### GetDisks

`func (o *HyperflexHxapVirtualMachine) GetDisks() []HyperflexVmDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *HyperflexHxapVirtualMachine) GetDisksOk() (*[]HyperflexVmDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *HyperflexHxapVirtualMachine) SetDisks(v []HyperflexVmDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *HyperflexHxapVirtualMachine) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetFailureReason

`func (o *HyperflexHxapVirtualMachine) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *HyperflexHxapVirtualMachine) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *HyperflexHxapVirtualMachine) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *HyperflexHxapVirtualMachine) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetGraphicConsoleUrl

`func (o *HyperflexHxapVirtualMachine) GetGraphicConsoleUrl() string`

GetGraphicConsoleUrl returns the GraphicConsoleUrl field if non-nil, zero value otherwise.

### GetGraphicConsoleUrlOk

`func (o *HyperflexHxapVirtualMachine) GetGraphicConsoleUrlOk() (*string, bool)`

GetGraphicConsoleUrlOk returns a tuple with the GraphicConsoleUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphicConsoleUrl

`func (o *HyperflexHxapVirtualMachine) SetGraphicConsoleUrl(v string)`

SetGraphicConsoleUrl sets GraphicConsoleUrl field to given value.

### HasGraphicConsoleUrl

`func (o *HyperflexHxapVirtualMachine) HasGraphicConsoleUrl() bool`

HasGraphicConsoleUrl returns a boolean if a field has been set.

### GetInterfaces

`func (o *HyperflexHxapVirtualMachine) GetInterfaces() []HyperflexVmInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *HyperflexHxapVirtualMachine) GetInterfacesOk() (*[]HyperflexVmInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *HyperflexHxapVirtualMachine) SetInterfaces(v []HyperflexVmInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *HyperflexHxapVirtualMachine) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *HyperflexHxapVirtualMachine) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *HyperflexHxapVirtualMachine) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *HyperflexHxapVirtualMachine) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *HyperflexHxapVirtualMachine) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetNetworkCount

`func (o *HyperflexHxapVirtualMachine) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *HyperflexHxapVirtualMachine) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *HyperflexHxapVirtualMachine) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *HyperflexHxapVirtualMachine) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetStartTime

`func (o *HyperflexHxapVirtualMachine) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *HyperflexHxapVirtualMachine) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *HyperflexHxapVirtualMachine) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *HyperflexHxapVirtualMachine) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *HyperflexHxapVirtualMachine) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HyperflexHxapVirtualMachine) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HyperflexHxapVirtualMachine) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HyperflexHxapVirtualMachine) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpTime

`func (o *HyperflexHxapVirtualMachine) GetUpTime() string`

GetUpTime returns the UpTime field if non-nil, zero value otherwise.

### GetUpTimeOk

`func (o *HyperflexHxapVirtualMachine) GetUpTimeOk() (*string, bool)`

GetUpTimeOk returns a tuple with the UpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpTime

`func (o *HyperflexHxapVirtualMachine) SetUpTime(v string)`

SetUpTime sets UpTime field to given value.

### HasUpTime

`func (o *HyperflexHxapVirtualMachine) HasUpTime() bool`

HasUpTime returns a boolean if a field has been set.

### GetCluster

`func (o *HyperflexHxapVirtualMachine) GetCluster() HyperflexHxapClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *HyperflexHxapVirtualMachine) GetClusterOk() (*HyperflexHxapClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *HyperflexHxapVirtualMachine) SetCluster(v HyperflexHxapClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *HyperflexHxapVirtualMachine) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *HyperflexHxapVirtualMachine) GetHost() HyperflexHxapHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *HyperflexHxapVirtualMachine) GetHostOk() (*HyperflexHxapHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *HyperflexHxapVirtualMachine) SetHost(v HyperflexHxapHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *HyperflexHxapVirtualMachine) HasHost() bool`

HasHost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


