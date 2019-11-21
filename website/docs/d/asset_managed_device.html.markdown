---
layout: "intersight"
page_title: "Intersight: intersight_asset_managed_device"
sidebar_current: "docs-intersight-data-source-assetManagedDevice"
description: |-
Attributes for Managed Device in Intersight and it maintains the relationship to the Device Connector Manager (Intersight appliance platform type tagged with 'Intersight Assist') Device Registration. It is possible to add Managed Devices within an Account and later associated to DeviceConnectorManager. Once attached to Device Connector Manager, Device Connector for the Managed Device type is started on the OnPrem and status related to it is maintained in the ManagedDeviceStatus MO. If the Device Connector Manager is deleted, Managed Device continues to exist and can be associated to a different Device Connector Manager subsequently.

---

# Data Source: intersight_asset_managed_device
Attributes for Managed Device in Intersight and it maintains the relationship to the Device Connector Manager (Intersight appliance platform type tagged with 'Intersight Assist') Device Registration. It is possible to add Managed Devices within an Account and later associated to DeviceConnectorManager. Once attached to Device Connector Manager, Device Connector for the Managed Device type is started on the OnPrem and status related to it is maintained in the ManagedDeviceStatus MO. If the Device Connector Manager is deleted, Managed Device continues to exist and can be associated to a different Device Connector Manager subsequently.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_type`:(string)Type of the Device such as VMware, Pure Storage.
* `is_enabled`:(bool)Device is Enabled/Disabled.
* `management_address`:(string)Management address of the device. It can be IPv4, IPv6 or FQDN.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name defined by the administrator for easier identification.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
