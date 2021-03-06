package purview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// DefaultAccountsClient is the creates a Microsoft.Purview management client.
type DefaultAccountsClient struct {
	BaseClient
}

// NewDefaultAccountsClient creates an instance of the DefaultAccountsClient client.
func NewDefaultAccountsClient(subscriptionID string) DefaultAccountsClient {
	return NewDefaultAccountsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDefaultAccountsClientWithBaseURI creates an instance of the DefaultAccountsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDefaultAccountsClientWithBaseURI(baseURI string, subscriptionID string) DefaultAccountsClient {
	return DefaultAccountsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the default account for the scope.
// Parameters:
// scopeTenantID - the tenant ID.
// scopeType - the scope for the default account.
// scope - the Id of the scope object, for example if the scope is "Subscription" then it is the ID of that
// subscription.
func (client DefaultAccountsClient) Get(ctx context.Context, scopeTenantID uuid.UUID, scopeType ScopeType, scope string) (result DefaultAccountPayload, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefaultAccountsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scopeTenantID, scopeType, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DefaultAccountsClient) GetPreparer(ctx context.Context, scopeTenantID uuid.UUID, scopeType ScopeType, scope string) (*http.Request, error) {
	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version":   APIVersion,
		"scopeTenantId": autorest.Encode("query", scopeTenantID),
		"scopeType":     autorest.Encode("query", scopeType),
	}
	if len(scope) > 0 {
		queryParameters["scope"] = autorest.Encode("query", scope)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Purview/getDefaultAccount"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefaultAccountsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefaultAccountsClient) GetResponder(resp *http.Response) (result DefaultAccountPayload, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove removes the default account from the scope.
// Parameters:
// scopeTenantID - the tenant ID.
// scopeType - the scope for the default account.
// scope - the Id of the scope object, for example if the scope is "Subscription" then it is the ID of that
// subscription.
func (client DefaultAccountsClient) Remove(ctx context.Context, scopeTenantID uuid.UUID, scopeType ScopeType, scope string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefaultAccountsClient.Remove")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RemovePreparer(ctx, scopeTenantID, scopeType, scope)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Remove", nil, "Failure preparing request")
		return
	}

	resp, err := client.RemoveSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Remove", resp, "Failure sending request")
		return
	}

	result, err = client.RemoveResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Remove", resp, "Failure responding to request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client DefaultAccountsClient) RemovePreparer(ctx context.Context, scopeTenantID uuid.UUID, scopeType ScopeType, scope string) (*http.Request, error) {
	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version":   APIVersion,
		"scopeTenantId": autorest.Encode("query", scopeTenantID),
		"scopeType":     autorest.Encode("query", scopeType),
	}
	if len(scope) > 0 {
		queryParameters["scope"] = autorest.Encode("query", scope)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Purview/removeDefaultAccount"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client DefaultAccountsClient) RemoveSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client DefaultAccountsClient) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Set sets the default account for the scope.
// Parameters:
// defaultAccountPayload - the payload containing the default account information and the scope.
func (client DefaultAccountsClient) Set(ctx context.Context, defaultAccountPayload DefaultAccountPayload) (result DefaultAccountPayload, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefaultAccountsClient.Set")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.SetPreparer(ctx, defaultAccountPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Set", nil, "Failure preparing request")
		return
	}

	resp, err := client.SetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Set", resp, "Failure sending request")
		return
	}

	result, err = client.SetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "purview.DefaultAccountsClient", "Set", resp, "Failure responding to request")
		return
	}

	return
}

// SetPreparer prepares the Set request.
func (client DefaultAccountsClient) SetPreparer(ctx context.Context, defaultAccountPayload DefaultAccountPayload) (*http.Request, error) {
	const APIVersion = "2021-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/providers/Microsoft.Purview/setDefaultAccount"),
		autorest.WithJSON(defaultAccountPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// SetSender sends the Set request. The method will close the
// http.Response Body if it receives an error.
func (client DefaultAccountsClient) SetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// SetResponder handles the response to the Set request. The method always
// closes the http.Response Body.
func (client DefaultAccountsClient) SetResponder(resp *http.Response) (result DefaultAccountPayload, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
