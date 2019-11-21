---
layout: "intersight"
page_title: "Intersight: intersight_terminal_audit_log"
sidebar_current: "docs-intersight-data-source-terminalAuditLog"
description: |-
Audit log of remote terminal user sessions.

---

# Data Source: intersight_terminal_audit_log
Audit log of remote terminal user sessions.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `end_time`:(string)The time the terminal was closed. If terminal has not closed, value is zero time.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `start_time`:(string)The time the terminal session was opened.
