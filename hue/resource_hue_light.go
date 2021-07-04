package hue

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/amimof/huego"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"unique_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique id of the device. The MAC address of the device with a unique endpoint id in the form: AA:BB:CC:DD:EE:FF:00:11-XX",
			},
			"light_index": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Hue light index (ID)",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "A unique, editable name given to the light.",
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
				MaxItems:    1,
				Computed:    true,
				Optional:    true,
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
							Optional:    true,
							Computed:    true,
							Description: "The brightness value to set the light to. Brightness is a scale from 1 (the minimum the light is capable of) to 254 (the maximum).",
						},
						"saturation": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Saturation of the light. 254 is the most saturated (colored) and 0 is the least saturated (white).",
						},
					},
				},
			},
		},
	}
}

func resourceHueLightCreateUpdate(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	uniqueId := d.Get("unique_id").(string)

	lightIndex, err := lightIndexFromUniqueId(client, uniqueId)
	if err != nil {
		return fmt.Errorf("Error create/update light : %v", err)
	}

	d.SetId(uniqueId)
	d.Set("light_index", lightIndex)

	updateConfig := huego.Light{}

	if name, ok := d.GetOk("name"); ok {
		updateConfig.Name = name.(string)
	}

	updatedLight, err := client.UpdateLight(lightIndex, updateConfig)
	if err != nil {
		return fmt.Errorf("Could not Create/Update Light with Index %d : %v", lightIndex, err)
	}

	updatedLightJson, err := json.Marshal(updatedLight)
	if err != nil {
		return fmt.Errorf("Could not Marshal Light with Index %d : %v", lightIndex, err)
	}

	log.Printf("[DEBUG] Updated light with Index %d : %s", lightIndex, updatedLightJson)

	if d.HasChange("state") {
		if err := resourceHueLightStateCreateUpdate(client, d); err != nil {
			return err
		}
	}

	return resourceHueLightRead(d, meta)
}

func resourceHueLightRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*huego.Bridge)
	uniqueId := d.Id()

	light_index, err := lightIndexFromUniqueId(client, uniqueId)
	if err != nil {
		return fmt.Errorf("Error parsing light Index from ID : %s : %v", uniqueId, err)
	}

	light, err := client.GetLight(light_index)
	if err != nil {
		return fmt.Errorf("Error retrieving light with Index %d : %v", light_index, err)
	}

	state := make([]map[string]interface{}, 0, 1)
	state = append(state, flattenLightState(light.State))

	d.Set("light_index", light.ID)
	d.Set("name", light.Name)
	d.Set("unique_id", light.UniqueID)
	d.Set("model_id", light.ModelID)
	d.Set("product_id", light.ProductID)
	d.Set("sw_version", light.SwVersion)
	d.Set("state", state)

	return nil

}

func resourceHueLightDelete(d *schema.ResourceData, meta interface{}) error {
	// We will not delete the light from the bridge, just from the state file
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

	if v, ok := currentState["brightness"]; ok {
		updateState.Bri = uint8(v.(int))
	}

	if v, ok := currentState["saturation"]; ok {
		updateState.Sat = uint8(v.(int))
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
