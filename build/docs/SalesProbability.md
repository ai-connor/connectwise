# SalesProbability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Probability** | **int32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesProbability

`func NewSalesProbability(probability int32, ) *SalesProbability`

NewSalesProbability instantiates a new SalesProbability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesProbabilityWithDefaults

`func NewSalesProbabilityWithDefaults() *SalesProbability`

NewSalesProbabilityWithDefaults instantiates a new SalesProbability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesProbability) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesProbability) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesProbability) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesProbability) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProbability

`func (o *SalesProbability) GetProbability() int32`

GetProbability returns the Probability field if non-nil, zero value otherwise.

### GetProbabilityOk

`func (o *SalesProbability) GetProbabilityOk() (*int32, bool)`

GetProbabilityOk returns a tuple with the Probability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbability

`func (o *SalesProbability) SetProbability(v int32)`

SetProbability sets Probability field to given value.


### GetInfo

`func (o *SalesProbability) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesProbability) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesProbability) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesProbability) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


