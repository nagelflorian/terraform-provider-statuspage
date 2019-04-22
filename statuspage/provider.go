package statuspage

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("STATUSPAGE_API_KEY", nil),
				Description: "The Statuspage API Key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"statuspage_page": resourcePage(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{APIToken: d.Get("api_key").(string)}
	return config.Client()
}
