package conoha

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/objectstorage/v1/containers"
)

func TestAccConohaContainer_Basic(t *testing.T) {
	var container containers.Container

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConohaContainerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckConohaContainerConfig_basic),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckConohaContainerExists("conoha_container.foobar", &container),
					testAccCheckConohaContainerAttributes(&container),
					resource.TestCheckResourceAttr(
						"conoha_container.foobar", "name", "foo"),
				),
			},
		},
	})
}

func TestAccConohaContainer_Updated(t *testing.T) {
	var container containers.Container

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckConohaContainerDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckConohaContainerConfig_basic),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckConohaContainerExists("conoha_container.foobar", &container),
					testAccCheckConohaContainerAttributes(&container),
					resource.TestCheckResourceAttr(
						"conoha_container.foobar", "name", "foo"),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(testAccCheckConohaContainerConfig_new_value),
				Check: resource.ComposeTestCheckFunc(
					// testAccCheckConohaContainerExists("conoha_container.foobar", &container),
					testAccCheckConohaContainerAttributesUpdated(&container),
					resource.TestCheckResourceAttr(
						"conoha_container.foobar", "name", "foo"),
				),
			},
		},
	})
}

func testAccCheckConohaContainerDestroy(s *terraform.State) error {
	provider := testAccProvider.Meta().(*gophercloud.ProviderClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "conoha_container" {
			continue
		}

		client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
			Region: "RegionOne",
		})
		if err != nil {
			return fmt.Errorf("error %s", err)
		}

		_, err = containers.Get(client, rs.Primary.Attributes["name"]).ExtractHeader()
		if err == nil {
			return fmt.Errorf("Record still exists")
		}
	}

	return nil
}

func testAccCheckConohaContainerAttributes(container *containers.Container) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if container.Name != "foo" {
			return fmt.Errorf("Bad content: %s", container.Name)
		}

		return nil
	}
}

func testAccCheckConohaContainerAttributesUpdated(container *containers.Container) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		if container.Name != "bar" {
			return fmt.Errorf("Bad content: %s", container.Name)
		}

		return nil
	}
}

func testAccCheckConohaContainerExists(n string, container *containers.Container) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No container ID is set")
		}

		provider := testAccProvider.Meta().(*gophercloud.ProviderClient)

		_, err = containers.Get(client, rs.Primary.Attributes["name"]).ExtractHeader()
		if err != nil {
			return err
		}

		if foundContainer.StringId() != rs.Primary.ID {
			return fmt.Errorf("container not found")
		}

		*container = *foundContainer

		return nil
	}
}

const testAccCheckConohaContainerConfig_basic = `
resource "conoha_container" "foobar" {
	name = "foo"
}`

const testAccCheckConohaContainerConfig_new_value = `
resource "conoha_container" "foobar" {
	name = "bar"
}`
