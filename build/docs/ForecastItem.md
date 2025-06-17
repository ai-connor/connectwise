# ForecastItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForecastDescription** | Pointer to **string** |  Max length: 50; | [optional] 
**Opportunity** | [**OpportunityReference**](OpportunityReference.md) |  | 
**Quantity** | Pointer to **float64** |  | [optional] 
**Status** | [**OpportunityStatusReference**](OpportunityStatusReference.md) |  | 
**CatalogItem** | Pointer to [**IvItemReference**](IvItemReference.md) |  | [optional] 
**ProductDescription** | Pointer to **string** |  | [optional] 
**ProductClass** | Pointer to **string** |  | [optional] 
**Revenue** | Pointer to **float64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Margin** | Pointer to **float64** |  | [optional] 
**Percentage** | Pointer to **int32** |  | [optional] 
**IncludeFlag** | Pointer to **bool** |  | [optional] 
**QuoteWerksDocNo** | Pointer to **string** |  Max length: 20; | [optional] 
**QuoteWerksDocName** | Pointer to **string** |  Max length: 255; | [optional] 
**QuoteWerksQuantity** | Pointer to **int32** |  | [optional] 
**ForecastType** | **NullableString** |  | 
**LinkFlag** | Pointer to **bool** |  | [optional] 
**RecurringRevenue** | Pointer to **float64** |  | [optional] 
**RecurringCost** | Pointer to **NullableFloat64** |  | [optional] 
**RecurringDateStart** | Pointer to **time.Time** |  | [optional] 
**RecurringDateEnd** | Pointer to **time.Time** |  | [optional] 
**BillCycle** | Pointer to [**BillingCycleReference**](BillingCycleReference.md) |  | [optional] 
**CycleBasis** | Pointer to **string** |  | [optional] 
**Cycles** | Pointer to **int32** |  | [optional] 
**RecurringFlag** | Pointer to **bool** |  | [optional] 
**SequenceNumber** | Pointer to **float64** |  | [optional] 
**SubNumber** | Pointer to **int32** |  | [optional] 
**TaxableFlag** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewForecastItem

`func NewForecastItem(opportunity OpportunityReference, status OpportunityStatusReference, forecastType NullableString, ) *ForecastItem`

NewForecastItem instantiates a new ForecastItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastItemWithDefaults

`func NewForecastItemWithDefaults() *ForecastItem`

NewForecastItemWithDefaults instantiates a new ForecastItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ForecastItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForecastItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForecastItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ForecastItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForecastDescription

`func (o *ForecastItem) GetForecastDescription() string`

GetForecastDescription returns the ForecastDescription field if non-nil, zero value otherwise.

### GetForecastDescriptionOk

`func (o *ForecastItem) GetForecastDescriptionOk() (*string, bool)`

GetForecastDescriptionOk returns a tuple with the ForecastDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastDescription

`func (o *ForecastItem) SetForecastDescription(v string)`

SetForecastDescription sets ForecastDescription field to given value.

### HasForecastDescription

`func (o *ForecastItem) HasForecastDescription() bool`

HasForecastDescription returns a boolean if a field has been set.

### GetOpportunity

`func (o *ForecastItem) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *ForecastItem) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *ForecastItem) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.


### GetQuantity

`func (o *ForecastItem) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ForecastItem) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ForecastItem) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ForecastItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetStatus

`func (o *ForecastItem) GetStatus() OpportunityStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ForecastItem) GetStatusOk() (*OpportunityStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ForecastItem) SetStatus(v OpportunityStatusReference)`

SetStatus sets Status field to given value.


### GetCatalogItem

`func (o *ForecastItem) GetCatalogItem() IvItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *ForecastItem) GetCatalogItemOk() (*IvItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *ForecastItem) SetCatalogItem(v IvItemReference)`

SetCatalogItem sets CatalogItem field to given value.

### HasCatalogItem

`func (o *ForecastItem) HasCatalogItem() bool`

