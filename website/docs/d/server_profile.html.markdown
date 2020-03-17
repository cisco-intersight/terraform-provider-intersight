---
layout: "intersight"
page_title: "Intersight: intersight_server_profile"
sidebar_current: "docs-intersight-data-source-serverProfile"
description: |-
A profile specifying configuration settings for a physical server.
---

# Data Source: intersight_server_profile
A profile specifying configuration settings for a physical server.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)"User initiated action. Each profile type has its own supported actions. For HyperFlex cluster profile, the supported actions are -- Validate, Deploy, Continue, Retry, Abort, Unassign For server profile, the support actions are -- Deploy, Unassign."
* `description`:(string)"Description of the profile."
* `is_pmc_deployed_secure_passphrase_set`:(bool)
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name of the concrete profile."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `pmc_deployed_secure_passphrase`:(string)"Secure passphrase that is already deployed on all the Persistent Memory Modules on the server. This deployed passphrase is required during deploy of server profile if secure passphrase is changed or security is disabled in the attached persistent memory policy."
* `type`:(string)"Defines the type of the profile. Accepted value is instance."
