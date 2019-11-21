---
layout: "intersight"
page_title: "Intersight: intersight_testcrypt_token_api"
sidebar_current: "docs-intersight-data-source-testcryptTokenApi"
description: |-
Mo to test encryption token APIs.

---

# Data Source: intersight_testcrypt_token_api
Mo to test encryption token APIs.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `session_id`:(string)The session identifier to be used to create/read/update or delete Encryption token. This could be a user's (web) session id or api key id, etc.
