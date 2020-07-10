# VirtualizationVirtualMachineAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | Action to be performed on a virtual machine (Create, PowerState, Migrate, Clone etc). | [optional] [default to "None"]
**ActionInfo** | Pointer to [**VirtualizationActionInfo**](virtualization.ActionInfo.md) |  | [optional] 
**AffinitySelectors** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**AntiAffinitySelectors** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**CloudInitConfig** | Pointer to [**VirtualizationCloudInitConfig**](virtualization.CloudInitConfig.md) |  | [optional] 
**Cpu** | Pointer to **int64** | Number of vCPUs allocated to virtual machine. | [optional] 
**Discovered** | Pointer to **bool** | Flag to indicate whether the configuration is created from inventory object. | [optional] [readonly] 
**Disk** | Pointer to [**[]VirtualizationVirtualMachineDisk**](virtualization.VirtualMachineDisk.md) |  | [optional] 
**GuestOs** | Pointer to **string** | Guest operating system running on virtual machine. | [optional] [default to "linux"]
**Interface** | Pointer to [**[]VirtualizationNetworkInterface**](virtualization.NetworkInterface.md) |  | [optional] 
**Labels** | Pointer to [**[]InfraMetaData**](infra.MetaData.md) |  | [optional] 
**Memory** | Pointer to **int64** | Virtual machine memory defined in mega bytes. | [optional] 
**Name** | Pointer to **string** | Virtual machine name contains only letters, numbers, allowed special character and must be unique. | [optional] 
**PowerState** | Pointer to **string** | Expected power state of virtual machine (PowerOn, PowerOff, Restart). | [optional] [default to "PowerOff"]
**Cluster** | Pointer to [**VirtualizationBaseClusterRelationship**](virtualization.BaseCluster.Relationship.md) |  | [optional] 
**Host** | Pointer to [**VirtualizationBaseHostRelationship**](virtualization.BaseHost.Relationship.md) |  | [optional] 
**Inventory** | Pointer to [**VirtualizationBaseVirtualMachineRelationship**](virtualization.BaseVirtualMachine.Relationship.md) |  | [optional] 
**RegisteredDevice** | Pointer to [**AssetDeviceRegistrationRelationship**](asset.DeviceRegistration.Relationship.md) |  | [optional] 
**WorkflowInfo** | Pointer to [**WorkflowWorkflowInfoRelationship**](workflow.WorkflowInfo.Relationship.md) |  | [optional] 

## Methods

### NewVirtualizationVirtualMachineAllOf

`func NewVirtualizationVirtualMachineAllOf() *VirtualizationVirtualMachineAllOf`

NewVirtualizationVirtualMachineAllOf instantiates a new VirtualizationVirtualMachineAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualizationVirtualMachineAllOfWithDefaults

`func NewVirtualizationVirtualMachineAllOfWithDefaults() *VirtualizationVirtualMachineAllOf`

NewVirtualizationVirtualMachineAllOfWithDefaults instantiates a new VirtualizationVirtualMachineAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VirtualizationVirtualMachineAllOf) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VirtualizationVirtualMachineAllOf) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VirtualizationVirtualMachineAllOf) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *VirtualizationVirtualMachineAllOf) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetActionInfo

`func (o *VirtualizationVirtualMachineAllOf) GetActionInfo() VirtualizationActionInfo`

GetActionInfo returns the ActionInfo field if non-nil, zero value otherwise.

### GetActionInfoOk

`func (o *VirtualizationVirtualMachineAllOf) GetActionInfoOk() (*VirtualizationActionInfo, bool)`

GetActionInfoOk returns a tuple with the ActionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionInfo

`func (o *VirtualizationVirtualMachineAllOf) SetActionInfo(v VirtualizationActionInfo)`

SetActionInfo sets ActionInfo field to given value.

### HasActionInfo

`func (o *VirtualizationVirtualMachineAllOf) HasActionInfo() bool`

HasActionInfo returns a boolean if a field has been set.

### GetAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) GetAffinitySelectors() []InfraMetaData`

