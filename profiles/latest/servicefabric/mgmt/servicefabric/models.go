//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicefabric

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/servicefabric/mgmt/2019-03-01/servicefabric"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ArmServicePackageActivationMode = original.ArmServicePackageActivationMode

const (
	ExclusiveProcess ArmServicePackageActivationMode = original.ExclusiveProcess
	SharedProcess    ArmServicePackageActivationMode = original.SharedProcess
)

type ArmUpgradeFailureAction = original.ArmUpgradeFailureAction

const (
	Manual   ArmUpgradeFailureAction = original.Manual
	Rollback ArmUpgradeFailureAction = original.Rollback
)

type ClusterState = original.ClusterState

const (
	AutoScale                 ClusterState = original.AutoScale
	BaselineUpgrade           ClusterState = original.BaselineUpgrade
	Deploying                 ClusterState = original.Deploying
	EnforcingClusterVersion   ClusterState = original.EnforcingClusterVersion
	Ready                     ClusterState = original.Ready
	UpdatingInfrastructure    ClusterState = original.UpdatingInfrastructure
	UpdatingUserCertificate   ClusterState = original.UpdatingUserCertificate
	UpdatingUserConfiguration ClusterState = original.UpdatingUserConfiguration
	UpgradeServiceUnreachable ClusterState = original.UpgradeServiceUnreachable
	WaitingForNodes           ClusterState = original.WaitingForNodes
)

type DurabilityLevel = original.DurabilityLevel

const (
	Bronze DurabilityLevel = original.Bronze
	Gold   DurabilityLevel = original.Gold
	Silver DurabilityLevel = original.Silver
)

type Environment = original.Environment

const (
	Linux   Environment = original.Linux
	Windows Environment = original.Windows
)

type MoveCost = original.MoveCost

const (
	High   MoveCost = original.High
	Low    MoveCost = original.Low
	Medium MoveCost = original.Medium
	Zero   MoveCost = original.Zero
)

type PartitionScheme = original.PartitionScheme

const (
	Invalid           PartitionScheme = original.Invalid
	Named             PartitionScheme = original.Named
	Singleton         PartitionScheme = original.Singleton
	UniformInt64Range PartitionScheme = original.UniformInt64Range
)

type PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeBasicPartitionSchemeDescription

const (
	PartitionSchemeNamed                      PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeNamed
	PartitionSchemePartitionSchemeDescription PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemePartitionSchemeDescription
	PartitionSchemeSingleton                  PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeSingleton
	PartitionSchemeUniformInt64Range          PartitionSchemeBasicPartitionSchemeDescription = original.PartitionSchemeUniformInt64Range
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Updating  ProvisioningState = original.Updating
)

type ReliabilityLevel = original.ReliabilityLevel

const (
	ReliabilityLevelBronze   ReliabilityLevel = original.ReliabilityLevelBronze
	ReliabilityLevelGold     ReliabilityLevel = original.ReliabilityLevelGold
	ReliabilityLevelNone     ReliabilityLevel = original.ReliabilityLevelNone
	ReliabilityLevelPlatinum ReliabilityLevel = original.ReliabilityLevelPlatinum
	ReliabilityLevelSilver   ReliabilityLevel = original.ReliabilityLevelSilver
)

type ReliabilityLevel1 = original.ReliabilityLevel1

const (
	ReliabilityLevel1Bronze   ReliabilityLevel1 = original.ReliabilityLevel1Bronze
	ReliabilityLevel1Gold     ReliabilityLevel1 = original.ReliabilityLevel1Gold
	ReliabilityLevel1None     ReliabilityLevel1 = original.ReliabilityLevel1None
	ReliabilityLevel1Platinum ReliabilityLevel1 = original.ReliabilityLevel1Platinum
	ReliabilityLevel1Silver   ReliabilityLevel1 = original.ReliabilityLevel1Silver
)

type ServiceCorrelationScheme = original.ServiceCorrelationScheme

