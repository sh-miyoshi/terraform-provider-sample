package sample

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type storageSpec struct {
	Size int `json:"size"`
}

type storage struct {
	ID   string      `json:"id"`
	Name string      `json:"name"`
	Spec storageSpec `json:"spec"`
}

func resourceStorage() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStorageCreate,
		ReadContext:   resourceStorageRead,
		UpdateContext: resourceStorageUpdate,
		DeleteContext: resourceStorageDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceStorageCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Call create handler",
		Detail:   "Debug message: Call create handler",
	})

	name := d.Get("name").(string)
	size := d.Get("size").(int)

	body, _ := json.Marshal(storage{
		Name: name,
		Spec: storageSpec{
			Size: size,
		},
	})
	httpRes, err := http.Post(appURL+"/storage", "application/json", bytes.NewBuffer(body))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	defer httpRes.Body.Close()

	var res storage
	json.NewDecoder(httpRes.Body).Decode(&res)

	d.SetId(res.ID)

	return diags
}

func resourceStorageRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Call read handler",
		Detail:   "Debug message: Call read handler",
	})

	id := d.Id()
	httpRes, err := http.Get(appURL + "/storage/" + id)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	defer httpRes.Body.Close()

	var res storage
	json.NewDecoder(httpRes.Body).Decode(&res)

	d.Set("name", res.Name)
	d.Set("size", res.Spec.Size)

	return diags
}

func resourceStorageUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Call update handler",
		Detail:   "Debug message: Call update handler",
	})

	if d.HasChange("name") || d.HasChange("size") {
		id := d.Id()

		name := d.Get("name").(string)
		size := d.Get("size").(int)

		body, _ := json.Marshal(storage{
			Name: name,
			Spec: storageSpec{
				Size: size,
			},
		})

		client := &http.Client{}
		req, _ := http.NewRequest(http.MethodPut, appURL+"/storage/"+id, bytes.NewBuffer(body))
		if _, err := client.Do(req); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	return diags
}

func resourceStorageDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Call delete handler",
		Detail:   "Debug message: Call delete handler",
	})

	id := d.Id()
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, appURL+"/storage/"+id, nil)
	if _, err := client.Do(req); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
