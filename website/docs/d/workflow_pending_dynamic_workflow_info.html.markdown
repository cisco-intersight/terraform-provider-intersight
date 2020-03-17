---
layout: "intersight"
page_title: "Intersight: intersight_workflow_pending_dynamic_workflow_info"
sidebar_current: "docs-intersight-data-source-workflowPendingDynamicWorkflowInfo"
description: |-
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
---

# Data Source: intersight_workflow_pending_dynamic_workflow_info
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"A name for the pending dynamic workflow."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `src`:(string)"The src is workflow owner service."
* `status`:(string)"The current status of the PendingDynamicWorkflowInfo."
* `wait_on_duplicate`:(bool)"When set to true workflow engine will wait for a duplicate to finish before starting a new one."
* `workflow_key`:(string)"This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup."
