---
layout: "intersight"
page_title: "Intersight: intersight_server_config_change_detail"
sidebar_current: "docs-intersight-data-source-serverConfigChangeDetail"
description: |-
The configuration change details are captured here.

---

# Data Source: intersight_server_config_change_detail
The configuration change details are captured here.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `message`:(string)Detailed description of the config change.
* `mod_status`:(string)Modification status of the mo in this config change.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
