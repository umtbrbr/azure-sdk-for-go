package billing

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/billing/mgmt/2018-03-01-preview/billing"

// DownloadURL a secure URL that can be used to download a PDF invoice until the URL expires.
type DownloadURL struct {
	// ExpiryTime - The time in UTC at which this download URL will expire.
	ExpiryTime *date.Time `json:"expiryTime,omitempty"`
	// URL - The URL to the PDF file.
	URL *string `json:"url,omitempty"`
}

// EnrollmentAccount an enrollment account resource.
type EnrollmentAccount struct {
	autorest.Response `json:"-"`
	// EnrollmentAccountProperties - An enrollment account.
	*EnrollmentAccountProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for EnrollmentAccount.
func (ea EnrollmentAccount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ea.EnrollmentAccountProperties != nil {
		objectMap["properties"] = ea.EnrollmentAccountProperties
	}
	if ea.ID != nil {
		objectMap["id"] = ea.ID
	}
	if ea.Name != nil {
		objectMap["name"] = ea.Name
	}
	if ea.Type != nil {
		objectMap["type"] = ea.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for EnrollmentAccount struct.
func (ea *EnrollmentAccount) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var enrollmentAccountProperties EnrollmentAccountProperties
				err = json.Unmarshal(*v, &enrollmentAccountProperties)
				if err != nil {
					return err
				}
				ea.EnrollmentAccountProperties = &enrollmentAccountProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ea.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ea.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ea.Type = &typeVar
			}
		}
	}

	return nil
}

// EnrollmentAccountListResult result of listing enrollment accounts.
type EnrollmentAccountListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of enrollment accounts.
	Value *[]EnrollmentAccount `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// EnrollmentAccountListResultIterator provides access to a complete listing of EnrollmentAccount values.
type EnrollmentAccountListResultIterator struct {
	i    int
	page EnrollmentAccountListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *EnrollmentAccountListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentAccountListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *EnrollmentAccountListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter EnrollmentAccountListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter EnrollmentAccountListResultIterator) Response() EnrollmentAccountListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter EnrollmentAccountListResultIterator) Value() EnrollmentAccount {
	if !iter.page.NotDone() {
		return EnrollmentAccount{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ealr EnrollmentAccountListResult) IsEmpty() bool {
	return ealr.Value == nil || len(*ealr.Value) == 0
}

// enrollmentAccountListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ealr EnrollmentAccountListResult) enrollmentAccountListResultPreparer(ctx context.Context) (*http.Request, error) {
	if ealr.NextLink == nil || len(to.String(ealr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ealr.NextLink)))
}

// EnrollmentAccountListResultPage contains a page of EnrollmentAccount values.
type EnrollmentAccountListResultPage struct {
	fn   func(context.Context, EnrollmentAccountListResult) (EnrollmentAccountListResult, error)
	ealr EnrollmentAccountListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *EnrollmentAccountListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/EnrollmentAccountListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ealr)
	if err != nil {
		return err
	}
	page.ealr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *EnrollmentAccountListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page EnrollmentAccountListResultPage) NotDone() bool {
	return !page.ealr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page EnrollmentAccountListResultPage) Response() EnrollmentAccountListResult {
	return page.ealr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page EnrollmentAccountListResultPage) Values() []EnrollmentAccount {
	if page.ealr.IsEmpty() {
		return nil
	}
	return *page.ealr.Value
}

// EnrollmentAccountProperties the properties of the enrollment account.
type EnrollmentAccountProperties struct {
	// PrincipalName - The account owner's principal name.
	PrincipalName *string `json:"principalName,omitempty"`
}

// ErrorDetails the details of the error.
type ErrorDetails struct {
	// Code - Error code.
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
	// Target - The target of the particular error.
	Target *string `json:"target,omitempty"`
}

// ErrorResponse error response indicates that the service is not able to process the incoming request. The
// reason is provided in the error message.
type ErrorResponse struct {
	// Error - The details of the error.
	Error *ErrorDetails `json:"error,omitempty"`
}

// Invoice an invoice resource can be used download a PDF version of an invoice.
type Invoice struct {
	autorest.Response `json:"-"`
	// InvoiceProperties - An invoice.
	*InvoiceProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Invoice.
func (i Invoice) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if i.InvoiceProperties != nil {
		objectMap["properties"] = i.InvoiceProperties
	}
	if i.ID != nil {
		objectMap["id"] = i.ID
	}
	if i.Name != nil {
		objectMap["name"] = i.Name
	}
	if i.Type != nil {
		objectMap["type"] = i.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Invoice struct.
func (i *Invoice) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var invoiceProperties InvoiceProperties
				err = json.Unmarshal(*v, &invoiceProperties)
				if err != nil {
					return err
				}
				i.InvoiceProperties = &invoiceProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				i.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				i.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				i.Type = &typeVar
			}
		}
	}

	return nil
}

// InvoicePricesheetGetFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type InvoicePricesheetGetFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *InvoicePricesheetGetFuture) Result(client InvoicePricesheetClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetGetFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("billing.InvoicePricesheetGetFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// InvoiceProperties the properties of the invoice.
type InvoiceProperties struct {
	// DownloadURL - A secure link to download the PDF version of an invoice. The link will cease to work after its expiry time is reached.
	DownloadURL *DownloadURL `json:"downloadUrl,omitempty"`
	// InvoicePeriodStartDate - The start of the date range covered by the invoice.
	InvoicePeriodStartDate *date.Date `json:"invoicePeriodStartDate,omitempty"`
	// InvoicePeriodEndDate - The end of the date range covered by the invoice.
	InvoicePeriodEndDate *date.Date `json:"invoicePeriodEndDate,omitempty"`
	// BillingPeriodIds - Array of billing perdiod ids that the invoice is attributed to.
	BillingPeriodIds *[]string `json:"billingPeriodIds,omitempty"`
}

// InvoicesListResult result of listing invoices. It contains a list of available invoices in reverse
// chronological order.
type InvoicesListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of invoices.
	Value *[]Invoice `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// InvoicesListResultIterator provides access to a complete listing of Invoice values.
