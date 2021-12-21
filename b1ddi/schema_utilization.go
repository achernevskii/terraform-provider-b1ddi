// Code generated by go-swagger; DO NOT EDIT.

package b1ddi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/models"
)

// IpamsvcUtilization Utilization
//
// The __Utilization__ object represents IP address usage statistics for an object.
//
// swagger:model ipamsvcUtilization
func schemaIpamsvcUtilization() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{

			// The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object.
			// Read Only: true
			// Maximum: 100
			"abandon_utilization": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The percentage of abandoned IP addresses relative to the total IP addresses available in the scope of the object.",
			},

			// The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).
			// Read Only: true
			"abandoned": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of IP addresses in the scope of the object which are in the abandoned state (issued by a DHCP server and then declined by the client).",
			},

			// The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.
			// Read Only: true
			"dynamic": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of IP addresses handed out by DHCP in the scope of the object. This includes all leased addresses, fixed addresses that are defined but not currently leased and abandoned leases.",
			},

			// The number of IP addresses available in the scope of the object.
			// Read Only: true
			"free": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of IP addresses available in the scope of the object.",
			},

			// The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.
			// Read Only: true
			"static": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of defined IP addresses such as reservations or DNS records. It can be computed as _static_ = _used_ - _dynamic_.",
			},

			// The total number of IP addresses available in the scope of the object.
			// Read Only: true
			"total": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total number of IP addresses available in the scope of the object.",
			},

			// The number of IP addresses used in the scope of the object.
			// Read Only: true
			"used": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The number of IP addresses used in the scope of the object.",
			},

			// The percentage of used IP addresses relative to the total IP addresses available in the scope of the object.
			// Read Only: true
			// Maximum: 100
			"utilization": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The percentage of used IP addresses relative to the total IP addresses available in the scope of the object.",
			},
		},
	}
}

func flattenIpamsvcUtilization(r *models.IpamsvcUtilization) []interface{} {
	if r == nil {
		return []interface{}{}
	}

	return []interface{}{
		map[string]interface{}{
			"abandon_utilization": r.AbandonUtilization,
			"abandoned":           r.Abandoned,
			"dynamic":             r.Dynamic,
			"free":                r.Free,
			"static":              r.Static,
			"total":               r.Total,
			"used":                r.Used,
			"utilization":         r.Utilization,
		},
	}
}
