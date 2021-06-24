package hue

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceHueLights() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHueLightsRead,
		Schema: map[string]*schema.Schema{
			"lights": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
				},
			},
		},
	}
}

func dataSourceHueLightsRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	d.SetId("")

	lights, err := client.GetLights()
	if err != nil {
		return fmt.Errorf("Error retrieving lights: %v", err)
	}

	fLights := flattenLights(lights)

	log.Printf("[INFO] found %d lights.", len(lights))

	log.Printf("[INFO] found light name %s.", fLights[0]["name"])

	d.Set("lights", fLights)

	return nil
}
