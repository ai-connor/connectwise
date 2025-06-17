# AgreementApplicationParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationUnit** | Pointer to [**AgreementApplicationUnit**](AgreementApplicationUnit.md) |  | [optional] 
**ApplicationLimit** | Pointer to [**AgreementApplicationLimit**](AgreementApplicationLimit.md) |  | [optional] 
**ApplicationLimitAmount** | Pointer to **float64** |  | [optional] 
**AvailablePer** | Pointer to [**AgreementApplicationAviablePer**](AgreementApplicationAviablePer.md) |  | [optional] 
**CoversTimeFlag** | Pointer to **bool** |  | [optional] 
**CoversExpensesFlag** | Pointer to **bool** |  | [optional] 
**CoversProductsFlag** | Pointer to **bool** |  | [optional] 
**CoversTaxFlag** | Pointer to **bool** |  | [optional] 
**CarryoverUnusedFlag** | Pointer to **bool** |  | [optional] 
**CarryOverDays** | Pointer to **NullableInt32** |  | [optional] 
**AllowOverrunsFlag** | Pointer to **bool** |  | [optional] 
**OverrunLimit** | Pointer to **NullableInt32** |  | [optional] 
**AgreementExpiresFlag** | Pointer to **bool** |  | [optional] 
**ChargeAdjustmentsFlag** | Pointer to **bool** |  | [optional] 
**PrepayFlag** | Pointer to **bool** |  | [optional] 
**AgrBillingCycle** | Pointer to [**AgreementApplicationBillingCycle**](AgreementApplicationBillingCycle.md) |  | [optional] 
**UserDefinedFieldValues** | Pointer to [**[]UserDefinedFieldValueModel**](UserDefinedFieldValueModel.md) |  | [optional] 

## Methods

### NewAgreementApplicationParameters

`func NewAgreementApplicationParameters() *AgreementApplicationParameters`

NewAgreementApplicationParameters instantiates a new AgreementApplicationParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementApplicationParametersWithDefaults

`func NewAgreementApplicationParametersWithDefaults() *AgreementApplicationParameters`

NewAgreementApplicationParametersWithDefaults instantiates a new AgreementApplicationParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationUnit

`func (o *AgreementApplicationParameters) GetApplicationUnit() AgreementApplicationUnit`

GetApplicationUnit returns the ApplicationUnit field if non-nil, zero value otherwise.

### GetApplicationUnitOk

`func (o *AgreementApplicationParameters) GetApplicationUnitOk() (*AgreementApplicationUnit, bool)`

GetApplicationUnitOk returns a tuple with the ApplicationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUnit

`func (o *AgreementApplicationParameters) SetApplicationUnit(v AgreementApplicationUnit)`

SetApplicationUnit sets ApplicationUnit field to given value.

### HasApplicationUnit

`func (o *AgreementApplicationParameters) HasApplicationUnit() bool`

HasApplicationUnit returns a boolean if a field has been set.

### GetApplicationLimit

`func (o *AgreementApplicationParameters) GetApplicationLimit() AgreementApplicationLimit`

GetApplicationLimit returns the ApplicationLimit field if non-nil, zero value otherwise.

### GetApplicationLimitOk

`func (o *AgreementApplicationParameters) GetApplicationLimitOk() (*AgreementApplicationLimit, bool)`

GetApplicationLimitOk returns a tuple with the ApplicationLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLimit

`func (o *AgreementApplicationParameters) SetApplicationLimit(v AgreementApplicationLimit)`

SetApplicationLimit sets ApplicationLimit field to given value.

### HasApplicationLimit

`func (o *AgreementApplicationParameters) HasApplicationLimit() bool`

HasApplicationLimit returns a boolean if a field has been set.

### GetApplicationLimitAmount

`func (o *AgreementApplicationParameters) GetApplicationLimitAmount() float64`

GetApplicationLimitAmount returns the ApplicationLimitAmount field if non-nil, zero value otherwise.

### GetApplicationLimitAmountOk

`func (o *AgreementApplicationParameters) GetApplicationLimitAmountOk() (*float64, bool)`

GetApplicationLimitAmountOk returns a tuple with the ApplicationLimitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLimitAmount

