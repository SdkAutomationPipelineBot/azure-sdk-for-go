//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package insights

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/monitor/mgmt/2019-06-01/insights"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type AggregationType = original.AggregationType

const (
	Average AggregationType = original.Average
	Count   AggregationType = original.Count
	Maximum AggregationType = original.Maximum
	Minimum AggregationType = original.Minimum
	None    AggregationType = original.None
	Total   AggregationType = original.Total
)

type AlertSeverity = original.AlertSeverity

const (
	Four  AlertSeverity = original.Four
	One   AlertSeverity = original.One
	Three AlertSeverity = original.Three
	Two   AlertSeverity = original.Two
	Zero  AlertSeverity = original.Zero
)

type BaselineSensitivity = original.BaselineSensitivity

const (
	High   BaselineSensitivity = original.High
	Low    BaselineSensitivity = original.Low
	Medium BaselineSensitivity = original.Medium
)

type CategoryType = original.CategoryType

const (
	Logs    CategoryType = original.Logs
	Metrics CategoryType = original.Metrics
)

type ComparisonOperationType = original.ComparisonOperationType

const (
	Equals             ComparisonOperationType = original.Equals
	GreaterThan        ComparisonOperationType = original.GreaterThan
	GreaterThanOrEqual ComparisonOperationType = original.GreaterThanOrEqual
	LessThan           ComparisonOperationType = original.LessThan
	LessThanOrEqual    ComparisonOperationType = original.LessThanOrEqual
	NotEquals          ComparisonOperationType = original.NotEquals
)

type ConditionOperator = original.ConditionOperator

const (
	ConditionOperatorGreaterThan        ConditionOperator = original.ConditionOperatorGreaterThan
	ConditionOperatorGreaterThanOrEqual ConditionOperator = original.ConditionOperatorGreaterThanOrEqual
	ConditionOperatorLessThan           ConditionOperator = original.ConditionOperatorLessThan
	ConditionOperatorLessThanOrEqual    ConditionOperator = original.ConditionOperatorLessThanOrEqual
)

type ConditionalOperator = original.ConditionalOperator

const (
	ConditionalOperatorEqual       ConditionalOperator = original.ConditionalOperatorEqual
	ConditionalOperatorGreaterThan ConditionalOperator = original.ConditionalOperatorGreaterThan
	ConditionalOperatorLessThan    ConditionalOperator = original.ConditionalOperatorLessThan
)

type CriterionType = original.CriterionType

const (
	CriterionTypeDynamicThresholdCriterion CriterionType = original.CriterionTypeDynamicThresholdCriterion
	CriterionTypeMultiMetricCriteria       CriterionType = original.CriterionTypeMultiMetricCriteria
	CriterionTypeStaticThresholdCriterion  CriterionType = original.CriterionTypeStaticThresholdCriterion
)

type DataStatus = original.DataStatus

const (
	NotPresent DataStatus = original.NotPresent
	Present    DataStatus = original.Present
)

type DynamicThresholdOperator = original.DynamicThresholdOperator

const (
	DynamicThresholdOperatorGreaterOrLessThan DynamicThresholdOperator = original.DynamicThresholdOperatorGreaterOrLessThan
	DynamicThresholdOperatorGreaterThan       DynamicThresholdOperator = original.DynamicThresholdOperatorGreaterThan
	DynamicThresholdOperatorLessThan          DynamicThresholdOperator = original.DynamicThresholdOperatorLessThan
)

type DynamicThresholdSensitivity = original.DynamicThresholdSensitivity

const (
	DynamicThresholdSensitivityHigh   DynamicThresholdSensitivity = original.DynamicThresholdSensitivityHigh
	DynamicThresholdSensitivityLow    DynamicThresholdSensitivity = original.DynamicThresholdSensitivityLow
	DynamicThresholdSensitivityMedium DynamicThresholdSensitivity = original.DynamicThresholdSensitivityMedium
)

type Enabled = original.Enabled

const (
	False Enabled = original.False
	True  Enabled = original.True
)

type EventLevel = original.EventLevel

const (
	EventLevelCritical      EventLevel = original.EventLevelCritical
	EventLevelError         EventLevel = original.EventLevelError
	EventLevelInformational EventLevel = original.EventLevelInformational
	EventLevelVerbose       EventLevel = original.EventLevelVerbose
	EventLevelWarning       EventLevel = original.EventLevelWarning
)

