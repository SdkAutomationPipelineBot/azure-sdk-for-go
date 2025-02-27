//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package migrate

import original "github.com/Azure/azure-sdk-for-go/services/preview/migrate/mgmt/2018-09-01-preview/migrate"

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type CleanupState = original.CleanupState

const (
	Completed  CleanupState = original.Completed
	Failed     CleanupState = original.Failed
	InProgress CleanupState = original.InProgress
	None       CleanupState = original.None
	Started    CleanupState = original.Started
)

type ContainerElementKind = original.ContainerElementKind

const (
	ContainerElementKindActionImport   ContainerElementKind = original.ContainerElementKindActionImport
	ContainerElementKindEntitySet      ContainerElementKind = original.ContainerElementKindEntitySet
	ContainerElementKindFunctionImport ContainerElementKind = original.ContainerElementKindFunctionImport
	ContainerElementKindNone           ContainerElementKind = original.ContainerElementKindNone
	ContainerElementKindSingleton      ContainerElementKind = original.ContainerElementKindSingleton
)

type ExpressionKind = original.ExpressionKind

const (
	ExpressionKindAnnotationPath             ExpressionKind = original.ExpressionKindAnnotationPath
	ExpressionKindBinaryConstant             ExpressionKind = original.ExpressionKindBinaryConstant
	ExpressionKindBooleanConstant            ExpressionKind = original.ExpressionKindBooleanConstant
	ExpressionKindCast                       ExpressionKind = original.ExpressionKindCast
	ExpressionKindCollection                 ExpressionKind = original.ExpressionKindCollection
	ExpressionKindDateConstant               ExpressionKind = original.ExpressionKindDateConstant
	ExpressionKindDateTimeOffsetConstant     ExpressionKind = original.ExpressionKindDateTimeOffsetConstant
	ExpressionKindDecimalConstant            ExpressionKind = original.ExpressionKindDecimalConstant
	ExpressionKindDurationConstant           ExpressionKind = original.ExpressionKindDurationConstant
	ExpressionKindEnumMember                 ExpressionKind = original.ExpressionKindEnumMember
	ExpressionKindFloatingConstant           ExpressionKind = original.ExpressionKindFloatingConstant
	ExpressionKindFunctionApplication        ExpressionKind = original.ExpressionKindFunctionApplication
	ExpressionKindGUIDConstant               ExpressionKind = original.ExpressionKindGUIDConstant
	ExpressionKindIf                         ExpressionKind = original.ExpressionKindIf
	ExpressionKindIntegerConstant            ExpressionKind = original.ExpressionKindIntegerConstant
	ExpressionKindIsType                     ExpressionKind = original.ExpressionKindIsType
	ExpressionKindLabeled                    ExpressionKind = original.ExpressionKindLabeled
	ExpressionKindLabeledExpressionReference ExpressionKind = original.ExpressionKindLabeledExpressionReference
	ExpressionKindNavigationPropertyPath     ExpressionKind = original.ExpressionKindNavigationPropertyPath
	ExpressionKindNone                       ExpressionKind = original.ExpressionKindNone
	ExpressionKindNull                       ExpressionKind = original.ExpressionKindNull
	ExpressionKindPath                       ExpressionKind = original.ExpressionKindPath
	ExpressionKindPropertyPath               ExpressionKind = original.ExpressionKindPropertyPath
	ExpressionKindRecord                     ExpressionKind = original.ExpressionKindRecord
	ExpressionKindStringConstant             ExpressionKind = original.ExpressionKindStringConstant
	ExpressionKindTimeOfDayConstant          ExpressionKind = original.ExpressionKindTimeOfDayConstant
)

type ExpressionKind1 = original.ExpressionKind1

