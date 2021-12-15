package b1ddi

import (
	"context"
	"github.com/go-openapi/swag"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc/subnet"
	"github.com/infobloxopen/b1ddi-go-client/models"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// IpamsvcSubnet Subnet
//
// A __Subnet__ object (_ipam/subnet_) is a set of contiguous IP addresses in the same IP space with no gap, expressed as an address and CIDR values. It represents a set of addresses from which addresses are assigned to network equipment interfaces.
//
// swagger:model ipamsvcSubnet
func resourceIpamsvcSubnet() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpamsvcSubnetCreate,
		ReadContext:   resourceIpamsvcSubnetRead,
		UpdateContext: resourceIpamsvcSubnetUpdate,
		DeleteContext: resourceIpamsvcSubnetDelete,
		Schema: map[string]*schema.Schema{

			// The address of the subnet in the form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.
			// Required: true
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The address of the subnet in the form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
			},

			// The Automated Scope Management configuration for the subnet.
			"asm_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcASMConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The Automated Scope Management configuration for the subnet.",
			},

			// Set to 1 to indicate that the subnet may run out of addresses.
			// Read Only: true
			"asm_scope_flag": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Set to 1 to indicate that the subnet may run out of addresses.",
			},

			// The CIDR of the subnet. This is required if _address_ does not include CIDR.
			// Maximum: 32
			// Minimum: 1
			"cidr": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The CIDR of the subnet. This is required if _address_ does not include CIDR.",
			},

			// The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8.
			"comment": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8.",
			},

			// Time when the object has been created.
			// Read Only: true
			// Format: date-time
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been created.",
			},

			// Controls who does the DDNS updates.
			//
			// Valid values are:
			// * _client_: DHCP server updates DNS if requested by client.
			// * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _ignore_: DHCP server always updates DNS, even if the client says not to.
			// * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.
			// * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.
			//
			// Defaults to _client_.
			"ddns_client_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Controls who does the DDNS updates.\n\nValid values are:\n* _client_: DHCP server updates DNS if requested by client.\n* _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _ignore_: DHCP server always updates DNS, even if the client says not to.\n* _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates.\n* _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.\n\nDefaults to _client_.",
			},

			// The domain suffix for DDNS updates. FQDN, may be empty.
			//
			// Defaults to empty.
			"ddns_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain suffix for DDNS updates. FQDN, may be empty.\n\nDefaults to empty.",
			},

			// Indicates if DDNS needs to generate a hostname when not supplied by the client.
			//
			// Defaults to _false_.
			"ddns_generate_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if DDNS needs to generate a hostname when not supplied by the client.\n\nDefaults to _false_.",
			},

			// The prefix used in the generation of an FQDN.
			//
			// When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].
			// where address-text is simply the lease IP address converted to a hyphenated string.
			//
			// Defaults to "myhost".
			"ddns_generated_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The prefix used in the generation of an FQDN.\n\nWhen generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix].\nwhere address-text is simply the lease IP address converted to a hyphenated string.\n\nDefaults to \"myhost\".",
			},

			// Determines if DDNS updates are enabled at the subnet level.
			// Defaults to _true_.
			"ddns_send_updates": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determines if DDNS updates are enabled at the subnet level.\nDefaults to _true_.",
			},

			// Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.
			//
			// Defaults to _false_.
			"ddns_update_on_renew": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.\n\nDefaults to _false_.",
			},

			// When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.
			//
			// When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.
			//
			// Defaults to _true_.
			"ddns_use_conflict_resolution": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.\n\nWhen false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.\n\nDefaults to _true_.",
			},

			// The DHCP configuration of the subnet that controls how leases are issued.
			"dhcp_config": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPConfig(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP configuration of the subnet that controls how leases are issued.",
			},

			// The resource identifier.
			"dhcp_host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The DHCP options of the subnet. This can either be a specific option or a group of options.
			"dhcp_options": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcOptionItem(),
				Optional:    true,
				Description: "The DHCP options of the subnet. This can either be a specific option or a group of options.",
			},

			// The utilization of IP addresses within the DHCP ranges of the subnet.
			// Read Only: true
			// ToDo add dhcp_utilization
			//"dhcp_utilization": {
			//	Type:        schema.TypeList,
			//	Elem:        schemaIpamsvcDHCPUtilization(),
			//	Computed:    true,
			//	Description: "The utilization of IP addresses within the DHCP ranges of the subnet.",
			//},

			// The configuration for header option filename field.
			"header_option_filename": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option filename field.",
			},

			// The configuration for header option server address field.
			"header_option_server_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server address field.",
			},

			// The configuration for header option server name field.
			"header_option_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The configuration for header option server name field.",
			},

			// The character to replace non-matching characters with, when hostname rewrite is enabled.
			//
			// Any single ASCII character.
			//
			// Defaults to "_".
			"hostname_rewrite_char": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The character to replace non-matching characters with, when hostname rewrite is enabled.\n\nAny single ASCII character.\n\nDefaults to \"_\".",
			},

			// Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.
			//
			// Defaults to _false_.
			"hostname_rewrite_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.\n\nDefaults to _false_.",
			},

			// The regex bracket expression to match valid characters.
			//
			// Must begin with "[" and end with "]" and be a compilable POSIX regex.
			//
			// Defaults to "[^a-zA-Z0-9_.]".
			"hostname_rewrite_regex": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The regex bracket expression to match valid characters.\n\nMust begin with \"[\" and end with \"]\" and be a compilable POSIX regex.\n\nDefaults to \"[^a-zA-Z0-9_.]\".",
			},

			// The list of the inheritance assigned hosts of the object.
			// Read Only: true
			// ToDo add inheritance_assigned_hosts
			//"inheritance_assigned_hosts": {
			//	Type:        schema.TypeList,
			//	Elem:        schemaInheritanceAssignedHost(),
			//	Computed:    true,
			//	Description: "The list of the inheritance assigned hosts of the object.",
			//},

			// The resource identifier.
			"inheritance_parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The DHCP inheritance configuration for the subnet.
			"inheritance_sources": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcDHCPInheritance(),
				MaxItems:    1,
				Optional:    true,
				Description: "The DHCP inheritance configuration for the subnet.",
			},

			// The name of the subnet. May contain 1 to 256 characters. Can include UTF-8.
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the subnet. May contain 1 to 256 characters. Can include UTF-8.",
			},

			// The resource identifier.
			"parent": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The resource identifier.",
			},

			// The type of protocol of the subnet (_ipv4_ or _ipv6_).
			// Read Only: true
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of protocol of the subnet (_ipv4_ or _ipv6_).",
			},

			// The resource identifier.
			// Required: true
			"space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The resource identifier.",
			},

			// The tags for the subnet in JSON format.
			"tags": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The tags for the subnet in JSON format.",
			},

			// The IP address utilization threshold settings for the subnet.
			"threshold": {
				Type:        schema.TypeList,
				Elem:        schemaIpamsvcUtilizationThreshold(),
				MaxItems:    1,
				Optional:    true,
				Description: "The IP address utilization threshold settings for the subnet.",
			},

			// Time when the object has been updated. Equals to _created_at_ if not updated after creation.
			// Read Only: true
			// Format: date-time
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
			},

			// The IP address utilization statistics of the subnet.
			// Read Only: true
			// ToDo Add utilization
			//"utilization": {
			//	Type:        schema.TypeList,
			//	Elem:        schemaIpamsvcUtilization(),
			//	Computed:    true,
			//	Description: "The IP address utilization statistics of the subnet.",
			//},
		},
	}
}

func resourceIpamsvcSubnetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	s := &models.IpamsvcSubnet{
		Address: swag.String(d.Get("address").(string)),
		Name:    d.Get("name").(string),
		Cidr:    int64(d.Get("cidr").(int)),
		Space:   swag.String(d.Get("space").(string)),
		Comment: d.Get("comment").(string),
	}

	resp, err := c.Subnet.SubnetCreate(&subnet.SubnetCreateParams{Body: s, Context: ctx}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.Payload.Result.ID)

	resourceIpamsvcSubnetRead(ctx, d, m)

	return diags
}

func resourceIpamsvcSubnetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	s, err := c.Subnet.SubnetRead(
		&subnet.SubnetReadParams{
			ID:      d.Id(),
			Context: ctx,
		},
		nil,
	)
	if err != nil {
		diag.FromErr(err)
	}

	err = d.Set("address", *s.Payload.Result.Address)
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("cidr", s.Payload.Result.Cidr)
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("protocol", s.Payload.Result.Protocol)
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("space", s.Payload.Result.Space)
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("name", s.Payload.Result.Name)
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("created_at", s.Payload.Result.CreatedAt.String())
	if err != nil {
		diag.FromErr(err)
	}
	err = d.Set("updated_at", s.Payload.Result.UpdatedAt.String())
	if err != nil {
		diag.FromErr(err)
	}

	return diags
}

func resourceIpamsvcSubnetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// ToDo Implement resourceIpamsvcSubnetUpdate function
	return diags
}

func resourceIpamsvcSubnetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*ipamsvc.IPAddressManagementAPI)

	var diags diag.Diagnostics

	_, err := c.Subnet.SubnetDelete(&subnet.SubnetDeleteParams{ID: d.Id(), Context: ctx}, nil)
	if err != nil {
		diag.FromErr(err)
	}

	d.SetId("")

	return diags
}