type MetricStatisticType = original.MetricStatisticType

const (
	MetricStatisticTypeAverage MetricStatisticType = original.MetricStatisticTypeAverage
	MetricStatisticTypeMax     MetricStatisticType = original.MetricStatisticTypeMax
	MetricStatisticTypeMin     MetricStatisticType = original.MetricStatisticTypeMin
	MetricStatisticTypeSum     MetricStatisticType = original.MetricStatisticTypeSum
)

type MetricTriggerType = original.MetricTriggerType

const (
	MetricTriggerTypeConsecutive MetricTriggerType = original.MetricTriggerTypeConsecutive
	MetricTriggerTypeTotal       MetricTriggerType = original.MetricTriggerTypeTotal
)

type OdataType = original.OdataType

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleManagementEventDataSource
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource          OdataType = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleMetricDataSource
	OdataTypeRuleDataSource                                                      OdataType = original.OdataTypeRuleDataSource
)

type OdataTypeBasicAction = original.OdataTypeBasicAction

const (
	OdataTypeAction                                                                                                                                   OdataTypeBasicAction = original.OdataTypeAction
	OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesAlertingAction    OdataTypeBasicAction = original.OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesAlertingAction
	OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesLogToMetricAction OdataTypeBasicAction = original.OdataTypeMicrosoftWindowsAzureManagementMonitoringAlertsModelsMicrosoftAppInsightsNexusDataContractsResourcesScheduledQueryRulesLogToMetricAction
)

type OdataTypeBasicMetricAlertCriteria = original.OdataTypeBasicMetricAlertCriteria

const (
	OdataTypeMetricAlertCriteria                                         OdataTypeBasicMetricAlertCriteria = original.OdataTypeMetricAlertCriteria
	OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria OdataTypeBasicMetricAlertCriteria = original.OdataTypeMicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria
	OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria   OdataTypeBasicMetricAlertCriteria = original.OdataTypeMicrosoftAzureMonitorSingleResourceMultipleMetricCriteria
	OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria    OdataTypeBasicMetricAlertCriteria = original.OdataTypeMicrosoftAzureMonitorWebtestLocationAvailabilityCriteria
)

type OdataTypeBasicRuleAction = original.OdataTypeBasicRuleAction

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction   OdataTypeBasicRuleAction = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleEmailAction
	OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction OdataTypeBasicRuleAction = original.OdataTypeMicrosoftAzureManagementInsightsModelsRuleWebhookAction
	OdataTypeRuleAction                                              OdataTypeBasicRuleAction = original.OdataTypeRuleAction
)

type OdataTypeBasicRuleCondition = original.OdataTypeBasicRuleCondition

const (
	OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsLocationThresholdRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition   OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsManagementEventRuleCondition
	OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition         OdataTypeBasicRuleCondition = original.OdataTypeMicrosoftAzureManagementInsightsModelsThresholdRuleCondition
	OdataTypeRuleCondition                                                        OdataTypeBasicRuleCondition = original.OdataTypeRuleCondition
)

type OnboardingStatus = original.OnboardingStatus

const (
	NotOnboarded OnboardingStatus = original.NotOnboarded
	Onboarded    OnboardingStatus = original.Onboarded
	Unknown      OnboardingStatus = original.Unknown
)

type Operator = original.Operator

const (
	OperatorEquals             Operator = original.OperatorEquals
	OperatorGreaterThan        Operator = original.OperatorGreaterThan
	OperatorGreaterThanOrEqual Operator = original.OperatorGreaterThanOrEqual
	OperatorLessThan           Operator = original.OperatorLessThan
	OperatorLessThanOrEqual    Operator = original.OperatorLessThanOrEqual
	OperatorNotEquals          Operator = original.OperatorNotEquals
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Deploying ProvisioningState = original.Deploying
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
)

type QueryType = original.QueryType

const (
	ResultCount QueryType = original.ResultCount
)

type ReceiverStatus = original.ReceiverStatus

