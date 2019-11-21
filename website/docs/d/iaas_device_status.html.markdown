---
layout: "intersight"
page_title: "Intersight: intersight_iaas_device_status"
sidebar_current: "docs-intersight-data-source-iaasDeviceStatus"
description: |-
List of infra accounts managed by UCSD.

---

# Data Source: intersight_iaas_device_status
List of infra accounts managed by UCSD.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_name`:(string)The UCSD infra account name. Account Name is created when UCSD admin adds any new infra account (Physical/Virtual/Compute/Network) to be managed by UCSD.
* `account_type`:(string)The UCSD Infra Account type.
* `connection_status`:(string)Describes about the connection status between the UCSD and the actual end device.
* `device_model`:(string)Describes about the device model.
* `device_vendor`:(string)Describes about the device vendor/manufacturer of the device.
* `device_version`:(string)Describes about the current firmware version running on the device.
* `ip_address`:(string)The IPAddress of the device.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pod`:(string)Describes about the pod to which this device belongs to in UCSD.
* `pod_type`:(string)Describes about the podType of Pod to which this device belongs to in UCSD.
