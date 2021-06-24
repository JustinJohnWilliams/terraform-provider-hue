package hue

import (
	"fmt"
	"log"

	"github.com/amimof/huego"
)

type Config struct {
	Host     string
	Username string
}

func (c *Config) Client() (*huego.Bridge, error) {

	if c.Host == "" || c.Username == "" {
		return nil, fmt.Errorf("Hue provider requires host and username.")
	}

	client := huego.New(c.Host, c.Username)
	lights, err := client.GetLights()
	if err != nil {
		return nil, fmt.Errorf("Provider Error %s", err)
	}

	log.Printf("[INFO] Found %d lights", len(lights))

	return client, nil
}