const (
	ExpressionKind1AnnotationPath             ExpressionKind1 = original.ExpressionKind1AnnotationPath
	ExpressionKind1BinaryConstant             ExpressionKind1 = original.ExpressionKind1BinaryConstant
	ExpressionKind1BooleanConstant            ExpressionKind1 = original.ExpressionKind1BooleanConstant
	ExpressionKind1Cast                       ExpressionKind1 = original.ExpressionKind1Cast
	ExpressionKind1Collection                 ExpressionKind1 = original.ExpressionKind1Collection
	ExpressionKind1DateConstant               ExpressionKind1 = original.ExpressionKind1DateConstant
	ExpressionKind1DateTimeOffsetConstant     ExpressionKind1 = original.ExpressionKind1DateTimeOffsetConstant
	ExpressionKind1DecimalConstant            ExpressionKind1 = original.ExpressionKind1DecimalConstant
	ExpressionKind1DurationConstant           ExpressionKind1 = original.ExpressionKind1DurationConstant
	ExpressionKind1EnumMember                 ExpressionKind1 = original.ExpressionKind1EnumMember
	ExpressionKind1FloatingConstant           ExpressionKind1 = original.ExpressionKind1FloatingConstant
	ExpressionKind1FunctionApplication        ExpressionKind1 = original.ExpressionKind1FunctionApplication
	ExpressionKind1GUIDConstant               ExpressionKind1 = original.ExpressionKind1GUIDConstant
	ExpressionKind1If                         ExpressionKind1 = original.ExpressionKind1If
	ExpressionKind1IntegerConstant            ExpressionKind1 = original.ExpressionKind1IntegerConstant
	ExpressionKind1IsType                     ExpressionKind1 = original.ExpressionKind1IsType
	ExpressionKind1Labeled                    ExpressionKind1 = original.ExpressionKind1Labeled
	ExpressionKind1LabeledExpressionReference ExpressionKind1 = original.ExpressionKind1LabeledExpressionReference
	ExpressionKind1NavigationPropertyPath     ExpressionKind1 = original.ExpressionKind1NavigationPropertyPath
	ExpressionKind1None                       ExpressionKind1 = original.ExpressionKind1None
	ExpressionKind1Null                       ExpressionKind1 = original.ExpressionKind1Null
	ExpressionKind1Path                       ExpressionKind1 = original.ExpressionKind1Path
	ExpressionKind1PropertyPath               ExpressionKind1 = original.ExpressionKind1PropertyPath
	ExpressionKind1Record                     ExpressionKind1 = original.ExpressionKind1Record
	ExpressionKind1StringConstant             ExpressionKind1 = original.ExpressionKind1StringConstant
	ExpressionKind1TimeOfDayConstant          ExpressionKind1 = original.ExpressionKind1TimeOfDayConstant
)

type Goal = original.Goal

const (
	Databases Goal = original.Databases
	Servers   Goal = original.Servers
)

type Goal1 = original.Goal1

const (
	Goal1Databases Goal1 = original.Goal1Databases
	Goal1Servers   Goal1 = original.Goal1Servers
)

type InstanceType = original.InstanceType

const (
	InstanceTypeDatabases              InstanceType = original.InstanceTypeDatabases
	InstanceTypeMigrateEventProperties InstanceType = original.InstanceTypeMigrateEventProperties
	InstanceTypeServers                InstanceType = original.InstanceTypeServers
)

type InstanceTypeBasicProjectSummary = original.InstanceTypeBasicProjectSummary

const (
	InstanceTypeBasicProjectSummaryInstanceTypeDatabases      InstanceTypeBasicProjectSummary = original.InstanceTypeBasicProjectSummaryInstanceTypeDatabases
	InstanceTypeBasicProjectSummaryInstanceTypeProjectSummary InstanceTypeBasicProjectSummary = original.InstanceTypeBasicProjectSummaryInstanceTypeProjectSummary
	InstanceTypeBasicProjectSummaryInstanceTypeServers        InstanceTypeBasicProjectSummary = original.InstanceTypeBasicProjectSummaryInstanceTypeServers
)

type InstanceTypeBasicSolutionSummary = original.InstanceTypeBasicSolutionSummary

