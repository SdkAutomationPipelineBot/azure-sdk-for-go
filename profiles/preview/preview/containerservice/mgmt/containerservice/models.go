//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package containerservice

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/containerservice/mgmt/2019-10-27-preview/containerservice"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AgentPoolType = original.AgentPoolType

const (
	AvailabilitySet         AgentPoolType = original.AvailabilitySet
	VirtualMachineScaleSets AgentPoolType = original.VirtualMachineScaleSets
)

type Kind = original.Kind

const (
	KindAADIdentityProvider                         Kind = original.KindAADIdentityProvider
	KindOpenShiftManagedClusterBaseIdentityProvider Kind = original.KindOpenShiftManagedClusterBaseIdentityProvider
)

type LoadBalancerSku = original.LoadBalancerSku

const (
	Basic    LoadBalancerSku = original.Basic
	Standard LoadBalancerSku = original.Standard
)

type NetworkPlugin = original.NetworkPlugin

const (
	Azure   NetworkPlugin = original.Azure
	Kubenet NetworkPlugin = original.Kubenet
)

type NetworkPolicy = original.NetworkPolicy

const (
	NetworkPolicyAzure  NetworkPolicy = original.NetworkPolicyAzure
	NetworkPolicyCalico NetworkPolicy = original.NetworkPolicyCalico
)

type OSType = original.OSType

const (
	Linux   OSType = original.Linux
	Windows OSType = original.Windows
)

type OpenShiftAgentPoolProfileRole = original.OpenShiftAgentPoolProfileRole

const (
	Compute OpenShiftAgentPoolProfileRole = original.Compute
	Infra   OpenShiftAgentPoolProfileRole = original.Infra
)

type OpenShiftContainerServiceVMSize = original.OpenShiftContainerServiceVMSize

const (
	StandardD16sV3 OpenShiftContainerServiceVMSize = original.StandardD16sV3
	StandardD2sV3  OpenShiftContainerServiceVMSize = original.StandardD2sV3
	StandardD32sV3 OpenShiftContainerServiceVMSize = original.StandardD32sV3
	StandardD4sV3  OpenShiftContainerServiceVMSize = original.StandardD4sV3
	StandardD64sV3 OpenShiftContainerServiceVMSize = original.StandardD64sV3
	StandardD8sV3  OpenShiftContainerServiceVMSize = original.StandardD8sV3
	StandardDS12V2 OpenShiftContainerServiceVMSize = original.StandardDS12V2
	StandardDS13V2 OpenShiftContainerServiceVMSize = original.StandardDS13V2
	StandardDS14V2 OpenShiftContainerServiceVMSize = original.StandardDS14V2
	StandardDS15V2 OpenShiftContainerServiceVMSize = original.StandardDS15V2
	StandardDS4V2  OpenShiftContainerServiceVMSize = original.StandardDS4V2
	StandardDS5V2  OpenShiftContainerServiceVMSize = original.StandardDS5V2
	StandardE16sV3 OpenShiftContainerServiceVMSize = original.StandardE16sV3
	StandardE20sV3 OpenShiftContainerServiceVMSize = original.StandardE20sV3
	StandardE32sV3 OpenShiftContainerServiceVMSize = original.StandardE32sV3
	StandardE4sV3  OpenShiftContainerServiceVMSize = original.StandardE4sV3
	StandardE64sV3 OpenShiftContainerServiceVMSize = original.StandardE64sV3
	StandardE8sV3  OpenShiftContainerServiceVMSize = original.StandardE8sV3
	StandardF16s   OpenShiftContainerServiceVMSize = original.StandardF16s
	StandardF16sV2 OpenShiftContainerServiceVMSize = original.StandardF16sV2
	StandardF32sV2 OpenShiftContainerServiceVMSize = original.StandardF32sV2
	StandardF64sV2 OpenShiftContainerServiceVMSize = original.StandardF64sV2
	StandardF72sV2 OpenShiftContainerServiceVMSize = original.StandardF72sV2
	StandardF8s    OpenShiftContainerServiceVMSize = original.StandardF8s
	StandardF8sV2  OpenShiftContainerServiceVMSize = original.StandardF8sV2
	StandardGS2    OpenShiftContainerServiceVMSize = original.StandardGS2
	StandardGS3    OpenShiftContainerServiceVMSize = original.StandardGS3
	StandardGS4    OpenShiftContainerServiceVMSize = original.StandardGS4
	StandardGS5    OpenShiftContainerServiceVMSize = original.StandardGS5
	StandardL16s   OpenShiftContainerServiceVMSize = original.StandardL16s
	StandardL32s   OpenShiftContainerServiceVMSize = original.StandardL32s
	StandardL4s    OpenShiftContainerServiceVMSize = original.StandardL4s
	StandardL8s    OpenShiftContainerServiceVMSize = original.StandardL8s
)

