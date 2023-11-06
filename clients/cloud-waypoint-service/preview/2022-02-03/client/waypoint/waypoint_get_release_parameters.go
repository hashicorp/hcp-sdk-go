// Code generated by go-swagger; DO NOT EDIT.

package waypoint

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
)

// NewWaypointGetReleaseParams creates a new WaypointGetReleaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointGetReleaseParams() *WaypointGetReleaseParams {
	return &WaypointGetReleaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointGetReleaseParamsWithTimeout creates a new WaypointGetReleaseParams object
// with the ability to set a timeout on a request.
func NewWaypointGetReleaseParamsWithTimeout(timeout time.Duration) *WaypointGetReleaseParams {
	return &WaypointGetReleaseParams{
		timeout: timeout,
	}
}

// NewWaypointGetReleaseParamsWithContext creates a new WaypointGetReleaseParams object
// with the ability to set a context for a request.
func NewWaypointGetReleaseParamsWithContext(ctx context.Context) *WaypointGetReleaseParams {
	return &WaypointGetReleaseParams{
		Context: ctx,
	}
}

// NewWaypointGetReleaseParamsWithHTTPClient creates a new WaypointGetReleaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointGetReleaseParamsWithHTTPClient(client *http.Client) *WaypointGetReleaseParams {
	return &WaypointGetReleaseParams{
		HTTPClient: client,
	}
}

/*
WaypointGetReleaseParams contains all the parameters to send to the API endpoint

	for the waypoint get release operation.

	Typically these are written to a http.Request.
*/
type WaypointGetReleaseParams struct {

	/* LoadDetails.

	     Load additional details about the release. These will become available
	in the Preload section.

	     Default: "NONE"
	*/
	LoadDetails *string

	/* NamespaceID.

	   namespace_id is the id of the namespace.
	*/
	NamespaceID string

	// RefID.
	RefID string

	// RefSequenceApplicationApplication.
	RefSequenceApplicationApplication *string

	// RefSequenceApplicationProject.
	RefSequenceApplicationProject *string

	// RefSequenceNumber.
	//
	// Format: uint64
	RefSequenceNumber *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint get release params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetReleaseParams) WithDefaults() *WaypointGetReleaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint get release params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointGetReleaseParams) SetDefaults() {
	var (
		loadDetailsDefault = string("NONE")
	)

	val := WaypointGetReleaseParams{
		LoadDetails: &loadDetailsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the waypoint get release params
func (o *WaypointGetReleaseParams) WithTimeout(timeout time.Duration) *WaypointGetReleaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint get release params
func (o *WaypointGetReleaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint get release params
func (o *WaypointGetReleaseParams) WithContext(ctx context.Context) *WaypointGetReleaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint get release params
func (o *WaypointGetReleaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint get release params
func (o *WaypointGetReleaseParams) WithHTTPClient(client *http.Client) *WaypointGetReleaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint get release params
func (o *WaypointGetReleaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLoadDetails adds the loadDetails to the waypoint get release params
func (o *WaypointGetReleaseParams) WithLoadDetails(loadDetails *string) *WaypointGetReleaseParams {
	o.SetLoadDetails(loadDetails)
	return o
}

// SetLoadDetails adds the loadDetails to the waypoint get release params
func (o *WaypointGetReleaseParams) SetLoadDetails(loadDetails *string) {
	o.LoadDetails = loadDetails
}

// WithNamespaceID adds the namespaceID to the waypoint get release params
func (o *WaypointGetReleaseParams) WithNamespaceID(namespaceID string) *WaypointGetReleaseParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint get release params
func (o *WaypointGetReleaseParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithRefID adds the refID to the waypoint get release params
func (o *WaypointGetReleaseParams) WithRefID(refID string) *WaypointGetReleaseParams {
	o.SetRefID(refID)
	return o
}

// SetRefID adds the refId to the waypoint get release params
func (o *WaypointGetReleaseParams) SetRefID(refID string) {
	o.RefID = refID
}

// WithRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get release params
func (o *WaypointGetReleaseParams) WithRefSequenceApplicationApplication(refSequenceApplicationApplication *string) *WaypointGetReleaseParams {
	o.SetRefSequenceApplicationApplication(refSequenceApplicationApplication)
	return o
}

// SetRefSequenceApplicationApplication adds the refSequenceApplicationApplication to the waypoint get release params
func (o *WaypointGetReleaseParams) SetRefSequenceApplicationApplication(refSequenceApplicationApplication *string) {
	o.RefSequenceApplicationApplication = refSequenceApplicationApplication
}

// WithRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get release params
func (o *WaypointGetReleaseParams) WithRefSequenceApplicationProject(refSequenceApplicationProject *string) *WaypointGetReleaseParams {
	o.SetRefSequenceApplicationProject(refSequenceApplicationProject)
	return o
}

// SetRefSequenceApplicationProject adds the refSequenceApplicationProject to the waypoint get release params
func (o *WaypointGetReleaseParams) SetRefSequenceApplicationProject(refSequenceApplicationProject *string) {
	o.RefSequenceApplicationProject = refSequenceApplicationProject
}

// WithRefSequenceNumber adds the refSequenceNumber to the waypoint get release params
func (o *WaypointGetReleaseParams) WithRefSequenceNumber(refSequenceNumber *string) *WaypointGetReleaseParams {
	o.SetRefSequenceNumber(refSequenceNumber)
	return o
}

// SetRefSequenceNumber adds the refSequenceNumber to the waypoint get release params
func (o *WaypointGetReleaseParams) SetRefSequenceNumber(refSequenceNumber *string) {
	o.RefSequenceNumber = refSequenceNumber
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointGetReleaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.LoadDetails != nil {

		// query param load_details
		var qrLoadDetails string

		if o.LoadDetails != nil {
			qrLoadDetails = *o.LoadDetails
		}
		qLoadDetails := qrLoadDetails
		if qLoadDetails != "" {

			if err := r.SetQueryParam("load_details", qLoadDetails); err != nil {
				return err
			}
		}
	}

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	// path param ref.id
	if err := r.SetPathParam("ref.id", o.RefID); err != nil {
		return err
	}

	if o.RefSequenceApplicationApplication != nil {

		// query param ref.sequence.application.application
		var qrRefSequenceApplicationApplication string

		if o.RefSequenceApplicationApplication != nil {
			qrRefSequenceApplicationApplication = *o.RefSequenceApplicationApplication
		}
		qRefSequenceApplicationApplication := qrRefSequenceApplicationApplication
		if qRefSequenceApplicationApplication != "" {

			if err := r.SetQueryParam("ref.sequence.application.application", qRefSequenceApplicationApplication); err != nil {
				return err
			}
		}
	}

	if o.RefSequenceApplicationProject != nil {

		// query param ref.sequence.application.project
		var qrRefSequenceApplicationProject string

		if o.RefSequenceApplicationProject != nil {
			qrRefSequenceApplicationProject = *o.RefSequenceApplicationProject
		}
		qRefSequenceApplicationProject := qrRefSequenceApplicationProject
		if qRefSequenceApplicationProject != "" {

			if err := r.SetQueryParam("ref.sequence.application.project", qRefSequenceApplicationProject); err != nil {
				return err
			}
		}
	}

	if o.RefSequenceNumber != nil {

		// query param ref.sequence.number
		var qrRefSequenceNumber string

		if o.RefSequenceNumber != nil {
			qrRefSequenceNumber = *o.RefSequenceNumber
		}
		qRefSequenceNumber := qrRefSequenceNumber
		if qRefSequenceNumber != "" {

			if err := r.SetQueryParam("ref.sequence.number", qRefSequenceNumber); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}