const (
	InstanceTypeBasicSolutionSummaryInstanceTypeDatabases       InstanceTypeBasicSolutionSummary = original.InstanceTypeBasicSolutionSummaryInstanceTypeDatabases
	InstanceTypeBasicSolutionSummaryInstanceTypeServers         InstanceTypeBasicSolutionSummary = original.InstanceTypeBasicSolutionSummaryInstanceTypeServers
	InstanceTypeBasicSolutionSummaryInstanceTypeSolutionSummary InstanceTypeBasicSolutionSummary = original.InstanceTypeBasicSolutionSummaryInstanceTypeSolutionSummary
)

type Kind = original.Kind

const (
	KindAggregatedCollectionPropertyNode  Kind = original.KindAggregatedCollectionPropertyNode
	KindAll                               Kind = original.KindAll
	KindAny                               Kind = original.KindAny
	KindBinaryOperator                    Kind = original.KindBinaryOperator
	KindCollectionComplexNode             Kind = original.KindCollectionComplexNode
	KindCollectionConstant                Kind = original.KindCollectionConstant
	KindCollectionFunctionCall            Kind = original.KindCollectionFunctionCall
	KindCollectionNavigationNode          Kind = original.KindCollectionNavigationNode
	KindCollectionOpenPropertyAccess      Kind = original.KindCollectionOpenPropertyAccess
	KindCollectionPropertyAccess          Kind = original.KindCollectionPropertyAccess
	KindCollectionPropertyNode            Kind = original.KindCollectionPropertyNode
	KindCollectionResourceCast            Kind = original.KindCollectionResourceCast
	KindCollectionResourceFunctionCall    Kind = original.KindCollectionResourceFunctionCall
	KindConstant                          Kind = original.KindConstant
	KindConvert                           Kind = original.KindConvert
	KindCount                             Kind = original.KindCount
	KindEntitySet                         Kind = original.KindEntitySet
	KindIn                                Kind = original.KindIn
	KindKeyLookup                         Kind = original.KindKeyLookup
	KindNamedFunctionParameter            Kind = original.KindNamedFunctionParameter
	KindNone                              Kind = original.KindNone
	KindNonResourceRangeVariableReference Kind = original.KindNonResourceRangeVariableReference
	KindParameterAlias                    Kind = original.KindParameterAlias
	KindResourceRangeVariableReference    Kind = original.KindResourceRangeVariableReference
	KindSearchTerm                        Kind = original.KindSearchTerm
	KindSingleComplexNode                 Kind = original.KindSingleComplexNode
	KindSingleNavigationNode              Kind = original.KindSingleNavigationNode
	KindSingleResourceCast                Kind = original.KindSingleResourceCast
	KindSingleResourceFunctionCall        Kind = original.KindSingleResourceFunctionCall
	KindSingleValueCast                   Kind = original.KindSingleValueCast
	KindSingleValueFunctionCall           Kind = original.KindSingleValueFunctionCall
	KindSingleValueOpenPropertyAccess     Kind = original.KindSingleValueOpenPropertyAccess
	KindSingleValuePropertyAccess         Kind = original.KindSingleValuePropertyAccess
	KindUnaryOperator                     Kind = original.KindUnaryOperator
)

type Kind1 = original.Kind1

const (
	Aggregate Kind1 = original.Aggregate
	Compute   Kind1 = original.Compute
	Filter    Kind1 = original.Filter
	GroupBy   Kind1 = original.GroupBy
)

type OnDelete = original.OnDelete

const (
	OnDeleteCascade OnDelete = original.OnDeleteCascade
	OnDeleteNone    OnDelete = original.OnDeleteNone
)

type PropertyKind = original.PropertyKind

const (
	PropertyKindNavigation PropertyKind = original.PropertyKindNavigation
	PropertyKindNone       PropertyKind = original.PropertyKindNone
	PropertyKindStructural PropertyKind = original.PropertyKindStructural
)

type PropertyKind1 = original.PropertyKind1

const (
	PropertyKind1Navigation PropertyKind1 = original.PropertyKind1Navigation
	PropertyKind1None       PropertyKind1 = original.PropertyKind1None
	PropertyKind1Structural PropertyKind1 = original.PropertyKind1Structural
)

type PropertyKind2 = original.PropertyKind2

