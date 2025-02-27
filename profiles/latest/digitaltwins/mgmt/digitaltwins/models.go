//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package digitaltwins

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/digitaltwins/mgmt/2020-12-01/digitaltwins"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AuthenticationType = original.AuthenticationType

const (
	IdentityBased AuthenticationType = original.IdentityBased
	KeyBased      AuthenticationType = original.KeyBased
)

type ConnectionPropertiesProvisioningState = original.ConnectionPropertiesProvisioningState

const (
	Approved     ConnectionPropertiesProvisioningState = original.Approved
	Disconnected ConnectionPropertiesProvisioningState = original.Disconnected
	Pending      ConnectionPropertiesProvisioningState = original.Pending
	Rejected     ConnectionPropertiesProvisioningState = original.Rejected
)

type EndpointProvisioningState = original.EndpointProvisioningState

const (
	Canceled     EndpointProvisioningState = original.Canceled
	Deleted      EndpointProvisioningState = original.Deleted
	Deleting     EndpointProvisioningState = original.Deleting
	Disabled     EndpointProvisioningState = original.Disabled
	Failed       EndpointProvisioningState = original.Failed
	Moving       EndpointProvisioningState = original.Moving
	Provisioning EndpointProvisioningState = original.Provisioning
	Restoring    EndpointProvisioningState = original.Restoring
	Succeeded    EndpointProvisioningState = original.Succeeded
	Suspending   EndpointProvisioningState = original.Suspending
	Warning      EndpointProvisioningState = original.Warning
)

type EndpointType = original.EndpointType

const (
	EndpointTypeDigitalTwinsEndpointResourceProperties EndpointType = original.EndpointTypeDigitalTwinsEndpointResourceProperties
	EndpointTypeEventGrid                              EndpointType = original.EndpointTypeEventGrid
	EndpointTypeEventHub                               EndpointType = original.EndpointTypeEventHub
	EndpointTypeServiceBus                             EndpointType = original.EndpointTypeServiceBus
)

type IdentityType = original.IdentityType

const (
	None           IdentityType = original.None
	SystemAssigned IdentityType = original.SystemAssigned
)

type PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatus

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusApproved
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusDisconnected
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusPending
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = original.PrivateLinkServiceConnectionStatusRejected
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateCanceled     ProvisioningState = original.ProvisioningStateCanceled
	ProvisioningStateDeleted      ProvisioningState = original.ProvisioningStateDeleted
	ProvisioningStateDeleting     ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed       ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoving       ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateProvisioning ProvisioningState = original.ProvisioningStateProvisioning
	ProvisioningStateRestoring    ProvisioningState = original.ProvisioningStateRestoring
	ProvisioningStateSucceeded    ProvisioningState = original.ProvisioningStateSucceeded
	ProvisioningStateSuspending   ProvisioningState = original.ProvisioningStateSuspending
	ProvisioningStateUpdating     ProvisioningState = original.ProvisioningStateUpdating
	ProvisioningStateWarning      ProvisioningState = original.ProvisioningStateWarning
)

