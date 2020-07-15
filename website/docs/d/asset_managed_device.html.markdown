
---
layout: "intersight"
page_title: "Intersight: intersight_asset_managed_device"
sidebar_current: "docs-intersight-data-source-asset-managed-device"
description: |-
Attributes for Managed Device in Intersight and it maintains the relationship to the Intersight Assist Device. Once added, Device Connector for the Managed Device type is started on the Intersight Assist and status related to it is maintained.
---

# Data Source: intersight_asset._managed_device
Attributes for Managed Device in Intersight and it maintains the relationship to the Intersight Assist Device. Once added, Device Connector for the Managed Device type is started on the Intersight Assist and status related to it is maintained.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `device_type`:(string) Type of the Device such as VMware, Pure Storage supported by Intersight Assist. 
* `ignore_cert`:(bool) Ignore Certificates with protocol https for connecting to the Managed Device. It is not used for other protocols. 
* `is_enabled`:(bool) Device is Enabled/Disabled. 
* `management_address`:(string) Management address of the device. It can be IPv4, IPv6 or FQDN. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `port`:(int) Port to use for connecting to the Managed Device. Port is optional. If not specified, default port for protocol is used. 
* `protocol`:(string) Protocol to use for connecting to the Managed Device. 
