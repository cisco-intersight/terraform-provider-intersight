---
layout: "intersight"
page_title: "Intersight: intersight_iam_certificate"
sidebar_current: "docs-intersight-data-source-iamCertificate"
description: |-
Holds a certificate, signed by a CAcert.

---

# Data Source: intersight_iam_certificate
Holds a certificate, signed by a CAcert.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `status`:(string)Status of the certificate.
