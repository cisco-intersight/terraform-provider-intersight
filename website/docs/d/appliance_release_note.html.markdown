---
layout: "intersight"
page_title: "Intersight: intersight_appliance_release_note"
sidebar_current: "docs-intersight-data-source-applianceReleaseNote"
description: |-
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
---

# Data Source: intersight_appliance_release_note
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `version`:(string)"Version number of the pending upgrade."
