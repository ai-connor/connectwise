# TaxCodeLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TaxLevel** | Pointer to **int32** |  | [optional] 
**TaxRate** | **NullableFloat64** |  | 
**RateType** | **NullableString** |  | 
**TaxableMax** | Pointer to **NullableFloat64** |  | [optional] 
**Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**TaxCodeXref** | Pointer to **string** |  Max length: 50; | [optional] 
**AgencyXref** | Pointer to **string** |  Max length: 100; | [optional] 
**TaxServicesFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxExpensesFlag** | Pointer to **NullableBool** |  | [optional] 
**TaxProductsFlag** | Pointer to **NullableBool** |  | [optional] 
**SingleUnitFlag** | Pointer to **NullableBool** |  | [optional] 
**SingleUnitMinimum** | Pointer to **NullableFloat64** |  | [optional] 
**SingleUnitMaximum** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxCodeLevel

`func NewTaxCodeLevel(taxRate NullableFloat64, rateType NullableString, ) *TaxCodeLevel`

NewTaxCodeLevel instantiates a new TaxCodeLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeLevelWithDefaults

`func NewTaxCodeLevelWithDefaults() *TaxCodeLevel`

NewTaxCodeLevelWithDefaults instantiates a new TaxCodeLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxCodeLevel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCodeLevel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCodeLevel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCodeLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTaxLevel

`func (o *TaxCodeLevel) GetTaxLevel() int32`

GetTaxLevel returns the TaxLevel field if non-nil, zero value otherwise.

### GetTaxLevelOk

`func (o *TaxCodeLevel) GetTaxLevelOk() (*int32, bool)`

GetTaxLevelOk returns a tuple with the TaxLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLevel

`func (o *TaxCodeLevel) SetTaxLevel(v int32)`

SetTaxLevel sets TaxLevel field to given value.

### HasTaxLevel

`func (o *TaxCodeLevel) HasTaxLevel() bool`

HasTaxLevel returns a boolean if a field has been set.

### GetTaxRate

`func (o *TaxCodeLevel) GetTaxRate() float64`

GetTaxRate returns the TaxRate field if non-nil, zero value otherwise.

### GetTaxRateOk

`func (o *TaxCodeLevel) GetTaxRateOk() (*float64, bool)`

GetTaxRateOk returns a tuple with the TaxRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxRate

`func (o *TaxCodeLevel) SetTaxRate(v float64)`

SetTaxRate sets TaxRate field to given value.


### SetTaxRateNil

`func (o *TaxCodeLevel) SetTaxRateNil(b bool)`

 SetTaxRateNil sets the value for TaxRate to be an explicit nil

### UnsetTaxRate
`func (o *TaxCodeLevel) UnsetTaxRate()`

UnsetTaxRate ensures that no value is present for TaxRate, not even an explicit nil
### GetRateType

`func (o *TaxCodeLevel) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *TaxCodeLevel) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *TaxCodeLevel) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *TaxCodeLevel) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *TaxCodeLevel) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetTaxableMax

`func (o *TaxCodeLevel) GetTaxableMax() float64`

GetTaxableMax returns the TaxableMax field if non-nil, zero value otherwise.

### GetTaxableMaxOk

`func (o *TaxCodeLevel) GetTaxableMaxOk() (*float64, bool)`

GetTaxableMaxOk returns a tuple with the TaxableMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableMax

`func (o *TaxCodeLevel) SetTaxableMax(v float64)`

SetTaxableMax sets TaxableMax field to given value.

### HasTaxableMax

`func (o *TaxCodeLevel) HasTaxableMax() bool`

HasTaxableMax returns a boolean if a field has been set.

### SetTaxableMaxNil

`func (o *TaxCodeLevel) SetTaxableMaxNil(b bool)`

 SetTaxableMaxNil sets the value for TaxableMax to be an explicit nil

### UnsetTaxableMax
`func (o *TaxCodeLevel) UnsetTaxableMax()`

UnsetTaxableMax ensures that no value is present for TaxableMax, not even an explicit nil
### GetCaption

`func (o *TaxCodeLevel) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *TaxCodeLevel) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *TaxCodeLevel) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *TaxCodeLevel) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetTaxCodeXref

`func (o *TaxCodeLevel) GetTaxCodeXref() string`

GetTaxCodeXref returns the TaxCodeXref field if non-nil, zero value otherwise.

