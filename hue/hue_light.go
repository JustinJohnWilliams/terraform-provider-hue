package hue

import (
	"fmt"

	"github.com/amimof/huego"
)

func flattenLights(input []huego.Light) []map[string]interface{} {

	var lights []map[string]interface{}

	for _, light := range input {
		lights = append(lights, map[string]interface{}{
			"name": light.Name,
		})
	}

	return lights
}

func flattenLightState(state *huego.State) interface{} {
	flattenstate := map[string]interface{}{
		"hue":        fmt.Sprint(state.Hue),
		"on":         fmt.Sprint(state.On),
		"color_mode": fmt.Sprint(state.ColorMode),
		"scene":      fmt.Sprint(state.Scene),
	}
	return flattenstate
}
