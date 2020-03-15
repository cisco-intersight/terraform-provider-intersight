---
layout: "intersight"
page_title: "Intersight: intersight_hyperflex_local_credential_policy"
sidebar_current: "docs-intersight-data-source-hyperflexLocalCredentialPolicy"
description: |-
A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.
---

# Data Source: intersight_hyperflex_local_credential_policy
A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `factory_hypervisor_password`:(bool)Indicates if Hypervisor password is the factory set default password. For HyperFlex Data Platform versions 3.0 or higher, enable this if the default password was not changed on the Hypervisor. It is required to supply a new custom Hypervisor password that will be applied to the Hypervisor during deployment. For HyperFlex Data Platform versions prior to 3.0 release, this setting has no effect and the default password will be used for initial install. The Hypervisor password should be changed after deployment.
* `hxdp_root_pwd`:(string)HyperFlex storage controller VM password must contain a minimum of 10 characters, with at least 1 lowercase, 1 uppercase, 1 numeric, and 1 of these -_@#$%^&*! special characters.
* `hypervisor_admin`:(string)Hypervisor administrator username must contain only alphanumeric characters. Use the root account for ESXi deployments.
* `hypervisor_admin_pwd`:(string)The ESXi root password. For HyperFlex Data Platform 3.0 or later, if the factory default password was not manually changed, you must set a new custom password. If the password was manually changed, you must not enable the factory default password property and provide the current hypervisor password. Note - All HyperFlex nodes require the same hypervisor password for installation. For HyperFlex Data Platform prior to 3.0, use the default password \"Cisco123\" for newly manufactured HyperFlex servers. A custom password should only be entered if hypervisor credentials were manually changed prior to deployment.
* `is_hxdp_root_pwd_set`:(bool)Indicates whether the value of the 'hxdpRootPwd' property has been set.
* `is_hypervisor_admin_pwd_set`:(bool)Indicates whether the value of the 'hypervisorAdminPwd' property has been set.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
