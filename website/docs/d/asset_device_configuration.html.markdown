---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_configuration"
sidebar_current: "docs-intersight-data-source-assetDeviceConfiguration"
description: |-
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.

---

# Data Source: intersight_asset_device_configuration
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `local_configuration_locked`:(bool)Specifies whether configuration through the platforms local management interface has been disabled, with only configuration through the Intersight service enabled.
* `log_level`:(string)The log level of the device connector service.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
