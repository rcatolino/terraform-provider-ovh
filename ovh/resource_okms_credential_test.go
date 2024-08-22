package ovh

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func kmsCredStateCommonChecks(resName string, credName string) []statecheck.StateCheck {
	urnRe := regexp.MustCompile("urn:v1:eu:identity:account:.*")
	pemRe := regexp.MustCompile("-----BEGIN CERTIFICATE-----.*")
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("name"),
			knownvalue.StringExact(credName)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("identity_urns"),
			knownvalue.ListExact([]knownvalue.Check{
				knownvalue.StringRegexp(urnRe),
			})),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("certificate_pem"),
			knownvalue.StringRegexp(pemRe)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("status"),
			knownvalue.StringExact("READY")),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("id"),
			knownvalue.NotNull()),
	}
}

func kmsCredStateCsrChecks(resName string) []statecheck.StateCheck {
	csrRe := regexp.MustCompile("-----BEGIN CERTIFICATE REQUEST-----.*")
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("csr"),
			knownvalue.StringRegexp(csrRe)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("private_key_pem"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("from_csr"),
			knownvalue.Bool(true)),
	}
}

func kmsCredStateNoCsrChecks(resName string) []statecheck.StateCheck {
	keyRe := regexp.MustCompile("-----BEGIN EC PRIVATE KEY-----.*")
	return []statecheck.StateCheck{
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("csr"),
			knownvalue.Null()),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("private_key_pem"),
			knownvalue.StringRegexp(keyRe)),
		statecheck.ExpectKnownValue(
			resName,
			tfjsonpath.New("from_csr"),
			knownvalue.Bool(false)),
	}
}

const confOkmsCredTest = `
data "ovh_me" "current_account" {
}

resource "ovh_okms" "kms" {
  ovh_subsidiary = "FR"
  display_name = "%[1]s"
  region = "EU_WEST_SBG"
}

resource "ovh_okms_credential" "cred" {
  okms_id = ovh_okms.kms.id
  name = "%[2]s"
  identity_urns = ["urn:v1:eu:identity:account:${data.ovh_me.current_account.nichandle}"]
}

resource "ovh_okms_credential" "credcsr" {
  okms_id = ovh_okms.kms.id
  name = "%[2]scsr"
  identity_urns = ["urn:v1:eu:identity:account:${data.ovh_me.current_account.nichandle}"]
  csr = <<EOT
-----BEGIN CERTIFICATE REQUEST-----
MIHLMHQCAQAwEjEQMA4GA1UEAwwHdGZ0ZXN0czBZMBMGByqGSM49AgEGCCqGSM49
AwEHA0IABKSlsmzuBjHxShCzX4L+esOR26pIbC89yekqT8vu6pbt1HxkhbCruKtD
LHKUsS1azH2KnLMquZCpDUhMJzdlPS6gADAKBggqhkjOPQQDAgNHADBEAiAb6TfV
uzoo2iM0ECxHzpVhEIGHSAdveyqPK8luLm9gmwIgUPgA8D/R3F/dvDtjBPuWkxwd
ujvZmkiGKQNew6HU1Q4=
-----END CERTIFICATE REQUEST-----
EOT
}
`

func TestAccResourceOkmsCredCreate(t *testing.T) {
	kmsName := acctest.RandomWithPrefix(test_prefix)
	credName := acctest.RandomWithPrefix(test_prefix)

	checks := kmsCredStateCommonChecks("ovh_okms_credential.credcsr", credName+"csr")
	checks = append(checks, kmsCredStateCommonChecks("ovh_okms_credential.cred", credName)...)
	checks = append(checks, kmsCredStateNoCsrChecks("ovh_okms_credential.cred")...)
	checks = append(checks, kmsCredStateCsrChecks("ovh_okms_credential.credcsr")...)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOrderOkms(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:            fmt.Sprintf(confOkmsCredTest, kmsName, credName),
				ConfigStateChecks: checks,
			},
		},
	})
}

const confOkmsCredImport = `
resource "ovh_okms_credential" "cred" {
  okms_id = "%[1]s"
}
`

func TestAccResourceOkmsCredImport(t *testing.T) {
	kmsId := os.Getenv("OVH_OKMS")
	credId := os.Getenv("OVH_OKMS_CREDENTIAL")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckOkmsCredential(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName:  "ovh_okms_credential.cred",
				ImportState:   true,
				ImportStateId: fmt.Sprintf("%s/%s", kmsId, credId),
				Config:        fmt.Sprintf(confOkmsCredImport, kmsId),
			},
		},
	})
}
