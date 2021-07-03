package hue

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name/IP of the Hue Bridge",
				DefaultFunc: schema.EnvDefaultFunc("HUE_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username of Hue Bridge",
				DefaultFunc: schema.EnvDefaultFunc("HUE_USER", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"hue_lights": dataSourceHueLights(),
			"hue_light":  dataSourceHueLight(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"hue_light": resourceHueLight(),
		},
	}
	p.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := p.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, terraformVersion)
	}
	return p
}

func providerConfigure(d *schema.ResourceData, terraformVersion string) (interface{}, error) {
	config := Config{
		Host:             d.Get("host").(string),
		Username:         d.Get("username").(string),
		terraformVersion: terraformVersion,
	}

	cfg, err := config.Client()
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
