package hdinsight

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
// Code generated by Microsoft (R) AutoRest Code Generator 2.2.22.0
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ApplicationName enumerates the values for application name.
type ApplicationName string

// AsyncOperationState enumerates the values for async operation state.
type AsyncOperationState string

const (
	// Failed specifies the failed state for async operation state.
	Failed AsyncOperationState = "Failed"
	// InProgress specifies the in progress state for async operation state.
	InProgress AsyncOperationState = "InProgress"
	// Succeeded specifies the succeeded state for async operation state.
	Succeeded AsyncOperationState = "Succeeded"
)

// ClusterProvisioningState enumerates the values for cluster provisioning state.
type ClusterProvisioningState string

const (
	// ClusterProvisioningStateCanceled specifies the cluster provisioning state canceled state for cluster provisioning
	// state.
	ClusterProvisioningStateCanceled ClusterProvisioningState = "Canceled"
	// ClusterProvisioningStateDeleting specifies the cluster provisioning state deleting state for cluster provisioning
	// state.
	ClusterProvisioningStateDeleting ClusterProvisioningState = "Deleting"
	// ClusterProvisioningStateFailed specifies the cluster provisioning state failed state for cluster provisioning state.
	ClusterProvisioningStateFailed ClusterProvisioningState = "Failed"
	// ClusterProvisioningStateInProgress specifies the cluster provisioning state in progress state for cluster
	// provisioning state.
	ClusterProvisioningStateInProgress ClusterProvisioningState = "InProgress"
	// ClusterProvisioningStateSucceeded specifies the cluster provisioning state succeeded state for cluster provisioning
	// state.
	ClusterProvisioningStateSucceeded ClusterProvisioningState = "Succeeded"
)

// Configurationname enumerates the values for configurationname.
type Configurationname string

const (
	// CoreSite specifies the core site state for configurationname.
	CoreSite Configurationname = "core-site"
	// Gateway specifies the gateway state for configurationname.
	Gateway Configurationname = "gateway"
)

// DirectoryType enumerates the values for directory type.
type DirectoryType string

const (
	// ActiveDirectory specifies the active directory state for directory type.
	ActiveDirectory DirectoryType = "ActiveDirectory"
)

// OSType enumerates the values for os type.
type OSType string

const (
	// Linux specifies the linux state for os type.
	Linux OSType = "Linux"
	// Windows specifies the windows state for os type.
	Windows OSType = "Windows"
)

// Tier enumerates the values for tier.
type Tier string

const (
	// Premium specifies the premium state for tier.
	Premium Tier = "Premium"
	// Standard specifies the standard state for tier.
	Standard Tier = "Standard"
)

// Application is hDInsight cluster application
type Application struct {
	autorest.Response `json:"-"`
	ID                *SubResource              `json:"id,omitempty"`
	Name              *string                   `json:"name,omitempty"`
	Type              *string                   `json:"type,omitempty"`
	Etag              *string                   `json:"etag,omitempty"`
	Tags              *map[string]*string       `json:"tags,omitempty"`
	Properties        *ApplicationGetProperties `json:"properties,omitempty"`
}

// ApplicationGetEndpoint is gets Application ssh endpoint
type ApplicationGetEndpoint struct {
	Location        *string `json:"location,omitempty"`
	DestinationPort *int32  `json:"destinationPort,omitempty"`
	PublicPort      *int32  `json:"publicPort,omitempty"`
}

// ApplicationGetHTTPSEndpoint is gets application Http endpoints.
type ApplicationGetHTTPSEndpoint struct {
	AdditionalProperties *map[string]*string `json:",omitempty"`
	AccessModes          *[]string           `json:"accessModes,omitempty"`
	Location             *string             `json:"location,omitempty"`
	DestinationPort      *int32              `json:"destinationPort,omitempty"`
	PublicPort           *int32              `json:"publicPort,omitempty"`
}

