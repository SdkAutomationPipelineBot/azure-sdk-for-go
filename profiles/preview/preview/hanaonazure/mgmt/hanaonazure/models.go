//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package hanaonazure

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/hanaonazure/mgmt/2017-11-03-preview/hanaonazure"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type HanaHardwareTypeNamesEnum = original.HanaHardwareTypeNamesEnum

const (
	CiscoUCS HanaHardwareTypeNamesEnum = original.CiscoUCS
	HPE      HanaHardwareTypeNamesEnum = original.HPE
)

type HanaInstancePowerStateEnum = original.HanaInstancePowerStateEnum

const (
	Restarting HanaInstancePowerStateEnum = original.Restarting
	Started    HanaInstancePowerStateEnum = original.Started
	Starting   HanaInstancePowerStateEnum = original.Starting
	Stopped    HanaInstancePowerStateEnum = original.Stopped
	Stopping   HanaInstancePowerStateEnum = original.Stopping
	Unknown    HanaInstancePowerStateEnum = original.Unknown
)

type HanaInstanceSizeNamesEnum = original.HanaInstanceSizeNamesEnum

const (
	S112    HanaInstanceSizeNamesEnum = original.S112
	S144    HanaInstanceSizeNamesEnum = original.S144
	S144m   HanaInstanceSizeNamesEnum = original.S144m
	S192    HanaInstanceSizeNamesEnum = original.S192
	S192m   HanaInstanceSizeNamesEnum = original.S192m
	S192xm  HanaInstanceSizeNamesEnum = original.S192xm
	S224    HanaInstanceSizeNamesEnum = original.S224
	S224m   HanaInstanceSizeNamesEnum = original.S224m
	S224om  HanaInstanceSizeNamesEnum = original.S224om
	S224oo  HanaInstanceSizeNamesEnum = original.S224oo
	S224oom HanaInstanceSizeNamesEnum = original.S224oom
	S224ooo HanaInstanceSizeNamesEnum = original.S224ooo
	S384    HanaInstanceSizeNamesEnum = original.S384
	S384m   HanaInstanceSizeNamesEnum = original.S384m
	S384xm  HanaInstanceSizeNamesEnum = original.S384xm
	S384xxm HanaInstanceSizeNamesEnum = original.S384xxm
	S448    HanaInstanceSizeNamesEnum = original.S448
	S448m   HanaInstanceSizeNamesEnum = original.S448m
	S448om  HanaInstanceSizeNamesEnum = original.S448om
	S448oo  HanaInstanceSizeNamesEnum = original.S448oo
	S448oom HanaInstanceSizeNamesEnum = original.S448oom
	S448ooo HanaInstanceSizeNamesEnum = original.S448ooo
	S576m   HanaInstanceSizeNamesEnum = original.S576m
	S576xm  HanaInstanceSizeNamesEnum = original.S576xm
	S672    HanaInstanceSizeNamesEnum = original.S672
	S672m   HanaInstanceSizeNamesEnum = original.S672m
	S672om  HanaInstanceSizeNamesEnum = original.S672om
	S672oo  HanaInstanceSizeNamesEnum = original.S672oo
	S672oom HanaInstanceSizeNamesEnum = original.S672oom
	S672ooo HanaInstanceSizeNamesEnum = original.S672ooo
	S72     HanaInstanceSizeNamesEnum = original.S72
	S72m    HanaInstanceSizeNamesEnum = original.S72m
	S768    HanaInstanceSizeNamesEnum = original.S768
	S768m   HanaInstanceSizeNamesEnum = original.S768m
	S768xm  HanaInstanceSizeNamesEnum = original.S768xm
	S896    HanaInstanceSizeNamesEnum = original.S896
	S896m   HanaInstanceSizeNamesEnum = original.S896m
	S896om  HanaInstanceSizeNamesEnum = original.S896om
	S896oo  HanaInstanceSizeNamesEnum = original.S896oo
	S896oom HanaInstanceSizeNamesEnum = original.S896oom
	S896ooo HanaInstanceSizeNamesEnum = original.S896ooo
	S96     HanaInstanceSizeNamesEnum = original.S96
	S960m   HanaInstanceSizeNamesEnum = original.S960m
)

type HanaProvisioningStatesEnum = original.HanaProvisioningStatesEnum

