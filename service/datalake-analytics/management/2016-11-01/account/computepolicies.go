package account

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
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ComputePoliciesClient is the creates an Azure Data Lake Analytics account
// management client.
type ComputePoliciesClient struct {
	ManagementClient
}

// NewComputePoliciesClient creates an instance of the ComputePoliciesClient
// client.
func NewComputePoliciesClient(subscriptionID string) ComputePoliciesClient {
	return NewComputePoliciesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewComputePoliciesClientWithBaseURI creates an instance of the
// ComputePoliciesClient client.
func NewComputePoliciesClientWithBaseURI(baseURI string, subscriptionID string) ComputePoliciesClient {
	return ComputePoliciesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates the specified compute policy. During
// update, the compute policy with the specified name will be replaced with
// this new compute policy. An account supports, at most, 50 policies
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to add or replace the compute policy. computePolicyName is
// the name of the compute policy to create or update. parameters is parameters
// supplied to create or update the compute policy. The max degree of
// parallelism per job property, min priority per job property, or both must be
// present.
func (client ComputePoliciesClient) CreateOrUpdate(resourceGroupName string, accountName string, computePolicyName string, parameters ComputePolicyCreateOrUpdateParameters) (result ComputePolicy, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.ComputePolicyPropertiesCreateParameters", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "parameters.ComputePolicyPropertiesCreateParameters.ObjectID", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.ComputePolicyPropertiesCreateParameters.MaxDegreeOfParallelismPerJob", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ComputePolicyPropertiesCreateParameters.MaxDegreeOfParallelismPerJob", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
					{Target: "parameters.ComputePolicyPropertiesCreateParameters.MinPriorityPerJob", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.ComputePolicyPropertiesCreateParameters.MinPriorityPerJob", Name: validation.InclusiveMinimum, Rule: 1, Chain: nil}}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "account.ComputePoliciesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, accountName, computePolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ComputePoliciesClient) CreateOrUpdatePreparer(resourceGroupName string, accountName string, computePolicyName string, parameters ComputePolicyCreateOrUpdateParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"computePolicyName": autorest.Encode("path", computePolicyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ComputePoliciesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ComputePoliciesClient) CreateOrUpdateResponder(resp *http.Response) (result ComputePolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified compute policy from the specified Data Lake
// Analytics account
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account from which to delete the compute policy. computePolicyName
// is the name of the compute policy to delete.
func (client ComputePoliciesClient) Delete(resourceGroupName string, accountName string, computePolicyName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, accountName, computePolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ComputePoliciesClient) DeletePreparer(resourceGroupName string, accountName string, computePolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"computePolicyName": autorest.Encode("path", computePolicyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ComputePoliciesClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ComputePoliciesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the specified Data Lake Analytics compute policy.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account from which to get the compute policy. computePolicyName is
// the name of the compute policy to retrieve.
func (client ComputePoliciesClient) Get(resourceGroupName string, accountName string, computePolicyName string) (result ComputePolicy, err error) {
	req, err := client.GetPreparer(resourceGroupName, accountName, computePolicyName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ComputePoliciesClient) GetPreparer(resourceGroupName string, accountName string, computePolicyName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"computePolicyName": autorest.Encode("path", computePolicyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ComputePoliciesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ComputePoliciesClient) GetResponder(resp *http.Response) (result ComputePolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccount lists the Data Lake Analytics compute policies within the
// specified Data Lake Analytics account. An account supports, at most, 50
// policies
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account from which to get the compute policies.
func (client ComputePoliciesClient) ListByAccount(resourceGroupName string, accountName string) (result ComputePolicyListResult, err error) {
	req, err := client.ListByAccountPreparer(resourceGroupName, accountName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", resp, "Failure sending request")
		return
	}

	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", resp, "Failure responding to request")
	}

	return
}

// ListByAccountPreparer prepares the ListByAccount request.
func (client ComputePoliciesClient) ListByAccountPreparer(resourceGroupName string, accountName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByAccountSender sends the ListByAccount request. The method will close the
// http.Response Body if it receives an error.
func (client ComputePoliciesClient) ListByAccountSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByAccountResponder handles the response to the ListByAccount request. The method always
// closes the http.Response Body.
func (client ComputePoliciesClient) ListByAccountResponder(resp *http.Response) (result ComputePolicyListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByAccountNextResults retrieves the next set of results, if any.
func (client ComputePoliciesClient) ListByAccountNextResults(lastResults ComputePolicyListResult) (result ComputePolicyListResult, err error) {
	req, err := lastResults.ComputePolicyListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByAccountSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", resp, "Failure sending next results request")
	}

	result, err = client.ListByAccountResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "ListByAccount", resp, "Failure responding to next results request")
	}

	return
}

// ListByAccountComplete gets all elements from the list without paging.
func (client ComputePoliciesClient) ListByAccountComplete(resourceGroupName string, accountName string, cancel <-chan struct{}) (<-chan ComputePolicy, <-chan error) {
	resultChan := make(chan ComputePolicy)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByAccount(resourceGroupName, accountName)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListByAccountNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}

// Update updates the specified compute policy.
//
// resourceGroupName is the name of the Azure resource group that contains the
// Data Lake Analytics account. accountName is the name of the Data Lake
// Analytics account to which to update the compute policy. computePolicyName
// is the name of the compute policy to update. parameters is parameters
// supplied to update the compute policy.
func (client ComputePoliciesClient) Update(resourceGroupName string, accountName string, computePolicyName string, parameters *ComputePolicy) (result ComputePolicy, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, accountName, computePolicyName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "account.ComputePoliciesClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ComputePoliciesClient) UpdatePreparer(resourceGroupName string, accountName string, computePolicyName string, parameters *ComputePolicy) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"computePolicyName": autorest.Encode("path", computePolicyName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-11-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataLakeAnalytics/accounts/{accountName}/computePolicies/{computePolicyName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if parameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(parameters))
	}
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ComputePoliciesClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ComputePoliciesClient) UpdateResponder(resp *http.Response) (result ComputePolicy, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}