type OrchestratorTypes = original.OrchestratorTypes

const (
	Custom     OrchestratorTypes = original.Custom
	DCOS       OrchestratorTypes = original.DCOS
	DockerCE   OrchestratorTypes = original.DockerCE
	Kubernetes OrchestratorTypes = original.Kubernetes
	Swarm      OrchestratorTypes = original.Swarm
)

type OutboundType = original.OutboundType

const (
	LoadBalancer       OutboundType = original.LoadBalancer
	UserDefinedRouting OutboundType = original.UserDefinedRouting
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	None           ResourceIdentityType = original.None
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type ScaleSetEvictionPolicy = original.ScaleSetEvictionPolicy

const (
	Deallocate ScaleSetEvictionPolicy = original.Deallocate
	Delete     ScaleSetEvictionPolicy = original.Delete
)

type ScaleSetPriority = original.ScaleSetPriority

const (
	Low     ScaleSetPriority = original.Low
	Regular ScaleSetPriority = original.Regular
)

type StorageProfileTypes = original.StorageProfileTypes

const (
	ManagedDisks   StorageProfileTypes = original.ManagedDisks
	StorageAccount StorageProfileTypes = original.StorageAccount
)

type VMSizeTypes = original.VMSizeTypes

const (
	VMSizeTypesStandardA1          VMSizeTypes = original.VMSizeTypesStandardA1
	VMSizeTypesStandardA10         VMSizeTypes = original.VMSizeTypesStandardA10
	VMSizeTypesStandardA11         VMSizeTypes = original.VMSizeTypesStandardA11
	VMSizeTypesStandardA1V2        VMSizeTypes = original.VMSizeTypesStandardA1V2
	VMSizeTypesStandardA2          VMSizeTypes = original.VMSizeTypesStandardA2
	VMSizeTypesStandardA2mV2       VMSizeTypes = original.VMSizeTypesStandardA2mV2
	VMSizeTypesStandardA2V2        VMSizeTypes = original.VMSizeTypesStandardA2V2
	VMSizeTypesStandardA3          VMSizeTypes = original.VMSizeTypesStandardA3
	VMSizeTypesStandardA4          VMSizeTypes = original.VMSizeTypesStandardA4
	VMSizeTypesStandardA4mV2       VMSizeTypes = original.VMSizeTypesStandardA4mV2
	VMSizeTypesStandardA4V2        VMSizeTypes = original.VMSizeTypesStandardA4V2
	VMSizeTypesStandardA5          VMSizeTypes = original.VMSizeTypesStandardA5
	VMSizeTypesStandardA6          VMSizeTypes = original.VMSizeTypesStandardA6
	VMSizeTypesStandardA7          VMSizeTypes = original.VMSizeTypesStandardA7
	VMSizeTypesStandardA8          VMSizeTypes = original.VMSizeTypesStandardA8
	VMSizeTypesStandardA8mV2       VMSizeTypes = original.VMSizeTypesStandardA8mV2
	VMSizeTypesStandardA8V2        VMSizeTypes = original.VMSizeTypesStandardA8V2
	VMSizeTypesStandardA9          VMSizeTypes = original.VMSizeTypesStandardA9
	VMSizeTypesStandardB2ms        VMSizeTypes = original.VMSizeTypesStandardB2ms
	VMSizeTypesStandardB2s         VMSizeTypes = original.VMSizeTypesStandardB2s
	VMSizeTypesStandardB4ms        VMSizeTypes = original.VMSizeTypesStandardB4ms
	VMSizeTypesStandardB8ms        VMSizeTypes = original.VMSizeTypesStandardB8ms
	VMSizeTypesStandardD1          VMSizeTypes = original.VMSizeTypesStandardD1
	VMSizeTypesStandardD11         VMSizeTypes = original.VMSizeTypesStandardD11
	VMSizeTypesStandardD11V2       VMSizeTypes = original.VMSizeTypesStandardD11V2
	VMSizeTypesStandardD11V2Promo  VMSizeTypes = original.VMSizeTypesStandardD11V2Promo
	VMSizeTypesStandardD12         VMSizeTypes = original.VMSizeTypesStandardD12
	VMSizeTypesStandardD12V2       VMSizeTypes = original.VMSizeTypesStandardD12V2
	VMSizeTypesStandardD12V2Promo  VMSizeTypes = original.VMSizeTypesStandardD12V2Promo
	VMSizeTypesStandardD13         VMSizeTypes = original.VMSizeTypesStandardD13
	VMSizeTypesStandardD13V2       VMSizeTypes = original.VMSizeTypesStandardD13V2
	VMSizeTypesStandardD13V2Promo  VMSizeTypes = original.VMSizeTypesStandardD13V2Promo
	VMSizeTypesStandardD14         VMSizeTypes = original.VMSizeTypesStandardD14
	VMSizeTypesStandardD14V2       VMSizeTypes = original.VMSizeTypesStandardD14V2
	VMSizeTypesStandardD14V2Promo  VMSizeTypes = original.VMSizeTypesStandardD14V2Promo
	VMSizeTypesStandardD15V2       VMSizeTypes = original.VMSizeTypesStandardD15V2
	VMSizeTypesStandardD16sV3      VMSizeTypes = original.VMSizeTypesStandardD16sV3
	VMSizeTypesStandardD16V3       VMSizeTypes = original.VMSizeTypesStandardD16V3
	VMSizeTypesStandardD1V2        VMSizeTypes = original.VMSizeTypesStandardD1V2
	VMSizeTypesStandardD2          VMSizeTypes = original.VMSizeTypesStandardD2
	VMSizeTypesStandardD2sV3       VMSizeTypes = original.VMSizeTypesStandardD2sV3
	VMSizeTypesStandardD2V2        VMSizeTypes = original.VMSizeTypesStandardD2V2
	VMSizeTypesStandardD2V2Promo   VMSizeTypes = original.VMSizeTypesStandardD2V2Promo
	VMSizeTypesStandardD2V3        VMSizeTypes = original.VMSizeTypesStandardD2V3
	VMSizeTypesStandardD3          VMSizeTypes = original.VMSizeTypesStandardD3
	VMSizeTypesStandardD32sV3      VMSizeTypes = original.VMSizeTypesStandardD32sV3
	VMSizeTypesStandardD32V3       VMSizeTypes = original.VMSizeTypesStandardD32V3
	VMSizeTypesStandardD3V2        VMSizeTypes = original.VMSizeTypesStandardD3V2
	VMSizeTypesStandardD3V2Promo   VMSizeTypes = original.VMSizeTypesStandardD3V2Promo
	VMSizeTypesStandardD4          VMSizeTypes = original.VMSizeTypesStandardD4
	VMSizeTypesStandardD4sV3       VMSizeTypes = original.VMSizeTypesStandardD4sV3
	VMSizeTypesStandardD4V2        VMSizeTypes = original.VMSizeTypesStandardD4V2
	VMSizeTypesStandardD4V2Promo   VMSizeTypes = original.VMSizeTypesStandardD4V2Promo
	VMSizeTypesStandardD4V3        VMSizeTypes = original.VMSizeTypesStandardD4V3
	VMSizeTypesStandardD5V2        VMSizeTypes = original.VMSizeTypesStandardD5V2
	VMSizeTypesStandardD5V2Promo   VMSizeTypes = original.VMSizeTypesStandardD5V2Promo
	VMSizeTypesStandardD64sV3      VMSizeTypes = original.VMSizeTypesStandardD64sV3
	VMSizeTypesStandardD64V3       VMSizeTypes = original.VMSizeTypesStandardD64V3
	VMSizeTypesStandardD8sV3       VMSizeTypes = original.VMSizeTypesStandardD8sV3
	VMSizeTypesStandardD8V3        VMSizeTypes = original.VMSizeTypesStandardD8V3
	VMSizeTypesStandardDS1         VMSizeTypes = original.VMSizeTypesStandardDS1
	VMSizeTypesStandardDS11        VMSizeTypes = original.VMSizeTypesStandardDS11
	VMSizeTypesStandardDS11V2      VMSizeTypes = original.VMSizeTypesStandardDS11V2
	VMSizeTypesStandardDS11V2Promo VMSizeTypes = original.VMSizeTypesStandardDS11V2Promo
	VMSizeTypesStandardDS12        VMSizeTypes = original.VMSizeTypesStandardDS12
	VMSizeTypesStandardDS12V2      VMSizeTypes = original.VMSizeTypesStandardDS12V2
	VMSizeTypesStandardDS12V2Promo VMSizeTypes = original.VMSizeTypesStandardDS12V2Promo
	VMSizeTypesStandardDS13        VMSizeTypes = original.VMSizeTypesStandardDS13
	VMSizeTypesStandardDS132V2     VMSizeTypes = original.VMSizeTypesStandardDS132V2
	VMSizeTypesStandardDS134V2     VMSizeTypes = original.VMSizeTypesStandardDS134V2
	VMSizeTypesStandardDS13V2      VMSizeTypes = original.VMSizeTypesStandardDS13V2
	VMSizeTypesStandardDS13V2Promo VMSizeTypes = original.VMSizeTypesStandardDS13V2Promo
	VMSizeTypesStandardDS14        VMSizeTypes = original.VMSizeTypesStandardDS14
	VMSizeTypesStandardDS144V2     VMSizeTypes = original.VMSizeTypesStandardDS144V2
	VMSizeTypesStandardDS148V2     VMSizeTypes = original.VMSizeTypesStandardDS148V2
	VMSizeTypesStandardDS14V2      VMSizeTypes = original.VMSizeTypesStandardDS14V2
	VMSizeTypesStandardDS14V2Promo VMSizeTypes = original.VMSizeTypesStandardDS14V2Promo
	VMSizeTypesStandardDS15V2      VMSizeTypes = original.VMSizeTypesStandardDS15V2
	VMSizeTypesStandardDS1V2       VMSizeTypes = original.VMSizeTypesStandardDS1V2
	VMSizeTypesStandardDS2         VMSizeTypes = original.VMSizeTypesStandardDS2
	VMSizeTypesStandardDS2V2       VMSizeTypes = original.VMSizeTypesStandardDS2V2
	VMSizeTypesStandardDS2V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS2V2Promo
	VMSizeTypesStandardDS3         VMSizeTypes = original.VMSizeTypesStandardDS3
	VMSizeTypesStandardDS3V2       VMSizeTypes = original.VMSizeTypesStandardDS3V2
	VMSizeTypesStandardDS3V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS3V2Promo
	VMSizeTypesStandardDS4         VMSizeTypes = original.VMSizeTypesStandardDS4
	VMSizeTypesStandardDS4V2       VMSizeTypes = original.VMSizeTypesStandardDS4V2
	VMSizeTypesStandardDS4V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS4V2Promo
	VMSizeTypesStandardDS5V2       VMSizeTypes = original.VMSizeTypesStandardDS5V2
	VMSizeTypesStandardDS5V2Promo  VMSizeTypes = original.VMSizeTypesStandardDS5V2Promo
	VMSizeTypesStandardE16sV3      VMSizeTypes = original.VMSizeTypesStandardE16sV3
	VMSizeTypesStandardE16V3       VMSizeTypes = original.VMSizeTypesStandardE16V3
	VMSizeTypesStandardE2sV3       VMSizeTypes = original.VMSizeTypesStandardE2sV3
	VMSizeTypesStandardE2V3        VMSizeTypes = original.VMSizeTypesStandardE2V3
	VMSizeTypesStandardE3216sV3    VMSizeTypes = original.VMSizeTypesStandardE3216sV3
	VMSizeTypesStandardE328sV3     VMSizeTypes = original.VMSizeTypesStandardE328sV3
	VMSizeTypesStandardE32sV3      VMSizeTypes = original.VMSizeTypesStandardE32sV3
	VMSizeTypesStandardE32V3       VMSizeTypes = original.VMSizeTypesStandardE32V3
	VMSizeTypesStandardE4sV3       VMSizeTypes = original.VMSizeTypesStandardE4sV3
	VMSizeTypesStandardE4V3        VMSizeTypes = original.VMSizeTypesStandardE4V3
	VMSizeTypesStandardE6416sV3    VMSizeTypes = original.VMSizeTypesStandardE6416sV3
	VMSizeTypesStandardE6432sV3    VMSizeTypes = original.VMSizeTypesStandardE6432sV3
	VMSizeTypesStandardE64sV3      VMSizeTypes = original.VMSizeTypesStandardE64sV3
	VMSizeTypesStandardE64V3       VMSizeTypes = original.VMSizeTypesStandardE64V3
	VMSizeTypesStandardE8sV3       VMSizeTypes = original.VMSizeTypesStandardE8sV3
	VMSizeTypesStandardE8V3        VMSizeTypes = original.VMSizeTypesStandardE8V3
	VMSizeTypesStandardF1          VMSizeTypes = original.VMSizeTypesStandardF1
	VMSizeTypesStandardF16         VMSizeTypes = original.VMSizeTypesStandardF16
	VMSizeTypesStandardF16s        VMSizeTypes = original.VMSizeTypesStandardF16s
	VMSizeTypesStandardF16sV2      VMSizeTypes = original.VMSizeTypesStandardF16sV2
	VMSizeTypesStandardF1s         VMSizeTypes = original.VMSizeTypesStandardF1s
	VMSizeTypesStandardF2          VMSizeTypes = original.VMSizeTypesStandardF2
	VMSizeTypesStandardF2s         VMSizeTypes = original.VMSizeTypesStandardF2s
	VMSizeTypesStandardF2sV2       VMSizeTypes = original.VMSizeTypesStandardF2sV2
	VMSizeTypesStandardF32sV2      VMSizeTypes = original.VMSizeTypesStandardF32sV2
	VMSizeTypesStandardF4          VMSizeTypes = original.VMSizeTypesStandardF4
	VMSizeTypesStandardF4s         VMSizeTypes = original.VMSizeTypesStandardF4s
	VMSizeTypesStandardF4sV2       VMSizeTypes = original.VMSizeTypesStandardF4sV2
	VMSizeTypesStandardF64sV2      VMSizeTypes = original.VMSizeTypesStandardF64sV2
	VMSizeTypesStandardF72sV2      VMSizeTypes = original.VMSizeTypesStandardF72sV2
	VMSizeTypesStandardF8          VMSizeTypes = original.VMSizeTypesStandardF8
	VMSizeTypesStandardF8s         VMSizeTypes = original.VMSizeTypesStandardF8s
	VMSizeTypesStandardF8sV2       VMSizeTypes = original.VMSizeTypesStandardF8sV2
	VMSizeTypesStandardG1          VMSizeTypes = original.VMSizeTypesStandardG1
	VMSizeTypesStandardG2          VMSizeTypes = original.VMSizeTypesStandardG2
	VMSizeTypesStandardG3          VMSizeTypes = original.VMSizeTypesStandardG3
	VMSizeTypesStandardG4          VMSizeTypes = original.VMSizeTypesStandardG4
	VMSizeTypesStandardG5          VMSizeTypes = original.VMSizeTypesStandardG5
	VMSizeTypesStandardGS1         VMSizeTypes = original.VMSizeTypesStandardGS1
	VMSizeTypesStandardGS2         VMSizeTypes = original.VMSizeTypesStandardGS2
	VMSizeTypesStandardGS3         VMSizeTypes = original.VMSizeTypesStandardGS3
	VMSizeTypesStandardGS4         VMSizeTypes = original.VMSizeTypesStandardGS4
	VMSizeTypesStandardGS44        VMSizeTypes = original.VMSizeTypesStandardGS44
	VMSizeTypesStandardGS48        VMSizeTypes = original.VMSizeTypesStandardGS48
	VMSizeTypesStandardGS5         VMSizeTypes = original.VMSizeTypesStandardGS5
	VMSizeTypesStandardGS516       VMSizeTypes = original.VMSizeTypesStandardGS516
	VMSizeTypesStandardGS58        VMSizeTypes = original.VMSizeTypesStandardGS58
	VMSizeTypesStandardH16         VMSizeTypes = original.VMSizeTypesStandardH16
	VMSizeTypesStandardH16m        VMSizeTypes = original.VMSizeTypesStandardH16m
	VMSizeTypesStandardH16mr       VMSizeTypes = original.VMSizeTypesStandardH16mr
	VMSizeTypesStandardH16r        VMSizeTypes = original.VMSizeTypesStandardH16r
	VMSizeTypesStandardH8          VMSizeTypes = original.VMSizeTypesStandardH8
	VMSizeTypesStandardH8m         VMSizeTypes = original.VMSizeTypesStandardH8m
	VMSizeTypesStandardL16s        VMSizeTypes = original.VMSizeTypesStandardL16s
	VMSizeTypesStandardL32s        VMSizeTypes = original.VMSizeTypesStandardL32s
	VMSizeTypesStandardL4s         VMSizeTypes = original.VMSizeTypesStandardL4s
	VMSizeTypesStandardL8s         VMSizeTypes = original.VMSizeTypesStandardL8s
	VMSizeTypesStandardM12832ms    VMSizeTypes = original.VMSizeTypesStandardM12832ms
	VMSizeTypesStandardM12864ms    VMSizeTypes = original.VMSizeTypesStandardM12864ms
	VMSizeTypesStandardM128ms      VMSizeTypes = original.VMSizeTypesStandardM128ms
	VMSizeTypesStandardM128s       VMSizeTypes = original.VMSizeTypesStandardM128s
	VMSizeTypesStandardM6416ms     VMSizeTypes = original.VMSizeTypesStandardM6416ms
	VMSizeTypesStandardM6432ms     VMSizeTypes = original.VMSizeTypesStandardM6432ms
	VMSizeTypesStandardM64ms       VMSizeTypes = original.VMSizeTypesStandardM64ms
	VMSizeTypesStandardM64s        VMSizeTypes = original.VMSizeTypesStandardM64s
	VMSizeTypesStandardNC12        VMSizeTypes = original.VMSizeTypesStandardNC12
	VMSizeTypesStandardNC12sV2     VMSizeTypes = original.VMSizeTypesStandardNC12sV2
	VMSizeTypesStandardNC12sV3     VMSizeTypes = original.VMSizeTypesStandardNC12sV3
	VMSizeTypesStandardNC24        VMSizeTypes = original.VMSizeTypesStandardNC24
	VMSizeTypesStandardNC24r       VMSizeTypes = original.VMSizeTypesStandardNC24r
	VMSizeTypesStandardNC24rsV2    VMSizeTypes = original.VMSizeTypesStandardNC24rsV2
	VMSizeTypesStandardNC24rsV3    VMSizeTypes = original.VMSizeTypesStandardNC24rsV3
	VMSizeTypesStandardNC24sV2     VMSizeTypes = original.VMSizeTypesStandardNC24sV2
	VMSizeTypesStandardNC24sV3     VMSizeTypes = original.VMSizeTypesStandardNC24sV3
	VMSizeTypesStandardNC6         VMSizeTypes = original.VMSizeTypesStandardNC6
	VMSizeTypesStandardNC6sV2      VMSizeTypes = original.VMSizeTypesStandardNC6sV2
	VMSizeTypesStandardNC6sV3      VMSizeTypes = original.VMSizeTypesStandardNC6sV3
	VMSizeTypesStandardND12s       VMSizeTypes = original.VMSizeTypesStandardND12s
	VMSizeTypesStandardND24rs      VMSizeTypes = original.VMSizeTypesStandardND24rs
	VMSizeTypesStandardND24s       VMSizeTypes = original.VMSizeTypesStandardND24s
	VMSizeTypesStandardND6s        VMSizeTypes = original.VMSizeTypesStandardND6s
	VMSizeTypesStandardNV12        VMSizeTypes = original.VMSizeTypesStandardNV12
	VMSizeTypesStandardNV24        VMSizeTypes = original.VMSizeTypesStandardNV24
	VMSizeTypesStandardNV6         VMSizeTypes = original.VMSizeTypesStandardNV6
)

type AccessProfile = original.AccessProfile
type AgentPool = original.AgentPool
type AgentPoolAvailableVersions = original.AgentPoolAvailableVersions
type AgentPoolAvailableVersionsProperties = original.AgentPoolAvailableVersionsProperties
type AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem = original.AgentPoolAvailableVersionsPropertiesAgentPoolVersionsItem
type AgentPoolListResult = original.AgentPoolListResult
type AgentPoolListResultIterator = original.AgentPoolListResultIterator
type AgentPoolListResultPage = original.AgentPoolListResultPage
type AgentPoolProfile = original.AgentPoolProfile
type AgentPoolUpgradeProfile = original.AgentPoolUpgradeProfile
type AgentPoolUpgradeProfileProperties = original.AgentPoolUpgradeProfileProperties
type AgentPoolUpgradeProfilePropertiesUpgradesItem = original.AgentPoolUpgradeProfilePropertiesUpgradesItem
type AgentPoolsClient = original.AgentPoolsClient
type AgentPoolsCreateOrUpdateFuture = original.AgentPoolsCreateOrUpdateFuture
type AgentPoolsDeleteFuture = original.AgentPoolsDeleteFuture
type BaseClient = original.BaseClient
type BasicOpenShiftManagedClusterBaseIdentityProvider = original.BasicOpenShiftManagedClusterBaseIdentityProvider
type CloudError = original.CloudError
type CloudErrorBody = original.CloudErrorBody
type ContainerService = original.ContainerService
type ContainerServicesClient = original.ContainerServicesClient
type ContainerServicesCreateOrUpdateFutureType = original.ContainerServicesCreateOrUpdateFutureType
type ContainerServicesDeleteFutureType = original.ContainerServicesDeleteFutureType
type CredentialResult = original.CredentialResult
type CredentialResults = original.CredentialResults
type CustomProfile = original.CustomProfile
type DiagnosticsProfile = original.DiagnosticsProfile
type KeyVaultSecretRef = original.KeyVaultSecretRef
type LinuxProfile = original.LinuxProfile
type ListResult = original.ListResult
type ListResultIterator = original.ListResultIterator
type ListResultPage = original.ListResultPage
type ManagedCluster = original.ManagedCluster
type ManagedClusterAADProfile = original.ManagedClusterAADProfile
type ManagedClusterAPIServerAccessProfile = original.ManagedClusterAPIServerAccessProfile
type ManagedClusterAccessProfile = original.ManagedClusterAccessProfile
type ManagedClusterAddonProfile = original.ManagedClusterAddonProfile
type ManagedClusterAddonProfileIdentity = original.ManagedClusterAddonProfileIdentity
type ManagedClusterAgentPoolProfile = original.ManagedClusterAgentPoolProfile
type ManagedClusterAgentPoolProfileProperties = original.ManagedClusterAgentPoolProfileProperties
type ManagedClusterIdentity = original.ManagedClusterIdentity
type ManagedClusterListResult = original.ManagedClusterListResult
type ManagedClusterListResultIterator = original.ManagedClusterListResultIterator
type ManagedClusterListResultPage = original.ManagedClusterListResultPage
type ManagedClusterLoadBalancerProfile = original.ManagedClusterLoadBalancerProfile
type ManagedClusterLoadBalancerProfileManagedOutboundIPs = original.ManagedClusterLoadBalancerProfileManagedOutboundIPs
type ManagedClusterLoadBalancerProfileOutboundIPPrefixes = original.ManagedClusterLoadBalancerProfileOutboundIPPrefixes
type ManagedClusterLoadBalancerProfileOutboundIPs = original.ManagedClusterLoadBalancerProfileOutboundIPs
type ManagedClusterPoolUpgradeProfile = original.ManagedClusterPoolUpgradeProfile
type ManagedClusterPoolUpgradeProfileUpgradesItem = original.ManagedClusterPoolUpgradeProfileUpgradesItem
type ManagedClusterProperties = original.ManagedClusterProperties
type ManagedClusterPropertiesIdentityProfileValue = original.ManagedClusterPropertiesIdentityProfileValue
type ManagedClusterServicePrincipalProfile = original.ManagedClusterServicePrincipalProfile
type ManagedClusterUpgradeProfile = original.ManagedClusterUpgradeProfile
type ManagedClusterUpgradeProfileProperties = original.ManagedClusterUpgradeProfileProperties
type ManagedClusterWindowsProfile = original.ManagedClusterWindowsProfile
type ManagedClustersClient = original.ManagedClustersClient
type ManagedClustersCreateOrUpdateFuture = original.ManagedClustersCreateOrUpdateFuture
type ManagedClustersDeleteFuture = original.ManagedClustersDeleteFuture
type ManagedClustersResetAADProfileFuture = original.ManagedClustersResetAADProfileFuture
type ManagedClustersResetServicePrincipalProfileFuture = original.ManagedClustersResetServicePrincipalProfileFuture
type ManagedClustersRotateClusterCertificatesFuture = original.ManagedClustersRotateClusterCertificatesFuture
type ManagedClustersUpdateTagsFuture = original.ManagedClustersUpdateTagsFuture
type MasterProfile = original.MasterProfile
type NetworkProfile = original.NetworkProfile
type NetworkProfileType = original.NetworkProfileType
type OpenShiftAPIProperties = original.OpenShiftAPIProperties
type OpenShiftManagedCluster = original.OpenShiftManagedCluster
type OpenShiftManagedClusterAADIdentityProvider = original.OpenShiftManagedClusterAADIdentityProvider
type OpenShiftManagedClusterAgentPoolProfile = original.OpenShiftManagedClusterAgentPoolProfile
type OpenShiftManagedClusterAuthProfile = original.OpenShiftManagedClusterAuthProfile
type OpenShiftManagedClusterBaseIdentityProvider = original.OpenShiftManagedClusterBaseIdentityProvider
type OpenShiftManagedClusterIdentityProvider = original.OpenShiftManagedClusterIdentityProvider
type OpenShiftManagedClusterListResult = original.OpenShiftManagedClusterListResult
type OpenShiftManagedClusterListResultIterator = original.OpenShiftManagedClusterListResultIterator
type OpenShiftManagedClusterListResultPage = original.OpenShiftManagedClusterListResultPage
type OpenShiftManagedClusterMasterPoolProfile = original.OpenShiftManagedClusterMasterPoolProfile
type OpenShiftManagedClusterMonitorProfile = original.OpenShiftManagedClusterMonitorProfile
type OpenShiftManagedClusterProperties = original.OpenShiftManagedClusterProperties
type OpenShiftManagedClustersClient = original.OpenShiftManagedClustersClient
type OpenShiftManagedClustersCreateOrUpdateFuture = original.OpenShiftManagedClustersCreateOrUpdateFuture
type OpenShiftManagedClustersDeleteFuture = original.OpenShiftManagedClustersDeleteFuture
type OpenShiftManagedClustersUpdateTagsFuture = original.OpenShiftManagedClustersUpdateTagsFuture
type OpenShiftRouterProfile = original.OpenShiftRouterProfile
type OperationListResult = original.OperationListResult
type OperationValue = original.OperationValue
type OperationValueDisplay = original.OperationValueDisplay
type OperationsClient = original.OperationsClient
type OrchestratorProfile = original.OrchestratorProfile
type OrchestratorProfileType = original.OrchestratorProfileType
type OrchestratorVersionProfile = original.OrchestratorVersionProfile
type OrchestratorVersionProfileListResult = original.OrchestratorVersionProfileListResult
type OrchestratorVersionProfileProperties = original.OrchestratorVersionProfileProperties
type Properties = original.Properties
type PurchasePlan = original.PurchasePlan
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type SSHConfiguration = original.SSHConfiguration
type SSHPublicKey = original.SSHPublicKey
type ServicePrincipalProfile = original.ServicePrincipalProfile
type SubResource = original.SubResource
type TagsObject = original.TagsObject
type UserAssignedIdentity = original.UserAssignedIdentity
type VMDiagnostics = original.VMDiagnostics
type WindowsProfile = original.WindowsProfile

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewAgentPoolListResultIterator(page AgentPoolListResultPage) AgentPoolListResultIterator {
	return original.NewAgentPoolListResultIterator(page)
}
func NewAgentPoolListResultPage(cur AgentPoolListResult, getNextPage func(context.Context, AgentPoolListResult) (AgentPoolListResult, error)) AgentPoolListResultPage {
	return original.NewAgentPoolListResultPage(cur, getNextPage)
}
func NewAgentPoolsClient(subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClient(subscriptionID)
}
func NewAgentPoolsClientWithBaseURI(baseURI string, subscriptionID string) AgentPoolsClient {
	return original.NewAgentPoolsClientWithBaseURI(baseURI, subscriptionID)
}
func NewContainerServicesClient(subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClient(subscriptionID)
}
func NewContainerServicesClientWithBaseURI(baseURI string, subscriptionID string) ContainerServicesClient {
	return original.NewContainerServicesClientWithBaseURI(baseURI, subscriptionID)
}
func NewListResultIterator(page ListResultPage) ListResultIterator {
	return original.NewListResultIterator(page)
}
func NewListResultPage(cur ListResult, getNextPage func(context.Context, ListResult) (ListResult, error)) ListResultPage {
	return original.NewListResultPage(cur, getNextPage)
}
func NewManagedClusterListResultIterator(page ManagedClusterListResultPage) ManagedClusterListResultIterator {
	return original.NewManagedClusterListResultIterator(page)
}
func NewManagedClusterListResultPage(cur ManagedClusterListResult, getNextPage func(context.Context, ManagedClusterListResult) (ManagedClusterListResult, error)) ManagedClusterListResultPage {
	return original.NewManagedClusterListResultPage(cur, getNextPage)
}
func NewManagedClustersClient(subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClient(subscriptionID)
}
func NewManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) ManagedClustersClient {
	return original.NewManagedClustersClientWithBaseURI(baseURI, subscriptionID)
}
func NewOpenShiftManagedClusterListResultIterator(page OpenShiftManagedClusterListResultPage) OpenShiftManagedClusterListResultIterator {
	return original.NewOpenShiftManagedClusterListResultIterator(page)
}
func NewOpenShiftManagedClusterListResultPage(cur OpenShiftManagedClusterListResult, getNextPage func(context.Context, OpenShiftManagedClusterListResult) (OpenShiftManagedClusterListResult, error)) OpenShiftManagedClusterListResultPage {
	return original.NewOpenShiftManagedClusterListResultPage(cur, getNextPage)
}
func NewOpenShiftManagedClustersClient(subscriptionID string) OpenShiftManagedClustersClient {
	return original.NewOpenShiftManagedClustersClient(subscriptionID)
}
func NewOpenShiftManagedClustersClientWithBaseURI(baseURI string, subscriptionID string) OpenShiftManagedClustersClient {
	return original.NewOpenShiftManagedClustersClientWithBaseURI(baseURI, subscriptionID)
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
func PossibleAgentPoolTypeValues() []AgentPoolType {
	return original.PossibleAgentPoolTypeValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleLoadBalancerSkuValues() []LoadBalancerSku {
	return original.PossibleLoadBalancerSkuValues()
}
func PossibleNetworkPluginValues() []NetworkPlugin {
	return original.PossibleNetworkPluginValues()
}
func PossibleNetworkPolicyValues() []NetworkPolicy {
	return original.PossibleNetworkPolicyValues()
}
func PossibleOSTypeValues() []OSType {
	return original.PossibleOSTypeValues()
}
func PossibleOpenShiftAgentPoolProfileRoleValues() []OpenShiftAgentPoolProfileRole {
	return original.PossibleOpenShiftAgentPoolProfileRoleValues()
}
func PossibleOpenShiftContainerServiceVMSizeValues() []OpenShiftContainerServiceVMSize {
	return original.PossibleOpenShiftContainerServiceVMSizeValues()
}
func PossibleOrchestratorTypesValues() []OrchestratorTypes {
	return original.PossibleOrchestratorTypesValues()
}
func PossibleOutboundTypeValues() []OutboundType {
	return original.PossibleOutboundTypeValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleScaleSetEvictionPolicyValues() []ScaleSetEvictionPolicy {
	return original.PossibleScaleSetEvictionPolicyValues()
}
func PossibleScaleSetPriorityValues() []ScaleSetPriority {
	return original.PossibleScaleSetPriorityValues()
}
func PossibleStorageProfileTypesValues() []StorageProfileTypes {
	return original.PossibleStorageProfileTypesValues()
}
func PossibleVMSizeTypesValues() []VMSizeTypes {
	return original.PossibleVMSizeTypesValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
