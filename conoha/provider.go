package conoha

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"tenant_name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_TENANT", nil),
				Description: "A ConoHa tenant name.",
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_USERNAME", nil),
				Description: "A ConoHa user name.",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CONOHA_PASSWORD", nil),
				Description: "A ConoHa user password.",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"conoha_container": resourceConohaContainer(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		TenantName: d.Get("tenant_name").(string),
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
	}

	return config.NewClient()
}
