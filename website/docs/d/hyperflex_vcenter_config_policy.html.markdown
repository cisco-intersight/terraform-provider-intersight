---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_vcenter_config_policy"
sidebar_current: "docs-intersight-data-source-hyperflexVcenterConfigPolicy"
description: |-
A policy specifying vCenter configuration.

---

# Data Source: intersight_hyperflex_vcenter_config_policy
A policy specifying vCenter configuration.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `data_center`:(string)The vCenter datacenter name.
* `description`:(string)Description of the policy.
* `hostname`:(string)The vCenter server FQDN or IP.
* `is_password_set`:(bool)
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `password`:(string)The password for authenticating with vCenter. Follow the corresponding password policy governed by vCenter.
* `sso_url`:(string)Overrides the default vCenter Single Sign-On URL. Do not specify unless instructed by Cisco TAC.
* `username`:(string)The vCenter username (e.g. administrator@vsphere.local).
