# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Type** | [**CampaignTypeReference**](CampaignTypeReference.md) |  | 
**SubType** | [**CampaignSubTypeReference**](CampaignSubTypeReference.md) |  | 
**Status** | Pointer to [**CampaignStatusReference**](CampaignStatusReference.md) |  | [optional] 
**StartDate** | **time.Time** |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Inactive** | Pointer to **NullableBool** |  | [optional] 
**InactiveDaysAfterEnd** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**DefaultGroup** | Pointer to [**GroupReference**](GroupReference.md) |  | [optional] 
**MarketingManagerDefaultTrackId** | Pointer to **NullableInt32** |  | [optional] 
**OpportunityDefaultTrackId** | Pointer to **NullableInt32** |  | [optional] 
**Impressions** | Pointer to **NullableInt32** |  | [optional] 
**BudgetRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**BudgetCost** | Pointer to **NullableFloat64** |  | [optional] 
**ActualCost** | Pointer to **NullableFloat64** |  | [optional] 
**BudgetGrossMargin** | Pointer to **NullableFloat64** |  | [optional] 
**BudgetROI** | Pointer to **NullableFloat64** |  | [optional] 
**ActualRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**ActualGrossMargin** | Pointer to **NullableFloat64** |  | [optional] 
**ActualROI** | Pointer to **NullableFloat64** |  | [optional] 
**EmailsSent** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCampaign

`func NewCampaign(name string, type_ CampaignTypeReference, subType CampaignSubTypeReference, startDate time.Time, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Campaign) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Campaign) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Campaign) GetType() CampaignTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Campaign) GetTypeOk() (*CampaignTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Campaign) SetType(v CampaignTypeReference)`

SetType sets Type field to given value.


### GetSubType

`func (o *Campaign) GetSubType() CampaignSubTypeReference`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Campaign) GetSubTypeOk() (*CampaignSubTypeReference, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Campaign) SetSubType(v CampaignSubTypeReference)`

SetSubType sets SubType field to given value.


### GetStatus

`func (o *Campaign) GetStatus() CampaignStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Campaign) GetStatusOk() (*CampaignStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Campaign) SetStatus(v CampaignStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Campaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartDate

`func (o *Campaign) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *Campaign) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *Campaign) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *Campaign) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *Campaign) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *Campaign) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *Campaign) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetLocationId

`func (o *Campaign) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *Campaign) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *Campaign) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *Campaign) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *Campaign) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *Campaign) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetMember

`func (o *Campaign) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *Campaign) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *Campaign) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *Campaign) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInactive

`func (o *Campaign) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Campaign) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *Campaign) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *Campaign) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### SetInactiveNil

`func (o *Campaign) SetInactiveNil(b bool)`

 SetInactiveNil sets the value for Inactive to be an explicit nil

### UnsetInactive
`func (o *Campaign) UnsetInactive()`

UnsetInactive ensures that no value is present for Inactive, not even an explicit nil
### GetInactiveDaysAfterEnd

`func (o *Campaign) GetInactiveDaysAfterEnd() int32`

GetInactiveDaysAfterEnd returns the InactiveDaysAfterEnd field if non-nil, zero value otherwise.

### GetInactiveDaysAfterEndOk

`func (o *Campaign) GetInactiveDaysAfterEndOk() (*int32, bool)`

GetInactiveDaysAfterEndOk returns a tuple with the InactiveDaysAfterEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDaysAfterEnd

`func (o *Campaign) SetInactiveDaysAfterEnd(v int32)`

SetInactiveDaysAfterEnd sets InactiveDaysAfterEnd field to given value.

### HasInactiveDaysAfterEnd

`func (o *Campaign) HasInactiveDaysAfterEnd() bool`

HasInactiveDaysAfterEnd returns a boolean if a field has been set.

### SetInactiveDaysAfterEndNil

`func (o *Campaign) SetInactiveDaysAfterEndNil(b bool)`

 SetInactiveDaysAfterEndNil sets the value for InactiveDaysAfterEnd to be an explicit nil

### UnsetInactiveDaysAfterEnd
`func (o *Campaign) UnsetInactiveDaysAfterEnd()`

