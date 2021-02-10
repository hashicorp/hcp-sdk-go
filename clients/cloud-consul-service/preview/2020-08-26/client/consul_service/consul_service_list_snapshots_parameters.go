// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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

// NewConsulServiceListSnapshotsParams creates a new ConsulServiceListSnapshotsParams object
// with the default values initialized.
func NewConsulServiceListSnapshotsParams() *ConsulServiceListSnapshotsParams {
	var ()
	return &ConsulServiceListSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsulServiceListSnapshotsParamsWithTimeout creates a new ConsulServiceListSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsulServiceListSnapshotsParamsWithTimeout(timeout time.Duration) *ConsulServiceListSnapshotsParams {
	var ()
	return &ConsulServiceListSnapshotsParams{

		timeout: timeout,
	}
}

// NewConsulServiceListSnapshotsParamsWithContext creates a new ConsulServiceListSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsulServiceListSnapshotsParamsWithContext(ctx context.Context) *ConsulServiceListSnapshotsParams {
	var ()
	return &ConsulServiceListSnapshotsParams{

		Context: ctx,
	}
}

// NewConsulServiceListSnapshotsParamsWithHTTPClient creates a new ConsulServiceListSnapshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsulServiceListSnapshotsParamsWithHTTPClient(client *http.Client) *ConsulServiceListSnapshotsParams {
	var ()
	return &ConsulServiceListSnapshotsParams{
		HTTPClient: client,
	}
}

/*ConsulServiceListSnapshotsParams contains all the parameters to send to the API endpoint
for the consul service list snapshots operation typically these are written to a http.Request
*/
type ConsulServiceListSnapshotsParams struct {

	/*PaginationNextPageToken
	  Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.

	*/
	PaginationNextPageToken *string
	/*PaginationPageSize
	  The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	*/
	PaginationPageSize *int64
	/*PaginationPreviousPageToken
	  Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.

	*/
	PaginationPreviousPageToken *string
	/*ResourceDescription
	  description is a human-friendly description for this link. This is
	used primarily for informational purposes such as error messages.

	*/
	ResourceDescription *string
	/*ResourceID
	  id is the identifier for this resource.

	*/
	ResourceID *string
	/*ResourceLocationOrganizationID
	  organization_id is the id of the organization.

	*/
	ResourceLocationOrganizationID string
	/*ResourceLocationProjectID
	  project_id is the projects id.

	*/
	ResourceLocationProjectID string
	/*ResourceLocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	ResourceLocationRegionProvider *string
	/*ResourceLocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	ResourceLocationRegionRegion *string
	/*ResourceType
	  type is the unique type of the resource. Each service publishes a
	unique set of types. The type value is recommended to be formatted
	in "<org>.<type>" such as "hashicorp.hvn". This is to prevent conflicts
	in the future, but any string value will work.

	*/
	ResourceType *string
	/*ResourceUUID
	  uuid is the unique UUID for this resource.

	*/
	ResourceUUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithTimeout(timeout time.Duration) *ConsulServiceListSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithContext(ctx context.Context) *ConsulServiceListSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithHTTPClient(client *http.Client) *ConsulServiceListSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ConsulServiceListSnapshotsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithPaginationPageSize(paginationPageSize *int64) *ConsulServiceListSnapshotsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ConsulServiceListSnapshotsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithResourceDescription adds the resourceDescription to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceDescription(resourceDescription *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceDescription(resourceDescription)
	return o
}

// SetResourceDescription adds the resourceDescription to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceDescription(resourceDescription *string) {
	o.ResourceDescription = resourceDescription
}

// WithResourceID adds the resourceID to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceID(resourceID *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceID(resourceID *string) {
	o.ResourceID = resourceID
}

// WithResourceLocationOrganizationID adds the resourceLocationOrganizationID to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceLocationOrganizationID(resourceLocationOrganizationID string) *ConsulServiceListSnapshotsParams {
	o.SetResourceLocationOrganizationID(resourceLocationOrganizationID)
	return o
}

// SetResourceLocationOrganizationID adds the resourceLocationOrganizationId to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceLocationOrganizationID(resourceLocationOrganizationID string) {
	o.ResourceLocationOrganizationID = resourceLocationOrganizationID
}

// WithResourceLocationProjectID adds the resourceLocationProjectID to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceLocationProjectID(resourceLocationProjectID string) *ConsulServiceListSnapshotsParams {
	o.SetResourceLocationProjectID(resourceLocationProjectID)
	return o
}