const (
	PropertyKind2Navigation PropertyKind2 = original.PropertyKind2Navigation
	PropertyKind2None       PropertyKind2 = original.PropertyKind2None
	PropertyKind2Structural PropertyKind2 = original.PropertyKind2Structural
)

type ProvisioningState = original.ProvisioningState

const (
	ProvisioningStateAccepted  ProvisioningState = original.ProvisioningStateAccepted
	ProvisioningStateCreating  ProvisioningState = original.ProvisioningStateCreating
	ProvisioningStateDeleting  ProvisioningState = original.ProvisioningStateDeleting
	ProvisioningStateFailed    ProvisioningState = original.ProvisioningStateFailed
	ProvisioningStateMoving    ProvisioningState = original.ProvisioningStateMoving
	ProvisioningStateSucceeded ProvisioningState = original.ProvisioningStateSucceeded
)

type Purpose = original.Purpose

const (
	Assessment Purpose = original.Assessment
	Discovery  Purpose = original.Discovery
	Migration  Purpose = original.Migration
)

type RefreshSummaryState = original.RefreshSummaryState

const (
	RefreshSummaryStateCompleted  RefreshSummaryState = original.RefreshSummaryStateCompleted
	RefreshSummaryStateFailed     RefreshSummaryState = original.RefreshSummaryStateFailed
	RefreshSummaryStateInProgress RefreshSummaryState = original.RefreshSummaryStateInProgress
	RefreshSummaryStateStarted    RefreshSummaryState = original.RefreshSummaryStateStarted
)

type RefreshSummaryState1 = original.RefreshSummaryState1

const (
	RefreshSummaryState1Completed  RefreshSummaryState1 = original.RefreshSummaryState1Completed
	RefreshSummaryState1Failed     RefreshSummaryState1 = original.RefreshSummaryState1Failed
	RefreshSummaryState1InProgress RefreshSummaryState1 = original.RefreshSummaryState1InProgress
	RefreshSummaryState1Started    RefreshSummaryState1 = original.RefreshSummaryState1Started
)

type SchemaElementKind = original.SchemaElementKind

const (
	SchemaElementKindAction          SchemaElementKind = original.SchemaElementKindAction
	SchemaElementKindEntityContainer SchemaElementKind = original.SchemaElementKindEntityContainer
	SchemaElementKindFunction        SchemaElementKind = original.SchemaElementKindFunction
	SchemaElementKindNone            SchemaElementKind = original.SchemaElementKindNone
	SchemaElementKindTerm            SchemaElementKind = original.SchemaElementKindTerm
	SchemaElementKindTypeDefinition  SchemaElementKind = original.SchemaElementKindTypeDefinition
)

type SchemaElementKind1 = original.SchemaElementKind1

const (
	SchemaElementKind1Action          SchemaElementKind1 = original.SchemaElementKind1Action
	SchemaElementKind1EntityContainer SchemaElementKind1 = original.SchemaElementKind1EntityContainer
	SchemaElementKind1Function        SchemaElementKind1 = original.SchemaElementKind1Function
	SchemaElementKind1None            SchemaElementKind1 = original.SchemaElementKind1None
	SchemaElementKind1Term            SchemaElementKind1 = original.SchemaElementKind1Term
	SchemaElementKind1TypeDefinition  SchemaElementKind1 = original.SchemaElementKind1TypeDefinition
)

type SchemaElementKind2 = original.SchemaElementKind2

const (
	SchemaElementKind2Action          SchemaElementKind2 = original.SchemaElementKind2Action
	SchemaElementKind2EntityContainer SchemaElementKind2 = original.SchemaElementKind2EntityContainer
	SchemaElementKind2Function        SchemaElementKind2 = original.SchemaElementKind2Function
	SchemaElementKind2None            SchemaElementKind2 = original.SchemaElementKind2None
	SchemaElementKind2Term            SchemaElementKind2 = original.SchemaElementKind2Term
	SchemaElementKind2TypeDefinition  SchemaElementKind2 = original.SchemaElementKind2TypeDefinition
)

type Status = original.Status

const (
	Active   Status = original.Active
	Inactive Status = original.Inactive
)

