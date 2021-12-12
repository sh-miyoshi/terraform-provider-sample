package sample

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	appURL string
)

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"app_url": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sample_storage": dataSourceStorage(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"sample_vm":      resourceVM(),
			"sample_storage": resourceStorage(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	defaultAppURL := "http://localhost:4567"

	appURL = d.Get("app_url").(string)
	if appURL == "" {
		appURL = defaultAppURL
	}

	return nil, diag.Diagnostics{}
}
