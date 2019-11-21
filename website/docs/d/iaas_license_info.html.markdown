---
layout: "intersight"
page_title: "Intersight: intersight_iaas_license_info"
sidebar_current: "docs-intersight-data-source-iaasLicenseInfo"
description: |-
Describes about license info currently available in UCSD.

---

# Data Source: intersight_iaas_license_info
Describes about license info currently available in UCSD.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `license_expiration_date`:(string)Licese expiration date.
* `license_type`:(string)License type of UCSD whether it is EVAL/Permanent/Subscription..
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
