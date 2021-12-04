package sample

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	appURL = "http://localhost:4567"
)

// Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"sample_vm":      resourceVM(),
			"sample_storage": resourceStorage(),
		},
	}
}