const (
	ServiceCorrelationSchemeAffinity           ServiceCorrelationScheme = original.ServiceCorrelationSchemeAffinity
	ServiceCorrelationSchemeAlignedAffinity    ServiceCorrelationScheme = original.ServiceCorrelationSchemeAlignedAffinity
	ServiceCorrelationSchemeInvalid            ServiceCorrelationScheme = original.ServiceCorrelationSchemeInvalid
	ServiceCorrelationSchemeNonAlignedAffinity ServiceCorrelationScheme = original.ServiceCorrelationSchemeNonAlignedAffinity
)

type ServiceKind = original.ServiceKind

const (
	ServiceKindInvalid   ServiceKind = original.ServiceKindInvalid
	ServiceKindStateful  ServiceKind = original.ServiceKindStateful
	ServiceKindStateless ServiceKind = original.ServiceKindStateless
)

type ServiceKindBasicServiceResourceProperties = original.ServiceKindBasicServiceResourceProperties

const (
	ServiceKindServiceResourceProperties ServiceKindBasicServiceResourceProperties = original.ServiceKindServiceResourceProperties
	ServiceKindStateful1                 ServiceKindBasicServiceResourceProperties = original.ServiceKindStateful1
	ServiceKindStateless1                ServiceKindBasicServiceResourceProperties = original.ServiceKindStateless1
)

type ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdateProperties

const (
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindServiceResourceUpdateProperties
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful                        ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateful
	ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless                       ServiceKindBasicServiceResourceUpdateProperties = original.ServiceKindBasicServiceResourceUpdatePropertiesServiceKindStateless
)

type ServiceLoadMetricWeight = original.ServiceLoadMetricWeight

const (
	ServiceLoadMetricWeightHigh   ServiceLoadMetricWeight = original.ServiceLoadMetricWeightHigh
	ServiceLoadMetricWeightLow    ServiceLoadMetricWeight = original.ServiceLoadMetricWeightLow
	ServiceLoadMetricWeightMedium ServiceLoadMetricWeight = original.ServiceLoadMetricWeightMedium
	ServiceLoadMetricWeightZero   ServiceLoadMetricWeight = original.ServiceLoadMetricWeightZero
)

type ServicePlacementPolicyType = original.ServicePlacementPolicyType

const (
	ServicePlacementPolicyTypeInvalid                    ServicePlacementPolicyType = original.ServicePlacementPolicyTypeInvalid
	ServicePlacementPolicyTypeInvalidDomain              ServicePlacementPolicyType = original.ServicePlacementPolicyTypeInvalidDomain
	ServicePlacementPolicyTypeNonPartiallyPlaceService   ServicePlacementPolicyType = original.ServicePlacementPolicyTypeNonPartiallyPlaceService
	ServicePlacementPolicyTypePreferredPrimaryDomain     ServicePlacementPolicyType = original.ServicePlacementPolicyTypePreferredPrimaryDomain
	ServicePlacementPolicyTypeRequiredDomain             ServicePlacementPolicyType = original.ServicePlacementPolicyTypeRequiredDomain
	ServicePlacementPolicyTypeRequiredDomainDistribution ServicePlacementPolicyType = original.ServicePlacementPolicyTypeRequiredDomainDistribution
)

type Type = original.Type

const (
	TypeServicePlacementPolicyDescription Type = original.TypeServicePlacementPolicyDescription
)

type UpgradeMode = original.UpgradeMode

const (
	UpgradeModeAutomatic UpgradeMode = original.UpgradeModeAutomatic
	UpgradeModeManual    UpgradeMode = original.UpgradeModeManual
)

type UpgradeMode1 = original.UpgradeMode1

const (
	UpgradeMode1Automatic UpgradeMode1 = original.UpgradeMode1Automatic
	UpgradeMode1Manual    UpgradeMode1 = original.UpgradeMode1Manual
)

type X509StoreName = original.X509StoreName

