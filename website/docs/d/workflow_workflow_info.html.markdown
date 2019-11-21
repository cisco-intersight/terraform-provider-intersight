---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_info"
sidebar_current: "docs-intersight-data-source-workflowWorkflowInfo"
description: |-
Contains information for a workflow execution which is a runtime instance of workflow.

---

# Data Source: intersight_workflow_workflow_info
Contains information for a workflow execution which is a runtime instance of workflow.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `action`:(string)The action of the workflow such as start, cancel, retry, pause.
* `cleanup_time`:(string)The time when the workflow info will be removed from database.
* `end_time`:(string)The time when the workflow reached a final state.
* `failed_workflow_cleanup_duration`:(int)The duration in hours after which the workflow info for failed, terminated or timed out workflow will be removed from database.
* `inst_id`:(string)A workflow instance Id which is the unique identified for the workflow execution.
* `internal`:(bool)Denotes if this workflow is internal and should be hidden from user view of running workflows.
* `meta_version`:(int)Version of the workflow metadata for which this workflow execution was started.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A name of the workflow execution instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `progress`:(float)This field indicates percentage of workflow task execution.
* `src`:(string)The source microservice name which is the owner for this workflow.
* `start_time`:(string)The time when the workflow was started for execution.
* `status`:(string)A status of the workflow (RUNNING, WAITING, COMPLETED, TIME_OUT, FAILED).
* `success_workflow_cleanup_duration`:(int)The duration in hours after which the workflow info for successful workflow will be removed from database.
* `type`:(string)A type of the workflow (serverconfig, ansible_monitoring).
* `user_id`:(string)The user identifier which indicates the user that started this workflow.
* `wait_reason`:(string)Denotes the reason workflow is in waiting status.
* `workflow_meta_type`:(string)The type of workflow meta. Derived from the workflow meta that is used to launch this workflow instance.
* `workflow_task_count`:(int)Total number of workflow tasks in this workflow.