type InvoicesListResultIterator struct {
	i    int
	page InvoicesListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *InvoicesListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *InvoicesListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter InvoicesListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter InvoicesListResultIterator) Response() InvoicesListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter InvoicesListResultIterator) Value() Invoice {
	if !iter.page.NotDone() {
		return Invoice{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (ilr InvoicesListResult) IsEmpty() bool {
	return ilr.Value == nil || len(*ilr.Value) == 0
}

// invoicesListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (ilr InvoicesListResult) invoicesListResultPreparer(ctx context.Context) (*http.Request, error) {
	if ilr.NextLink == nil || len(to.String(ilr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(ilr.NextLink)))
}

// InvoicesListResultPage contains a page of Invoice values.
type InvoicesListResultPage struct {
	fn  func(context.Context, InvoicesListResult) (InvoicesListResult, error)
	ilr InvoicesListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *InvoicesListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicesListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.ilr)
	if err != nil {
		return err
	}
	page.ilr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *InvoicesListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page InvoicesListResultPage) NotDone() bool {
	return !page.ilr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page InvoicesListResultPage) Response() InvoicesListResult {
	return page.ilr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page InvoicesListResultPage) Values() []Invoice {
	if page.ilr.IsEmpty() {
		return nil
	}
	return *page.ilr.Value
}

// Operation a Billing REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.Billing.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: Invoice, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: Read, write, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result listing billing operations. It contains a list of operations and a URL link
// to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of billing operations supported by the Microsoft.Billing resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Period a billing period resource.
type Period struct {
	autorest.Response `json:"-"`
	// PeriodProperties - A billing period.
	*PeriodProperties `json:"properties,omitempty"`
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Period.
func (p Period) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if p.PeriodProperties != nil {
		objectMap["properties"] = p.PeriodProperties
	}
	if p.ID != nil {
		objectMap["id"] = p.ID
	}
	if p.Name != nil {
		objectMap["name"] = p.Name
	}
	if p.Type != nil {
		objectMap["type"] = p.Type
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Period struct.
func (p *Period) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var periodProperties PeriodProperties
				err = json.Unmarshal(*v, &periodProperties)
				if err != nil {
					return err
				}
				p.PeriodProperties = &periodProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				p.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				p.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				p.Type = &typeVar
			}
		}
	}

	return nil
}

// PeriodProperties the properties of the billing period.
type PeriodProperties struct {
	// BillingPeriodStartDate - The start of the date range covered by the billing period.
	BillingPeriodStartDate *date.Date `json:"billingPeriodStartDate,omitempty"`
	// BillingPeriodEndDate - The end of the date range covered by the billing period.
	BillingPeriodEndDate *date.Date `json:"billingPeriodEndDate,omitempty"`
	// InvoiceIds - Array of invoice ids that associated with.
	InvoiceIds *[]string `json:"invoiceIds,omitempty"`
}

// PeriodsListResult result of listing billing periods. It contains a list of available billing periods in
// reverse chronological order.
type PeriodsListResult struct {
	autorest.Response `json:"-"`
	// Value - The list of billing periods.
	Value *[]Period `json:"value,omitempty"`
	// NextLink - The link (url) to the next page of results.
	NextLink *string `json:"nextLink,omitempty"`
}

// PeriodsListResultIterator provides access to a complete listing of Period values.
type PeriodsListResultIterator struct {
	i    int
	page PeriodsListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *PeriodsListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeriodsListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *PeriodsListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter PeriodsListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter PeriodsListResultIterator) Response() PeriodsListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter PeriodsListResultIterator) Value() Period {
	if !iter.page.NotDone() {
		return Period{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (plr PeriodsListResult) IsEmpty() bool {
	return plr.Value == nil || len(*plr.Value) == 0
}

// periodsListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (plr PeriodsListResult) periodsListResultPreparer(ctx context.Context) (*http.Request, error) {
	if plr.NextLink == nil || len(to.String(plr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(plr.NextLink)))
}

// PeriodsListResultPage contains a page of Period values.
type PeriodsListResultPage struct {
	fn  func(context.Context, PeriodsListResult) (PeriodsListResult, error)
	plr PeriodsListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *PeriodsListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PeriodsListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.plr)
	if err != nil {
		return err
	}
	page.plr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *PeriodsListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page PeriodsListResultPage) NotDone() bool {
	return !page.plr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page PeriodsListResultPage) Response() PeriodsListResult {
	return page.plr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page PeriodsListResultPage) Values() []Period {
	if page.plr.IsEmpty() {
		return nil
	}
	return *page.plr.Value
}

// Resource the Resource model definition.
type Resource struct {
	// ID - Resource Id.
	ID *string `json:"id,omitempty"`
	// Name - Resource name.
	Name *string `json:"name,omitempty"`
	// Type - Resource type.
	Type *string `json:"type,omitempty"`
}
