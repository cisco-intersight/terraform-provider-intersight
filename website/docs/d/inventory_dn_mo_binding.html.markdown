---
layout: "intersight"
page_title: "Intersight: intersight_inventory_dn_mo_binding"
sidebar_current: "docs-intersight-data-source-inventoryDnMoBinding"
description: |-
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
---

# Data Source: intersight_inventory_dn_mo_binding
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `dn`:(string)"Distinguished Name of the UCSM object."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `target_mo_id`:(string)"The property targetMoId represents the Binding target's MoId."
* `target_mo_type`:(string)"The property targetMoType represents the Binding target's Mo type."
