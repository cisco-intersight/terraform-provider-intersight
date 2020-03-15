---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_virtual_machine"
sidebar_current: "docs-intersight-data-source-virtualizationVmwareVirtualMachine"
description: |-
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
---

# Data Source: intersight_virtualization_vmware_virtual_machine
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `annotation`:(string)List of annotations provided to this VM by user; Can be long.
* `boot_time`:(string)Time when this vm booted up.
* `cpu_hot_add_enabled`:(bool)Is the capability to add CPUs to a running VM enabled.
* `config_name`:(string)The named config for this vm, might be the same as guestHostname.
* `connection_state`:(string)Is virtual machine connected to vCenter. Values are connected, disconnected, orphaned, inaccessible, invalid.
* `default_power_off_type`:(string)Indication of how the VM will be powered off (soft, hard, etc).
* `dhcp_enabled`:(bool)Is Dhcp used for IP/DNS on this vm.
* `folder`:(string)The folder name associated with this VM.
* `guest_state`:(string)State of the guest OS on this VM. Is it running, notRunning, etc.
* `hypervisor_type`:(string)Type of hypervisor where the virtual machine is hosted, for example VMware ESXi.
* `identity`:(string)The internally generated identity of this vm. This entity is not manipulated by users. It aids in uniquely identifying the virtual machine object. For VMware, this is MOR (managed object reference).
* `instance_uuid`:(string)UUID assigned by vCenter to every VM.
* `is_template`:(bool)If true, indicates that the entity refers to a template of a virtual machine and not a real virtual machine.
* `memory_hot_add_enabled`:(bool)Adding memory to a running VM.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)User-provided name to identify the virtual machine.
* `network_count`:(int)How many networks are used by this VM.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `power_state`:(string)Power state of the virtual machine.
* `protected_vm`:(bool)Is this a protected VM. VMs can be in protection groups.
* `remote_display_vnc_enabled`:(bool)Is support for a remote VNC access enabled.
* `resource_pool`:(string)Name of the resource pool to which this vm belongs (optional).
* `resource_pool_owner`:(string)Who owns the resource pool.
* `resource_pool_parent`:(string)The parent of the current resource pool to which this VM belongs.
* `tool_running_status`:(string)Are guest tools running on this vm. Set to (guestToolNotRunning,  guestToolsRunning).
* `tools_version`:(string)The version of the guest tools, usually not specified.
* `uuid`:(string)The uuid of this virtual machine. The uuid is internally generated and not user specified.
* `vm_disk_count`:(int)How many disks are assigned to this VM.
* `vm_overall_status`:(string)The operational state (there are many possible states, Available, Provisioned, Maintenance mode, Deleting, etc.) of this vm.
* `vm_path`:(string)Example - [datastore3] VCSA-134/VCSA-134.vmx.
* `vm_version`:(string)Information about the version of this VM (vmx-09, vmx-11 etc.).
* `vm_vnic_count`:(int)How many vnics are present.
* `vnic_device_config_id`:(string)Information related to the guest info's vnic virtual device (this is a comma separated list).
