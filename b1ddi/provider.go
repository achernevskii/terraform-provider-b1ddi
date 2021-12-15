package b1ddi

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/infobloxopen/b1ddi-go-client/ipamsvc"

	httptransport "github.com/go-openapi/runtime/client"
	b1ddiclient "github.com/infobloxopen/b1ddi-go-client/client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_HOST", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("B1DDI_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"b1ddi_ip_space":      resourceIpamsvcIPSpace(),
			"b1ddi_subnet":        resourceIpamsvcSubnet(),
			"b1ddi_fixed_address": resourceIpamsvcFixedAddress(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	host := d.Get("host").(string)
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if host == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without API host",
			// ToDo add detail Detail: "Please set token via"
		})
		return nil, diags
	}

	if token == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to initialise B1DDI client without API token",
			// ToDo add detail Detail: "Please set token via"
		})
		return nil, diags
	}

	// create the transport
	transport := httptransport.New(
		// Todo configure schemes via terraform
		host, "api/ddi/v1", nil,
	)

	// Create default auth header for all API requests
	tokenAuth := b1ddiclient.B1DDIToken(token)
	transport.DefaultAuthentication = tokenAuth

	// create the API client
	c := ipamsvc.New(transport, strfmt.Default)

	return c, diags
}