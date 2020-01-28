resource "intersight_server_profile" "server1" {
  name = "server1"
  tags {
    key   = "server"
    value = "demo"
  }
  organization {
    object_type = "organization.Organization"
    moid = "5e2540956972652d301b0a65"
  }
}