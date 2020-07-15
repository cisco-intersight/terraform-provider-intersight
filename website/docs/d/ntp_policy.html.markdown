
---
layout: "intersight"
page_title: "Intersight: intersight_ntp_policy"
sidebar_current: "docs-intersight-data-source-ntp-policy"
description: |-
Policy to configure the NTP Servers.
---

# Data Source: intersight_ntp._policy
Policy to configure the NTP Servers.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `description`:(string) Description of the policy. 
* `enabled`:(bool) State of NTP service on the endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `name`:(string) Name of the concrete policy. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `timezone`:(string) Timezone of services on the endpoint. 
