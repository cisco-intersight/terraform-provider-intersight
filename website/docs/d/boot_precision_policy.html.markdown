---
layout: "intersight"
page_title: "Intersight: intersight_boot_precision_policy"
sidebar_current: "docs-intersight-data-source-bootPrecisionPolicy"
description: |-
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
---

# Data Source: intersight_boot_precision_policy
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `configured_boot_mode`:(string)Sets the BIOS boot mode. UEFI uses the GUID Partition Table (GPT) whereas Legacy mode uses the Master Boot Record (MBR) partitioning scheme.
* `description`:(string)Description of the policy.
* `enforce_uefi_secure_boot`:(bool)If UEFI secure boot is enabled, the boot mode is set to UEFI by default. Secure boot enforces that device boots using only software that is trusted by the Original Equipment Manufacturer (OEM).
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