GetAffinitySelectors returns the AffinitySelectors field if non-nil, zero value otherwise.

### GetAffinitySelectorsOk

`func (o *VirtualizationVirtualMachineAllOf) GetAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAffinitySelectorsOk returns a tuple with the AffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) SetAffinitySelectors(v []InfraMetaData)`

SetAffinitySelectors sets AffinitySelectors field to given value.

### HasAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) HasAffinitySelectors() bool`

HasAffinitySelectors returns a boolean if a field has been set.

### GetAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) GetAntiAffinitySelectors() []InfraMetaData`

GetAntiAffinitySelectors returns the AntiAffinitySelectors field if non-nil, zero value otherwise.

### GetAntiAffinitySelectorsOk

`func (o *VirtualizationVirtualMachineAllOf) GetAntiAffinitySelectorsOk() (*[]InfraMetaData, bool)`

GetAntiAffinitySelectorsOk returns a tuple with the AntiAffinitySelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) SetAntiAffinitySelectors(v []InfraMetaData)`

SetAntiAffinitySelectors sets AntiAffinitySelectors field to given value.

### HasAntiAffinitySelectors

`func (o *VirtualizationVirtualMachineAllOf) HasAntiAffinitySelectors() bool`

HasAntiAffinitySelectors returns a boolean if a field has been set.

### GetCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) GetCloudInitConfig() VirtualizationCloudInitConfig`

GetCloudInitConfig returns the CloudInitConfig field if non-nil, zero value otherwise.

### GetCloudInitConfigOk

`func (o *VirtualizationVirtualMachineAllOf) GetCloudInitConfigOk() (*VirtualizationCloudInitConfig, bool)`

GetCloudInitConfigOk returns a tuple with the CloudInitConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) SetCloudInitConfig(v VirtualizationCloudInitConfig)`

SetCloudInitConfig sets CloudInitConfig field to given value.

### HasCloudInitConfig

`func (o *VirtualizationVirtualMachineAllOf) HasCloudInitConfig() bool`

HasCloudInitConfig returns a boolean if a field has been set.

### GetCpu

`func (o *VirtualizationVirtualMachineAllOf) GetCpu() int64`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *VirtualizationVirtualMachineAllOf) GetCpuOk() (*int64, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *VirtualizationVirtualMachineAllOf) SetCpu(v int64)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *VirtualizationVirtualMachineAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiscovered

`func (o *VirtualizationVirtualMachineAllOf) GetDiscovered() bool`

GetDiscovered returns the Discovered field if non-nil, zero value otherwise.

### GetDiscoveredOk

`func (o *VirtualizationVirtualMachineAllOf) GetDiscoveredOk() (*bool, bool)`

GetDiscoveredOk returns a tuple with the Discovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovered

`func (o *VirtualizationVirtualMachineAllOf) SetDiscovered(v bool)`

SetDiscovered sets Discovered field to given value.

### HasDiscovered

`func (o *VirtualizationVirtualMachineAllOf) HasDiscovered() bool`

HasDiscovered returns a boolean if a field has been set.

### GetDisk

`func (o *VirtualizationVirtualMachineAllOf) GetDisk() []VirtualizationVirtualMachineDisk`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VirtualizationVirtualMachineAllOf) GetDiskOk() (*[]VirtualizationVirtualMachineDisk, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VirtualizationVirtualMachineAllOf) SetDisk(v []VirtualizationVirtualMachineDisk)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *VirtualizationVirtualMachineAllOf) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGuestOs

`func (o *VirtualizationVirtualMachineAllOf) GetGuestOs() string`

GetGuestOs returns the GuestOs field if non-nil, zero value otherwise.

### GetGuestOsOk

`func (o *VirtualizationVirtualMachineAllOf) GetGuestOsOk() (*string, bool)`

GetGuestOsOk returns a tuple with the GuestOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestOs

`func (o *VirtualizationVirtualMachineAllOf) SetGuestOs(v string)`

SetGuestOs sets GuestOs field to given value.

### HasGuestOs

`func (o *VirtualizationVirtualMachineAllOf) HasGuestOs() bool`

HasGuestOs returns a boolean if a field has been set.

### GetInterface

`func (o *VirtualizationVirtualMachineAllOf) GetInterface() []VirtualizationNetworkInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *VirtualizationVirtualMachineAllOf) GetInterfaceOk() (*[]VirtualizationNetworkInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *VirtualizationVirtualMachineAllOf) SetInterface(v []VirtualizationNetworkInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *VirtualizationVirtualMachineAllOf) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetLabels

