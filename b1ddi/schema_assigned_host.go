// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// InheritanceAssignedHost AssignedHost
//
// _ddi/assigned_host_ represents a BloxOne DDI host assigned to one of the following:
//  * Subnet (_ipam/subnet_)
//  * Range (_ipam/range_)
//  * Fixed Address (_dhcp/fixed_address_)
//  * Authoritative Zone (_dns/auth_zone_)
//
// swagger:model inheritanceAssignedHost
func schemaInheritanceAssignedHost() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The human-readable display name for the host referred to by _ophid_.
			// Read Only: true
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The human-readable display name for the host referred to by _ophid_.",
			},

			// The resource identifier.
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The on-prem host ID.
			// Read Only: true
			"ophid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The on-prem host ID.",
			},
		},
	}
}