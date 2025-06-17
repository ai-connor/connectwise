# InventoryOnHand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CatalogItem** | Pointer to [**CatalogItemReference**](CatalogItemReference.md) |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**OnHand** | Pointer to **NullableInt32** |  | [optional] 
**SerialNumbers** | Pointer to [**[]OnHandSerialNumberReference**](OnHandSerialNumberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInventoryOnHand

`func NewInventoryOnHand() *InventoryOnHand`

NewInventoryOnHand instantiates a new InventoryOnHand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryOnHandWithDefaults

`func NewInventoryOnHandWithDefaults() *InventoryOnHand`

NewInventoryOnHandWithDefaults instantiates a new InventoryOnHand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InventoryOnHand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryOnHand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryOnHand) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InventoryOnHand) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCatalogItem

`func (o *InventoryOnHand) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *InventoryOnHand) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *InventoryOnHand) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.

### HasCatalogItem

`func (o *InventoryOnHand) HasCatalogItem() bool`

HasCatalogItem returns a boolean if a field has been set.

### GetWarehouse

`func (o *InventoryOnHand) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *InventoryOnHand) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *InventoryOnHand) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *InventoryOnHand) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *InventoryOnHand) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *InventoryOnHand) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *InventoryOnHand) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *InventoryOnHand) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetOnHand

`func (o *InventoryOnHand) GetOnHand() int32`

GetOnHand returns the OnHand field if non-nil, zero value otherwise.

### GetOnHandOk

`func (o *InventoryOnHand) GetOnHandOk() (*int32, bool)`

GetOnHandOk returns a tuple with the OnHand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHand

`func (o *InventoryOnHand) SetOnHand(v int32)`

SetOnHand sets OnHand field to given value.

### HasOnHand

`func (o *InventoryOnHand) HasOnHand() bool`

HasOnHand returns a boolean if a field has been set.

### SetOnHandNil

`func (o *InventoryOnHand) SetOnHandNil(b bool)`

 SetOnHandNil sets the value for OnHand to be an explicit nil

### UnsetOnHand
`func (o *InventoryOnHand) UnsetOnHand()`

UnsetOnHand ensures that no value is present for OnHand, not even an explicit nil
### GetSerialNumbers

`func (o *InventoryOnHand) GetSerialNumbers() []OnHandSerialNumberReference`

GetSerialNumbers returns the SerialNumbers field if non-nil, zero value otherwise.

### GetSerialNumbersOk

`func (o *InventoryOnHand) GetSerialNumbersOk() (*[]OnHandSerialNumberReference, bool)`

GetSerialNumbersOk returns a tuple with the SerialNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumbers

`func (o *InventoryOnHand) SetSerialNumbers(v []OnHandSerialNumberReference)`

SetSerialNumbers sets SerialNumbers field to given value.

### HasSerialNumbers

`func (o *InventoryOnHand) HasSerialNumbers() bool`

HasSerialNumbers returns a boolean if a field has been set.

### GetInfo

`func (o *InventoryOnHand) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InventoryOnHand) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InventoryOnHand) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InventoryOnHand) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


