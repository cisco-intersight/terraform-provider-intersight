---
layout: "intersight"
page_title: "Intersight: intersight_firmware_upgrade"
sidebar_current: "docs-intersight-data-source-firmwareUpgrade"
description: |-
Firmware upgrade operation that downloads the image from Cisco/appliance/user provided HTTP repository or use the image from a network share and upgrade. The direct download is used for upgrade to use the image from Cisco repository or appliance repository. The network share is used for upgrade to use the image from a network share in user data center.

---

# Data Source: intersight_firmware_upgrade
Firmware upgrade operation that downloads the image from Cisco/appliance/user provided HTTP repository or use the image from a network share and upgrade. The direct download is used for upgrade to use the image from Cisco repository or appliance repository. The network share is used for upgrade to use the image from a network share in user data center.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `upgrade_type`:(string)Desired upgrade mode to choose either direct download based upgrade or network share upgrade.
