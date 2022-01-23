package b1ddi

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/address"
	"testing"
)

func TestAccResourceAddress_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBasicTestStep(),
		},
	})
}

func resourceAddressBasicTestStep() resource.TestStep {
	return resource.TestStep{
		Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						space = b1ddi_ip_space.tf_acc_test_space.id
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
		Check: resource.ComposeAggregateTestCheckFunc(
			testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
			testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
			testAccAddressExists("b1ddi_address.tf_acc_test_address"),
			testCheckAddressInSpace("b1ddi_address.tf_acc_test_address", "b1ddi_ip_space.tf_acc_test_space"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "address", "192.168.1.15"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "comment", "This Address is created by terraform provider acceptance test"),
			resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "created_at"),
			resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "dhcp_info"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "host", ""),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "hwaddr", ""),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "interface", ""),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "names.%", "0"),
			resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "parent"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "protocol", "ip4"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "range", ""),
			resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "space"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "state", "used"),
			resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "tags"),
			resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "updated_at"),
			resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "usage.0", "IPAM RESERVED"),
		),
	}
}

func TestAccResourceAddress_full_config(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						#host = ""
						hwaddr = "00:00:00:00:00:00"
						interface = "eth0"
						#names {
						#	name = "tf_acc_test_address"
						#	type = "test_name"
						#}
						space = b1ddi_ip_space.tf_acc_test_space.id
						tags = {
							TestType = "Acceptance"
						}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testAccAddressExists("b1ddi_address.tf_acc_test_address"),
					testCheckAddressInSpace("b1ddi_address.tf_acc_test_address", "b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "comment", "This Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "created_at"),
					resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "dhcp_info"),
					// ToDo Add check for the custom host
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "host", ""),

					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "hwaddr", "00:00:00:00:00:00"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "interface", "eth0"),
					// ToDo Add check for custom names
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "names.%", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "parent"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "protocol", "ip4"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "range", ""),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "state", "used"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "updated_at"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "usage.0", "IPAM RESERVED"),
				),
			},
		},
	})
}

func TestAccResourceAddress_update_space(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_ip_space" "tf_acc_new_test_space" {
  						name = "tf_acc_new_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_new_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_new_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						space = b1ddi_ip_space.tf_acc_new_test_space.id
						depends_on = [b1ddi_subnet.tf_acc_new_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_new_test_subnet"),
					testAccAddressExists("b1ddi_address.tf_acc_test_address"),
					testCheckAddressInSpace("b1ddi_address.tf_acc_test_address", "b1ddi_ip_space.tf_acc_new_test_space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "comment", "This Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "created_at"),
					resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "dhcp_info"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "host", ""),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "hwaddr", ""),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "interface", ""),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "names.%", "0"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "parent"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "protocol", "ip4"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "range", ""),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "state", "used"),
					resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "tags"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "updated_at"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "usage.0", "IPAM RESERVED"),
				),
			},
		},
	})
}

func TestAccResourceAddress_update(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			resourceAddressBasicTestStep(),
			{
				Config: fmt.Sprintf(`
					resource "b1ddi_ip_space" "tf_acc_test_space" {
  						name = "tf_acc_test_space"
  						comment = "This IP Space is created by terraform provider acceptance test"
					}
					resource "b1ddi_subnet" "tf_acc_test_subnet" {
						name = "tf_acc_test_subnet"						
						address = "192.168.1.0"
						space = b1ddi_ip_space.tf_acc_test_space.id
						cidr = 24
  						comment = "This Subnet is created by terraform provider acceptance test"
					}
					resource "b1ddi_address" "tf_acc_test_address" {
						address = "192.168.1.15"
						comment = "This Address is created by terraform provider acceptance test"
						#host = ""
						hwaddr = "00:00:00:00:00:00"
						interface = "eth0"
						#names {
						#	name = "tf_acc_test_address"
						#	type = "test_name"
						#}
						space = b1ddi_ip_space.tf_acc_test_space.id
						tags = {
							TestType = "Acceptance"
						}
						depends_on = [b1ddi_subnet.tf_acc_test_subnet]
					}`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccAddressExists("b1ddi_address.tf_acc_test_address"),
					testCheckIPSpaceExists("b1ddi_ip_space.tf_acc_test_space"),
					testCheckSubnetExists("b1ddi_subnet.tf_acc_test_subnet"),
					testCheckAddressInSpace("b1ddi_address.tf_acc_test_address", "b1ddi_ip_space.tf_acc_test_space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "address", "192.168.1.15"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "comment", "This Address is created by terraform provider acceptance test"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "created_at"),
					resource.TestCheckNoResourceAttr("b1ddi_address.tf_acc_test_address", "dhcp_info"),
					// ToDo Add check for the custom host
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "host", ""),

					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "hwaddr", "00:00:00:00:00:00"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "interface", "eth0"),
					// ToDo Add check for custom names
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "names.%", "0"),

					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "parent"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "protocol", "ip4"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "range", ""),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "space"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "state", "used"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "tags.%", "1"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "tags.TestType", "Acceptance"),
					resource.TestCheckResourceAttrSet("b1ddi_address.tf_acc_test_address", "updated_at"),
					resource.TestCheckResourceAttr("b1ddi_address.tf_acc_test_address", "usage.0", "IPAM RESERVED"),
				),
			},
		},
	})
}

func testAccAddressExists(resPath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, found := s.RootModule().Resources[resPath]
		if !found {
			return fmt.Errorf("not found %s", resPath)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", resPath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.Address.AddressRead(
			&address.AddressReadParams{ID: res.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}
		if resp.Payload.Result.ID != res.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				res.Primary.ID)
		}
		return nil
	}
}

// Checks if the specified Address resides in the specified IP Space
func testCheckAddressInSpace(addressPath, spacePath string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		addressResource, found := s.RootModule().Resources[addressPath]
		if !found {
			return fmt.Errorf("not found %s", addressPath)
		}
		if addressResource.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", addressPath)
		}
		space, found := s.RootModule().Resources[spacePath]
		if !found {
			return fmt.Errorf("not found %s", spacePath)
		}
		if space.Primary.ID == "" {
			return fmt.Errorf("ID for %s is not set", spacePath)
		}

		cli := testAccProvider.Meta().(*b1ddiclient.Client)

		resp, err := cli.IPAddressManagementAPI.Address.AddressRead(
			&address.AddressReadParams{ID: addressResource.Primary.ID, Context: context.TODO()},
			nil,
		)
		if err != nil {
			return err
		}

		if resp.Payload.Result.ID != addressResource.Primary.ID {
			return fmt.Errorf(
				"'id' does not match: \n got: '%s', \nexpected: '%s'",
				resp.Payload.Result.ID,
				addressResource.Primary.ID)
		}

		if *resp.Payload.Result.Space != space.Primary.ID {
			return fmt.Errorf(
				"'space' does not match: \n got: '%s', \nexpected: '%s'",
				*resp.Payload.Result.Space,
				space.Primary.ID)
		}

		return nil
	}
}
