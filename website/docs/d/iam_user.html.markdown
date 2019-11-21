---
layout: "intersight"
page_title: "Intersight: intersight_iam_user"
sidebar_current: "docs-intersight-data-source-iamUser"
description: |-
The Intersight account user.

---

# Data Source: intersight_iam_user
The Intersight account user.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `client_ip_address`:(string)IP address from which the user last logged in to Intersight.
* `email`:(string)Email of the user. Users are added to Intersight using the email configured in the IdP.
* `first_name`:(string)First name of the user. This field is populated from the IdP attributes received after authentication.
* `last_login_time`:(string)Last successful login time for user.
* `last_name`:(string)Last name of the user. This field is populated from the IdP attributes received after authentication.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)UserID as configured in the IdP.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `user_type`:(string)Type of the User. If a user is added manually by specifying the email address, or has logged in using groups, based on the IdP attributes received during authentication. If added manually, the user type will be static, otherwise dynamic.