const (
	AddressBook          X509StoreName = original.AddressBook
	AuthRoot             X509StoreName = original.AuthRoot
	CertificateAuthority X509StoreName = original.CertificateAuthority
	Disallowed           X509StoreName = original.Disallowed
	My                   X509StoreName = original.My
	Root                 X509StoreName = original.Root
	TrustedPeople        X509StoreName = original.TrustedPeople
	TrustedPublisher     X509StoreName = original.TrustedPublisher
)

type X509StoreName1 = original.X509StoreName1

const (
	X509StoreName1AddressBook          X509StoreName1 = original.X509StoreName1AddressBook
	X509StoreName1AuthRoot             X509StoreName1 = original.X509StoreName1AuthRoot
	X509StoreName1CertificateAuthority X509StoreName1 = original.X509StoreName1CertificateAuthority
	X509StoreName1Disallowed           X509StoreName1 = original.X509StoreName1Disallowed
	X509StoreName1My                   X509StoreName1 = original.X509StoreName1My
	X509StoreName1Root                 X509StoreName1 = original.X509StoreName1Root
	X509StoreName1TrustedPeople        X509StoreName1 = original.X509StoreName1TrustedPeople
	X509StoreName1TrustedPublisher     X509StoreName1 = original.X509StoreName1TrustedPublisher
)

