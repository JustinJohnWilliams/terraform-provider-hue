---
page_title: "Philips Hue Bridge Provider"
subcategory: ""
description: |-
---

# Philips Hue Bridge Provider

A provider to mange Philips Hue Bridge configuration and devices.

## Example Usage

```hcl
terraform {
  required_providers {
    hue = {
      source  = "AliAllomani/hue"
      version = "~> 0.0"
    }
  }
}
```

## Authentication: Get Hue Bridge IP and Username

Follow [Hue Developer's Get Started](https://developers.meethue.com/develop/get-started-2/) Guide to retrieve Hue Bridge IP Address and create API Username.

## Argument Reference

- `host` (Optional) (String) The Hue Bridge Hostname or IP Address. It must be provided, but
  it can also be sourced from the `HUE_HOST` environment variable.
- `username` (Optional) (String) The API Username of Hue Bridge. It must be provided, but
  it can also be sourced from the `HUE_USER` environment variable.
