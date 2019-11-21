---
layout: "intersight"
page_title: "Intersight: intersight_license_customer_op"
sidebar_current: "docs-intersight-data-source-licenseCustomerOp"
description: |-
Customer operation object to refresh the registration or re-authenticate, pre-created.

---

# Data Source: intersight_license_customer_op
Customer operation object to refresh the registration or re-authenticate, pre-created.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `deregister_device`:(bool)Trigger de-registration/disable.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `renew_authorization`:(bool)Trigger renew authorization.
* `renew_id_certificate`:(bool)Trigger renew registration.
* `show_agent_tech_support`:(bool)Trigger show tech support feature.
