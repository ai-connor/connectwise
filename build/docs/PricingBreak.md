# PricingBreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DetailId** | Pointer to **NullableInt32** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**QuantityStart** | **NullableFloat64** |  | 
**QuantityEnd** | Pointer to **NullableFloat64** |  | [optional] 
**Unlimited** | Pointer to **bool** |  | [optional] 
**PriceMethod** | **NullableString** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPricingBreak

`func NewPricingBreak(quantityStart NullableFloat64, priceMethod NullableString, ) *PricingBreak`

NewPricingBreak instantiates a new PricingBreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingBreakWithDefaults

`func NewPricingBreakWithDefaults() *PricingBreak`

NewPricingBreakWithDefaults instantiates a new PricingBreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingBreak) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingBreak) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingBreak) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PricingBreak) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDetailId

`func (o *PricingBreak) GetDetailId() int32`

GetDetailId returns the DetailId field if non-nil, zero value otherwise.

### GetDetailIdOk

`func (o *PricingBreak) GetDetailIdOk() (*int32, bool)`

GetDetailIdOk returns a tuple with the DetailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailId

`func (o *PricingBreak) SetDetailId(v int32)`

SetDetailId sets DetailId field to given value.

### HasDetailId

`func (o *PricingBreak) HasDetailId() bool`

HasDetailId returns a boolean if a field has been set.

### SetDetailIdNil

`func (o *PricingBreak) SetDetailIdNil(b bool)`

 SetDetailIdNil sets the value for DetailId to be an explicit nil

### UnsetDetailId
`func (o *PricingBreak) UnsetDetailId()`

UnsetDetailId ensures that no value is present for DetailId, not even an explicit nil
### GetAmount

`func (o *PricingBreak) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PricingBreak) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PricingBreak) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PricingBreak) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *PricingBreak) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *PricingBreak) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetQuantityStart

`func (o *PricingBreak) GetQuantityStart() float64`

GetQuantityStart returns the QuantityStart field if non-nil, zero value otherwise.

### GetQuantityStartOk

`func (o *PricingBreak) GetQuantityStartOk() (*float64, bool)`

GetQuantityStartOk returns a tuple with the QuantityStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityStart

`func (o *PricingBreak) SetQuantityStart(v float64)`

SetQuantityStart sets QuantityStart field to given value.


### SetQuantityStartNil

`func (o *PricingBreak) SetQuantityStartNil(b bool)`

 SetQuantityStartNil sets the value for QuantityStart to be an explicit nil

### UnsetQuantityStart
`func (o *PricingBreak) UnsetQuantityStart()`

UnsetQuantityStart ensures that no value is present for QuantityStart, not even an explicit nil
### GetQuantityEnd

`func (o *PricingBreak) GetQuantityEnd() float64`

GetQuantityEnd returns the QuantityEnd field if non-nil, zero value otherwise.

### GetQuantityEndOk

`func (o *PricingBreak) GetQuantityEndOk() (*float64, bool)`

GetQuantityEndOk returns a tuple with the QuantityEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityEnd

`func (o *PricingBreak) SetQuantityEnd(v float64)`

SetQuantityEnd sets QuantityEnd field to given value.

### HasQuantityEnd

`func (o *PricingBreak) HasQuantityEnd() bool`

HasQuantityEnd returns a boolean if a field has been set.

### SetQuantityEndNil

`func (o *PricingBreak) SetQuantityEndNil(b bool)`

 SetQuantityEndNil sets the value for QuantityEnd to be an explicit nil

### UnsetQuantityEnd
`func (o *PricingBreak) UnsetQuantityEnd()`

UnsetQuantityEnd ensures that no value is present for QuantityEnd, not even an explicit nil
### GetUnlimited

`func (o *PricingBreak) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *PricingBreak) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *PricingBreak) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.

### HasUnlimited

`func (o *PricingBreak) HasUnlimited() bool`

HasUnlimited returns a boolean if a field has been set.

### GetPriceMethod

`func (o *PricingBreak) GetPriceMethod() string`

GetPriceMethod returns the PriceMethod field if non-nil, zero value otherwise.

### GetPriceMethodOk

`func (o *PricingBreak) GetPriceMethodOk() (*string, bool)`

GetPriceMethodOk returns a tuple with the PriceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMethod

`func (o *PricingBreak) SetPriceMethod(v string)`

SetPriceMethod sets PriceMethod field to given value.


### SetPriceMethodNil

`func (o *PricingBreak) SetPriceMethodNil(b bool)`

 SetPriceMethodNil sets the value for PriceMethod to be an explicit nil

### UnsetPriceMethod
`func (o *PricingBreak) UnsetPriceMethod()`

UnsetPriceMethod ensures that no value is present for PriceMethod, not even an explicit nil
### GetInfo

`func (o *PricingBreak) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PricingBreak) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PricingBreak) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PricingBreak) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