### GetTaxCodeXrefOk

`func (o *TaxCodeLevel) GetTaxCodeXrefOk() (*string, bool)`

GetTaxCodeXrefOk returns a tuple with the TaxCodeXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCodeXref

`func (o *TaxCodeLevel) SetTaxCodeXref(v string)`

SetTaxCodeXref sets TaxCodeXref field to given value.

### HasTaxCodeXref

`func (o *TaxCodeLevel) HasTaxCodeXref() bool`

HasTaxCodeXref returns a boolean if a field has been set.

### GetAgencyXref

`func (o *TaxCodeLevel) GetAgencyXref() string`

GetAgencyXref returns the AgencyXref field if non-nil, zero value otherwise.

### GetAgencyXrefOk

`func (o *TaxCodeLevel) GetAgencyXrefOk() (*string, bool)`

GetAgencyXrefOk returns a tuple with the AgencyXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgencyXref

`func (o *TaxCodeLevel) SetAgencyXref(v string)`

SetAgencyXref sets AgencyXref field to given value.

### HasAgencyXref

`func (o *TaxCodeLevel) HasAgencyXref() bool`

HasAgencyXref returns a boolean if a field has been set.

### GetTaxServicesFlag

`func (o *TaxCodeLevel) GetTaxServicesFlag() bool`

GetTaxServicesFlag returns the TaxServicesFlag field if non-nil, zero value otherwise.

### GetTaxServicesFlagOk

`func (o *TaxCodeLevel) GetTaxServicesFlagOk() (*bool, bool)`

GetTaxServicesFlagOk returns a tuple with the TaxServicesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxServicesFlag

`func (o *TaxCodeLevel) SetTaxServicesFlag(v bool)`

SetTaxServicesFlag sets TaxServicesFlag field to given value.

### HasTaxServicesFlag

`func (o *TaxCodeLevel) HasTaxServicesFlag() bool`

HasTaxServicesFlag returns a boolean if a field has been set.

### SetTaxServicesFlagNil

`func (o *TaxCodeLevel) SetTaxServicesFlagNil(b bool)`

 SetTaxServicesFlagNil sets the value for TaxServicesFlag to be an explicit nil

### UnsetTaxServicesFlag
`func (o *TaxCodeLevel) UnsetTaxServicesFlag()`

UnsetTaxServicesFlag ensures that no value is present for TaxServicesFlag, not even an explicit nil
### GetTaxExpensesFlag

`func (o *TaxCodeLevel) GetTaxExpensesFlag() bool`

GetTaxExpensesFlag returns the TaxExpensesFlag field if non-nil, zero value otherwise.

### GetTaxExpensesFlagOk

`func (o *TaxCodeLevel) GetTaxExpensesFlagOk() (*bool, bool)`

GetTaxExpensesFlagOk returns a tuple with the TaxExpensesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxExpensesFlag

`func (o *TaxCodeLevel) SetTaxExpensesFlag(v bool)`

SetTaxExpensesFlag sets TaxExpensesFlag field to given value.

### HasTaxExpensesFlag

`func (o *TaxCodeLevel) HasTaxExpensesFlag() bool`

HasTaxExpensesFlag returns a boolean if a field has been set.

### SetTaxExpensesFlagNil

`func (o *TaxCodeLevel) SetTaxExpensesFlagNil(b bool)`

 SetTaxExpensesFlagNil sets the value for TaxExpensesFlag to be an explicit nil

### UnsetTaxExpensesFlag
`func (o *TaxCodeLevel) UnsetTaxExpensesFlag()`

UnsetTaxExpensesFlag ensures that no value is present for TaxExpensesFlag, not even an explicit nil
### GetTaxProductsFlag

`func (o *TaxCodeLevel) GetTaxProductsFlag() bool`

GetTaxProductsFlag returns the TaxProductsFlag field if non-nil, zero value otherwise.

### GetTaxProductsFlagOk

`func (o *TaxCodeLevel) GetTaxProductsFlagOk() (*bool, bool)`

GetTaxProductsFlagOk returns a tuple with the TaxProductsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxProductsFlag

`func (o *TaxCodeLevel) SetTaxProductsFlag(v bool)`

SetTaxProductsFlag sets TaxProductsFlag field to given value.

### HasTaxProductsFlag

`func (o *TaxCodeLevel) HasTaxProductsFlag() bool`

