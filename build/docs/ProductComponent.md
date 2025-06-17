# ProductComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  Required On Updates; | [optional] 
**Quantity** | **NullableFloat64** |  | 
**CatalogItem** | [**CatalogItemReference**](CatalogItemReference.md) |  | 
**HidePriceFlag** | Pointer to **NullableBool** |  | [optional] 
**HideItemIdentifierFlag** | Pointer to **NullableBool** |  | [optional] 
**HideDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**HideQuantityFlag** | Pointer to **NullableBool** |  | [optional] 
**HideExtendedPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**ParentProductItem** | Pointer to [**ProductItemReference**](ProductItemReference.md) |  | [optional] 
**ProductItem** | Pointer to [**ProductItemReference**](ProductItemReference.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProductComponent

`func NewProductComponent(quantity NullableFloat64, catalogItem CatalogItemReference, ) *ProductComponent`

NewProductComponent instantiates a new ProductComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductComponentWithDefaults

`func NewProductComponentWithDefaults() *ProductComponent`

NewProductComponentWithDefaults instantiates a new ProductComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ProductComponent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ProductComponent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ProductComponent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *ProductComponent) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *ProductComponent) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ProductComponent) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuantity

`func (o *ProductComponent) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductComponent) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductComponent) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *ProductComponent) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ProductComponent) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetCatalogItem

`func (o *ProductComponent) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *ProductComponent) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *ProductComponent) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.


### GetHidePriceFlag

`func (o *ProductComponent) GetHidePriceFlag() bool`

GetHidePriceFlag returns the HidePriceFlag field if non-nil, zero value otherwise.

### GetHidePriceFlagOk

`func (o *ProductComponent) GetHidePriceFlagOk() (*bool, bool)`

GetHidePriceFlagOk returns a tuple with the HidePriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePriceFlag

`func (o *ProductComponent) SetHidePriceFlag(v bool)`

SetHidePriceFlag sets HidePriceFlag field to given value.

### HasHidePriceFlag

`func (o *ProductComponent) HasHidePriceFlag() bool`

HasHidePriceFlag returns a boolean if a field has been set.

### SetHidePriceFlagNil

`func (o *ProductComponent) SetHidePriceFlagNil(b bool)`

 SetHidePriceFlagNil sets the value for HidePriceFlag to be an explicit nil

### UnsetHidePriceFlag
`func (o *ProductComponent) UnsetHidePriceFlag()`

UnsetHidePriceFlag ensures that no value is present for HidePriceFlag, not even an explicit nil
### GetHideItemIdentifierFlag

`func (o *ProductComponent) GetHideItemIdentifierFlag() bool`

GetHideItemIdentifierFlag returns the HideItemIdentifierFlag field if non-nil, zero value otherwise.

### GetHideItemIdentifierFlagOk

`func (o *ProductComponent) GetHideItemIdentifierFlagOk() (*bool, bool)`

GetHideItemIdentifierFlagOk returns a tuple with the HideItemIdentifierFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideItemIdentifierFlag

`func (o *ProductComponent) SetHideItemIdentifierFlag(v bool)`

SetHideItemIdentifierFlag sets HideItemIdentifierFlag field to given value.

### HasHideItemIdentifierFlag

`func (o *ProductComponent) HasHideItemIdentifierFlag() bool`

HasHideItemIdentifierFlag returns a boolean if a field has been set.

### SetHideItemIdentifierFlagNil

`func (o *ProductComponent) SetHideItemIdentifierFlagNil(b bool)`

 SetHideItemIdentifierFlagNil sets the value for HideItemIdentifierFlag to be an explicit nil

### UnsetHideItemIdentifierFlag
`func (o *ProductComponent) UnsetHideItemIdentifierFlag()`

UnsetHideItemIdentifierFlag ensures that no value is present for HideItemIdentifierFlag, not even an explicit nil
### GetHideDescriptionFlag

`func (o *ProductComponent) GetHideDescriptionFlag() bool`

GetHideDescriptionFlag returns the HideDescriptionFlag field if non-nil, zero value otherwise.

### GetHideDescriptionFlagOk

`func (o *ProductComponent) GetHideDescriptionFlagOk() (*bool, bool)`

GetHideDescriptionFlagOk returns a tuple with the HideDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDescriptionFlag

`func (o *ProductComponent) SetHideDescriptionFlag(v bool)`

SetHideDescriptionFlag sets HideDescriptionFlag field to given value.

### HasHideDescriptionFlag

`func (o *ProductComponent) HasHideDescriptionFlag() bool`

HasHideDescriptionFlag returns a boolean if a field has been set.

### SetHideDescriptionFlagNil