const (
	ReceiverStatusDisabled     ReceiverStatus = original.ReceiverStatusDisabled
	ReceiverStatusEnabled      ReceiverStatus = original.ReceiverStatusEnabled
	ReceiverStatusNotSpecified ReceiverStatus = original.ReceiverStatusNotSpecified
)

type RecurrenceFrequency = original.RecurrenceFrequency

const (
	RecurrenceFrequencyDay    RecurrenceFrequency = original.RecurrenceFrequencyDay
	RecurrenceFrequencyHour   RecurrenceFrequency = original.RecurrenceFrequencyHour
	RecurrenceFrequencyMinute RecurrenceFrequency = original.RecurrenceFrequencyMinute
	RecurrenceFrequencyMonth  RecurrenceFrequency = original.RecurrenceFrequencyMonth
	RecurrenceFrequencyNone   RecurrenceFrequency = original.RecurrenceFrequencyNone
	RecurrenceFrequencySecond RecurrenceFrequency = original.RecurrenceFrequencySecond
	RecurrenceFrequencyWeek   RecurrenceFrequency = original.RecurrenceFrequencyWeek
	RecurrenceFrequencyYear   RecurrenceFrequency = original.RecurrenceFrequencyYear
)

type ResultType = original.ResultType

const (
	Data     ResultType = original.Data
	Metadata ResultType = original.Metadata
)

type ScaleDirection = original.ScaleDirection

const (
	ScaleDirectionDecrease ScaleDirection = original.ScaleDirectionDecrease
	ScaleDirectionIncrease ScaleDirection = original.ScaleDirectionIncrease
	ScaleDirectionNone     ScaleDirection = original.ScaleDirectionNone
)

type ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationType

const (
	ScaleRuleMetricDimensionOperationTypeEquals    ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationTypeEquals
	ScaleRuleMetricDimensionOperationTypeNotEquals ScaleRuleMetricDimensionOperationType = original.ScaleRuleMetricDimensionOperationTypeNotEquals
)

type ScaleType = original.ScaleType

const (
	ChangeCount        ScaleType = original.ChangeCount
	ExactCount         ScaleType = original.ExactCount
	PercentChangeCount ScaleType = original.PercentChangeCount
)

type Sensitivity = original.Sensitivity

const (
	SensitivityHigh   Sensitivity = original.SensitivityHigh
	SensitivityLow    Sensitivity = original.SensitivityLow
	SensitivityMedium Sensitivity = original.SensitivityMedium
)

type TimeAggregationOperator = original.TimeAggregationOperator

const (
	TimeAggregationOperatorAverage TimeAggregationOperator = original.TimeAggregationOperatorAverage
	TimeAggregationOperatorLast    TimeAggregationOperator = original.TimeAggregationOperatorLast
	TimeAggregationOperatorMaximum TimeAggregationOperator = original.TimeAggregationOperatorMaximum
	TimeAggregationOperatorMinimum TimeAggregationOperator = original.TimeAggregationOperatorMinimum
	TimeAggregationOperatorTotal   TimeAggregationOperator = original.TimeAggregationOperatorTotal
)

type TimeAggregationType = original.TimeAggregationType

const (
	TimeAggregationTypeAverage TimeAggregationType = original.TimeAggregationTypeAverage
	TimeAggregationTypeCount   TimeAggregationType = original.TimeAggregationTypeCount
	TimeAggregationTypeLast    TimeAggregationType = original.TimeAggregationTypeLast
	TimeAggregationTypeMaximum TimeAggregationType = original.TimeAggregationTypeMaximum
	TimeAggregationTypeMinimum TimeAggregationType = original.TimeAggregationTypeMinimum
	TimeAggregationTypeTotal   TimeAggregationType = original.TimeAggregationTypeTotal
)

type Unit = original.Unit

const (
	UnitBitsPerSecond  Unit = original.UnitBitsPerSecond
	UnitBytes          Unit = original.UnitBytes
	UnitByteSeconds    Unit = original.UnitByteSeconds
	UnitBytesPerSecond Unit = original.UnitBytesPerSecond
	UnitCores          Unit = original.UnitCores
	UnitCount          Unit = original.UnitCount
	UnitCountPerSecond Unit = original.UnitCountPerSecond
	UnitMilliCores     Unit = original.UnitMilliCores
	UnitMilliSeconds   Unit = original.UnitMilliSeconds
	UnitNanoCores      Unit = original.UnitNanoCores
	UnitPercent        Unit = original.UnitPercent
	UnitSeconds        Unit = original.UnitSeconds
	UnitUnspecified    Unit = original.UnitUnspecified
)

