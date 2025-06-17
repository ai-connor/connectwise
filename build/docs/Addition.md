# Addition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Product** | [**IvItemReference**](IvItemReference.md) |  | 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**LessIncluded** | Pointer to **NullableFloat64** |  | [optional] 
**UnitPrice** | Pointer to **NullableFloat64** |  | [optional] 
**UnitCost** | Pointer to **NullableFloat64** |  | [optional] 
**BillCustomer** | **NullableString** |  | 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**CancelledDate** | Pointer to **time.Time** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**SerialNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**InvoiceDescription** | Pointer to **string** |  Max length: 6000; | [optional] 
**PurchaseItemFlag** | Pointer to **NullableBool** |  | [optional] 
**SpecialOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**BilledQuantity** | Pointer to **NullableFloat64** |  | [optional] 
**Uom** | Pointer to **string** |  | [optional] 
**ExtPrice** | Pointer to **NullableFloat64** |  | [optional] 
**ExtCost** | Pointer to **NullableFloat64** |  | [optional] 
**SequenceNumber** | Pointer to **NullableFloat64** |  | [optional] 
**Margin** | Pointer to **NullableFloat64** |  | [optional] 
**ProrateCost** | Pointer to **NullableFloat64** |  | [optional] 
**ProratePrice** | Pointer to **NullableFloat64** |  | [optional] 
**ExtendedProrateCost** | Pointer to **NullableFloat64** |  | [optional] 
**ExtendedProratePrice** | Pointer to **NullableFloat64** |  | [optional] 
**ProrateCurrentPeriodFlag** | Pointer to **NullableBool** |  | [optional] 
**Opportunity** | Pointer to [**OpportunityReference**](OpportunityReference.md) |  | [optional] 
**AgreementStatus** | Pointer to **NullableString** |  | [optional] 
**InvoiceGrouping** | Pointer to [**InvoiceGroupingReference**](InvoiceGroupingReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewAddition

`func NewAddition(product IvItemReference, billCustomer NullableString, ) *Addition`

NewAddition instantiates a new Addition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdditionWithDefaults

`func NewAdditionWithDefaults() *Addition`

NewAdditionWithDefaults instantiates a new Addition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Addition) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Addition) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Addition) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Addition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProduct

`func (o *Addition) GetProduct() IvItemReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *Addition) GetProductOk() (*IvItemReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *Addition) SetProduct(v IvItemReference)`

SetProduct sets Product field to given value.


### GetQuantity

`func (o *Addition) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Addition) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Addition) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Addition) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *Addition) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *Addition) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetLessIncluded

`func (o *Addition) GetLessIncluded() float64`

GetLessIncluded returns the LessIncluded field if non-nil, zero value otherwise.

### GetLessIncludedOk

`func (o *Addition) GetLessIncludedOk() (*float64, bool)`

GetLessIncludedOk returns a tuple with the LessIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLessIncluded

`func (o *Addition) SetLessIncluded(v float64)`

SetLessIncluded sets LessIncluded field to given value.

### HasLessIncluded

`func (o *Addition) HasLessIncluded() bool`

HasLessIncluded returns a boolean if a field has been set.

### SetLessIncludedNil

`func (o *Addition) SetLessIncludedNil(b bool)`

 SetLessIncludedNil sets the value for LessIncluded to be an explicit nil

### UnsetLessIncluded
`func (o *Addition) UnsetLessIncluded()`

UnsetLessIncluded ensures that no value is present for LessIncluded, not even an explicit nil
### GetUnitPrice

`func (o *Addition) GetUnitPrice() float64`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *Addition) GetUnitPriceOk() (*float64, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *Addition) SetUnitPrice(v float64)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *Addition) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.

### SetUnitPriceNil

`func (o *Addition) SetUnitPriceNil(b bool)`

 SetUnitPriceNil sets the value for UnitPrice to be an explicit nil

### UnsetUnitPrice
`func (o *Addition) UnsetUnitPrice()`

UnsetUnitPrice ensures that no value is present for UnitPrice, not even an explicit nil
### GetUnitCost

`func (o *Addition) GetUnitCost() float64`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *Addition) GetUnitCostOk() (*float64, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *Addition) SetUnitCost(v float64)`

SetUnitCost sets UnitCost field to given value.