HasTaxProductsFlag returns a boolean if a field has been set.

### SetTaxProductsFlagNil

`func (o *TaxCodeLevel) SetTaxProductsFlagNil(b bool)`

 SetTaxProductsFlagNil sets the value for TaxProductsFlag to be an explicit nil

### UnsetTaxProductsFlag
`func (o *TaxCodeLevel) UnsetTaxProductsFlag()`

UnsetTaxProductsFlag ensures that no value is present for TaxProductsFlag, not even an explicit nil
### GetSingleUnitFlag

`func (o *TaxCodeLevel) GetSingleUnitFlag() bool`

GetSingleUnitFlag returns the SingleUnitFlag field if non-nil, zero value otherwise.

### GetSingleUnitFlagOk

`func (o *TaxCodeLevel) GetSingleUnitFlagOk() (*bool, bool)`

GetSingleUnitFlagOk returns a tuple with the SingleUnitFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUnitFlag

`func (o *TaxCodeLevel) SetSingleUnitFlag(v bool)`

SetSingleUnitFlag sets SingleUnitFlag field to given value.

### HasSingleUnitFlag

`func (o *TaxCodeLevel) HasSingleUnitFlag() bool`

HasSingleUnitFlag returns a boolean if a field has been set.

### SetSingleUnitFlagNil

`func (o *TaxCodeLevel) SetSingleUnitFlagNil(b bool)`

 SetSingleUnitFlagNil sets the value for SingleUnitFlag to be an explicit nil

### UnsetSingleUnitFlag
`func (o *TaxCodeLevel) UnsetSingleUnitFlag()`

UnsetSingleUnitFlag ensures that no value is present for SingleUnitFlag, not even an explicit nil
### GetSingleUnitMinimum

`func (o *TaxCodeLevel) GetSingleUnitMinimum() float64`

GetSingleUnitMinimum returns the SingleUnitMinimum field if non-nil, zero value otherwise.

### GetSingleUnitMinimumOk

`func (o *TaxCodeLevel) GetSingleUnitMinimumOk() (*float64, bool)`

GetSingleUnitMinimumOk returns a tuple with the SingleUnitMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUnitMinimum

`func (o *TaxCodeLevel) SetSingleUnitMinimum(v float64)`

SetSingleUnitMinimum sets SingleUnitMinimum field to given value.

### HasSingleUnitMinimum

`func (o *TaxCodeLevel) HasSingleUnitMinimum() bool`

HasSingleUnitMinimum returns a boolean if a field has been set.

### SetSingleUnitMinimumNil

`func (o *TaxCodeLevel) SetSingleUnitMinimumNil(b bool)`

 SetSingleUnitMinimumNil sets the value for SingleUnitMinimum to be an explicit nil

### UnsetSingleUnitMinimum
`func (o *TaxCodeLevel) UnsetSingleUnitMinimum()`

UnsetSingleUnitMinimum ensures that no value is present for SingleUnitMinimum, not even an explicit nil
### GetSingleUnitMaximum

`func (o *TaxCodeLevel) GetSingleUnitMaximum() float64`

GetSingleUnitMaximum returns the SingleUnitMaximum field if non-nil, zero value otherwise.

### GetSingleUnitMaximumOk

`func (o *TaxCodeLevel) GetSingleUnitMaximumOk() (*float64, bool)`

GetSingleUnitMaximumOk returns a tuple with the SingleUnitMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleUnitMaximum

`func (o *TaxCodeLevel) SetSingleUnitMaximum(v float64)`

SetSingleUnitMaximum sets SingleUnitMaximum field to given value.

### HasSingleUnitMaximum

`func (o *TaxCodeLevel) HasSingleUnitMaximum() bool`

HasSingleUnitMaximum returns a boolean if a field has been set.

### SetSingleUnitMaximumNil

`func (o *TaxCodeLevel) SetSingleUnitMaximumNil(b bool)`

 SetSingleUnitMaximumNil sets the value for SingleUnitMaximum to be an explicit nil

### UnsetSingleUnitMaximum
`func (o *TaxCodeLevel) UnsetSingleUnitMaximum()`

UnsetSingleUnitMaximum ensures that no value is present for SingleUnitMaximum, not even an explicit nil
### GetInfo

`func (o *TaxCodeLevel) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxCodeLevel) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxCodeLevel) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxCodeLevel) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


