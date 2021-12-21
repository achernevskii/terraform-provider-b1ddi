// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcFixedAddressInheritance FixedAddressInheritance
//
// The __FixedAddressInheritance__ object specifies how and which fields _FixedAddress_ object inherits from the parent.
//
// swagger:model ipamsvcFixedAddressInheritance
func schemaIpamsvcFixedAddressInheritance() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The inheritance configuration for _dhcp_options_ field.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcInheritedDHCPOptionList(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _dhcp_options_ field.",
			},

			// The inheritance configuration for _header_option_filename_ field.
			"header_option_filename": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_filename_ field.",
			},

			// The inheritance configuration for _header_option_server_address_ field.
			"header_option_server_address": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_address_ field.",
			},

			// The inheritance configuration for _header_option_server_name_ field.
			"header_option_server_name": {
				Type:        schema.TypeList,
				Elem:        schemaInheritanceInheritedString(),
				MaxItems:    1,
				Optional:    true,
				Description: "The inheritance configuration for _header_option_server_name_ field.",
			},
		},
	}
}
