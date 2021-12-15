// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcInheritedDDNSHostnameBlock InheritedDDNSHostnameBlock
//
// The inheritance configuration block for dynamic DNS hostname.
//
// swagger:model ipamsvcInheritedDDNSHostnameBlock
func schemaIpamsvcInheritedDDNSHostnameBlock() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance setting.
			//
			// Valid values are:
			// * _inherit_: Use the inherited value.
			// * _override_: Use the value set in the object.
			//
			// Defaults to _inherit_.
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The inheritance setting.\n\nValid values are:\n* _inherit_: Use the inherited value.\n* _override_: Use the value set in the object.\n\nDefaults to _inherit_.",
			},

			// The human-readable display name for the object referred to by _source_.
			// Read Only: true
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable display name for the object referred to by _source_.",
			},

			// The resource identifier.
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The inherited value.
			// Read Only: true
			"value": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDDNSHostnameBlock(),
				Computed:    true,
				Description: "The inherited value.",
			},
		},
	}
}

func flattenIpamsvcInheritedDDNSHostnameBlock(r *models.IpamsvcInheritedDDNSHostnameBlock) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	res := make(map[string]interface{})

	res["action"] = r.Action
	res["display_name"] = r.DisplayName
	res["source"] = r.Source
	res["value"] = flattenIpamsvcDDNSHostnameBlock(r.Value)

	return []interface{}{res}
}