`func (o *ProductComponent) SetHideDescriptionFlagNil(b bool)`

 SetHideDescriptionFlagNil sets the value for HideDescriptionFlag to be an explicit nil

### UnsetHideDescriptionFlag
`func (o *ProductComponent) UnsetHideDescriptionFlag()`

UnsetHideDescriptionFlag ensures that no value is present for HideDescriptionFlag, not even an explicit nil
### GetHideQuantityFlag

`func (o *ProductComponent) GetHideQuantityFlag() bool`

GetHideQuantityFlag returns the HideQuantityFlag field if non-nil, zero value otherwise.

### GetHideQuantityFlagOk

`func (o *ProductComponent) GetHideQuantityFlagOk() (*bool, bool)`

GetHideQuantityFlagOk returns a tuple with the HideQuantityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideQuantityFlag

`func (o *ProductComponent) SetHideQuantityFlag(v bool)`

SetHideQuantityFlag sets HideQuantityFlag field to given value.

### HasHideQuantityFlag

`func (o *ProductComponent) HasHideQuantityFlag() bool`

HasHideQuantityFlag returns a boolean if a field has been set.

### SetHideQuantityFlagNil

`func (o *ProductComponent) SetHideQuantityFlagNil(b bool)`

 SetHideQuantityFlagNil sets the value for HideQuantityFlag to be an explicit nil

### UnsetHideQuantityFlag
`func (o *ProductComponent) UnsetHideQuantityFlag()`

UnsetHideQuantityFlag ensures that no value is present for HideQuantityFlag, not even an explicit nil
### GetHideExtendedPriceFlag

`func (o *ProductComponent) GetHideExtendedPriceFlag() bool`

GetHideExtendedPriceFlag returns the HideExtendedPriceFlag field if non-nil, zero value otherwise.

### GetHideExtendedPriceFlagOk

`func (o *ProductComponent) GetHideExtendedPriceFlagOk() (*bool, bool)`

GetHideExtendedPriceFlagOk returns a tuple with the HideExtendedPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideExtendedPriceFlag

`func (o *ProductComponent) SetHideExtendedPriceFlag(v bool)`

SetHideExtendedPriceFlag sets HideExtendedPriceFlag field to given value.

### HasHideExtendedPriceFlag

`func (o *ProductComponent) HasHideExtendedPriceFlag() bool`

HasHideExtendedPriceFlag returns a boolean if a field has been set.

### SetHideExtendedPriceFlagNil

`func (o *ProductComponent) SetHideExtendedPriceFlagNil(b bool)`

 SetHideExtendedPriceFlagNil sets the value for HideExtendedPriceFlag to be an explicit nil

### UnsetHideExtendedPriceFlag
`func (o *ProductComponent) UnsetHideExtendedPriceFlag()`

UnsetHideExtendedPriceFlag ensures that no value is present for HideExtendedPriceFlag, not even an explicit nil
### GetVendor

`func (o *ProductComponent) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ProductComponent) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ProductComponent) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ProductComponent) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetParentProductItem

`func (o *ProductComponent) GetParentProductItem() ProductItemReference`

GetParentProductItem returns the ParentProductItem field if non-nil, zero value otherwise.

### GetParentProductItemOk

`func (o *ProductComponent) GetParentProductItemOk() (*ProductItemReference, bool)`

GetParentProductItemOk returns a tuple with the ParentProductItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentProductItem

`func (o *ProductComponent) SetParentProductItem(v ProductItemReference)`

SetParentProductItem sets ParentProductItem field to given value.

### HasParentProductItem

`func (o *ProductComponent) HasParentProductItem() bool`

HasParentProductItem returns a boolean if a field has been set.

### GetProductItem

`func (o *ProductComponent) GetProductItem() ProductItemReference`

GetProductItem returns the ProductItem field if non-nil, zero value otherwise.

### GetProductItemOk

`func (o *ProductComponent) GetProductItemOk() (*ProductItemReference, bool)`

GetProductItemOk returns a tuple with the ProductItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItem

`func (o *ProductComponent) SetProductItem(v ProductItemReference)`

SetProductItem sets ProductItem field to given value.

### HasProductItem

`func (o *ProductComponent) HasProductItem() bool`

HasProductItem returns a boolean if a field has been set.

### GetPrice

`func (o *ProductComponent) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductComponent) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductComponent) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductComponent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ProductComponent) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ProductComponent) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCost

`func (o *ProductComponent) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProductComponent) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProductComponent) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProductComponent) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *ProductComponent) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *ProductComponent) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetInfo

`func (o *ProductComponent) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProductComponent) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProductComponent) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProductComponent) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


