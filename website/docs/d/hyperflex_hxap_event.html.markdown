
---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_hxap_event"
sidebar_current: "docs-intersight-data-source-hyperflex-hxap-event"
description: |-
Events associated with HyperFlex Application Platform compute cluster.
---

# Data Source: intersight_hyperflex._hxap_event
Events associated with HyperFlex Application Platform compute cluster.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `first_time`:(string) First timestamp of the event occurrence. 
* `identity`:(string) Internally generated identity (UUID) of this event. 
* `last_time`:(string) Last timestamp of the event occurrence. 
* `message`:(string) Full description of the event. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `owner_name`:(string) Name of the owner with which event is associated. 
* `owner_type`:(string) Type of the object with which event is associated (Host, Cluster, VM). 
* `owner_uuid`:(string) UUID of the owner with which event is associated. 
* `severity`:(string) Severity level of the event (Info/Warning/Critical). 
