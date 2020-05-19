---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex._node_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflex.NodeConfigPolicy"
description: |-
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
---

# Data Source: intersight_hyperflex._node_config_policy
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `description`:(string)"Description of the policy."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `node_name_prefix`:(string)"The node name prefix that is used to automatically generate the default hostname for each server.\nA dash (-) will be appended to the prefix followed by the node number to form a hostname.\nThis default naming scheme can be manually overridden in the node configuration.\nThe maximum length of a prefix is 60, must only contain alphanumeric characters or dash (-), and must\nstart with an alphanumeric character."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
