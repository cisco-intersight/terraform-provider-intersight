---
layout: "intersight"
page_title: "Intersight: intersight_sol_policy"
sidebar_current: "docs-intersight-data-source-solPolicy"
description: |-
Policy for configuring Serial Over LAN settings on endpoint.

---

# Data Source: intersight_sol_policy
Policy for configuring Serial Over LAN settings on endpoint.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `baud_rate`:(int)Baud Rate used for Serial Over LAN communication.
* `com_port`:(string)Serial port through which the system routes Serial Over LAN communication. This field is available only on some Cisco UCS C-Series servers. If it is unavailable, the server uses COM port 0 by default.
* `description`:(string)Description of the policy.
* `enabled`:(bool)State of Serial Over LAN service on the endpoint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `ssh_port`:(int)SSH port used to access Serial Over LAN directly. Enables bypassing Cisco IMC shell to provide direct access to Serial Over LAN.
