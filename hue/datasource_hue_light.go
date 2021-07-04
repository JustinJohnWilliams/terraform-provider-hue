package hue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHueLight() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHueLightRead,
		Schema: map[string]*schema.Schema{
			"unique_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique id of the device. The MAC address of the device with a unique endpoint id in the form: AA:BB:CC:DD:EE:FF:00:11-XX",
			},

			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A unique, editable name given to the light.",
			},
			"light_index": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Hue light index (ID)",
			},
			"model_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The hardware model of the light.",
			},
			"product_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Product ID of the Light",
			},
			"sw_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An identifier for the software version running on the light.",
			},
			"state": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The current state of the Light",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hue": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The hue value to set light to.The hue value is a wrapping value between 0 and 65535. Both 0 and 65535 are red, 25500 is green and 46920 is blue.",
						},
						"color_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Indicates the color mode in which the light is working, this is the last command type it received. Values are “hs” for Hue and Saturation, “xy” for XY and “ct” for Color Temperature. This parameter is only present when the light supports at least one of the values",
						},
						"on": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "On/Off state of the light. On=true, Off=false",
						},
						"scene": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The Scene name",
						},
						"brightness": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The brightness value to set the light to. Brightness is a scale from 1 (the minimum the light is capable of) to 254 (the maximum).",
						},
						"saturation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Saturation of the light. 254 is the most saturated (colored) and 0 is the least saturated (white).",
						},
					},
				},
			},
		},
	}
}

func dataSourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	uniqueId := d.Get("unique_id").(string)
	d.SetId(uniqueId)

	return resourceHueLightRead(d, meta)
}
