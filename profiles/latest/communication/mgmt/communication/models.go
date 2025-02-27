//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package communication

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/communication/mgmt/2020-08-20/communication"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ActionType = original.ActionType

const (
	Internal ActionType = original.Internal
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type KeyType = original.KeyType

const (
	Primary   KeyType = original.Primary
	Secondary KeyType = original.Secondary
)

type Origin = original.Origin

const (
	OriginSystem     Origin = original.OriginSystem
	OriginUser       Origin = original.OriginUser
	OriginUsersystem Origin = original.OriginUsersystem
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Moving    ProvisioningState = original.Moving
	Running   ProvisioningState = original.Running
	Succeeded ProvisioningState = original.Succeeded
	Unknown   ProvisioningState = original.Unknown
	Updating  ProvisioningState = original.Updating
)

type Status = original.Status

const (
	StatusCanceled  Status = original.StatusCanceled
	StatusCreating  Status = original.StatusCreating
	StatusDeleting  Status = original.StatusDeleting
	StatusFailed    Status = original.StatusFailed
	StatusMoving    Status = original.StatusMoving
	StatusSucceeded Status = original.StatusSucceeded
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ErrorAdditionalInfo = original.ErrorAdditionalInfo
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type LinkNotificationHubParameters = original.LinkNotificationHubParameters
type LinkedNotificationHub = original.LinkedNotificationHub
type LocationResource = original.LocationResource
type NameAvailability = original.NameAvailability
type NameAvailabilityParameters = original.NameAvailabilityParameters
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationStatus = original.OperationStatus
type OperationStatusesClient = original.OperationStatusesClient
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type RegenerateKeyParameters = original.RegenerateKeyParameters
type Resource = original.Resource
type ServiceClient = original.ServiceClient
type ServiceCreateOrUpdateFuture = original.ServiceCreateOrUpdateFuture
type ServiceDeleteFuture = original.ServiceDeleteFuture
type ServiceKeys = original.ServiceKeys
type ServiceProperties = original.ServiceProperties
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceListIterator = original.ServiceResourceListIterator
type ServiceResourceListPage = original.ServiceResourceListPage
type SystemData = original.SystemData
type TaggedResource = original.TaggedResource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationStatusesClient(subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClient(subscriptionID)
}
func NewOperationStatusesClientWithBaseURI(baseURI string, subscriptionID string) OperationStatusesClient {
	return original.NewOperationStatusesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceClient(subscriptionID string) ServiceClient {
	return original.NewServiceClient(subscriptionID)
}
func NewServiceClientWithBaseURI(baseURI string, subscriptionID string) ServiceClient {
	return original.NewServiceClientWithBaseURI(baseURI, subscriptionID)
}
func NewServiceResourceListIterator(page ServiceResourceListPage) ServiceResourceListIterator {
	return original.NewServiceResourceListIterator(page)
}
func NewServiceResourceListPage(cur ServiceResourceList, getNextPage func(context.Context, ServiceResourceList) (ServiceResourceList, error)) ServiceResourceListPage {
	return original.NewServiceResourceListPage(cur, getNextPage)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleActionTypeValues() []ActionType {
	return original.PossibleActionTypeValues()
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleKeyTypeValues() []KeyType {
	return original.PossibleKeyTypeValues()
}
func PossibleOriginValues() []Origin {
	return original.PossibleOriginValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
