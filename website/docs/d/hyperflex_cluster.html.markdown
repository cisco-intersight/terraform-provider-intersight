---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_cluster"
sidebar_current: "docs-intersight-data-source-hyperflexCluster"
description: |-
A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes
of the cluster.

---

# Data Source: intersight_hyperflex_cluster
A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes
of the cluster.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cluster_name`:(string)The name of this HyperFlex cluster.
* `cluster_type`:(int)The storage type of this cluster (All Flash or Hybrid).
* `cluster_uuid`:(string)The unique identifier for this HyperFlex cluster.
* `compute_node_count`:(int)The number of compute nodes that belong to this cluster.
* `converged_node_count`:(int)The number of converged nodes that belong to this cluster.
* `deployment_type`:(string)The deployment type of the HyperFlex cluster.The cluster can have one of the following configurations:1. Datacenter: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes on a single site.2. Stretched Cluster: The HyperFlex cluster consists of UCS Fabric Interconnect-attached nodes distributed across multiple sites.3. Edge: The HyperFlex cluster consists of 2-4 standalone nodes.If the cluster is running a HyperFlex Data Platform version less than 4.0 or if the deployment type cannot be determined,the deployment type is set as 'NA' (not available).
* `device_id`:(string)The unique identifier of the device registration that represents this HyperFlex cluster's connection to Intersight.
* `flt_aggr`:(int)The number of yellow (warning) and red (critical) alarms stored as an aggregate.The first 16 bits indicate the number of red alarms, and the last 16 bits contain the number of yellow alarms.
* `hx_version`:(string)The HyperFlex Data Platform version of this cluster.
* `hxdp_build_version`:(string)The version and build number of the HyperFlex Data Platform for this cluster.After a cluster upgrade, this version string will be updated on the next inventory cycle to reflectthe newly installed version.
* `hypervisor_type`:(string)The type of hypervisor running on this cluster.
* `hypervisor_version`:(string)The version of hypervisor running on this cluster.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `vm_count`:(int)The number of virtual machines present on this cluster.
