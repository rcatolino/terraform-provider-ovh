package ovh

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func kmsResourceStateChecks(displayName string) []statecheck.StateCheck {
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("iam.display_name"),
			knownvalue.StringExact(displayName)),
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("display_name"),
			knownvalue.StringExact(displayName)),
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("kmip_endpoint"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("swagger_endpoint"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("rest_endpoint"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			"ovh_okms.kms",
			tfjsonpath.New("public_ca"),
			knownvalue.NotNull()),
	}
}

const confOkmsResourceOrder = `
resource "ovh_okms" "newkms" {
  ovh_subsidiary = "FR"
  display_name = "%s"
  region = "EU_WEST_SBG"
}
`

func TestAccResourceOkmsOrder(t *testing.T) {
	compareValuesSame := compare.CompareValue(compare.ValuesSame())
	displayName := acctest.RandomWithPrefix(test_prefix)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOrderOkms(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(confOkmsResourceOrder, displayName),
				ConfigStateChecks: append(
					kmsResourceStateChecks(displayName),
					statecheck.ExpectKnownValue(
						"ovh_okms.kms",
						tfjsonpath.New("id"),
						knownvalue.NotNull()),
					compareValuesSame.AddStateValue(
						"ovh_okms.kms",
						tfjsonpath.New("id")),
				),
			},
			{
				Config: fmt.Sprintf(confOkmsResourceOrder, displayName+"new"),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectResourceAction(
							"ovh_okms.kms",
							plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: append(
					kmsResourceStateChecks(displayName),
					compareValuesSame.AddStateValue(
						"ovh_okms.kms",
						tfjsonpath.New("id")),
				),
			},
		},
	})
}

const confOkmsResourceUpdate = `
resource "ovh_okms" "kms" {
	display_name = "%s"
	ovh_subsidiary = "FR"
	region = "EU_WEST_SBG"
}
`

func TestAccResourceOkmsUpdate(t *testing.T) {
	displayName := acctest.RandomWithPrefix(test_prefix)
	kmsId := os.Getenv("OVH_OKMS")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOkms(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:        fmt.Sprintf(confOkmsResourceUpdate, displayName),
				ResourceName:  "ovh_okms.kms",
				ImportState:   true,
				ImportStateId: kmsId,
				//ImportStatePersist: true,
			},
			{
				Config: fmt.Sprintf(confOkmsResourceUpdate, displayName),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{
						plancheck.ExpectNonEmptyPlan(),
						plancheck.ExpectKnownValue(
							"ovh_okms.kms",
							tfjsonpath.New("id"),
							knownvalue.StringExact(kmsId)),
						plancheck.ExpectResourceAction(
							"ovh_okms.kms",
							plancheck.ResourceActionUpdate),
					},
				},
				ConfigStateChecks: append(
					kmsResourceStateChecks(displayName),
					statecheck.ExpectKnownValue(
						"ovh_okms.kms",
						tfjsonpath.New("id"),
						knownvalue.StringExact(kmsId)),
				),
			},
			/*
				{
					Config: fmt.Sprintf(confOkmsResourceUpdate, displayName+"2"),
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{
							plancheck.ExpectKnownValue(
								"ovh_okms.kms",
								tfjsonpath.New("id"),
								knownvalue.StringExact(kmsId)),
						},
					},
					ConfigStateChecks: append(
						kmsResourceStateChecks(displayName+"2"),
						statecheck.ExpectKnownValue(
							"ovh_okms.kms",
							tfjsonpath.New("id"),
							knownvalue.StringExact(kmsId)),
					),
				},
			*/
		},
	})
}