HasCatalogItem returns a boolean if a field has been set.

### GetProductDescription

`func (o *ForecastItem) GetProductDescription() string`

GetProductDescription returns the ProductDescription field if non-nil, zero value otherwise.

### GetProductDescriptionOk

`func (o *ForecastItem) GetProductDescriptionOk() (*string, bool)`

GetProductDescriptionOk returns a tuple with the ProductDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductDescription

`func (o *ForecastItem) SetProductDescription(v string)`

SetProductDescription sets ProductDescription field to given value.

### HasProductDescription

`func (o *ForecastItem) HasProductDescription() bool`

HasProductDescription returns a boolean if a field has been set.

### GetProductClass

`func (o *ForecastItem) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *ForecastItem) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *ForecastItem) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *ForecastItem) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### GetRevenue

`func (o *ForecastItem) GetRevenue() float64`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *ForecastItem) GetRevenueOk() (*float64, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *ForecastItem) SetRevenue(v float64)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *ForecastItem) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### GetCost

`func (o *ForecastItem) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ForecastItem) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ForecastItem) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ForecastItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ForecastItem) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ForecastItem) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetMargin

`func (o *ForecastItem) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *ForecastItem) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *ForecastItem) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *ForecastItem) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetPercentage

`func (o *ForecastItem) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *ForecastItem) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *ForecastItem) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *ForecastItem) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetIncludeFlag

`func (o *ForecastItem) GetIncludeFlag() bool`

GetIncludeFlag returns the IncludeFlag field if non-nil, zero value otherwise.

### GetIncludeFlagOk

`func (o *ForecastItem) GetIncludeFlagOk() (*bool, bool)`

GetIncludeFlagOk returns a tuple with the IncludeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFlag

`func (o *ForecastItem) SetIncludeFlag(v bool)`

SetIncludeFlag sets IncludeFlag field to given value.

### HasIncludeFlag

`func (o *ForecastItem) HasIncludeFlag() bool`

HasIncludeFlag returns a boolean if a field has been set.

### GetQuoteWerksDocNo

`func (o *ForecastItem) GetQuoteWerksDocNo() string`

GetQuoteWerksDocNo returns the QuoteWerksDocNo field if non-nil, zero value otherwise.

### GetQuoteWerksDocNoOk

`func (o *ForecastItem) GetQuoteWerksDocNoOk() (*string, bool)`

GetQuoteWerksDocNoOk returns a tuple with the QuoteWerksDocNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteWerksDocNo

`func (o *ForecastItem) SetQuoteWerksDocNo(v string)`

SetQuoteWerksDocNo sets QuoteWerksDocNo field to given value.

### HasQuoteWerksDocNo

`func (o *ForecastItem) HasQuoteWerksDocNo() bool`

HasQuoteWerksDocNo returns a boolean if a field has been set.

### GetQuoteWerksDocName

`func (o *ForecastItem) GetQuoteWerksDocName() string`

GetQuoteWerksDocName returns the QuoteWerksDocName field if non-nil, zero value otherwise.

### GetQuoteWerksDocNameOk

`func (o *ForecastItem) GetQuoteWerksDocNameOk() (*string, bool)`

GetQuoteWerksDocNameOk returns a tuple with the QuoteWerksDocName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteWerksDocName

`func (o *ForecastItem) SetQuoteWerksDocName(v string)`

SetQuoteWerksDocName sets QuoteWerksDocName field to given value.

### HasQuoteWerksDocName

`func (o *ForecastItem) HasQuoteWerksDocName() bool`

HasQuoteWerksDocName returns a boolean if a field has been set.

### GetQuoteWerksQuantity

`func (o *ForecastItem) GetQuoteWerksQuantity() int32`

GetQuoteWerksQuantity returns the QuoteWerksQuantity field if non-nil, zero value otherwise.

### GetQuoteWerksQuantityOk

`func (o *ForecastItem) GetQuoteWerksQuantityOk() (*int32, bool)`

