---
layout: "intersight"
page_title: "Intersight: intersight_niaapi_dcnm_cco_post"
sidebar_current: "docs-intersight-data-source-niaapiDcnmCcoPost"
description: |-
The post reporting a new release is available for DCNM.

---

# Data Source: intersight_niaapi_dcnm_cco_post
The post reporting a new release is available for DCNM.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `post_date`:(string)The date when this new release notice is posted.
* `post_type`:(string)The document type of this post.
* `postid`:(string)Identificator of this inbox post.
* `revision`:(string)Revision number of this notice.
