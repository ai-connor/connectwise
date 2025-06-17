# AdjustmentDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CatalogItem** | [**CatalogItemReference**](CatalogItemReference.md) |  | 
**Description** | Pointer to **string** |  Max length: 50; | [optional] 
**QuantityOnHand** | Pointer to **NullableFloat64** |  | [optional] 
**UnitCost** | Pointer to **NullableFloat64** |  | [optional] 
**Warehouse** | [**WarehouseReference**](WarehouseReference.md) |  | 
**WarehouseBin** | [**WarehouseBinReference**](WarehouseBinReference.md) |  | 
**QuantityAdjusted** | **NullableInt32** |  | 
**SerialNumber** | Pointer to **string** |  Max length: 1000; | [optional] 
**Adjustment** | Pointer to [**AdjustmentReference**](AdjustmentReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAdjustmentDetail

`func NewAdjustmentDetail(catalogItem CatalogItemReference, warehouse WarehouseReference, warehouseBin WarehouseBinReference, quantityAdjusted NullableInt32, ) *AdjustmentDetail`

NewAdjustmentDetail instantiates a new AdjustmentDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustmentDetailWithDefaults

`func NewAdjustmentDetailWithDefaults() *AdjustmentDetail`

NewAdjustmentDetailWithDefaults instantiates a new AdjustmentDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AdjustmentDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdjustmentDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdjustmentDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AdjustmentDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCatalogItem

`func (o *AdjustmentDetail) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *AdjustmentDetail) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *AdjustmentDetail) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.


### GetDescription

`func (o *AdjustmentDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AdjustmentDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AdjustmentDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AdjustmentDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantityOnHand

`func (o *AdjustmentDetail) GetQuantityOnHand() float64`

GetQuantityOnHand returns the QuantityOnHand field if non-nil, zero value otherwise.

### GetQuantityOnHandOk

`func (o *AdjustmentDetail) GetQuantityOnHandOk() (*float64, bool)`

GetQuantityOnHandOk returns a tuple with the QuantityOnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityOnHand

`func (o *AdjustmentDetail) SetQuantityOnHand(v float64)`

SetQuantityOnHand sets QuantityOnHand field to given value.

### HasQuantityOnHand

`func (o *AdjustmentDetail) HasQuantityOnHand() bool`

HasQuantityOnHand returns a boolean if a field has been set.

### SetQuantityOnHandNil

`func (o *AdjustmentDetail) SetQuantityOnHandNil(b bool)`

 SetQuantityOnHandNil sets the value for QuantityOnHand to be an explicit nil

### UnsetQuantityOnHand
`func (o *AdjustmentDetail) UnsetQuantityOnHand()`

UnsetQuantityOnHand ensures that no value is present for QuantityOnHand, not even an explicit nil
### GetUnitCost

`func (o *AdjustmentDetail) GetUnitCost() float64`

GetUnitCost returns the UnitCost field if non-nil, zero value otherwise.

### GetUnitCostOk

`func (o *AdjustmentDetail) GetUnitCostOk() (*float64, bool)`

GetUnitCostOk returns a tuple with the UnitCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitCost

`func (o *AdjustmentDetail) SetUnitCost(v float64)`

SetUnitCost sets UnitCost field to given value.

### HasUnitCost

`func (o *AdjustmentDetail) HasUnitCost() bool`

HasUnitCost returns a boolean if a field has been set.

### SetUnitCostNil

`func (o *AdjustmentDetail) SetUnitCostNil(b bool)`

 SetUnitCostNil sets the value for UnitCost to be an explicit nil

### UnsetUnitCost
`func (o *AdjustmentDetail) UnsetUnitCost()`

UnsetUnitCost ensures that no value is present for UnitCost, not even an explicit nil
### GetWarehouse

`func (o *AdjustmentDetail) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *AdjustmentDetail) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *AdjustmentDetail) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.


### GetWarehouseBin

`func (o *AdjustmentDetail) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *AdjustmentDetail) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *AdjustmentDetail) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.


### GetQuantityAdjusted

`func (o *AdjustmentDetail) GetQuantityAdjusted() int32`

GetQuantityAdjusted returns the QuantityAdjusted field if non-nil, zero value otherwise.

### GetQuantityAdjustedOk

`func (o *AdjustmentDetail) GetQuantityAdjustedOk() (*int32, bool)`

GetQuantityAdjustedOk returns a tuple with the QuantityAdjusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAdjusted

`func (o *AdjustmentDetail) SetQuantityAdjusted(v int32)`

SetQuantityAdjusted sets QuantityAdjusted field to given value.


### SetQuantityAdjustedNil

`func (o *AdjustmentDetail) SetQuantityAdjustedNil(b bool)`

 SetQuantityAdjustedNil sets the value for QuantityAdjusted to be an explicit nil

### UnsetQuantityAdjusted
`func (o *AdjustmentDetail) UnsetQuantityAdjusted()`

UnsetQuantityAdjusted ensures that no value is present for QuantityAdjusted, not even an explicit nil
### GetSerialNumber

`func (o *AdjustmentDetail) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AdjustmentDetail) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AdjustmentDetail) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AdjustmentDetail) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetAdjustment

`func (o *AdjustmentDetail) GetAdjustment() AdjustmentReference`

GetAdjustment returns the Adjustment field if non-nil, zero value otherwise.

### GetAdjustmentOk

`func (o *AdjustmentDetail) GetAdjustmentOk() (*AdjustmentReference, bool)`

GetAdjustmentOk returns a tuple with the Adjustment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustment

`func (o *AdjustmentDetail) SetAdjustment(v AdjustmentReference)`

SetAdjustment sets Adjustment field to given value.

### HasAdjustment

`func (o *AdjustmentDetail) HasAdjustment() bool`

HasAdjustment returns a boolean if a field has been set.

### GetInfo

`func (o *AdjustmentDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AdjustmentDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AdjustmentDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AdjustmentDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


