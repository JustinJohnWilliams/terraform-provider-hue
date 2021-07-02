package hue

import (
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

	return lights
}

func flattenLightState(state *huego.State) map[string]interface{} {
	flattenstate := map[string]interface{}{
		"hue":        state.Hue,
		"on":         state.On,
		"color_mode": state.ColorMode,
		"scene":      state.Scene,
	}
	return flattenstate
}
