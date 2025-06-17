# ServiceSignoff

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**VisibleLogoFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyInfoFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingTermsFlag** | Pointer to **NullableBool** |  | [optional] 
**SummaryFlag** | Pointer to **NullableBool** |  | [optional] 
**DiscussionFlag** | Pointer to **NullableBool** |  | [optional] 
**TaskFlag** | Pointer to **NullableBool** | On add/post, if this is set to true but no value is set for task, task is defaulted to ServiceTasks.All. | [optional] 
**Task** | Pointer to **NullableString** | On add/post, if this is set but no value is set for taskFlag, taskFlag is set to true. | [optional] 
**ConfigurationsFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**ResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeFlag** | Pointer to **NullableBool** | On add/post, if any time related flag is set to true, this is also set to true. | [optional] 
**TimeMemberFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeDateFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeStartEndFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeBillFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeHoursFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeRateFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeExtendedAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeWorkTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeAgreementFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeManualFlag** | Pointer to **NullableBool** |  | [optional] 
**TimeManualEntry** | Pointer to **NullableInt32** |  | [optional] 
**TimeTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseFlag** | Pointer to **NullableBool** | On add/post, if any expense related flag is set to true, this is also set to true. | [optional] 
**ExpenseDateFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseMemberFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseTypeFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseBillFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseAgreementFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseNotesFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseManualFlag** | Pointer to **NullableBool** |  | [optional] 
**ExpenseManualEntry** | Pointer to **NullableInt32** |  | [optional] 
**ProductFlag** | Pointer to **NullableBool** | On add/post, if any product related flag is set to true, this is also set to true. | [optional] 
**ProductDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductBillFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductQuantityFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductExtendedAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductAgreementFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductManualFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductManualEntry** | Pointer to **NullableInt32** |  | [optional] 
**ProductTaxFlag** | Pointer to **NullableBool** |  | [optional] 
**TechnicianSignoffFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomerSignoffTextFlag** | Pointer to **NullableBool** | On add/post, if customerSignoffText.Length &gt; 0, this is set to true. | [optional] 
**CustomerSignoffText** | Pointer to **string** |  Max length: 4000; | [optional] 
**CustomerSignoffFieldsFlag** | Pointer to **NullableBool** |  | [optional] 
**BillingMethodsTextFlag** | Pointer to **NullableBool** | On add/post, if billingMethodsText.Length &gt; 0, this is set to true. | [optional] 
**BillingMethodsText** | Pointer to **string** |  Max length: 2000; | [optional] 
**CreditCardFieldsFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFFFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceSignoff

`func NewServiceSignoff(name string, ) *ServiceSignoff`

NewServiceSignoff instantiates a new ServiceSignoff object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSignoffWithDefaults

`func NewServiceSignoffWithDefaults() *ServiceSignoff`

NewServiceSignoffWithDefaults instantiates a new ServiceSignoff object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSignoff) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSignoff) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSignoff) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceSignoff) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceSignoff) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceSignoff) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceSignoff) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *ServiceSignoff) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ServiceSignoff) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ServiceSignoff) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ServiceSignoff) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ServiceSignoff) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ServiceSignoff) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetVisibleLogoFlag

`func (o *ServiceSignoff) GetVisibleLogoFlag() bool`

GetVisibleLogoFlag returns the VisibleLogoFlag field if non-nil, zero value otherwise.

### GetVisibleLogoFlagOk

`func (o *ServiceSignoff) GetVisibleLogoFlagOk() (*bool, bool)`

GetVisibleLogoFlagOk returns a tuple with the VisibleLogoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleLogoFlag

`func (o *ServiceSignoff) SetVisibleLogoFlag(v bool)`

SetVisibleLogoFlag sets VisibleLogoFlag field to given value.

### HasVisibleLogoFlag

`func (o *ServiceSignoff) HasVisibleLogoFlag() bool`

HasVisibleLogoFlag returns a boolean if a field has been set.

### SetVisibleLogoFlagNil

`func (o *ServiceSignoff) SetVisibleLogoFlagNil(b bool)`

 SetVisibleLogoFlagNil sets the value for VisibleLogoFlag to be an explicit nil

### UnsetVisibleLogoFlag
`func (o *ServiceSignoff) UnsetVisibleLogoFlag()`

UnsetVisibleLogoFlag ensures that no value is present for VisibleLogoFlag, not even an explicit nil
### GetCompanyInfoFlag

`func (o *ServiceSignoff) GetCompanyInfoFlag() bool`

GetCompanyInfoFlag returns the CompanyInfoFlag field if non-nil, zero value otherwise.

### GetCompanyInfoFlagOk

`func (o *ServiceSignoff) GetCompanyInfoFlagOk() (*bool, bool)`

GetCompanyInfoFlagOk returns a tuple with the CompanyInfoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyInfoFlag

`func (o *ServiceSignoff) SetCompanyInfoFlag(v bool)`

SetCompanyInfoFlag sets CompanyInfoFlag field to given value.

### HasCompanyInfoFlag

`func (o *ServiceSignoff) HasCompanyInfoFlag() bool`

HasCompanyInfoFlag returns a boolean if a field has been set.

### SetCompanyInfoFlagNil

`func (o *ServiceSignoff) SetCompanyInfoFlagNil(b bool)`

 SetCompanyInfoFlagNil sets the value for CompanyInfoFlag to be an explicit nil

### UnsetCompanyInfoFlag
`func (o *ServiceSignoff) UnsetCompanyInfoFlag()`

UnsetCompanyInfoFlag ensures that no value is present for CompanyInfoFlag, not even an explicit nil
### GetBillingTermsFlag

`func (o *ServiceSignoff) GetBillingTermsFlag() bool`

GetBillingTermsFlag returns the BillingTermsFlag field if non-nil, zero value otherwise.

### GetBillingTermsFlagOk

`func (o *ServiceSignoff) GetBillingTermsFlagOk() (*bool, bool)`

GetBillingTermsFlagOk returns a tuple with the BillingTermsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTermsFlag

`func (o *ServiceSignoff) SetBillingTermsFlag(v bool)`

SetBillingTermsFlag sets BillingTermsFlag field to given value.

### HasBillingTermsFlag

`func (o *ServiceSignoff) HasBillingTermsFlag() bool`

HasBillingTermsFlag returns a boolean if a field has been set.

### SetBillingTermsFlagNil

`func (o *ServiceSignoff) SetBillingTermsFlagNil(b bool)`

 SetBillingTermsFlagNil sets the value for BillingTermsFlag to be an explicit nil

### UnsetBillingTermsFlag
`func (o *ServiceSignoff) UnsetBillingTermsFlag()`

UnsetBillingTermsFlag ensures that no value is present for BillingTermsFlag, not even an explicit nil
### GetSummaryFlag

`func (o *ServiceSignoff) GetSummaryFlag() bool`

GetSummaryFlag returns the SummaryFlag field if non-nil, zero value otherwise.

### GetSummaryFlagOk

`func (o *ServiceSignoff) GetSummaryFlagOk() (*bool, bool)`

GetSummaryFlagOk returns a tuple with the SummaryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryFlag

`func (o *ServiceSignoff) SetSummaryFlag(v bool)`

SetSummaryFlag sets SummaryFlag field to given value.

### HasSummaryFlag

`func (o *ServiceSignoff) HasSummaryFlag() bool`

HasSummaryFlag returns a boolean if a field has been set.

### SetSummaryFlagNil

`func (o *ServiceSignoff) SetSummaryFlagNil(b bool)`

 SetSummaryFlagNil sets the value for SummaryFlag to be an explicit nil

### UnsetSummaryFlag
`func (o *ServiceSignoff) UnsetSummaryFlag()`

UnsetSummaryFlag ensures that no value is present for SummaryFlag, not even an explicit nil
### GetDiscussionFlag

`func (o *ServiceSignoff) GetDiscussionFlag() bool`

GetDiscussionFlag returns the DiscussionFlag field if non-nil, zero value otherwise.