### HasUnitCost

`func (o *Addition) HasUnitCost() bool`

HasUnitCost returns a boolean if a field has been set.

### SetUnitCostNil

`func (o *Addition) SetUnitCostNil(b bool)`

 SetUnitCostNil sets the value for UnitCost to be an explicit nil

### UnsetUnitCost
`func (o *Addition) UnsetUnitCost()`

UnsetUnitCost ensures that no value is present for UnitCost, not even an explicit nil
### GetBillCustomer

`func (o *Addition) GetBillCustomer() string`

GetBillCustomer returns the BillCustomer field if non-nil, zero value otherwise.

### GetBillCustomerOk

`func (o *Addition) GetBillCustomerOk() (*string, bool)`

GetBillCustomerOk returns a tuple with the BillCustomer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCustomer

`func (o *Addition) SetBillCustomer(v string)`

SetBillCustomer sets BillCustomer field to given value.


### SetBillCustomerNil

`func (o *Addition) SetBillCustomerNil(b bool)`

 SetBillCustomerNil sets the value for BillCustomer to be an explicit nil

### UnsetBillCustomer
`func (o *Addition) UnsetBillCustomer()`

UnsetBillCustomer ensures that no value is present for BillCustomer, not even an explicit nil
### GetEffectiveDate

`func (o *Addition) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *Addition) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *Addition) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *Addition) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCancelledDate

`func (o *Addition) GetCancelledDate() time.Time`

GetCancelledDate returns the CancelledDate field if non-nil, zero value otherwise.

### GetCancelledDateOk

`func (o *Addition) GetCancelledDateOk() (*time.Time, bool)`

GetCancelledDateOk returns a tuple with the CancelledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledDate

`func (o *Addition) SetCancelledDate(v time.Time)`

SetCancelledDate sets CancelledDate field to given value.

### HasCancelledDate

`func (o *Addition) HasCancelledDate() bool`

HasCancelledDate returns a boolean if a field has been set.

### GetTaxableFlag

`func (o *Addition) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *Addition) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *Addition) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *Addition) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *Addition) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *Addition) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetSerialNumber

`func (o *Addition) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Addition) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Addition) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Addition) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetInvoiceDescription

`func (o *Addition) GetInvoiceDescription() string`

GetInvoiceDescription returns the InvoiceDescription field if non-nil, zero value otherwise.

### GetInvoiceDescriptionOk

`func (o *Addition) GetInvoiceDescriptionOk() (*string, bool)`

GetInvoiceDescriptionOk returns a tuple with the InvoiceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceDescription

`func (o *Addition) SetInvoiceDescription(v string)`

SetInvoiceDescription sets InvoiceDescription field to given value.

### HasInvoiceDescription

`func (o *Addition) HasInvoiceDescription() bool`

HasInvoiceDescription returns a boolean if a field has been set.

### GetPurchaseItemFlag

`func (o *Addition) GetPurchaseItemFlag() bool`

GetPurchaseItemFlag returns the PurchaseItemFlag field if non-nil, zero value otherwise.

### GetPurchaseItemFlagOk

`func (o *Addition) GetPurchaseItemFlagOk() (*bool, bool)`

GetPurchaseItemFlagOk returns a tuple with the PurchaseItemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseItemFlag

`func (o *Addition) SetPurchaseItemFlag(v bool)`

SetPurchaseItemFlag sets PurchaseItemFlag field to given value.

### HasPurchaseItemFlag

`func (o *Addition) HasPurchaseItemFlag() bool`

HasPurchaseItemFlag returns a boolean if a field has been set.

### SetPurchaseItemFlagNil

`func (o *Addition) SetPurchaseItemFlagNil(b bool)`

 SetPurchaseItemFlagNil sets the value for PurchaseItemFlag to be an explicit nil

### UnsetPurchaseItemFlag
`func (o *Addition) UnsetPurchaseItemFlag()`

UnsetPurchaseItemFlag ensures that no value is present for PurchaseItemFlag, not even an explicit nil
### GetSpecialOrderFlag

`func (o *Addition) GetSpecialOrderFlag() bool`

GetSpecialOrderFlag returns the SpecialOrderFlag field if non-nil, zero value otherwise.

### GetSpecialOrderFlagOk

`func (o *Addition) GetSpecialOrderFlagOk() (*bool, bool)`

GetSpecialOrderFlagOk returns a tuple with the SpecialOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOrderFlag

`func (o *Addition) SetSpecialOrderFlag(v bool)`

SetSpecialOrderFlag sets SpecialOrderFlag field to given value.

### HasSpecialOrderFlag

`func (o *Addition) HasSpecialOrderFlag() bool`

HasSpecialOrderFlag returns a boolean if a field has been set.

### SetSpecialOrderFlagNil

`func (o *Addition) SetSpecialOrderFlagNil(b bool)`

 SetSpecialOrderFlagNil sets the value for SpecialOrderFlag to be an explicit nil

### UnsetSpecialOrderFlag
`func (o *Addition) UnsetSpecialOrderFlag()`

UnsetSpecialOrderFlag ensures that no value is present for SpecialOrderFlag, not even an explicit nil
### GetAgreementId

`func (o *Addition) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *Addition) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *Addition) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *Addition) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *Addition) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *Addition) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetDescription

`func (o *Addition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Addition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Addition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Addition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBilledQuantity

`func (o *Addition) GetBilledQuantity() float64`

GetBilledQuantity returns the BilledQuantity field if non-nil, zero value otherwise.

### GetBilledQuantityOk

`func (o *Addition) GetBilledQuantityOk() (*float64, bool)`

GetBilledQuantityOk returns a tuple with the BilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledQuantity

`func (o *Addition) SetBilledQuantity(v float64)`

SetBilledQuantity sets BilledQuantity field to given value.

### HasBilledQuantity

`func (o *Addition) HasBilledQuantity() bool`

HasBilledQuantity returns a boolean if a field has been set.

### SetBilledQuantityNil

`func (o *Addition) SetBilledQuantityNil(b bool)`

 SetBilledQuantityNil sets the value for BilledQuantity to be an explicit nil

### UnsetBilledQuantity
`func (o *Addition) UnsetBilledQuantity()`

UnsetBilledQuantity ensures that no value is present for BilledQuantity, not even an explicit nil
### GetUom

`func (o *Addition) GetUom() string`

GetUom returns the Uom field if non-nil, zero value otherwise.

### GetUomOk

`func (o *Addition) GetUomOk() (*string, bool)`

GetUomOk returns a tuple with the Uom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUom

`func (o *Addition) SetUom(v string)`

SetUom sets Uom field to given value.

### HasUom

`func (o *Addition) HasUom() bool`

HasUom returns a boolean if a field has been set.

### GetExtPrice

`func (o *Addition) GetExtPrice() float64`

GetExtPrice returns the ExtPrice field if non-nil, zero value otherwise.

### GetExtPriceOk

`func (o *Addition) GetExtPriceOk() (*float64, bool)`

GetExtPriceOk returns a tuple with the ExtPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtPrice

`func (o *Addition) SetExtPrice(v float64)`

SetExtPrice sets ExtPrice field to given value.

### HasExtPrice

`func (o *Addition) HasExtPrice() bool`

HasExtPrice returns a boolean if a field has been set.

### SetExtPriceNil

`func (o *Addition) SetExtPriceNil(b bool)`

 SetExtPriceNil sets the value for ExtPrice to be an explicit nil

### UnsetExtPrice
`func (o *Addition) UnsetExtPrice()`

UnsetExtPrice ensures that no value is present for ExtPrice, not even an explicit nil
### GetExtCost

`func (o *Addition) GetExtCost() float64`

GetExtCost returns the ExtCost field if non-nil, zero value otherwise.

### GetExtCostOk

`func (o *Addition) GetExtCostOk() (*float64, bool)`

GetExtCostOk returns a tuple with the ExtCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtCost

`func (o *Addition) SetExtCost(v float64)`

SetExtCost sets ExtCost field to given value.

### HasExtCost

`func (o *Addition) HasExtCost() bool`

HasExtCost returns a boolean if a field has been set.

### SetExtCostNil

`func (o *Addition) SetExtCostNil(b bool)`

 SetExtCostNil sets the value for ExtCost to be an explicit nil

### UnsetExtCost
`func (o *Addition) UnsetExtCost()`

UnsetExtCost ensures that no value is present for ExtCost, not even an explicit nil
### GetSequenceNumber

`func (o *Addition) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *Addition) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *Addition) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *Addition) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *Addition) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *Addition) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetMargin

