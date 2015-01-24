package conoha

import (
	"fmt"
	"log"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/objectstorage/v1/containers"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceConohaContainer() *schema.Resource {
	return &schema.Resource{
		Create: resourceConohaContainerCreate,
		Read:   resourceConohaContainerRead,
		Delete: resourceConohaContainerDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceConohaContainerCreate(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*gophercloud.ProviderClient)

	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	name := d.Get("name").(string)
	// TODO: Use metadata
	opts := containers.CreateOpts{
		Metadata: map[string]string{
			"name": name,
		},
	}
	_, err = containers.Create(client, name, opts).ExtractHeader()
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	d.Set("name", name)
	d.SetId(name)
	log.Printf("[INFO] ConoHa container ID: %s", d.Id())

	return resourceConohaContainerRead(d, meta)
}

func resourceConohaContainerRead(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*gophercloud.ProviderClient)

	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	name := d.Get("name").(string)
	_, err = containers.Get(client, name).ExtractHeader()
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	d.Set("name", name)
	d.SetId(name)

	return nil
}

func resourceConohaContainerDelete(d *schema.ResourceData, meta interface{}) error {
	provider := meta.(*gophercloud.ProviderClient)

	client, err := openstack.NewObjectStorageV1(provider, gophercloud.EndpointOpts{
		Region: "RegionOne",
	})
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	name := d.Get("name").(string)
	_, err = containers.Delete(client, name).ExtractHeader()
	if err != nil {
		return fmt.Errorf("error %s", err)
	}

	return nil
}