`func (o *AgreementApplicationParameters) SetApplicationLimitAmount(v float64)`

SetApplicationLimitAmount sets ApplicationLimitAmount field to given value.

### HasApplicationLimitAmount

`func (o *AgreementApplicationParameters) HasApplicationLimitAmount() bool`

HasApplicationLimitAmount returns a boolean if a field has been set.

### GetAvailablePer

`func (o *AgreementApplicationParameters) GetAvailablePer() AgreementApplicationAviablePer`

GetAvailablePer returns the AvailablePer field if non-nil, zero value otherwise.

### GetAvailablePerOk

`func (o *AgreementApplicationParameters) GetAvailablePerOk() (*AgreementApplicationAviablePer, bool)`

GetAvailablePerOk returns a tuple with the AvailablePer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailablePer

`func (o *AgreementApplicationParameters) SetAvailablePer(v AgreementApplicationAviablePer)`

SetAvailablePer sets AvailablePer field to given value.

### HasAvailablePer

`func (o *AgreementApplicationParameters) HasAvailablePer() bool`

HasAvailablePer returns a boolean if a field has been set.

### GetCoversTimeFlag

`func (o *AgreementApplicationParameters) GetCoversTimeFlag() bool`

GetCoversTimeFlag returns the CoversTimeFlag field if non-nil, zero value otherwise.

### GetCoversTimeFlagOk

`func (o *AgreementApplicationParameters) GetCoversTimeFlagOk() (*bool, bool)`

GetCoversTimeFlagOk returns a tuple with the CoversTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoversTimeFlag

`func (o *AgreementApplicationParameters) SetCoversTimeFlag(v bool)`

SetCoversTimeFlag sets CoversTimeFlag field to given value.

### HasCoversTimeFlag

`func (o *AgreementApplicationParameters) HasCoversTimeFlag() bool`

HasCoversTimeFlag returns a boolean if a field has been set.

### GetCoversExpensesFlag

`func (o *AgreementApplicationParameters) GetCoversExpensesFlag() bool`

GetCoversExpensesFlag returns the CoversExpensesFlag field if non-nil, zero value otherwise.

### GetCoversExpensesFlagOk

`func (o *AgreementApplicationParameters) GetCoversExpensesFlagOk() (*bool, bool)`

GetCoversExpensesFlagOk returns a tuple with the CoversExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoversExpensesFlag

`func (o *AgreementApplicationParameters) SetCoversExpensesFlag(v bool)`

SetCoversExpensesFlag sets CoversExpensesFlag field to given value.

### HasCoversExpensesFlag

`func (o *AgreementApplicationParameters) HasCoversExpensesFlag() bool`

HasCoversExpensesFlag returns a boolean if a field has been set.

### GetCoversProductsFlag

`func (o *AgreementApplicationParameters) GetCoversProductsFlag() bool`

GetCoversProductsFlag returns the CoversProductsFlag field if non-nil, zero value otherwise.

### GetCoversProductsFlagOk

`func (o *AgreementApplicationParameters) GetCoversProductsFlagOk() (*bool, bool)`

GetCoversProductsFlagOk returns a tuple with the CoversProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoversProductsFlag

`func (o *AgreementApplicationParameters) SetCoversProductsFlag(v bool)`

SetCoversProductsFlag sets CoversProductsFlag field to given value.

### HasCoversProductsFlag

`func (o *AgreementApplicationParameters) HasCoversProductsFlag() bool`

HasCoversProductsFlag returns a boolean if a field has been set.

### GetCoversTaxFlag

`func (o *AgreementApplicationParameters) GetCoversTaxFlag() bool`

GetCoversTaxFlag returns the CoversTaxFlag field if non-nil, zero value otherwise.

### GetCoversTaxFlagOk

`func (o *AgreementApplicationParameters) GetCoversTaxFlagOk() (*bool, bool)`

GetCoversTaxFlagOk returns a tuple with the CoversTaxFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoversTaxFlag

`func (o *AgreementApplicationParameters) SetCoversTaxFlag(v bool)`

SetCoversTaxFlag sets CoversTaxFlag field to given value.

### HasCoversTaxFlag

`func (o *AgreementApplicationParameters) HasCoversTaxFlag() bool`

