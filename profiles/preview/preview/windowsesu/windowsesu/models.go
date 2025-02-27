//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package windowsesu

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/windowsesu/2019-09-16-preview/windowsesu"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type OsType = original.OsType

const (
	Windows7            OsType = original.Windows7
	WindowsServer2008   OsType = original.WindowsServer2008
	WindowsServer2008R2 OsType = original.WindowsServer2008R2
)

type ProvisioningState = original.ProvisioningState

const (
	Accepted     ProvisioningState = original.Accepted
	Canceled     ProvisioningState = original.Canceled
	Failed       ProvisioningState = original.Failed
	Provisioning ProvisioningState = original.Provisioning
	Succeeded    ProvisioningState = original.Succeeded
)

type SupportType = original.SupportType

const (
	PremiumAssurance      SupportType = original.PremiumAssurance
	SupplementalServicing SupportType = original.SupplementalServicing
)

type AzureEntityResource = original.AzureEntityResource
type BaseClient = original.BaseClient
type ErrorDefinition = original.ErrorDefinition
type ErrorResponse = original.ErrorResponse
type MultipleActivationKey = original.MultipleActivationKey
type MultipleActivationKeyList = original.MultipleActivationKeyList
type MultipleActivationKeyListIterator = original.MultipleActivationKeyListIterator
type MultipleActivationKeyListPage = original.MultipleActivationKeyListPage
type MultipleActivationKeyProperties = original.MultipleActivationKeyProperties
type MultipleActivationKeyUpdate = original.MultipleActivationKeyUpdate
type MultipleActivationKeysClient = original.MultipleActivationKeysClient
type MultipleActivationKeysCreateFuture = original.MultipleActivationKeysCreateFuture
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationList = original.OperationList
type OperationListIterator = original.OperationListIterator
type OperationListPage = original.OperationListPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type TrackedResource = original.TrackedResource

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewMultipleActivationKeyListIterator(page MultipleActivationKeyListPage) MultipleActivationKeyListIterator {
	return original.NewMultipleActivationKeyListIterator(page)
}
func NewMultipleActivationKeyListPage(cur MultipleActivationKeyList, getNextPage func(context.Context, MultipleActivationKeyList) (MultipleActivationKeyList, error)) MultipleActivationKeyListPage {
	return original.NewMultipleActivationKeyListPage(cur, getNextPage)
}
func NewMultipleActivationKeysClient(subscriptionID string) MultipleActivationKeysClient {
	return original.NewMultipleActivationKeysClient(subscriptionID)
}
func NewMultipleActivationKeysClientWithBaseURI(baseURI string, subscriptionID string) MultipleActivationKeysClient {
	return original.NewMultipleActivationKeysClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListIterator(page OperationListPage) OperationListIterator {
	return original.NewOperationListIterator(page)
}
func NewOperationListPage(cur OperationList, getNextPage func(context.Context, OperationList) (OperationList, error)) OperationListPage {
	return original.NewOperationListPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleOsTypeValues() []OsType {
	return original.PossibleOsTypeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleSupportTypeValues() []SupportType {
	return original.PossibleSupportTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
