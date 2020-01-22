---
layout: "intersight"
page_title: "Intersight: intersight_os_install"
sidebar_current: "docs-intersight-data-source-osInstall"
description: |-
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.

---

# Data Source: intersight_os_install
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)User provided description about the OS install configuration.
* `install_method`:(string)The install method to be used for OS installation - vMedia, iPXE. Only vMedia is supported as of now.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name of the OS install configuration.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
