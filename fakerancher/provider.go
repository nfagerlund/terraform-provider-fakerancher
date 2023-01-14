package fakerancher

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{}, // who cares
		ResourcesMap: map[string]*schema.Resource{
			"fakerancher_cluster": resourceFakerancherCluster(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
