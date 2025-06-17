# MinimumStockByWarehouse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Warehouse** | [**WarehouseReference**](WarehouseReference.md) |  | 
**MinimumStock** | **NullableInt32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMinimumStockByWarehouse

`func NewMinimumStockByWarehouse(warehouse WarehouseReference, minimumStock NullableInt32, ) *MinimumStockByWarehouse`

NewMinimumStockByWarehouse instantiates a new MinimumStockByWarehouse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimumStockByWarehouseWithDefaults

`func NewMinimumStockByWarehouseWithDefaults() *MinimumStockByWarehouse`

NewMinimumStockByWarehouseWithDefaults instantiates a new MinimumStockByWarehouse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MinimumStockByWarehouse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MinimumStockByWarehouse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MinimumStockByWarehouse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MinimumStockByWarehouse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWarehouse

`func (o *MinimumStockByWarehouse) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *MinimumStockByWarehouse) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *MinimumStockByWarehouse) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.


### GetMinimumStock

`func (o *MinimumStockByWarehouse) GetMinimumStock() int32`

GetMinimumStock returns the MinimumStock field if non-nil, zero value otherwise.

### GetMinimumStockOk

`func (o *MinimumStockByWarehouse) GetMinimumStockOk() (*int32, bool)`

GetMinimumStockOk returns a tuple with the MinimumStock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumStock

`func (o *MinimumStockByWarehouse) SetMinimumStock(v int32)`

SetMinimumStock sets MinimumStock field to given value.


### SetMinimumStockNil

`func (o *MinimumStockByWarehouse) SetMinimumStockNil(b bool)`

 SetMinimumStockNil sets the value for MinimumStock to be an explicit nil

### UnsetMinimumStock
`func (o *MinimumStockByWarehouse) UnsetMinimumStock()`

UnsetMinimumStock ensures that no value is present for MinimumStock, not even an explicit nil
### GetInfo

`func (o *MinimumStockByWarehouse) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MinimumStockByWarehouse) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MinimumStockByWarehouse) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MinimumStockByWarehouse) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


