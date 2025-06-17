# CatalogPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItem** | Pointer to [**CatalogItemReference**](CatalogItemReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 

## Methods

### NewCatalogPricing

`func NewCatalogPricing() *CatalogPricing`

NewCatalogPricing instantiates a new CatalogPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogPricingWithDefaults

`func NewCatalogPricingWithDefaults() *CatalogPricing`

NewCatalogPricingWithDefaults instantiates a new CatalogPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogItem

`func (o *CatalogPricing) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *CatalogPricing) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *CatalogPricing) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.

### HasCatalogItem

`func (o *CatalogPricing) HasCatalogItem() bool`

HasCatalogItem returns a boolean if a field has been set.

### GetCompany

`func (o *CatalogPricing) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CatalogPricing) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CatalogPricing) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CatalogPricing) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *CatalogPricing) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CatalogPricing) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CatalogPricing) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CatalogPricing) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetQuantity

`func (o *CatalogPricing) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogPricing) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogPricing) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CatalogPricing) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetDate

`func (o *CatalogPricing) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CatalogPricing) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CatalogPricing) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *CatalogPricing) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPrice

`func (o *CatalogPricing) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogPricing) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogPricing) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogPricing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogPricing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogPricing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


