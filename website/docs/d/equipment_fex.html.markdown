
---
layout: "intersight"
page_title: "Intersight: intersight_equipment_fex"
sidebar_current: "docs-intersight-data-source-equipment-fex"
description: |-
Fabric Extender which can mutiplex traffic from the host facing ports.
---

# Data Source: intersight_equipment._fex
Fabric Extender which can mutiplex traffic from the host facing ports.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string) The concrete type of this complex type. Its value must be the same as the 'objectType' property.The OpenAPI document references this property as a discriminator value. 
* `connection_status`:(string) Connectivity Status of FEX/IOM to Switch - A or B or AB. 
* `description`:(string) This field is to provide description for the iocard module model. 
* `device_mo_id`:(string) The database identifier of the registered device of an object. 
* `discovery_state`:(string) Discovery state of IO card or fabric extender. 
* `dn`:(string) The Distinguished Name unambiguously identifies an object in the system. 
* `model`:(string) This field identifies the model of the given component. 
* `module_id`:(int) Module Identifier for the IO module. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `object_type`:(string) The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path. 
* `oper_state`:(string) Operational state of IO card or fabric extender. 
* `part_number`:(string) Part Number identifier for the IO module. 
* `pid`:(string) This field identifies the Product ID for the IO module. 
* `presence`:(string) This field identifies the Presence state of the IO card module. 
* `product_name`:(string) This field identifies the Product Name for the iocard module model. 
* `revision`:(string) This field identifies the revision of the given component. 
* `rn`:(string) The Relative Name uniquely identifies an object within a given context. 
* `serial`:(string) This field identifies the serial of the given component. 
* `sku`:(string) This field identifies the Stock Keeping Unit for the IO card module. 
* `vendor`:(string) This field identifies the vendor of the given component. 
* `version`:(string) This field identifies the version of the IO card module. 
* `vid`:(string) This field identifies the Vendor ID for the IO card module. 
