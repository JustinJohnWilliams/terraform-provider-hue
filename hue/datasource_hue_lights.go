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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"light_index": {
							Type:     schema.TypeInt,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceHueLightsRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	d.SetId("lights")

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
