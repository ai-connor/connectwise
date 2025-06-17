# ProductTypeExemption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ProductType** | [**ProductTypeReference**](ProductTypeReference.md) |  | 
**TaxableLevels** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProductTypeExemption

`func NewProductTypeExemption(productType ProductTypeReference, ) *ProductTypeExemption`

NewProductTypeExemption instantiates a new ProductTypeExemption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductTypeExemptionWithDefaults

`func NewProductTypeExemptionWithDefaults() *ProductTypeExemption`

NewProductTypeExemptionWithDefaults instantiates a new ProductTypeExemption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductTypeExemption) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductTypeExemption) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductTypeExemption) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductTypeExemption) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProductType

`func (o *ProductTypeExemption) GetProductType() ProductTypeReference`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *ProductTypeExemption) GetProductTypeOk() (*ProductTypeReference, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *ProductTypeExemption) SetProductType(v ProductTypeReference)`

SetProductType sets ProductType field to given value.


### GetTaxableLevels

`func (o *ProductTypeExemption) GetTaxableLevels() []int32`

GetTaxableLevels returns the TaxableLevels field if non-nil, zero value otherwise.

### GetTaxableLevelsOk

`func (o *ProductTypeExemption) GetTaxableLevelsOk() (*[]int32, bool)`

GetTaxableLevelsOk returns a tuple with the TaxableLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableLevels

`func (o *ProductTypeExemption) SetTaxableLevels(v []int32)`

SetTaxableLevels sets TaxableLevels field to given value.

### HasTaxableLevels

`func (o *ProductTypeExemption) HasTaxableLevels() bool`

HasTaxableLevels returns a boolean if a field has been set.

### GetInfo

`func (o *ProductTypeExemption) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProductTypeExemption) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProductTypeExemption) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProductTypeExemption) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


