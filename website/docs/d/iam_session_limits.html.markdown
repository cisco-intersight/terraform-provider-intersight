---
layout: "intersight"
page_title: "Intersight: intersight_iam_session_limits"
sidebar_current: "docs-intersight-data-source-iamSessionLimits"
description: |-
The session related configuration limits.

---

# Data Source: intersight_iam_session_limits
The session related configuration limits.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `idle_time_out`:(int)The idle timeout interval for the web session in seconds. The default value is 1800 seconds. When a session is not refreshed for this duration, backend will mark the session as idle and remove the session.
* `maximum_limit`:(int)The maximum number of sessions allowed in an account. The default value is 128.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `per_user_limit`:(int)The maximum number of sessions allowed per user. Default value is 32.
* `session_time_out`:(int)The session expiry duration in seconds. The default value is 57600 seconds or 16 hours.