HasCoversTaxFlag returns a boolean if a field has been set.

### GetCarryoverUnusedFlag

`func (o *AgreementApplicationParameters) GetCarryoverUnusedFlag() bool`

GetCarryoverUnusedFlag returns the CarryoverUnusedFlag field if non-nil, zero value otherwise.

### GetCarryoverUnusedFlagOk

`func (o *AgreementApplicationParameters) GetCarryoverUnusedFlagOk() (*bool, bool)`

GetCarryoverUnusedFlagOk returns a tuple with the CarryoverUnusedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryoverUnusedFlag

`func (o *AgreementApplicationParameters) SetCarryoverUnusedFlag(v bool)`

SetCarryoverUnusedFlag sets CarryoverUnusedFlag field to given value.

### HasCarryoverUnusedFlag

`func (o *AgreementApplicationParameters) HasCarryoverUnusedFlag() bool`

HasCarryoverUnusedFlag returns a boolean if a field has been set.

### GetCarryOverDays

`func (o *AgreementApplicationParameters) GetCarryOverDays() int32`

GetCarryOverDays returns the CarryOverDays field if non-nil, zero value otherwise.

### GetCarryOverDaysOk

`func (o *AgreementApplicationParameters) GetCarryOverDaysOk() (*int32, bool)`

GetCarryOverDaysOk returns a tuple with the CarryOverDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarryOverDays

`func (o *AgreementApplicationParameters) SetCarryOverDays(v int32)`

SetCarryOverDays sets CarryOverDays field to given value.

### HasCarryOverDays

`func (o *AgreementApplicationParameters) HasCarryOverDays() bool`

HasCarryOverDays returns a boolean if a field has been set.

### SetCarryOverDaysNil

`func (o *AgreementApplicationParameters) SetCarryOverDaysNil(b bool)`

 SetCarryOverDaysNil sets the value for CarryOverDays to be an explicit nil

### UnsetCarryOverDays
`func (o *AgreementApplicationParameters) UnsetCarryOverDays()`

UnsetCarryOverDays ensures that no value is present for CarryOverDays, not even an explicit nil
### GetAllowOverrunsFlag

`func (o *AgreementApplicationParameters) GetAllowOverrunsFlag() bool`

GetAllowOverrunsFlag returns the AllowOverrunsFlag field if non-nil, zero value otherwise.

### GetAllowOverrunsFlagOk

`func (o *AgreementApplicationParameters) GetAllowOverrunsFlagOk() (*bool, bool)`

GetAllowOverrunsFlagOk returns a tuple with the AllowOverrunsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverrunsFlag

`func (o *AgreementApplicationParameters) SetAllowOverrunsFlag(v bool)`

SetAllowOverrunsFlag sets AllowOverrunsFlag field to given value.

### HasAllowOverrunsFlag

`func (o *AgreementApplicationParameters) HasAllowOverrunsFlag() bool`

HasAllowOverrunsFlag returns a boolean if a field has been set.

### GetOverrunLimit

`func (o *AgreementApplicationParameters) GetOverrunLimit() int32`

GetOverrunLimit returns the OverrunLimit field if non-nil, zero value otherwise.

### GetOverrunLimitOk

`func (o *AgreementApplicationParameters) GetOverrunLimitOk() (*int32, bool)`

GetOverrunLimitOk returns a tuple with the OverrunLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrunLimit

`func (o *AgreementApplicationParameters) SetOverrunLimit(v int32)`

SetOverrunLimit sets OverrunLimit field to given value.

### HasOverrunLimit

`func (o *AgreementApplicationParameters) HasOverrunLimit() bool`

HasOverrunLimit returns a boolean if a field has been set.

### SetOverrunLimitNil

`func (o *AgreementApplicationParameters) SetOverrunLimitNil(b bool)`

 SetOverrunLimitNil sets the value for OverrunLimit to be an explicit nil

### UnsetOverrunLimit
`func (o *AgreementApplicationParameters) UnsetOverrunLimit()`

UnsetOverrunLimit ensures that no value is present for OverrunLimit, not even an explicit nil
### GetAgreementExpiresFlag

`func (o *AgreementApplicationParameters) GetAgreementExpiresFlag() bool`

