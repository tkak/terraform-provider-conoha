package conoha

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/tkak/conoha"
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

	client := meta.(*conoha.Client)

	c := &conoha.Container{
		Name: d.Get("name").(string),
	}

	err := client.CreateContainer(c)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	return resourceConohaContainerRead(d, meta)
}

func resourceConohaContainerRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*conoha.Client)

	c := &conoha.Container{
		Name: d.Get("name").(string),
	}

	err := client.ReadContainer(c)
	if err != nil {
		return fmt.Errorf("Error reading container: %s", err)
	}

	d.Set("name", c.Name)
	d.SetId(c.Name)

	return nil
}

func resourceConohaContainerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*conoha.Client)

	c := &conoha.Container{
		Name: d.Get("name").(string),
	}

	err := client.DeleteContainer(c)
	if err != nil {
		fmt.Printf("Error %s", err)
	}

	return nil
}
