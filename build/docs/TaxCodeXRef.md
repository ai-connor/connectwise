# TaxCodeXRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**LevelOne** | Pointer to **NullableString** |  | [optional] 
**LevelTwo** | Pointer to **NullableString** |  | [optional] 
**LevelThree** | Pointer to **NullableString** |  | [optional] 
**LevelFour** | Pointer to **NullableString** |  | [optional] 
**LevelFive** | Pointer to **NullableString** |  | [optional] 
**LevelSix** | Pointer to **NullableString** |  | [optional] 
**TaxCode** | Pointer to [**TaxCodeReference**](TaxCodeReference.md) |  | [optional] 
**TaxableLevels** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTaxCodeXRef

`func NewTaxCodeXRef(description string, ) *TaxCodeXRef`

NewTaxCodeXRef instantiates a new TaxCodeXRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxCodeXRefWithDefaults

`func NewTaxCodeXRefWithDefaults() *TaxCodeXRef`

NewTaxCodeXRefWithDefaults instantiates a new TaxCodeXRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TaxCodeXRef) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaxCodeXRef) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaxCodeXRef) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TaxCodeXRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *TaxCodeXRef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TaxCodeXRef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TaxCodeXRef) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDefaultFlag

`func (o *TaxCodeXRef) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *TaxCodeXRef) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *TaxCodeXRef) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *TaxCodeXRef) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *TaxCodeXRef) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *TaxCodeXRef) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetLevelOne

`func (o *TaxCodeXRef) GetLevelOne() string`

GetLevelOne returns the LevelOne field if non-nil, zero value otherwise.

### GetLevelOneOk

`func (o *TaxCodeXRef) GetLevelOneOk() (*string, bool)`

GetLevelOneOk returns a tuple with the LevelOne field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOne

`func (o *TaxCodeXRef) SetLevelOne(v string)`

SetLevelOne sets LevelOne field to given value.

### HasLevelOne

`func (o *TaxCodeXRef) HasLevelOne() bool`

HasLevelOne returns a boolean if a field has been set.

### SetLevelOneNil

`func (o *TaxCodeXRef) SetLevelOneNil(b bool)`

 SetLevelOneNil sets the value for LevelOne to be an explicit nil

### UnsetLevelOne
`func (o *TaxCodeXRef) UnsetLevelOne()`

UnsetLevelOne ensures that no value is present for LevelOne, not even an explicit nil
### GetLevelTwo

`func (o *TaxCodeXRef) GetLevelTwo() string`

GetLevelTwo returns the LevelTwo field if non-nil, zero value otherwise.

### GetLevelTwoOk

`func (o *TaxCodeXRef) GetLevelTwoOk() (*string, bool)`

GetLevelTwoOk returns a tuple with the LevelTwo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTwo

`func (o *TaxCodeXRef) SetLevelTwo(v string)`

SetLevelTwo sets LevelTwo field to given value.

### HasLevelTwo

`func (o *TaxCodeXRef) HasLevelTwo() bool`

HasLevelTwo returns a boolean if a field has been set.

### SetLevelTwoNil

`func (o *TaxCodeXRef) SetLevelTwoNil(b bool)`

 SetLevelTwoNil sets the value for LevelTwo to be an explicit nil

### UnsetLevelTwo
`func (o *TaxCodeXRef) UnsetLevelTwo()`

UnsetLevelTwo ensures that no value is present for LevelTwo, not even an explicit nil
### GetLevelThree

`func (o *TaxCodeXRef) GetLevelThree() string`

GetLevelThree returns the LevelThree field if non-nil, zero value otherwise.

### GetLevelThreeOk

`func (o *TaxCodeXRef) GetLevelThreeOk() (*string, bool)`

GetLevelThreeOk returns a tuple with the LevelThree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelThree

`func (o *TaxCodeXRef) SetLevelThree(v string)`

SetLevelThree sets LevelThree field to given value.

### HasLevelThree

`func (o *TaxCodeXRef) HasLevelThree() bool`

HasLevelThree returns a boolean if a field has been set.

### SetLevelThreeNil

`func (o *TaxCodeXRef) SetLevelThreeNil(b bool)`

 SetLevelThreeNil sets the value for LevelThree to be an explicit nil

### UnsetLevelThree
`func (o *TaxCodeXRef) UnsetLevelThree()`

UnsetLevelThree ensures that no value is present for LevelThree, not even an explicit nil
### GetLevelFour

`func (o *TaxCodeXRef) GetLevelFour() string`

GetLevelFour returns the LevelFour field if non-nil, zero value otherwise.

### GetLevelFourOk

`func (o *TaxCodeXRef) GetLevelFourOk() (*string, bool)`

GetLevelFourOk returns a tuple with the LevelFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFour

`func (o *TaxCodeXRef) SetLevelFour(v string)`

SetLevelFour sets LevelFour field to given value.

### HasLevelFour

`func (o *TaxCodeXRef) HasLevelFour() bool`

HasLevelFour returns a boolean if a field has been set.

### SetLevelFourNil

`func (o *TaxCodeXRef) SetLevelFourNil(b bool)`

 SetLevelFourNil sets the value for LevelFour to be an explicit nil

### UnsetLevelFour
`func (o *TaxCodeXRef) UnsetLevelFour()`

UnsetLevelFour ensures that no value is present for LevelFour, not even an explicit nil
### GetLevelFive

`func (o *TaxCodeXRef) GetLevelFive() string`

GetLevelFive returns the LevelFive field if non-nil, zero value otherwise.

### GetLevelFiveOk

`func (o *TaxCodeXRef) GetLevelFiveOk() (*string, bool)`

GetLevelFiveOk returns a tuple with the LevelFive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelFive

`func (o *TaxCodeXRef) SetLevelFive(v string)`

SetLevelFive sets LevelFive field to given value.

### HasLevelFive

`func (o *TaxCodeXRef) HasLevelFive() bool`

HasLevelFive returns a boolean if a field has been set.

### SetLevelFiveNil

`func (o *TaxCodeXRef) SetLevelFiveNil(b bool)`

 SetLevelFiveNil sets the value for LevelFive to be an explicit nil

### UnsetLevelFive
`func (o *TaxCodeXRef) UnsetLevelFive()`

UnsetLevelFive ensures that no value is present for LevelFive, not even an explicit nil
### GetLevelSix

`func (o *TaxCodeXRef) GetLevelSix() string`

GetLevelSix returns the LevelSix field if non-nil, zero value otherwise.

### GetLevelSixOk

`func (o *TaxCodeXRef) GetLevelSixOk() (*string, bool)`

GetLevelSixOk returns a tuple with the LevelSix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelSix

`func (o *TaxCodeXRef) SetLevelSix(v string)`

SetLevelSix sets LevelSix field to given value.

### HasLevelSix

`func (o *TaxCodeXRef) HasLevelSix() bool`

HasLevelSix returns a boolean if a field has been set.

### SetLevelSixNil

`func (o *TaxCodeXRef) SetLevelSixNil(b bool)`

 SetLevelSixNil sets the value for LevelSix to be an explicit nil

### UnsetLevelSix
`func (o *TaxCodeXRef) UnsetLevelSix()`

UnsetLevelSix ensures that no value is present for LevelSix, not even an explicit nil
### GetTaxCode

`func (o *TaxCodeXRef) GetTaxCode() TaxCodeReference`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *TaxCodeXRef) GetTaxCodeOk() (*TaxCodeReference, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *TaxCodeXRef) SetTaxCode(v TaxCodeReference)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *TaxCodeXRef) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### GetTaxableLevels

`func (o *TaxCodeXRef) GetTaxableLevels() []int32`

GetTaxableLevels returns the TaxableLevels field if non-nil, zero value otherwise.

### GetTaxableLevelsOk

`func (o *TaxCodeXRef) GetTaxableLevelsOk() (*[]int32, bool)`

GetTaxableLevelsOk returns a tuple with the TaxableLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableLevels

`func (o *TaxCodeXRef) SetTaxableLevels(v []int32)`

SetTaxableLevels sets TaxableLevels field to given value.

### HasTaxableLevels

`func (o *TaxCodeXRef) HasTaxableLevels() bool`

HasTaxableLevels returns a boolean if a field has been set.

### GetInfo

`func (o *TaxCodeXRef) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TaxCodeXRef) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TaxCodeXRef) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TaxCodeXRef) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


