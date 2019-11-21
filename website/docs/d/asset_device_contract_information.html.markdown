---
layout: "intersight"
page_title: "Intersight: intersight_asset_device_contract_information"
sidebar_current: "docs-intersight-data-source-assetDeviceContractInformation"
description: |-
Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.

---

# Data Source: intersight_asset_device_contract_information
Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `contract_status`:(string)Calculated contract status that is derived based on the service line status and contract end date. It is different from serviceLineStatus property. serviceLineStatus gives us ACTIVE, OVERDUE, EXPIRED. These are transformed into Active, Expiring Soon and Not Covered.
* `covered_product_line_end_date`:(string)End date of the covered product line. The coverage end date is fetched from Cisco SN2INFO API.
* `device_id`:(string)Unique identifier of the Cisco device. This information is used to query Cisco APIx SN2INFO and CCWR databases.
* `device_type`:(string)Type used to classify the device in Cisco Intersight. Currently supported values are Server and FabricInterconnect. This will be expanded to support more types in future.
* `is_valid`:(bool)Validates if the device is a genuine Cisco device. Validated is done using the Cisco SN2INFO APIs.
* `item_type`:(string)Item type of this specific Cisco device. example \"Chassis\".
* `maintenance_purchase_order_number`:(string)Maintenance purchase order number for the Cisco device.
* `maintenance_sales_order_number`:(string)Maintenance sales order number for the Cisco device.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `platform_type`:(string)The platform type of the Cisco device.
* `purchase_order_number`:(string)Purchase order number for the Cisco device. It is a unique number assigned for every purchase.
* `sales_order_number`:(string)Sales order number for the Cisco device. It is a unique number assigned for every sale.
* `service_description`:(string)The type of service contract that covers the Cisco device.
* `service_end_date`:(string)End date for the Cisco service contract that covers this Cisco device.
* `service_level`:(string)The type of service contract that covers the Cisco device.
* `service_sku`:(string)The SKU of the service contract that covers the Cisco device.
* `service_start_date`:(string)Start date for the Cisco service contract that covers this Cisco device.
* `warranty_end_date`:(string)End date for the warranty that covers the Cisco device.
* `warranty_type`:(string)Type of warranty that covers the Cisco device.
