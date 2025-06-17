# CatalogComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | Pointer to **NullableInt32** |  | [optional] 
**Quantity** | **NullableFloat64** |  | 
**CatalogItem** | [**CatalogItemReference**](CatalogItemReference.md) |  | 
**HidePriceFlag** | Pointer to **NullableBool** |  | [optional] 
**HideItemIdentifierFlag** | Pointer to **NullableBool** |  | [optional] 
**HideDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**HideQuantityFlag** | Pointer to **NullableBool** |  | [optional] 
**HideExtendedPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**ParentCatalogItem** | Pointer to [**CatalogItemReference**](CatalogItemReference.md) |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCatalogComponent

`func NewCatalogComponent(quantity NullableFloat64, catalogItem CatalogItemReference, ) *CatalogComponent`

NewCatalogComponent instantiates a new CatalogComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogComponentWithDefaults

`func NewCatalogComponentWithDefaults() *CatalogComponent`

NewCatalogComponentWithDefaults instantiates a new CatalogComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogComponent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogComponent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogComponent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogComponent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *CatalogComponent) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *CatalogComponent) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *CatalogComponent) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *CatalogComponent) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### SetSequenceNumberNil

`func (o *CatalogComponent) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *CatalogComponent) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetQuantity

`func (o *CatalogComponent) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogComponent) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogComponent) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.


### SetQuantityNil

`func (o *CatalogComponent) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CatalogComponent) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetCatalogItem

`func (o *CatalogComponent) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *CatalogComponent) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *CatalogComponent) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.


### GetHidePriceFlag

`func (o *CatalogComponent) GetHidePriceFlag() bool`

GetHidePriceFlag returns the HidePriceFlag field if non-nil, zero value otherwise.

### GetHidePriceFlagOk

`func (o *CatalogComponent) GetHidePriceFlagOk() (*bool, bool)`

GetHidePriceFlagOk returns a tuple with the HidePriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidePriceFlag

`func (o *CatalogComponent) SetHidePriceFlag(v bool)`

SetHidePriceFlag sets HidePriceFlag field to given value.

### HasHidePriceFlag

`func (o *CatalogComponent) HasHidePriceFlag() bool`

HasHidePriceFlag returns a boolean if a field has been set.

### SetHidePriceFlagNil

`func (o *CatalogComponent) SetHidePriceFlagNil(b bool)`

 SetHidePriceFlagNil sets the value for HidePriceFlag to be an explicit nil

### UnsetHidePriceFlag
`func (o *CatalogComponent) UnsetHidePriceFlag()`

UnsetHidePriceFlag ensures that no value is present for HidePriceFlag, not even an explicit nil
### GetHideItemIdentifierFlag

`func (o *CatalogComponent) GetHideItemIdentifierFlag() bool`

GetHideItemIdentifierFlag returns the HideItemIdentifierFlag field if non-nil, zero value otherwise.

### GetHideItemIdentifierFlagOk

`func (o *CatalogComponent) GetHideItemIdentifierFlagOk() (*bool, bool)`

GetHideItemIdentifierFlagOk returns a tuple with the HideItemIdentifierFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideItemIdentifierFlag

`func (o *CatalogComponent) SetHideItemIdentifierFlag(v bool)`

SetHideItemIdentifierFlag sets HideItemIdentifierFlag field to given value.

### HasHideItemIdentifierFlag

`func (o *CatalogComponent) HasHideItemIdentifierFlag() bool`

HasHideItemIdentifierFlag returns a boolean if a field has been set.

### SetHideItemIdentifierFlagNil

`func (o *CatalogComponent) SetHideItemIdentifierFlagNil(b bool)`

 SetHideItemIdentifierFlagNil sets the value for HideItemIdentifierFlag to be an explicit nil

### UnsetHideItemIdentifierFlag
`func (o *CatalogComponent) UnsetHideItemIdentifierFlag()`

UnsetHideItemIdentifierFlag ensures that no value is present for HideItemIdentifierFlag, not even an explicit nil
### GetHideDescriptionFlag

`func (o *CatalogComponent) GetHideDescriptionFlag() bool`

GetHideDescriptionFlag returns the HideDescriptionFlag field if non-nil, zero value otherwise.

### GetHideDescriptionFlagOk

`func (o *CatalogComponent) GetHideDescriptionFlagOk() (*bool, bool)`

GetHideDescriptionFlagOk returns a tuple with the HideDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDescriptionFlag

