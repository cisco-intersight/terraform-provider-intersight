---
layout: "intersight"
page_title: "Intersight: intersight_tam_advisory_instance"
sidebar_current: "docs-intersight-data-source-tamAdvisoryInstance"
description: |-
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.

---

# Data Source: intersight_tam_advisory_instance
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `affected_object_moid`:(string)Moid of the Intersight MO affected by the alert.
* `affected_object_type`:(string)Object type of the Intersight MO affected by the alert.
* `last_state_change_time`:(string)Timestamp when a state change was observed on this advisory instnace.
* `last_verified_time`:(string)Timestamp when this advisory was last evaluated.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `state`:(string)Current state of the advisory instance (Active/Cleared/Unknown etc.).
