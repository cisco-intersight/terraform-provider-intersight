resource "intersight_softwarerepository_operating_system_file" "osf1" {
  version = "ESXi 6.5 U2"
  description = "ESXi6.5U2 without answers"
  name = "ESXi6.5 w/o cifs 21"
  source {
    additional_properties = jsonencode({
      RemoteIp = "10.197.110.132"
      RemoteShare = "/nfs"
      RemoteFile = "esx65u2.iso"
    })
    object_type = "softwarerepository.NfsServer"
  }
  vendor = "VMware"
}

resource "intersight_softwarerepository_operating_system_file" "osf2" {
  source {
    additional_properties = jsonencode({
      RemoteIp = "10.225.79.79"
      RemoteShare = "/Public/iso"
      RemoteFile = "esx6u3.iso"
      Username = "user"
      Password = var.os_file_password_1
      MountOption = "sec=ntlm"
    })
    object_type = "softwarerepository.CifsServer"
  }
  name = "ESXi6.0-CIFS"
  vendor = "VMware"
  version = "ESXi 6.0 U3"
  description = "ESXi6.0-CIFS"
}