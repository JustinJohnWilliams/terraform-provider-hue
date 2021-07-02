package hue

import (
	"encoding/json"
	"fmt"
	"log"

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
			"light_index": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
				Optional: true,
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

func resourceHueLightCreateUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	light_index := d.Get("light_index").(int)

	updateConfig := huego.Light{}

	if name, ok := d.GetOk("name"); ok {
		updateConfig.Name = name.(string)
	}

	updatedLight, err := client.UpdateLight(light_index, updateConfig)
	if err != nil {
		return fmt.Errorf("Could not Create/Update Light with Index %d : %v", light_index, err)
	}

	updatedLightJson, err := json.Marshal(updatedLight)
	if err != nil {
		return fmt.Errorf("Could not Marshal Light with Index %d : %v", light_index, err)
	}

	log.Printf("[DEBUG] Updated light with Index %d : %s", light_index, updatedLightJson)

	if d.HasChange("state") {
		if err := resourceHueLightStateCreateUpdate(client, d); err != nil {
			return err
		}
	}

	d.SetId(fmt.Sprint(light_index))

	return resourceHueLightRead(d, meta)
}

func resourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	light_index := d.Get("light_index").(int)

	light, err := client.GetLight(light_index)
	if err != nil {
		return fmt.Errorf("Error retrieving light with Index %d : %v", light_index, err)
	}

	state := make([]map[string]interface{}, 0, 1)
	state = append(state, flattenLightState(light.State))

	d.Set("name", light.Name)
	d.Set("unique_id", light.UniqueID)
	d.Set("model_id", light.ModelID)
	d.Set("product_id", light.ProductID)
	d.Set("sw_version", light.SwVersion)
	d.Set("state", state)

	return nil

}

func resourceHueLightDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func resourceHueLightStateCreateUpdate(client *huego.Bridge, d *schema.ResourceData) error {

	light_index := d.Get("light_index").(int)

	state := d.Get("state").([]interface{})
	updateState := huego.State{}

	var currentState map[string]interface{}
	if state[0] != nil {
		currentState = state[0].(map[string]interface{})
	} else {
		currentState = make(map[string]interface{})
	}

	if v, ok := currentState["hue"]; ok {
		updateState.Hue = uint16(v.(int))
	}

	if v, ok := currentState["on"]; ok {
		updateState.On = v.(bool)
	}

	if v, ok := currentState["scene"]; ok {
		updateState.Scene = v.(string)
	}

	if v, ok := currentState["color_mode"]; ok {
		updateState.ColorMode = v.(string)
	}

	updatedLightState, err := client.SetLightState(light_index, updateState)
	if err != nil {
		return fmt.Errorf("Could not Create/Update Light State with Index %d : %v", light_index, err)
	}

	updatedLightStateJson, err := json.Marshal(updatedLightState)
	if err != nil {
		return fmt.Errorf("Could not Marshal Light State with Index %d : %v", light_index, err)
	}

	log.Printf("[DEBUG] Updated light state with Index %d : %s", light_index, updatedLightStateJson)

	return nil
}
