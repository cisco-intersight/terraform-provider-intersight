resource "intersight_firmware_server_configuration_utility_distributable" "scu1" {
  name = "SCU-6.0.4c nfs"
  source {
    object_type = "softwarerepository.CifsServer"
    additional_properties = jsonencode({
      RemoteIp: "10.225.79.79"
      RemoteShare: "/Public/iso"
      RemoteFile: "scu-604c.iso"
      Username: "user"
      Password: "ChangeMe"
      MountOption: "sec=ntlm"
    })
  }
  vendor = "Cisco"
  version = "6.0.(4c)"
  supported_models = [
    "C-series",
  ]
  description = "Cisco SCU-6.0(4c)"
  catalog {
    moid = "5dde4cb06567612d3063467a"
  }
}
