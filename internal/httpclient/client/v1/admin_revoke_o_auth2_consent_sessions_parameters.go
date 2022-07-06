// Code generated by go-swagger; DO NOT EDIT.

package v1

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

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// NewAdminRevokeOAuth2ConsentSessionsParams creates a new AdminRevokeOAuth2ConsentSessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminRevokeOAuth2ConsentSessionsParams() *AdminRevokeOAuth2ConsentSessionsParams {
	return &AdminRevokeOAuth2ConsentSessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminRevokeOAuth2ConsentSessionsParamsWithTimeout creates a new AdminRevokeOAuth2ConsentSessionsParams object
// with the ability to set a timeout on a request.
func NewAdminRevokeOAuth2ConsentSessionsParamsWithTimeout(timeout time.Duration) *AdminRevokeOAuth2ConsentSessionsParams {
	return &AdminRevokeOAuth2ConsentSessionsParams{
		timeout: timeout,
	}
}

// NewAdminRevokeOAuth2ConsentSessionsParamsWithContext creates a new AdminRevokeOAuth2ConsentSessionsParams object
// with the ability to set a context for a request.
func NewAdminRevokeOAuth2ConsentSessionsParamsWithContext(ctx context.Context) *AdminRevokeOAuth2ConsentSessionsParams {
	return &AdminRevokeOAuth2ConsentSessionsParams{
		Context: ctx,
	}
}

// NewAdminRevokeOAuth2ConsentSessionsParamsWithHTTPClient creates a new AdminRevokeOAuth2ConsentSessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminRevokeOAuth2ConsentSessionsParamsWithHTTPClient(client *http.Client) *AdminRevokeOAuth2ConsentSessionsParams {
	return &AdminRevokeOAuth2ConsentSessionsParams{
		HTTPClient: client,
	}
}

/* AdminRevokeOAuth2ConsentSessionsParams contains all the parameters to send to the API endpoint
   for the admin revoke o auth2 consent sessions operation.

   Typically these are written to a http.Request.
*/
type AdminRevokeOAuth2ConsentSessionsParams struct {

	/* All.

	   If set to `true` deletes all consent sessions by the Subject that have been granted.
	*/
	All *bool

	/* Client.

	   If set, deletes only those consent sessions by the Subject that have been granted to the specified OAuth 2.0 Client ID
	*/
	Client *string

	/* Subject.

	   The subject (Subject) whose consent sessions should be deleted.
	*/
	Subject string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin revoke o auth2 consent sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithDefaults() *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin revoke o auth2 consent sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithTimeout(timeout time.Duration) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithContext(ctx context.Context) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithHTTPClient(client *http.Client) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAll adds the all to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithAll(all *bool) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetAll(all)
	return o
}

// SetAll adds the all to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetAll(all *bool) {
	o.All = all
}

// WithClient adds the client to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithClient(client *string) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetClient(client)
	return o
}

// SetClient adds the client to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetClient(client *string) {
	o.Client = client
}

// WithSubject adds the subject to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) WithSubject(subject string) *AdminRevokeOAuth2ConsentSessionsParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the admin revoke o auth2 consent sessions params
func (o *AdminRevokeOAuth2ConsentSessionsParams) SetSubject(subject string) {
	o.Subject = subject
}

// WriteToRequest writes these params to a swagger request
func (o *AdminRevokeOAuth2ConsentSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.All != nil {

		// query param all
		var qrAll bool

		if o.All != nil {
			qrAll = *o.All
		}
		qAll := swag.FormatBool(qrAll)
		if qAll != "" {

			if err := r.SetQueryParam("all", qAll); err != nil {
				return err
			}
		}
	}

	if o.Client != nil {

		// query param client
		var qrClient string

		if o.Client != nil {
			qrClient = *o.Client
		}
		qClient := qrClient
		if qClient != "" {

			if err := r.SetQueryParam("client", qClient); err != nil {
				return err
			}
		}
	}

	// query param subject
	qrSubject := o.Subject
	qSubject := qrSubject
	if qSubject != "" {

		if err := r.SetQueryParam("subject", qSubject); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}