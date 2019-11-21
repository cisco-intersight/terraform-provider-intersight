---
layout: "intersight"
page_title: "Intersight: intersight_workflow_workflow_meta"
sidebar_current: "docs-intersight-data-source-workflowWorkflowMeta"
description: |-
Contains a workflow definition which is a sequence of tasks to execute.

---

# Data Source: intersight_workflow_workflow_meta
Contains a workflow definition which is a sequence of tasks to execute.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)The description for the workflow.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The name given to the workflow.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `schema_version`:(int)The Conductor schema version that decides what attribute can be supported.
* `src`:(string)The src is workflow owner service.
* `type`:(string)The type of workflow definition.
* `version`:(int)The version for the workflow so we can support multiple versions for the same workflow name.
* `wait_on_duplicate`:(bool)Parameter decides if workflows will wait for a duplicate to finish before starting a new one.
