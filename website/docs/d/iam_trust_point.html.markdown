---
layout: "intersight"
page_title: "Intersight: intersight_iam_trust_point"
sidebar_current: "docs-intersight-data-source-iamTrustPoint"
description: |-
To affirm the identity of trusted source.
Allows import of third-party CA certificates in X.509 (CER) format.
It can be a root CA or an trust chain that leads to a root CA.

---

# Data Source: intersight_iam_trust_point
To affirm the identity of trusted source.
Allows import of third-party CA certificates in X.509 (CER) format.
It can be a root CA or an trust chain that leads to a root CA.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `chain`:(string)The certificate information for this trusted point. The certificate must be in Base64 encoded X.509 (CER) format.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
