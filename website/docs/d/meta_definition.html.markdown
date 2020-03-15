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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `parent_class`:(string)The fully-qualified name of the parent metaclass in the class inheritance hierarchy.
* `permission_supported`:(bool)Boolean flag to specify whether instances of this class type can be specified in permissions for instance based access control. Permissions can be created for entire Intersight account or to a subset of resources (instance based access control). In the first release, permissions are supported for entire account or for a subset of organizations.
* `rbac_resource`:(bool)Boolean flag to specify whether instances of this class type can be assigned to resource groups that are part of an organization for access control. Inventoried physical/logical objects which needs access control should have rbacResource=yes. These objects are not part of any organization by default like device registrations and should be assigned to organizations for access control. Profiles, policies, workflow definitions which are created by specifying organization need not have this flag set.
* `rest_path`:(string)Restful URL path for the meta.
* `version`:(string)The version of the service that defines the meta-data.
