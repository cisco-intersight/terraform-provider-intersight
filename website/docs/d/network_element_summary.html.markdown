---
layout: "intersight"
page_title: "Intersight: intersight_network_element_summary"
sidebar_current: "docs-intersight-data-source-networkElementSummary"
description: |-
View MO which aggregates information pertaining to a network element from mutiple MOs.

---

# Data Source: intersight_network_element_summary
View MO which aggregates information pertaining to a network element from mutiple MOs.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `admin_inband_interface_state`:(string)
* `device_mo_id`:(string)
* `dn`:(string)
* `fault_summary`:(int)
* `firmware`:(string)Running firmware information.
* `ipv4_address`:(string)IP version 4 address is saved in this property.
* `inband_ip_address`:(string)
* `inband_ip_gateway`:(string)
* `inband_ip_mask`:(string)
* `inband_vlan`:(int)
* `model`:(string)This field identifies the model of the given component.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the ElementSummary object is saved in this property.
* `num_ether_ports`:(int)Total number of Ethernet ports.
* `num_ether_ports_configured`:(int)Total number of configured Ethernet ports.
* `num_ether_ports_link_up`:(int)Total number of Ethernet ports which are UP.
* `num_expansion_modules`:(int)Total number of expansion modules.
* `num_fc_ports`:(int)Total number of FC ports.
* `num_fc_ports_configured`:(int)Total number of configured FC ports.
* `num_fc_ports_link_up`:(int)Total number of FC ports which are UP.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `out_of_band_ip_address`:(string)
* `out_of_band_ip_gateway`:(string)
* `out_of_band_ip_mask`:(string)
* `out_of_band_mac`:(string)
* `revision`:(string)
* `rn`:(string)
* `serial`:(string)This field identifies the serial of the given component.
* `source_object_type`:(string)Specifies the source object type for View MO.
* `switch_id`:(string)
* `vendor`:(string)This field identifies the vendor of the given component.
* `version`:(string)Version holds the firmware version related information.
