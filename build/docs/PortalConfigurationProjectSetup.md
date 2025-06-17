# PortalConfigurationProjectSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PortalConfig** | Pointer to [**PortalConfigurationReference**](PortalConfigurationReference.md) |  | [optional] 
**ProjectNameFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**StatusFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectManagerFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingMethodFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactFlag** | Pointer to **NullableBool** |  | [optional] 
**EstimatedStartFlag** | Pointer to **NullableBool** |  | [optional] 
**EstimatedEndFlag** | Pointer to **NullableBool** |  | [optional] 
**DescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**LastUpdatedFlag** | Pointer to **NullableBool** |  | [optional] 
**OnlyDisplay** | **NullableString** |  | 
**TimeMaterialBudgetHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialScheduledStartFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialScheduledFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialScheduledHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialActualStartFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialActualFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialActualHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialBillFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeMaterialAssignedFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeBudgetHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeScheduledStartFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeScheduledFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeScheduledHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeActualStartFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeActualFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeActualHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeBillFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**FixedFeeAssignedFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueBudgetHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueScheduledStartFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueScheduledFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueScheduledHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueActualStartFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueActualFinishFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueActualHrsFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueBillFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueStatusFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectIssueAssignedFlag** | Pointer to **NullableBool** |  | [optional] 
**ProjectDetailTotalHoursFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPortalConfigurationProjectSetup

`func NewPortalConfigurationProjectSetup(onlyDisplay NullableString, ) *PortalConfigurationProjectSetup`

NewPortalConfigurationProjectSetup instantiates a new PortalConfigurationProjectSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalConfigurationProjectSetupWithDefaults

`func NewPortalConfigurationProjectSetupWithDefaults() *PortalConfigurationProjectSetup`

NewPortalConfigurationProjectSetupWithDefaults instantiates a new PortalConfigurationProjectSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortalConfigurationProjectSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortalConfigurationProjectSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortalConfigurationProjectSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PortalConfigurationProjectSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortalConfig

`func (o *PortalConfigurationProjectSetup) GetPortalConfig() PortalConfigurationReference`

GetPortalConfig returns the PortalConfig field if non-nil, zero value otherwise.

### GetPortalConfigOk

`func (o *PortalConfigurationProjectSetup) GetPortalConfigOk() (*PortalConfigurationReference, bool)`

GetPortalConfigOk returns a tuple with the PortalConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalConfig

`func (o *PortalConfigurationProjectSetup) SetPortalConfig(v PortalConfigurationReference)`

SetPortalConfig sets PortalConfig field to given value.

### HasPortalConfig

`func (o *PortalConfigurationProjectSetup) HasPortalConfig() bool`

HasPortalConfig returns a boolean if a field has been set.

### GetProjectNameFlag

`func (o *PortalConfigurationProjectSetup) GetProjectNameFlag() bool`

GetProjectNameFlag returns the ProjectNameFlag field if non-nil, zero value otherwise.

### GetProjectNameFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectNameFlagOk() (*bool, bool)`

GetProjectNameFlagOk returns a tuple with the ProjectNameFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectNameFlag

`func (o *PortalConfigurationProjectSetup) SetProjectNameFlag(v bool)`

SetProjectNameFlag sets ProjectNameFlag field to given value.

### HasProjectNameFlag

`func (o *PortalConfigurationProjectSetup) HasProjectNameFlag() bool`

HasProjectNameFlag returns a boolean if a field has been set.

### SetProjectNameFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectNameFlagNil(b bool)`

 SetProjectNameFlagNil sets the value for ProjectNameFlag to be an explicit nil

### UnsetProjectNameFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectNameFlag()`

UnsetProjectNameFlag ensures that no value is present for ProjectNameFlag, not even an explicit nil
### GetProjectTypeFlag

`func (o *PortalConfigurationProjectSetup) GetProjectTypeFlag() bool`

GetProjectTypeFlag returns the ProjectTypeFlag field if non-nil, zero value otherwise.

### GetProjectTypeFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectTypeFlagOk() (*bool, bool)`

GetProjectTypeFlagOk returns a tuple with the ProjectTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectTypeFlag

`func (o *PortalConfigurationProjectSetup) SetProjectTypeFlag(v bool)`

SetProjectTypeFlag sets ProjectTypeFlag field to given value.

### HasProjectTypeFlag

`func (o *PortalConfigurationProjectSetup) HasProjectTypeFlag() bool`

HasProjectTypeFlag returns a boolean if a field has been set.

### SetProjectTypeFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectTypeFlagNil(b bool)`

 SetProjectTypeFlagNil sets the value for ProjectTypeFlag to be an explicit nil

### UnsetProjectTypeFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectTypeFlag()`

UnsetProjectTypeFlag ensures that no value is present for ProjectTypeFlag, not even an explicit nil
### GetStatusFlag

`func (o *PortalConfigurationProjectSetup) GetStatusFlag() bool`

GetStatusFlag returns the StatusFlag field if non-nil, zero value otherwise.

### GetStatusFlagOk

`func (o *PortalConfigurationProjectSetup) GetStatusFlagOk() (*bool, bool)`

GetStatusFlagOk returns a tuple with the StatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusFlag

`func (o *PortalConfigurationProjectSetup) SetStatusFlag(v bool)`

SetStatusFlag sets StatusFlag field to given value.

### HasStatusFlag

`func (o *PortalConfigurationProjectSetup) HasStatusFlag() bool`

HasStatusFlag returns a boolean if a field has been set.

### SetStatusFlagNil

`func (o *PortalConfigurationProjectSetup) SetStatusFlagNil(b bool)`

 SetStatusFlagNil sets the value for StatusFlag to be an explicit nil

### UnsetStatusFlag
`func (o *PortalConfigurationProjectSetup) UnsetStatusFlag()`

UnsetStatusFlag ensures that no value is present for StatusFlag, not even an explicit nil
### GetProjectManagerFlag

`func (o *PortalConfigurationProjectSetup) GetProjectManagerFlag() bool`

GetProjectManagerFlag returns the ProjectManagerFlag field if non-nil, zero value otherwise.

### GetProjectManagerFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectManagerFlagOk() (*bool, bool)`

GetProjectManagerFlagOk returns a tuple with the ProjectManagerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectManagerFlag

