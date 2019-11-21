---
layout: "intersight"
page_title: "Intersight: intersight_telemetry_time_series"
sidebar_current: "docs-intersight-data-source-telemetryTimeSeries"
description: |-
Exposes a REST endpoint for performing queries against Druid time series data.
View Telemetry allows [POST of a Druid query](http://druid.io/docs/latest/querying/querying).
Manage Telemetry allows [READ of broker status](http://druid.io/docs/latest/operations/api-reference.html#broker).

---

# Data Source: intersight_telemetry_time_series
Exposes a REST endpoint for performing queries against Druid time series data.
View Telemetry allows [POST of a Druid query](http://druid.io/docs/latest/querying/querying).
Manage Telemetry allows [READ of broker status](http://druid.io/docs/latest/operations/api-reference.html#broker).

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
