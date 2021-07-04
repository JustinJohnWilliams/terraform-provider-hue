package hue

import (
	"fmt"
	"log"
	"sort"

	"github.com/amimof/huego"
)

func flattenLights(input []huego.Light) []map[string]interface{} {

	var lights []map[string]interface{}
	var state []map[string]interface{}

	for _, light := range input {

		state = make([]map[string]interface{}, 0, 1)
		state = append(state, flattenLightState(light.State))

		lights = append(lights, map[string]interface{}{
			"light_index": light.ID,
			"name":        light.Name,
			"model_id":    light.ModelID,
			"product_id":  light.ProductID,
			"unique_id":   light.UniqueID,
			"sw_version":  light.SwVersion,
			"state":       state,
		})
	}

	sort.Slice(lights, func(i, j int) bool {
		return lights[i]["light_index"].(int) < lights[j]["light_index"].(int)
	})

	return lights
}

func flattenLightState(state *huego.State) map[string]interface{} {
	flattenstate := map[string]interface{}{
		"hue":        state.Hue,
		"on":         state.On,
		"color_mode": state.ColorMode,
		"scene":      state.Scene,
		"brightness": state.Bri,
		"saturation": state.Sat,
	}
	return flattenstate
}

func lightIndexFromUniqueId(client *huego.Bridge, uniqueID string) (int, error) {

	lights, err := client.GetLights()
	if err != nil {
		return 0, fmt.Errorf("Error retrieving lights: %v", err)
	}

	for _, light := range lights {
		if light.UniqueID == uniqueID {
			log.Printf("[INFO] Found Light Index %d for Light Unique ID: %s", light.ID, light.UniqueID)
			return light.ID, nil
		}
	}

	return 0, fmt.Errorf("Cloud not find a light with Unique ID : %s", uniqueID)

}
