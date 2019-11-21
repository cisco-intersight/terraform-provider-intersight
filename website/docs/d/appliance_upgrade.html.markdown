---
layout: "intersight"
page_title: "Intersight: intersight_appliance_upgrade"
sidebar_current: "docs-intersight-data-source-applianceUpgrade"
description: |-
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.

---

# Data Source: intersight_appliance_upgrade
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `active`:(bool)Indicates if the software upgrade is active or not.
* `auto_created`:(bool)Indicates that the request was automatically created by the system.
* `description`:(string)Description of the software upgrade.
* `elapsed_time`:(int)Elapsed time in seconds during the software upgrade.
* `end_time`:(string)End date of the software upgrade.
* `fingerprint`:(string)Software upgrade manifest's fingerprint.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `start_time`:(string)Start date of the software upgrade. UI can modify startTime to re-schedule an upgrade.
* `status`:(string)Status of the Intersight Appliance's software upgrade.
* `version`:(string)Software upgrade manifest's version.
