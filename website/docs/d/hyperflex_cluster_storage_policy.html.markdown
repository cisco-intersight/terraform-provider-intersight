---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster_storage_policy"
sidebar_current: "docs-intersight-data-source-hyperflexClusterStoragePolicy"
description: |-
A policy specifying HyperFlex cluster storage settings (optional).

---

# Data Source: intersight_hyperflex_cluster_storage_policy
A policy specifying HyperFlex cluster storage settings (optional).

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `disk_partition_cleanup`:(bool)If enabled, formats existing disk partitions (destroys all user data).
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `vdi_optimization`:(bool)Enable or disable VDI optimization (hybrid HyperFlex systems only).
