//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package managementgroups

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/resources/mgmt/2018-03-01-preview/managementgroups"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type InheritedPermissions = original.InheritedPermissions

const (
	Delete   InheritedPermissions = original.Delete
	Edit     InheritedPermissions = original.Edit
	Noaccess InheritedPermissions = original.Noaccess
	View     InheritedPermissions = original.View
)

type Permissions = original.Permissions

const (
	PermissionsDelete   Permissions = original.PermissionsDelete
	PermissionsEdit     Permissions = original.PermissionsEdit
	PermissionsNoaccess Permissions = original.PermissionsNoaccess
	PermissionsView     Permissions = original.PermissionsView
)

type Permissions1 = original.Permissions1

const (
	Permissions1Delete   Permissions1 = original.Permissions1Delete
	Permissions1Edit     Permissions1 = original.Permissions1Edit
	Permissions1Noaccess Permissions1 = original.Permissions1Noaccess
	Permissions1View     Permissions1 = original.Permissions1View
)

type ProvisioningState = original.ProvisioningState

const (
	Updating ProvisioningState = original.Updating
)

type Reason = original.Reason

const (
	AlreadyExists Reason = original.AlreadyExists
	Invalid       Reason = original.Invalid
)

type Status = original.Status

const (
	Cancelled                Status = original.Cancelled
	Completed                Status = original.Completed
	Failed                   Status = original.Failed
	NotStarted               Status = original.NotStarted
	NotStartedButGroupsExist Status = original.NotStartedButGroupsExist
	Started                  Status = original.Started
)

type Type = original.Type

const (
	ProvidersMicrosoftManagementmanagementGroups Type = original.ProvidersMicrosoftManagementmanagementGroups
)

type Type1 = original.Type1

const (
	Type1ProvidersMicrosoftManagementmanagementGroups Type1 = original.Type1ProvidersMicrosoftManagementmanagementGroups
	Type1Subscriptions                                Type1 = original.Type1Subscriptions
)

type Type2 = original.Type2

const (
	Type2ProvidersMicrosoftManagementmanagementGroups Type2 = original.Type2ProvidersMicrosoftManagementmanagementGroups
	Type2Subscriptions                                Type2 = original.Type2Subscriptions
)

type BaseClient = original.BaseClient
type CheckNameAvailabilityRequest = original.CheckNameAvailabilityRequest
type CheckNameAvailabilityResult = original.CheckNameAvailabilityResult
type ChildInfo = original.ChildInfo
type Client = original.Client
type CreateManagementGroupChildInfo = original.CreateManagementGroupChildInfo
type CreateManagementGroupDetails = original.CreateManagementGroupDetails
type CreateManagementGroupProperties = original.CreateManagementGroupProperties
type CreateManagementGroupRequest = original.CreateManagementGroupRequest
type CreateOrUpdateFuture = original.CreateOrUpdateFuture
type CreateParentGroupInfo = original.CreateParentGroupInfo
type DeleteFuture = original.DeleteFuture
type DescendantInfo = original.DescendantInfo
type DescendantInfoProperties = original.DescendantInfoProperties
type DescendantListResult = original.DescendantListResult
type DescendantListResultIterator = original.DescendantListResultIterator
type DescendantListResultPage = original.DescendantListResultPage
type DescendantParentGroupInfo = original.DescendantParentGroupInfo
type Details = original.Details
type EntitiesClient = original.EntitiesClient
type EntityHierarchyItem = original.EntityHierarchyItem
type EntityHierarchyItemProperties = original.EntityHierarchyItemProperties
type EntityInfo = original.EntityInfo
type EntityInfoProperties = original.EntityInfoProperties
type EntityListResult = original.EntityListResult
type EntityListResultIterator = original.EntityListResultIterator
type EntityListResultPage = original.EntityListResultPage
type EntityParentGroupInfo = original.EntityParentGroupInfo
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Info = original.Info
type InfoProperties = original.InfoProperties
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagementGroup = original.ManagementGroup
type Operation = original.Operation
type OperationDisplayProperties = original.OperationDisplayProperties
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResults = original.OperationResults
type OperationResultsProperties = original.OperationResultsProperties
type OperationsClient = original.OperationsClient
type ParentGroupInfo = original.ParentGroupInfo
type PatchManagementGroupRequest = original.PatchManagementGroupRequest
type Properties = original.Properties
type SetObject = original.SetObject
type SubscriptionsClient = original.SubscriptionsClient
type TenantBackfillStatusResult = original.TenantBackfillStatusResult

func New() BaseClient {
	return original.New()
}
func NewClient() Client {
	return original.NewClient()
}
func NewClientWithBaseURI(baseURI string) Client {
	return original.NewClientWithBaseURI(baseURI)
}
func NewDescendantListResultIterator(page DescendantListResultPage) DescendantListResultIterator {
	return original.NewDescendantListResultIterator(page)
}
func NewDescendantListResultPage(cur DescendantListResult, getNextPage func(context.Context, DescendantListResult) (DescendantListResult, error)) DescendantListResultPage {
	return original.NewDescendantListResultPage(cur, getNextPage)
}
func NewEntitiesClient() EntitiesClient {
	return original.NewEntitiesClient()
}
func NewEntitiesClientWithBaseURI(baseURI string) EntitiesClient {
	return original.NewEntitiesClientWithBaseURI(baseURI)
}
func NewEntityListResultIterator(page EntityListResultPage) EntityListResultIterator {
	return original.NewEntityListResultIterator(page)
}
func NewEntityListResultPage(cur EntityListResult, getNextPage func(context.Context, EntityListResult) (EntityListResult, error)) EntityListResultPage {
	return original.NewEntityListResultPage(cur, getNextPage)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient() OperationsClient {
	return original.NewOperationsClient()
}
func NewOperationsClientWithBaseURI(baseURI string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI)
}
func NewSubscriptionsClient() SubscriptionsClient {
	return original.NewSubscriptionsClient()
}
func NewSubscriptionsClientWithBaseURI(baseURI string) SubscriptionsClient {
	return original.NewSubscriptionsClientWithBaseURI(baseURI)
}
func NewWithBaseURI(baseURI string) BaseClient {
	return original.NewWithBaseURI(baseURI)
}
func PossibleInheritedPermissionsValues() []InheritedPermissions {
	return original.PossibleInheritedPermissionsValues()
}
func PossiblePermissions1Values() []Permissions1 {
	return original.PossiblePermissions1Values()
}
func PossiblePermissionsValues() []Permissions {
	return original.PossiblePermissionsValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReasonValues() []Reason {
	return original.PossibleReasonValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleType1Values() []Type1 {
	return original.PossibleType1Values()
}
func PossibleType2Values() []Type2 {
	return original.PossibleType2Values()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
