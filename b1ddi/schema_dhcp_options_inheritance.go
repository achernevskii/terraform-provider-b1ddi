// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcDHCPOptionsInheritance DHCPOptionsInheritance
//
// The inheritance configuration that specifies how the _dhcp_options_ field is inherited from the parent object.
//
// swagger:model ipamsvcDHCPOptionsInheritance
func schemaIpamsvcDHCPOptionsInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			// The inheritance configuration for the _dhcp_options_ field.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPOptionList(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for the _dhcp_options_ field.",
			},
		},
	}
}

func flattenIpamsvcDHCPOptionsInheritance(r *models.IpamsvcDHCPOptionsInheritance) []interface{} {
	if r == nil {
		return nil
	}
	return []interface{}{
		map[string]interface{}{
			"dhcp_options": flattenIpamsvcInheritedDHCPOptionList(r.DhcpOptions),
		},
	}
}

func expandIpamsvcDHCPOptionsInheritance(d []interface{}) *models.IpamsvcDHCPOptionsInheritance {
	if len(d) == 0 || d[0] == nil {
		return nil
	}
	in := d[0].(map[string]interface{})

	return &models.IpamsvcDHCPOptionsInheritance{
		DhcpOptions: expandIpamsvcInheritedDHCPOptionList(in["dhcp_options"].([]interface{})),
	}
}