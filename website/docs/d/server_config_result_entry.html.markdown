---
layout: "intersight"
page_title: "Intersight: intersight_server_config_result_entry"
sidebar_current: "docs-intersight-data-source-serverConfigResultEntry"
description: |-
The profile configuration (deploy, validation) results details information.

---

# Data Source: intersight_server_config_result_entry
The profile configuration (deploy, validation) results details information.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `completed_time`:(string)The completed time of the task in installer.
* `message`:(string)Localized message based on the locale setting of the user's context.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `owner_id`:(string)
* `state`:(string)Values  -- ok, ok-with-warning, errored.
* `type`:(string)Indicates if the result is reported during the logical model validation/resource allocation phase. or the configuration applying phase. Values -- validation, config.