GetAgreementExpiresFlag returns the AgreementExpiresFlag field if non-nil, zero value otherwise.

### GetAgreementExpiresFlagOk

`func (o *AgreementApplicationParameters) GetAgreementExpiresFlagOk() (*bool, bool)`

GetAgreementExpiresFlagOk returns a tuple with the AgreementExpiresFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementExpiresFlag

`func (o *AgreementApplicationParameters) SetAgreementExpiresFlag(v bool)`

SetAgreementExpiresFlag sets AgreementExpiresFlag field to given value.

### HasAgreementExpiresFlag

`func (o *AgreementApplicationParameters) HasAgreementExpiresFlag() bool`

HasAgreementExpiresFlag returns a boolean if a field has been set.

### GetChargeAdjustmentsFlag

`func (o *AgreementApplicationParameters) GetChargeAdjustmentsFlag() bool`

GetChargeAdjustmentsFlag returns the ChargeAdjustmentsFlag field if non-nil, zero value otherwise.

### GetChargeAdjustmentsFlagOk

`func (o *AgreementApplicationParameters) GetChargeAdjustmentsFlagOk() (*bool, bool)`

GetChargeAdjustmentsFlagOk returns a tuple with the ChargeAdjustmentsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeAdjustmentsFlag

`func (o *AgreementApplicationParameters) SetChargeAdjustmentsFlag(v bool)`

SetChargeAdjustmentsFlag sets ChargeAdjustmentsFlag field to given value.

### HasChargeAdjustmentsFlag

`func (o *AgreementApplicationParameters) HasChargeAdjustmentsFlag() bool`

HasChargeAdjustmentsFlag returns a boolean if a field has been set.

### GetPrepayFlag

`func (o *AgreementApplicationParameters) GetPrepayFlag() bool`

GetPrepayFlag returns the PrepayFlag field if non-nil, zero value otherwise.

### GetPrepayFlagOk

`func (o *AgreementApplicationParameters) GetPrepayFlagOk() (*bool, bool)`

GetPrepayFlagOk returns a tuple with the PrepayFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepayFlag

`func (o *AgreementApplicationParameters) SetPrepayFlag(v bool)`

SetPrepayFlag sets PrepayFlag field to given value.

### HasPrepayFlag

`func (o *AgreementApplicationParameters) HasPrepayFlag() bool`

HasPrepayFlag returns a boolean if a field has been set.

### GetAgrBillingCycle

`func (o *AgreementApplicationParameters) GetAgrBillingCycle() AgreementApplicationBillingCycle`

GetAgrBillingCycle returns the AgrBillingCycle field if non-nil, zero value otherwise.

### GetAgrBillingCycleOk

`func (o *AgreementApplicationParameters) GetAgrBillingCycleOk() (*AgreementApplicationBillingCycle, bool)`

GetAgrBillingCycleOk returns a tuple with the AgrBillingCycle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgrBillingCycle

`func (o *AgreementApplicationParameters) SetAgrBillingCycle(v AgreementApplicationBillingCycle)`

SetAgrBillingCycle sets AgrBillingCycle field to given value.

### HasAgrBillingCycle

`func (o *AgreementApplicationParameters) HasAgrBillingCycle() bool`

HasAgrBillingCycle returns a boolean if a field has been set.

### GetUserDefinedFieldValues

`func (o *AgreementApplicationParameters) GetUserDefinedFieldValues() []UserDefinedFieldValueModel`

GetUserDefinedFieldValues returns the UserDefinedFieldValues field if non-nil, zero value otherwise.

### GetUserDefinedFieldValuesOk

`func (o *AgreementApplicationParameters) GetUserDefinedFieldValuesOk() (*[]UserDefinedFieldValueModel, bool)`

GetUserDefinedFieldValuesOk returns a tuple with the UserDefinedFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFieldValues

`func (o *AgreementApplicationParameters) SetUserDefinedFieldValues(v []UserDefinedFieldValueModel)`

SetUserDefinedFieldValues sets UserDefinedFieldValues field to given value.

### HasUserDefinedFieldValues

`func (o *AgreementApplicationParameters) HasUserDefinedFieldValues() bool`

HasUserDefinedFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


