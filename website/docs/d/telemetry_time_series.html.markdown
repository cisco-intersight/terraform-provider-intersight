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
* `moid`:(string)"The unique identifier of this Managed Object instance."
* `object_type`:(string)"The fully-qualified type of this managed object, i.e. the class name.\nThis property is optional. The ObjectType is implied from the URL path.\nIf specified, the value of objectType must match the class name specified in the URL path."