GetQuoteWerksQuantityOk returns a tuple with the QuoteWerksQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteWerksQuantity

`func (o *ForecastItem) SetQuoteWerksQuantity(v int32)`

SetQuoteWerksQuantity sets QuoteWerksQuantity field to given value.

### HasQuoteWerksQuantity

`func (o *ForecastItem) HasQuoteWerksQuantity() bool`

HasQuoteWerksQuantity returns a boolean if a field has been set.

### GetForecastType

`func (o *ForecastItem) GetForecastType() string`

GetForecastType returns the ForecastType field if non-nil, zero value otherwise.

### GetForecastTypeOk

`func (o *ForecastItem) GetForecastTypeOk() (*string, bool)`

GetForecastTypeOk returns a tuple with the ForecastType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastType

`func (o *ForecastItem) SetForecastType(v string)`

SetForecastType sets ForecastType field to given value.


### SetForecastTypeNil

`func (o *ForecastItem) SetForecastTypeNil(b bool)`

 SetForecastTypeNil sets the value for ForecastType to be an explicit nil

### UnsetForecastType
`func (o *ForecastItem) UnsetForecastType()`

UnsetForecastType ensures that no value is present for ForecastType, not even an explicit nil
### GetLinkFlag

`func (o *ForecastItem) GetLinkFlag() bool`

GetLinkFlag returns the LinkFlag field if non-nil, zero value otherwise.

### GetLinkFlagOk

`func (o *ForecastItem) GetLinkFlagOk() (*bool, bool)`

GetLinkFlagOk returns a tuple with the LinkFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkFlag

`func (o *ForecastItem) SetLinkFlag(v bool)`

SetLinkFlag sets LinkFlag field to given value.

### HasLinkFlag

`func (o *ForecastItem) HasLinkFlag() bool`

HasLinkFlag returns a boolean if a field has been set.

### GetRecurringRevenue

`func (o *ForecastItem) GetRecurringRevenue() float64`

GetRecurringRevenue returns the RecurringRevenue field if non-nil, zero value otherwise.

### GetRecurringRevenueOk

`func (o *ForecastItem) GetRecurringRevenueOk() (*float64, bool)`

GetRecurringRevenueOk returns a tuple with the RecurringRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRevenue

`func (o *ForecastItem) SetRecurringRevenue(v float64)`

SetRecurringRevenue sets RecurringRevenue field to given value.

### HasRecurringRevenue

`func (o *ForecastItem) HasRecurringRevenue() bool`

HasRecurringRevenue returns a boolean if a field has been set.

### GetRecurringCost

`func (o *ForecastItem) GetRecurringCost() float64`

GetRecurringCost returns the RecurringCost field if non-nil, zero value otherwise.

### GetRecurringCostOk

`func (o *ForecastItem) GetRecurringCostOk() (*float64, bool)`

GetRecurringCostOk returns a tuple with the RecurringCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCost

`func (o *ForecastItem) SetRecurringCost(v float64)`

SetRecurringCost sets RecurringCost field to given value.

### HasRecurringCost

`func (o *ForecastItem) HasRecurringCost() bool`

HasRecurringCost returns a boolean if a field has been set.

### SetRecurringCostNil

`func (o *ForecastItem) SetRecurringCostNil(b bool)`

 SetRecurringCostNil sets the value for RecurringCost to be an explicit nil

### UnsetRecurringCost
`func (o *ForecastItem) UnsetRecurringCost()`

UnsetRecurringCost ensures that no value is present for RecurringCost, not even an explicit nil
### GetRecurringDateStart

`func (o *ForecastItem) GetRecurringDateStart() time.Time`

GetRecurringDateStart returns the RecurringDateStart field if non-nil, zero value otherwise.

### GetRecurringDateStartOk

`func (o *ForecastItem) GetRecurringDateStartOk() (*time.Time, bool)`

GetRecurringDateStartOk returns a tuple with the RecurringDateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDateStart

`func (o *ForecastItem) SetRecurringDateStart(v time.Time)`