### GetDiscussionFlagOk

`func (o *ServiceSignoff) GetDiscussionFlagOk() (*bool, bool)`

GetDiscussionFlagOk returns a tuple with the DiscussionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionFlag

`func (o *ServiceSignoff) SetDiscussionFlag(v bool)`

SetDiscussionFlag sets DiscussionFlag field to given value.

### HasDiscussionFlag

`func (o *ServiceSignoff) HasDiscussionFlag() bool`

HasDiscussionFlag returns a boolean if a field has been set.

### SetDiscussionFlagNil

`func (o *ServiceSignoff) SetDiscussionFlagNil(b bool)`

 SetDiscussionFlagNil sets the value for DiscussionFlag to be an explicit nil

### UnsetDiscussionFlag
`func (o *ServiceSignoff) UnsetDiscussionFlag()`

UnsetDiscussionFlag ensures that no value is present for DiscussionFlag, not even an explicit nil
### GetTaskFlag

`func (o *ServiceSignoff) GetTaskFlag() bool`

GetTaskFlag returns the TaskFlag field if non-nil, zero value otherwise.

### GetTaskFlagOk

`func (o *ServiceSignoff) GetTaskFlagOk() (*bool, bool)`

GetTaskFlagOk returns a tuple with the TaskFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskFlag

`func (o *ServiceSignoff) SetTaskFlag(v bool)`

SetTaskFlag sets TaskFlag field to given value.

### HasTaskFlag

`func (o *ServiceSignoff) HasTaskFlag() bool`

HasTaskFlag returns a boolean if a field has been set.

### SetTaskFlagNil

`func (o *ServiceSignoff) SetTaskFlagNil(b bool)`

 SetTaskFlagNil sets the value for TaskFlag to be an explicit nil

### UnsetTaskFlag
`func (o *ServiceSignoff) UnsetTaskFlag()`

UnsetTaskFlag ensures that no value is present for TaskFlag, not even an explicit nil
### GetTask

`func (o *ServiceSignoff) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ServiceSignoff) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ServiceSignoff) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ServiceSignoff) HasTask() bool`

HasTask returns a boolean if a field has been set.

### SetTaskNil

`func (o *ServiceSignoff) SetTaskNil(b bool)`

 SetTaskNil sets the value for Task to be an explicit nil

### UnsetTask
`func (o *ServiceSignoff) UnsetTask()`

UnsetTask ensures that no value is present for Task, not even an explicit nil
### GetConfigurationsFlag

`func (o *ServiceSignoff) GetConfigurationsFlag() bool`

GetConfigurationsFlag returns the ConfigurationsFlag field if non-nil, zero value otherwise.

### GetConfigurationsFlagOk

`func (o *ServiceSignoff) GetConfigurationsFlagOk() (*bool, bool)`

GetConfigurationsFlagOk returns a tuple with the ConfigurationsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationsFlag

`func (o *ServiceSignoff) SetConfigurationsFlag(v bool)`

SetConfigurationsFlag sets ConfigurationsFlag field to given value.

### HasConfigurationsFlag

`func (o *ServiceSignoff) HasConfigurationsFlag() bool`

HasConfigurationsFlag returns a boolean if a field has been set.

### SetConfigurationsFlagNil

`func (o *ServiceSignoff) SetConfigurationsFlagNil(b bool)`

 SetConfigurationsFlagNil sets the value for ConfigurationsFlag to be an explicit nil

### UnsetConfigurationsFlag
`func (o *ServiceSignoff) UnsetConfigurationsFlag()`

UnsetConfigurationsFlag ensures that no value is present for ConfigurationsFlag, not even an explicit nil
### GetInternalNotesFlag

`func (o *ServiceSignoff) GetInternalNotesFlag() bool`

GetInternalNotesFlag returns the InternalNotesFlag field if non-nil, zero value otherwise.

### GetInternalNotesFlagOk

`func (o *ServiceSignoff) GetInternalNotesFlagOk() (*bool, bool)`

GetInternalNotesFlagOk returns a tuple with the InternalNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalNotesFlag

`func (o *ServiceSignoff) SetInternalNotesFlag(v bool)`

SetInternalNotesFlag sets InternalNotesFlag field to given value.

### HasInternalNotesFlag

`func (o *ServiceSignoff) HasInternalNotesFlag() bool`

HasInternalNotesFlag returns a boolean if a field has been set.

### SetInternalNotesFlagNil

`func (o *ServiceSignoff) SetInternalNotesFlagNil(b bool)`

 SetInternalNotesFlagNil sets the value for InternalNotesFlag to be an explicit nil

### UnsetInternalNotesFlag
`func (o *ServiceSignoff) UnsetInternalNotesFlag()`

UnsetInternalNotesFlag ensures that no value is present for InternalNotesFlag, not even an explicit nil
### GetResolutionFlag

