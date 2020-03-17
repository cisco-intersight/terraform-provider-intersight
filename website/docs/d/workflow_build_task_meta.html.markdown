---
layout: "intersight"
page_title: "Intersight: intersight_workflow_build_task_meta"
sidebar_current: "docs-intersight-data-source-workflowBuildTaskMeta"
description: |-
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
---

# Data Source: intersight_workflow_build_task_meta
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"Name for the BuildTaskMeta instance used to created a dynamic workflow."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `src`:(string)"Microservice owner for the task in this workflow."
* `task_type`:(string)"The type of the task within this workflow."
* `workflow_type`:(string)"The type for the dynamic workflow."
