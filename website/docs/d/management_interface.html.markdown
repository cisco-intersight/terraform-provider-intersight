---
layout: "intersight"
page_title: "Intersight: intersight_management_interface"
sidebar_current: "docs-intersight-data-source-managementInterface"
description: |-
Interface that provides access to the management controller.

---

# Data Source: intersight_management_interface
Interface that provides access to the management controller.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_mo_id`:(string)
* `dn`:(string)
* `gateway`:(string)Default gateway for the interface.
* `host_name`:(string)Hostname configured for the interface.
* `ip_address`:(string)IP address of the interface.
* `mac_address`:(string)MAC address configured for the interface.
* `mask`:(string)Netmask for the interface.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rn`:(string)
* `virtual_host_name`:(string)Virtual hostname configured for the interface in case of clustered environment.