type Tool = original.Tool

const (
	Carbonite                  Tool = original.Carbonite
	Cloudamize                 Tool = original.Cloudamize
	CorentTech                 Tool = original.CorentTech
	DatabaseMigrationService   Tool = original.DatabaseMigrationService
	DataMigrationAssistant     Tool = original.DataMigrationAssistant
	ServerAssessment           Tool = original.ServerAssessment
	ServerAssessmentV1         Tool = original.ServerAssessmentV1
	ServerDiscovery            Tool = original.ServerDiscovery
	ServerMigration            Tool = original.ServerMigration
	ServerMigrationReplication Tool = original.ServerMigrationReplication
	Turbonomic                 Tool = original.Turbonomic
	Zerto                      Tool = original.Zerto
)

type Tool1 = original.Tool1

const (
	Tool1Carbonite                  Tool1 = original.Tool1Carbonite
	Tool1Cloudamize                 Tool1 = original.Tool1Cloudamize
	Tool1CorentTech                 Tool1 = original.Tool1CorentTech
	Tool1DatabaseMigrationService   Tool1 = original.Tool1DatabaseMigrationService
	Tool1DataMigrationAssistant     Tool1 = original.Tool1DataMigrationAssistant
	Tool1ServerAssessment           Tool1 = original.Tool1ServerAssessment
	Tool1ServerAssessmentV1         Tool1 = original.Tool1ServerAssessmentV1
	Tool1ServerDiscovery            Tool1 = original.Tool1ServerDiscovery
	Tool1ServerMigration            Tool1 = original.Tool1ServerMigration
	Tool1ServerMigrationReplication Tool1 = original.Tool1ServerMigrationReplication
	Tool1Turbonomic                 Tool1 = original.Tool1Turbonomic
	Tool1Zerto                      Tool1 = original.Tool1Zerto
)

type TypeKind = original.TypeKind

const (
	TypeKindCollection      TypeKind = original.TypeKindCollection
	TypeKindComplex         TypeKind = original.TypeKindComplex
	TypeKindEntity          TypeKind = original.TypeKindEntity
	TypeKindEntityReference TypeKind = original.TypeKindEntityReference
	TypeKindEnum            TypeKind = original.TypeKindEnum
	TypeKindNone            TypeKind = original.TypeKindNone
	TypeKindPath            TypeKind = original.TypeKindPath
	TypeKindPrimitive       TypeKind = original.TypeKindPrimitive
	TypeKindTypeDefinition  TypeKind = original.TypeKindTypeDefinition
	TypeKindUntyped         TypeKind = original.TypeKindUntyped
)

type TypeKind1 = original.TypeKind1

const (
	TypeKind1Collection      TypeKind1 = original.TypeKind1Collection
	TypeKind1Complex         TypeKind1 = original.TypeKind1Complex
	TypeKind1Entity          TypeKind1 = original.TypeKind1Entity
	TypeKind1EntityReference TypeKind1 = original.TypeKind1EntityReference
	TypeKind1Enum            TypeKind1 = original.TypeKind1Enum
	TypeKind1None            TypeKind1 = original.TypeKind1None
	TypeKind1Path            TypeKind1 = original.TypeKind1Path
	TypeKind1Primitive       TypeKind1 = original.TypeKind1Primitive
	TypeKind1TypeDefinition  TypeKind1 = original.TypeKind1TypeDefinition
	TypeKind1Untyped         TypeKind1 = original.TypeKind1Untyped
)

