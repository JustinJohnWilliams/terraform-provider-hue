---
page_title: "Data Source: hue_light"
subcategory: "Lights"
description: |-
---

# Data Source: hue_light

Use this data source to access information about an existing light.

## Example Usage

```hcl
data "hue_light" "example" {
  unique_id = "00:17:88:01:03:97:02:b8-0b"
}
```

## Arguments Reference

- `unique_id` - (Required) (String) The Unique Id of the light. The MAC address of the device with a unique endpoint id in the form: AA:BB:CC:DD:EE:FF:00:11-XX

## Attribute Reference

- `id` - (String) The ID of this resource.
- `light_index` - (Number) Hue light index.
- `name` - (String) A unique, editable name given to the light.
- `state` - A block of current state of the Light (see below).
- `model_id` - (String) The hardware model of the light.
- `product_id` - (String) The Product ID of the Light
- `sw_version` - (String) An identifier for the software version running on the light.

---

A `state` block exports the following:

- `brightness` - (Number) The brightness value to set the light to. Brightness is a scale from 1 (the minimum the light is capable of) to 254 (the maximum).
- `color_mode` - (String) Indicates the color mode in which the light is working, this is the last command type it received. Values are “hs” for Hue and Saturation, “xy” for XY and “ct” for Color Temperature. This parameter is only present when the light supports at least one of the values
- `hue` - (Number) The hue value to set light to.The hue value is a wrapping value between 0 and 65535. Both 0 and 65535 are red, 25500 is green and 46920 is blue.
- `on` - (Boolean) On/Off state of the light. On=true, Off=false
- `saturation` - (Number) Saturation of the light. 254 is the most saturated (colored) and 0 is the least saturated (white).
- `scene` - (String) The Scene name

---
