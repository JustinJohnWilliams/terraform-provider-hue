package hue

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHueLight() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHueLightRead,
		Schema: map[string]*schema.Schema{
			"light_index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"model_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sw_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"hue": {
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"color_mode": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"on": {
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
						},
						"scene": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	light_index := d.Get("light_index").(int)
	d.SetId(fmt.Sprint(light_index))

	return resourceHueLightRead(d, meta)
}