SetRecurringDateStart sets RecurringDateStart field to given value.

### HasRecurringDateStart

`func (o *ForecastItem) HasRecurringDateStart() bool`

HasRecurringDateStart returns a boolean if a field has been set.

### GetRecurringDateEnd

`func (o *ForecastItem) GetRecurringDateEnd() time.Time`

GetRecurringDateEnd returns the RecurringDateEnd field if non-nil, zero value otherwise.

### GetRecurringDateEndOk

`func (o *ForecastItem) GetRecurringDateEndOk() (*time.Time, bool)`

GetRecurringDateEndOk returns a tuple with the RecurringDateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringDateEnd

`func (o *ForecastItem) SetRecurringDateEnd(v time.Time)`

SetRecurringDateEnd sets RecurringDateEnd field to given value.

### HasRecurringDateEnd

`func (o *ForecastItem) HasRecurringDateEnd() bool`

HasRecurringDateEnd returns a boolean if a field has been set.

### GetBillCycle

`func (o *ForecastItem) GetBillCycle() BillingCycleReference`

GetBillCycle returns the BillCycle field if non-nil, zero value otherwise.

### GetBillCycleOk

`func (o *ForecastItem) GetBillCycleOk() (*BillingCycleReference, bool)`

GetBillCycleOk returns a tuple with the BillCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCycle

`func (o *ForecastItem) SetBillCycle(v BillingCycleReference)`

SetBillCycle sets BillCycle field to given value.

### HasBillCycle

`func (o *ForecastItem) HasBillCycle() bool`

HasBillCycle returns a boolean if a field has been set.

### GetCycleBasis

`func (o *ForecastItem) GetCycleBasis() string`

GetCycleBasis returns the CycleBasis field if non-nil, zero value otherwise.

### GetCycleBasisOk

`func (o *ForecastItem) GetCycleBasisOk() (*string, bool)`

GetCycleBasisOk returns a tuple with the CycleBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleBasis

`func (o *ForecastItem) SetCycleBasis(v string)`

SetCycleBasis sets CycleBasis field to given value.

### HasCycleBasis

`func (o *ForecastItem) HasCycleBasis() bool`

HasCycleBasis returns a boolean if a field has been set.

### GetCycles

`func (o *ForecastItem) GetCycles() int32`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *ForecastItem) GetCyclesOk() (*int32, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *ForecastItem) SetCycles(v int32)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *ForecastItem) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### GetRecurringFlag

`func (o *ForecastItem) GetRecurringFlag() bool`

GetRecurringFlag returns the RecurringFlag field if non-nil, zero value otherwise.

### GetRecurringFlagOk

`func (o *ForecastItem) GetRecurringFlagOk() (*bool, bool)`

GetRecurringFlagOk returns a tuple with the RecurringFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringFlag

`func (o *ForecastItem) SetRecurringFlag(v bool)`

SetRecurringFlag sets RecurringFlag field to given value.

### HasRecurringFlag

`func (o *ForecastItem) HasRecurringFlag() bool`

HasRecurringFlag returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ForecastItem) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ForecastItem) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ForecastItem) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ForecastItem) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetSubNumber

`func (o *ForecastItem) GetSubNumber() int32`

GetSubNumber returns the SubNumber field if non-nil, zero value otherwise.

### GetSubNumberOk

`func (o *ForecastItem) GetSubNumberOk() (*int32, bool)`

GetSubNumberOk returns a tuple with the SubNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubNumber

`func (o *ForecastItem) SetSubNumber(v int32)`

SetSubNumber sets SubNumber field to given value.

### HasSubNumber

`func (o *ForecastItem) HasSubNumber() bool`

HasSubNumber returns a boolean if a field has been set.

### GetTaxableFlag

`func (o *ForecastItem) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *ForecastItem) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *ForecastItem) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *ForecastItem) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### GetInfo

`func (o *ForecastItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ForecastItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ForecastItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ForecastItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