type PublicNetworkAccess = original.PublicNetworkAccess

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = original.PublicNetworkAccessDisabled
	PublicNetworkAccessEnabled  PublicNetworkAccess = original.PublicNetworkAccessEnabled
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type BaseClient = original.BaseClient
type BasicEndpointResourceProperties = original.BasicEndpointResourceProperties
type CheckNameRequest = original.CheckNameRequest
type CheckNameResult = original.CheckNameResult
type Client = original.Client
type ConnectionProperties = original.ConnectionProperties
type ConnectionPropertiesPrivateEndpoint = original.ConnectionPropertiesPrivateEndpoint
type ConnectionPropertiesPrivateLinkServiceConnectionState = original.ConnectionPropertiesPrivateLinkServiceConnectionState
type ConnectionState = original.ConnectionState
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type DeleteFuture = original.DeleteFuture
type Description = original.Description
type DescriptionListResult = original.DescriptionListResult
type DescriptionListResultIterator = original.DescriptionListResultIterator
type DescriptionListResultPage = original.DescriptionListResultPage
type EndpointClient = original.EndpointClient
type EndpointCreateOrUpdateFuture = original.EndpointCreateOrUpdateFuture
type EndpointDeleteFuture = original.EndpointDeleteFuture
type EndpointResource = original.EndpointResource
type EndpointResourceListResult = original.EndpointResourceListResult
type EndpointResourceListResultIterator = original.EndpointResourceListResultIterator
type EndpointResourceListResultPage = original.EndpointResourceListResultPage
type EndpointResourceProperties = original.EndpointResourceProperties
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type EventGrid = original.EventGrid
type EventHub = original.EventHub
type ExternalResource = original.ExternalResource
type GroupIDInformation = original.GroupIDInformation
type GroupIDInformationProperties = original.GroupIDInformationProperties
type GroupIDInformationPropertiesModel = original.GroupIDInformationPropertiesModel
type GroupIDInformationResponse = original.GroupIDInformationResponse
type Identity = original.Identity
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PatchDescription = original.PatchDescription
type PatchProperties = original.PatchProperties
type PrivateEndpoint = original.PrivateEndpoint
type PrivateEndpointConnection = original.PrivateEndpointConnection
type PrivateEndpointConnectionProperties = original.PrivateEndpointConnectionProperties
type PrivateEndpointConnectionsClient = original.PrivateEndpointConnectionsClient
type PrivateEndpointConnectionsCreateOrUpdateFuture = original.PrivateEndpointConnectionsCreateOrUpdateFuture
type PrivateEndpointConnectionsDeleteFuture = original.PrivateEndpointConnectionsDeleteFuture
type PrivateEndpointConnectionsResponse = original.PrivateEndpointConnectionsResponse
type PrivateLinkResourcesClient = original.PrivateLinkResourcesClient
type Properties = original.Properties
type Resource = original.Resource
type ServiceBus = original.ServiceBus
type UpdateFuture = original.UpdateFuture

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClient(subscriptionID string) Client {
	return original.NewClient(subscriptionID)
}
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return original.NewClientWithBaseURI(baseURI, subscriptionID)
}
func NewDescriptionListResultIterator(page DescriptionListResultPage) DescriptionListResultIterator {
	return original.NewDescriptionListResultIterator(page)
}
func NewDescriptionListResultPage(cur DescriptionListResult, getNextPage func(context.Context, DescriptionListResult) (DescriptionListResult, error)) DescriptionListResultPage {
	return original.NewDescriptionListResultPage(cur, getNextPage)
}
func NewEndpointClient(subscriptionID string) EndpointClient {
	return original.NewEndpointClient(subscriptionID)
}
func NewEndpointClientWithBaseURI(baseURI string, subscriptionID string) EndpointClient {
	return original.NewEndpointClientWithBaseURI(baseURI, subscriptionID)
}
func NewEndpointResourceListResultIterator(page EndpointResourceListResultPage) EndpointResourceListResultIterator {
	return original.NewEndpointResourceListResultIterator(page)
}
func NewEndpointResourceListResultPage(cur EndpointResourceListResult, getNextPage func(context.Context, EndpointResourceListResult) (EndpointResourceListResult, error)) EndpointResourceListResultPage {
	return original.NewEndpointResourceListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateEndpointConnectionsClient(subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClient(subscriptionID)
}
func NewPrivateEndpointConnectionsClientWithBaseURI(baseURI string, subscriptionID string) PrivateEndpointConnectionsClient {
	return original.NewPrivateEndpointConnectionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPrivateLinkResourcesClient(subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClient(subscriptionID)
}
func NewPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) PrivateLinkResourcesClient {
	return original.NewPrivateLinkResourcesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAuthenticationTypeValues() []AuthenticationType {
	return original.PossibleAuthenticationTypeValues()
}
func PossibleConnectionPropertiesProvisioningStateValues() []ConnectionPropertiesProvisioningState {
	return original.PossibleConnectionPropertiesProvisioningStateValues()
}
func PossibleEndpointProvisioningStateValues() []EndpointProvisioningState {
	return original.PossibleEndpointProvisioningStateValues()
}
func PossibleEndpointTypeValues() []EndpointType {
	return original.PossibleEndpointTypeValues()
}
func PossibleIdentityTypeValues() []IdentityType {
	return original.PossibleIdentityTypeValues()
}
func PossiblePrivateLinkServiceConnectionStatusValues() []PrivateLinkServiceConnectionStatus {
	return original.PossiblePrivateLinkServiceConnectionStatusValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return original.PossiblePublicNetworkAccessValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