`func (o *PortalConfigurationProjectSetup) SetProjectManagerFlag(v bool)`

SetProjectManagerFlag sets ProjectManagerFlag field to given value.

### HasProjectManagerFlag

`func (o *PortalConfigurationProjectSetup) HasProjectManagerFlag() bool`

HasProjectManagerFlag returns a boolean if a field has been set.

### SetProjectManagerFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectManagerFlagNil(b bool)`

 SetProjectManagerFlagNil sets the value for ProjectManagerFlag to be an explicit nil

### UnsetProjectManagerFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectManagerFlag()`

UnsetProjectManagerFlag ensures that no value is present for ProjectManagerFlag, not even an explicit nil
### GetBillingMethodFlag

`func (o *PortalConfigurationProjectSetup) GetBillingMethodFlag() bool`

GetBillingMethodFlag returns the BillingMethodFlag field if non-nil, zero value otherwise.

### GetBillingMethodFlagOk

`func (o *PortalConfigurationProjectSetup) GetBillingMethodFlagOk() (*bool, bool)`

GetBillingMethodFlagOk returns a tuple with the BillingMethodFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethodFlag

`func (o *PortalConfigurationProjectSetup) SetBillingMethodFlag(v bool)`

SetBillingMethodFlag sets BillingMethodFlag field to given value.

### HasBillingMethodFlag

`func (o *PortalConfigurationProjectSetup) HasBillingMethodFlag() bool`

HasBillingMethodFlag returns a boolean if a field has been set.

### SetBillingMethodFlagNil

`func (o *PortalConfigurationProjectSetup) SetBillingMethodFlagNil(b bool)`

 SetBillingMethodFlagNil sets the value for BillingMethodFlag to be an explicit nil

### UnsetBillingMethodFlag
`func (o *PortalConfigurationProjectSetup) UnsetBillingMethodFlag()`

UnsetBillingMethodFlag ensures that no value is present for BillingMethodFlag, not even an explicit nil
### GetContactFlag

`func (o *PortalConfigurationProjectSetup) GetContactFlag() bool`

GetContactFlag returns the ContactFlag field if non-nil, zero value otherwise.

### GetContactFlagOk

`func (o *PortalConfigurationProjectSetup) GetContactFlagOk() (*bool, bool)`

GetContactFlagOk returns a tuple with the ContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactFlag

`func (o *PortalConfigurationProjectSetup) SetContactFlag(v bool)`

SetContactFlag sets ContactFlag field to given value.

### HasContactFlag

`func (o *PortalConfigurationProjectSetup) HasContactFlag() bool`

HasContactFlag returns a boolean if a field has been set.

### SetContactFlagNil

`func (o *PortalConfigurationProjectSetup) SetContactFlagNil(b bool)`

 SetContactFlagNil sets the value for ContactFlag to be an explicit nil

### UnsetContactFlag
`func (o *PortalConfigurationProjectSetup) UnsetContactFlag()`

UnsetContactFlag ensures that no value is present for ContactFlag, not even an explicit nil
### GetEstimatedStartFlag

`func (o *PortalConfigurationProjectSetup) GetEstimatedStartFlag() bool`

GetEstimatedStartFlag returns the EstimatedStartFlag field if non-nil, zero value otherwise.

### GetEstimatedStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetEstimatedStartFlagOk() (*bool, bool)`

GetEstimatedStartFlagOk returns a tuple with the EstimatedStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedStartFlag

`func (o *PortalConfigurationProjectSetup) SetEstimatedStartFlag(v bool)`

SetEstimatedStartFlag sets EstimatedStartFlag field to given value.

### HasEstimatedStartFlag

`func (o *PortalConfigurationProjectSetup) HasEstimatedStartFlag() bool`

HasEstimatedStartFlag returns a boolean if a field has been set.

### SetEstimatedStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetEstimatedStartFlagNil(b bool)`

 SetEstimatedStartFlagNil sets the value for EstimatedStartFlag to be an explicit nil

### UnsetEstimatedStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetEstimatedStartFlag()`

UnsetEstimatedStartFlag ensures that no value is present for EstimatedStartFlag, not even an explicit nil
### GetEstimatedEndFlag

`func (o *PortalConfigurationProjectSetup) GetEstimatedEndFlag() bool`

GetEstimatedEndFlag returns the EstimatedEndFlag field if non-nil, zero value otherwise.

### GetEstimatedEndFlagOk

`func (o *PortalConfigurationProjectSetup) GetEstimatedEndFlagOk() (*bool, bool)`

GetEstimatedEndFlagOk returns a tuple with the EstimatedEndFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedEndFlag

`func (o *PortalConfigurationProjectSetup) SetEstimatedEndFlag(v bool)`

SetEstimatedEndFlag sets EstimatedEndFlag field to given value.

### HasEstimatedEndFlag

`func (o *PortalConfigurationProjectSetup) HasEstimatedEndFlag() bool`

HasEstimatedEndFlag returns a boolean if a field has been set.

### SetEstimatedEndFlagNil

`func (o *PortalConfigurationProjectSetup) SetEstimatedEndFlagNil(b bool)`

 SetEstimatedEndFlagNil sets the value for EstimatedEndFlag to be an explicit nil

### UnsetEstimatedEndFlag
`func (o *PortalConfigurationProjectSetup) UnsetEstimatedEndFlag()`

UnsetEstimatedEndFlag ensures that no value is present for EstimatedEndFlag, not even an explicit nil
### GetDescriptionFlag

`func (o *PortalConfigurationProjectSetup) GetDescriptionFlag() bool`

GetDescriptionFlag returns the DescriptionFlag field if non-nil, zero value otherwise.

### GetDescriptionFlagOk

`func (o *PortalConfigurationProjectSetup) GetDescriptionFlagOk() (*bool, bool)`

GetDescriptionFlagOk returns a tuple with the DescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionFlag

`func (o *PortalConfigurationProjectSetup) SetDescriptionFlag(v bool)`

SetDescriptionFlag sets DescriptionFlag field to given value.

### HasDescriptionFlag

`func (o *PortalConfigurationProjectSetup) HasDescriptionFlag() bool`

HasDescriptionFlag returns a boolean if a field has been set.

### SetDescriptionFlagNil

