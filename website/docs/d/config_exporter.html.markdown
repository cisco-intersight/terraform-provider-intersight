---
layout: "intersight"
page_title: "Intersight: intersight_config_exporter"
sidebar_current: "docs-intersight-data-source-configExporter"
description: |-
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
---

# Data Source: intersight_config_exporter
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `class_id`:(string)"The concrete type of this complex type. Its value must be the same as the 'objectType' property.\nThe OpenAPI document references this property as a discriminator value."
* `download_path`:(string)"Pre-signed URL to download the exported package, if the export operation has completed successfully. Regenerated during a GET request, if the existing pre-signed URL has expired."
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `name`:(string)"An identifier for the exporter instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
* `status`:(string)"Status of the export operation."
* `status_message`:(string)"Status message associated with failures or progress indication."
