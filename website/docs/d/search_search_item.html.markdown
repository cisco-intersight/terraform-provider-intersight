---
layout: "intersight"
page_title: "Intersight: intersight_search_search_item"
sidebar_current: "docs-intersight-data-source-searchSearchItem"
description: |-
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.

---

# Data Source: intersight_search_search_item
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
