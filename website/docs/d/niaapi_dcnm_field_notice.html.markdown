---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_field_notice"
sidebar_current: "docs-intersight-data-source-niaapiDcnmFieldNotice"
description: |-
The field notice reporting bug and related software or hardware for DCNM.

---

# Data Source: intersight_niaapi_dcnm_field_notice
The field notice reporting bug and related software or hardware for DCNM.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `bugid`:(string)Bug Id associated with this notice.
* `field_notice_desc`:(string)Field notice Description.
* `field_notice_id`:(string)Fieldnotice Id of this notice.
* `field_notice_url`:(string)Field notice URL link to the notice webpage.
* `headline`:(string)The headline of this field notice.
* `hwpid`:(string)Hardware PID for affected models.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `sw_release`:(string)Software Release number for affected versions.
* `workaround_url`:(string)URL of workaround of this notice.
