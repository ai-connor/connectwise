# CatalogItemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ProductClass** | Pointer to **NullableString** |  | [optional] 
**SerializedCostFlag** | Pointer to **NullableBool** |  | [optional] 
**Price** | Pointer to **NullableFloat64** |  | [optional] 
**Cost** | Pointer to **NullableFloat64** |  | [optional] 
**TaxableFlag** | Pointer to **NullableBool** |  | [optional] 
**DropShipFlag** | Pointer to **NullableBool** |  | [optional] 
**SpecialOrderFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomerDescription** | Pointer to **string** |  | [optional] 
**ManufacturerPartNumber** | Pointer to **string** |  | [optional] 
**VendorSku** | Pointer to **string** |  | [optional] 
**BillableOption** | Pointer to **NullableString** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCatalogItemInfo

`func NewCatalogItemInfo() *CatalogItemInfo`

NewCatalogItemInfo instantiates a new CatalogItemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemInfoWithDefaults

`func NewCatalogItemInfoWithDefaults() *CatalogItemInfo`

NewCatalogItemInfoWithDefaults instantiates a new CatalogItemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogItemInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogItemInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogItemInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogItemInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *CatalogItemInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CatalogItemInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CatalogItemInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CatalogItemInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetDescription

`func (o *CatalogItemInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CatalogItemInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CatalogItemInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CatalogItemInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *CatalogItemInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CatalogItemInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CatalogItemInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CatalogItemInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CatalogItemInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CatalogItemInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetProductClass

`func (o *CatalogItemInfo) GetProductClass() string`

GetProductClass returns the ProductClass field if non-nil, zero value otherwise.

### GetProductClassOk

`func (o *CatalogItemInfo) GetProductClassOk() (*string, bool)`

GetProductClassOk returns a tuple with the ProductClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductClass

`func (o *CatalogItemInfo) SetProductClass(v string)`

SetProductClass sets ProductClass field to given value.

### HasProductClass

`func (o *CatalogItemInfo) HasProductClass() bool`

HasProductClass returns a boolean if a field has been set.

### SetProductClassNil

`func (o *CatalogItemInfo) SetProductClassNil(b bool)`

 SetProductClassNil sets the value for ProductClass to be an explicit nil

### UnsetProductClass
`func (o *CatalogItemInfo) UnsetProductClass()`

UnsetProductClass ensures that no value is present for ProductClass, not even an explicit nil
### GetSerializedCostFlag

`func (o *CatalogItemInfo) GetSerializedCostFlag() bool`

GetSerializedCostFlag returns the SerializedCostFlag field if non-nil, zero value otherwise.

### GetSerializedCostFlagOk

`func (o *CatalogItemInfo) GetSerializedCostFlagOk() (*bool, bool)`

GetSerializedCostFlagOk returns a tuple with the SerializedCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedCostFlag

`func (o *CatalogItemInfo) SetSerializedCostFlag(v bool)`

SetSerializedCostFlag sets SerializedCostFlag field to given value.

### HasSerializedCostFlag

`func (o *CatalogItemInfo) HasSerializedCostFlag() bool`

HasSerializedCostFlag returns a boolean if a field has been set.

### SetSerializedCostFlagNil

`func (o *CatalogItemInfo) SetSerializedCostFlagNil(b bool)`

 SetSerializedCostFlagNil sets the value for SerializedCostFlag to be an explicit nil

### UnsetSerializedCostFlag
`func (o *CatalogItemInfo) UnsetSerializedCostFlag()`

UnsetSerializedCostFlag ensures that no value is present for SerializedCostFlag, not even an explicit nil
### GetPrice

`func (o *CatalogItemInfo) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogItemInfo) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogItemInfo) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogItemInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CatalogItemInfo) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CatalogItemInfo) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCost

`func (o *CatalogItemInfo) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *CatalogItemInfo) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *CatalogItemInfo) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *CatalogItemInfo) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *CatalogItemInfo) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *CatalogItemInfo) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetTaxableFlag

`func (o *CatalogItemInfo) GetTaxableFlag() bool`

GetTaxableFlag returns the TaxableFlag field if non-nil, zero value otherwise.

### GetTaxableFlagOk

`func (o *CatalogItemInfo) GetTaxableFlagOk() (*bool, bool)`

GetTaxableFlagOk returns a tuple with the TaxableFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxableFlag

`func (o *CatalogItemInfo) SetTaxableFlag(v bool)`

SetTaxableFlag sets TaxableFlag field to given value.

### HasTaxableFlag

`func (o *CatalogItemInfo) HasTaxableFlag() bool`

HasTaxableFlag returns a boolean if a field has been set.

### SetTaxableFlagNil

`func (o *CatalogItemInfo) SetTaxableFlagNil(b bool)`

 SetTaxableFlagNil sets the value for TaxableFlag to be an explicit nil

### UnsetTaxableFlag
`func (o *CatalogItemInfo) UnsetTaxableFlag()`

UnsetTaxableFlag ensures that no value is present for TaxableFlag, not even an explicit nil
### GetDropShipFlag

`func (o *CatalogItemInfo) GetDropShipFlag() bool`

GetDropShipFlag returns the DropShipFlag field if non-nil, zero value otherwise.

### GetDropShipFlagOk

`func (o *CatalogItemInfo) GetDropShipFlagOk() (*bool, bool)`

GetDropShipFlagOk returns a tuple with the DropShipFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropShipFlag

`func (o *CatalogItemInfo) SetDropShipFlag(v bool)`

SetDropShipFlag sets DropShipFlag field to given value.

### HasDropShipFlag

`func (o *CatalogItemInfo) HasDropShipFlag() bool`

HasDropShipFlag returns a boolean if a field has been set.

### SetDropShipFlagNil

`func (o *CatalogItemInfo) SetDropShipFlagNil(b bool)`

 SetDropShipFlagNil sets the value for DropShipFlag to be an explicit nil

### UnsetDropShipFlag
`func (o *CatalogItemInfo) UnsetDropShipFlag()`

UnsetDropShipFlag ensures that no value is present for DropShipFlag, not even an explicit nil
### GetSpecialOrderFlag

`func (o *CatalogItemInfo) GetSpecialOrderFlag() bool`

GetSpecialOrderFlag returns the SpecialOrderFlag field if non-nil, zero value otherwise.

### GetSpecialOrderFlagOk

`func (o *CatalogItemInfo) GetSpecialOrderFlagOk() (*bool, bool)`

GetSpecialOrderFlagOk returns a tuple with the SpecialOrderFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialOrderFlag

`func (o *CatalogItemInfo) SetSpecialOrderFlag(v bool)`

SetSpecialOrderFlag sets SpecialOrderFlag field to given value.

### HasSpecialOrderFlag

`func (o *CatalogItemInfo) HasSpecialOrderFlag() bool`

HasSpecialOrderFlag returns a boolean if a field has been set.

### SetSpecialOrderFlagNil

`func (o *CatalogItemInfo) SetSpecialOrderFlagNil(b bool)`

 SetSpecialOrderFlagNil sets the value for SpecialOrderFlag to be an explicit nil

### UnsetSpecialOrderFlag
`func (o *CatalogItemInfo) UnsetSpecialOrderFlag()`

UnsetSpecialOrderFlag ensures that no value is present for SpecialOrderFlag, not even an explicit nil
### GetCustomerDescription

`func (o *CatalogItemInfo) GetCustomerDescription() string`

GetCustomerDescription returns the CustomerDescription field if non-nil, zero value otherwise.

### GetCustomerDescriptionOk

`func (o *CatalogItemInfo) GetCustomerDescriptionOk() (*string, bool)`

GetCustomerDescriptionOk returns a tuple with the CustomerDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerDescription

`func (o *CatalogItemInfo) SetCustomerDescription(v string)`

SetCustomerDescription sets CustomerDescription field to given value.

### HasCustomerDescription

`func (o *CatalogItemInfo) HasCustomerDescription() bool`

HasCustomerDescription returns a boolean if a field has been set.

### GetManufacturerPartNumber

`func (o *CatalogItemInfo) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *CatalogItemInfo) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *CatalogItemInfo) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *CatalogItemInfo) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetVendorSku

`func (o *CatalogItemInfo) GetVendorSku() string`

GetVendorSku returns the VendorSku field if non-nil, zero value otherwise.

### GetVendorSkuOk

`func (o *CatalogItemInfo) GetVendorSkuOk() (*string, bool)`

GetVendorSkuOk returns a tuple with the VendorSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSku

`func (o *CatalogItemInfo) SetVendorSku(v string)`

SetVendorSku sets VendorSku field to given value.

### HasVendorSku

`func (o *CatalogItemInfo) HasVendorSku() bool`

HasVendorSku returns a boolean if a field has been set.

### GetBillableOption

`func (o *CatalogItemInfo) GetBillableOption() string`

GetBillableOption returns the BillableOption field if non-nil, zero value otherwise.

### GetBillableOptionOk

`func (o *CatalogItemInfo) GetBillableOptionOk() (*string, bool)`

GetBillableOptionOk returns a tuple with the BillableOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableOption

`func (o *CatalogItemInfo) SetBillableOption(v string)`

SetBillableOption sets BillableOption field to given value.

### HasBillableOption

`func (o *CatalogItemInfo) HasBillableOption() bool`

HasBillableOption returns a boolean if a field has been set.

### SetBillableOptionNil

`func (o *CatalogItemInfo) SetBillableOptionNil(b bool)`

 SetBillableOptionNil sets the value for BillableOption to be an explicit nil

### UnsetBillableOption
`func (o *CatalogItemInfo) UnsetBillableOption()`

UnsetBillableOption ensures that no value is present for BillableOption, not even an explicit nil
### GetInfo

`func (o *CatalogItemInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CatalogItemInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CatalogItemInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CatalogItemInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