`func (o *ServiceSignoff) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *ServiceSignoff) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *ServiceSignoff) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *ServiceSignoff) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *ServiceSignoff) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *ServiceSignoff) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetTimeFlag

`func (o *ServiceSignoff) GetTimeFlag() bool`

GetTimeFlag returns the TimeFlag field if non-nil, zero value otherwise.

### GetTimeFlagOk

`func (o *ServiceSignoff) GetTimeFlagOk() (*bool, bool)`

GetTimeFlagOk returns a tuple with the TimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFlag

`func (o *ServiceSignoff) SetTimeFlag(v bool)`

SetTimeFlag sets TimeFlag field to given value.

### HasTimeFlag

`func (o *ServiceSignoff) HasTimeFlag() bool`

HasTimeFlag returns a boolean if a field has been set.

### SetTimeFlagNil

`func (o *ServiceSignoff) SetTimeFlagNil(b bool)`

 SetTimeFlagNil sets the value for TimeFlag to be an explicit nil

### UnsetTimeFlag
`func (o *ServiceSignoff) UnsetTimeFlag()`

UnsetTimeFlag ensures that no value is present for TimeFlag, not even an explicit nil
### GetTimeMemberFlag

`func (o *ServiceSignoff) GetTimeMemberFlag() bool`

GetTimeMemberFlag returns the TimeMemberFlag field if non-nil, zero value otherwise.

### GetTimeMemberFlagOk

`func (o *ServiceSignoff) GetTimeMemberFlagOk() (*bool, bool)`

GetTimeMemberFlagOk returns a tuple with the TimeMemberFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMemberFlag

`func (o *ServiceSignoff) SetTimeMemberFlag(v bool)`

SetTimeMemberFlag sets TimeMemberFlag field to given value.

### HasTimeMemberFlag

`func (o *ServiceSignoff) HasTimeMemberFlag() bool`

HasTimeMemberFlag returns a boolean if a field has been set.

### SetTimeMemberFlagNil

`func (o *ServiceSignoff) SetTimeMemberFlagNil(b bool)`

 SetTimeMemberFlagNil sets the value for TimeMemberFlag to be an explicit nil

### UnsetTimeMemberFlag
`func (o *ServiceSignoff) UnsetTimeMemberFlag()`

UnsetTimeMemberFlag ensures that no value is present for TimeMemberFlag, not even an explicit nil
### GetTimeDateFlag

`func (o *ServiceSignoff) GetTimeDateFlag() bool`

GetTimeDateFlag returns the TimeDateFlag field if non-nil, zero value otherwise.

### GetTimeDateFlagOk

`func (o *ServiceSignoff) GetTimeDateFlagOk() (*bool, bool)`

GetTimeDateFlagOk returns a tuple with the TimeDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDateFlag

`func (o *ServiceSignoff) SetTimeDateFlag(v bool)`

SetTimeDateFlag sets TimeDateFlag field to given value.

### HasTimeDateFlag

`func (o *ServiceSignoff) HasTimeDateFlag() bool`

HasTimeDateFlag returns a boolean if a field has been set.

### SetTimeDateFlagNil

`func (o *ServiceSignoff) SetTimeDateFlagNil(b bool)`

 SetTimeDateFlagNil sets the value for TimeDateFlag to be an explicit nil

### UnsetTimeDateFlag
`func (o *ServiceSignoff) UnsetTimeDateFlag()`

UnsetTimeDateFlag ensures that no value is present for TimeDateFlag, not even an explicit nil
### GetTimeStartEndFlag

`func (o *ServiceSignoff) GetTimeStartEndFlag() bool`

GetTimeStartEndFlag returns the TimeStartEndFlag field if non-nil, zero value otherwise.

### GetTimeStartEndFlagOk

`func (o *ServiceSignoff) GetTimeStartEndFlagOk() (*bool, bool)`

GetTimeStartEndFlagOk returns a tuple with the TimeStartEndFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStartEndFlag

`func (o *ServiceSignoff) SetTimeStartEndFlag(v bool)`

SetTimeStartEndFlag sets TimeStartEndFlag field to given value.

### HasTimeStartEndFlag

`func (o *ServiceSignoff) HasTimeStartEndFlag() bool`

HasTimeStartEndFlag returns a boolean if a field has been set.

### SetTimeStartEndFlagNil

`func (o *ServiceSignoff) SetTimeStartEndFlagNil(b bool)`

 SetTimeStartEndFlagNil sets the value for TimeStartEndFlag to be an explicit nil

### UnsetTimeStartEndFlag
`func (o *ServiceSignoff) UnsetTimeStartEndFlag()`

UnsetTimeStartEndFlag ensures that no value is present for TimeStartEndFlag, not even an explicit nil
### GetTimeBillFlag

`func (o *ServiceSignoff) GetTimeBillFlag() bool`

GetTimeBillFlag returns the TimeBillFlag field if non-nil, zero value otherwise.

### GetTimeBillFlagOk

`func (o *ServiceSignoff) GetTimeBillFlagOk() (*bool, bool)`

GetTimeBillFlagOk returns a tuple with the TimeBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBillFlag

`func (o *ServiceSignoff) SetTimeBillFlag(v bool)`

SetTimeBillFlag sets TimeBillFlag field to given value.

### HasTimeBillFlag

`func (o *ServiceSignoff) HasTimeBillFlag() bool`

HasTimeBillFlag returns a boolean if a field has been set.

### SetTimeBillFlagNil

`func (o *ServiceSignoff) SetTimeBillFlagNil(b bool)`

 SetTimeBillFlagNil sets the value for TimeBillFlag to be an explicit nil

### UnsetTimeBillFlag
`func (o *ServiceSignoff) UnsetTimeBillFlag()`

UnsetTimeBillFlag ensures that no value is present for TimeBillFlag, not even an explicit nil
### GetTimeHoursFlag

`func (o *ServiceSignoff) GetTimeHoursFlag() bool`

GetTimeHoursFlag returns the TimeHoursFlag field if non-nil, zero value otherwise.

### GetTimeHoursFlagOk

`func (o *ServiceSignoff) GetTimeHoursFlagOk() (*bool, bool)`

GetTimeHoursFlagOk returns a tuple with the TimeHoursFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeHoursFlag

`func (o *ServiceSignoff) SetTimeHoursFlag(v bool)`

SetTimeHoursFlag sets TimeHoursFlag field to given value.

### HasTimeHoursFlag

`func (o *ServiceSignoff) HasTimeHoursFlag() bool`

HasTimeHoursFlag returns a boolean if a field has been set.

### SetTimeHoursFlagNil

`func (o *ServiceSignoff) SetTimeHoursFlagNil(b bool)`

 SetTimeHoursFlagNil sets the value for TimeHoursFlag to be an explicit nil

### UnsetTimeHoursFlag
`func (o *ServiceSignoff) UnsetTimeHoursFlag()`

UnsetTimeHoursFlag ensures that no value is present for TimeHoursFlag, not even an explicit nil
### GetTimeRateFlag

`func (o *ServiceSignoff) GetTimeRateFlag() bool`

GetTimeRateFlag returns the TimeRateFlag field if non-nil, zero value otherwise.

### GetTimeRateFlagOk

`func (o *ServiceSignoff) GetTimeRateFlagOk() (*bool, bool)`

GetTimeRateFlagOk returns a tuple with the TimeRateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRateFlag

`func (o *ServiceSignoff) SetTimeRateFlag(v bool)`

SetTimeRateFlag sets TimeRateFlag field to given value.

### HasTimeRateFlag

`func (o *ServiceSignoff) HasTimeRateFlag() bool`

HasTimeRateFlag returns a boolean if a field has been set.

### SetTimeRateFlagNil

`func (o *ServiceSignoff) SetTimeRateFlagNil(b bool)`

 SetTimeRateFlagNil sets the value for TimeRateFlag to be an explicit nil

### UnsetTimeRateFlag
`func (o *ServiceSignoff) UnsetTimeRateFlag()`

UnsetTimeRateFlag ensures that no value is present for TimeRateFlag, not even an explicit nil
### GetTimeExtendedAmountFlag

`func (o *ServiceSignoff) GetTimeExtendedAmountFlag() bool`

GetTimeExtendedAmountFlag returns the TimeExtendedAmountFlag field if non-nil, zero value otherwise.

### GetTimeExtendedAmountFlagOk

`func (o *ServiceSignoff) GetTimeExtendedAmountFlagOk() (*bool, bool)`

GetTimeExtendedAmountFlagOk returns a tuple with the TimeExtendedAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeExtendedAmountFlag

`func (o *ServiceSignoff) SetTimeExtendedAmountFlag(v bool)`

SetTimeExtendedAmountFlag sets TimeExtendedAmountFlag field to given value.

### HasTimeExtendedAmountFlag

`func (o *ServiceSignoff) HasTimeExtendedAmountFlag() bool`

HasTimeExtendedAmountFlag returns a boolean if a field has been set.

### SetTimeExtendedAmountFlagNil

`func (o *ServiceSignoff) SetTimeExtendedAmountFlagNil(b bool)`

 SetTimeExtendedAmountFlagNil sets the value for TimeExtendedAmountFlag to be an explicit nil

### UnsetTimeExtendedAmountFlag
`func (o *ServiceSignoff) UnsetTimeExtendedAmountFlag()`

UnsetTimeExtendedAmountFlag ensures that no value is present for TimeExtendedAmountFlag, not even an explicit nil
### GetTimeWorkTypeFlag

`func (o *ServiceSignoff) GetTimeWorkTypeFlag() bool`

GetTimeWorkTypeFlag returns the TimeWorkTypeFlag field if non-nil, zero value otherwise.

### GetTimeWorkTypeFlagOk

`func (o *ServiceSignoff) GetTimeWorkTypeFlagOk() (*bool, bool)`

GetTimeWorkTypeFlagOk returns a tuple with the TimeWorkTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeWorkTypeFlag

`func (o *ServiceSignoff) SetTimeWorkTypeFlag(v bool)`

SetTimeWorkTypeFlag sets TimeWorkTypeFlag field to given value.

### HasTimeWorkTypeFlag

`func (o *ServiceSignoff) HasTimeWorkTypeFlag() bool`

HasTimeWorkTypeFlag returns a boolean if a field has been set.

### SetTimeWorkTypeFlagNil

`func (o *ServiceSignoff) SetTimeWorkTypeFlagNil(b bool)`

 SetTimeWorkTypeFlagNil sets the value for TimeWorkTypeFlag to be an explicit nil

### UnsetTimeWorkTypeFlag
`func (o *ServiceSignoff) UnsetTimeWorkTypeFlag()`

UnsetTimeWorkTypeFlag ensures that no value is present for TimeWorkTypeFlag, not even an explicit nil
### GetTimeAgreementFlag

`func (o *ServiceSignoff) GetTimeAgreementFlag() bool`

GetTimeAgreementFlag returns the TimeAgreementFlag field if non-nil, zero value otherwise.

### GetTimeAgreementFlagOk

`func (o *ServiceSignoff) GetTimeAgreementFlagOk() (*bool, bool)`

GetTimeAgreementFlagOk returns a tuple with the TimeAgreementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAgreementFlag

`func (o *ServiceSignoff) SetTimeAgreementFlag(v bool)`

SetTimeAgreementFlag sets TimeAgreementFlag field to given value.

### HasTimeAgreementFlag

`func (o *ServiceSignoff) HasTimeAgreementFlag() bool`

HasTimeAgreementFlag returns a boolean if a field has been set.

### SetTimeAgreementFlagNil

`func (o *ServiceSignoff) SetTimeAgreementFlagNil(b bool)`

 SetTimeAgreementFlagNil sets the value for TimeAgreementFlag to be an explicit nil

### UnsetTimeAgreementFlag
`func (o *ServiceSignoff) UnsetTimeAgreementFlag()`

UnsetTimeAgreementFlag ensures that no value is present for TimeAgreementFlag, not even an explicit nil
### GetTimeNotesFlag

`func (o *ServiceSignoff) GetTimeNotesFlag() bool`

GetTimeNotesFlag returns the TimeNotesFlag field if non-nil, zero value otherwise.

### GetTimeNotesFlagOk

`func (o *ServiceSignoff) GetTimeNotesFlagOk() (*bool, bool)`

GetTimeNotesFlagOk returns a tuple with the TimeNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeNotesFlag

`func (o *ServiceSignoff) SetTimeNotesFlag(v bool)`

SetTimeNotesFlag sets TimeNotesFlag field to given value.

### HasTimeNotesFlag

`func (o *ServiceSignoff) HasTimeNotesFlag() bool`

HasTimeNotesFlag returns a boolean if a field has been set.

### SetTimeNotesFlagNil

`func (o *ServiceSignoff) SetTimeNotesFlagNil(b bool)`

 SetTimeNotesFlagNil sets the value for TimeNotesFlag to be an explicit nil

### UnsetTimeNotesFlag
`func (o *ServiceSignoff) UnsetTimeNotesFlag()`

UnsetTimeNotesFlag ensures that no value is present for TimeNotesFlag, not even an explicit nil
### GetTimeManualFlag

`func (o *ServiceSignoff) GetTimeManualFlag() bool`

GetTimeManualFlag returns the TimeManualFlag field if non-nil, zero value otherwise.

### GetTimeManualFlagOk

`func (o *ServiceSignoff) GetTimeManualFlagOk() (*bool, bool)`

GetTimeManualFlagOk returns a tuple with the TimeManualFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeManualFlag

`func (o *ServiceSignoff) SetTimeManualFlag(v bool)`

SetTimeManualFlag sets TimeManualFlag field to given value.

### HasTimeManualFlag

`func (o *ServiceSignoff) HasTimeManualFlag() bool`

HasTimeManualFlag returns a boolean if a field has been set.

### SetTimeManualFlagNil

`func (o *ServiceSignoff) SetTimeManualFlagNil(b bool)`

 SetTimeManualFlagNil sets the value for TimeManualFlag to be an explicit nil

### UnsetTimeManualFlag
`func (o *ServiceSignoff) UnsetTimeManualFlag()`

UnsetTimeManualFlag ensures that no value is present for TimeManualFlag, not even an explicit nil
### GetTimeManualEntry

`func (o *ServiceSignoff) GetTimeManualEntry() int32`

GetTimeManualEntry returns the TimeManualEntry field if non-nil, zero value otherwise.

### GetTimeManualEntryOk

`func (o *ServiceSignoff) GetTimeManualEntryOk() (*int32, bool)`

GetTimeManualEntryOk returns a tuple with the TimeManualEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeManualEntry

`func (o *ServiceSignoff) SetTimeManualEntry(v int32)`

SetTimeManualEntry sets TimeManualEntry field to given value.

### HasTimeManualEntry

`func (o *ServiceSignoff) HasTimeManualEntry() bool`

HasTimeManualEntry returns a boolean if a field has been set.

### SetTimeManualEntryNil

`func (o *ServiceSignoff) SetTimeManualEntryNil(b bool)`

 SetTimeManualEntryNil sets the value for TimeManualEntry to be an explicit nil

### UnsetTimeManualEntry
`func (o *ServiceSignoff) UnsetTimeManualEntry()`

UnsetTimeManualEntry ensures that no value is present for TimeManualEntry, not even an explicit nil
### GetTimeTaxFlag

`func (o *ServiceSignoff) GetTimeTaxFlag() bool`

GetTimeTaxFlag returns the TimeTaxFlag field if non-nil, zero value otherwise.

### GetTimeTaxFlagOk

`func (o *ServiceSignoff) GetTimeTaxFlagOk() (*bool, bool)`

GetTimeTaxFlagOk returns a tuple with the TimeTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTaxFlag

`func (o *ServiceSignoff) SetTimeTaxFlag(v bool)`

SetTimeTaxFlag sets TimeTaxFlag field to given value.

### HasTimeTaxFlag

`func (o *ServiceSignoff) HasTimeTaxFlag() bool`

HasTimeTaxFlag returns a boolean if a field has been set.

### SetTimeTaxFlagNil

`func (o *ServiceSignoff) SetTimeTaxFlagNil(b bool)`

 SetTimeTaxFlagNil sets the value for TimeTaxFlag to be an explicit nil

### UnsetTimeTaxFlag
`func (o *ServiceSignoff) UnsetTimeTaxFlag()`

UnsetTimeTaxFlag ensures that no value is present for TimeTaxFlag, not even an explicit nil
### GetExpenseFlag

`func (o *ServiceSignoff) GetExpenseFlag() bool`

GetExpenseFlag returns the ExpenseFlag field if non-nil, zero value otherwise.

### GetExpenseFlagOk

`func (o *ServiceSignoff) GetExpenseFlagOk() (*bool, bool)`

GetExpenseFlagOk returns a tuple with the ExpenseFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseFlag

`func (o *ServiceSignoff) SetExpenseFlag(v bool)`

SetExpenseFlag sets ExpenseFlag field to given value.

### HasExpenseFlag

`func (o *ServiceSignoff) HasExpenseFlag() bool`

HasExpenseFlag returns a boolean if a field has been set.

### SetExpenseFlagNil

`func (o *ServiceSignoff) SetExpenseFlagNil(b bool)`

 SetExpenseFlagNil sets the value for ExpenseFlag to be an explicit nil

### UnsetExpenseFlag
`func (o *ServiceSignoff) UnsetExpenseFlag()`

UnsetExpenseFlag ensures that no value is present for ExpenseFlag, not even an explicit nil
### GetExpenseDateFlag

`func (o *ServiceSignoff) GetExpenseDateFlag() bool`

GetExpenseDateFlag returns the ExpenseDateFlag field if non-nil, zero value otherwise.

### GetExpenseDateFlagOk

`func (o *ServiceSignoff) GetExpenseDateFlagOk() (*bool, bool)`

GetExpenseDateFlagOk returns a tuple with the ExpenseDateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseDateFlag

`func (o *ServiceSignoff) SetExpenseDateFlag(v bool)`

SetExpenseDateFlag sets ExpenseDateFlag field to given value.

### HasExpenseDateFlag

`func (o *ServiceSignoff) HasExpenseDateFlag() bool`

HasExpenseDateFlag returns a boolean if a field has been set.

### SetExpenseDateFlagNil

`func (o *ServiceSignoff) SetExpenseDateFlagNil(b bool)`

 SetExpenseDateFlagNil sets the value for ExpenseDateFlag to be an explicit nil

### UnsetExpenseDateFlag
`func (o *ServiceSignoff) UnsetExpenseDateFlag()`

UnsetExpenseDateFlag ensures that no value is present for ExpenseDateFlag, not even an explicit nil
### GetExpenseMemberFlag

`func (o *ServiceSignoff) GetExpenseMemberFlag() bool`

GetExpenseMemberFlag returns the ExpenseMemberFlag field if non-nil, zero value otherwise.

### GetExpenseMemberFlagOk

`func (o *ServiceSignoff) GetExpenseMemberFlagOk() (*bool, bool)`

GetExpenseMemberFlagOk returns a tuple with the ExpenseMemberFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseMemberFlag

`func (o *ServiceSignoff) SetExpenseMemberFlag(v bool)`

SetExpenseMemberFlag sets ExpenseMemberFlag field to given value.

### HasExpenseMemberFlag

`func (o *ServiceSignoff) HasExpenseMemberFlag() bool`

HasExpenseMemberFlag returns a boolean if a field has been set.

### SetExpenseMemberFlagNil

`func (o *ServiceSignoff) SetExpenseMemberFlagNil(b bool)`

 SetExpenseMemberFlagNil sets the value for ExpenseMemberFlag to be an explicit nil

### UnsetExpenseMemberFlag
`func (o *ServiceSignoff) UnsetExpenseMemberFlag()`

UnsetExpenseMemberFlag ensures that no value is present for ExpenseMemberFlag, not even an explicit nil
### GetExpenseTypeFlag

`func (o *ServiceSignoff) GetExpenseTypeFlag() bool`

GetExpenseTypeFlag returns the ExpenseTypeFlag field if non-nil, zero value otherwise.

### GetExpenseTypeFlagOk

`func (o *ServiceSignoff) GetExpenseTypeFlagOk() (*bool, bool)`

GetExpenseTypeFlagOk returns a tuple with the ExpenseTypeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTypeFlag

`func (o *ServiceSignoff) SetExpenseTypeFlag(v bool)`

SetExpenseTypeFlag sets ExpenseTypeFlag field to given value.

### HasExpenseTypeFlag

`func (o *ServiceSignoff) HasExpenseTypeFlag() bool`

HasExpenseTypeFlag returns a boolean if a field has been set.

### SetExpenseTypeFlagNil

`func (o *ServiceSignoff) SetExpenseTypeFlagNil(b bool)`

 SetExpenseTypeFlagNil sets the value for ExpenseTypeFlag to be an explicit nil

### UnsetExpenseTypeFlag
`func (o *ServiceSignoff) UnsetExpenseTypeFlag()`

UnsetExpenseTypeFlag ensures that no value is present for ExpenseTypeFlag, not even an explicit nil
### GetExpenseBillFlag

`func (o *ServiceSignoff) GetExpenseBillFlag() bool`

GetExpenseBillFlag returns the ExpenseBillFlag field if non-nil, zero value otherwise.

### GetExpenseBillFlagOk

`func (o *ServiceSignoff) GetExpenseBillFlagOk() (*bool, bool)`

GetExpenseBillFlagOk returns a tuple with the ExpenseBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseBillFlag

`func (o *ServiceSignoff) SetExpenseBillFlag(v bool)`

SetExpenseBillFlag sets ExpenseBillFlag field to given value.

### HasExpenseBillFlag

`func (o *ServiceSignoff) HasExpenseBillFlag() bool`

HasExpenseBillFlag returns a boolean if a field has been set.

### SetExpenseBillFlagNil

`func (o *ServiceSignoff) SetExpenseBillFlagNil(b bool)`

 SetExpenseBillFlagNil sets the value for ExpenseBillFlag to be an explicit nil

### UnsetExpenseBillFlag
`func (o *ServiceSignoff) UnsetExpenseBillFlag()`

UnsetExpenseBillFlag ensures that no value is present for ExpenseBillFlag, not even an explicit nil
### GetExpenseAmountFlag

`func (o *ServiceSignoff) GetExpenseAmountFlag() bool`

GetExpenseAmountFlag returns the ExpenseAmountFlag field if non-nil, zero value otherwise.

### GetExpenseAmountFlagOk

`func (o *ServiceSignoff) GetExpenseAmountFlagOk() (*bool, bool)`

GetExpenseAmountFlagOk returns a tuple with the ExpenseAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAmountFlag

`func (o *ServiceSignoff) SetExpenseAmountFlag(v bool)`

SetExpenseAmountFlag sets ExpenseAmountFlag field to given value.

### HasExpenseAmountFlag

`func (o *ServiceSignoff) HasExpenseAmountFlag() bool`

HasExpenseAmountFlag returns a boolean if a field has been set.

### SetExpenseAmountFlagNil

`func (o *ServiceSignoff) SetExpenseAmountFlagNil(b bool)`

 SetExpenseAmountFlagNil sets the value for ExpenseAmountFlag to be an explicit nil

### UnsetExpenseAmountFlag
`func (o *ServiceSignoff) UnsetExpenseAmountFlag()`

UnsetExpenseAmountFlag ensures that no value is present for ExpenseAmountFlag, not even an explicit nil
### GetExpenseAgreementFlag

`func (o *ServiceSignoff) GetExpenseAgreementFlag() bool`

GetExpenseAgreementFlag returns the ExpenseAgreementFlag field if non-nil, zero value otherwise.

### GetExpenseAgreementFlagOk

`func (o *ServiceSignoff) GetExpenseAgreementFlagOk() (*bool, bool)`

GetExpenseAgreementFlagOk returns a tuple with the ExpenseAgreementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseAgreementFlag

`func (o *ServiceSignoff) SetExpenseAgreementFlag(v bool)`

SetExpenseAgreementFlag sets ExpenseAgreementFlag field to given value.

### HasExpenseAgreementFlag

`func (o *ServiceSignoff) HasExpenseAgreementFlag() bool`

HasExpenseAgreementFlag returns a boolean if a field has been set.

### SetExpenseAgreementFlagNil

`func (o *ServiceSignoff) SetExpenseAgreementFlagNil(b bool)`

 SetExpenseAgreementFlagNil sets the value for ExpenseAgreementFlag to be an explicit nil

### UnsetExpenseAgreementFlag
`func (o *ServiceSignoff) UnsetExpenseAgreementFlag()`

UnsetExpenseAgreementFlag ensures that no value is present for ExpenseAgreementFlag, not even an explicit nil
### GetExpenseNotesFlag

`func (o *ServiceSignoff) GetExpenseNotesFlag() bool`

GetExpenseNotesFlag returns the ExpenseNotesFlag field if non-nil, zero value otherwise.

### GetExpenseNotesFlagOk

`func (o *ServiceSignoff) GetExpenseNotesFlagOk() (*bool, bool)`

GetExpenseNotesFlagOk returns a tuple with the ExpenseNotesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseNotesFlag

`func (o *ServiceSignoff) SetExpenseNotesFlag(v bool)`

SetExpenseNotesFlag sets ExpenseNotesFlag field to given value.

### HasExpenseNotesFlag

`func (o *ServiceSignoff) HasExpenseNotesFlag() bool`

HasExpenseNotesFlag returns a boolean if a field has been set.

### SetExpenseNotesFlagNil

`func (o *ServiceSignoff) SetExpenseNotesFlagNil(b bool)`

 SetExpenseNotesFlagNil sets the value for ExpenseNotesFlag to be an explicit nil

### UnsetExpenseNotesFlag
`func (o *ServiceSignoff) UnsetExpenseNotesFlag()`

UnsetExpenseNotesFlag ensures that no value is present for ExpenseNotesFlag, not even an explicit nil
### GetExpenseTaxFlag

`func (o *ServiceSignoff) GetExpenseTaxFlag() bool`

GetExpenseTaxFlag returns the ExpenseTaxFlag field if non-nil, zero value otherwise.

### GetExpenseTaxFlagOk

`func (o *ServiceSignoff) GetExpenseTaxFlagOk() (*bool, bool)`

GetExpenseTaxFlagOk returns a tuple with the ExpenseTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseTaxFlag

`func (o *ServiceSignoff) SetExpenseTaxFlag(v bool)`

SetExpenseTaxFlag sets ExpenseTaxFlag field to given value.

### HasExpenseTaxFlag

`func (o *ServiceSignoff) HasExpenseTaxFlag() bool`

HasExpenseTaxFlag returns a boolean if a field has been set.

### SetExpenseTaxFlagNil

`func (o *ServiceSignoff) SetExpenseTaxFlagNil(b bool)`

 SetExpenseTaxFlagNil sets the value for ExpenseTaxFlag to be an explicit nil

### UnsetExpenseTaxFlag
`func (o *ServiceSignoff) UnsetExpenseTaxFlag()`

UnsetExpenseTaxFlag ensures that no value is present for ExpenseTaxFlag, not even an explicit nil
### GetExpenseManualFlag

`func (o *ServiceSignoff) GetExpenseManualFlag() bool`

GetExpenseManualFlag returns the ExpenseManualFlag field if non-nil, zero value otherwise.

### GetExpenseManualFlagOk

`func (o *ServiceSignoff) GetExpenseManualFlagOk() (*bool, bool)`

GetExpenseManualFlagOk returns a tuple with the ExpenseManualFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseManualFlag

`func (o *ServiceSignoff) SetExpenseManualFlag(v bool)`

SetExpenseManualFlag sets ExpenseManualFlag field to given value.

### HasExpenseManualFlag

`func (o *ServiceSignoff) HasExpenseManualFlag() bool`

HasExpenseManualFlag returns a boolean if a field has been set.

### SetExpenseManualFlagNil

`func (o *ServiceSignoff) SetExpenseManualFlagNil(b bool)`

 SetExpenseManualFlagNil sets the value for ExpenseManualFlag to be an explicit nil

### UnsetExpenseManualFlag
`func (o *ServiceSignoff) UnsetExpenseManualFlag()`

UnsetExpenseManualFlag ensures that no value is present for ExpenseManualFlag, not even an explicit nil
### GetExpenseManualEntry

`func (o *ServiceSignoff) GetExpenseManualEntry() int32`

GetExpenseManualEntry returns the ExpenseManualEntry field if non-nil, zero value otherwise.

### GetExpenseManualEntryOk

`func (o *ServiceSignoff) GetExpenseManualEntryOk() (*int32, bool)`

GetExpenseManualEntryOk returns a tuple with the ExpenseManualEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseManualEntry

`func (o *ServiceSignoff) SetExpenseManualEntry(v int32)`

SetExpenseManualEntry sets ExpenseManualEntry field to given value.

### HasExpenseManualEntry

`func (o *ServiceSignoff) HasExpenseManualEntry() bool`

HasExpenseManualEntry returns a boolean if a field has been set.

### SetExpenseManualEntryNil

`func (o *ServiceSignoff) SetExpenseManualEntryNil(b bool)`

 SetExpenseManualEntryNil sets the value for ExpenseManualEntry to be an explicit nil

### UnsetExpenseManualEntry
`func (o *ServiceSignoff) UnsetExpenseManualEntry()`

UnsetExpenseManualEntry ensures that no value is present for ExpenseManualEntry, not even an explicit nil
### GetProductFlag

`func (o *ServiceSignoff) GetProductFlag() bool`

GetProductFlag returns the ProductFlag field if non-nil, zero value otherwise.

### GetProductFlagOk

`func (o *ServiceSignoff) GetProductFlagOk() (*bool, bool)`

GetProductFlagOk returns a tuple with the ProductFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductFlag

`func (o *ServiceSignoff) SetProductFlag(v bool)`

SetProductFlag sets ProductFlag field to given value.

### HasProductFlag

`func (o *ServiceSignoff) HasProductFlag() bool`

HasProductFlag returns a boolean if a field has been set.

### SetProductFlagNil

`func (o *ServiceSignoff) SetProductFlagNil(b bool)`

 SetProductFlagNil sets the value for ProductFlag to be an explicit nil

### UnsetProductFlag
`func (o *ServiceSignoff) UnsetProductFlag()`

UnsetProductFlag ensures that no value is present for ProductFlag, not even an explicit nil
### GetProductDescriptionFlag

`func (o *ServiceSignoff) GetProductDescriptionFlag() bool`

GetProductDescriptionFlag returns the ProductDescriptionFlag field if non-nil, zero value otherwise.

### GetProductDescriptionFlagOk

`func (o *ServiceSignoff) GetProductDescriptionFlagOk() (*bool, bool)`

GetProductDescriptionFlagOk returns a tuple with the ProductDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescriptionFlag

`func (o *ServiceSignoff) SetProductDescriptionFlag(v bool)`

SetProductDescriptionFlag sets ProductDescriptionFlag field to given value.

### HasProductDescriptionFlag

`func (o *ServiceSignoff) HasProductDescriptionFlag() bool`

HasProductDescriptionFlag returns a boolean if a field has been set.

### SetProductDescriptionFlagNil

`func (o *ServiceSignoff) SetProductDescriptionFlagNil(b bool)`

 SetProductDescriptionFlagNil sets the value for ProductDescriptionFlag to be an explicit nil

### UnsetProductDescriptionFlag
`func (o *ServiceSignoff) UnsetProductDescriptionFlag()`

UnsetProductDescriptionFlag ensures that no value is present for ProductDescriptionFlag, not even an explicit nil
### GetProductBillFlag

`func (o *ServiceSignoff) GetProductBillFlag() bool`

GetProductBillFlag returns the ProductBillFlag field if non-nil, zero value otherwise.

### GetProductBillFlagOk

`func (o *ServiceSignoff) GetProductBillFlagOk() (*bool, bool)`

GetProductBillFlagOk returns a tuple with the ProductBillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductBillFlag

`func (o *ServiceSignoff) SetProductBillFlag(v bool)`

SetProductBillFlag sets ProductBillFlag field to given value.

### HasProductBillFlag

`func (o *ServiceSignoff) HasProductBillFlag() bool`

HasProductBillFlag returns a boolean if a field has been set.

### SetProductBillFlagNil

`func (o *ServiceSignoff) SetProductBillFlagNil(b bool)`

 SetProductBillFlagNil sets the value for ProductBillFlag to be an explicit nil

### UnsetProductBillFlag
`func (o *ServiceSignoff) UnsetProductBillFlag()`

UnsetProductBillFlag ensures that no value is present for ProductBillFlag, not even an explicit nil
### GetProductQuantityFlag

`func (o *ServiceSignoff) GetProductQuantityFlag() bool`

GetProductQuantityFlag returns the ProductQuantityFlag field if non-nil, zero value otherwise.

### GetProductQuantityFlagOk

`func (o *ServiceSignoff) GetProductQuantityFlagOk() (*bool, bool)`

GetProductQuantityFlagOk returns a tuple with the ProductQuantityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductQuantityFlag

`func (o *ServiceSignoff) SetProductQuantityFlag(v bool)`

SetProductQuantityFlag sets ProductQuantityFlag field to given value.

### HasProductQuantityFlag

`func (o *ServiceSignoff) HasProductQuantityFlag() bool`

HasProductQuantityFlag returns a boolean if a field has been set.

### SetProductQuantityFlagNil

`func (o *ServiceSignoff) SetProductQuantityFlagNil(b bool)`

 SetProductQuantityFlagNil sets the value for ProductQuantityFlag to be an explicit nil

### UnsetProductQuantityFlag
`func (o *ServiceSignoff) UnsetProductQuantityFlag()`

UnsetProductQuantityFlag ensures that no value is present for ProductQuantityFlag, not even an explicit nil
### GetProductPriceFlag

`func (o *ServiceSignoff) GetProductPriceFlag() bool`

GetProductPriceFlag returns the ProductPriceFlag field if non-nil, zero value otherwise.

### GetProductPriceFlagOk

`func (o *ServiceSignoff) GetProductPriceFlagOk() (*bool, bool)`

GetProductPriceFlagOk returns a tuple with the ProductPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductPriceFlag

`func (o *ServiceSignoff) SetProductPriceFlag(v bool)`

SetProductPriceFlag sets ProductPriceFlag field to given value.

### HasProductPriceFlag

`func (o *ServiceSignoff) HasProductPriceFlag() bool`

HasProductPriceFlag returns a boolean if a field has been set.

### SetProductPriceFlagNil

`func (o *ServiceSignoff) SetProductPriceFlagNil(b bool)`

 SetProductPriceFlagNil sets the value for ProductPriceFlag to be an explicit nil

### UnsetProductPriceFlag
`func (o *ServiceSignoff) UnsetProductPriceFlag()`

UnsetProductPriceFlag ensures that no value is present for ProductPriceFlag, not even an explicit nil
### GetProductExtendedAmountFlag

`func (o *ServiceSignoff) GetProductExtendedAmountFlag() bool`

GetProductExtendedAmountFlag returns the ProductExtendedAmountFlag field if non-nil, zero value otherwise.

### GetProductExtendedAmountFlagOk

`func (o *ServiceSignoff) GetProductExtendedAmountFlagOk() (*bool, bool)`

GetProductExtendedAmountFlagOk returns a tuple with the ProductExtendedAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductExtendedAmountFlag

`func (o *ServiceSignoff) SetProductExtendedAmountFlag(v bool)`

SetProductExtendedAmountFlag sets ProductExtendedAmountFlag field to given value.

### HasProductExtendedAmountFlag

`func (o *ServiceSignoff) HasProductExtendedAmountFlag() bool`

HasProductExtendedAmountFlag returns a boolean if a field has been set.

### SetProductExtendedAmountFlagNil

`func (o *ServiceSignoff) SetProductExtendedAmountFlagNil(b bool)`

 SetProductExtendedAmountFlagNil sets the value for ProductExtendedAmountFlag to be an explicit nil

### UnsetProductExtendedAmountFlag
`func (o *ServiceSignoff) UnsetProductExtendedAmountFlag()`

UnsetProductExtendedAmountFlag ensures that no value is present for ProductExtendedAmountFlag, not even an explicit nil
### GetProductAgreementFlag

`func (o *ServiceSignoff) GetProductAgreementFlag() bool`

GetProductAgreementFlag returns the ProductAgreementFlag field if non-nil, zero value otherwise.

### GetProductAgreementFlagOk

`func (o *ServiceSignoff) GetProductAgreementFlagOk() (*bool, bool)`

GetProductAgreementFlagOk returns a tuple with the ProductAgreementFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductAgreementFlag

`func (o *ServiceSignoff) SetProductAgreementFlag(v bool)`

SetProductAgreementFlag sets ProductAgreementFlag field to given value.

### HasProductAgreementFlag

`func (o *ServiceSignoff) HasProductAgreementFlag() bool`

HasProductAgreementFlag returns a boolean if a field has been set.

### SetProductAgreementFlagNil

`func (o *ServiceSignoff) SetProductAgreementFlagNil(b bool)`

 SetProductAgreementFlagNil sets the value for ProductAgreementFlag to be an explicit nil

### UnsetProductAgreementFlag
`func (o *ServiceSignoff) UnsetProductAgreementFlag()`

UnsetProductAgreementFlag ensures that no value is present for ProductAgreementFlag, not even an explicit nil
### GetProductManualFlag

`func (o *ServiceSignoff) GetProductManualFlag() bool`

GetProductManualFlag returns the ProductManualFlag field if non-nil, zero value otherwise.

### GetProductManualFlagOk

`func (o *ServiceSignoff) GetProductManualFlagOk() (*bool, bool)`

GetProductManualFlagOk returns a tuple with the ProductManualFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManualFlag

`func (o *ServiceSignoff) SetProductManualFlag(v bool)`

SetProductManualFlag sets ProductManualFlag field to given value.

### HasProductManualFlag

`func (o *ServiceSignoff) HasProductManualFlag() bool`

HasProductManualFlag returns a boolean if a field has been set.

### SetProductManualFlagNil

`func (o *ServiceSignoff) SetProductManualFlagNil(b bool)`

 SetProductManualFlagNil sets the value for ProductManualFlag to be an explicit nil

### UnsetProductManualFlag
`func (o *ServiceSignoff) UnsetProductManualFlag()`

UnsetProductManualFlag ensures that no value is present for ProductManualFlag, not even an explicit nil
### GetProductManualEntry

`func (o *ServiceSignoff) GetProductManualEntry() int32`

GetProductManualEntry returns the ProductManualEntry field if non-nil, zero value otherwise.

### GetProductManualEntryOk

`func (o *ServiceSignoff) GetProductManualEntryOk() (*int32, bool)`

GetProductManualEntryOk returns a tuple with the ProductManualEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductManualEntry

`func (o *ServiceSignoff) SetProductManualEntry(v int32)`

SetProductManualEntry sets ProductManualEntry field to given value.

### HasProductManualEntry

`func (o *ServiceSignoff) HasProductManualEntry() bool`

HasProductManualEntry returns a boolean if a field has been set.

### SetProductManualEntryNil

`func (o *ServiceSignoff) SetProductManualEntryNil(b bool)`

 SetProductManualEntryNil sets the value for ProductManualEntry to be an explicit nil

### UnsetProductManualEntry
`func (o *ServiceSignoff) UnsetProductManualEntry()`

UnsetProductManualEntry ensures that no value is present for ProductManualEntry, not even an explicit nil
### GetProductTaxFlag

`func (o *ServiceSignoff) GetProductTaxFlag() bool`

GetProductTaxFlag returns the ProductTaxFlag field if non-nil, zero value otherwise.

### GetProductTaxFlagOk

`func (o *ServiceSignoff) GetProductTaxFlagOk() (*bool, bool)`

GetProductTaxFlagOk returns a tuple with the ProductTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductTaxFlag

`func (o *ServiceSignoff) SetProductTaxFlag(v bool)`

SetProductTaxFlag sets ProductTaxFlag field to given value.

### HasProductTaxFlag

`func (o *ServiceSignoff) HasProductTaxFlag() bool`

HasProductTaxFlag returns a boolean if a field has been set.

### SetProductTaxFlagNil

`func (o *ServiceSignoff) SetProductTaxFlagNil(b bool)`

 SetProductTaxFlagNil sets the value for ProductTaxFlag to be an explicit nil

### UnsetProductTaxFlag
`func (o *ServiceSignoff) UnsetProductTaxFlag()`

UnsetProductTaxFlag ensures that no value is present for ProductTaxFlag, not even an explicit nil
### GetTechnicianSignoffFlag

`func (o *ServiceSignoff) GetTechnicianSignoffFlag() bool`

GetTechnicianSignoffFlag returns the TechnicianSignoffFlag field if non-nil, zero value otherwise.

### GetTechnicianSignoffFlagOk

`func (o *ServiceSignoff) GetTechnicianSignoffFlagOk() (*bool, bool)`

GetTechnicianSignoffFlagOk returns a tuple with the TechnicianSignoffFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicianSignoffFlag

`func (o *ServiceSignoff) SetTechnicianSignoffFlag(v bool)`

SetTechnicianSignoffFlag sets TechnicianSignoffFlag field to given value.

### HasTechnicianSignoffFlag

`func (o *ServiceSignoff) HasTechnicianSignoffFlag() bool`

HasTechnicianSignoffFlag returns a boolean if a field has been set.

### SetTechnicianSignoffFlagNil

`func (o *ServiceSignoff) SetTechnicianSignoffFlagNil(b bool)`

 SetTechnicianSignoffFlagNil sets the value for TechnicianSignoffFlag to be an explicit nil

### UnsetTechnicianSignoffFlag
`func (o *ServiceSignoff) UnsetTechnicianSignoffFlag()`

UnsetTechnicianSignoffFlag ensures that no value is present for TechnicianSignoffFlag, not even an explicit nil
### GetCustomerSignoffTextFlag

`func (o *ServiceSignoff) GetCustomerSignoffTextFlag() bool`

GetCustomerSignoffTextFlag returns the CustomerSignoffTextFlag field if non-nil, zero value otherwise.

### GetCustomerSignoffTextFlagOk

`func (o *ServiceSignoff) GetCustomerSignoffTextFlagOk() (*bool, bool)`

GetCustomerSignoffTextFlagOk returns a tuple with the CustomerSignoffTextFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSignoffTextFlag

`func (o *ServiceSignoff) SetCustomerSignoffTextFlag(v bool)`

SetCustomerSignoffTextFlag sets CustomerSignoffTextFlag field to given value.

### HasCustomerSignoffTextFlag

`func (o *ServiceSignoff) HasCustomerSignoffTextFlag() bool`

HasCustomerSignoffTextFlag returns a boolean if a field has been set.

### SetCustomerSignoffTextFlagNil

`func (o *ServiceSignoff) SetCustomerSignoffTextFlagNil(b bool)`

 SetCustomerSignoffTextFlagNil sets the value for CustomerSignoffTextFlag to be an explicit nil

### UnsetCustomerSignoffTextFlag
`func (o *ServiceSignoff) UnsetCustomerSignoffTextFlag()`

UnsetCustomerSignoffTextFlag ensures that no value is present for CustomerSignoffTextFlag, not even an explicit nil
### GetCustomerSignoffText

`func (o *ServiceSignoff) GetCustomerSignoffText() string`

GetCustomerSignoffText returns the CustomerSignoffText field if non-nil, zero value otherwise.

### GetCustomerSignoffTextOk

`func (o *ServiceSignoff) GetCustomerSignoffTextOk() (*string, bool)`

GetCustomerSignoffTextOk returns a tuple with the CustomerSignoffText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSignoffText

`func (o *ServiceSignoff) SetCustomerSignoffText(v string)`

SetCustomerSignoffText sets CustomerSignoffText field to given value.

### HasCustomerSignoffText

`func (o *ServiceSignoff) HasCustomerSignoffText() bool`

HasCustomerSignoffText returns a boolean if a field has been set.

### GetCustomerSignoffFieldsFlag

`func (o *ServiceSignoff) GetCustomerSignoffFieldsFlag() bool`

GetCustomerSignoffFieldsFlag returns the CustomerSignoffFieldsFlag field if non-nil, zero value otherwise.

### GetCustomerSignoffFieldsFlagOk

`func (o *ServiceSignoff) GetCustomerSignoffFieldsFlagOk() (*bool, bool)`

GetCustomerSignoffFieldsFlagOk returns a tuple with the CustomerSignoffFieldsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerSignoffFieldsFlag

`func (o *ServiceSignoff) SetCustomerSignoffFieldsFlag(v bool)`

SetCustomerSignoffFieldsFlag sets CustomerSignoffFieldsFlag field to given value.

### HasCustomerSignoffFieldsFlag

`func (o *ServiceSignoff) HasCustomerSignoffFieldsFlag() bool`

HasCustomerSignoffFieldsFlag returns a boolean if a field has been set.

### SetCustomerSignoffFieldsFlagNil

`func (o *ServiceSignoff) SetCustomerSignoffFieldsFlagNil(b bool)`

 SetCustomerSignoffFieldsFlagNil sets the value for CustomerSignoffFieldsFlag to be an explicit nil

### UnsetCustomerSignoffFieldsFlag
`func (o *ServiceSignoff) UnsetCustomerSignoffFieldsFlag()`

UnsetCustomerSignoffFieldsFlag ensures that no value is present for CustomerSignoffFieldsFlag, not even an explicit nil
### GetBillingMethodsTextFlag

`func (o *ServiceSignoff) GetBillingMethodsTextFlag() bool`

GetBillingMethodsTextFlag returns the BillingMethodsTextFlag field if non-nil, zero value otherwise.

### GetBillingMethodsTextFlagOk

`func (o *ServiceSignoff) GetBillingMethodsTextFlagOk() (*bool, bool)`

GetBillingMethodsTextFlagOk returns a tuple with the BillingMethodsTextFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethodsTextFlag

`func (o *ServiceSignoff) SetBillingMethodsTextFlag(v bool)`

SetBillingMethodsTextFlag sets BillingMethodsTextFlag field to given value.

### HasBillingMethodsTextFlag

`func (o *ServiceSignoff) HasBillingMethodsTextFlag() bool`

HasBillingMethodsTextFlag returns a boolean if a field has been set.

### SetBillingMethodsTextFlagNil

`func (o *ServiceSignoff) SetBillingMethodsTextFlagNil(b bool)`

 SetBillingMethodsTextFlagNil sets the value for BillingMethodsTextFlag to be an explicit nil

### UnsetBillingMethodsTextFlag
`func (o *ServiceSignoff) UnsetBillingMethodsTextFlag()`

UnsetBillingMethodsTextFlag ensures that no value is present for BillingMethodsTextFlag, not even an explicit nil
### GetBillingMethodsText

`func (o *ServiceSignoff) GetBillingMethodsText() string`

GetBillingMethodsText returns the BillingMethodsText field if non-nil, zero value otherwise.

### GetBillingMethodsTextOk

`func (o *ServiceSignoff) GetBillingMethodsTextOk() (*string, bool)`

GetBillingMethodsTextOk returns a tuple with the BillingMethodsText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMethodsText

`func (o *ServiceSignoff) SetBillingMethodsText(v string)`

SetBillingMethodsText sets BillingMethodsText field to given value.

### HasBillingMethodsText

`func (o *ServiceSignoff) HasBillingMethodsText() bool`

HasBillingMethodsText returns a boolean if a field has been set.

### GetCreditCardFieldsFlag

`func (o *ServiceSignoff) GetCreditCardFieldsFlag() bool`

GetCreditCardFieldsFlag returns the CreditCardFieldsFlag field if non-nil, zero value otherwise.

### GetCreditCardFieldsFlagOk

`func (o *ServiceSignoff) GetCreditCardFieldsFlagOk() (*bool, bool)`

GetCreditCardFieldsFlagOk returns a tuple with the CreditCardFieldsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCardFieldsFlag

`func (o *ServiceSignoff) SetCreditCardFieldsFlag(v bool)`

SetCreditCardFieldsFlag sets CreditCardFieldsFlag field to given value.

### HasCreditCardFieldsFlag

`func (o *ServiceSignoff) HasCreditCardFieldsFlag() bool`

HasCreditCardFieldsFlag returns a boolean if a field has been set.

### SetCreditCardFieldsFlagNil

`func (o *ServiceSignoff) SetCreditCardFieldsFlagNil(b bool)`

 SetCreditCardFieldsFlagNil sets the value for CreditCardFieldsFlag to be an explicit nil

### UnsetCreditCardFieldsFlag
`func (o *ServiceSignoff) UnsetCreditCardFieldsFlag()`

UnsetCreditCardFieldsFlag ensures that no value is present for CreditCardFieldsFlag, not even an explicit nil
### GetDefaultFFFlag

`func (o *ServiceSignoff) GetDefaultFFFlag() bool`

GetDefaultFFFlag returns the DefaultFFFlag field if non-nil, zero value otherwise.

### GetDefaultFFFlagOk

`func (o *ServiceSignoff) GetDefaultFFFlagOk() (*bool, bool)`

GetDefaultFFFlagOk returns a tuple with the DefaultFFFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFFFlag

`func (o *ServiceSignoff) SetDefaultFFFlag(v bool)`

SetDefaultFFFlag sets DefaultFFFlag field to given value.

### HasDefaultFFFlag

`func (o *ServiceSignoff) HasDefaultFFFlag() bool`

HasDefaultFFFlag returns a boolean if a field has been set.

### SetDefaultFFFlagNil

`func (o *ServiceSignoff) SetDefaultFFFlagNil(b bool)`

 SetDefaultFFFlagNil sets the value for DefaultFFFlag to be an explicit nil

### UnsetDefaultFFFlag
`func (o *ServiceSignoff) UnsetDefaultFFFlag()`

UnsetDefaultFFFlag ensures that no value is present for DefaultFFFlag, not even an explicit nil
### GetInfo

`func (o *ServiceSignoff) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceSignoff) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceSignoff) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceSignoff) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


