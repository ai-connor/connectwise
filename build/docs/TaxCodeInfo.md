# TaxCodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **string** |  | [optional] 
**CancelDate** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOneRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelTwoRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelThreeRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFourRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelFiveRate** | Pointer to **NullableFloat64** |  | [optional] 
**LevelSixRate** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxCodeInfo

`func NewTaxCodeInfo() *TaxCodeInfo`

NewTaxCodeInfo instantiates a new TaxCodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeInfoWithDefaults

`func NewTaxCodeInfoWithDefaults() *TaxCodeInfo`

NewTaxCodeInfoWithDefaults instantiates a new TaxCodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxCodeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCodeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCodeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCodeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *TaxCodeInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *TaxCodeInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *TaxCodeInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *TaxCodeInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *TaxCodeInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxCodeInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxCodeInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TaxCodeInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *TaxCodeInfo) GetEffectiveDate() string`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *TaxCodeInfo) GetEffectiveDateOk() (*string, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *TaxCodeInfo) SetEffectiveDate(v string)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *TaxCodeInfo) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCancelDate

`func (o *TaxCodeInfo) GetCancelDate() string`

GetCancelDate returns the CancelDate field if non-nil, zero value otherwise.

### GetCancelDateOk

`func (o *TaxCodeInfo) GetCancelDateOk() (*string, bool)`

GetCancelDateOk returns a tuple with the CancelDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelDate

`func (o *TaxCodeInfo) SetCancelDate(v string)`

SetCancelDate sets CancelDate field to given value.

### HasCancelDate

`func (o *TaxCodeInfo) HasCancelDate() bool`

HasCancelDate returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *TaxCodeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *TaxCodeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *TaxCodeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *TaxCodeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *TaxCodeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *TaxCodeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetLevelOneRate

`func (o *TaxCodeInfo) GetLevelOneRate() float64`

GetLevelOneRate returns the LevelOneRate field if non-nil, zero value otherwise.

### GetLevelOneRateOk

`func (o *TaxCodeInfo) GetLevelOneRateOk() (*float64, bool)`

GetLevelOneRateOk returns a tuple with the LevelOneRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOneRate

`func (o *TaxCodeInfo) SetLevelOneRate(v float64)`

SetLevelOneRate sets LevelOneRate field to given value.

### HasLevelOneRate

`func (o *TaxCodeInfo) HasLevelOneRate() bool`

HasLevelOneRate returns a boolean if a field has been set.

### SetLevelOneRateNil

`func (o *TaxCodeInfo) SetLevelOneRateNil(b bool)`

 SetLevelOneRateNil sets the value for LevelOneRate to be an explicit nil

### UnsetLevelOneRate
`func (o *TaxCodeInfo) UnsetLevelOneRate()`

UnsetLevelOneRate ensures that no value is present for LevelOneRate, not even an explicit nil
### GetLevelTwoRate

`func (o *TaxCodeInfo) GetLevelTwoRate() float64`

GetLevelTwoRate returns the LevelTwoRate field if non-nil, zero value otherwise.

### GetLevelTwoRateOk

`func (o *TaxCodeInfo) GetLevelTwoRateOk() (*float64, bool)`

GetLevelTwoRateOk returns a tuple with the LevelTwoRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwoRate

`func (o *TaxCodeInfo) SetLevelTwoRate(v float64)`

SetLevelTwoRate sets LevelTwoRate field to given value.

### HasLevelTwoRate

`func (o *TaxCodeInfo) HasLevelTwoRate() bool`

HasLevelTwoRate returns a boolean if a field has been set.

### SetLevelTwoRateNil

`func (o *TaxCodeInfo) SetLevelTwoRateNil(b bool)`

 SetLevelTwoRateNil sets the value for LevelTwoRate to be an explicit nil

### UnsetLevelTwoRate
`func (o *TaxCodeInfo) UnsetLevelTwoRate()`

UnsetLevelTwoRate ensures that no value is present for LevelTwoRate, not even an explicit nil
### GetLevelThreeRate

`func (o *TaxCodeInfo) GetLevelThreeRate() float64`

GetLevelThreeRate returns the LevelThreeRate field if non-nil, zero value otherwise.

### GetLevelThreeRateOk

`func (o *TaxCodeInfo) GetLevelThreeRateOk() (*float64, bool)`

GetLevelThreeRateOk returns a tuple with the LevelThreeRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThreeRate

`func (o *TaxCodeInfo) SetLevelThreeRate(v float64)`

SetLevelThreeRate sets LevelThreeRate field to given value.

### HasLevelThreeRate

`func (o *TaxCodeInfo) HasLevelThreeRate() bool`

HasLevelThreeRate returns a boolean if a field has been set.

### SetLevelThreeRateNil

`func (o *TaxCodeInfo) SetLevelThreeRateNil(b bool)`

 SetLevelThreeRateNil sets the value for LevelThreeRate to be an explicit nil

### UnsetLevelThreeRate
`func (o *TaxCodeInfo) UnsetLevelThreeRate()`

UnsetLevelThreeRate ensures that no value is present for LevelThreeRate, not even an explicit nil
### GetLevelFourRate

`func (o *TaxCodeInfo) GetLevelFourRate() float64`

GetLevelFourRate returns the LevelFourRate field if non-nil, zero value otherwise.

### GetLevelFourRateOk

`func (o *TaxCodeInfo) GetLevelFourRateOk() (*float64, bool)`

GetLevelFourRateOk returns a tuple with the LevelFourRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFourRate

`func (o *TaxCodeInfo) SetLevelFourRate(v float64)`

SetLevelFourRate sets LevelFourRate field to given value.

### HasLevelFourRate

`func (o *TaxCodeInfo) HasLevelFourRate() bool`

HasLevelFourRate returns a boolean if a field has been set.

### SetLevelFourRateNil

`func (o *TaxCodeInfo) SetLevelFourRateNil(b bool)`

 SetLevelFourRateNil sets the value for LevelFourRate to be an explicit nil

### UnsetLevelFourRate
`func (o *TaxCodeInfo) UnsetLevelFourRate()`

UnsetLevelFourRate ensures that no value is present for LevelFourRate, not even an explicit nil
### GetLevelFiveRate

`func (o *TaxCodeInfo) GetLevelFiveRate() float64`

GetLevelFiveRate returns the LevelFiveRate field if non-nil, zero value otherwise.

### GetLevelFiveRateOk

`func (o *TaxCodeInfo) GetLevelFiveRateOk() (*float64, bool)`

GetLevelFiveRateOk returns a tuple with the LevelFiveRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFiveRate

`func (o *TaxCodeInfo) SetLevelFiveRate(v float64)`

SetLevelFiveRate sets LevelFiveRate field to given value.

### HasLevelFiveRate

`func (o *TaxCodeInfo) HasLevelFiveRate() bool`

HasLevelFiveRate returns a boolean if a field has been set.

### SetLevelFiveRateNil

`func (o *TaxCodeInfo) SetLevelFiveRateNil(b bool)`

 SetLevelFiveRateNil sets the value for LevelFiveRate to be an explicit nil

### UnsetLevelFiveRate
`func (o *TaxCodeInfo) UnsetLevelFiveRate()`

UnsetLevelFiveRate ensures that no value is present for LevelFiveRate, not even an explicit nil
### GetLevelSixRate

`func (o *TaxCodeInfo) GetLevelSixRate() float64`

GetLevelSixRate returns the LevelSixRate field if non-nil, zero value otherwise.

### GetLevelSixRateOk

`func (o *TaxCodeInfo) GetLevelSixRateOk() (*float64, bool)`

GetLevelSixRateOk returns a tuple with the LevelSixRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSixRate

`func (o *TaxCodeInfo) SetLevelSixRate(v float64)`

SetLevelSixRate sets LevelSixRate field to given value.

### HasLevelSixRate

`func (o *TaxCodeInfo) HasLevelSixRate() bool`

HasLevelSixRate returns a boolean if a field has been set.

### SetLevelSixRateNil

`func (o *TaxCodeInfo) SetLevelSixRateNil(b bool)`

 SetLevelSixRateNil sets the value for LevelSixRate to be an explicit nil

### UnsetLevelSixRate
`func (o *TaxCodeInfo) UnsetLevelSixRate()`

UnsetLevelSixRate ensures that no value is present for LevelSixRate, not even an explicit nil
### GetInfo

`func (o *TaxCodeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxCodeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxCodeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxCodeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


