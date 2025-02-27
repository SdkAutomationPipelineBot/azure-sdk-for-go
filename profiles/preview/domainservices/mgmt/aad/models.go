//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package aad

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/domainservices/mgmt/2020-01-01/aad"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ExternalAccess = original.ExternalAccess

const (
	Disabled ExternalAccess = original.Disabled
	Enabled  ExternalAccess = original.Enabled
)

type FilteredSync = original.FilteredSync

const (
	FilteredSyncDisabled FilteredSync = original.FilteredSyncDisabled
	FilteredSyncEnabled  FilteredSync = original.FilteredSyncEnabled
)

type Ldaps = original.Ldaps

const (
	LdapsDisabled Ldaps = original.LdapsDisabled
	LdapsEnabled  Ldaps = original.LdapsEnabled
)

type NotifyDcAdmins = original.NotifyDcAdmins

const (
	NotifyDcAdminsDisabled NotifyDcAdmins = original.NotifyDcAdminsDisabled
	NotifyDcAdminsEnabled  NotifyDcAdmins = original.NotifyDcAdminsEnabled
)

type NotifyGlobalAdmins = original.NotifyGlobalAdmins

const (
	NotifyGlobalAdminsDisabled NotifyGlobalAdmins = original.NotifyGlobalAdminsDisabled
	NotifyGlobalAdminsEnabled  NotifyGlobalAdmins = original.NotifyGlobalAdminsEnabled
)

type NtlmV1 = original.NtlmV1

const (
	NtlmV1Disabled NtlmV1 = original.NtlmV1Disabled
	NtlmV1Enabled  NtlmV1 = original.NtlmV1Enabled
)

type SyncKerberosPasswords = original.SyncKerberosPasswords

const (
	SyncKerberosPasswordsDisabled SyncKerberosPasswords = original.SyncKerberosPasswordsDisabled
	SyncKerberosPasswordsEnabled  SyncKerberosPasswords = original.SyncKerberosPasswordsEnabled
)

type SyncNtlmPasswords = original.SyncNtlmPasswords

const (
	SyncNtlmPasswordsDisabled SyncNtlmPasswords = original.SyncNtlmPasswordsDisabled
	SyncNtlmPasswordsEnabled  SyncNtlmPasswords = original.SyncNtlmPasswordsEnabled
)

type SyncOnPremPasswords = original.SyncOnPremPasswords

const (
	SyncOnPremPasswordsDisabled SyncOnPremPasswords = original.SyncOnPremPasswordsDisabled
	SyncOnPremPasswordsEnabled  SyncOnPremPasswords = original.SyncOnPremPasswordsEnabled
)

type TLSV1 = original.TLSV1

const (
	TLSV1Disabled TLSV1 = original.TLSV1Disabled
	TLSV1Enabled  TLSV1 = original.TLSV1Enabled
)