// ApplicationGetProperties is hDInsight cluster application.
type ApplicationGetProperties struct {
	ComputeProfile         *ComputeProfile                `json:"computeProfile,omitempty"`
	InstallScriptActions   *[]RuntimeScriptAction         `json:"installScriptActions,omitempty"`
	UninstallScriptActions *[]RuntimeScriptAction         `json:"uninstallScriptActions,omitempty"`
	HTTPSEndpoints         *[]ApplicationGetHTTPSEndpoint `json:"httpsEndpoints,omitempty"`
	SSHEndpoints           *[]ApplicationGetEndpoint      `json:"sshEndpoints,omitempty"`
	ProvisioningState      *string                        `json:"provisioningState,omitempty"`
	ApplicationType        *string                        `json:"applicationType,omitempty"`
	ApplicationState       *string                        `json:"applicationState,omitempty"`
	Errors                 *[]Errors                      `json:"errors,omitempty"`
	CreatedDate            *string                        `json:"createdDate,omitempty"`
	MarketplaceIdentifier  *string                        `json:"marketplaceIdentifier,omitempty"`
	AdditionalProperties   *string                        `json:"additionalProperties,omitempty"`
}

// ApplicationListResult is result of the request to list cluster Applications. It contains a list of operations and a
// URL link to get the next set of results.
type ApplicationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Application `json:"value,omitempty"`
	NextLink          *string        `json:"nextLink,omitempty"`
}

// ApplicationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ApplicationListResult) ApplicationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// CapabilitiesResult is the Get Capabilities operation response.
type CapabilitiesResult struct {
	autorest.Response `json:"-"`
	Versions          *map[string]*VersionsCapability `json:"versions,omitempty"`
	Regions           *map[string]*RegionsCapability  `json:"regions,omitempty"`
	Vmsizes           *map[string]*VMSizesCapability  `json:"vmsizes,omitempty"`
	VmsizeFilters     *[]VMSizeCompatibilityFilter    `json:"vmsize_filters,omitempty"`
	Features          *[]string                       `json:"features,omitempty"`
	Quota             *QuotaCapability                `json:"quota,omitempty"`
}

// Cluster is describes the cluster.
type Cluster struct {
	autorest.Response `json:"-"`
	ID                *string               `json:"id,omitempty"`
	Name              *string               `json:"name,omitempty"`
	Type              *string               `json:"type,omitempty"`
	Location          *string               `json:"location,omitempty"`
	Tags              *map[string]*string   `json:"tags,omitempty"`
	Etag              *string               `json:"etag,omitempty"`
	Properties        *ClusterGetProperties `json:"properties,omitempty"`
}

// ClusterCreateParametersExtended is the CreateCluster request parameters.
type ClusterCreateParametersExtended struct {
	Location   *string                  `json:"location,omitempty"`
	Tags       *map[string]*string      `json:"tags,omitempty"`
	Properties *ClusterCreateProperties `json:"properties,omitempty"`
}

// ClusterCreateProperties is the cluster create parameters.
type ClusterCreateProperties struct {
	ClusterVersion    *string            `json:"clusterVersion,omitempty"`
	OsType            OSType             `json:"osType,omitempty"`
	Tier              Tier               `json:"tier,omitempty"`
	ClusterDefinition *ClusterDefinition `json:"clusterDefinition,omitempty"`
	SecurityProfile   *SecurityProfile   `json:"securityProfile,omitempty"`
	ComputeProfile    *ComputeProfile    `json:"computeProfile,omitempty"`
	StorageProfile    *StorageProfile    `json:"storageProfile,omitempty"`
}

// ClusterDefinition is the cluste definition.
type ClusterDefinition struct {
	Blueprint        *string                 `json:"blueprint,omitempty"`
	Kind             *string                 `json:"kind,omitempty"`
	ComponentVersion *map[string]*string     `json:"componentVersion,omitempty"`
	Configurations   *map[string]interface{} `json:"configurations,omitempty"`
}

// ClusterGetProperties is the properties of cluster.
type ClusterGetProperties struct {
	ClusterVersion        *string                  `json:"clusterVersion,omitempty"`
	OsType                OSType                   `json:"osType,omitempty"`
	Tier                  Tier                     `json:"tier,omitempty"`
	ClusterDefinition     *ClusterDefinition       `json:"clusterDefinition,omitempty"`
	SecurityProfile       *SecurityProfile         `json:"securityProfile,omitempty"`
	ComputeProfile        *ComputeProfile          `json:"computeProfile,omitempty"`
	ProvisioningState     ClusterProvisioningState `json:"provisioningState,omitempty"`
	CreatedDate           *string                  `json:"createdDate,omitempty"`
	ClusterState          *string                  `json:"clusterState,omitempty"`
	QuotaInfo             *QuotaInfo               `json:"quotaInfo,omitempty"`
	Errors                *[]Errors                `json:"errors,omitempty"`
	ConnectivityEndpoints *[]ConnectivityEndpoint  `json:"connectivityEndpoints,omitempty"`
}

