---
layout: "intersight"
page_title: "Intersight: intersight_ssh_policy"
sidebar_current: "docs-intersight-data-source-sshPolicy"
description: |-
Secure shell policy on the endpoint.
---

# Data Source: intersight_ssh_policy
Secure shell policy on the endpoint.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of SSH service on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `port`:(int)Port used for secure shell access.
* `timeout`:(int)Number of seconds to wait before the system considers a SSH request to have timed out.
