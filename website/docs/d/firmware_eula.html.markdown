---
layout: "intersight"
page_title: "Intersight: intersight_firmware_eula"
sidebar_current: "docs-intersight-data-source-firmwareEula"
description: |-
EULA (End User License Agreement) acceptance status for an account to use cisco.com, to download software.

---

# Data Source: intersight_firmware_eula
EULA (End User License Agreement) acceptance status for an account to use cisco.com, to download software.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `accepted`:(bool)EULA acceptance status for the account.
* `content`:(string)EULA acceptance form content provided by cisco.com.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
