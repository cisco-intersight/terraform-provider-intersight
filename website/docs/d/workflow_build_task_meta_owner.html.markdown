---
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta_owner"
sidebar_current: "docs-intersight-data-source-workflowBuildTaskMetaOwner"
description: |-
Contains the list of dynamic workflow types that a microservice supports.

---

# Data Source: intersight_workflow_build_task_meta_owner
Contains the list of dynamic workflow types that a microservice supports.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `service`:(string)The microservice owner responsible for the tasks.