UnsetInactiveDaysAfterEnd ensures that no value is present for InactiveDaysAfterEnd, not even an explicit nil
### GetNotes

`func (o *Campaign) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Campaign) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Campaign) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Campaign) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *Campaign) GetDefaultGroup() GroupReference`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *Campaign) GetDefaultGroupOk() (*GroupReference, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *Campaign) SetDefaultGroup(v GroupReference)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *Campaign) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetMarketingManagerDefaultTrackId

`func (o *Campaign) GetMarketingManagerDefaultTrackId() int32`

GetMarketingManagerDefaultTrackId returns the MarketingManagerDefaultTrackId field if non-nil, zero value otherwise.

### GetMarketingManagerDefaultTrackIdOk

`func (o *Campaign) GetMarketingManagerDefaultTrackIdOk() (*int32, bool)`

GetMarketingManagerDefaultTrackIdOk returns a tuple with the MarketingManagerDefaultTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingManagerDefaultTrackId

`func (o *Campaign) SetMarketingManagerDefaultTrackId(v int32)`

SetMarketingManagerDefaultTrackId sets MarketingManagerDefaultTrackId field to given value.

### HasMarketingManagerDefaultTrackId

`func (o *Campaign) HasMarketingManagerDefaultTrackId() bool`

HasMarketingManagerDefaultTrackId returns a boolean if a field has been set.

### SetMarketingManagerDefaultTrackIdNil

`func (o *Campaign) SetMarketingManagerDefaultTrackIdNil(b bool)`

 SetMarketingManagerDefaultTrackIdNil sets the value for MarketingManagerDefaultTrackId to be an explicit nil

### UnsetMarketingManagerDefaultTrackId
`func (o *Campaign) UnsetMarketingManagerDefaultTrackId()`

UnsetMarketingManagerDefaultTrackId ensures that no value is present for MarketingManagerDefaultTrackId, not even an explicit nil
### GetOpportunityDefaultTrackId

`func (o *Campaign) GetOpportunityDefaultTrackId() int32`

GetOpportunityDefaultTrackId returns the OpportunityDefaultTrackId field if non-nil, zero value otherwise.

### GetOpportunityDefaultTrackIdOk

`func (o *Campaign) GetOpportunityDefaultTrackIdOk() (*int32, bool)`

GetOpportunityDefaultTrackIdOk returns a tuple with the OpportunityDefaultTrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunityDefaultTrackId

`func (o *Campaign) SetOpportunityDefaultTrackId(v int32)`

SetOpportunityDefaultTrackId sets OpportunityDefaultTrackId field to given value.

### HasOpportunityDefaultTrackId

`func (o *Campaign) HasOpportunityDefaultTrackId() bool`

HasOpportunityDefaultTrackId returns a boolean if a field has been set.

### SetOpportunityDefaultTrackIdNil

`func (o *Campaign) SetOpportunityDefaultTrackIdNil(b bool)`

 SetOpportunityDefaultTrackIdNil sets the value for OpportunityDefaultTrackId to be an explicit nil

### UnsetOpportunityDefaultTrackId
`func (o *Campaign) UnsetOpportunityDefaultTrackId()`

UnsetOpportunityDefaultTrackId ensures that no value is present for OpportunityDefaultTrackId, not even an explicit nil
### GetImpressions

