package hue

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceHueLights_basic(t *testing.T) {

	resourceName := "data.hue_lights.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHueLightsConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "lights.0.state.0.on", "true"),
					resource.TestCheckResourceAttr(resourceName, "lights.0.light_index", "1"),
				),
			},
		},
	})
}

func testAccDataSourceHueLightsConfig() string {
	return `
data "hue_lights" "test" {
}
`
}
