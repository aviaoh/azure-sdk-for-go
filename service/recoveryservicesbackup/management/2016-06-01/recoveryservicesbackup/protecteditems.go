package recoveryservicesbackup

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
	"net/http"
)

// ProtectedItemsClient is the client for the ProtectedItems methods of the
// Recoveryservicesbackup service.
type ProtectedItemsClient struct {
	ManagementClient
}

// NewProtectedItemsClient creates an instance of the ProtectedItemsClient
// client.
func NewProtectedItemsClient(subscriptionID string) ProtectedItemsClient {
	return NewProtectedItemsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProtectedItemsClientWithBaseURI creates an instance of the
// ProtectedItemsClient client.
func NewProtectedItemsClientWithBaseURI(baseURI string, subscriptionID string) ProtectedItemsClient {
	return ProtectedItemsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate this operation enables an item to be backed up, or modifies
// the existing backup policy information for an item that has been backed up.
// This is an asynchronous operation. To learn the status of the operation,
// call the GetItemOperationResult API.
//
// vaultName is the name of the Recovery Services vault. resourceGroupName is
// the name of the resource group associated with the Recovery Services vault.
// fabricName is the fabric name associated with the backup item. containerName
// is the container name associated with the backup item. protectedItemName is
// the name of the backup item. resourceProtectedItem is the resource backup
// item.
func (client ProtectedItemsClient) CreateOrUpdate(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceProtectedItem ProtectedItemResource) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdatePreparer(vaultName, resourceGroupName, fabricName, containerName, protectedItemName, resourceProtectedItem)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ProtectedItemsClient) CreateOrUpdatePreparer(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, resourceProtectedItem ProtectedItemResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", pathParameters),
		autorest.WithJSON(resourceProtectedItem),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ProtectedItemsClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete used to disable the backup job for an item within a container. This
// is an asynchronous operation. To learn the status of the request, call the
// GetItemOperationResult API.
//
// vaultName is the name of the Recovery Services vault. resourceGroupName is
// the name of the resource group associated with the Recovery Services vault.
// fabricName is the fabric name associated with the backup item. containerName
// is the container name associated with the backup item. protectedItemName is
// the backup item to be deleted.
func (client ProtectedItemsClient) Delete(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(vaultName, resourceGroupName, fabricName, containerName, protectedItemName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ProtectedItemsClient) DeletePreparer(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ProtectedItemsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get provides the details of the backup item. This is an asynchronous
// operation. To know the status of the operation, call the
// GetItemOperationResult API.
//
// vaultName is the name of the Recovery Services vault. resourceGroupName is
// the name of the resource group associated with the Recovery Services vault.
// fabricName is the fabric name associated with the backup item. containerName
// is the container name associated with the backup item. protectedItemName is
// the backup item name used in this GET operation. filter is expand eq
// {extendedinfo}. This filter enables you to choose (or filter) specific items
// in the list of backup items.
func (client ProtectedItemsClient) Get(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (result ProtectedItemResource, err error) {
	req, err := client.GetPreparer(vaultName, resourceGroupName, fabricName, containerName, protectedItemName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProtectedItemsClient) GetPreparer(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProtectedItemsClient) GetResponder(resp *http.Response) (result ProtectedItemResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List provides a pageable list of all items in a subscription, that can be
// protected.
//
// vaultName is the name of the Recovery Services vault. resourceGroupName is
// the name of the resource group associated with the Recovery Services vault.
// filter is itemType eq { VM , FileFolder , AzureSqlDb , SQLDB , Exchange ,
// Sharepoint , DPMUnknown } and providerType eq { AzureIaasVM, MAB, DPM,
// AzureBackupServer, AzureSql } and policyName eq {policyname} and
// containerName eq {containername} and backupManagementType eq { AzureIaasVM,
// MAB, DPM, AzureBackupServer, AzureSql }. skipToken is the Skip Token filter.
func (client ProtectedItemsClient) List(vaultName string, resourceGroupName string, filter string, skipToken string) (result ProtectedItemResourceList, err error) {
	req, err := client.ListPreparer(vaultName, resourceGroupName, filter, skipToken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ProtectedItemsClient) ListPreparer(vaultName string, resourceGroupName string, filter string, skipToken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-06-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(skipToken) > 0 {
		queryParameters["$skipToken"] = autorest.Encode("query", skipToken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupProtectedItems", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ProtectedItemsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ProtectedItemsClient) ListResponder(resp *http.Response) (result ProtectedItemResourceList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ProtectedItemsClient) ListNextResults(lastResults ProtectedItemResourceList) (result ProtectedItemResourceList, err error) {
	req, err := lastResults.ProtectedItemResourceListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recoveryservicesbackup.ProtectedItemsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ProtectedItemsClient) ListComplete(vaultName string, resourceGroupName string, filter string, skipToken string, cancel <-chan struct{}) (<-chan ProtectedItemResource, <-chan error) {
	resultChan := make(chan ProtectedItemResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(vaultName, resourceGroupName, filter, skipToken)
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
			list, err = client.ListNextResults(list)
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