//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package attestation

import original "github.com/Azure/azure-sdk-for-go/services/attestation/mgmt/2020-10-01/attestation"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CreatedByType = original.CreatedByType

const (
	Application     CreatedByType = original.Application
	Key             CreatedByType = original.Key
	ManagedIdentity CreatedByType = original.ManagedIdentity
	User            CreatedByType = original.User
)

type ServiceStatus = original.ServiceStatus

const (
	Error    ServiceStatus = original.Error
	NotReady ServiceStatus = original.NotReady
	Ready    ServiceStatus = original.Ready
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type JSONWebKey = original.JSONWebKey
type JSONWebKeySet = original.JSONWebKeySet
type OperationList = original.OperationList
type OperationsClient = original.OperationsClient
type OperationsDefinition = original.OperationsDefinition
type OperationsDisplayDefinition = original.OperationsDisplayDefinition
type Provider = original.Provider
type ProviderListResult = original.ProviderListResult
type ProvidersClient = original.ProvidersClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServiceCreationParams = original.ServiceCreationParams
type ServiceCreationSpecificParams = original.ServiceCreationSpecificParams
type ServicePatchParams = original.ServicePatchParams
type StatusResult = original.StatusResult
type SystemData = original.SystemData
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProvidersClient(subscriptionID string) ProvidersClient {
	return original.NewProvidersClient(subscriptionID)
}
func NewProvidersClientWithBaseURI(baseURI string, subscriptionID string) ProvidersClient {
	return original.NewProvidersClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleCreatedByTypeValues() []CreatedByType {
	return original.PossibleCreatedByTypeValues()
}
func PossibleServiceStatusValues() []ServiceStatus {
	return original.PossibleServiceStatusValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