const (
	Accepted  HanaProvisioningStatesEnum = original.Accepted
	Creating  HanaProvisioningStatesEnum = original.Creating
	Deleting  HanaProvisioningStatesEnum = original.Deleting
	Failed    HanaProvisioningStatesEnum = original.Failed
	Migrating HanaProvisioningStatesEnum = original.Migrating
	Succeeded HanaProvisioningStatesEnum = original.Succeeded
	Updating  HanaProvisioningStatesEnum = original.Updating
)

type BaseClient = original.BaseClient
type Disk = original.Disk
type Display = original.Display
type ErrorResponse = original.ErrorResponse
type HanaInstance = original.HanaInstance
type HanaInstanceProperties = original.HanaInstanceProperties
type HanaInstancesClient = original.HanaInstancesClient
type HanaInstancesCreateFuture = original.HanaInstancesCreateFuture
type HanaInstancesDeleteFuture = original.HanaInstancesDeleteFuture
type HanaInstancesListResult = original.HanaInstancesListResult
type HanaInstancesListResultIterator = original.HanaInstancesListResultIterator
type HanaInstancesListResultPage = original.HanaInstancesListResultPage
type HanaInstancesRestartFuture = original.HanaInstancesRestartFuture
type HanaInstancesShutdownFuture = original.HanaInstancesShutdownFuture
type HanaInstancesStartFuture = original.HanaInstancesStartFuture
type HardwareProfile = original.HardwareProfile
type IPAddress = original.IPAddress
type MonitoringDetails = original.MonitoringDetails
type NetworkProfile = original.NetworkProfile
type OSProfile = original.OSProfile
type Operation = original.Operation
type OperationList = original.OperationList
type OperationsClient = original.OperationsClient
type Resource = original.Resource
type SapMonitor = original.SapMonitor
type SapMonitorListResult = original.SapMonitorListResult
type SapMonitorListResultIterator = original.SapMonitorListResultIterator
type SapMonitorListResultPage = original.SapMonitorListResultPage
type SapMonitorProperties = original.SapMonitorProperties
type SapMonitorsClient = original.SapMonitorsClient
type SapMonitorsCreateFuture = original.SapMonitorsCreateFuture
type SapMonitorsDeleteFuture = original.SapMonitorsDeleteFuture
type StorageProfile = original.StorageProfile
type Tags = original.Tags

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewHanaInstancesClient(subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClient(subscriptionID)
}
func NewHanaInstancesClientWithBaseURI(baseURI string, subscriptionID string) HanaInstancesClient {
	return original.NewHanaInstancesClientWithBaseURI(baseURI, subscriptionID)
}
func NewHanaInstancesListResultIterator(page HanaInstancesListResultPage) HanaInstancesListResultIterator {
	return original.NewHanaInstancesListResultIterator(page)
}
func NewHanaInstancesListResultPage(cur HanaInstancesListResult, getNextPage func(context.Context, HanaInstancesListResult) (HanaInstancesListResult, error)) HanaInstancesListResultPage {
	return original.NewHanaInstancesListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewSapMonitorListResultIterator(page SapMonitorListResultPage) SapMonitorListResultIterator {
	return original.NewSapMonitorListResultIterator(page)
}
func NewSapMonitorListResultPage(cur SapMonitorListResult, getNextPage func(context.Context, SapMonitorListResult) (SapMonitorListResult, error)) SapMonitorListResultPage {
	return original.NewSapMonitorListResultPage(cur, getNextPage)
}
func NewSapMonitorsClient(subscriptionID string) SapMonitorsClient {
	return original.NewSapMonitorsClient(subscriptionID)
}
func NewSapMonitorsClientWithBaseURI(baseURI string, subscriptionID string) SapMonitorsClient {
	return original.NewSapMonitorsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleHanaHardwareTypeNamesEnumValues() []HanaHardwareTypeNamesEnum {
	return original.PossibleHanaHardwareTypeNamesEnumValues()
}
func PossibleHanaInstancePowerStateEnumValues() []HanaInstancePowerStateEnum {
	return original.PossibleHanaInstancePowerStateEnumValues()
}
func PossibleHanaInstanceSizeNamesEnumValues() []HanaInstanceSizeNamesEnum {
	return original.PossibleHanaInstanceSizeNamesEnumValues()
}
func PossibleHanaProvisioningStatesEnumValues() []HanaProvisioningStatesEnum {
	return original.PossibleHanaProvisioningStatesEnumValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