`func (o *PortalConfigurationProjectSetup) SetDescriptionFlagNil(b bool)`

 SetDescriptionFlagNil sets the value for DescriptionFlag to be an explicit nil

### UnsetDescriptionFlag
`func (o *PortalConfigurationProjectSetup) UnsetDescriptionFlag()`

UnsetDescriptionFlag ensures that no value is present for DescriptionFlag, not even an explicit nil
### GetLastUpdatedFlag

`func (o *PortalConfigurationProjectSetup) GetLastUpdatedFlag() bool`

GetLastUpdatedFlag returns the LastUpdatedFlag field if non-nil, zero value otherwise.

### GetLastUpdatedFlagOk

`func (o *PortalConfigurationProjectSetup) GetLastUpdatedFlagOk() (*bool, bool)`

GetLastUpdatedFlagOk returns a tuple with the LastUpdatedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedFlag

`func (o *PortalConfigurationProjectSetup) SetLastUpdatedFlag(v bool)`

SetLastUpdatedFlag sets LastUpdatedFlag field to given value.

### HasLastUpdatedFlag

`func (o *PortalConfigurationProjectSetup) HasLastUpdatedFlag() bool`

HasLastUpdatedFlag returns a boolean if a field has been set.

### SetLastUpdatedFlagNil

`func (o *PortalConfigurationProjectSetup) SetLastUpdatedFlagNil(b bool)`

 SetLastUpdatedFlagNil sets the value for LastUpdatedFlag to be an explicit nil

### UnsetLastUpdatedFlag
`func (o *PortalConfigurationProjectSetup) UnsetLastUpdatedFlag()`

UnsetLastUpdatedFlag ensures that no value is present for LastUpdatedFlag, not even an explicit nil
### GetOnlyDisplay

`func (o *PortalConfigurationProjectSetup) GetOnlyDisplay() string`

GetOnlyDisplay returns the OnlyDisplay field if non-nil, zero value otherwise.

### GetOnlyDisplayOk

`func (o *PortalConfigurationProjectSetup) GetOnlyDisplayOk() (*string, bool)`

GetOnlyDisplayOk returns a tuple with the OnlyDisplay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyDisplay

`func (o *PortalConfigurationProjectSetup) SetOnlyDisplay(v string)`

SetOnlyDisplay sets OnlyDisplay field to given value.


### SetOnlyDisplayNil

`func (o *PortalConfigurationProjectSetup) SetOnlyDisplayNil(b bool)`

 SetOnlyDisplayNil sets the value for OnlyDisplay to be an explicit nil

### UnsetOnlyDisplay
`func (o *PortalConfigurationProjectSetup) UnsetOnlyDisplay()`

UnsetOnlyDisplay ensures that no value is present for OnlyDisplay, not even an explicit nil
### GetTimeMaterialBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialBudgetHrsFlag() bool`

GetTimeMaterialBudgetHrsFlag returns the TimeMaterialBudgetHrsFlag field if non-nil, zero value otherwise.

### GetTimeMaterialBudgetHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialBudgetHrsFlagOk() (*bool, bool)`

GetTimeMaterialBudgetHrsFlagOk returns a tuple with the TimeMaterialBudgetHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialBudgetHrsFlag(v bool)`

SetTimeMaterialBudgetHrsFlag sets TimeMaterialBudgetHrsFlag field to given value.

### HasTimeMaterialBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialBudgetHrsFlag() bool`

HasTimeMaterialBudgetHrsFlag returns a boolean if a field has been set.

### SetTimeMaterialBudgetHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialBudgetHrsFlagNil(b bool)`

 SetTimeMaterialBudgetHrsFlagNil sets the value for TimeMaterialBudgetHrsFlag to be an explicit nil

### UnsetTimeMaterialBudgetHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialBudgetHrsFlag()`

UnsetTimeMaterialBudgetHrsFlag ensures that no value is present for TimeMaterialBudgetHrsFlag, not even an explicit nil
### GetTimeMaterialScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledStartFlag() bool`

GetTimeMaterialScheduledStartFlag returns the TimeMaterialScheduledStartFlag field if non-nil, zero value otherwise.

### GetTimeMaterialScheduledStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledStartFlagOk() (*bool, bool)`

GetTimeMaterialScheduledStartFlagOk returns a tuple with the TimeMaterialScheduledStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledStartFlag(v bool)`

SetTimeMaterialScheduledStartFlag sets TimeMaterialScheduledStartFlag field to given value.

### HasTimeMaterialScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialScheduledStartFlag() bool`

HasTimeMaterialScheduledStartFlag returns a boolean if a field has been set.

### SetTimeMaterialScheduledStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledStartFlagNil(b bool)`

 SetTimeMaterialScheduledStartFlagNil sets the value for TimeMaterialScheduledStartFlag to be an explicit nil

### UnsetTimeMaterialScheduledStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialScheduledStartFlag()`

UnsetTimeMaterialScheduledStartFlag ensures that no value is present for TimeMaterialScheduledStartFlag, not even an explicit nil
### GetTimeMaterialScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledFinishFlag() bool`

GetTimeMaterialScheduledFinishFlag returns the TimeMaterialScheduledFinishFlag field if non-nil, zero value otherwise.

### GetTimeMaterialScheduledFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledFinishFlagOk() (*bool, bool)`

GetTimeMaterialScheduledFinishFlagOk returns a tuple with the TimeMaterialScheduledFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledFinishFlag(v bool)`

SetTimeMaterialScheduledFinishFlag sets TimeMaterialScheduledFinishFlag field to given value.

### HasTimeMaterialScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialScheduledFinishFlag() bool`

HasTimeMaterialScheduledFinishFlag returns a boolean if a field has been set.

### SetTimeMaterialScheduledFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledFinishFlagNil(b bool)`

 SetTimeMaterialScheduledFinishFlagNil sets the value for TimeMaterialScheduledFinishFlag to be an explicit nil

### UnsetTimeMaterialScheduledFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialScheduledFinishFlag()`

UnsetTimeMaterialScheduledFinishFlag ensures that no value is present for TimeMaterialScheduledFinishFlag, not even an explicit nil
### GetTimeMaterialScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledHrsFlag() bool`

GetTimeMaterialScheduledHrsFlag returns the TimeMaterialScheduledHrsFlag field if non-nil, zero value otherwise.

### GetTimeMaterialScheduledHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialScheduledHrsFlagOk() (*bool, bool)`

GetTimeMaterialScheduledHrsFlagOk returns a tuple with the TimeMaterialScheduledHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledHrsFlag(v bool)`

SetTimeMaterialScheduledHrsFlag sets TimeMaterialScheduledHrsFlag field to given value.

### HasTimeMaterialScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialScheduledHrsFlag() bool`

HasTimeMaterialScheduledHrsFlag returns a boolean if a field has been set.

### SetTimeMaterialScheduledHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialScheduledHrsFlagNil(b bool)`

 SetTimeMaterialScheduledHrsFlagNil sets the value for TimeMaterialScheduledHrsFlag to be an explicit nil

### UnsetTimeMaterialScheduledHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialScheduledHrsFlag()`

UnsetTimeMaterialScheduledHrsFlag ensures that no value is present for TimeMaterialScheduledHrsFlag, not even an explicit nil
### GetTimeMaterialActualStartFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualStartFlag() bool`

GetTimeMaterialActualStartFlag returns the TimeMaterialActualStartFlag field if non-nil, zero value otherwise.

### GetTimeMaterialActualStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualStartFlagOk() (*bool, bool)`

GetTimeMaterialActualStartFlagOk returns a tuple with the TimeMaterialActualStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialActualStartFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualStartFlag(v bool)`

SetTimeMaterialActualStartFlag sets TimeMaterialActualStartFlag field to given value.

### HasTimeMaterialActualStartFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialActualStartFlag() bool`

HasTimeMaterialActualStartFlag returns a boolean if a field has been set.

### SetTimeMaterialActualStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualStartFlagNil(b bool)`

 SetTimeMaterialActualStartFlagNil sets the value for TimeMaterialActualStartFlag to be an explicit nil

### UnsetTimeMaterialActualStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialActualStartFlag()`

UnsetTimeMaterialActualStartFlag ensures that no value is present for TimeMaterialActualStartFlag, not even an explicit nil
### GetTimeMaterialActualFinishFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualFinishFlag() bool`

GetTimeMaterialActualFinishFlag returns the TimeMaterialActualFinishFlag field if non-nil, zero value otherwise.

### GetTimeMaterialActualFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualFinishFlagOk() (*bool, bool)`

GetTimeMaterialActualFinishFlagOk returns a tuple with the TimeMaterialActualFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialActualFinishFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualFinishFlag(v bool)`

SetTimeMaterialActualFinishFlag sets TimeMaterialActualFinishFlag field to given value.

### HasTimeMaterialActualFinishFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialActualFinishFlag() bool`

HasTimeMaterialActualFinishFlag returns a boolean if a field has been set.

### SetTimeMaterialActualFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualFinishFlagNil(b bool)`

 SetTimeMaterialActualFinishFlagNil sets the value for TimeMaterialActualFinishFlag to be an explicit nil

### UnsetTimeMaterialActualFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialActualFinishFlag()`

UnsetTimeMaterialActualFinishFlag ensures that no value is present for TimeMaterialActualFinishFlag, not even an explicit nil
### GetTimeMaterialActualHrsFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualHrsFlag() bool`

GetTimeMaterialActualHrsFlag returns the TimeMaterialActualHrsFlag field if non-nil, zero value otherwise.

### GetTimeMaterialActualHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialActualHrsFlagOk() (*bool, bool)`

GetTimeMaterialActualHrsFlagOk returns a tuple with the TimeMaterialActualHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialActualHrsFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualHrsFlag(v bool)`

SetTimeMaterialActualHrsFlag sets TimeMaterialActualHrsFlag field to given value.

### HasTimeMaterialActualHrsFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialActualHrsFlag() bool`

HasTimeMaterialActualHrsFlag returns a boolean if a field has been set.

### SetTimeMaterialActualHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialActualHrsFlagNil(b bool)`

 SetTimeMaterialActualHrsFlagNil sets the value for TimeMaterialActualHrsFlag to be an explicit nil

### UnsetTimeMaterialActualHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialActualHrsFlag()`

UnsetTimeMaterialActualHrsFlag ensures that no value is present for TimeMaterialActualHrsFlag, not even an explicit nil
### GetTimeMaterialBillFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialBillFlag() bool`

GetTimeMaterialBillFlag returns the TimeMaterialBillFlag field if non-nil, zero value otherwise.

### GetTimeMaterialBillFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialBillFlagOk() (*bool, bool)`

GetTimeMaterialBillFlagOk returns a tuple with the TimeMaterialBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialBillFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialBillFlag(v bool)`

SetTimeMaterialBillFlag sets TimeMaterialBillFlag field to given value.

### HasTimeMaterialBillFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialBillFlag() bool`

HasTimeMaterialBillFlag returns a boolean if a field has been set.

### SetTimeMaterialBillFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialBillFlagNil(b bool)`

 SetTimeMaterialBillFlagNil sets the value for TimeMaterialBillFlag to be an explicit nil

### UnsetTimeMaterialBillFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialBillFlag()`

UnsetTimeMaterialBillFlag ensures that no value is present for TimeMaterialBillFlag, not even an explicit nil
### GetTimeMaterialStatusFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialStatusFlag() bool`

GetTimeMaterialStatusFlag returns the TimeMaterialStatusFlag field if non-nil, zero value otherwise.

### GetTimeMaterialStatusFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialStatusFlagOk() (*bool, bool)`

GetTimeMaterialStatusFlagOk returns a tuple with the TimeMaterialStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialStatusFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialStatusFlag(v bool)`

SetTimeMaterialStatusFlag sets TimeMaterialStatusFlag field to given value.

### HasTimeMaterialStatusFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialStatusFlag() bool`

HasTimeMaterialStatusFlag returns a boolean if a field has been set.

### SetTimeMaterialStatusFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialStatusFlagNil(b bool)`

 SetTimeMaterialStatusFlagNil sets the value for TimeMaterialStatusFlag to be an explicit nil

### UnsetTimeMaterialStatusFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialStatusFlag()`