`func (o *Campaign) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *Campaign) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *Campaign) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *Campaign) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### SetImpressionsNil

`func (o *Campaign) SetImpressionsNil(b bool)`

 SetImpressionsNil sets the value for Impressions to be an explicit nil

### UnsetImpressions
`func (o *Campaign) UnsetImpressions()`

UnsetImpressions ensures that no value is present for Impressions, not even an explicit nil
### GetBudgetRevenue

`func (o *Campaign) GetBudgetRevenue() float64`

GetBudgetRevenue returns the BudgetRevenue field if non-nil, zero value otherwise.

### GetBudgetRevenueOk

`func (o *Campaign) GetBudgetRevenueOk() (*float64, bool)`

GetBudgetRevenueOk returns a tuple with the BudgetRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetRevenue

`func (o *Campaign) SetBudgetRevenue(v float64)`

SetBudgetRevenue sets BudgetRevenue field to given value.

### HasBudgetRevenue

`func (o *Campaign) HasBudgetRevenue() bool`

HasBudgetRevenue returns a boolean if a field has been set.

### SetBudgetRevenueNil

`func (o *Campaign) SetBudgetRevenueNil(b bool)`

 SetBudgetRevenueNil sets the value for BudgetRevenue to be an explicit nil

### UnsetBudgetRevenue
`func (o *Campaign) UnsetBudgetRevenue()`

UnsetBudgetRevenue ensures that no value is present for BudgetRevenue, not even an explicit nil
### GetBudgetCost

`func (o *Campaign) GetBudgetCost() float64`

GetBudgetCost returns the BudgetCost field if non-nil, zero value otherwise.

### GetBudgetCostOk

`func (o *Campaign) GetBudgetCostOk() (*float64, bool)`

GetBudgetCostOk returns a tuple with the BudgetCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetCost

`func (o *Campaign) SetBudgetCost(v float64)`

SetBudgetCost sets BudgetCost field to given value.

### HasBudgetCost

`func (o *Campaign) HasBudgetCost() bool`

HasBudgetCost returns a boolean if a field has been set.

### SetBudgetCostNil

`func (o *Campaign) SetBudgetCostNil(b bool)`

 SetBudgetCostNil sets the value for BudgetCost to be an explicit nil

### UnsetBudgetCost
`func (o *Campaign) UnsetBudgetCost()`

UnsetBudgetCost ensures that no value is present for BudgetCost, not even an explicit nil
### GetActualCost

`func (o *Campaign) GetActualCost() float64`

GetActualCost returns the ActualCost field if non-nil, zero value otherwise.

### GetActualCostOk

`func (o *Campaign) GetActualCostOk() (*float64, bool)`

GetActualCostOk returns a tuple with the ActualCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualCost

`func (o *Campaign) SetActualCost(v float64)`

SetActualCost sets ActualCost field to given value.

### HasActualCost

`func (o *Campaign) HasActualCost() bool`

HasActualCost returns a boolean if a field has been set.

### SetActualCostNil

`func (o *Campaign) SetActualCostNil(b bool)`

 SetActualCostNil sets the value for ActualCost to be an explicit nil

### UnsetActualCost
`func (o *Campaign) UnsetActualCost()`

UnsetActualCost ensures that no value is present for ActualCost, not even an explicit nil
### GetBudgetGrossMargin

`func (o *Campaign) GetBudgetGrossMargin() float64`

GetBudgetGrossMargin returns the BudgetGrossMargin field if non-nil, zero value otherwise.

### GetBudgetGrossMarginOk

`func (o *Campaign) GetBudgetGrossMarginOk() (*float64, bool)`

GetBudgetGrossMarginOk returns a tuple with the BudgetGrossMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetGrossMargin

`func (o *Campaign) SetBudgetGrossMargin(v float64)`

SetBudgetGrossMargin sets BudgetGrossMargin field to given value.

### HasBudgetGrossMargin

`func (o *Campaign) HasBudgetGrossMargin() bool`

HasBudgetGrossMargin returns a boolean if a field has been set.

### SetBudgetGrossMarginNil

`func (o *Campaign) SetBudgetGrossMarginNil(b bool)`

 SetBudgetGrossMarginNil sets the value for BudgetGrossMargin to be an explicit nil

### UnsetBudgetGrossMargin
`func (o *Campaign) UnsetBudgetGrossMargin()`

UnsetBudgetGrossMargin ensures that no value is present for BudgetGrossMargin, not even an explicit nil
### GetBudgetROI

`func (o *Campaign) GetBudgetROI() float64`

GetBudgetROI returns the BudgetROI field if non-nil, zero value otherwise.

### GetBudgetROIOk

`func (o *Campaign) GetBudgetROIOk() (*float64, bool)`

GetBudgetROIOk returns a tuple with the BudgetROI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetROI

`func (o *Campaign) SetBudgetROI(v float64)`

SetBudgetROI sets BudgetROI field to given value.

### HasBudgetROI

`func (o *Campaign) HasBudgetROI() bool`

HasBudgetROI returns a boolean if a field has been set.

### SetBudgetROINil

`func (o *Campaign) SetBudgetROINil(b bool)`

 SetBudgetROINil sets the value for BudgetROI to be an explicit nil

### UnsetBudgetROI
`func (o *Campaign) UnsetBudgetROI()`

UnsetBudgetROI ensures that no value is present for BudgetROI, not even an explicit nil
### GetActualRevenue

`func (o *Campaign) GetActualRevenue() float64`

GetActualRevenue returns the ActualRevenue field if non-nil, zero value otherwise.

### GetActualRevenueOk

`func (o *Campaign) GetActualRevenueOk() (*float64, bool)`

GetActualRevenueOk returns a tuple with the ActualRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualRevenue

`func (o *Campaign) SetActualRevenue(v float64)`

SetActualRevenue sets ActualRevenue field to given value.

### HasActualRevenue

`func (o *Campaign) HasActualRevenue() bool`

HasActualRevenue returns a boolean if a field has been set.

### SetActualRevenueNil

`func (o *Campaign) SetActualRevenueNil(b bool)`

 SetActualRevenueNil sets the value for ActualRevenue to be an explicit nil

### UnsetActualRevenue
`func (o *Campaign) UnsetActualRevenue()`

UnsetActualRevenue ensures that no value is present for ActualRevenue, not even an explicit nil
### GetActualGrossMargin

`func (o *Campaign) GetActualGrossMargin() float64`

GetActualGrossMargin returns the ActualGrossMargin field if non-nil, zero value otherwise.

### GetActualGrossMarginOk

`func (o *Campaign) GetActualGrossMarginOk() (*float64, bool)`

GetActualGrossMarginOk returns a tuple with the ActualGrossMargin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualGrossMargin

`func (o *Campaign) SetActualGrossMargin(v float64)`

SetActualGrossMargin sets ActualGrossMargin field to given value.

### HasActualGrossMargin

`func (o *Campaign) HasActualGrossMargin() bool`

HasActualGrossMargin returns a boolean if a field has been set.

### SetActualGrossMarginNil

`func (o *Campaign) SetActualGrossMarginNil(b bool)`

 SetActualGrossMarginNil sets the value for ActualGrossMargin to be an explicit nil

### UnsetActualGrossMargin
`func (o *Campaign) UnsetActualGrossMargin()`

UnsetActualGrossMargin ensures that no value is present for ActualGrossMargin, not even an explicit nil
### GetActualROI

`func (o *Campaign) GetActualROI() float64`

GetActualROI returns the ActualROI field if non-nil, zero value otherwise.

### GetActualROIOk

`func (o *Campaign) GetActualROIOk() (*float64, bool)`

GetActualROIOk returns a tuple with the ActualROI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualROI

`func (o *Campaign) SetActualROI(v float64)`

SetActualROI sets ActualROI field to given value.

### HasActualROI

`func (o *Campaign) HasActualROI() bool`

HasActualROI returns a boolean if a field has been set.

### SetActualROINil

`func (o *Campaign) SetActualROINil(b bool)`

 SetActualROINil sets the value for ActualROI to be an explicit nil

### UnsetActualROI
`func (o *Campaign) UnsetActualROI()`

UnsetActualROI ensures that no value is present for ActualROI, not even an explicit nil
### GetEmailsSent

`func (o *Campaign) GetEmailsSent() int32`

GetEmailsSent returns the EmailsSent field if non-nil, zero value otherwise.

### GetEmailsSentOk

`func (o *Campaign) GetEmailsSentOk() (*int32, bool)`

GetEmailsSentOk returns a tuple with the EmailsSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsSent

`func (o *Campaign) SetEmailsSent(v int32)`

SetEmailsSent sets EmailsSent field to given value.

### HasEmailsSent

`func (o *Campaign) HasEmailsSent() bool`

HasEmailsSent returns a boolean if a field has been set.

### SetEmailsSentNil

`func (o *Campaign) SetEmailsSentNil(b bool)`

 SetEmailsSentNil sets the value for EmailsSent to be an explicit nil

### UnsetEmailsSent
`func (o *Campaign) UnsetEmailsSent()`

UnsetEmailsSent ensures that no value is present for EmailsSent, not even an explicit nil
### GetInfo

`func (o *Campaign) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Campaign) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Campaign) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Campaign) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


