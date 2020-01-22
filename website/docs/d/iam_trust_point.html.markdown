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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
