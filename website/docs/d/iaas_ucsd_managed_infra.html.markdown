---
layout: "intersight"
page_title: "Intersight: intersight_iaas_ucsd_managed_infra"
sidebar_current: "docs-intersight-data-source-iaasUcsdManagedInfra"
description: |-
Describes about UCSD Managed infrastructure statistics.

---

# Data Source: intersight_iaas_ucsd_managed_infra
Describes about UCSD Managed infrastructure statistics.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `advanced_catalog_count`:(int)Total advanced catalogs in UCSD.
* `bm_catalog_count`:(int)Total bare metal catalogs in UCSD.
* `container_catalog_count`:(int)Total service container catalogs in UCSD.
* `esxi_host_count`:(int)Total ESXi hosts in UCSD.
* `external_group_count`:(int)Total external (Ldap) groups in UCSD.
* `hyperv_host_count`:(int)Total HyperV hosts in UCSD.
* `local_group_count`:(int)Total local groups in UCSD.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `standard_catalog_count`:(int)Total standard catalogs in UCSD.
* `user_count`:(int)Total user accounts in UCSD.
* `vm_count`:(int)Total VMs in UCSD.
* `vdc_count`:(int)Total virtual datacenters in UCSD.