`func (o *VirtualizationVirtualMachineAllOf) GetLabels() []InfraMetaData`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *VirtualizationVirtualMachineAllOf) GetLabelsOk() (*[]InfraMetaData, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *VirtualizationVirtualMachineAllOf) SetLabels(v []InfraMetaData)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *VirtualizationVirtualMachineAllOf) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMemory

`func (o *VirtualizationVirtualMachineAllOf) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *VirtualizationVirtualMachineAllOf) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *VirtualizationVirtualMachineAllOf) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *VirtualizationVirtualMachineAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetName

`func (o *VirtualizationVirtualMachineAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualizationVirtualMachineAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualizationVirtualMachineAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualizationVirtualMachineAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPowerState

`func (o *VirtualizationVirtualMachineAllOf) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VirtualizationVirtualMachineAllOf) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VirtualizationVirtualMachineAllOf) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *VirtualizationVirtualMachineAllOf) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetCluster

`func (o *VirtualizationVirtualMachineAllOf) GetCluster() VirtualizationBaseClusterRelationship`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VirtualizationVirtualMachineAllOf) GetClusterOk() (*VirtualizationBaseClusterRelationship, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VirtualizationVirtualMachineAllOf) SetCluster(v VirtualizationBaseClusterRelationship)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VirtualizationVirtualMachineAllOf) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetHost

`func (o *VirtualizationVirtualMachineAllOf) GetHost() VirtualizationBaseHostRelationship`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VirtualizationVirtualMachineAllOf) GetHostOk() (*VirtualizationBaseHostRelationship, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VirtualizationVirtualMachineAllOf) SetHost(v VirtualizationBaseHostRelationship)`

SetHost sets Host field to given value.

### HasHost

`func (o *VirtualizationVirtualMachineAllOf) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetInventory

`func (o *VirtualizationVirtualMachineAllOf) GetInventory() VirtualizationBaseVirtualMachineRelationship`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *VirtualizationVirtualMachineAllOf) GetInventoryOk() (*VirtualizationBaseVirtualMachineRelationship, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *VirtualizationVirtualMachineAllOf) SetInventory(v VirtualizationBaseVirtualMachineRelationship)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *VirtualizationVirtualMachineAllOf) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship`

GetRegisteredDevice returns the RegisteredDevice field if non-nil, zero value otherwise.

### GetRegisteredDeviceOk

`func (o *VirtualizationVirtualMachineAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool)`

GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship)`

SetRegisteredDevice sets RegisteredDevice field to given value.

### HasRegisteredDevice

`func (o *VirtualizationVirtualMachineAllOf) HasRegisteredDevice() bool`

HasRegisteredDevice returns a boolean if a field has been set.

### GetWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) GetWorkflowInfo() WorkflowWorkflowInfoRelationship`

GetWorkflowInfo returns the WorkflowInfo field if non-nil, zero value otherwise.

### GetWorkflowInfoOk

`func (o *VirtualizationVirtualMachineAllOf) GetWorkflowInfoOk() (*WorkflowWorkflowInfoRelationship, bool)`

GetWorkflowInfoOk returns a tuple with the WorkflowInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) SetWorkflowInfo(v WorkflowWorkflowInfoRelationship)`

SetWorkflowInfo sets WorkflowInfo field to given value.

### HasWorkflowInfo

`func (o *VirtualizationVirtualMachineAllOf) HasWorkflowInfo() bool`

HasWorkflowInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


