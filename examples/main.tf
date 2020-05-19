provider "intersight" {
  apikey = "5db9f3567564612d3084dcc3/5db9fa327564612d3084e469/5ebce3857564612d30488203"
  secretkeyfile = "/Users/aanimish/Downloads/SecretKey.txt"
  endpoint = "vvb-ad-onprem-infranode1.cisco.com"
}

//provider "intersight" {
//  apikey    = "5ebc49e17564612d30cda178/5ebc4d2f7564612d30cdb584/5ebe454f7564612d302763ae"
//  secretkeyfile = "/Users/aanimish/Downloads/SecretKey-aanimish.txt"
//  endpoint = "aanimish-appliance.cisco.com"
//}

//resource "intersight_boot_precision_policy" "boot_precision1" {
//  name                     = "boot_precision1"
//  description              = "test policy"
//  configured_boot_mode     = "Legacy"
//  enforce_uefi_secure_boot = false
//  organization {
//    object_type = "organization.Organization"
//    moid = "5e14d0966972652d30638167"
//  }
//  boot_devices {
//    enabled     = true
//    name        = "scu-device-hdd"
//    object_type = "boot.LocalDisk"
//    additional_properties = jsonencode({
//      Slot = "MRAID"
//      Bootloader = {
//        Description = ""
//        Name        = ""
//        ObjectType  = "boot.Bootloader"
//        Path        = ""
//      }
//    })
//  }
//  boot_devices {
//    enabled     = true
//    name        = "NIIODCIMCDVD"
//    object_type = "boot.VirtualMedia"
//    additional_properties = jsonencode({
//      Subtype = "cimc-mapped-dvd"
//    })
//  }
//  boot_devices {
//    enabled     = true
//    name        = "hdd"
//    object_type = "boot.LocalDisk"
//    additional_properties = jsonencode({
//      Slot = "MRAID"
//      Bootloader = {
//        Description = ""
//        Name        = ""
//        ObjectType  = "boot.Bootloader"
//        Path        = ""
//      }
//    })
//  }
////  profiles {
////    moid        = intersight_server_profile.server1.id
////    object_type = "server.Profile"
////  }
//
//}

//resource "intersight_kvm_policy" "kvm1" {
//  name = "aaa-kvm1"
//  description = "demo kvm policy"
//  enabled = true
//  maximum_sessions = 3
//  remote_port = 2069
//  enable_video_encryption = true
//  enable_local_server_video = true
//  organization {
//    moid = "5e14d0966972652d30638167"
//  }
//  profiles {
//    moid = intersight_server_profile.server1.id
//  }
//}
//
//resource "intersight_server_profile" "server1" {
//  name = "aaa-server1"
//  tags {
//    key = "server"
//    value = "demo"
//  }
//  organization {
//    moid = "5e14d0966972652d30638167"
//  }
//}

data "intersight_server_profile" "s1" {
  name = "goldprofile"
}