// ClusterListPersistedScriptActionsResult is list PersistedScriptActions operations response.
type ClusterListPersistedScriptActionsResult struct {
	Value    *[]RuntimeScriptAction `json:"value,omitempty"`
	NextLink *string                `json:"nextLink,omitempty"`
}

// ClusterListResult is the List Cluster operation response.
type ClusterListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Cluster `json:"value,omitempty"`
	NextLink          *string    `json:"nextLink,omitempty"`
}

// ClusterListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ClusterListResult) ClusterListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ClusterListRuntimeScriptActionDetailResult is the ListScriptExecutionHistory response.
type ClusterListRuntimeScriptActionDetailResult struct {
	Value    *[]RuntimeScriptActionDetail `json:"value,omitempty"`
	NextLink *string                      `json:"nextLink,omitempty"`
}

// ClusterPatchParameters is the PatchCluster request parameters
type ClusterPatchParameters struct {
	Tags *map[string]*string `json:"tags,omitempty"`
}

// ClusterResizeParameters is the Resize Cluster request parameters.
type ClusterResizeParameters struct {
	TargetInstanceCount *int32 `json:"targetInstanceCount,omitempty"`
}

// ComputeProfile is describes the compute profile.
type ComputeProfile struct {
	Roles *[]Role `json:"roles,omitempty"`
}