UnsetTimeMaterialStatusFlag ensures that no value is present for TimeMaterialStatusFlag, not even an explicit nil
### GetTimeMaterialAssignedFlag

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialAssignedFlag() bool`

GetTimeMaterialAssignedFlag returns the TimeMaterialAssignedFlag field if non-nil, zero value otherwise.

### GetTimeMaterialAssignedFlagOk

`func (o *PortalConfigurationProjectSetup) GetTimeMaterialAssignedFlagOk() (*bool, bool)`

GetTimeMaterialAssignedFlagOk returns a tuple with the TimeMaterialAssignedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMaterialAssignedFlag

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialAssignedFlag(v bool)`

SetTimeMaterialAssignedFlag sets TimeMaterialAssignedFlag field to given value.

### HasTimeMaterialAssignedFlag

`func (o *PortalConfigurationProjectSetup) HasTimeMaterialAssignedFlag() bool`

HasTimeMaterialAssignedFlag returns a boolean if a field has been set.

### SetTimeMaterialAssignedFlagNil

`func (o *PortalConfigurationProjectSetup) SetTimeMaterialAssignedFlagNil(b bool)`

 SetTimeMaterialAssignedFlagNil sets the value for TimeMaterialAssignedFlag to be an explicit nil

### UnsetTimeMaterialAssignedFlag
`func (o *PortalConfigurationProjectSetup) UnsetTimeMaterialAssignedFlag()`

UnsetTimeMaterialAssignedFlag ensures that no value is present for TimeMaterialAssignedFlag, not even an explicit nil
### GetFixedFeeBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeBudgetHrsFlag() bool`

GetFixedFeeBudgetHrsFlag returns the FixedFeeBudgetHrsFlag field if non-nil, zero value otherwise.

### GetFixedFeeBudgetHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeBudgetHrsFlagOk() (*bool, bool)`

GetFixedFeeBudgetHrsFlagOk returns a tuple with the FixedFeeBudgetHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeBudgetHrsFlag(v bool)`

SetFixedFeeBudgetHrsFlag sets FixedFeeBudgetHrsFlag field to given value.

### HasFixedFeeBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeBudgetHrsFlag() bool`

HasFixedFeeBudgetHrsFlag returns a boolean if a field has been set.

### SetFixedFeeBudgetHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeBudgetHrsFlagNil(b bool)`

 SetFixedFeeBudgetHrsFlagNil sets the value for FixedFeeBudgetHrsFlag to be an explicit nil

### UnsetFixedFeeBudgetHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeBudgetHrsFlag()`

UnsetFixedFeeBudgetHrsFlag ensures that no value is present for FixedFeeBudgetHrsFlag, not even an explicit nil
### GetFixedFeeScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledStartFlag() bool`

GetFixedFeeScheduledStartFlag returns the FixedFeeScheduledStartFlag field if non-nil, zero value otherwise.

### GetFixedFeeScheduledStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledStartFlagOk() (*bool, bool)`

GetFixedFeeScheduledStartFlagOk returns a tuple with the FixedFeeScheduledStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledStartFlag(v bool)`

SetFixedFeeScheduledStartFlag sets FixedFeeScheduledStartFlag field to given value.

### HasFixedFeeScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeScheduledStartFlag() bool`

HasFixedFeeScheduledStartFlag returns a boolean if a field has been set.

### SetFixedFeeScheduledStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledStartFlagNil(b bool)`

 SetFixedFeeScheduledStartFlagNil sets the value for FixedFeeScheduledStartFlag to be an explicit nil

### UnsetFixedFeeScheduledStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeScheduledStartFlag()`

UnsetFixedFeeScheduledStartFlag ensures that no value is present for FixedFeeScheduledStartFlag, not even an explicit nil
### GetFixedFeeScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledFinishFlag() bool`

GetFixedFeeScheduledFinishFlag returns the FixedFeeScheduledFinishFlag field if non-nil, zero value otherwise.

### GetFixedFeeScheduledFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledFinishFlagOk() (*bool, bool)`

GetFixedFeeScheduledFinishFlagOk returns a tuple with the FixedFeeScheduledFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledFinishFlag(v bool)`

SetFixedFeeScheduledFinishFlag sets FixedFeeScheduledFinishFlag field to given value.

### HasFixedFeeScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeScheduledFinishFlag() bool`

HasFixedFeeScheduledFinishFlag returns a boolean if a field has been set.

### SetFixedFeeScheduledFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledFinishFlagNil(b bool)`

 SetFixedFeeScheduledFinishFlagNil sets the value for FixedFeeScheduledFinishFlag to be an explicit nil

### UnsetFixedFeeScheduledFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeScheduledFinishFlag()`

UnsetFixedFeeScheduledFinishFlag ensures that no value is present for FixedFeeScheduledFinishFlag, not even an explicit nil
### GetFixedFeeScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledHrsFlag() bool`

GetFixedFeeScheduledHrsFlag returns the FixedFeeScheduledHrsFlag field if non-nil, zero value otherwise.

### GetFixedFeeScheduledHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeScheduledHrsFlagOk() (*bool, bool)`

GetFixedFeeScheduledHrsFlagOk returns a tuple with the FixedFeeScheduledHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledHrsFlag(v bool)`

SetFixedFeeScheduledHrsFlag sets FixedFeeScheduledHrsFlag field to given value.

### HasFixedFeeScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeScheduledHrsFlag() bool`

HasFixedFeeScheduledHrsFlag returns a boolean if a field has been set.

### SetFixedFeeScheduledHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeScheduledHrsFlagNil(b bool)`

 SetFixedFeeScheduledHrsFlagNil sets the value for FixedFeeScheduledHrsFlag to be an explicit nil

### UnsetFixedFeeScheduledHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeScheduledHrsFlag()`

UnsetFixedFeeScheduledHrsFlag ensures that no value is present for FixedFeeScheduledHrsFlag, not even an explicit nil
### GetFixedFeeActualStartFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualStartFlag() bool`

GetFixedFeeActualStartFlag returns the FixedFeeActualStartFlag field if non-nil, zero value otherwise.

### GetFixedFeeActualStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualStartFlagOk() (*bool, bool)`

GetFixedFeeActualStartFlagOk returns a tuple with the FixedFeeActualStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeActualStartFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualStartFlag(v bool)`

SetFixedFeeActualStartFlag sets FixedFeeActualStartFlag field to given value.

### HasFixedFeeActualStartFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeActualStartFlag() bool`

