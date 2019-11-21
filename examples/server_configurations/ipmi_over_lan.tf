resource "intersight_ipmioverlan_policy" "ipmi1" {
  name           = "ipmi1"
  description    = "demo ipmi policy"
  enabled        = true
  privilege      = "admin"
  encryption_key = var.encryption_key
  profiles {
    moid        = intersight_server_profile.server1.id
    object_type = "server.Profile"
  }
}

/*
SAMPLE PAYLOAD
-----------------
IpmioverlanPolicyApi: {
  "Name": "AUTO_IPMI_POLICY_ADMIN_CRR",
  "Description": "IPMI Admin Testing",
  "Tags": [{"Key": "ipmiadmin", "Value": "555"}],
  "Enabled": True,
  "Privilege": "admin",
  "EncryptionKey": "345653"
}
*/