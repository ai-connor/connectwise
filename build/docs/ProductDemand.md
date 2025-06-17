# ProductDemand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductRecId** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewProductDemand

`func NewProductDemand() *ProductDemand`

NewProductDemand instantiates a new ProductDemand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductDemandWithDefaults

`func NewProductDemandWithDefaults() *ProductDemand`

NewProductDemandWithDefaults instantiates a new ProductDemand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductRecId

`func (o *ProductDemand) GetProductRecId() int32`

GetProductRecId returns the ProductRecId field if non-nil, zero value otherwise.

### GetProductRecIdOk

`func (o *ProductDemand) GetProductRecIdOk() (*int32, bool)`

GetProductRecIdOk returns a tuple with the ProductRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductRecId

`func (o *ProductDemand) SetProductRecId(v int32)`

SetProductRecId sets ProductRecId field to given value.

### HasProductRecId

`func (o *ProductDemand) HasProductRecId() bool`

HasProductRecId returns a boolean if a field has been set.

### GetQuantity

`func (o *ProductDemand) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductDemand) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductDemand) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductDemand) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCost

`func (o *ProductDemand) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductDemand) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductDemand) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProductDemand) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProductDemand) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProductDemand) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


