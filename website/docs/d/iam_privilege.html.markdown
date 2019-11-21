---
layout: "intersight"
page_title: "Intersight: intersight_iam_privilege"
sidebar_current: "docs-intersight-data-source-iamPrivilege"
description: |-
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.

---

# Data Source: intersight_iam_privilege
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `hostname_prefix`:(string)The hostname prefix of the resource corresponding to this privilege. For example 'sentry' in https://sentry.intersight.com .
* `method`:(string)The API method on the rest resource corresponding to privilege. For example READ, CREATE, UPDATE etc.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the privilege reported by microservice.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `rest_path`:(string)The REST API path of the resource corresponding to this privilege. For example /v1/iam/Accounts, /v1/iam/Sessions.
* `url_prefix`:(string)The URL path prefix of the resource corresponding to this privilege. For example /devops/kibana, /devops/grafana etc.