type BaseClient = original.BaseClient
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ContainerAccount = original.ContainerAccount
type DomainSecuritySettings = original.DomainSecuritySettings
type DomainService = original.DomainService
type DomainServiceListResult = original.DomainServiceListResult
type DomainServiceListResultIterator = original.DomainServiceListResultIterator
type DomainServiceListResultPage = original.DomainServiceListResultPage
type DomainServiceOperationsClient = original.DomainServiceOperationsClient
type DomainServiceProperties = original.DomainServiceProperties
type DomainServicesClient = original.DomainServicesClient
type DomainServicesCreateOrUpdateFuture = original.DomainServicesCreateOrUpdateFuture
type DomainServicesDeleteFuture = original.DomainServicesDeleteFuture
type DomainServicesUpdateFuture = original.DomainServicesUpdateFuture
type ForestTrust = original.ForestTrust
type HealthAlert = original.HealthAlert
type HealthMonitor = original.HealthMonitor
type LdapsSettings = original.LdapsSettings
type NotificationSettings = original.NotificationSettings
type OperationDisplayInfo = original.OperationDisplayInfo
type OperationEntity = original.OperationEntity
type OperationEntityListResult = original.OperationEntityListResult
type OperationEntityListResultIterator = original.OperationEntityListResultIterator
type OperationEntityListResultPage = original.OperationEntityListResultPage
type OuContainer = original.OuContainer
type OuContainerClient = original.OuContainerClient
type OuContainerCreateFuture = original.OuContainerCreateFuture
type OuContainerDeleteFuture = original.OuContainerDeleteFuture
type OuContainerListResult = original.OuContainerListResult
type OuContainerListResultIterator = original.OuContainerListResultIterator
type OuContainerListResultPage = original.OuContainerListResultPage
type OuContainerOperationsClient = original.OuContainerOperationsClient
type OuContainerProperties = original.OuContainerProperties
type OuContainerUpdateFuture = original.OuContainerUpdateFuture
type ReplicaSet = original.ReplicaSet
type Resource = original.Resource
type ResourceForestSettings = original.ResourceForestSettings

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewDomainServiceListResultIterator(page DomainServiceListResultPage) DomainServiceListResultIterator {
	return original.NewDomainServiceListResultIterator(page)
}
func NewDomainServiceListResultPage(cur DomainServiceListResult, getNextPage func(context.Context, DomainServiceListResult) (DomainServiceListResult, error)) DomainServiceListResultPage {
	return original.NewDomainServiceListResultPage(cur, getNextPage)
}
func NewDomainServiceOperationsClient(subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClient(subscriptionID)
}
func NewDomainServiceOperationsClientWithBaseURI(baseURI string, subscriptionID string) DomainServiceOperationsClient {
	return original.NewDomainServiceOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDomainServicesClient(subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClient(subscriptionID)
}
func NewDomainServicesClientWithBaseURI(baseURI string, subscriptionID string) DomainServicesClient {
	return original.NewDomainServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationEntityListResultIterator(page OperationEntityListResultPage) OperationEntityListResultIterator {
	return original.NewOperationEntityListResultIterator(page)
}
func NewOperationEntityListResultPage(cur OperationEntityListResult, getNextPage func(context.Context, OperationEntityListResult) (OperationEntityListResult, error)) OperationEntityListResultPage {
	return original.NewOperationEntityListResultPage(cur, getNextPage)
}
func NewOuContainerClient(subscriptionID string) OuContainerClient {
	return original.NewOuContainerClient(subscriptionID)
}
func NewOuContainerClientWithBaseURI(baseURI string, subscriptionID string) OuContainerClient {
	return original.NewOuContainerClientWithBaseURI(baseURI, subscriptionID)
}
func NewOuContainerListResultIterator(page OuContainerListResultPage) OuContainerListResultIterator {
	return original.NewOuContainerListResultIterator(page)
}
func NewOuContainerListResultPage(cur OuContainerListResult, getNextPage func(context.Context, OuContainerListResult) (OuContainerListResult, error)) OuContainerListResultPage {
	return original.NewOuContainerListResultPage(cur, getNextPage)
}
func NewOuContainerOperationsClient(subscriptionID string) OuContainerOperationsClient {
	return original.NewOuContainerOperationsClient(subscriptionID)
}
func NewOuContainerOperationsClientWithBaseURI(baseURI string, subscriptionID string) OuContainerOperationsClient {
	return original.NewOuContainerOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleExternalAccessValues() []ExternalAccess {
	return original.PossibleExternalAccessValues()
}
func PossibleFilteredSyncValues() []FilteredSync {
	return original.PossibleFilteredSyncValues()
}
func PossibleLdapsValues() []Ldaps {
	return original.PossibleLdapsValues()
}
func PossibleNotifyDcAdminsValues() []NotifyDcAdmins {
	return original.PossibleNotifyDcAdminsValues()
}
func PossibleNotifyGlobalAdminsValues() []NotifyGlobalAdmins {
	return original.PossibleNotifyGlobalAdminsValues()
}
func PossibleNtlmV1Values() []NtlmV1 {
	return original.PossibleNtlmV1Values()
}
func PossibleSyncKerberosPasswordsValues() []SyncKerberosPasswords {
	return original.PossibleSyncKerberosPasswordsValues()
}
func PossibleSyncNtlmPasswordsValues() []SyncNtlmPasswords {
	return original.PossibleSyncNtlmPasswordsValues()
}
func PossibleSyncOnPremPasswordsValues() []SyncOnPremPasswords {
	return original.PossibleSyncOnPremPasswordsValues()
}
func PossibleTLSV1Values() []TLSV1 {
	return original.PossibleTLSV1Values()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