// SetResourceLocationProjectID adds the resourceLocationProjectId to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceLocationProjectID(resourceLocationProjectID string) {
	o.ResourceLocationProjectID = resourceLocationProjectID
}

// WithResourceLocationRegionProvider adds the resourceLocationRegionProvider to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceLocationRegionProvider(resourceLocationRegionProvider *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceLocationRegionProvider(resourceLocationRegionProvider)
	return o
}

// SetResourceLocationRegionProvider adds the resourceLocationRegionProvider to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceLocationRegionProvider(resourceLocationRegionProvider *string) {
	o.ResourceLocationRegionProvider = resourceLocationRegionProvider
}

// WithResourceLocationRegionRegion adds the resourceLocationRegionRegion to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceLocationRegionRegion(resourceLocationRegionRegion *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceLocationRegionRegion(resourceLocationRegionRegion)
	return o
}

// SetResourceLocationRegionRegion adds the resourceLocationRegionRegion to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceLocationRegionRegion(resourceLocationRegionRegion *string) {
	o.ResourceLocationRegionRegion = resourceLocationRegionRegion
}

// WithResourceType adds the resourceType to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceType(resourceType *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceType(resourceType)
	return o
}

// SetResourceType adds the resourceType to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceType(resourceType *string) {
	o.ResourceType = resourceType
}

// WithResourceUUID adds the resourceUUID to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) WithResourceUUID(resourceUUID *string) *ConsulServiceListSnapshotsParams {
	o.SetResourceUUID(resourceUUID)
	return o
}

// SetResourceUUID adds the resourceUuid to the consul service list snapshots params
func (o *ConsulServiceListSnapshotsParams) SetResourceUUID(resourceUUID *string) {
	o.ResourceUUID = resourceUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ConsulServiceListSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ResourceDescription != nil {

		// query param resource.description
		var qrResourceDescription string
		if o.ResourceDescription != nil {
			qrResourceDescription = *o.ResourceDescription
		}
		qResourceDescription := qrResourceDescription
		if qResourceDescription != "" {
			if err := r.SetQueryParam("resource.description", qResourceDescription); err != nil {
				return err
			}
		}

	}

	if o.ResourceID != nil {

		// query param resource.id
		var qrResourceID string
		if o.ResourceID != nil {
			qrResourceID = *o.ResourceID
		}
		qResourceID := qrResourceID
		if qResourceID != "" {
			if err := r.SetQueryParam("resource.id", qResourceID); err != nil {
				return err
			}
		}

	}

	// path param resource.location.organization_id
	if err := r.SetPathParam("resource.location.organization_id", o.ResourceLocationOrganizationID); err != nil {
		return err
	}

	// path param resource.location.project_id
	if err := r.SetPathParam("resource.location.project_id", o.ResourceLocationProjectID); err != nil {
		return err
	}

	if o.ResourceLocationRegionProvider != nil {

		// query param resource.location.region.provider
		var qrResourceLocationRegionProvider string
		if o.ResourceLocationRegionProvider != nil {
			qrResourceLocationRegionProvider = *o.ResourceLocationRegionProvider
		}
		qResourceLocationRegionProvider := qrResourceLocationRegionProvider
		if qResourceLocationRegionProvider != "" {
			if err := r.SetQueryParam("resource.location.region.provider", qResourceLocationRegionProvider); err != nil {
				return err
			}
		}

	}

	if o.ResourceLocationRegionRegion != nil {

		// query param resource.location.region.region
		var qrResourceLocationRegionRegion string
		if o.ResourceLocationRegionRegion != nil {
			qrResourceLocationRegionRegion = *o.ResourceLocationRegionRegion
		}
		qResourceLocationRegionRegion := qrResourceLocationRegionRegion
		if qResourceLocationRegionRegion != "" {
			if err := r.SetQueryParam("resource.location.region.region", qResourceLocationRegionRegion); err != nil {
				return err
			}
		}

	}

	if o.ResourceType != nil {

		// query param resource.type
		var qrResourceType string
		if o.ResourceType != nil {
			qrResourceType = *o.ResourceType
		}
		qResourceType := qrResourceType
		if qResourceType != "" {
			if err := r.SetQueryParam("resource.type", qResourceType); err != nil {
				return err
			}
		}

	}

	if o.ResourceUUID != nil {

		// query param resource.uuid
		var qrResourceUUID string
		if o.ResourceUUID != nil {
			qrResourceUUID = *o.ResourceUUID
		}
		qResourceUUID := qrResourceUUID
		if qResourceUUID != "" {
			if err := r.SetQueryParam("resource.uuid", qResourceUUID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}