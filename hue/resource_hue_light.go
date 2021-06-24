package hue

import (
	"fmt"

	"github.com/amimof/huego"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceHueLight() *schema.Resource {
	return &schema.Resource{
		Create: resourceHueLightCreateUpdate,
		Read:   resourceHueLightRead,
		Update: resourceHueLightCreateUpdate,
		Delete: resourceHueLightDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"index": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uniqueid": {
				Type:     schema.TypeString,
				Computed: true,
			},

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
	}
}

func resourceHueLightCreateUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	id := d.Get("index").(int)

	config := huego.Light{}

	if name, ok := d.GetOk("name"); ok {
		config.Name = name.(string)
	}

	// if hue, ok := d.GetOk("hue"); ok {
	// 	config.State.Hue = uint16(hue.(int))
	// }

	if on, ok := d.GetOk("on"); ok {
		config.State.On = on.(bool)
	}

	if _, err := client.UpdateLight(id, config); err != nil {
		return fmt.Errorf("Could not create Light with Index %d : %v", id, err)
	}

	return resourceHueLightRead(d, meta)
}

func resourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	id := d.Get("index").(int)

	d.SetId(fmt.Sprint(id))

	light, err := client.GetLight(id)
	if err != nil {
		return fmt.Errorf("Error retrieving light: %v", err)
	}

	d.Set("name", light.Name)
	d.Set("uniqueid", light.UniqueID)
	d.Set("on", light.State.On)
	d.Set("hue", light.State.Hue)
	d.Set("scene", light.State.Scene)
	d.Set("color_mode", light.State.ColorMode)

	return nil

}

func resourceHueLightDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
