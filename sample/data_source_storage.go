package sample

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceStorage() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceStorageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceStorageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Call data source read handler",
		Detail:   "Debug message: Call data source storage read handler",
	})

	name := d.Get("name").(string)

	var res []storage
	httpRes, err := http.Get(appURL + "/storage")
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	defer httpRes.Body.Close()
	json.NewDecoder(httpRes.Body).Decode(&res)

	for _, s := range res {
		if s.Name == name {
			d.SetId(s.ID)
			d.Set("name", s.Name)
			d.Set("size", s.Spec.Size)
			return diags
		}
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "No such storage",
		Detail:   fmt.Sprintf("Failed to find the storage %s", name),
	})

	return diags
}
