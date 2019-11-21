---
layout: "intersight"
page_title: "Intersight: intersight_license_account_license_data"
sidebar_current: "docs-intersight-data-source-licenseAccountLicenseData"
description: |-
License information for an account.

---

# Data Source: intersight_license_account_license_data
License information for an account.

## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `account_id`:(string)Root user's ID of the account.
* `agent_data`:(string)Agent trusted store data.
* `auth_expire_time`:(string)Authorization expiration time.
* `auth_initial_time`:(string)Intial authorization time.
* `auth_next_time`:(string)Next time for the authorization.
* `category`:(string)Account license data category name.
* `group`:(string)Account license data group name.
* `last_sync`:(string)Specifies last sync time with SA.
* `last_updated_time`:(string)Record's last update datetime.
* `license_state`:(string)Aggregrated mode for the agent.
* `license_tech_support_info`:(string)Tech-support info of a smart-agent.
* `moid`:(string)The unique identifier of this Managed Object instance.
* `object_type`:(string)The concrete type of this complex type.The ObjectType property must be set explicitly by API clients when the type is ambiguous. In all other cases, the ObjectType is optional. The type is ambiguous when a managed object contains an array of nested documents, and the documents in the arrayare heterogeneous, i.e. the array can contain nested documents of different types.
* `register_expire_time`:(string)Registration exipiration time.
* `register_initial_time`:(string)Initial time of registration.
* `register_next_time`:(string)Next time for the license registration.
* `registration_status`:(string)Registration status of a smart-agent.
* `renew_failure_string`:(string)License renewal failure message.
* `smart_account`:(string)Name of the smart account.
* `sync_status`:(string)Current sync status for the account.
* `virtual_account`:(string)Name of the virtual account.
