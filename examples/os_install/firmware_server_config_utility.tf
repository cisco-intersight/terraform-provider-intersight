resource "intersight_firmware_server_configuration_utility_distributable" "scu1" {
  name = "SCU-6.0.2a nfs"
  source {
    object_type = "softwarerepository.NfsServer"
    additional_properties = jsonencode({
      RemoteFile = "scu-602a.iso"
      RemoteIp = "10.197.110.132"
      RemoteShare = "/nfs"
    })
  }
  vendor = "Cisco"
  version = "6.0.(2a)"
  supported_models = [
    "C-series",
  ]
  description = "Cisco SCU-6.0(2a)"
}

resource "intersight_firmware_server_configuration_utility_distributable" "scu2" {
  name = "SCU-CIFS"
  source {
    object_type = "softwarerepository.CifsServer"
    additional_properties = jsonencode({
      RemoteIp: "10.225.79.79"
      RemoteShare: "/Public/iso"
      RemoteFile: "scu-602a.iso"
      Username: "user"
      Password: var.scu_password_1
      MountOption: "sec=ntlm"
    })
  }
  version = "6.0(2a)"
  supported_models = [
    "C-Series",
  ]
  description = "SCU-CIFS"
}

resource "intersight_firmware_server_configuration_utility_distributable" "scu3" {
  name = "SCU-HTTP-auth"
  supported_models = [
    "C-series",
  ]
  description = "SCU-HTTP-auth"
  source {
    object_type = "softwarerepository.HttpServer"
    additional_properties = jsonencode({
      LocationLink: "http://10.105.219.224/auth/scu-602a.iso"
      Username: "user1"
      Password: var.scu_password_2
    })
  }
  version = "6.0(2a)"
}

/*
SAMPLE PAYLOAD
-----------------
{
  "Description": "Cisco SCU-6.0(2a)",
  "Name": "SCU-6.0.2a nfs",
  "Source": {
    "ObjectType": "softwarerepository.NfsServer",
    "RemoteIp": "10.197.110.132",
    "RemoteShare": "/nfs",
    "RemoteFile": "scu-602a.iso"
  },
  "Vendor": "Cisco",
  "Version": "6.0.(2a)",
  "SupportedModels": ["C-series"]
}

{
    "Name": "SCU-CIFS",
    "Version": "6.0(2a)",
    "SupportedModels": [
        "C-Series"
    ],
    "Description": "SCU-CIFS",
    "Tags": [],
    "Source": {
        "ObjectType": "softwarerepository.CifsServer",
        "RemoteIp": "10.225.79.79",
        "RemoteShare": "/Public/iso",
        "RemoteFile": "scu-602a.iso",
        "Username": "user",
        "Password": "changeMe",
        "MountOption": "sec=ntlm"
    }
}

{
    "Name": "SCU-HTTP-auth",
    "Version": "6.0(2a)",
    "SupportedModels": [
        "C-series"
    ],
    "Description": "SCU-HTTP-auth",
    "Tags": [],
    "Source": {
        "ObjectType": "softwarerepository.HttpServer",
        "LocationLink": "http://10.105.219.224/auth/scu-602a.iso",
        "Username": "user1",
        "Password": "changeMe"
    }
}
*/