HasFixedFeeActualStartFlag returns a boolean if a field has been set.

### SetFixedFeeActualStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualStartFlagNil(b bool)`

 SetFixedFeeActualStartFlagNil sets the value for FixedFeeActualStartFlag to be an explicit nil

### UnsetFixedFeeActualStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeActualStartFlag()`

UnsetFixedFeeActualStartFlag ensures that no value is present for FixedFeeActualStartFlag, not even an explicit nil
### GetFixedFeeActualFinishFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualFinishFlag() bool`

GetFixedFeeActualFinishFlag returns the FixedFeeActualFinishFlag field if non-nil, zero value otherwise.

### GetFixedFeeActualFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualFinishFlagOk() (*bool, bool)`

GetFixedFeeActualFinishFlagOk returns a tuple with the FixedFeeActualFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeActualFinishFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualFinishFlag(v bool)`

SetFixedFeeActualFinishFlag sets FixedFeeActualFinishFlag field to given value.

### HasFixedFeeActualFinishFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeActualFinishFlag() bool`

HasFixedFeeActualFinishFlag returns a boolean if a field has been set.

### SetFixedFeeActualFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualFinishFlagNil(b bool)`

 SetFixedFeeActualFinishFlagNil sets the value for FixedFeeActualFinishFlag to be an explicit nil

### UnsetFixedFeeActualFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeActualFinishFlag()`

UnsetFixedFeeActualFinishFlag ensures that no value is present for FixedFeeActualFinishFlag, not even an explicit nil
### GetFixedFeeActualHrsFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualHrsFlag() bool`

GetFixedFeeActualHrsFlag returns the FixedFeeActualHrsFlag field if non-nil, zero value otherwise.

### GetFixedFeeActualHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeActualHrsFlagOk() (*bool, bool)`

GetFixedFeeActualHrsFlagOk returns a tuple with the FixedFeeActualHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeActualHrsFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualHrsFlag(v bool)`

SetFixedFeeActualHrsFlag sets FixedFeeActualHrsFlag field to given value.

### HasFixedFeeActualHrsFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeActualHrsFlag() bool`

HasFixedFeeActualHrsFlag returns a boolean if a field has been set.

### SetFixedFeeActualHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeActualHrsFlagNil(b bool)`

 SetFixedFeeActualHrsFlagNil sets the value for FixedFeeActualHrsFlag to be an explicit nil

### UnsetFixedFeeActualHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeActualHrsFlag()`

UnsetFixedFeeActualHrsFlag ensures that no value is present for FixedFeeActualHrsFlag, not even an explicit nil
### GetFixedFeeBillFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeBillFlag() bool`

GetFixedFeeBillFlag returns the FixedFeeBillFlag field if non-nil, zero value otherwise.

### GetFixedFeeBillFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeBillFlagOk() (*bool, bool)`

GetFixedFeeBillFlagOk returns a tuple with the FixedFeeBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeBillFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeBillFlag(v bool)`

SetFixedFeeBillFlag sets FixedFeeBillFlag field to given value.

### HasFixedFeeBillFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeBillFlag() bool`

HasFixedFeeBillFlag returns a boolean if a field has been set.

### SetFixedFeeBillFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeBillFlagNil(b bool)`

 SetFixedFeeBillFlagNil sets the value for FixedFeeBillFlag to be an explicit nil

### UnsetFixedFeeBillFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeBillFlag()`

UnsetFixedFeeBillFlag ensures that no value is present for FixedFeeBillFlag, not even an explicit nil
### GetFixedFeeStatusFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeStatusFlag() bool`

GetFixedFeeStatusFlag returns the FixedFeeStatusFlag field if non-nil, zero value otherwise.

### GetFixedFeeStatusFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeStatusFlagOk() (*bool, bool)`

GetFixedFeeStatusFlagOk returns a tuple with the FixedFeeStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeStatusFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeStatusFlag(v bool)`

SetFixedFeeStatusFlag sets FixedFeeStatusFlag field to given value.

### HasFixedFeeStatusFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeStatusFlag() bool`

HasFixedFeeStatusFlag returns a boolean if a field has been set.

### SetFixedFeeStatusFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeStatusFlagNil(b bool)`

 SetFixedFeeStatusFlagNil sets the value for FixedFeeStatusFlag to be an explicit nil

### UnsetFixedFeeStatusFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeStatusFlag()`

UnsetFixedFeeStatusFlag ensures that no value is present for FixedFeeStatusFlag, not even an explicit nil
### GetFixedFeeAssignedFlag

`func (o *PortalConfigurationProjectSetup) GetFixedFeeAssignedFlag() bool`

GetFixedFeeAssignedFlag returns the FixedFeeAssignedFlag field if non-nil, zero value otherwise.

### GetFixedFeeAssignedFlagOk

`func (o *PortalConfigurationProjectSetup) GetFixedFeeAssignedFlagOk() (*bool, bool)`

GetFixedFeeAssignedFlagOk returns a tuple with the FixedFeeAssignedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedFeeAssignedFlag

`func (o *PortalConfigurationProjectSetup) SetFixedFeeAssignedFlag(v bool)`

SetFixedFeeAssignedFlag sets FixedFeeAssignedFlag field to given value.

### HasFixedFeeAssignedFlag

`func (o *PortalConfigurationProjectSetup) HasFixedFeeAssignedFlag() bool`

HasFixedFeeAssignedFlag returns a boolean if a field has been set.

### SetFixedFeeAssignedFlagNil

`func (o *PortalConfigurationProjectSetup) SetFixedFeeAssignedFlagNil(b bool)`

 SetFixedFeeAssignedFlagNil sets the value for FixedFeeAssignedFlag to be an explicit nil

### UnsetFixedFeeAssignedFlag
`func (o *PortalConfigurationProjectSetup) UnsetFixedFeeAssignedFlag()`

UnsetFixedFeeAssignedFlag ensures that no value is present for FixedFeeAssignedFlag, not even an explicit nil
### GetProjectIssueBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueBudgetHrsFlag() bool`

GetProjectIssueBudgetHrsFlag returns the ProjectIssueBudgetHrsFlag field if non-nil, zero value otherwise.

### GetProjectIssueBudgetHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueBudgetHrsFlagOk() (*bool, bool)`

