resource "intersight_adapter_config_policy" "adapter_config1" {
  name        = "adapter_config1"
  description = "test policy"
  settings {
    slot_id = "1"
    eth_settings {
      lldp_enabled = true
    }
    fc_settings {
      fip_enabled = true
    }
  }
  settings {
    slot_id = "MLOM"
    eth_settings {
      lldp_enabled = true
    }
    fc_settings {
      fip_enabled = true
    }
  }
}

/*
SAMPLE PAYLOAD
-----------------
AdapterConfigPolicyApi: {
  "Name": "AUTO_TEST_ACP1",
  "Settings": [
    {
      "SlotId": "1",
      "EthSettings": {"LldpEnabled": True},
      "FcSettings": {"FipEnabled": True}
    }
  ]
}
*/