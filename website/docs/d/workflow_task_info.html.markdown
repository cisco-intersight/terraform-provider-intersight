---
layout: "intersight"
page_title: "Intersight: intersight_workflow_task_info"
sidebar_current: "docs-intersight-data-source-workflowTaskInfo"
description: |-
Task instance which represents the run time instance of a task within a workflow.

---

# Data Source: intersight_workflow_task_info
Task instance which represents the run time instance of a task within a workflow.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)The task description and this is the description that was added when the task was included into the workflow.
* `end_time`:(string)The time stamp when the task reached a final state.
* `failure_reason`:(string)Description of the reason why the task failed.
* `inst_id`:(string)The current running task instance Id.
* `internal`:(bool)Denotes whether or not this is an internal task.  Internal tasks will be hidden from the UI within a workflow.
* `label`:(string)User friendly short label to describe this task instance in the workflow.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Task definition name which specifies the task type.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `ref_name`:(string)The task reference name to ensure its unique inside a workflow.
* `retry_count`:(int)A counter for number of retries.
* `start_time`:(string)The time stamp when the task started execution.
* `status`:(string)The status of the task and this will specify if the task is running or has reached a final state.
