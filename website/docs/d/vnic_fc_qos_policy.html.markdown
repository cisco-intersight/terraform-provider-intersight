---
layout: "intersight"
page_title: "Intersight: intersight_vnic_fc_qos_policy"
sidebar_current: "docs-intersight-data-source-vnicFcQosPolicy"
description: |-
A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.
---

# Data Source: intersight_vnic_fc_qos_policy
A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters you can also specify additional controls like burst and rate on the outgoing traffic.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `cos`:(int)"Class of Service to be associated to the traffic on the virtual interface."
* `description`:(string)"Description of the policy."
* `max_data_field_size`:(int)"The maximum size of the Fibre Channel frame payload bytes that the virtual interface supports."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete policy."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `rate_limit`:(int)"The value in Mbps to use for limiting the data rate on the virtual interface. Setting this to zero will turn rate limiting off."