type ApplyClause = original.ApplyClause
type AssessmentDetails = original.AssessmentDetails
type BaseClient = original.BaseClient
type BasicEventProperties = original.BasicEventProperties
type BasicProjectSummary = original.BasicProjectSummary
type BasicSolutionSummary = original.BasicSolutionSummary
type Database = original.Database
type DatabaseAssessmentDetails = original.DatabaseAssessmentDetails
type DatabaseCollection = original.DatabaseCollection
type DatabaseInstance = original.DatabaseInstance
type DatabaseInstanceCollection = original.DatabaseInstanceCollection
type DatabaseInstanceDiscoveryDetails = original.DatabaseInstanceDiscoveryDetails
type DatabaseInstanceProperties = original.DatabaseInstanceProperties
type DatabaseInstanceSummary = original.DatabaseInstanceSummary
type DatabaseInstancesClient = original.DatabaseInstancesClient
type DatabaseMigrateEventProperties = original.DatabaseMigrateEventProperties
type DatabaseProjectSummary = original.DatabaseProjectSummary
type DatabaseProperties = original.DatabaseProperties
type DatabasesClient = original.DatabasesClient
type DatabasesSolutionSummary = original.DatabasesSolutionSummary
type DefaultQuerySettings = original.DefaultQuerySettings
type DiscoveryDetails = original.DiscoveryDetails
type EdmReferentialConstraintPropertyPair = original.EdmReferentialConstraintPropertyPair
type Event = original.Event
type EventCollection = original.EventCollection
type EventProperties = original.EventProperties
type EventsClient = original.EventsClient
type FilterClause = original.FilterClause
type FilterQueryOption = original.FilterQueryOption
type IEdmEntityContainer = original.IEdmEntityContainer
type IEdmEntityContainerElement = original.IEdmEntityContainerElement
type IEdmExpression = original.IEdmExpression
type IEdmModel = original.IEdmModel
type IEdmNavigationProperty = original.IEdmNavigationProperty
type IEdmNavigationPropertyBinding = original.IEdmNavigationPropertyBinding
type IEdmNavigationSource = original.IEdmNavigationSource
type IEdmPathExpression = original.IEdmPathExpression
type IEdmProperty = original.IEdmProperty
type IEdmReferentialConstraint = original.IEdmReferentialConstraint
type IEdmSchemaElement = original.IEdmSchemaElement
type IEdmStructuralProperty = original.IEdmStructuralProperty
type IEdmStructuredType = original.IEdmStructuredType
type IEdmTerm = original.IEdmTerm
type IEdmType = original.IEdmType
type IEdmTypeReference = original.IEdmTypeReference
type IEdmVocabularyAnnotation = original.IEdmVocabularyAnnotation
type Machine = original.Machine
type MachineCollection = original.MachineCollection
type MachineMigrateEventProperties = original.MachineMigrateEventProperties
type MachineProperties = original.MachineProperties
type MachinesClient = original.MachinesClient
type MigrationDetails = original.MigrationDetails
type ODataPath = original.ODataPath
type ODataPathSegment = original.ODataPathSegment
type ODataQueryContext = original.ODataQueryContext
type ODataQueryOptions1 = original.ODataQueryOptions1
type ODataRawQueryOptions = original.ODataRawQueryOptions
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationResultList = original.OperationResultList
type OperationsClient = original.OperationsClient
type Project = original.Project
type ProjectProperties = original.ProjectProperties
type ProjectSummary = original.ProjectSummary
type ProjectTags = original.ProjectTags
type ProjectsClient = original.ProjectsClient
type RangeVariable = original.RangeVariable
type RefreshSummaryInput = original.RefreshSummaryInput
type RefreshSummaryResult = original.RefreshSummaryResult
type RegisterToolInput = original.RegisterToolInput
type RegistrationResult = original.RegistrationResult
type SelectExpandClause = original.SelectExpandClause
type ServersProjectSummary = original.ServersProjectSummary
type ServersSolutionSummary = original.ServersSolutionSummary
type SingleValueNode = original.SingleValueNode
type Solution = original.Solution
type SolutionConfig = original.SolutionConfig
type SolutionDetails = original.SolutionDetails
type SolutionProperties = original.SolutionProperties
type SolutionSummary = original.SolutionSummary
type SolutionsClient = original.SolutionsClient
type SolutionsCollection = original.SolutionsCollection
type TransformationNode = original.TransformationNode

