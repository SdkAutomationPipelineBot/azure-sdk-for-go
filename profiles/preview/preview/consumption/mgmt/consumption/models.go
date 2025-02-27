//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package consumption

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/consumption/mgmt/2017-12-30-preview/consumption"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type OperatorType = original.OperatorType

const (
	EqualTo              OperatorType = original.EqualTo
	GreaterThan          OperatorType = original.GreaterThan
	GreaterThanOrEqualTo OperatorType = original.GreaterThanOrEqualTo
)

type TimeGrainType = original.TimeGrainType

const (
	Annually  TimeGrainType = original.Annually
	Monthly   TimeGrainType = original.Monthly
	Quarterly TimeGrainType = original.Quarterly
)

type BaseClient = original.BaseClient
type Budget = original.Budget
type BudgetProperties = original.BudgetProperties
type BudgetTimePeriod = original.BudgetTimePeriod
type BudgetsClient = original.BudgetsClient
type BudgetsListResult = original.BudgetsListResult
type CurrentSpend = original.CurrentSpend
type ErrorDetails = original.ErrorDetails
type ErrorResponse = original.ErrorResponse
type Notification = original.Notification
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type ProxyResource = original.ProxyResource

func New(subscriptionID string, name string) BaseClient {
	return original.New(subscriptionID, name)
}
func NewBudgetsClient(subscriptionID string, name string) BudgetsClient {
	return original.NewBudgetsClient(subscriptionID, name)
}
func NewBudgetsClientWithBaseURI(baseURI string, subscriptionID string, name string) BudgetsClient {
	return original.NewBudgetsClientWithBaseURI(baseURI, subscriptionID, name)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(cur OperationListResult, getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(cur, getNextPage)
}
func NewOperationsClient(subscriptionID string, name string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, name)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, name string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, name)
}
func NewWithBaseURI(baseURI string, subscriptionID string, name string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, name)
}
func PossibleOperatorTypeValues() []OperatorType {
	return original.PossibleOperatorTypeValues()
}
func PossibleTimeGrainTypeValues() []TimeGrainType {
	return original.PossibleTimeGrainTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