`func (o *CatalogComponent) SetHideDescriptionFlag(v bool)`

SetHideDescriptionFlag sets HideDescriptionFlag field to given value.

### HasHideDescriptionFlag

`func (o *CatalogComponent) HasHideDescriptionFlag() bool`

HasHideDescriptionFlag returns a boolean if a field has been set.

### SetHideDescriptionFlagNil

`func (o *CatalogComponent) SetHideDescriptionFlagNil(b bool)`

 SetHideDescriptionFlagNil sets the value for HideDescriptionFlag to be an explicit nil

### UnsetHideDescriptionFlag
`func (o *CatalogComponent) UnsetHideDescriptionFlag()`

UnsetHideDescriptionFlag ensures that no value is present for HideDescriptionFlag, not even an explicit nil
### GetHideQuantityFlag

`func (o *CatalogComponent) GetHideQuantityFlag() bool`

GetHideQuantityFlag returns the HideQuantityFlag field if non-nil, zero value otherwise.

### GetHideQuantityFlagOk

`func (o *CatalogComponent) GetHideQuantityFlagOk() (*bool, bool)`

GetHideQuantityFlagOk returns a tuple with the HideQuantityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideQuantityFlag

`func (o *CatalogComponent) SetHideQuantityFlag(v bool)`

SetHideQuantityFlag sets HideQuantityFlag field to given value.

### HasHideQuantityFlag

`func (o *CatalogComponent) HasHideQuantityFlag() bool`

HasHideQuantityFlag returns a boolean if a field has been set.

### SetHideQuantityFlagNil

`func (o *CatalogComponent) SetHideQuantityFlagNil(b bool)`

 SetHideQuantityFlagNil sets the value for HideQuantityFlag to be an explicit nil

### UnsetHideQuantityFlag
`func (o *CatalogComponent) UnsetHideQuantityFlag()`

UnsetHideQuantityFlag ensures that no value is present for HideQuantityFlag, not even an explicit nil
### GetHideExtendedPriceFlag

`func (o *CatalogComponent) GetHideExtendedPriceFlag() bool`

GetHideExtendedPriceFlag returns the HideExtendedPriceFlag field if non-nil, zero value otherwise.

### GetHideExtendedPriceFlagOk

`func (o *CatalogComponent) GetHideExtendedPriceFlagOk() (*bool, bool)`

GetHideExtendedPriceFlagOk returns a tuple with the HideExtendedPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideExtendedPriceFlag

`func (o *CatalogComponent) SetHideExtendedPriceFlag(v bool)`

SetHideExtendedPriceFlag sets HideExtendedPriceFlag field to given value.

### HasHideExtendedPriceFlag

`func (o *CatalogComponent) HasHideExtendedPriceFlag() bool`

HasHideExtendedPriceFlag returns a boolean if a field has been set.

### SetHideExtendedPriceFlagNil

`func (o *CatalogComponent) SetHideExtendedPriceFlagNil(b bool)`

 SetHideExtendedPriceFlagNil sets the value for HideExtendedPriceFlag to be an explicit nil

### UnsetHideExtendedPriceFlag
`func (o *CatalogComponent) UnsetHideExtendedPriceFlag()`

UnsetHideExtendedPriceFlag ensures that no value is present for HideExtendedPriceFlag, not even an explicit nil
### GetParentCatalogItem

`func (o *CatalogComponent) GetParentCatalogItem() CatalogItemReference`

GetParentCatalogItem returns the ParentCatalogItem field if non-nil, zero value otherwise.

### GetParentCatalogItemOk

`func (o *CatalogComponent) GetParentCatalogItemOk() (*CatalogItemReference, bool)`

GetParentCatalogItemOk returns a tuple with the ParentCatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCatalogItem

`func (o *CatalogComponent) SetParentCatalogItem(v CatalogItemReference)`

SetParentCatalogItem sets ParentCatalogItem field to given value.

### HasParentCatalogItem

`func (o *CatalogComponent) HasParentCatalogItem() bool`

HasParentCatalogItem returns a boolean if a field has been set.

### GetPrice

`func (o *CatalogComponent) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogComponent) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogComponent) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogComponent) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogComponent) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogComponent) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCost

`func (o *CatalogComponent) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogComponent) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogComponent) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CatalogComponent) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CatalogComponent) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CatalogComponent) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetInfo

`func (o *CatalogComponent) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CatalogComponent) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CatalogComponent) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CatalogComponent) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


