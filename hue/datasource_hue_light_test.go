package hue

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceHueLight_basic(t *testing.T) {

	resourceName := "data.hue_light.test"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceHueLightConfig("65535", "240"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "state.0.on", "true"),
					resource.TestCheckResourceAttr(resourceName, "light_index", "1"),
				),
			},
		},
	})
}

func testAccDataSourceHueLightConfig(hue string, bri string) string {
	return fmt.Sprintf(`

	%s

data "hue_light" "test" {
  unique_id = hue_light.test.id
}
`, testAccResourceHueLightConfig(hue, bri))
}
