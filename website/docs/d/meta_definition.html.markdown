---
layout: "intersight"
page_title: "Intersight: intersight_meta_definition"
sidebar_current: "docs-intersight-data-source-metaDefinition"
description: |-
The meta-data of managed objects and complex types.

---

# Data Source: intersight_meta_definition
The meta-data of managed objects and complex types.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `is_concrete`:(bool)Boolean flag to specify whether the meta class is a concrete class or not.
* `meta_type`:(string)Indicates whether the meta class is a complex type or managed object.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)The fully-qualified class name of the Managed Object or complex type. For example, \"compute:Blade\" where the Managed Object is \"Blade\" and the package is 'compute'.
* `namespace`:(string)The namespace of the meta.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `parent_class`:(string)The fully-qualified name of the parent metaclass in the class inheritance hierarchy.
* `rest_path`:(string)Restful URL path for the meta.
* `version`:(string)The version of the service that defines the meta-data.
