package hue

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccResourceHueLight_basic(t *testing.T) {

	resourceName := "hue_light.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		CheckDestroy:      testAccResourceCheckHueLightDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceHueLightConfig("65535", "240"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state.0.on", "true"),
					resource.TestCheckResourceAttr(resourceName, "light_index", "1"),
					resource.TestCheckResourceAttr(resourceName, "state.0.brightness", "240"),
				),
			},
			{
				Config: testAccResourceHueLightConfig("46920", "100"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state.0.on", "true"),
					resource.TestCheckResourceAttr(resourceName, "light_index", "1"),
					resource.TestCheckResourceAttr(resourceName, "state.0.brightness", "100"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceCheckHueLightDestroy(s *terraform.State) error {
	// ignore check-destroy as the lights will not be removed from the bridge
	return nil
}

func testAccResourceHueLightConfig(hue string, bri string) string {
	return fmt.Sprintf(`

data "hue_lights" "test" {}

resource "hue_light" "test" {
  unique_id = data.hue_lights.test.lights.0.unique_id

  state {
    on         = true
    hue        = "%s"
    brightness = "%s"
  }
}
`, hue, bri)
}
