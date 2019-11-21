---
layout: "intersight"
page_title: "Intersight: intersight_hcl_exempted_catalog"
sidebar_current: "docs-intersight-data-source-hclExemptedCatalog"
description: |-
Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.

---

# Data Source: intersight_hcl_exempted_catalog
Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `comments`:(string)Reason for the exemption.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A unique descriptive name of the exemption.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `os_vendor`:(string)Vendor of the Operating System.
* `os_version`:(string)Version of the Operating system.
* `processor_name`:(string)Name of the processor supported for the server.
* `product_type`:(string)Type of the product/adapter say PT for Pass Through controllers.
* `server_pid`:(string)Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
* `ucs_version`:(string)Version of the UCS software.
* `version_type`:(string)Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
