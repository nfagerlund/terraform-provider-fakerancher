package fakerancher

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFakerancherCluster() *schema.Resource {
	return &schema.Resource{
		Create: clusterCreate,
		Read:   clusterRead,
		Delete: clusterDelete,
		Update: clusterUpdate,
		Schema: clusterFields(),
	}
}

func clusterCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(fmt.Sprintf("%d", rand.Int()))
	return nil
}

func clusterRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func clusterDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func clusterUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}
