---
layout: "intersight"
page_title: "Intersight: intersight_forecast_instance"
sidebar_current: "docs-intersight-data-source-forecastInstance"
description: |-
Entity representing forecast result for instance of managed object, ie, data source.
---

# Data Source: intersight_forecast_instance
Entity representing forecast result for instance of managed object, ie, data source.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `device_id`:(string)The Moid of the Intersight managed device instance for which regression model is derived.
* `full_cap_days`:(int)The number of days remaining before the device reaches its full functional capacity.
* `metric_name`:(string)The name of the metric for which regression model is generated.
* `min_days_for_forecast`:(int)The minimum number of days the HyperFlex cluster should be up for computing forecast.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `threshold_days`:(int)The number of days remaining before the device reaches the specified threshold for the metric as defined in definition.