type Action = original.Action
type ActionGroup = original.ActionGroup
type ActionGroupList = original.ActionGroupList
type ActionGroupPatch = original.ActionGroupPatch
type ActionGroupPatchBody = original.ActionGroupPatchBody
type ActionGroupResource = original.ActionGroupResource
type ActionGroupsClient = original.ActionGroupsClient
type ActivityLogAlert = original.ActivityLogAlert
type ActivityLogAlertActionGroup = original.ActivityLogAlertActionGroup
type ActivityLogAlertActionList = original.ActivityLogAlertActionList
type ActivityLogAlertAllOfCondition = original.ActivityLogAlertAllOfCondition
type ActivityLogAlertLeafCondition = original.ActivityLogAlertLeafCondition
type ActivityLogAlertList = original.ActivityLogAlertList
type ActivityLogAlertPatch = original.ActivityLogAlertPatch
type ActivityLogAlertPatchBody = original.ActivityLogAlertPatchBody
type ActivityLogAlertResource = original.ActivityLogAlertResource
type ActivityLogAlertsClient = original.ActivityLogAlertsClient
type ActivityLogsClient = original.ActivityLogsClient
type AlertRule = original.AlertRule
type AlertRuleIncidentsClient = original.AlertRuleIncidentsClient
type AlertRuleResource = original.AlertRuleResource
type AlertRuleResourceCollection = original.AlertRuleResourceCollection
type AlertRuleResourcePatch = original.AlertRuleResourcePatch
type AlertRulesClient = original.AlertRulesClient
type AlertingAction = original.AlertingAction
type ArmRoleReceiver = original.ArmRoleReceiver
type AutomationRunbookReceiver = original.AutomationRunbookReceiver
type AutoscaleNotification = original.AutoscaleNotification
type AutoscaleProfile = original.AutoscaleProfile
type AutoscaleSetting = original.AutoscaleSetting
type AutoscaleSettingResource = original.AutoscaleSettingResource
type AutoscaleSettingResourceCollection = original.AutoscaleSettingResourceCollection
type AutoscaleSettingResourceCollectionIterator = original.AutoscaleSettingResourceCollectionIterator
type AutoscaleSettingResourceCollectionPage = original.AutoscaleSettingResourceCollectionPage
type AutoscaleSettingResourcePatch = original.AutoscaleSettingResourcePatch
type AutoscaleSettingsClient = original.AutoscaleSettingsClient
type AzNsActionGroup = original.AzNsActionGroup
type AzureAppPushReceiver = original.AzureAppPushReceiver
type AzureFunctionReceiver = original.AzureFunctionReceiver
type BaseClient = original.BaseClient
type Baseline = original.Baseline
type BaselineMetadata = original.BaselineMetadata
type BaselineMetadataValue = original.BaselineMetadataValue
type BaselineProperties = original.BaselineProperties
type BaselineResponse = original.BaselineResponse
type BaselinesClient = original.BaselinesClient
type BasicAction = original.BasicAction
type BasicMetricAlertCriteria = original.BasicMetricAlertCriteria
type BasicMultiMetricCriteria = original.BasicMultiMetricCriteria
type BasicRuleAction = original.BasicRuleAction
type BasicRuleCondition = original.BasicRuleCondition
type BasicRuleDataSource = original.BasicRuleDataSource
type CalculateBaselineResponse = original.CalculateBaselineResponse
type Criteria = original.Criteria
type DataContainer = original.DataContainer
type DiagnosticSettings = original.DiagnosticSettings
type DiagnosticSettingsCategory = original.DiagnosticSettingsCategory
type DiagnosticSettingsCategoryClient = original.DiagnosticSettingsCategoryClient
type DiagnosticSettingsCategoryResource = original.DiagnosticSettingsCategoryResource
type DiagnosticSettingsCategoryResourceCollection = original.DiagnosticSettingsCategoryResourceCollection
type DiagnosticSettingsClient = original.DiagnosticSettingsClient
type DiagnosticSettingsResource = original.DiagnosticSettingsResource
type DiagnosticSettingsResourceCollection = original.DiagnosticSettingsResourceCollection
type Dimension = original.Dimension
type DynamicMetricCriteria = original.DynamicMetricCriteria
type DynamicThresholdFailingPeriods = original.DynamicThresholdFailingPeriods
type EmailNotification = original.EmailNotification
type EmailReceiver = original.EmailReceiver
type EnableRequest = original.EnableRequest
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type EventCategoriesClient = original.EventCategoriesClient
type EventCategoryCollection = original.EventCategoryCollection
type EventData = original.EventData
type EventDataCollection = original.EventDataCollection
type EventDataCollectionIterator = original.EventDataCollectionIterator
type EventDataCollectionPage = original.EventDataCollectionPage
type HTTPRequestInfo = original.HTTPRequestInfo
type Incident = original.Incident
type IncidentListResult = original.IncidentListResult
type ItsmReceiver = original.ItsmReceiver
type LocalizableString = original.LocalizableString
type LocationThresholdRuleCondition = original.LocationThresholdRuleCondition
type LogMetricTrigger = original.LogMetricTrigger
type LogProfileCollection = original.LogProfileCollection
type LogProfileProperties = original.LogProfileProperties
type LogProfileResource = original.LogProfileResource
type LogProfileResourcePatch = original.LogProfileResourcePatch
type LogProfilesClient = original.LogProfilesClient
type LogSearchRule = original.LogSearchRule
type LogSearchRulePatch = original.LogSearchRulePatch
type LogSearchRuleResource = original.LogSearchRuleResource
type LogSearchRuleResourceCollection = original.LogSearchRuleResourceCollection
type LogSearchRuleResourcePatch = original.LogSearchRuleResourcePatch
type LogSettings = original.LogSettings
type LogToMetricAction = original.LogToMetricAction
type LogicAppReceiver = original.LogicAppReceiver
type ManagementEventAggregationCondition = original.ManagementEventAggregationCondition
type ManagementEventRuleCondition = original.ManagementEventRuleCondition
type MetadataValue = original.MetadataValue
type Metric = original.Metric
type MetricAlertAction = original.MetricAlertAction
type MetricAlertCriteria = original.MetricAlertCriteria
type MetricAlertMultipleResourceMultipleMetricCriteria = original.MetricAlertMultipleResourceMultipleMetricCriteria
type MetricAlertProperties = original.MetricAlertProperties
type MetricAlertResource = original.MetricAlertResource
type MetricAlertResourceCollection = original.MetricAlertResourceCollection
type MetricAlertResourcePatch = original.MetricAlertResourcePatch
type MetricAlertSingleResourceMultipleMetricCriteria = original.MetricAlertSingleResourceMultipleMetricCriteria
type MetricAlertStatus = original.MetricAlertStatus
type MetricAlertStatusCollection = original.MetricAlertStatusCollection
type MetricAlertStatusProperties = original.MetricAlertStatusProperties
type MetricAlertsClient = original.MetricAlertsClient
type MetricAlertsStatusClient = original.MetricAlertsStatusClient
type MetricAvailability = original.MetricAvailability
type MetricBaselineClient = original.MetricBaselineClient
type MetricBaselinesProperties = original.MetricBaselinesProperties
type MetricBaselinesResponse = original.MetricBaselinesResponse
type MetricCriteria = original.MetricCriteria
type MetricDefinition = original.MetricDefinition
type MetricDefinitionCollection = original.MetricDefinitionCollection
type MetricDefinitionsClient = original.MetricDefinitionsClient
type MetricDimension = original.MetricDimension
type MetricNamespace = original.MetricNamespace
type MetricNamespaceCollection = original.MetricNamespaceCollection
type MetricNamespaceName = original.MetricNamespaceName
type MetricNamespacesClient = original.MetricNamespacesClient
type MetricSettings = original.MetricSettings
type MetricSingleDimension = original.MetricSingleDimension
type MetricTrigger = original.MetricTrigger
type MetricValue = original.MetricValue
type MetricsClient = original.MetricsClient
type MultiMetricCriteria = original.MultiMetricCriteria
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationsClient = original.OperationsClient
type ProxyOnlyResource = original.ProxyOnlyResource
type ProxyResource = original.ProxyResource
type Recurrence = original.Recurrence
type RecurrentSchedule = original.RecurrentSchedule
type Resource = original.Resource
type Response = original.Response
type ResponseWithError = original.ResponseWithError
type RetentionPolicy = original.RetentionPolicy
type RuleAction = original.RuleAction
type RuleCondition = original.RuleCondition
type RuleDataSource = original.RuleDataSource
type RuleEmailAction = original.RuleEmailAction
type RuleManagementEventClaimsDataSource = original.RuleManagementEventClaimsDataSource
type RuleManagementEventDataSource = original.RuleManagementEventDataSource
type RuleMetricDataSource = original.RuleMetricDataSource
type RuleWebhookAction = original.RuleWebhookAction
type ScaleAction = original.ScaleAction
type ScaleCapacity = original.ScaleCapacity
type ScaleRule = original.ScaleRule
type ScaleRuleMetricDimension = original.ScaleRuleMetricDimension
type Schedule = original.Schedule
type ScheduledQueryRulesClient = original.ScheduledQueryRulesClient
type SenderAuthorization = original.SenderAuthorization
type SingleBaseline = original.SingleBaseline
type SingleMetricBaseline = original.SingleMetricBaseline
type SmsReceiver = original.SmsReceiver
type Source = original.Source
type TenantActivityLogsClient = original.TenantActivityLogsClient
type ThresholdRuleCondition = original.ThresholdRuleCondition
type TimeSeriesBaseline = original.TimeSeriesBaseline
type TimeSeriesElement = original.TimeSeriesElement
type TimeSeriesInformation = original.TimeSeriesInformation
type TimeWindow = original.TimeWindow
type TriggerCondition = original.TriggerCondition
type VMInsightsClient = original.VMInsightsClient
type VMInsightsOnboardingStatus = original.VMInsightsOnboardingStatus
type VMInsightsOnboardingStatusProperties = original.VMInsightsOnboardingStatusProperties
type VoiceReceiver = original.VoiceReceiver
type WebhookNotification = original.WebhookNotification
type WebhookReceiver = original.WebhookReceiver
type WebtestLocationAvailabilityCriteria = original.WebtestLocationAvailabilityCriteria
type WorkspaceInfo = original.WorkspaceInfo
type WorkspaceInfoProperties = original.WorkspaceInfoProperties

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewActionGroupsClient(subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClient(subscriptionID)
}
func NewActionGroupsClientWithBaseURI(baseURI string, subscriptionID string) ActionGroupsClient {
	return original.NewActionGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActivityLogAlertsClient(subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClient(subscriptionID)
}
func NewActivityLogAlertsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogAlertsClient {
	return original.NewActivityLogAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewActivityLogsClient(subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClient(subscriptionID)
}
func NewActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) ActivityLogsClient {
	return original.NewActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRuleIncidentsClient(subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClient(subscriptionID)
}
func NewAlertRuleIncidentsClientWithBaseURI(baseURI string, subscriptionID string) AlertRuleIncidentsClient {
	return original.NewAlertRuleIncidentsClientWithBaseURI(baseURI, subscriptionID)
}
func NewAlertRulesClient(subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClient(subscriptionID)
}
func NewAlertRulesClientWithBaseURI(baseURI string, subscriptionID string) AlertRulesClient {
	return original.NewAlertRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewAutoscaleSettingResourceCollectionIterator(page AutoscaleSettingResourceCollectionPage) AutoscaleSettingResourceCollectionIterator {
	return original.NewAutoscaleSettingResourceCollectionIterator(page)
}
func NewAutoscaleSettingResourceCollectionPage(cur AutoscaleSettingResourceCollection, getNextPage func(context.Context, AutoscaleSettingResourceCollection) (AutoscaleSettingResourceCollection, error)) AutoscaleSettingResourceCollectionPage {
	return original.NewAutoscaleSettingResourceCollectionPage(cur, getNextPage)
}
func NewAutoscaleSettingsClient(subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClient(subscriptionID)
}
func NewAutoscaleSettingsClientWithBaseURI(baseURI string, subscriptionID string) AutoscaleSettingsClient {
	return original.NewAutoscaleSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewBaselinesClient(subscriptionID string) BaselinesClient {
	return original.NewBaselinesClient(subscriptionID)
}
func NewBaselinesClientWithBaseURI(baseURI string, subscriptionID string) BaselinesClient {
	return original.NewBaselinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewDiagnosticSettingsCategoryClient(subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClient(subscriptionID)
}
func NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsCategoryClient {
	return original.NewDiagnosticSettingsCategoryClientWithBaseURI(baseURI, subscriptionID)
}
func NewDiagnosticSettingsClient(subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClient(subscriptionID)
}
func NewDiagnosticSettingsClientWithBaseURI(baseURI string, subscriptionID string) DiagnosticSettingsClient {
	return original.NewDiagnosticSettingsClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventCategoriesClient(subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClient(subscriptionID)
}
func NewEventCategoriesClientWithBaseURI(baseURI string, subscriptionID string) EventCategoriesClient {
	return original.NewEventCategoriesClientWithBaseURI(baseURI, subscriptionID)
}
func NewEventDataCollectionIterator(page EventDataCollectionPage) EventDataCollectionIterator {
	return original.NewEventDataCollectionIterator(page)
}
func NewEventDataCollectionPage(cur EventDataCollection, getNextPage func(context.Context, EventDataCollection) (EventDataCollection, error)) EventDataCollectionPage {
	return original.NewEventDataCollectionPage(cur, getNextPage)
}
func NewLogProfilesClient(subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClient(subscriptionID)
}
func NewLogProfilesClientWithBaseURI(baseURI string, subscriptionID string) LogProfilesClient {
	return original.NewLogProfilesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricAlertsClient(subscriptionID string) MetricAlertsClient {
	return original.NewMetricAlertsClient(subscriptionID)
}
func NewMetricAlertsClientWithBaseURI(baseURI string, subscriptionID string) MetricAlertsClient {
	return original.NewMetricAlertsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricAlertsStatusClient(subscriptionID string) MetricAlertsStatusClient {
	return original.NewMetricAlertsStatusClient(subscriptionID)
}
func NewMetricAlertsStatusClientWithBaseURI(baseURI string, subscriptionID string) MetricAlertsStatusClient {
	return original.NewMetricAlertsStatusClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricBaselineClient(subscriptionID string) MetricBaselineClient {
	return original.NewMetricBaselineClient(subscriptionID)
}
func NewMetricBaselineClientWithBaseURI(baseURI string, subscriptionID string) MetricBaselineClient {
	return original.NewMetricBaselineClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricDefinitionsClient(subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClient(subscriptionID)
}
func NewMetricDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) MetricDefinitionsClient {
	return original.NewMetricDefinitionsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricNamespacesClient(subscriptionID string) MetricNamespacesClient {
	return original.NewMetricNamespacesClient(subscriptionID)
}
func NewMetricNamespacesClientWithBaseURI(baseURI string, subscriptionID string) MetricNamespacesClient {
	return original.NewMetricNamespacesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMetricsClient(subscriptionID string) MetricsClient {
	return original.NewMetricsClient(subscriptionID)
}
func NewMetricsClientWithBaseURI(baseURI string, subscriptionID string) MetricsClient {
	return original.NewMetricsClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewScheduledQueryRulesClient(subscriptionID string) ScheduledQueryRulesClient {
	return original.NewScheduledQueryRulesClient(subscriptionID)
}
func NewScheduledQueryRulesClientWithBaseURI(baseURI string, subscriptionID string) ScheduledQueryRulesClient {
	return original.NewScheduledQueryRulesClientWithBaseURI(baseURI, subscriptionID)
}
func NewTenantActivityLogsClient(subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClient(subscriptionID)
}
func NewTenantActivityLogsClientWithBaseURI(baseURI string, subscriptionID string) TenantActivityLogsClient {
	return original.NewTenantActivityLogsClientWithBaseURI(baseURI, subscriptionID)
}
func NewVMInsightsClient(subscriptionID string) VMInsightsClient {
	return original.NewVMInsightsClient(subscriptionID)
}
func NewVMInsightsClientWithBaseURI(baseURI string, subscriptionID string) VMInsightsClient {
	return original.NewVMInsightsClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAggregationTypeValues() []AggregationType {
	return original.PossibleAggregationTypeValues()
}
func PossibleAlertSeverityValues() []AlertSeverity {
	return original.PossibleAlertSeverityValues()
}
func PossibleBaselineSensitivityValues() []BaselineSensitivity {
	return original.PossibleBaselineSensitivityValues()
}
func PossibleCategoryTypeValues() []CategoryType {
	return original.PossibleCategoryTypeValues()
}
func PossibleComparisonOperationTypeValues() []ComparisonOperationType {
	return original.PossibleComparisonOperationTypeValues()
}
func PossibleConditionOperatorValues() []ConditionOperator {
	return original.PossibleConditionOperatorValues()
}
func PossibleConditionalOperatorValues() []ConditionalOperator {
	return original.PossibleConditionalOperatorValues()
}
func PossibleCriterionTypeValues() []CriterionType {
	return original.PossibleCriterionTypeValues()
}
func PossibleDataStatusValues() []DataStatus {
	return original.PossibleDataStatusValues()
}
func PossibleDynamicThresholdOperatorValues() []DynamicThresholdOperator {
	return original.PossibleDynamicThresholdOperatorValues()
}
func PossibleDynamicThresholdSensitivityValues() []DynamicThresholdSensitivity {
	return original.PossibleDynamicThresholdSensitivityValues()
}
func PossibleEnabledValues() []Enabled {
	return original.PossibleEnabledValues()
}
func PossibleEventLevelValues() []EventLevel {
	return original.PossibleEventLevelValues()
}
func PossibleMetricStatisticTypeValues() []MetricStatisticType {
	return original.PossibleMetricStatisticTypeValues()
}
func PossibleMetricTriggerTypeValues() []MetricTriggerType {
	return original.PossibleMetricTriggerTypeValues()
}
func PossibleOdataTypeBasicActionValues() []OdataTypeBasicAction {
	return original.PossibleOdataTypeBasicActionValues()
}
func PossibleOdataTypeBasicMetricAlertCriteriaValues() []OdataTypeBasicMetricAlertCriteria {
	return original.PossibleOdataTypeBasicMetricAlertCriteriaValues()
}
func PossibleOdataTypeBasicRuleActionValues() []OdataTypeBasicRuleAction {
	return original.PossibleOdataTypeBasicRuleActionValues()
}
func PossibleOdataTypeBasicRuleConditionValues() []OdataTypeBasicRuleCondition {
	return original.PossibleOdataTypeBasicRuleConditionValues()
}
func PossibleOdataTypeValues() []OdataType {
	return original.PossibleOdataTypeValues()
}
func PossibleOnboardingStatusValues() []OnboardingStatus {
	return original.PossibleOnboardingStatusValues()
}
func PossibleOperatorValues() []Operator {
	return original.PossibleOperatorValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleQueryTypeValues() []QueryType {
	return original.PossibleQueryTypeValues()
}
func PossibleReceiverStatusValues() []ReceiverStatus {
	return original.PossibleReceiverStatusValues()
}
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return original.PossibleRecurrenceFrequencyValues()
}
func PossibleResultTypeValues() []ResultType {
	return original.PossibleResultTypeValues()
}
func PossibleScaleDirectionValues() []ScaleDirection {
	return original.PossibleScaleDirectionValues()
}
func PossibleScaleRuleMetricDimensionOperationTypeValues() []ScaleRuleMetricDimensionOperationType {
	return original.PossibleScaleRuleMetricDimensionOperationTypeValues()
}
func PossibleScaleTypeValues() []ScaleType {
	return original.PossibleScaleTypeValues()
}
func PossibleSensitivityValues() []Sensitivity {
	return original.PossibleSensitivityValues()
}
func PossibleTimeAggregationOperatorValues() []TimeAggregationOperator {
	return original.PossibleTimeAggregationOperatorValues()
}
func PossibleTimeAggregationTypeValues() []TimeAggregationType {
	return original.PossibleTimeAggregationTypeValues()
}
func PossibleUnitValues() []Unit {
	return original.PossibleUnitValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/2020-09-01"
}
func Version() string {
	return original.Version()
}
