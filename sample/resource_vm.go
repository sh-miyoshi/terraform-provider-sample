package sample

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type vmSpec struct {
	CPU    int `json:"cpu"`
	Memory int `json:"memory"`
}

type vm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Spec vmSpec `json:"spec"`
}

func resourceVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVMCreate,
		ReadContext:   resourceVMRead,
		UpdateContext: resourceVMUpdate,
		DeleteContext: resourceVMDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"memory": {
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	name := d.Get("name").(string)
	cpu := d.Get("cpu").(int)
	memory := d.Get("memory").(int)

	body, _ := json.Marshal(vm{
		Name: name,
		Spec: vmSpec{
			CPU:    cpu,
			Memory: memory,
		},
	})
	httpRes, err := http.Post(appURL+"/vm", "application/json", bytes.NewBuffer(body))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to request to server",
			Detail:   err.Error(),
		})
		return diags
	}
	defer httpRes.Body.Close()

	var res vm
	json.NewDecoder(httpRes.Body).Decode(&res)

	d.SetId(res.ID)

	return diags
}

func resourceVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}

func resourceVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	return diags
}