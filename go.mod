module github.com/cisco-intersight/terraform-provider-intersight

go 1.13

require (
	github.com/bflad/tfproviderlint v0.14.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/stretchr/testify v1.4.0 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

require github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk v0.0.0

replace github.com/cisco-intersight/terraform-provider-intersight/intersight_gosdk v0.0.0 => ./intersight_gosdk
