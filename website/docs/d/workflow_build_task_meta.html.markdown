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
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name for the BuildTaskMeta instance used to created a dynamic workflow.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `src`:(string)Microservice owner for the task in this workflow.
* `task_type`:(string)The type of the task within this workflow.
* `workflow_type`:(string)The type for the dynamic workflow.
