---
layout: "intersight"
page_title: "Intersight: intersight_deviceconnector_policy"
sidebar_current: "docs-intersight-data-source-deviceconnectorPolicy"
description: |-
Device Connector Policy.

---

# Data Source: intersight_deviceconnector_policy
Device Connector Policy.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `lockout_enabled`:(bool)Enables configuration lockout on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
