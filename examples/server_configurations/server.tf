resource "intersight_server_profile" "server1" {
  name = "server1"
  tags {
    key   = "server"
    value = "demo"
  }
  assigned_server {
    moid        = "5d7b2f8e6176752d30511ee6"
    object_type = "compute.RackUnit"
  }
}