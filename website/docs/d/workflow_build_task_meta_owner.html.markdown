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
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `service`:(string)The microservice owner responsible for the tasks.
