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
* `admin_inband_interface_state`:(string)The administrative state of the network Element inband management interface.
* `device_mo_id`:(string)
* `dn`:(string)The Distinguished Name unambiguously identifies an object in the system.
* `fault_summary`:(int)
* `firmware`:(string)Running firmware information.
* `ipv4_address`:(string)IP version 4 address is saved in this property.
* `inband_ip_address`:(string)The IP address of the network Element inband management interface.
* `inband_ip_gateway`:(string)The default gateway of the network Element inband management interface.
* `inband_ip_mask`:(string)The network mask of the network Element inband management interface.
* `inband_vlan`:(int)The VLAN ID of the network Element inband management interface.
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
* `object_type`:(string)The fully-qualified type of this managed object, i.e. the class name.This property is optional. The ObjectType is implied from the URL path.If specified, the value of objectType must match the class name specified in the URL path.
* `out_of_band_ip_address`:(string)The IP address of the network Element out-of-band management interface.
* `out_of_band_ip_gateway`:(string)The default gateway of the network Element out-of-band management interface.
* `out_of_band_ip_mask`:(string)The network mask of the network Element out-of-band management interface.
* `out_of_band_ipv4_address`:(string)The IPv4 address of the network Element out-of-band management interface.
* `out_of_band_ipv4_gateway`:(string)The default IPv4 gateway of the network Element out-of-band management interface.
* `out_of_band_ipv4_mask`:(string)The network mask of the network Element out-of-band management interface.
* `out_of_band_ipv6_address`:(string)The IPv6 address of the network Element out-of-band management interface.
* `out_of_band_ipv6_gateway`:(string)The default IPv6 gateway of the network Element out-of-band management interface.
* `out_of_band_ipv6_prefix`:(string)The network mask of the network Element out-of-band management interface.
* `out_of_band_mac`:(string)The MAC address of the network Element out-of-band management interface.
* `revision`:(string)
* `rn`:(string)The Relative Name uniquely identifies an object within a given context.
* `serial`:(string)This field identifies the serial of the given component.
* `source_object_type`:(string)The source object type of this view MO.
* `switch_id`:(string)The Switch Id of the network Element.
* `vendor`:(string)This field identifies the vendor of the given component.
* `version`:(string)Version holds the firmware version related information.
