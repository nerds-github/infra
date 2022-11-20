// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/devbookhq/packages/firecracker-task-driver/internal/client/models"
)

// NewPatchBalloonStatsIntervalParams creates a new PatchBalloonStatsIntervalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchBalloonStatsIntervalParams() *PatchBalloonStatsIntervalParams {
	return &PatchBalloonStatsIntervalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchBalloonStatsIntervalParamsWithTimeout creates a new PatchBalloonStatsIntervalParams object
// with the ability to set a timeout on a request.
func NewPatchBalloonStatsIntervalParamsWithTimeout(timeout time.Duration) *PatchBalloonStatsIntervalParams {
	return &PatchBalloonStatsIntervalParams{
		timeout: timeout,
	}
}

// NewPatchBalloonStatsIntervalParamsWithContext creates a new PatchBalloonStatsIntervalParams object
// with the ability to set a context for a request.
func NewPatchBalloonStatsIntervalParamsWithContext(ctx context.Context) *PatchBalloonStatsIntervalParams {
	return &PatchBalloonStatsIntervalParams{
		Context: ctx,
	}
}

// NewPatchBalloonStatsIntervalParamsWithHTTPClient creates a new PatchBalloonStatsIntervalParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchBalloonStatsIntervalParamsWithHTTPClient(client *http.Client) *PatchBalloonStatsIntervalParams {
	return &PatchBalloonStatsIntervalParams{
		HTTPClient: client,
	}
}

/* PatchBalloonStatsIntervalParams contains all the parameters to send to the API endpoint
   for the patch balloon stats interval operation.

   Typically these are written to a http.Request.
*/
type PatchBalloonStatsIntervalParams struct {

	/* Body.

	   Balloon properties
	*/
	Body *models.BalloonStatsUpdate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch balloon stats interval params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchBalloonStatsIntervalParams) WithDefaults() *PatchBalloonStatsIntervalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch balloon stats interval params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchBalloonStatsIntervalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithTimeout(timeout time.Duration) *PatchBalloonStatsIntervalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithContext(ctx context.Context) *PatchBalloonStatsIntervalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithHTTPClient(client *http.Client) *PatchBalloonStatsIntervalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) WithBody(body *models.BalloonStatsUpdate) *PatchBalloonStatsIntervalParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch balloon stats interval params
func (o *PatchBalloonStatsIntervalParams) SetBody(body *models.BalloonStatsUpdate) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchBalloonStatsIntervalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
