---
layout: "intersight"
page_title: "Intersight: intersight_workflow._workflow_info"
sidebar_current: "docs-intersight-data-source-workflow.WorkflowInfo"
description: |-
Contains information for a workflow execution which is a runtime instance of workflow.
---

# Data Source: intersight_workflow._workflow_info
Contains information for a workflow execution which is a runtime instance of workflow.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)"The action of the workflow such as start, cancel, retry, pause."
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `cleanup_time`:(string)"The time when the workflow info will be removed from database."
* `end_time`:(string)"The time when the workflow reached a final state."
* `failed_workflow_cleanup_duration`:(int)"The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database."
* `inst_id`:(string)"A workflow instance Id which is the unique identified for the workflow execution."
* `internal`:(bool)"Denotes if this workflow is internal and should be hidden from user view of running workflows."
* `last_action`:(string)"The last action that was issued on the workflow is saved in this field."
* `meta_version`:(int)"Version of the workflow metadata for which this workflow execution was started."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"A name of the workflow execution instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `progress`:(float)"This field indicates percentage of workflow task execution."
* `retry_from_task_name`:(string)"This field is applicable when Retry action is issued for a workflow which is in a final state. When this field is not specified then the workflow will retry from the start of the workflow. When this field is specified then the workflow will be retried from the specified task. The field should carry the task name which is the unique name of the task within the workflow. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration."
* `src`:(string)"The source microservice name which is the owner for this workflow."
* `start_time`:(string)"The time when the workflow was started for execution."
* `status`:(string)"A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED)."
* `success_workflow_cleanup_duration`:(int)"The duration in hours after which the workflow info for successful workflow will be removed from database."
* `trace_id`:(string)"The trace id to keep track of workflow execution."
* `type`:(string)"A type of the workflow (serverconfig, ansible_monitoring)."
* `user_id`:(string)"The user identifier which indicates the user that started this workflow."
* `wait_reason`:(string)"Denotes the reason workflow is in waiting status."
* `workflow_meta_type`:(string)"The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance."
* `workflow_task_count`:(int)"Total number of workflow tasks in this workflow."
