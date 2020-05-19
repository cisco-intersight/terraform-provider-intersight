resource "intersight_snmp_policy" "snmp1" {
  name                    = "snmp1"
  description             = "testing smtp policy"
  enabled                 = true
  snmp_port               = 1983
  access_community_string = "1234cisco"
  community_access        = "Disabled"
  trap_community          = "TrapCommunity"
  sys_contact             = "aanimish"
  sys_location            = "Karnataka"
  engine_id               = "vvb"
  snmp_users {
    name             = "aarush1234"
    privacy_type     = "AES"
    auth_password    = var.auth_password
    privacy_password = var.privacy_password
    security_level   = "AuthPriv"
    auth_type        = "SHA"
  }
  snmp_traps {
    destination = "10.123.110.81"
    enabled     = false
    port        = 660
    type        = "Trap"
    user        = "aarush1234"
    version     = "V3"
  }
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}

/*
SAMPLE PAYLOAD
-----------------
SnmpPolicyApi: {
    "Name": "AUTO_SNMP_POLICY_CRR",
    "Enabled": True,
    "Description": "SNMP Test policy",
    "Tags": [{"Key": "snmp", "Value": "snmp_policy"}],
    "SnmpPort": 1983,
    "AccessCommunityString": "1234cisco",
    "CommunityAccess": "Disabled",
    "TrapCommunity": "TrapCommunity",
    "SysContact": "guptaraghav",
    "SysLocation": "Karnataka",
    "EngineId": "vvb",
    "SnmpUsers": [{"Name": "aarush1234", "SecurityLevel": "AuthPriv",
                   "AuthType": "SHA", "AuthPassword": "cisco1234",
                   "PrivacyType": "AES", "PrivacyPassword": "cisco1234"}],
    "SnmpTraps": [{"Enabled": True, "Version": "V3",
                   "Type": "Trap", "User": "aarush1234",
                   "Destination": "10.123.110.81", "Port": 660}]
}
*/