// ConnectivityEndpoint is the connectivity properties
type ConnectivityEndpoint struct {
	Name     *string `json:"name,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Location *string `json:"location,omitempty"`
	Port     *int32  `json:"port,omitempty"`
}

// Errors is the error message associated with the cluster creation.
type Errors struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// ExecuteScriptActionParameters is describes the script actions on a running cluster.
type ExecuteScriptActionParameters struct {
	ScriptActions    *[]RuntimeScriptAction `json:"scriptActions,omitempty"`
	PersistOnSuccess *string                `json:"persistOnSuccess,omitempty"`
}

// Extension is cluster monitoring extensions
type Extension struct {
	autorest.Response `json:"-"`
	WorkspaceID       *string `json:"workspaceId,omitempty"`
	PrimaryKey        *string `json:"primaryKey,omitempty"`
}

// HardwareProfile is describes the hardware profile.
type HardwareProfile struct {
	VMSize *string `json:"vmSize,omitempty"`
}

// HTTPConnectivitySettings is the payload for a Configure HTTP settings request.
type HTTPConnectivitySettings struct {
	autorest.Response `json:"-"`
	EnabledCredential *string `json:"restAuthCredential.isEnabled,omitempty"`
	Username          *string `json:"restAuthCredential.username,omitempty"`
	Password          *string `json:"restAuthCredential.password,omitempty"`
}

// HTTPSettingsParameters is the payload for a Configure HTTP settings request.
type HTTPSettingsParameters struct {
	RestAuthCredentialIsEnabled *string `json:"restAuthCredential.isEnabled,omitempty"`
	RestAuthCredentialUsername  *string `json:"restAuthCredential.username,omitempty"`
	RestAuthCredentialPassword  *string `json:"restAuthCredential.password,omitempty"`
}

// LinuxOperatingSystemProfile is the ssh username, password, and ssh public key.
type LinuxOperatingSystemProfile struct {
	Username   *string     `json:"username,omitempty"`
	Password   *string     `json:"password,omitempty"`
	SSHProfile *SSHProfile `json:"sshProfile,omitempty"`
}

// Operation is hDInsight REST API operation
type Operation struct {
	Name    *string           `json:"name,omitempty"`
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay is the object that represents the operation.
type OperationDisplay struct {
	Provider  *string `json:"provider,omitempty"`
	Resource  *string `json:"resource,omitempty"`
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult is result of the request to list HDInsight operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Operation `json:"value,omitempty"`
	NextLink          *string      `json:"nextLink,omitempty"`
}

// OperationListResultPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client OperationListResult) OperationListResultPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// OperationResource is the azure async operation response.
type OperationResource struct {
	Status AsyncOperationState `json:"status,omitempty"`
	Error  *Errors             `json:"error,omitempty"`
}

// OsProfile is the Windows operation systems profile, and configure remote desktop settings.
type OsProfile struct {
	WindowsOperatingSystemProfile *WindowsOperatingSystemProfile `json:"windowsOperatingSystemProfile,omitempty"`
	LinuxOperatingSystemProfile   *LinuxOperatingSystemProfile   `json:"linuxOperatingSystemProfile,omitempty"`
}

// QuotaCapability is the regional quota capability.
type QuotaCapability struct {
	RegionalQuotas *[]RegionalQuotaCapability `json:"regionalQuotas,omitempty"`
}

// QuotaInfo is gets or sets Quota properties for the cluster.
type QuotaInfo struct {
	CoresUsed *int32 `json:"coresUsed,omitempty"`
}

// RdpSettings is the RDP settings for the windows cluster.
type RdpSettings struct {
	Username   *string    `json:"username,omitempty"`
	Password   *string    `json:"password,omitempty"`
	ExpiryDate *date.Date `json:"expiryDate,omitempty"`
}

// RDPSettingsParameters is parameters specifying the data factory gateway definition for a create or update operation.
type RDPSettingsParameters struct {
	OsProfile *OsProfile `json:"osProfile,omitempty"`
}

// RegionalQuotaCapability is the regional quota capacity.
type RegionalQuotaCapability struct {
	RegionName     *string `json:"region_name,omitempty"`
	CoresUsed      *int64  `json:"cores_used,omitempty"`
	CoresAvailable *int64  `json:"cores_available,omitempty"`
}

// RegionsCapability is the regions capability.
type RegionsCapability struct {
	Available *[]string `json:"available,omitempty"`
}

// Resource is the resource definition.
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// Role is describes a role on the cluster.
type Role struct {
	Name                  *string                `json:"name,omitempty"`
	MinInstanceCount      *int32                 `json:"minInstanceCount,omitempty"`
	TargetInstanceCount   *int32                 `json:"targetInstanceCount,omitempty"`
	HardwareProfile       *HardwareProfile       `json:"hardwareProfile,omitempty"`
	OsProfile             *OsProfile             `json:"osProfile,omitempty"`
	VirtualNetworkProfile *VirtualNetworkProfile `json:"virtualNetworkProfile,omitempty"`
	ScriptActions         *[]ScriptAction        `json:"scriptActions,omitempty"`
}

// RuntimeScriptAction is describes a script action on a running cluster.
type RuntimeScriptAction struct {
	Name            *string   `json:"name,omitempty"`
	URI             *string   `json:"uri,omitempty"`
	Parameters      *string   `json:"parameters,omitempty"`
	Roles           *[]string `json:"roles,omitempty"`
	ApplicationName *string   `json:"applicationName,omitempty"`
}

// RuntimeScriptActionDetail is describes the execution details of a script action.
type RuntimeScriptActionDetail struct {
	autorest.Response `json:"-"`
	ScriptExecutionID *int64                          `json:"scriptExecutionId,omitempty"`
	StartTime         *string                         `json:"startTime,omitempty"`
	EndTime           *string                         `json:"endTime,omitempty"`
	Status            *string                         `json:"status,omitempty"`
	Operation         *string                         `json:"operation,omitempty"`
	ExecutionSummary  *[]ScriptActionExecutionSummary `json:"executionSummary,omitempty"`
	DebugInformation  *string                         `json:"debugInformation,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	URI               *string                         `json:"uri,omitempty"`
	Parameters        *string                         `json:"parameters,omitempty"`
	Roles             *[]string                       `json:"roles,omitempty"`
	ApplicationName   *string                         `json:"applicationName,omitempty"`
}

// ScriptAction is describes a script action on role on the cluster.
type ScriptAction struct {
	Name       *string `json:"name,omitempty"`
	URI        *string `json:"uri,omitempty"`
	Parameters *string `json:"parameters,omitempty"`
}

// ScriptActionExecutionHistoryList is the ListScriptExecutionHistory response.
type ScriptActionExecutionHistoryList struct {
	autorest.Response `json:"-"`
	Value             *[]RuntimeScriptActionDetail `json:"value,omitempty"`
	NextLink          *string                      `json:"nextLink,omitempty"`
}

// ScriptActionExecutionHistoryListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ScriptActionExecutionHistoryList) ScriptActionExecutionHistoryListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// ScriptActionExecutionSummary is describes the execution summary of a script action.
type ScriptActionExecutionSummary struct {
	Status        *string `json:"status,omitempty"`
	InstanceCount *int32  `json:"instanceCount,omitempty"`
}

// ScriptActionPersistedGetResponseSpec is the persisted script action for cluster
type ScriptActionPersistedGetResponseSpec struct {
	Name            *string   `json:"name,omitempty"`
	URI             *string   `json:"uri,omitempty"`
	Parameters      *string   `json:"parameters,omitempty"`
	Roles           *[]string `json:"roles,omitempty"`
	ApplicationName *string   `json:"applicationName,omitempty"`
}

// ScriptActionsList is all persisted script action for the cluster.
type ScriptActionsList struct {
	autorest.Response `json:"-"`
	Value             *[]RuntimeScriptActionDetail `json:"value,omitempty"`
	NextLink          *string                      `json:"nextLink,omitempty"`
}

// ScriptActionsListPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client ScriptActionsList) ScriptActionsListPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// SecurityProfile is the security profile which contains Ssh public key for the HDInsight cluster.
type SecurityProfile struct {
	DirectoryType        DirectoryType `json:"directoryType,omitempty"`
	Domain               *string       `json:"domain,omitempty"`
	OrganizationalUnitDN *string       `json:"organizationalUnitDN,omitempty"`
	LdapsUrls            *[]string     `json:"ldapsUrls,omitempty"`
	DomainUsername       *string       `json:"domainUsername,omitempty"`
	DomainUserPassword   *string       `json:"domainUserPassword,omitempty"`
	ClusterUsersGroupDNS *[]string     `json:"clusterUsersGroupDNs,omitempty"`
}

// SSHProfile is the list of Ssh public keys.
type SSHProfile struct {
	PublicKeys *[]SSHPublicKey `json:"publicKeys,omitempty"`
}

// SSHPublicKey is the Ssh public key for the cluster nodes.
type SSHPublicKey struct {
	CertificateData *string `json:"certificateData,omitempty"`
}

// StorageAccount is describes the storage Account.
type StorageAccount struct {
	Name      *string `json:"name,omitempty"`
	IsDefault *bool   `json:"isDefault,omitempty"`
	Container *string `json:"container,omitempty"`
	Key       *string `json:"key,omitempty"`
}

// StorageProfile is describes the storage profile.
type StorageProfile struct {
	Storageaccounts *[]StorageAccount `json:"storageaccounts,omitempty"`
}

// SubResource is the sub resource definition.
type SubResource struct {
	ID *string `json:"id,omitempty"`
}

// VersionsCapability is the version capability.
type VersionsCapability struct {
	Available *[]VersionSpec `json:"available,omitempty"`
}

// VersionSpec is gets or sets Version spec properties.
type VersionSpec struct {
	FriendlyName      *string             `json:"friendlyName,omitempty"`
	DisplayName       *string             `json:"displayName,omitempty"`
	IsDefault         *string             `json:"isDefault,omitempty"`
	ComponentVersions *map[string]*string `json:"componentVersions,omitempty"`
}

// VirtualNetworkProfile is the Virtual network properties.
type VirtualNetworkProfile struct {
	ID     *string `json:"id,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
}

// VMSizeCompatibilityFilter is the virtual machine type compatibility filter.
type VMSizeCompatibilityFilter struct {
	FilterMode      *string   `json:"FilterMode,omitempty"`
	Regions         *[]string `json:"Regions,omitempty"`
	ClusterFlavors  *[]string `json:"ClusterFlavors,omitempty"`
	NodeTypes       *[]string `json:"NodeTypes,omitempty"`
	ClusterVersions *[]string `json:"ClusterVersions,omitempty"`
	Vmsizes         *[]string `json:"vmsizes,omitempty"`
}

// VMSizesCapability is the virtual machine sizes capability.
type VMSizesCapability struct {
	Available *[]string `json:"available,omitempty"`
}

// WindowsOperatingSystemProfile is the Windows operation system settings.
type WindowsOperatingSystemProfile struct {
	RdpSettings *RdpSettings `json:"rdpSettings,omitempty"`
}