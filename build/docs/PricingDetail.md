# PricingDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Product** | Pointer to [**CatalogItemReference**](CatalogItemReference.md) |  | [optional] 
**Category** | Pointer to [**ProductCategoryReference**](ProductCategoryReference.md) |  | [optional] 
**SubCategory** | Pointer to [**ProductSubCategoryReference**](ProductSubCategoryReference.md) |  | [optional] 
**StartDate** | **time.Time** |  | 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**NoEndDate** | Pointer to **bool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPricingDetail

`func NewPricingDetail(startDate time.Time, ) *PricingDetail`

NewPricingDetail instantiates a new PricingDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingDetailWithDefaults

`func NewPricingDetailWithDefaults() *PricingDetail`

NewPricingDetailWithDefaults instantiates a new PricingDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PricingDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PricingDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProduct

`func (o *PricingDetail) GetProduct() CatalogItemReference`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *PricingDetail) GetProductOk() (*CatalogItemReference, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *PricingDetail) SetProduct(v CatalogItemReference)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *PricingDetail) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCategory

`func (o *PricingDetail) GetCategory() ProductCategoryReference`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PricingDetail) GetCategoryOk() (*ProductCategoryReference, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PricingDetail) SetCategory(v ProductCategoryReference)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PricingDetail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetSubCategory

`func (o *PricingDetail) GetSubCategory() ProductSubCategoryReference`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *PricingDetail) GetSubCategoryOk() (*ProductSubCategoryReference, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *PricingDetail) SetSubCategory(v ProductSubCategoryReference)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *PricingDetail) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### GetStartDate

`func (o *PricingDetail) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *PricingDetail) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *PricingDetail) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *PricingDetail) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *PricingDetail) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *PricingDetail) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *PricingDetail) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetNoEndDate

`func (o *PricingDetail) GetNoEndDate() bool`

GetNoEndDate returns the NoEndDate field if non-nil, zero value otherwise.

### GetNoEndDateOk

`func (o *PricingDetail) GetNoEndDateOk() (*bool, bool)`

GetNoEndDateOk returns a tuple with the NoEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoEndDate

`func (o *PricingDetail) SetNoEndDate(v bool)`

SetNoEndDate sets NoEndDate field to given value.

### HasNoEndDate

`func (o *PricingDetail) HasNoEndDate() bool`

HasNoEndDate returns a boolean if a field has been set.

### GetInfo

`func (o *PricingDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PricingDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PricingDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PricingDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


