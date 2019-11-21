---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_apic_release_recommend"
sidebar_current: "docs-intersight-data-source-niaapiApicReleaseRecommend"
description: |-
The recommend version information for each release on APIC.

---

# Data Source: intersight_niaapi_apic_release_recommend
The recommend version information for each release on APIC.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `cll`:(string)Current long-lived release.
* `crr`:(string)Customer recommended releases.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `pid`:(string)Hardware model identificator.
* `ull`:(string)Upcoming long-lived release.
