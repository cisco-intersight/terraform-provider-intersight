---
layout: "intersight"
page_title: "Intersight: intersight_config_importer"
sidebar_current: "docs-intersight-data-source-configImporter"
description: |-
All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
---

# Data Source: intersight_config_importer
All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `import_path`:(string)"The path to the archive in Intersight storage that has all the MOs\nto be imported."
* `import_source`:(string)"The source of the archive in Intersight storage that has all the MOs\nto be imported."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"An identifier for the importer instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `skip_integrity_checks`:(bool)"Specifies whether integrity checks must be skipped."
* `status`:(string)"Status of the import operation."
* `status_message`:(string)"Status message associated with failures or progress indication."
