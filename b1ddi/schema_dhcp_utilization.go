// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcDHCPUtilization DHCPUtilization
//
// The __DHCPUtilization__ object represents DHCP utilization statistics for an object.
//
// swagger:model ipamsvcDHCPUtilization
func schemaIpamsvcDHCPUtilization() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_.
			// Read Only: true
			"dhcp_free": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total free IP addresses in the DHCP ranges in the scope of this object. It can be computed as _dhcp_total_ - _dhcp_used_.",
			},

			// The total IP addresses available in the DHCP ranges in the scope of this object.
			// Read Only: true
			"dhcp_total": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total IP addresses available in the DHCP ranges in the scope of this object.",
			},

			// The total IP addresses marked as used in the DHCP ranges in the scope of this object.
			// Read Only: true
			"dhcp_used": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total IP addresses marked as used in the DHCP ranges in the scope of this object.",
			},

			// The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object.
			// Read Only: true
			// Maximum: 100
			"dhcp_utilization": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The percentage of used IP addresses relative to the total IP addresses available in the DHCP ranges in the scope of this object.",
			},
		},
	}
}

func flattenIpamsvcDHCPUtilization(r *models.IpamsvcDHCPUtilization) []interface{} {
	if r == nil {
		return []interface{}{}
	}
	return []interface{}{
		map[string]interface{}{
			"dhcp_free":        r.DhcpFree,
			"dhcp_total":       r.DhcpTotal,
			"dhcp_used":        r.DhcpUsed,
			"dhcp_utilization": r.DhcpUtilization,
		},
	}

}
