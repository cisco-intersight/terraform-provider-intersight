package intersight

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_API_KEY", nil),
				Description: "API Key Id",
			},
			"secretkeyfile": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_SECRET_KEY_FILE_PATH", nil),
				Description: "Secret Key File Path",
			},
			"endpoint": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INTERSIGHT_ENDPOINT_URL", nil),
				Description: "Endpoint URL",
			},
		},
		DataSourcesMap: GetDataSourceMapping(),
		ResourcesMap:   GetResourceMapping(),
		ConfigureFunc:  configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		ApiKey:        d.Get("apikey").(string),
		SecretKeyFile: d.Get("secretkeyfile").(string),
		Endpoint:      d.Get("endpoint").(string),
	}
	return &config, nil
}
