---
layout: "intersight"
page_title: "Intersight: intersight_vmedia_policy"
sidebar_current: "docs-intersight-data-source-vmediaPolicy"
description: |-
Policy to configure virtual media settings on endpoint.

---

# Data Source: intersight_vmedia_policy
Policy to configure virtual media settings on endpoint.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of the Virtual Media service on the endpoint.
* `encryption`:(bool)If enabled, allows encryption of all Virtual Media communications.
* `low_power_usb`:(bool)If enabled, the virtual drives appear on the boot selection menu after mapping the image and rebooting the host.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
