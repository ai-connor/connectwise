# ProductRecurring

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecurringRevenue** | Pointer to **NullableFloat64** |  | [optional] 
**RecurringCost** | Pointer to **NullableFloat64** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** | The Recurring End Date is calculated based on the             start date, number of cycles, and cycle type. | [optional] 
**BillCycleId** | Pointer to **NullableInt32** |  | [optional] 
**BillCycle** | Pointer to [**BillingCycleReference**](BillingCycleReference.md) |  | [optional] 
**Cycles** | Pointer to **NullableInt32** |  | [optional] 
**CycleType** | Pointer to **NullableString** |  | [optional] 
**AgreementType** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 

## Methods

### NewProductRecurring

`func NewProductRecurring() *ProductRecurring`

NewProductRecurring instantiates a new ProductRecurring object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductRecurringWithDefaults

`func NewProductRecurringWithDefaults() *ProductRecurring`

NewProductRecurringWithDefaults instantiates a new ProductRecurring object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecurringRevenue

`func (o *ProductRecurring) GetRecurringRevenue() float64`

GetRecurringRevenue returns the RecurringRevenue field if non-nil, zero value otherwise.

### GetRecurringRevenueOk

`func (o *ProductRecurring) GetRecurringRevenueOk() (*float64, bool)`

GetRecurringRevenueOk returns a tuple with the RecurringRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringRevenue

`func (o *ProductRecurring) SetRecurringRevenue(v float64)`

SetRecurringRevenue sets RecurringRevenue field to given value.

### HasRecurringRevenue

`func (o *ProductRecurring) HasRecurringRevenue() bool`

HasRecurringRevenue returns a boolean if a field has been set.

### SetRecurringRevenueNil

`func (o *ProductRecurring) SetRecurringRevenueNil(b bool)`

 SetRecurringRevenueNil sets the value for RecurringRevenue to be an explicit nil

### UnsetRecurringRevenue
`func (o *ProductRecurring) UnsetRecurringRevenue()`

UnsetRecurringRevenue ensures that no value is present for RecurringRevenue, not even an explicit nil
### GetRecurringCost

`func (o *ProductRecurring) GetRecurringCost() float64`

GetRecurringCost returns the RecurringCost field if non-nil, zero value otherwise.

### GetRecurringCostOk

`func (o *ProductRecurring) GetRecurringCostOk() (*float64, bool)`

GetRecurringCostOk returns a tuple with the RecurringCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringCost

`func (o *ProductRecurring) SetRecurringCost(v float64)`

SetRecurringCost sets RecurringCost field to given value.

### HasRecurringCost

`func (o *ProductRecurring) HasRecurringCost() bool`

HasRecurringCost returns a boolean if a field has been set.

### SetRecurringCostNil

`func (o *ProductRecurring) SetRecurringCostNil(b bool)`

 SetRecurringCostNil sets the value for RecurringCost to be an explicit nil

### UnsetRecurringCost
`func (o *ProductRecurring) UnsetRecurringCost()`

UnsetRecurringCost ensures that no value is present for RecurringCost, not even an explicit nil
### GetStartDate

`func (o *ProductRecurring) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ProductRecurring) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ProductRecurring) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ProductRecurring) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ProductRecurring) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ProductRecurring) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ProductRecurring) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ProductRecurring) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetBillCycleId

`func (o *ProductRecurring) GetBillCycleId() int32`

GetBillCycleId returns the BillCycleId field if non-nil, zero value otherwise.

### GetBillCycleIdOk

`func (o *ProductRecurring) GetBillCycleIdOk() (*int32, bool)`

GetBillCycleIdOk returns a tuple with the BillCycleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCycleId

`func (o *ProductRecurring) SetBillCycleId(v int32)`

SetBillCycleId sets BillCycleId field to given value.

### HasBillCycleId

`func (o *ProductRecurring) HasBillCycleId() bool`

HasBillCycleId returns a boolean if a field has been set.

### SetBillCycleIdNil

`func (o *ProductRecurring) SetBillCycleIdNil(b bool)`

 SetBillCycleIdNil sets the value for BillCycleId to be an explicit nil

### UnsetBillCycleId
`func (o *ProductRecurring) UnsetBillCycleId()`

UnsetBillCycleId ensures that no value is present for BillCycleId, not even an explicit nil
### GetBillCycle

`func (o *ProductRecurring) GetBillCycle() BillingCycleReference`

GetBillCycle returns the BillCycle field if non-nil, zero value otherwise.

### GetBillCycleOk

`func (o *ProductRecurring) GetBillCycleOk() (*BillingCycleReference, bool)`

GetBillCycleOk returns a tuple with the BillCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCycle

`func (o *ProductRecurring) SetBillCycle(v BillingCycleReference)`

SetBillCycle sets BillCycle field to given value.

### HasBillCycle

`func (o *ProductRecurring) HasBillCycle() bool`

HasBillCycle returns a boolean if a field has been set.

### GetCycles

`func (o *ProductRecurring) GetCycles() int32`

GetCycles returns the Cycles field if non-nil, zero value otherwise.

### GetCyclesOk

`func (o *ProductRecurring) GetCyclesOk() (*int32, bool)`

GetCyclesOk returns a tuple with the Cycles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycles

`func (o *ProductRecurring) SetCycles(v int32)`

SetCycles sets Cycles field to given value.

### HasCycles

`func (o *ProductRecurring) HasCycles() bool`

HasCycles returns a boolean if a field has been set.

### SetCyclesNil

`func (o *ProductRecurring) SetCyclesNil(b bool)`

 SetCyclesNil sets the value for Cycles to be an explicit nil

### UnsetCycles
`func (o *ProductRecurring) UnsetCycles()`

UnsetCycles ensures that no value is present for Cycles, not even an explicit nil
### GetCycleType

`func (o *ProductRecurring) GetCycleType() string`

GetCycleType returns the CycleType field if non-nil, zero value otherwise.

### GetCycleTypeOk

`func (o *ProductRecurring) GetCycleTypeOk() (*string, bool)`

GetCycleTypeOk returns a tuple with the CycleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleType

`func (o *ProductRecurring) SetCycleType(v string)`

SetCycleType sets CycleType field to given value.

### HasCycleType

`func (o *ProductRecurring) HasCycleType() bool`

HasCycleType returns a boolean if a field has been set.

### SetCycleTypeNil

`func (o *ProductRecurring) SetCycleTypeNil(b bool)`

 SetCycleTypeNil sets the value for CycleType to be an explicit nil

### UnsetCycleType
`func (o *ProductRecurring) UnsetCycleType()`

UnsetCycleType ensures that no value is present for CycleType, not even an explicit nil
### GetAgreementType

`func (o *ProductRecurring) GetAgreementType() AgreementTypeReference`

GetAgreementType returns the AgreementType field if non-nil, zero value otherwise.

### GetAgreementTypeOk

`func (o *ProductRecurring) GetAgreementTypeOk() (*AgreementTypeReference, bool)`

GetAgreementTypeOk returns a tuple with the AgreementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementType

`func (o *ProductRecurring) SetAgreementType(v AgreementTypeReference)`

SetAgreementType sets AgreementType field to given value.

### HasAgreementType

`func (o *ProductRecurring) HasAgreementType() bool`

HasAgreementType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