type ApplicationDeltaHealthPolicy = original.ApplicationDeltaHealthPolicy
type ApplicationHealthPolicy = original.ApplicationHealthPolicy
type ApplicationMetricDescription = original.ApplicationMetricDescription
type ApplicationResource = original.ApplicationResource
type ApplicationResourceList = original.ApplicationResourceList
type ApplicationResourceProperties = original.ApplicationResourceProperties
type ApplicationResourceUpdate = original.ApplicationResourceUpdate
type ApplicationResourceUpdateProperties = original.ApplicationResourceUpdateProperties
type ApplicationTypeResource = original.ApplicationTypeResource
type ApplicationTypeResourceList = original.ApplicationTypeResourceList
type ApplicationTypeResourceProperties = original.ApplicationTypeResourceProperties
type ApplicationTypeVersionResource = original.ApplicationTypeVersionResource
type ApplicationTypeVersionResourceList = original.ApplicationTypeVersionResourceList
type ApplicationTypeVersionResourceProperties = original.ApplicationTypeVersionResourceProperties
type ApplicationTypeVersionsClient = original.ApplicationTypeVersionsClient
type ApplicationTypeVersionsCreateOrUpdateFuture = original.ApplicationTypeVersionsCreateOrUpdateFuture
type ApplicationTypeVersionsDeleteFuture = original.ApplicationTypeVersionsDeleteFuture
type ApplicationTypesClient = original.ApplicationTypesClient
type ApplicationTypesDeleteFuture = original.ApplicationTypesDeleteFuture
type ApplicationUpgradePolicy = original.ApplicationUpgradePolicy
type ApplicationsClient = original.ApplicationsClient
type ApplicationsCreateOrUpdateFuture = original.ApplicationsCreateOrUpdateFuture
type ApplicationsDeleteFuture = original.ApplicationsDeleteFuture
type ApplicationsUpdateFuture = original.ApplicationsUpdateFuture
type ArmApplicationHealthPolicy = original.ArmApplicationHealthPolicy
type ArmRollingUpgradeMonitoringPolicy = original.ArmRollingUpgradeMonitoringPolicy
type ArmServiceTypeHealthPolicy = original.ArmServiceTypeHealthPolicy
type AvailableOperationDisplay = original.AvailableOperationDisplay
type AzureActiveDirectory = original.AzureActiveDirectory
type BaseClient = original.BaseClient
type BasicPartitionSchemeDescription = original.BasicPartitionSchemeDescription
type BasicServicePlacementPolicyDescription = original.BasicServicePlacementPolicyDescription
type BasicServiceResourceProperties = original.BasicServiceResourceProperties
type BasicServiceResourceUpdateProperties = original.BasicServiceResourceUpdateProperties
type CertificateDescription = original.CertificateDescription
type ClientCertificateCommonName = original.ClientCertificateCommonName
type ClientCertificateThumbprint = original.ClientCertificateThumbprint
type Cluster = original.Cluster
type ClusterCodeVersionsListResult = original.ClusterCodeVersionsListResult
type ClusterCodeVersionsResult = original.ClusterCodeVersionsResult
type ClusterHealthPolicy = original.ClusterHealthPolicy
type ClusterListResult = original.ClusterListResult
type ClusterProperties = original.ClusterProperties
type ClusterPropertiesUpdateParameters = original.ClusterPropertiesUpdateParameters
type ClusterUpdateParameters = original.ClusterUpdateParameters
type ClusterUpgradeDeltaHealthPolicy = original.ClusterUpgradeDeltaHealthPolicy
type ClusterUpgradePolicy = original.ClusterUpgradePolicy
type ClusterVersionDetails = original.ClusterVersionDetails
type ClusterVersionsClient = original.ClusterVersionsClient
type ClustersClient = original.ClustersClient
type ClustersCreateOrUpdateFuture = original.ClustersCreateOrUpdateFuture
type ClustersUpdateFuture = original.ClustersUpdateFuture
type DiagnosticsStorageAccountConfig = original.DiagnosticsStorageAccountConfig
type EndpointRangeDescription = original.EndpointRangeDescription
type ErrorModel = original.ErrorModel
type ErrorModelError = original.ErrorModelError
type NamedPartitionSchemeDescription = original.NamedPartitionSchemeDescription
type NodeTypeDescription = original.NodeTypeDescription
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationResult = original.OperationResult
type OperationsClient = original.OperationsClient
type PartitionSchemeDescription = original.PartitionSchemeDescription
type ProxyResource = original.ProxyResource
type Resource = original.Resource
type ServerCertificateCommonName = original.ServerCertificateCommonName
type ServerCertificateCommonNames = original.ServerCertificateCommonNames
type ServiceCorrelationDescription = original.ServiceCorrelationDescription
type ServiceLoadMetricDescription = original.ServiceLoadMetricDescription
type ServicePlacementPolicyDescription = original.ServicePlacementPolicyDescription
type ServiceResource = original.ServiceResource
type ServiceResourceList = original.ServiceResourceList
type ServiceResourceProperties = original.ServiceResourceProperties
type ServiceResourcePropertiesBase = original.ServiceResourcePropertiesBase
type ServiceResourceUpdate = original.ServiceResourceUpdate
type ServiceResourceUpdateProperties = original.ServiceResourceUpdateProperties
type ServiceTypeDeltaHealthPolicy = original.ServiceTypeDeltaHealthPolicy
type ServiceTypeHealthPolicy = original.ServiceTypeHealthPolicy
type ServicesClient = original.ServicesClient
type ServicesCreateOrUpdateFuture = original.ServicesCreateOrUpdateFuture
type ServicesDeleteFuture = original.ServicesDeleteFuture
type ServicesUpdateFuture = original.ServicesUpdateFuture
type SettingsParameterDescription = original.SettingsParameterDescription
type SettingsSectionDescription = original.SettingsSectionDescription
type SingletonPartitionSchemeDescription = original.SingletonPartitionSchemeDescription
type StatefulServiceProperties = original.StatefulServiceProperties
type StatefulServiceUpdateProperties = original.StatefulServiceUpdateProperties
type StatelessServiceProperties = original.StatelessServiceProperties
type StatelessServiceUpdateProperties = original.StatelessServiceUpdateProperties
type UniformInt64RangePartitionSchemeDescription = original.UniformInt64RangePartitionSchemeDescription

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewApplicationTypeVersionsClient(subscriptionID string) ApplicationTypeVersionsClient {
	return original.NewApplicationTypeVersionsClient(subscriptionID)
}
func NewApplicationTypeVersionsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationTypeVersionsClient {
	return original.NewApplicationTypeVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationTypesClient(subscriptionID string) ApplicationTypesClient {
	return original.NewApplicationTypesClient(subscriptionID)
}
func NewApplicationTypesClientWithBaseURI(baseURI string, subscriptionID string) ApplicationTypesClient {
	return original.NewApplicationTypesClientWithBaseURI(baseURI, subscriptionID)
}
func NewApplicationsClient(subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClient(subscriptionID)
}
func NewApplicationsClientWithBaseURI(baseURI string, subscriptionID string) ApplicationsClient {
	return original.NewApplicationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClusterVersionsClient(subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClient(subscriptionID)
}
func NewClusterVersionsClientWithBaseURI(baseURI string, subscriptionID string) ClusterVersionsClient {
	return original.NewClusterVersionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewClustersClient(subscriptionID string) ClustersClient {
	return original.NewClustersClient(subscriptionID)
}
func NewClustersClientWithBaseURI(baseURI string, subscriptionID string) ClustersClient {
	return original.NewClustersClientWithBaseURI(baseURI, subscriptionID)
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
func NewServicesClient(subscriptionID string) ServicesClient {
	return original.NewServicesClient(subscriptionID)
}
func NewServicesClientWithBaseURI(baseURI string, subscriptionID string) ServicesClient {
	return original.NewServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleArmServicePackageActivationModeValues() []ArmServicePackageActivationMode {
	return original.PossibleArmServicePackageActivationModeValues()
}
func PossibleArmUpgradeFailureActionValues() []ArmUpgradeFailureAction {
	return original.PossibleArmUpgradeFailureActionValues()
}
func PossibleClusterStateValues() []ClusterState {
	return original.PossibleClusterStateValues()
}
func PossibleDurabilityLevelValues() []DurabilityLevel {
	return original.PossibleDurabilityLevelValues()
}
func PossibleEnvironmentValues() []Environment {
	return original.PossibleEnvironmentValues()
}
func PossibleMoveCostValues() []MoveCost {
	return original.PossibleMoveCostValues()
}
func PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues() []PartitionSchemeBasicPartitionSchemeDescription {
	return original.PossiblePartitionSchemeBasicPartitionSchemeDescriptionValues()
}
func PossiblePartitionSchemeValues() []PartitionScheme {
	return original.PossiblePartitionSchemeValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleReliabilityLevel1Values() []ReliabilityLevel1 {
	return original.PossibleReliabilityLevel1Values()
}
func PossibleReliabilityLevelValues() []ReliabilityLevel {
	return original.PossibleReliabilityLevelValues()
}
func PossibleServiceCorrelationSchemeValues() []ServiceCorrelationScheme {
	return original.PossibleServiceCorrelationSchemeValues()
}
func PossibleServiceKindBasicServiceResourcePropertiesValues() []ServiceKindBasicServiceResourceProperties {
	return original.PossibleServiceKindBasicServiceResourcePropertiesValues()
}
func PossibleServiceKindBasicServiceResourceUpdatePropertiesValues() []ServiceKindBasicServiceResourceUpdateProperties {
	return original.PossibleServiceKindBasicServiceResourceUpdatePropertiesValues()
}
func PossibleServiceKindValues() []ServiceKind {
	return original.PossibleServiceKindValues()
}
func PossibleServiceLoadMetricWeightValues() []ServiceLoadMetricWeight {
	return original.PossibleServiceLoadMetricWeightValues()
}
func PossibleServicePlacementPolicyTypeValues() []ServicePlacementPolicyType {
	return original.PossibleServicePlacementPolicyTypeValues()
}
func PossibleTypeValues() []Type {
	return original.PossibleTypeValues()
}
func PossibleUpgradeMode1Values() []UpgradeMode1 {
	return original.PossibleUpgradeMode1Values()
}
func PossibleUpgradeModeValues() []UpgradeMode {
	return original.PossibleUpgradeModeValues()
}
func PossibleX509StoreName1Values() []X509StoreName1 {
	return original.PossibleX509StoreName1Values()
}
func PossibleX509StoreNameValues() []X509StoreName {
	return original.PossibleX509StoreNameValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
