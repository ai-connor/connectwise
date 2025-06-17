# AgreementRecurringParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingCycle** | Pointer to [**GenericNameIdDTO**](GenericNameIdDTO.md) |  | [optional] 
**CycleBase** | Pointer to [**GenericNameIdDTO**](GenericNameIdDTO.md) |  | [optional] 
**AGRAmount** | Pointer to **NullableFloat64** |  | [optional] 
**Taxable** | Pointer to **bool** |  | [optional] 
**ChildrenAmount** | Pointer to **float64** |  | [optional] 
**AdditionsAmount** | Pointer to **float64** |  | [optional] 
**TotalAmount** | Pointer to **float64** |  | [optional] 
**AGRProrate** | Pointer to **NullableFloat64** |  | [optional] 
**BillStartDate** | Pointer to **string** |  | [optional] 
**TaxCode** | Pointer to [**GenericNameIdDTO**](GenericNameIdDTO.md) |  | [optional] 
**Terms** | Pointer to [**GenericNameIdDTO**](GenericNameIdDTO.md) |  | [optional] 
**ProrateFlag** | Pointer to **bool** |  | [optional] 
**InvoiceProratedAdditionsFlag** | Pointer to **bool** |  | [optional] 
**RestrictDownpayment** | Pointer to **bool** |  | [optional] 
**Currency** | Pointer to [**GenericNameIdDTO**](GenericNameIdDTO.md) |  | [optional] 
**AutoInvoiceFlag** | Pointer to **bool** |  | [optional] 
**UserDefinedFieldValues** | Pointer to [**[]UserDefinedFieldValueModel**](UserDefinedFieldValueModel.md) |  | [optional] 

## Methods

### NewAgreementRecurringParameters

`func NewAgreementRecurringParameters() *AgreementRecurringParameters`

NewAgreementRecurringParameters instantiates a new AgreementRecurringParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementRecurringParametersWithDefaults

`func NewAgreementRecurringParametersWithDefaults() *AgreementRecurringParameters`

NewAgreementRecurringParametersWithDefaults instantiates a new AgreementRecurringParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingCycle

`func (o *AgreementRecurringParameters) GetBillingCycle() GenericNameIdDTO`

GetBillingCycle returns the BillingCycle field if non-nil, zero value otherwise.

### GetBillingCycleOk

`func (o *AgreementRecurringParameters) GetBillingCycleOk() (*GenericNameIdDTO, bool)`

GetBillingCycleOk returns a tuple with the BillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingCycle

`func (o *AgreementRecurringParameters) SetBillingCycle(v GenericNameIdDTO)`

SetBillingCycle sets BillingCycle field to given value.

### HasBillingCycle

`func (o *AgreementRecurringParameters) HasBillingCycle() bool`

HasBillingCycle returns a boolean if a field has been set.

### GetCycleBase

`func (o *AgreementRecurringParameters) GetCycleBase() GenericNameIdDTO`

GetCycleBase returns the CycleBase field if non-nil, zero value otherwise.

### GetCycleBaseOk

`func (o *AgreementRecurringParameters) GetCycleBaseOk() (*GenericNameIdDTO, bool)`

GetCycleBaseOk returns a tuple with the CycleBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleBase

`func (o *AgreementRecurringParameters) SetCycleBase(v GenericNameIdDTO)`

SetCycleBase sets CycleBase field to given value.

### HasCycleBase

`func (o *AgreementRecurringParameters) HasCycleBase() bool`

HasCycleBase returns a boolean if a field has been set.

### GetAGRAmount

`func (o *AgreementRecurringParameters) GetAGRAmount() float64`

GetAGRAmount returns the AGRAmount field if non-nil, zero value otherwise.

### GetAGRAmountOk

`func (o *AgreementRecurringParameters) GetAGRAmountOk() (*float64, bool)`

GetAGRAmountOk returns a tuple with the AGRAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAGRAmount

`func (o *AgreementRecurringParameters) SetAGRAmount(v float64)`

SetAGRAmount sets AGRAmount field to given value.

### HasAGRAmount

`func (o *AgreementRecurringParameters) HasAGRAmount() bool`

HasAGRAmount returns a boolean if a field has been set.

### SetAGRAmountNil

`func (o *AgreementRecurringParameters) SetAGRAmountNil(b bool)`

 SetAGRAmountNil sets the value for AGRAmount to be an explicit nil

### UnsetAGRAmount
`func (o *AgreementRecurringParameters) UnsetAGRAmount()`

UnsetAGRAmount ensures that no value is present for AGRAmount, not even an explicit nil
### GetTaxable

