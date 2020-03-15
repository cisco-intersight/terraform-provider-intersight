---
layout: "intersight"
page_title: "Intersight: intersight_kvm_policy"
sidebar_current: "docs-intersight-data-source-kvmPolicy"
description: |-
Policy to configure KVM Launch settings.
---

# Data Source: intersight_kvm_policy
Policy to configure KVM Launch settings.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `description`:(string)Description of the policy.
* `enable_local_server_video`:(bool)If enabled, displays KVM session on any monitor attached to the server.
* `enable_video_encryption`:(bool)If enabled, encrypts all video information sent through KVM.
* `enabled`:(bool)State of the vKVM service on the endpoint.
* `maximum_sessions`:(int)The maximum number of concurrent KVM sessions allowed.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `remote_port`:(int)The port used for KVM communication.
