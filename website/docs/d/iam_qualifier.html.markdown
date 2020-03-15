---
layout: "intersight"
page_title: "Intersight: intersight_iam_qualifier"
sidebar_current: "docs-intersight-data-source-iamQualifier"
description: |-
The qualifier defines how a user qualifies to be part of a user group.
---

# Data Source: intersight_iam_qualifier
The qualifier defines how a user qualifies to be part of a user group.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the SAML attribute used to qualify a user to user group. By default this is memberOf attribute in SAML assertion.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
