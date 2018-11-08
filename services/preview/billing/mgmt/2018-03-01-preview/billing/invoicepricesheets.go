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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InvoicePricesheetsClient is the billing client provides access to billing resources for Azure subscriptions.
type InvoicePricesheetsClient struct {
	BaseClient
}

// NewInvoicePricesheetsClient creates an instance of the InvoicePricesheetsClient client.
func NewInvoicePricesheetsClient(subscriptionID string) InvoicePricesheetsClient {
	return NewInvoicePricesheetsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoicePricesheetsClientWithBaseURI creates an instance of the InvoicePricesheetsClient client.
func NewInvoicePricesheetsClientWithBaseURI(baseURI string, subscriptionID string) InvoicePricesheetsClient {
	return InvoicePricesheetsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Download get pricesheet data for invoice id (invoiceName).
// Parameters:
// billingAccountID - azure Billing Account ID.
// invoiceName - the name of an invoice resource.
func (client InvoicePricesheetsClient) Download(ctx context.Context, billingAccountID string, invoiceName string) (result InvoicePricesheetsDownloadFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoicePricesheetsClient.Download")
		defer func() {
			sc := -1
			if result.Response() != nil {
				sc = result.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DownloadPreparer(ctx, billingAccountID, invoiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetsClient", "Download", nil, "Failure preparing request")
		return
	}

	result, err = client.DownloadSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoicePricesheetsClient", "Download", result.Response(), "Failure sending request")
		return
	}

	return
}

// DownloadPreparer prepares the Download request.
func (client InvoicePricesheetsClient) DownloadPreparer(ctx context.Context, billingAccountID string, invoiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
		"invoiceName":      autorest.Encode("path", invoiceName),
	}

	const APIVersion = "2018-03-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/invoices/{invoiceName}/pricesheets/default/download", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DownloadSender sends the Download request. The method will close the
// http.Response Body if it receives an error.
func (client InvoicePricesheetsClient) DownloadSender(req *http.Request) (future InvoicePricesheetsDownloadFuture, err error) {
	var resp *http.Response
	resp, err = autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if err != nil {
		return
	}
	future.Future, err = azure.NewFutureFromResponse(resp)
	return
}

// DownloadResponder handles the response to the Download request. The method always
// closes the http.Response Body.
func (client InvoicePricesheetsClient) DownloadResponder(resp *http.Response) (result DownloadURL, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