GetProjectIssueBudgetHrsFlagOk returns a tuple with the ProjectIssueBudgetHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueBudgetHrsFlag(v bool)`

SetProjectIssueBudgetHrsFlag sets ProjectIssueBudgetHrsFlag field to given value.

### HasProjectIssueBudgetHrsFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueBudgetHrsFlag() bool`

HasProjectIssueBudgetHrsFlag returns a boolean if a field has been set.

### SetProjectIssueBudgetHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueBudgetHrsFlagNil(b bool)`

 SetProjectIssueBudgetHrsFlagNil sets the value for ProjectIssueBudgetHrsFlag to be an explicit nil

### UnsetProjectIssueBudgetHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueBudgetHrsFlag()`

UnsetProjectIssueBudgetHrsFlag ensures that no value is present for ProjectIssueBudgetHrsFlag, not even an explicit nil
### GetProjectIssueScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledStartFlag() bool`

GetProjectIssueScheduledStartFlag returns the ProjectIssueScheduledStartFlag field if non-nil, zero value otherwise.

### GetProjectIssueScheduledStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledStartFlagOk() (*bool, bool)`

GetProjectIssueScheduledStartFlagOk returns a tuple with the ProjectIssueScheduledStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledStartFlag(v bool)`

SetProjectIssueScheduledStartFlag sets ProjectIssueScheduledStartFlag field to given value.

### HasProjectIssueScheduledStartFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueScheduledStartFlag() bool`

HasProjectIssueScheduledStartFlag returns a boolean if a field has been set.

### SetProjectIssueScheduledStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledStartFlagNil(b bool)`

 SetProjectIssueScheduledStartFlagNil sets the value for ProjectIssueScheduledStartFlag to be an explicit nil

### UnsetProjectIssueScheduledStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueScheduledStartFlag()`

UnsetProjectIssueScheduledStartFlag ensures that no value is present for ProjectIssueScheduledStartFlag, not even an explicit nil
### GetProjectIssueScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledFinishFlag() bool`

GetProjectIssueScheduledFinishFlag returns the ProjectIssueScheduledFinishFlag field if non-nil, zero value otherwise.

### GetProjectIssueScheduledFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledFinishFlagOk() (*bool, bool)`

GetProjectIssueScheduledFinishFlagOk returns a tuple with the ProjectIssueScheduledFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledFinishFlag(v bool)`

SetProjectIssueScheduledFinishFlag sets ProjectIssueScheduledFinishFlag field to given value.

### HasProjectIssueScheduledFinishFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueScheduledFinishFlag() bool`

HasProjectIssueScheduledFinishFlag returns a boolean if a field has been set.

### SetProjectIssueScheduledFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledFinishFlagNil(b bool)`

 SetProjectIssueScheduledFinishFlagNil sets the value for ProjectIssueScheduledFinishFlag to be an explicit nil

### UnsetProjectIssueScheduledFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueScheduledFinishFlag()`

UnsetProjectIssueScheduledFinishFlag ensures that no value is present for ProjectIssueScheduledFinishFlag, not even an explicit nil
### GetProjectIssueScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledHrsFlag() bool`

GetProjectIssueScheduledHrsFlag returns the ProjectIssueScheduledHrsFlag field if non-nil, zero value otherwise.

### GetProjectIssueScheduledHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueScheduledHrsFlagOk() (*bool, bool)`

GetProjectIssueScheduledHrsFlagOk returns a tuple with the ProjectIssueScheduledHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledHrsFlag(v bool)`

SetProjectIssueScheduledHrsFlag sets ProjectIssueScheduledHrsFlag field to given value.

### HasProjectIssueScheduledHrsFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueScheduledHrsFlag() bool`

HasProjectIssueScheduledHrsFlag returns a boolean if a field has been set.

### SetProjectIssueScheduledHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueScheduledHrsFlagNil(b bool)`

 SetProjectIssueScheduledHrsFlagNil sets the value for ProjectIssueScheduledHrsFlag to be an explicit nil

### UnsetProjectIssueScheduledHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueScheduledHrsFlag()`

UnsetProjectIssueScheduledHrsFlag ensures that no value is present for ProjectIssueScheduledHrsFlag, not even an explicit nil
### GetProjectIssueActualStartFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualStartFlag() bool`

GetProjectIssueActualStartFlag returns the ProjectIssueActualStartFlag field if non-nil, zero value otherwise.

### GetProjectIssueActualStartFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualStartFlagOk() (*bool, bool)`

GetProjectIssueActualStartFlagOk returns a tuple with the ProjectIssueActualStartFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueActualStartFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualStartFlag(v bool)`

SetProjectIssueActualStartFlag sets ProjectIssueActualStartFlag field to given value.

### HasProjectIssueActualStartFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueActualStartFlag() bool`

HasProjectIssueActualStartFlag returns a boolean if a field has been set.

### SetProjectIssueActualStartFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualStartFlagNil(b bool)`

 SetProjectIssueActualStartFlagNil sets the value for ProjectIssueActualStartFlag to be an explicit nil

### UnsetProjectIssueActualStartFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueActualStartFlag()`

UnsetProjectIssueActualStartFlag ensures that no value is present for ProjectIssueActualStartFlag, not even an explicit nil
### GetProjectIssueActualFinishFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualFinishFlag() bool`

GetProjectIssueActualFinishFlag returns the ProjectIssueActualFinishFlag field if non-nil, zero value otherwise.

### GetProjectIssueActualFinishFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualFinishFlagOk() (*bool, bool)`

GetProjectIssueActualFinishFlagOk returns a tuple with the ProjectIssueActualFinishFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueActualFinishFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualFinishFlag(v bool)`

SetProjectIssueActualFinishFlag sets ProjectIssueActualFinishFlag field to given value.

### HasProjectIssueActualFinishFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueActualFinishFlag() bool`

HasProjectIssueActualFinishFlag returns a boolean if a field has been set.

### SetProjectIssueActualFinishFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualFinishFlagNil(b bool)`

 SetProjectIssueActualFinishFlagNil sets the value for ProjectIssueActualFinishFlag to be an explicit nil

### UnsetProjectIssueActualFinishFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueActualFinishFlag()`