`func (o *AgreementRecurringParameters) GetTaxable() bool`

GetTaxable returns the Taxable field if non-nil, zero value otherwise.

### GetTaxableOk

`func (o *AgreementRecurringParameters) GetTaxableOk() (*bool, bool)`

GetTaxableOk returns a tuple with the Taxable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxable

`func (o *AgreementRecurringParameters) SetTaxable(v bool)`

SetTaxable sets Taxable field to given value.

### HasTaxable

`func (o *AgreementRecurringParameters) HasTaxable() bool`

HasTaxable returns a boolean if a field has been set.

### GetChildrenAmount

`func (o *AgreementRecurringParameters) GetChildrenAmount() float64`

GetChildrenAmount returns the ChildrenAmount field if non-nil, zero value otherwise.

### GetChildrenAmountOk

`func (o *AgreementRecurringParameters) GetChildrenAmountOk() (*float64, bool)`

GetChildrenAmountOk returns a tuple with the ChildrenAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenAmount

`func (o *AgreementRecurringParameters) SetChildrenAmount(v float64)`

SetChildrenAmount sets ChildrenAmount field to given value.

### HasChildrenAmount

`func (o *AgreementRecurringParameters) HasChildrenAmount() bool`

HasChildrenAmount returns a boolean if a field has been set.

### GetAdditionsAmount

`func (o *AgreementRecurringParameters) GetAdditionsAmount() float64`

GetAdditionsAmount returns the AdditionsAmount field if non-nil, zero value otherwise.

### GetAdditionsAmountOk

`func (o *AgreementRecurringParameters) GetAdditionsAmountOk() (*float64, bool)`

GetAdditionsAmountOk returns a tuple with the AdditionsAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionsAmount

`func (o *AgreementRecurringParameters) SetAdditionsAmount(v float64)`

SetAdditionsAmount sets AdditionsAmount field to given value.

### HasAdditionsAmount

`func (o *AgreementRecurringParameters) HasAdditionsAmount() bool`

HasAdditionsAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *AgreementRecurringParameters) GetTotalAmount() float64`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *AgreementRecurringParameters) GetTotalAmountOk() (*float64, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *AgreementRecurringParameters) SetTotalAmount(v float64)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *AgreementRecurringParameters) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetAGRProrate

`func (o *AgreementRecurringParameters) GetAGRProrate() float64`

GetAGRProrate returns the AGRProrate field if non-nil, zero value otherwise.

### GetAGRProrateOk

`func (o *AgreementRecurringParameters) GetAGRProrateOk() (*float64, bool)`

GetAGRProrateOk returns a tuple with the AGRProrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAGRProrate

`func (o *AgreementRecurringParameters) SetAGRProrate(v float64)`

SetAGRProrate sets AGRProrate field to given value.

### HasAGRProrate

`func (o *AgreementRecurringParameters) HasAGRProrate() bool`

HasAGRProrate returns a boolean if a field has been set.

### SetAGRProrateNil

`func (o *AgreementRecurringParameters) SetAGRProrateNil(b bool)`

 SetAGRProrateNil sets the value for AGRProrate to be an explicit nil

### UnsetAGRProrate
`func (o *AgreementRecurringParameters) UnsetAGRProrate()`

UnsetAGRProrate ensures that no value is present for AGRProrate, not even an explicit nil
### GetBillStartDate

`func (o *AgreementRecurringParameters) GetBillStartDate() string`

GetBillStartDate returns the BillStartDate field if non-nil, zero value otherwise.

### GetBillStartDateOk

`func (o *AgreementRecurringParameters) GetBillStartDateOk() (*string, bool)`

GetBillStartDateOk returns a tuple with the BillStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillStartDate

`func (o *AgreementRecurringParameters) SetBillStartDate(v string)`

SetBillStartDate sets BillStartDate field to given value.

### HasBillStartDate

`func (o *AgreementRecurringParameters) HasBillStartDate() bool`

HasBillStartDate returns a boolean if a field has been set.

### GetTaxCode

`func (o *AgreementRecurringParameters) GetTaxCode() GenericNameIdDTO`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *AgreementRecurringParameters) GetTaxCodeOk() (*GenericNameIdDTO, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *AgreementRecurringParameters) SetTaxCode(v GenericNameIdDTO)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *AgreementRecurringParameters) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTerms

`func (o *AgreementRecurringParameters) GetTerms() GenericNameIdDTO`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *AgreementRecurringParameters) GetTermsOk() (*GenericNameIdDTO, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *AgreementRecurringParameters) SetTerms(v GenericNameIdDTO)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *AgreementRecurringParameters) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetProrateFlag

`func (o *AgreementRecurringParameters) GetProrateFlag() bool`

GetProrateFlag returns the ProrateFlag field if non-nil, zero value otherwise.

### GetProrateFlagOk

`func (o *AgreementRecurringParameters) GetProrateFlagOk() (*bool, bool)`

GetProrateFlagOk returns a tuple with the ProrateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateFlag

`func (o *AgreementRecurringParameters) SetProrateFlag(v bool)`

SetProrateFlag sets ProrateFlag field to given value.

### HasProrateFlag

`func (o *AgreementRecurringParameters) HasProrateFlag() bool`

HasProrateFlag returns a boolean if a field has been set.

### GetInvoiceProratedAdditionsFlag

`func (o *AgreementRecurringParameters) GetInvoiceProratedAdditionsFlag() bool`

GetInvoiceProratedAdditionsFlag returns the InvoiceProratedAdditionsFlag field if non-nil, zero value otherwise.

### GetInvoiceProratedAdditionsFlagOk

`func (o *AgreementRecurringParameters) GetInvoiceProratedAdditionsFlagOk() (*bool, bool)`

GetInvoiceProratedAdditionsFlagOk returns a tuple with the InvoiceProratedAdditionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceProratedAdditionsFlag

`func (o *AgreementRecurringParameters) SetInvoiceProratedAdditionsFlag(v bool)`

SetInvoiceProratedAdditionsFlag sets InvoiceProratedAdditionsFlag field to given value.

### HasInvoiceProratedAdditionsFlag

`func (o *AgreementRecurringParameters) HasInvoiceProratedAdditionsFlag() bool`

HasInvoiceProratedAdditionsFlag returns a boolean if a field has been set.

### GetRestrictDownpayment

`func (o *AgreementRecurringParameters) GetRestrictDownpayment() bool`

GetRestrictDownpayment returns the RestrictDownpayment field if non-nil, zero value otherwise.

### GetRestrictDownpaymentOk

`func (o *AgreementRecurringParameters) GetRestrictDownpaymentOk() (*bool, bool)`

GetRestrictDownpaymentOk returns a tuple with the RestrictDownpayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictDownpayment

`func (o *AgreementRecurringParameters) SetRestrictDownpayment(v bool)`

SetRestrictDownpayment sets RestrictDownpayment field to given value.

### HasRestrictDownpayment

`func (o *AgreementRecurringParameters) HasRestrictDownpayment() bool`

HasRestrictDownpayment returns a boolean if a field has been set.

### GetCurrency

`func (o *AgreementRecurringParameters) GetCurrency() GenericNameIdDTO`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AgreementRecurringParameters) GetCurrencyOk() (*GenericNameIdDTO, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AgreementRecurringParameters) SetCurrency(v GenericNameIdDTO)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AgreementRecurringParameters) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetAutoInvoiceFlag

`func (o *AgreementRecurringParameters) GetAutoInvoiceFlag() bool`

GetAutoInvoiceFlag returns the AutoInvoiceFlag field if non-nil, zero value otherwise.

### GetAutoInvoiceFlagOk

`func (o *AgreementRecurringParameters) GetAutoInvoiceFlagOk() (*bool, bool)`

GetAutoInvoiceFlagOk returns a tuple with the AutoInvoiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoInvoiceFlag

`func (o *AgreementRecurringParameters) SetAutoInvoiceFlag(v bool)`

SetAutoInvoiceFlag sets AutoInvoiceFlag field to given value.

### HasAutoInvoiceFlag

`func (o *AgreementRecurringParameters) HasAutoInvoiceFlag() bool`

HasAutoInvoiceFlag returns a boolean if a field has been set.

### GetUserDefinedFieldValues

`func (o *AgreementRecurringParameters) GetUserDefinedFieldValues() []UserDefinedFieldValueModel`

GetUserDefinedFieldValues returns the UserDefinedFieldValues field if non-nil, zero value otherwise.

### GetUserDefinedFieldValuesOk

`func (o *AgreementRecurringParameters) GetUserDefinedFieldValuesOk() (*[]UserDefinedFieldValueModel, bool)`

GetUserDefinedFieldValuesOk returns a tuple with the UserDefinedFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFieldValues

`func (o *AgreementRecurringParameters) SetUserDefinedFieldValues(v []UserDefinedFieldValueModel)`

SetUserDefinedFieldValues sets UserDefinedFieldValues field to given value.

### HasUserDefinedFieldValues

`func (o *AgreementRecurringParameters) HasUserDefinedFieldValues() bool`

HasUserDefinedFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


