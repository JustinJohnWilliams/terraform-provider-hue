package hue

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHueLight() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHueLightRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uniqueid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeMap,
				Computed: true,
			},
		},
	}
}

func dataSourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	id := d.Get("id").(int)

	d.SetId(fmt.Sprint(id))

	light, err := client.GetLight(id)
	if err != nil {
		return fmt.Errorf("Error retrieving light: %v", err)
	}

	d.Set("name", light.Name)
	d.Set("uniqueid", light.UniqueID)
	d.Set("state", flattenLightState(light.State))

	return nil
}
