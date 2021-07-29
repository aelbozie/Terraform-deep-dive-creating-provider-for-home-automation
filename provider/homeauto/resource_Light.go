package homeauto

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLight() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceLightCreate,
		ReadContext:   resourceLightRead,
		UpdateContext: resourceLightUpdate,
		DeleteContext: resourceLightDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "on",
			},
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Light",
			},
			"color_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "hs",
			},
			"brightness": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  255,
			},
			"white_value": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"supported_features": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  147,
			},
			"hs_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
			"rgb_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"xy_color": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeFloat,
				},
			},
		},
	}
}
func resourceLightCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	item := LightItem{
		EntityID: d.Get("entity_id").(string),
		State:    d.Get("state").(string),
	}
	err := StartLight(item, *c)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId(item.EntityID)
	return resourceLightRead(ctx, d, m)
}
func resourceLightRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	lightID := d.Id()

	light, err := GetLight(lightID, *c)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("state", light.State); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
func resourceLightUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	item := LightItem{
		EntityID: d.Get("entity_id").(string),
		State:    d.Get("state").(string),
	}

	err := StartLight(item, *c)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceLightRead(ctx, d, m)
}
func resourceLightDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)
	var diags diag.Diagnostics
	lightID := d.Id()

	err := DeleteLight(lightID, *c)
	if err != nil {
		return diag.FromErr(err)
	}
	d.SetId("")
	return diags
}
