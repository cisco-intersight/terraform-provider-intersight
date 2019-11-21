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
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)A name for the pending dynamic workflow.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `src`:(string)The src is workflow owner service.
* `status`:(string)The current status of the PendingDynamicWorkflowInfo.
* `wait_on_duplicate`:(bool)When set to true workflow engine will wait for a duplicate to finish before starting a new one.
* `workflow_key`:(string)This key contains workflow, initiator and target name. Workflow engine uses the key to do workflow dedup.
