---
layout: "intersight"
page_title: "Intersight: intersight_networkconfig_policy"
sidebar_current: "docs-intersight-data-source-networkconfigPolicy"
description: |-
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.

---

# Data Source: intersight_networkconfig_policy
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `alternate_ipv4dns_server`:(string)IP address of the secondary DNS server.
* `alternate_ipv6dns_server`:(string)IP address of the secondary DNS server.
* `description`:(string)Description of the policy.
* `dynamic_dns_domain`:(string)The domain name appended to a hostname for a Dynamic DNS (DDNS) update. If left blank, only a hostname is sent to the DDNS update request.
* `enable_dynamic_dns`:(bool)If enabled, updates the resource records to the DNS from Cisco IMC.
* `enable_ipv6`:(bool)If enabled, allows to configure IPv6 properties.
* `enable_ipv4dns_from_dhcp`:(bool)If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv4 in Cisco IMC to enable it.
* `enable_ipv6dns_from_dhcp`:(bool)If enabled, Cisco IMC retrieves the DNS server addresses from DHCP. Use DHCP field must be enabled for IPv6 in Cisco IMC to enable it.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `name`:(string)Name of the concrete policy.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `preferred_ipv4dns_server`:(string)IP address of the primary DNS server.
* `preferred_ipv6dns_server`:(string)IP address of the primary DNS server.
