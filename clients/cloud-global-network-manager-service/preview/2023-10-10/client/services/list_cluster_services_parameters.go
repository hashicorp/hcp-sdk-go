// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewListClusterServicesParams creates a new ListClusterServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListClusterServicesParams() *ListClusterServicesParams {
	return &ListClusterServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListClusterServicesParamsWithTimeout creates a new ListClusterServicesParams object
// with the ability to set a timeout on a request.
func NewListClusterServicesParamsWithTimeout(timeout time.Duration) *ListClusterServicesParams {
	return &ListClusterServicesParams{
		timeout: timeout,
	}
}

// NewListClusterServicesParamsWithContext creates a new ListClusterServicesParams object
// with the ability to set a context for a request.
func NewListClusterServicesParamsWithContext(ctx context.Context) *ListClusterServicesParams {
	return &ListClusterServicesParams{
		Context: ctx,
	}
}

// NewListClusterServicesParamsWithHTTPClient creates a new ListClusterServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListClusterServicesParamsWithHTTPClient(client *http.Client) *ListClusterServicesParams {
	return &ListClusterServicesParams{
		HTTPClient: client,
	}
}

/*
ListClusterServicesParams contains all the parameters to send to the API endpoint

	for the list cluster services operation.

	Typically these are written to a http.Request.
*/
type ListClusterServicesParams struct {

	/* ClusterResourceName.

	   Cluster resource name will be of the form `consul/project/<project_id>/cluster/<cluster_id>`
	*/
	ClusterResourceName string

	/* ExternalSources.

	   Query param filter: `external_sources`. Returns a service when service's external sources overlap with the external sources specified.
	*/
	ExternalSources []string

	/* InMesh.

	   Query param filter: whether the service is in the service mesh or not. Possible values: `true` or `false`.
	*/
	InMesh *bool

	/* Kind.

	   Query param filter: `kind` of the service. This can be combination of `typical`, `connect-proxy`, `destination`, `api-gateway`, `ingress-gateway`, `terminating-gateway`, `mesh-gateway`.
	*/
	Kind []string

	/* MtlsMode.

	   Query param filter: mTLS mode of the service. Possible value: `permissive`, `strict`.
	*/
	MtlsMode []string

	/* Namespace.

	   Query param filter: `namespace` of the service.
	*/
	Namespace *string

	/* OrderBy.

	   Sorts the services based on a field. Allowed fields: `name`, `health`. The value needs to be of the format <{name/health} {asc/desc}> For example: `name asc`, `health desc`.
	*/
	OrderBy []string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this parameter to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The maximum number of results per page to return. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned, which you can use to get the next page of results in subsequent
	requests. A value of zero causes `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this parameter to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* Partition.

	   Query param filter: `partition` of the service.
	*/
	Partition *string

	/* Query.

	   Search query to filter by. Searches across service `name`, `cluster_id`, `partition`, `namespace`, `sameness_group`, `tags`
	*/
	Query *string

	/* Status.

	   Query param filter: `status` of the service. This can be combination of `passing`, `warning`, `critical`, `none`.
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list cluster services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterServicesParams) WithDefaults() *ListClusterServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list cluster services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListClusterServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list cluster services params
func (o *ListClusterServicesParams) WithTimeout(timeout time.Duration) *ListClusterServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list cluster services params
func (o *ListClusterServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list cluster services params
func (o *ListClusterServicesParams) WithContext(ctx context.Context) *ListClusterServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list cluster services params
func (o *ListClusterServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list cluster services params
func (o *ListClusterServicesParams) WithHTTPClient(client *http.Client) *ListClusterServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list cluster services params
func (o *ListClusterServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterResourceName adds the clusterResourceName to the list cluster services params
func (o *ListClusterServicesParams) WithClusterResourceName(clusterResourceName string) *ListClusterServicesParams {
	o.SetClusterResourceName(clusterResourceName)
	return o
}

// SetClusterResourceName adds the clusterResourceName to the list cluster services params
func (o *ListClusterServicesParams) SetClusterResourceName(clusterResourceName string) {
	o.ClusterResourceName = clusterResourceName
}

// WithExternalSources adds the externalSources to the list cluster services params
func (o *ListClusterServicesParams) WithExternalSources(externalSources []string) *ListClusterServicesParams {
	o.SetExternalSources(externalSources)
	return o
}

// SetExternalSources adds the externalSources to the list cluster services params
func (o *ListClusterServicesParams) SetExternalSources(externalSources []string) {
	o.ExternalSources = externalSources
}

// WithInMesh adds the inMesh to the list cluster services params
func (o *ListClusterServicesParams) WithInMesh(inMesh *bool) *ListClusterServicesParams {
	o.SetInMesh(inMesh)
	return o
}

// SetInMesh adds the inMesh to the list cluster services params
func (o *ListClusterServicesParams) SetInMesh(inMesh *bool) {
	o.InMesh = inMesh
}

// WithKind adds the kind to the list cluster services params
func (o *ListClusterServicesParams) WithKind(kind []string) *ListClusterServicesParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the list cluster services params
func (o *ListClusterServicesParams) SetKind(kind []string) {
	o.Kind = kind
}

// WithMtlsMode adds the mtlsMode to the list cluster services params
func (o *ListClusterServicesParams) WithMtlsMode(mtlsMode []string) *ListClusterServicesParams {
	o.SetMtlsMode(mtlsMode)
	return o
}

// SetMtlsMode adds the mtlsMode to the list cluster services params
func (o *ListClusterServicesParams) SetMtlsMode(mtlsMode []string) {
	o.MtlsMode = mtlsMode
}

// WithNamespace adds the namespace to the list cluster services params
func (o *ListClusterServicesParams) WithNamespace(namespace *string) *ListClusterServicesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the list cluster services params
func (o *ListClusterServicesParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithOrderBy adds the orderBy to the list cluster services params
func (o *ListClusterServicesParams) WithOrderBy(orderBy []string) *ListClusterServicesParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the list cluster services params
func (o *ListClusterServicesParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list cluster services params
func (o *ListClusterServicesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListClusterServicesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list cluster services params
func (o *ListClusterServicesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list cluster services params
func (o *ListClusterServicesParams) WithPaginationPageSize(paginationPageSize *int64) *ListClusterServicesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list cluster services params
func (o *ListClusterServicesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list cluster services params
func (o *ListClusterServicesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListClusterServicesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list cluster services params
func (o *ListClusterServicesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPartition adds the partition to the list cluster services params
func (o *ListClusterServicesParams) WithPartition(partition *string) *ListClusterServicesParams {
	o.SetPartition(partition)
	return o
}

// SetPartition adds the partition to the list cluster services params
func (o *ListClusterServicesParams) SetPartition(partition *string) {
	o.Partition = partition
}

// WithQuery adds the query to the list cluster services params
func (o *ListClusterServicesParams) WithQuery(query *string) *ListClusterServicesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list cluster services params
func (o *ListClusterServicesParams) SetQuery(query *string) {
	o.Query = query
}

// WithStatus adds the status to the list cluster services params
func (o *ListClusterServicesParams) WithStatus(status []string) *ListClusterServicesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list cluster services params
func (o *ListClusterServicesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListClusterServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_resource_name
	if err := r.SetPathParam("cluster_resource_name", o.ClusterResourceName); err != nil {
		return err
	}

	if o.ExternalSources != nil {

		// binding items for external_sources
		joinedExternalSources := o.bindParamExternalSources(reg)

		// query array param external_sources
		if err := r.SetQueryParam("external_sources", joinedExternalSources...); err != nil {
			return err
		}
	}

	if o.InMesh != nil {

		// query param in_mesh
		var qrInMesh bool

		if o.InMesh != nil {
			qrInMesh = *o.InMesh
		}
		qInMesh := swag.FormatBool(qrInMesh)
		if qInMesh != "" {

			if err := r.SetQueryParam("in_mesh", qInMesh); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// binding items for kind
		joinedKind := o.bindParamKind(reg)

		// query array param kind
		if err := r.SetQueryParam("kind", joinedKind...); err != nil {
			return err
		}
	}

	if o.MtlsMode != nil {

		// binding items for mtls_mode
		joinedMtlsMode := o.bindParamMtlsMode(reg)

		// query array param mtls_mode
		if err := r.SetQueryParam("mtls_mode", joinedMtlsMode...); err != nil {
			return err
		}
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.Partition != nil {

		// query param partition
		var qrPartition string

		if o.Partition != nil {
			qrPartition = *o.Partition
		}
		qPartition := qrPartition
		if qPartition != "" {

			if err := r.SetQueryParam("partition", qPartition); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListClusterServices binds the parameter external_sources
func (o *ListClusterServicesParams) bindParamExternalSources(formats strfmt.Registry) []string {
	externalSourcesIR := o.ExternalSources

	var externalSourcesIC []string
	for _, externalSourcesIIR := range externalSourcesIR { // explode []string

		externalSourcesIIV := externalSourcesIIR // string as string
		externalSourcesIC = append(externalSourcesIC, externalSourcesIIV)
	}

	// items.CollectionFormat: "multi"
	externalSourcesIS := swag.JoinByFormat(externalSourcesIC, "multi")

	return externalSourcesIS
}

// bindParamListClusterServices binds the parameter kind
func (o *ListClusterServicesParams) bindParamKind(formats strfmt.Registry) []string {
	kindIR := o.Kind

	var kindIC []string
	for _, kindIIR := range kindIR { // explode []string

		kindIIV := kindIIR // string as string
		kindIC = append(kindIC, kindIIV)
	}

	// items.CollectionFormat: "multi"
	kindIS := swag.JoinByFormat(kindIC, "multi")

	return kindIS
}

// bindParamListClusterServices binds the parameter mtls_mode
func (o *ListClusterServicesParams) bindParamMtlsMode(formats strfmt.Registry) []string {
	mtlsModeIR := o.MtlsMode

	var mtlsModeIC []string
	for _, mtlsModeIIR := range mtlsModeIR { // explode []string

		mtlsModeIIV := mtlsModeIIR // string as string
		mtlsModeIC = append(mtlsModeIC, mtlsModeIIV)
	}

	// items.CollectionFormat: "multi"
	mtlsModeIS := swag.JoinByFormat(mtlsModeIC, "multi")

	return mtlsModeIS
}

// bindParamListClusterServices binds the parameter order_by
func (o *ListClusterServicesParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "multi"
	orderByIS := swag.JoinByFormat(orderByIC, "multi")

	return orderByIS
}

// bindParamListClusterServices binds the parameter status
func (o *ListClusterServicesParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "multi"
	statusIS := swag.JoinByFormat(statusIC, "multi")

	return statusIS
}