UnsetProjectIssueActualFinishFlag ensures that no value is present for ProjectIssueActualFinishFlag, not even an explicit nil
### GetProjectIssueActualHrsFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualHrsFlag() bool`

GetProjectIssueActualHrsFlag returns the ProjectIssueActualHrsFlag field if non-nil, zero value otherwise.

### GetProjectIssueActualHrsFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueActualHrsFlagOk() (*bool, bool)`

GetProjectIssueActualHrsFlagOk returns a tuple with the ProjectIssueActualHrsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueActualHrsFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualHrsFlag(v bool)`

SetProjectIssueActualHrsFlag sets ProjectIssueActualHrsFlag field to given value.

### HasProjectIssueActualHrsFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueActualHrsFlag() bool`

HasProjectIssueActualHrsFlag returns a boolean if a field has been set.

### SetProjectIssueActualHrsFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueActualHrsFlagNil(b bool)`

 SetProjectIssueActualHrsFlagNil sets the value for ProjectIssueActualHrsFlag to be an explicit nil

### UnsetProjectIssueActualHrsFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueActualHrsFlag()`

UnsetProjectIssueActualHrsFlag ensures that no value is present for ProjectIssueActualHrsFlag, not even an explicit nil
### GetProjectIssueBillFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueBillFlag() bool`

GetProjectIssueBillFlag returns the ProjectIssueBillFlag field if non-nil, zero value otherwise.

### GetProjectIssueBillFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueBillFlagOk() (*bool, bool)`

GetProjectIssueBillFlagOk returns a tuple with the ProjectIssueBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueBillFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueBillFlag(v bool)`

SetProjectIssueBillFlag sets ProjectIssueBillFlag field to given value.

### HasProjectIssueBillFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueBillFlag() bool`

HasProjectIssueBillFlag returns a boolean if a field has been set.

### SetProjectIssueBillFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueBillFlagNil(b bool)`

 SetProjectIssueBillFlagNil sets the value for ProjectIssueBillFlag to be an explicit nil

### UnsetProjectIssueBillFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueBillFlag()`

UnsetProjectIssueBillFlag ensures that no value is present for ProjectIssueBillFlag, not even an explicit nil
### GetProjectIssueStatusFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueStatusFlag() bool`

GetProjectIssueStatusFlag returns the ProjectIssueStatusFlag field if non-nil, zero value otherwise.

### GetProjectIssueStatusFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueStatusFlagOk() (*bool, bool)`

GetProjectIssueStatusFlagOk returns a tuple with the ProjectIssueStatusFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueStatusFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueStatusFlag(v bool)`

SetProjectIssueStatusFlag sets ProjectIssueStatusFlag field to given value.

### HasProjectIssueStatusFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueStatusFlag() bool`

HasProjectIssueStatusFlag returns a boolean if a field has been set.

### SetProjectIssueStatusFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueStatusFlagNil(b bool)`

 SetProjectIssueStatusFlagNil sets the value for ProjectIssueStatusFlag to be an explicit nil

### UnsetProjectIssueStatusFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueStatusFlag()`

UnsetProjectIssueStatusFlag ensures that no value is present for ProjectIssueStatusFlag, not even an explicit nil
### GetProjectIssueAssignedFlag

`func (o *PortalConfigurationProjectSetup) GetProjectIssueAssignedFlag() bool`

GetProjectIssueAssignedFlag returns the ProjectIssueAssignedFlag field if non-nil, zero value otherwise.

### GetProjectIssueAssignedFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectIssueAssignedFlagOk() (*bool, bool)`

GetProjectIssueAssignedFlagOk returns a tuple with the ProjectIssueAssignedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIssueAssignedFlag

`func (o *PortalConfigurationProjectSetup) SetProjectIssueAssignedFlag(v bool)`

SetProjectIssueAssignedFlag sets ProjectIssueAssignedFlag field to given value.

### HasProjectIssueAssignedFlag

`func (o *PortalConfigurationProjectSetup) HasProjectIssueAssignedFlag() bool`

HasProjectIssueAssignedFlag returns a boolean if a field has been set.

### SetProjectIssueAssignedFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectIssueAssignedFlagNil(b bool)`

 SetProjectIssueAssignedFlagNil sets the value for ProjectIssueAssignedFlag to be an explicit nil

### UnsetProjectIssueAssignedFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectIssueAssignedFlag()`

UnsetProjectIssueAssignedFlag ensures that no value is present for ProjectIssueAssignedFlag, not even an explicit nil
### GetProjectDetailTotalHoursFlag

`func (o *PortalConfigurationProjectSetup) GetProjectDetailTotalHoursFlag() bool`

GetProjectDetailTotalHoursFlag returns the ProjectDetailTotalHoursFlag field if non-nil, zero value otherwise.

### GetProjectDetailTotalHoursFlagOk

`func (o *PortalConfigurationProjectSetup) GetProjectDetailTotalHoursFlagOk() (*bool, bool)`

GetProjectDetailTotalHoursFlagOk returns a tuple with the ProjectDetailTotalHoursFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDetailTotalHoursFlag

`func (o *PortalConfigurationProjectSetup) SetProjectDetailTotalHoursFlag(v bool)`

SetProjectDetailTotalHoursFlag sets ProjectDetailTotalHoursFlag field to given value.

### HasProjectDetailTotalHoursFlag

`func (o *PortalConfigurationProjectSetup) HasProjectDetailTotalHoursFlag() bool`

HasProjectDetailTotalHoursFlag returns a boolean if a field has been set.

### SetProjectDetailTotalHoursFlagNil

`func (o *PortalConfigurationProjectSetup) SetProjectDetailTotalHoursFlagNil(b bool)`

 SetProjectDetailTotalHoursFlagNil sets the value for ProjectDetailTotalHoursFlag to be an explicit nil

### UnsetProjectDetailTotalHoursFlag
`func (o *PortalConfigurationProjectSetup) UnsetProjectDetailTotalHoursFlag()`

UnsetProjectDetailTotalHoursFlag ensures that no value is present for ProjectDetailTotalHoursFlag, not even an explicit nil
### GetInfo

`func (o *PortalConfigurationProjectSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PortalConfigurationProjectSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PortalConfigurationProjectSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PortalConfigurationProjectSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


