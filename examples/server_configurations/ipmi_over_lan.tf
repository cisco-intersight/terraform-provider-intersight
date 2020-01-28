resource "intersight_ipmioverlan_policy" "ipmi1" {
  name           = "ipmi1"
  description    = "demo ipmi policy"
  enabled        = true
  privilege      = "admin"
  encryption_key = var.encryption_key
  organization {
    object_type = "organization.Organization"
    moid = "5e2540956972652d301b0a65"
  }
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
  "EncryptionKey": "345653",
  "Profiles" : [{
    "Moid":"567890qwert4567",
    "ObjectType":"server.Profile"
  }]
}
*/