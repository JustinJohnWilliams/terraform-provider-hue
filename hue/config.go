package hue

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
)

type Config struct {
	Host             string
	Username         string
	terraformVersion string
}

func (c *Config) Client() (*huego.Bridge, error) {

	if c.Host == "" || c.Username == "" {
		return nil, fmt.Errorf("Hue provider requires host and username.")
	}

	client := huego.New(c.Host, c.Username)
	config, err := client.GetConfig()
	if err != nil {
		return nil, fmt.Errorf("Provider Error, Could not get Hue Bridge configuration: %s", err)
	}

	log.Printf("[INFO] Found Bridge ID: %s", config.BridgeID)

	return client, nil
}
