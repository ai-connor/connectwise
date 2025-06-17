# OnHandSerialNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Serial** | Pointer to **string** |  | [optional] 
**CatalogItem** | Pointer to [**CatalogItemReference**](CatalogItemReference.md) |  | [optional] 
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**WarehouseBin** | Pointer to [**WarehouseBinReference**](WarehouseBinReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOnHandSerialNumber

`func NewOnHandSerialNumber() *OnHandSerialNumber`

NewOnHandSerialNumber instantiates a new OnHandSerialNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnHandSerialNumberWithDefaults

`func NewOnHandSerialNumberWithDefaults() *OnHandSerialNumber`

NewOnHandSerialNumberWithDefaults instantiates a new OnHandSerialNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnHandSerialNumber) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnHandSerialNumber) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnHandSerialNumber) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OnHandSerialNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSerial

`func (o *OnHandSerialNumber) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *OnHandSerialNumber) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *OnHandSerialNumber) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *OnHandSerialNumber) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetCatalogItem

`func (o *OnHandSerialNumber) GetCatalogItem() CatalogItemReference`

GetCatalogItem returns the CatalogItem field if non-nil, zero value otherwise.

### GetCatalogItemOk

`func (o *OnHandSerialNumber) GetCatalogItemOk() (*CatalogItemReference, bool)`

GetCatalogItemOk returns a tuple with the CatalogItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItem

`func (o *OnHandSerialNumber) SetCatalogItem(v CatalogItemReference)`

SetCatalogItem sets CatalogItem field to given value.

### HasCatalogItem

`func (o *OnHandSerialNumber) HasCatalogItem() bool`

HasCatalogItem returns a boolean if a field has been set.

### GetWarehouse

`func (o *OnHandSerialNumber) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *OnHandSerialNumber) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *OnHandSerialNumber) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *OnHandSerialNumber) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetWarehouseBin

`func (o *OnHandSerialNumber) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *OnHandSerialNumber) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *OnHandSerialNumber) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.

### HasWarehouseBin

`func (o *OnHandSerialNumber) HasWarehouseBin() bool`

HasWarehouseBin returns a boolean if a field has been set.

### GetInfo

`func (o *OnHandSerialNumber) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OnHandSerialNumber) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OnHandSerialNumber) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OnHandSerialNumber) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


