package ovh

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/compare"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func kmsServiceKeyStateCommonChecks(resName string, keyName string) []statecheck.StateCheck {
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("created_at"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("deactivation_reason"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("id"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("name"),
			knownvalue.StringExact(keyName)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("state"),
			knownvalue.StringExact("ACTIVE")),
	}
}

func kmsServiceKeyStateSymmetricChecks(resName string) []statecheck.StateCheck {
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("context"),
			knownvalue.NotNull()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("curve"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("size"),
			knownvalue.Int64Exact(256)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("type"),
			knownvalue.StringExact("oct")),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("operations"),
			knownvalue.SetExact(
				[]knownvalue.Check{
					knownvalue.StringExact("encrypt"),
					knownvalue.StringExact("decrypt"),
				},
			)),
	}
}

func kmsServiceKeyStateRsaChecks(resName string) []statecheck.StateCheck {
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("context"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("curve"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("size"),
			knownvalue.Int64Exact(2048)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("type"),
			knownvalue.StringExact("RSA")),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("operations"),
			knownvalue.SetExact(
				[]knownvalue.Check{
					knownvalue.StringExact("sign"),
					knownvalue.StringExact("verify"),
				},
			)),
	}
}

func kmsServiceKeyStateECChecks(resName string) []statecheck.StateCheck {
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("context"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("curve"),
			knownvalue.StringExact("P-256")),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("size"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("type"),
			knownvalue.StringExact("EC")),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("operations"),
			knownvalue.SetExact(
				[]knownvalue.Check{
					knownvalue.StringExact("sign"),
					knownvalue.StringExact("verify"),
				},
			)),
	}
}

const confOkmsServiceKeyTest = `
data "ovh_me" "current_account" {
}

resource "ovh_okms" "kms" {
  ovh_subsidiary = "FR"
  display_name = "%[1]s"
  region = "EU_WEST_SBG"
}

resource "ovh_okms_service_key" "key_symetric" {
  okms_id    = ovh_okms.kms.id
  name       = "%[1]s-sk-oct"
  type       = "oct"
  size       = 256
  operations = ["encrypt", "decrypt"]
  context    = "ctx"
}

resource "ovh_okms_service_key" "key_rsa" {
  okms_id    = ovh_okms.kms.id
  name       = "%[1]s-sk-rsa"
  type       = "RSA"
  size       = 2048
  operations = ["sign", "verify"]
}

resource "ovh_okms_service_key" "key_ecdsa" {
  okms_id    = ovh_okms.kms.id
  name       = "%[1]s-sk-ecdsa"
  type = "EC"
  curve = "P-256"
  operations = ["sign", "verify"]
}
`

func TestAccResourceOkmsServiceKeyCreate(t *testing.T) {
	resName := fmt.Sprintf("test-tf-%d", acctest.RandIntRange(10000, 99999))

	checks := []statecheck.StateCheck{
		statecheck.CompareValuePairs(
			"ovh_okms.kms",
			tfjsonpath.New("id"),
			"ovh_okms_service_key.key_symetric",
			tfjsonpath.New("okms_id"),
			compare.ValuesSame(),
		),
		statecheck.CompareValuePairs(
			"ovh_okms.kms",
			tfjsonpath.New("id"),
			"ovh_okms_service_key.key_rsa",
			tfjsonpath.New("okms_id"),
			compare.ValuesSame(),
		),
		statecheck.CompareValuePairs(
			"ovh_okms.kms",
			tfjsonpath.New("id"),
			"ovh_okms_service_key.key_ecdsa",
			tfjsonpath.New("okms_id"),
			compare.ValuesSame(),
		),
	}
	checks = append(checks, kmsServiceKeyStateCommonChecks("ovh_okms_service_key.key_symetric", resName+"-sk-oct")...)
	checks = append(checks, kmsServiceKeyStateCommonChecks("ovh_okms_service_key.key_rsa", resName+"-sk-rsa")...)
	checks = append(checks, kmsServiceKeyStateCommonChecks("ovh_okms_service_key.key_ecdsa", resName+"-sk-ecdsa")...)
	checks = append(checks, kmsServiceKeyStateSymmetricChecks("ovh_okms_service_key.key_symetric")...)
	checks = append(checks, kmsServiceKeyStateRsaChecks("ovh_okms_service_key.key_rsa")...)
	checks = append(checks, kmsServiceKeyStateECChecks("ovh_okms_service_key.key_ecdsa")...)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOrderOkms(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:            fmt.Sprintf(confOkmsServiceKeyTest, resName),
				ConfigStateChecks: checks,
			},
		},
	})
}
