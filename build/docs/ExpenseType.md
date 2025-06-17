# ExpenseType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**AmountCaption** | **string** |  | 
**ReimbursementRate** | Pointer to **NullableFloat64** |  | [optional] 
**BillExpenses** | **NullableString** |  | 
**InvoiceMarkupOption** | **NullableString** |  | 
**InvoiceMarkupAmount** | Pointer to **NullableFloat64** |  | [optional] 
**AdvancedAmountFlag** | Pointer to **NullableBool** |  | [optional] 
**MileageFlag** | Pointer to **NullableBool** |  | [optional] 
**QuantityFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**MaxAmount** | Pointer to **NullableFloat64** |  | [optional] 
**IntegrationXRef** | Pointer to **string** |  Max length: 50; | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseType

`func NewExpenseType(name string, amountCaption string, billExpenses NullableString, invoiceMarkupOption NullableString, ) *ExpenseType`

NewExpenseType instantiates a new ExpenseType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTypeWithDefaults

`func NewExpenseTypeWithDefaults() *ExpenseType`

NewExpenseTypeWithDefaults instantiates a new ExpenseType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExpenseType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseType) SetName(v string)`

SetName sets Name field to given value.


### GetAmountCaption

`func (o *ExpenseType) GetAmountCaption() string`

GetAmountCaption returns the AmountCaption field if non-nil, zero value otherwise.

### GetAmountCaptionOk

`func (o *ExpenseType) GetAmountCaptionOk() (*string, bool)`

GetAmountCaptionOk returns a tuple with the AmountCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCaption

`func (o *ExpenseType) SetAmountCaption(v string)`

SetAmountCaption sets AmountCaption field to given value.


### GetReimbursementRate

`func (o *ExpenseType) GetReimbursementRate() float64`

GetReimbursementRate returns the ReimbursementRate field if non-nil, zero value otherwise.

### GetReimbursementRateOk

`func (o *ExpenseType) GetReimbursementRateOk() (*float64, bool)`

GetReimbursementRateOk returns a tuple with the ReimbursementRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReimbursementRate

`func (o *ExpenseType) SetReimbursementRate(v float64)`

SetReimbursementRate sets ReimbursementRate field to given value.

### HasReimbursementRate

`func (o *ExpenseType) HasReimbursementRate() bool`

HasReimbursementRate returns a boolean if a field has been set.

### SetReimbursementRateNil

`func (o *ExpenseType) SetReimbursementRateNil(b bool)`

 SetReimbursementRateNil sets the value for ReimbursementRate to be an explicit nil

### UnsetReimbursementRate
`func (o *ExpenseType) UnsetReimbursementRate()`

UnsetReimbursementRate ensures that no value is present for ReimbursementRate, not even an explicit nil
### GetBillExpenses

`func (o *ExpenseType) GetBillExpenses() string`

GetBillExpenses returns the BillExpenses field if non-nil, zero value otherwise.

### GetBillExpensesOk

`func (o *ExpenseType) GetBillExpensesOk() (*string, bool)`

GetBillExpensesOk returns a tuple with the BillExpenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillExpenses

`func (o *ExpenseType) SetBillExpenses(v string)`

SetBillExpenses sets BillExpenses field to given value.


### SetBillExpensesNil

`func (o *ExpenseType) SetBillExpensesNil(b bool)`

 SetBillExpensesNil sets the value for BillExpenses to be an explicit nil

### UnsetBillExpenses
`func (o *ExpenseType) UnsetBillExpenses()`

UnsetBillExpenses ensures that no value is present for BillExpenses, not even an explicit nil
### GetInvoiceMarkupOption

`func (o *ExpenseType) GetInvoiceMarkupOption() string`

GetInvoiceMarkupOption returns the InvoiceMarkupOption field if non-nil, zero value otherwise.

### GetInvoiceMarkupOptionOk

`func (o *ExpenseType) GetInvoiceMarkupOptionOk() (*string, bool)`

GetInvoiceMarkupOptionOk returns a tuple with the InvoiceMarkupOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMarkupOption

`func (o *ExpenseType) SetInvoiceMarkupOption(v string)`

SetInvoiceMarkupOption sets InvoiceMarkupOption field to given value.


### SetInvoiceMarkupOptionNil

`func (o *ExpenseType) SetInvoiceMarkupOptionNil(b bool)`

 SetInvoiceMarkupOptionNil sets the value for InvoiceMarkupOption to be an explicit nil

### UnsetInvoiceMarkupOption
`func (o *ExpenseType) UnsetInvoiceMarkupOption()`

UnsetInvoiceMarkupOption ensures that no value is present for InvoiceMarkupOption, not even an explicit nil
### GetInvoiceMarkupAmount

`func (o *ExpenseType) GetInvoiceMarkupAmount() float64`

GetInvoiceMarkupAmount returns the InvoiceMarkupAmount field if non-nil, zero value otherwise.

### GetInvoiceMarkupAmountOk

`func (o *ExpenseType) GetInvoiceMarkupAmountOk() (*float64, bool)`

GetInvoiceMarkupAmountOk returns a tuple with the InvoiceMarkupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceMarkupAmount

`func (o *ExpenseType) SetInvoiceMarkupAmount(v float64)`

SetInvoiceMarkupAmount sets InvoiceMarkupAmount field to given value.

### HasInvoiceMarkupAmount

`func (o *ExpenseType) HasInvoiceMarkupAmount() bool`

HasInvoiceMarkupAmount returns a boolean if a field has been set.

### SetInvoiceMarkupAmountNil

`func (o *ExpenseType) SetInvoiceMarkupAmountNil(b bool)`

 SetInvoiceMarkupAmountNil sets the value for InvoiceMarkupAmount to be an explicit nil

### UnsetInvoiceMarkupAmount
`func (o *ExpenseType) UnsetInvoiceMarkupAmount()`

UnsetInvoiceMarkupAmount ensures that no value is present for InvoiceMarkupAmount, not even an explicit nil
### GetAdvancedAmountFlag

`func (o *ExpenseType) GetAdvancedAmountFlag() bool`

GetAdvancedAmountFlag returns the AdvancedAmountFlag field if non-nil, zero value otherwise.

### GetAdvancedAmountFlagOk

`func (o *ExpenseType) GetAdvancedAmountFlagOk() (*bool, bool)`

GetAdvancedAmountFlagOk returns a tuple with the AdvancedAmountFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedAmountFlag

`func (o *ExpenseType) SetAdvancedAmountFlag(v bool)`

SetAdvancedAmountFlag sets AdvancedAmountFlag field to given value.

### HasAdvancedAmountFlag

`func (o *ExpenseType) HasAdvancedAmountFlag() bool`

HasAdvancedAmountFlag returns a boolean if a field has been set.

### SetAdvancedAmountFlagNil

`func (o *ExpenseType) SetAdvancedAmountFlagNil(b bool)`

 SetAdvancedAmountFlagNil sets the value for AdvancedAmountFlag to be an explicit nil

### UnsetAdvancedAmountFlag
`func (o *ExpenseType) UnsetAdvancedAmountFlag()`

UnsetAdvancedAmountFlag ensures that no value is present for AdvancedAmountFlag, not even an explicit nil
### GetMileageFlag

`func (o *ExpenseType) GetMileageFlag() bool`

GetMileageFlag returns the MileageFlag field if non-nil, zero value otherwise.

### GetMileageFlagOk

`func (o *ExpenseType) GetMileageFlagOk() (*bool, bool)`

GetMileageFlagOk returns a tuple with the MileageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMileageFlag

`func (o *ExpenseType) SetMileageFlag(v bool)`

SetMileageFlag sets MileageFlag field to given value.

### HasMileageFlag

`func (o *ExpenseType) HasMileageFlag() bool`

HasMileageFlag returns a boolean if a field has been set.

### SetMileageFlagNil

`func (o *ExpenseType) SetMileageFlagNil(b bool)`

 SetMileageFlagNil sets the value for MileageFlag to be an explicit nil

### UnsetMileageFlag
`func (o *ExpenseType) UnsetMileageFlag()`

UnsetMileageFlag ensures that no value is present for MileageFlag, not even an explicit nil
### GetQuantityFlag

`func (o *ExpenseType) GetQuantityFlag() bool`

GetQuantityFlag returns the QuantityFlag field if non-nil, zero value otherwise.

### GetQuantityFlagOk

`func (o *ExpenseType) GetQuantityFlagOk() (*bool, bool)`

GetQuantityFlagOk returns a tuple with the QuantityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityFlag

`func (o *ExpenseType) SetQuantityFlag(v bool)`

SetQuantityFlag sets QuantityFlag field to given value.

### HasQuantityFlag

`func (o *ExpenseType) HasQuantityFlag() bool`

HasQuantityFlag returns a boolean if a field has been set.

### SetQuantityFlagNil

`func (o *ExpenseType) SetQuantityFlagNil(b bool)`

 SetQuantityFlagNil sets the value for QuantityFlag to be an explicit nil

### UnsetQuantityFlag
`func (o *ExpenseType) UnsetQuantityFlag()`

UnsetQuantityFlag ensures that no value is present for QuantityFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ExpenseType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ExpenseType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ExpenseType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ExpenseType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ExpenseType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ExpenseType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetMaxAmount

`func (o *ExpenseType) GetMaxAmount() float64`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *ExpenseType) GetMaxAmountOk() (*float64, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *ExpenseType) SetMaxAmount(v float64)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *ExpenseType) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### SetMaxAmountNil

`func (o *ExpenseType) SetMaxAmountNil(b bool)`

 SetMaxAmountNil sets the value for MaxAmount to be an explicit nil

### UnsetMaxAmount
`func (o *ExpenseType) UnsetMaxAmount()`

UnsetMaxAmount ensures that no value is present for MaxAmount, not even an explicit nil
### GetIntegrationXRef

`func (o *ExpenseType) GetIntegrationXRef() string`

GetIntegrationXRef returns the IntegrationXRef field if non-nil, zero value otherwise.

### GetIntegrationXRefOk

`func (o *ExpenseType) GetIntegrationXRefOk() (*string, bool)`

GetIntegrationXRefOk returns a tuple with the IntegrationXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXRef

`func (o *ExpenseType) SetIntegrationXRef(v string)`

SetIntegrationXRef sets IntegrationXRef field to given value.

### HasIntegrationXRef

`func (o *ExpenseType) HasIntegrationXRef() bool`

HasIntegrationXRef returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ExpenseType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ExpenseType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ExpenseType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ExpenseType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ExpenseType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ExpenseType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *ExpenseType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


