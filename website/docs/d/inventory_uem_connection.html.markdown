
---
layout: "intersight"
page_title: "Intersight: intersight_inventory_uem_connection"
sidebar_current: "docs-intersight-data-source-inventory-uem-connection"
description: |-
UemConnection MO maintains the current Uem connection status for a given UEM endpoint.
---

# Data Source: intersight_inventory._uem_connection
UemConnection MO maintains the current Uem connection status for a given UEM endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connection_status`:(string) Connections status of Uem endpoint. 
* `ep_type`:(string) Type of Uem endpoint (BMC, CMC or VIC). 
* `ip`:(string) The IP address of the Uem endpoint. 
* `member_identity`:(string) The member identity of the UEM connection being inventoried. 
* `model`:(string) The model of the Uem endpoint. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `serial`:(string) The serial number of the Uem endpoint. 
* `target_mo_id`:(string) The MoId address of the Uem endpoint. 
* `vendor`:(string) The vendor of the Uem endpoint. 
