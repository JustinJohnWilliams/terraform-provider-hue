---
page_title: "Data Source: hue_lights"
subcategory: "Lights"
description: |-
---

# Data Source: hue_lights

Use this data source to access information about all existing lights.

## Example Usage

```hcl
data "hue_lights" "example" {
}
```

## Attribute Reference

- `lights` - A list of all the lights available as [hue_light data source](./light.md) objects.
