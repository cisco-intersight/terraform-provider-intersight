---
layout: "intersight"
page_title: "Intersight: intersight_virtualization_vmware_vcenter"
sidebar_current: "docs-intersight-data-source-virtualizationVmwareVcenter"
description: |-
VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
---

# Data Source: intersight_virtualization_vmware_vcenter
VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `identity`:(string)Identity of the hypervisor (not manipulated by user). It could be a UUID too. Example - c917093f-5443-4748-bc09-eec72ded7608.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The user provided name for the hypervisor manager (for example, vCenterIreland). Usually, this name is subject to manipulations by user. It is not the identity of the hypervisor.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `version`:(string)Release version of the Hypervisor Manger (VMware vCenter Server 6.0.0 build-4541947).