func New(subscriptionID string, acceptLanguage string) BaseClient {
	return original.New(subscriptionID, acceptLanguage)
}
func NewDatabaseInstancesClient(subscriptionID string, acceptLanguage string) DatabaseInstancesClient {
	return original.NewDatabaseInstancesClient(subscriptionID, acceptLanguage)
}
func NewDatabaseInstancesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) DatabaseInstancesClient {
	return original.NewDatabaseInstancesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewDatabasesClient(subscriptionID string, acceptLanguage string) DatabasesClient {
	return original.NewDatabasesClient(subscriptionID, acceptLanguage)
}
func NewDatabasesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) DatabasesClient {
	return original.NewDatabasesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewEventsClient(subscriptionID string, acceptLanguage string) EventsClient {
	return original.NewEventsClient(subscriptionID, acceptLanguage)
}
func NewEventsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) EventsClient {
	return original.NewEventsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewMachinesClient(subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClient(subscriptionID, acceptLanguage)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewOperationsClient(subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClient(subscriptionID, acceptLanguage)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewProjectsClient(subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClient(subscriptionID, acceptLanguage)
}
func NewProjectsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) ProjectsClient {
	return original.NewProjectsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewSolutionsClient(subscriptionID string, acceptLanguage string) SolutionsClient {
	return original.NewSolutionsClient(subscriptionID, acceptLanguage)
}
func NewSolutionsClientWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) SolutionsClient {
	return original.NewSolutionsClientWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func NewWithBaseURI(baseURI string, subscriptionID string, acceptLanguage string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID, acceptLanguage)
}
func PossibleCleanupStateValues() []CleanupState {
	return original.PossibleCleanupStateValues()
}
func PossibleContainerElementKindValues() []ContainerElementKind {
	return original.PossibleContainerElementKindValues()
}
func PossibleExpressionKind1Values() []ExpressionKind1 {
	return original.PossibleExpressionKind1Values()
}
func PossibleExpressionKindValues() []ExpressionKind {
	return original.PossibleExpressionKindValues()
}
func PossibleGoal1Values() []Goal1 {
	return original.PossibleGoal1Values()
}
func PossibleGoalValues() []Goal {
	return original.PossibleGoalValues()
}
func PossibleInstanceTypeBasicProjectSummaryValues() []InstanceTypeBasicProjectSummary {
	return original.PossibleInstanceTypeBasicProjectSummaryValues()
}
func PossibleInstanceTypeBasicSolutionSummaryValues() []InstanceTypeBasicSolutionSummary {
	return original.PossibleInstanceTypeBasicSolutionSummaryValues()
}
func PossibleInstanceTypeValues() []InstanceType {
	return original.PossibleInstanceTypeValues()
}
func PossibleKind1Values() []Kind1 {
	return original.PossibleKind1Values()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleOnDeleteValues() []OnDelete {
	return original.PossibleOnDeleteValues()
}
func PossiblePropertyKind1Values() []PropertyKind1 {
	return original.PossiblePropertyKind1Values()
}
func PossiblePropertyKind2Values() []PropertyKind2 {
	return original.PossiblePropertyKind2Values()
}
func PossiblePropertyKindValues() []PropertyKind {
	return original.PossiblePropertyKindValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossiblePurposeValues() []Purpose {
	return original.PossiblePurposeValues()
}
func PossibleRefreshSummaryState1Values() []RefreshSummaryState1 {
	return original.PossibleRefreshSummaryState1Values()
}
func PossibleRefreshSummaryStateValues() []RefreshSummaryState {
	return original.PossibleRefreshSummaryStateValues()
}
func PossibleSchemaElementKind1Values() []SchemaElementKind1 {
	return original.PossibleSchemaElementKind1Values()
}
func PossibleSchemaElementKind2Values() []SchemaElementKind2 {
	return original.PossibleSchemaElementKind2Values()
}
func PossibleSchemaElementKindValues() []SchemaElementKind {
	return original.PossibleSchemaElementKindValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func PossibleTool1Values() []Tool1 {
	return original.PossibleTool1Values()
}
func PossibleToolValues() []Tool {
	return original.PossibleToolValues()
}
func PossibleTypeKind1Values() []TypeKind1 {
	return original.PossibleTypeKind1Values()
}
func PossibleTypeKindValues() []TypeKind {
	return original.PossibleTypeKindValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
