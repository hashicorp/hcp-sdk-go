// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/activation_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/billing_account_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/contract_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/f_c_p_management_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/invoice_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/product_service"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-billing/preview/2020-11-05/client/statement_service"
)

// Default cloud billing HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.cloud.hashicorp.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new cloud billing HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CloudBilling {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloud billing HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CloudBilling {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloud billing client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CloudBilling {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CloudBilling)
	cli.Transport = transport
	cli.ActivationService = activation_service.New(transport, formats)
	cli.BillingAccountService = billing_account_service.New(transport, formats)
	cli.ContractService = contract_service.New(transport, formats)
	cli.FcpManagementService = f_c_p_management_service.New(transport, formats)
	cli.InvoiceService = invoice_service.New(transport, formats)
	cli.ProductService = product_service.New(transport, formats)
	cli.StatementService = statement_service.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CloudBilling is a client for cloud billing
type CloudBilling struct {
	ActivationService activation_service.ClientService

	BillingAccountService billing_account_service.ClientService

	ContractService contract_service.ClientService

	FcpManagementService f_c_p_management_service.ClientService

	InvoiceService invoice_service.ClientService

	ProductService product_service.ClientService

	StatementService statement_service.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CloudBilling) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.ActivationService.SetTransport(transport)
	c.BillingAccountService.SetTransport(transport)
	c.ContractService.SetTransport(transport)
	c.FcpManagementService.SetTransport(transport)
	c.InvoiceService.SetTransport(transport)
	c.ProductService.SetTransport(transport)
	c.StatementService.SetTransport(transport)
}
