resource "intersight_server_profile" "server1" {
  name = "server1"
  tags {
    key   = "server"
    value = "demo"
  }
  organization {
    object_type = "organization.Organization"
    moid = var.organization
  }
}