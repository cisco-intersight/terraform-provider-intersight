---
layout: "intersight"
page_title: "Intersight: intersight_hcl_driver_image"
sidebar_current: "docs-intersight-data-source-hclDriverImage"
description: |-
Collection used to store the driver ISO urls for each server based on how it is managed.

---

# Data Source: intersight_hcl_driver_image
Collection used to store the driver ISO urls for each server based on how it is managed.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `driver_iso_url`:(string)URL of the driver ISO images.
* `management_type`:(string)Type of the UCS version indicating whether it is a UCSM release vesion or a IMC release.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `server_pid`:(string)Three part ID representing the server model as returned by UCSM/CIMC XML APIs.
