---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_definition"
sidebar_current: "docs-intersight-data-source-tamAdvisoryDefinition"
description: |-
An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
---

# Data Source: intersight_tam_advisory_definition
An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advisory_id`:(string)"Cisco generated identifier for the published security advisory."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `date_published`:(string)"Date when the security advisory was first published by Cisco."
* `date_updated`:(string)"Date when the security advisory was last updated by Cisco."
* `description`:(string)"Brief description of the advisory details."
* `external_url`:(string)"A link to an external URL describing security Advisory in more details."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"A user defined name for the Intersight Advisory."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `recommendation`:(string)"Recommended action to resolve the security advisory."
* `state`:(string)"Current state of the advisory."
* `type`:(string)"The type (field notice, security advisory etc.) of Intersight advisory."
* `version`:(string)"Cisco assigned advisory version after latest revision."
* `workaround`:(string)"Workarounds available for the advisory."
