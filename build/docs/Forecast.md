# Forecast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ForecastItems** | Pointer to [**[]ForecastItem**](ForecastItem.md) |  | [optional] 
**ProductRevenue** | Pointer to [**ProductRevenueReference**](ProductRevenueReference.md) |  | [optional] 
**ServiceRevenue** | Pointer to [**ServiceRevenueReference**](ServiceRevenueReference.md) |  | [optional] 
**AgreementRevenue** | Pointer to [**AgreementRevenueReference**](AgreementRevenueReference.md) |  | [optional] 
**TimeRevenue** | Pointer to [**TimeRevenueReference**](TimeRevenueReference.md) |  | [optional] 
**ExpenseRevenue** | Pointer to [**ExpenseRevenueReference**](ExpenseRevenueReference.md) |  | [optional] 
**ForecastRevenueTotals** | Pointer to [**ForecastRevenueReference**](ForecastRevenueReference.md) |  | [optional] 
**InclusiveRevenueTotals** | Pointer to [**InclusiveRevenueReference**](InclusiveRevenueReference.md) |  | [optional] 
**RecurringTotal** | Pointer to **NullableFloat64** |  | [optional] 
**WonRevenue** | Pointer to [**WonRevenueReference**](WonRevenueReference.md) |  | [optional] 
**LostRevenue** | Pointer to [**LostRevenueReference**](LostRevenueReference.md) |  | [optional] 
**OpenRevenue** | Pointer to [**OpenRevenueReference**](OpenRevenueReference.md) |  | [optional] 
**OtherRevenue1** | Pointer to [**Other1RevenueReference**](Other1RevenueReference.md) |  | [optional] 
**OtherRevenue2** | Pointer to [**Other2RevenueReference**](Other2RevenueReference.md) |  | [optional] 
**SalesTaxRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**ForecastTotalWithTaxes** | Pointer to **NullableFloat64** |  | [optional] 
**ExpectedProbability** | Pointer to **int32** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**BillingTerms** | Pointer to [**BillingTermsReference**](BillingTermsReference.md) |  | [optional] 
**Currency** | Pointer to [**CurrencyReference**](CurrencyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewForecast

`func NewForecast() *Forecast`

NewForecast instantiates a new Forecast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForecastWithDefaults

`func NewForecastWithDefaults() *Forecast`

NewForecastWithDefaults instantiates a new Forecast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Forecast) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Forecast) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Forecast) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Forecast) HasId() bool`

HasId returns a boolean if a field has been set.

### GetForecastItems

`func (o *Forecast) GetForecastItems() []ForecastItem`

GetForecastItems returns the ForecastItems field if non-nil, zero value otherwise.

### GetForecastItemsOk

`func (o *Forecast) GetForecastItemsOk() (*[]ForecastItem, bool)`

GetForecastItemsOk returns a tuple with the ForecastItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastItems

`func (o *Forecast) SetForecastItems(v []ForecastItem)`

SetForecastItems sets ForecastItems field to given value.

### HasForecastItems

`func (o *Forecast) HasForecastItems() bool`

HasForecastItems returns a boolean if a field has been set.

### GetProductRevenue

`func (o *Forecast) GetProductRevenue() ProductRevenueReference`

GetProductRevenue returns the ProductRevenue field if non-nil, zero value otherwise.

### GetProductRevenueOk

`func (o *Forecast) GetProductRevenueOk() (*ProductRevenueReference, bool)`

GetProductRevenueOk returns a tuple with the ProductRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRevenue

`func (o *Forecast) SetProductRevenue(v ProductRevenueReference)`

SetProductRevenue sets ProductRevenue field to given value.

### HasProductRevenue

`func (o *Forecast) HasProductRevenue() bool`

HasProductRevenue returns a boolean if a field has been set.

### GetServiceRevenue

`func (o *Forecast) GetServiceRevenue() ServiceRevenueReference`

GetServiceRevenue returns the ServiceRevenue field if non-nil, zero value otherwise.

### GetServiceRevenueOk

`func (o *Forecast) GetServiceRevenueOk() (*ServiceRevenueReference, bool)`

GetServiceRevenueOk returns a tuple with the ServiceRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRevenue

`func (o *Forecast) SetServiceRevenue(v ServiceRevenueReference)`

SetServiceRevenue sets ServiceRevenue field to given value.

### HasServiceRevenue

`func (o *Forecast) HasServiceRevenue() bool`

HasServiceRevenue returns a boolean if a field has been set.

### GetAgreementRevenue

`func (o *Forecast) GetAgreementRevenue() AgreementRevenueReference`

GetAgreementRevenue returns the AgreementRevenue field if non-nil, zero value otherwise.

### GetAgreementRevenueOk

`func (o *Forecast) GetAgreementRevenueOk() (*AgreementRevenueReference, bool)`

GetAgreementRevenueOk returns a tuple with the AgreementRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementRevenue

`func (o *Forecast) SetAgreementRevenue(v AgreementRevenueReference)`

SetAgreementRevenue sets AgreementRevenue field to given value.

### HasAgreementRevenue

`func (o *Forecast) HasAgreementRevenue() bool`

HasAgreementRevenue returns a boolean if a field has been set.

### GetTimeRevenue

`func (o *Forecast) GetTimeRevenue() TimeRevenueReference`

GetTimeRevenue returns the TimeRevenue field if non-nil, zero value otherwise.

### GetTimeRevenueOk

`func (o *Forecast) GetTimeRevenueOk() (*TimeRevenueReference, bool)`

GetTimeRevenueOk returns a tuple with the TimeRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRevenue

`func (o *Forecast) SetTimeRevenue(v TimeRevenueReference)`

SetTimeRevenue sets TimeRevenue field to given value.

### HasTimeRevenue

`func (o *Forecast) HasTimeRevenue() bool`

HasTimeRevenue returns a boolean if a field has been set.

### GetExpenseRevenue

`func (o *Forecast) GetExpenseRevenue() ExpenseRevenueReference`

GetExpenseRevenue returns the ExpenseRevenue field if non-nil, zero value otherwise.

### GetExpenseRevenueOk

`func (o *Forecast) GetExpenseRevenueOk() (*ExpenseRevenueReference, bool)`

GetExpenseRevenueOk returns a tuple with the ExpenseRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenseRevenue

`func (o *Forecast) SetExpenseRevenue(v ExpenseRevenueReference)`

SetExpenseRevenue sets ExpenseRevenue field to given value.

### HasExpenseRevenue

`func (o *Forecast) HasExpenseRevenue() bool`

HasExpenseRevenue returns a boolean if a field has been set.

### GetForecastRevenueTotals

`func (o *Forecast) GetForecastRevenueTotals() ForecastRevenueReference`

GetForecastRevenueTotals returns the ForecastRevenueTotals field if non-nil, zero value otherwise.

### GetForecastRevenueTotalsOk

`func (o *Forecast) GetForecastRevenueTotalsOk() (*ForecastRevenueReference, bool)`

GetForecastRevenueTotalsOk returns a tuple with the ForecastRevenueTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastRevenueTotals

`func (o *Forecast) SetForecastRevenueTotals(v ForecastRevenueReference)`

SetForecastRevenueTotals sets ForecastRevenueTotals field to given value.

### HasForecastRevenueTotals

`func (o *Forecast) HasForecastRevenueTotals() bool`

HasForecastRevenueTotals returns a boolean if a field has been set.

### GetInclusiveRevenueTotals

`func (o *Forecast) GetInclusiveRevenueTotals() InclusiveRevenueReference`

GetInclusiveRevenueTotals returns the InclusiveRevenueTotals field if non-nil, zero value otherwise.

### GetInclusiveRevenueTotalsOk

`func (o *Forecast) GetInclusiveRevenueTotalsOk() (*InclusiveRevenueReference, bool)`

GetInclusiveRevenueTotalsOk returns a tuple with the InclusiveRevenueTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInclusiveRevenueTotals

`func (o *Forecast) SetInclusiveRevenueTotals(v InclusiveRevenueReference)`

SetInclusiveRevenueTotals sets InclusiveRevenueTotals field to given value.

### HasInclusiveRevenueTotals

`func (o *Forecast) HasInclusiveRevenueTotals() bool`

HasInclusiveRevenueTotals returns a boolean if a field has been set.

### GetRecurringTotal

`func (o *Forecast) GetRecurringTotal() float64`

GetRecurringTotal returns the RecurringTotal field if non-nil, zero value otherwise.

### GetRecurringTotalOk

`func (o *Forecast) GetRecurringTotalOk() (*float64, bool)`

GetRecurringTotalOk returns a tuple with the RecurringTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringTotal

`func (o *Forecast) SetRecurringTotal(v float64)`

SetRecurringTotal sets RecurringTotal field to given value.

### HasRecurringTotal

`func (o *Forecast) HasRecurringTotal() bool`

HasRecurringTotal returns a boolean if a field has been set.

### SetRecurringTotalNil

`func (o *Forecast) SetRecurringTotalNil(b bool)`

 SetRecurringTotalNil sets the value for RecurringTotal to be an explicit nil

### UnsetRecurringTotal
`func (o *Forecast) UnsetRecurringTotal()`

UnsetRecurringTotal ensures that no value is present for RecurringTotal, not even an explicit nil
### GetWonRevenue

`func (o *Forecast) GetWonRevenue() WonRevenueReference`

GetWonRevenue returns the WonRevenue field if non-nil, zero value otherwise.

### GetWonRevenueOk

`func (o *Forecast) GetWonRevenueOk() (*WonRevenueReference, bool)`

GetWonRevenueOk returns a tuple with the WonRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWonRevenue

`func (o *Forecast) SetWonRevenue(v WonRevenueReference)`

SetWonRevenue sets WonRevenue field to given value.

### HasWonRevenue

`func (o *Forecast) HasWonRevenue() bool`

HasWonRevenue returns a boolean if a field has been set.

### GetLostRevenue

`func (o *Forecast) GetLostRevenue() LostRevenueReference`

GetLostRevenue returns the LostRevenue field if non-nil, zero value otherwise.

### GetLostRevenueOk

`func (o *Forecast) GetLostRevenueOk() (*LostRevenueReference, bool)`

GetLostRevenueOk returns a tuple with the LostRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostRevenue

`func (o *Forecast) SetLostRevenue(v LostRevenueReference)`

SetLostRevenue sets LostRevenue field to given value.

### HasLostRevenue

`func (o *Forecast) HasLostRevenue() bool`

HasLostRevenue returns a boolean if a field has been set.

### GetOpenRevenue

`func (o *Forecast) GetOpenRevenue() OpenRevenueReference`

GetOpenRevenue returns the OpenRevenue field if non-nil, zero value otherwise.

### GetOpenRevenueOk

`func (o *Forecast) GetOpenRevenueOk() (*OpenRevenueReference, bool)`

GetOpenRevenueOk returns a tuple with the OpenRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenRevenue

`func (o *Forecast) SetOpenRevenue(v OpenRevenueReference)`

SetOpenRevenue sets OpenRevenue field to given value.

### HasOpenRevenue

`func (o *Forecast) HasOpenRevenue() bool`

HasOpenRevenue returns a boolean if a field has been set.

### GetOtherRevenue1

`func (o *Forecast) GetOtherRevenue1() Other1RevenueReference`

GetOtherRevenue1 returns the OtherRevenue1 field if non-nil, zero value otherwise.

### GetOtherRevenue1Ok

`func (o *Forecast) GetOtherRevenue1Ok() (*Other1RevenueReference, bool)`

GetOtherRevenue1Ok returns a tuple with the OtherRevenue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherRevenue1

`func (o *Forecast) SetOtherRevenue1(v Other1RevenueReference)`

SetOtherRevenue1 sets OtherRevenue1 field to given value.

### HasOtherRevenue1

`func (o *Forecast) HasOtherRevenue1() bool`

HasOtherRevenue1 returns a boolean if a field has been set.

### GetOtherRevenue2

`func (o *Forecast) GetOtherRevenue2() Other2RevenueReference`

GetOtherRevenue2 returns the OtherRevenue2 field if non-nil, zero value otherwise.

### GetOtherRevenue2Ok

`func (o *Forecast) GetOtherRevenue2Ok() (*Other2RevenueReference, bool)`

GetOtherRevenue2Ok returns a tuple with the OtherRevenue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherRevenue2

`func (o *Forecast) SetOtherRevenue2(v Other2RevenueReference)`

SetOtherRevenue2 sets OtherRevenue2 field to given value.

### HasOtherRevenue2

`func (o *Forecast) HasOtherRevenue2() bool`

HasOtherRevenue2 returns a boolean if a field has been set.

### GetSalesTaxRevenue

`func (o *Forecast) GetSalesTaxRevenue() float64`

GetSalesTaxRevenue returns the SalesTaxRevenue field if non-nil, zero value otherwise.

### GetSalesTaxRevenueOk

`func (o *Forecast) GetSalesTaxRevenueOk() (*float64, bool)`

GetSalesTaxRevenueOk returns a tuple with the SalesTaxRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesTaxRevenue

`func (o *Forecast) SetSalesTaxRevenue(v float64)`

SetSalesTaxRevenue sets SalesTaxRevenue field to given value.

### HasSalesTaxRevenue

`func (o *Forecast) HasSalesTaxRevenue() bool`

HasSalesTaxRevenue returns a boolean if a field has been set.

### SetSalesTaxRevenueNil

`func (o *Forecast) SetSalesTaxRevenueNil(b bool)`

 SetSalesTaxRevenueNil sets the value for SalesTaxRevenue to be an explicit nil

### UnsetSalesTaxRevenue
`func (o *Forecast) UnsetSalesTaxRevenue()`

UnsetSalesTaxRevenue ensures that no value is present for SalesTaxRevenue, not even an explicit nil
### GetForecastTotalWithTaxes

`func (o *Forecast) GetForecastTotalWithTaxes() float64`

GetForecastTotalWithTaxes returns the ForecastTotalWithTaxes field if non-nil, zero value otherwise.

### GetForecastTotalWithTaxesOk

`func (o *Forecast) GetForecastTotalWithTaxesOk() (*float64, bool)`

GetForecastTotalWithTaxesOk returns a tuple with the ForecastTotalWithTaxes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForecastTotalWithTaxes

`func (o *Forecast) SetForecastTotalWithTaxes(v float64)`

SetForecastTotalWithTaxes sets ForecastTotalWithTaxes field to given value.

### HasForecastTotalWithTaxes

`func (o *Forecast) HasForecastTotalWithTaxes() bool`

HasForecastTotalWithTaxes returns a boolean if a field has been set.

### SetForecastTotalWithTaxesNil

`func (o *Forecast) SetForecastTotalWithTaxesNil(b bool)`

 SetForecastTotalWithTaxesNil sets the value for ForecastTotalWithTaxes to be an explicit nil

### UnsetForecastTotalWithTaxes
`func (o *Forecast) UnsetForecastTotalWithTaxes()`

UnsetForecastTotalWithTaxes ensures that no value is present for ForecastTotalWithTaxes, not even an explicit nil
### GetExpectedProbability

`func (o *Forecast) GetExpectedProbability() int32`

GetExpectedProbability returns the ExpectedProbability field if non-nil, zero value otherwise.

### GetExpectedProbabilityOk

`func (o *Forecast) GetExpectedProbabilityOk() (*int32, bool)`

GetExpectedProbabilityOk returns a tuple with the ExpectedProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedProbability

`func (o *Forecast) SetExpectedProbability(v int32)`

SetExpectedProbability sets ExpectedProbability field to given value.

### HasExpectedProbability

`func (o *Forecast) HasExpectedProbability() bool`

HasExpectedProbability returns a boolean if a field has been set.

### GetTaxCode

`func (o *Forecast) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *Forecast) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *Forecast) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *Forecast) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetBillingTerms

`func (o *Forecast) GetBillingTerms() BillingTermsReference`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *Forecast) GetBillingTermsOk() (*BillingTermsReference, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *Forecast) SetBillingTerms(v BillingTermsReference)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *Forecast) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### GetCurrency

`func (o *Forecast) GetCurrency() CurrencyReference`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Forecast) GetCurrencyOk() (*CurrencyReference, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Forecast) SetCurrency(v CurrencyReference)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Forecast) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetInfo

`func (o *Forecast) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Forecast) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Forecast) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Forecast) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


