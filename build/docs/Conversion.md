# Conversion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Quantity** | Pointer to **NullableFloat64** |  | [optional] 
**UomType** | [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | 
**ParentUOM** | Pointer to [**UnitOfMeasureReference**](UnitOfMeasureReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConversion

`func NewConversion(uomType UnitOfMeasureReference, ) *Conversion`

NewConversion instantiates a new Conversion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionWithDefaults

`func NewConversionWithDefaults() *Conversion`

NewConversionWithDefaults instantiates a new Conversion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Conversion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conversion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conversion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Conversion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetQuantity

`func (o *Conversion) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Conversion) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Conversion) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Conversion) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *Conversion) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *Conversion) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetUomType

`func (o *Conversion) GetUomType() UnitOfMeasureReference`

GetUomType returns the UomType field if non-nil, zero value otherwise.

### GetUomTypeOk

`func (o *Conversion) GetUomTypeOk() (*UnitOfMeasureReference, bool)`

GetUomTypeOk returns a tuple with the UomType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUomType

`func (o *Conversion) SetUomType(v UnitOfMeasureReference)`

SetUomType sets UomType field to given value.


### GetParentUOM

`func (o *Conversion) GetParentUOM() UnitOfMeasureReference`

GetParentUOM returns the ParentUOM field if non-nil, zero value otherwise.

### GetParentUOMOk

`func (o *Conversion) GetParentUOMOk() (*UnitOfMeasureReference, bool)`

GetParentUOMOk returns a tuple with the ParentUOM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentUOM

`func (o *Conversion) SetParentUOM(v UnitOfMeasureReference)`

SetParentUOM sets ParentUOM field to given value.

### HasParentUOM

`func (o *Conversion) HasParentUOM() bool`

HasParentUOM returns a boolean if a field has been set.

### GetInfo

`func (o *Conversion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Conversion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Conversion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Conversion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