`func (o *Addition) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *Addition) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *Addition) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *Addition) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### SetMarginNil

`func (o *Addition) SetMarginNil(b bool)`

 SetMarginNil sets the value for Margin to be an explicit nil

### UnsetMargin
`func (o *Addition) UnsetMargin()`

UnsetMargin ensures that no value is present for Margin, not even an explicit nil
### GetProrateCost

`func (o *Addition) GetProrateCost() float64`

GetProrateCost returns the ProrateCost field if non-nil, zero value otherwise.

### GetProrateCostOk

`func (o *Addition) GetProrateCostOk() (*float64, bool)`

GetProrateCostOk returns a tuple with the ProrateCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateCost

`func (o *Addition) SetProrateCost(v float64)`

SetProrateCost sets ProrateCost field to given value.

### HasProrateCost

`func (o *Addition) HasProrateCost() bool`

HasProrateCost returns a boolean if a field has been set.

### SetProrateCostNil

`func (o *Addition) SetProrateCostNil(b bool)`

 SetProrateCostNil sets the value for ProrateCost to be an explicit nil

### UnsetProrateCost
`func (o *Addition) UnsetProrateCost()`

UnsetProrateCost ensures that no value is present for ProrateCost, not even an explicit nil
### GetProratePrice

`func (o *Addition) GetProratePrice() float64`

GetProratePrice returns the ProratePrice field if non-nil, zero value otherwise.

### GetProratePriceOk

`func (o *Addition) GetProratePriceOk() (*float64, bool)`

GetProratePriceOk returns a tuple with the ProratePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProratePrice

`func (o *Addition) SetProratePrice(v float64)`

SetProratePrice sets ProratePrice field to given value.

### HasProratePrice

`func (o *Addition) HasProratePrice() bool`

HasProratePrice returns a boolean if a field has been set.

### SetProratePriceNil

`func (o *Addition) SetProratePriceNil(b bool)`

 SetProratePriceNil sets the value for ProratePrice to be an explicit nil

### UnsetProratePrice
`func (o *Addition) UnsetProratePrice()`

UnsetProratePrice ensures that no value is present for ProratePrice, not even an explicit nil
### GetExtendedProrateCost

`func (o *Addition) GetExtendedProrateCost() float64`

GetExtendedProrateCost returns the ExtendedProrateCost field if non-nil, zero value otherwise.

### GetExtendedProrateCostOk

`func (o *Addition) GetExtendedProrateCostOk() (*float64, bool)`

GetExtendedProrateCostOk returns a tuple with the ExtendedProrateCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProrateCost

`func (o *Addition) SetExtendedProrateCost(v float64)`

SetExtendedProrateCost sets ExtendedProrateCost field to given value.

### HasExtendedProrateCost

`func (o *Addition) HasExtendedProrateCost() bool`

HasExtendedProrateCost returns a boolean if a field has been set.

### SetExtendedProrateCostNil

`func (o *Addition) SetExtendedProrateCostNil(b bool)`

 SetExtendedProrateCostNil sets the value for ExtendedProrateCost to be an explicit nil

### UnsetExtendedProrateCost
`func (o *Addition) UnsetExtendedProrateCost()`

UnsetExtendedProrateCost ensures that no value is present for ExtendedProrateCost, not even an explicit nil
### GetExtendedProratePrice

`func (o *Addition) GetExtendedProratePrice() float64`

GetExtendedProratePrice returns the ExtendedProratePrice field if non-nil, zero value otherwise.

### GetExtendedProratePriceOk

`func (o *Addition) GetExtendedProratePriceOk() (*float64, bool)`

GetExtendedProratePriceOk returns a tuple with the ExtendedProratePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedProratePrice

`func (o *Addition) SetExtendedProratePrice(v float64)`

SetExtendedProratePrice sets ExtendedProratePrice field to given value.

### HasExtendedProratePrice

`func (o *Addition) HasExtendedProratePrice() bool`

HasExtendedProratePrice returns a boolean if a field has been set.

### SetExtendedProratePriceNil

`func (o *Addition) SetExtendedProratePriceNil(b bool)`

 SetExtendedProratePriceNil sets the value for ExtendedProratePrice to be an explicit nil

### UnsetExtendedProratePrice
`func (o *Addition) UnsetExtendedProratePrice()`

UnsetExtendedProratePrice ensures that no value is present for ExtendedProratePrice, not even an explicit nil
### GetProrateCurrentPeriodFlag

`func (o *Addition) GetProrateCurrentPeriodFlag() bool`

GetProrateCurrentPeriodFlag returns the ProrateCurrentPeriodFlag field if non-nil, zero value otherwise.

### GetProrateCurrentPeriodFlagOk

`func (o *Addition) GetProrateCurrentPeriodFlagOk() (*bool, bool)`

GetProrateCurrentPeriodFlagOk returns a tuple with the ProrateCurrentPeriodFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProrateCurrentPeriodFlag

`func (o *Addition) SetProrateCurrentPeriodFlag(v bool)`

SetProrateCurrentPeriodFlag sets ProrateCurrentPeriodFlag field to given value.

### HasProrateCurrentPeriodFlag

`func (o *Addition) HasProrateCurrentPeriodFlag() bool`

HasProrateCurrentPeriodFlag returns a boolean if a field has been set.

### SetProrateCurrentPeriodFlagNil

`func (o *Addition) SetProrateCurrentPeriodFlagNil(b bool)`

 SetProrateCurrentPeriodFlagNil sets the value for ProrateCurrentPeriodFlag to be an explicit nil

### UnsetProrateCurrentPeriodFlag
`func (o *Addition) UnsetProrateCurrentPeriodFlag()`

UnsetProrateCurrentPeriodFlag ensures that no value is present for ProrateCurrentPeriodFlag, not even an explicit nil
### GetOpportunity

`func (o *Addition) GetOpportunity() OpportunityReference`

GetOpportunity returns the Opportunity field if non-nil, zero value otherwise.

### GetOpportunityOk

`func (o *Addition) GetOpportunityOk() (*OpportunityReference, bool)`

GetOpportunityOk returns a tuple with the Opportunity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunity

`func (o *Addition) SetOpportunity(v OpportunityReference)`

SetOpportunity sets Opportunity field to given value.

### HasOpportunity

`func (o *Addition) HasOpportunity() bool`

HasOpportunity returns a boolean if a field has been set.

### GetAgreementStatus

`func (o *Addition) GetAgreementStatus() string`

GetAgreementStatus returns the AgreementStatus field if non-nil, zero value otherwise.

### GetAgreementStatusOk

`func (o *Addition) GetAgreementStatusOk() (*string, bool)`

GetAgreementStatusOk returns a tuple with the AgreementStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementStatus

`func (o *Addition) SetAgreementStatus(v string)`

SetAgreementStatus sets AgreementStatus field to given value.

### HasAgreementStatus

`func (o *Addition) HasAgreementStatus() bool`

HasAgreementStatus returns a boolean if a field has been set.

### SetAgreementStatusNil

`func (o *Addition) SetAgreementStatusNil(b bool)`

 SetAgreementStatusNil sets the value for AgreementStatus to be an explicit nil

### UnsetAgreementStatus
`func (o *Addition) UnsetAgreementStatus()`

UnsetAgreementStatus ensures that no value is present for AgreementStatus, not even an explicit nil
### GetInvoiceGrouping

`func (o *Addition) GetInvoiceGrouping() InvoiceGroupingReference`

GetInvoiceGrouping returns the InvoiceGrouping field if non-nil, zero value otherwise.

### GetInvoiceGroupingOk

`func (o *Addition) GetInvoiceGroupingOk() (*InvoiceGroupingReference, bool)`

GetInvoiceGroupingOk returns a tuple with the InvoiceGrouping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceGrouping

`func (o *Addition) SetInvoiceGrouping(v InvoiceGroupingReference)`

SetInvoiceGrouping sets InvoiceGrouping field to given value.

### HasInvoiceGrouping

`func (o *Addition) HasInvoiceGrouping() bool`

HasInvoiceGrouping returns a boolean if a field has been set.

### GetInfo

`func (o *Addition) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Addition) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Addition) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Addition) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *Addition) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *Addition) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *Addition) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *Addition) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


