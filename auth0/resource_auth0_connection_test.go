package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccConnection(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "auth0"),
				),
			},
		},
	})
}

const testAccConnectionConfig = `
provider "auth0" {}

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection"
	strategy = "auth0"
	options = {
		password_policy = "fair"
	}
}
`