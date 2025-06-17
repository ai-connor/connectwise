# ProductPickingShippingDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PickedQuantity** | Pointer to **NullableInt32** |  | [optional] 
**ShippedQuantity** | Pointer to **NullableInt32** |  | [optional] 
**Warehouse** | [**WarehouseReference**](WarehouseReference.md) |  | 
**WarehouseBin** | [**WarehouseBinReference**](WarehouseBinReference.md) |  | 
**ShipmentMethod** | Pointer to [**ShipmentMethodReference**](ShipmentMethodReference.md) |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**SerialNumberIds** | Pointer to **[]int32** |  | [optional] 
**TrackingNumber** | Pointer to **string** |  | [optional] 
**ProductItem** | Pointer to [**ProductItemReference**](ProductItemReference.md) |  | [optional] 
**LineNumber** | Pointer to **NullableInt32** |  | [optional] 
**Quantity** | Pointer to **NullableInt32** |  | [optional] 
**ExpectedArrivalDate** | Pointer to **time.Time** |  | [optional] 
**ShipmentDate** | Pointer to **time.Time** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProductPickingShippingDetail

`func NewProductPickingShippingDetail(warehouse WarehouseReference, warehouseBin WarehouseBinReference, ) *ProductPickingShippingDetail`

NewProductPickingShippingDetail instantiates a new ProductPickingShippingDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductPickingShippingDetailWithDefaults

`func NewProductPickingShippingDetailWithDefaults() *ProductPickingShippingDetail`

NewProductPickingShippingDetailWithDefaults instantiates a new ProductPickingShippingDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProductPickingShippingDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductPickingShippingDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductPickingShippingDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProductPickingShippingDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPickedQuantity

`func (o *ProductPickingShippingDetail) GetPickedQuantity() int32`

GetPickedQuantity returns the PickedQuantity field if non-nil, zero value otherwise.

### GetPickedQuantityOk

`func (o *ProductPickingShippingDetail) GetPickedQuantityOk() (*int32, bool)`

GetPickedQuantityOk returns a tuple with the PickedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickedQuantity

`func (o *ProductPickingShippingDetail) SetPickedQuantity(v int32)`

SetPickedQuantity sets PickedQuantity field to given value.

### HasPickedQuantity

`func (o *ProductPickingShippingDetail) HasPickedQuantity() bool`

HasPickedQuantity returns a boolean if a field has been set.

### SetPickedQuantityNil

`func (o *ProductPickingShippingDetail) SetPickedQuantityNil(b bool)`

 SetPickedQuantityNil sets the value for PickedQuantity to be an explicit nil

### UnsetPickedQuantity
`func (o *ProductPickingShippingDetail) UnsetPickedQuantity()`

UnsetPickedQuantity ensures that no value is present for PickedQuantity, not even an explicit nil
### GetShippedQuantity

`func (o *ProductPickingShippingDetail) GetShippedQuantity() int32`

GetShippedQuantity returns the ShippedQuantity field if non-nil, zero value otherwise.

### GetShippedQuantityOk

`func (o *ProductPickingShippingDetail) GetShippedQuantityOk() (*int32, bool)`

GetShippedQuantityOk returns a tuple with the ShippedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedQuantity

`func (o *ProductPickingShippingDetail) SetShippedQuantity(v int32)`

SetShippedQuantity sets ShippedQuantity field to given value.

### HasShippedQuantity

`func (o *ProductPickingShippingDetail) HasShippedQuantity() bool`

HasShippedQuantity returns a boolean if a field has been set.

### SetShippedQuantityNil

`func (o *ProductPickingShippingDetail) SetShippedQuantityNil(b bool)`

 SetShippedQuantityNil sets the value for ShippedQuantity to be an explicit nil

### UnsetShippedQuantity
`func (o *ProductPickingShippingDetail) UnsetShippedQuantity()`

UnsetShippedQuantity ensures that no value is present for ShippedQuantity, not even an explicit nil
### GetWarehouse

`func (o *ProductPickingShippingDetail) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *ProductPickingShippingDetail) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *ProductPickingShippingDetail) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.


### GetWarehouseBin

`func (o *ProductPickingShippingDetail) GetWarehouseBin() WarehouseBinReference`

GetWarehouseBin returns the WarehouseBin field if non-nil, zero value otherwise.

### GetWarehouseBinOk

`func (o *ProductPickingShippingDetail) GetWarehouseBinOk() (*WarehouseBinReference, bool)`

GetWarehouseBinOk returns a tuple with the WarehouseBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseBin

`func (o *ProductPickingShippingDetail) SetWarehouseBin(v WarehouseBinReference)`

SetWarehouseBin sets WarehouseBin field to given value.


### GetShipmentMethod

`func (o *ProductPickingShippingDetail) GetShipmentMethod() ShipmentMethodReference`

GetShipmentMethod returns the ShipmentMethod field if non-nil, zero value otherwise.

### GetShipmentMethodOk

`func (o *ProductPickingShippingDetail) GetShipmentMethodOk() (*ShipmentMethodReference, bool)`

GetShipmentMethodOk returns a tuple with the ShipmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentMethod

`func (o *ProductPickingShippingDetail) SetShipmentMethod(v ShipmentMethodReference)`

SetShipmentMethod sets ShipmentMethod field to given value.

### HasShipmentMethod

`func (o *ProductPickingShippingDetail) HasShipmentMethod() bool`

HasShipmentMethod returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ProductPickingShippingDetail) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ProductPickingShippingDetail) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ProductPickingShippingDetail) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ProductPickingShippingDetail) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSerialNumberIds

`func (o *ProductPickingShippingDetail) GetSerialNumberIds() []int32`

GetSerialNumberIds returns the SerialNumberIds field if non-nil, zero value otherwise.

### GetSerialNumberIdsOk

`func (o *ProductPickingShippingDetail) GetSerialNumberIdsOk() (*[]int32, bool)`

GetSerialNumberIdsOk returns a tuple with the SerialNumberIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumberIds

`func (o *ProductPickingShippingDetail) SetSerialNumberIds(v []int32)`

SetSerialNumberIds sets SerialNumberIds field to given value.

### HasSerialNumberIds

`func (o *ProductPickingShippingDetail) HasSerialNumberIds() bool`

HasSerialNumberIds returns a boolean if a field has been set.

### GetTrackingNumber

`func (o *ProductPickingShippingDetail) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ProductPickingShippingDetail) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ProductPickingShippingDetail) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ProductPickingShippingDetail) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### GetProductItem

`func (o *ProductPickingShippingDetail) GetProductItem() ProductItemReference`

GetProductItem returns the ProductItem field if non-nil, zero value otherwise.

### GetProductItemOk

`func (o *ProductPickingShippingDetail) GetProductItemOk() (*ProductItemReference, bool)`

GetProductItemOk returns a tuple with the ProductItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItem

`func (o *ProductPickingShippingDetail) SetProductItem(v ProductItemReference)`

SetProductItem sets ProductItem field to given value.

### HasProductItem

`func (o *ProductPickingShippingDetail) HasProductItem() bool`

HasProductItem returns a boolean if a field has been set.

### GetLineNumber

`func (o *ProductPickingShippingDetail) GetLineNumber() int32`

GetLineNumber returns the LineNumber field if non-nil, zero value otherwise.

### GetLineNumberOk

`func (o *ProductPickingShippingDetail) GetLineNumberOk() (*int32, bool)`

GetLineNumberOk returns a tuple with the LineNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumber

`func (o *ProductPickingShippingDetail) SetLineNumber(v int32)`

SetLineNumber sets LineNumber field to given value.

### HasLineNumber

`func (o *ProductPickingShippingDetail) HasLineNumber() bool`

HasLineNumber returns a boolean if a field has been set.

### SetLineNumberNil

`func (o *ProductPickingShippingDetail) SetLineNumberNil(b bool)`

 SetLineNumberNil sets the value for LineNumber to be an explicit nil

### UnsetLineNumber
`func (o *ProductPickingShippingDetail) UnsetLineNumber()`

UnsetLineNumber ensures that no value is present for LineNumber, not even an explicit nil
### GetQuantity

`func (o *ProductPickingShippingDetail) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ProductPickingShippingDetail) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ProductPickingShippingDetail) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ProductPickingShippingDetail) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *ProductPickingShippingDetail) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *ProductPickingShippingDetail) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetExpectedArrivalDate

`func (o *ProductPickingShippingDetail) GetExpectedArrivalDate() time.Time`

GetExpectedArrivalDate returns the ExpectedArrivalDate field if non-nil, zero value otherwise.

### GetExpectedArrivalDateOk

`func (o *ProductPickingShippingDetail) GetExpectedArrivalDateOk() (*time.Time, bool)`

GetExpectedArrivalDateOk returns a tuple with the ExpectedArrivalDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedArrivalDate

`func (o *ProductPickingShippingDetail) SetExpectedArrivalDate(v time.Time)`

SetExpectedArrivalDate sets ExpectedArrivalDate field to given value.

### HasExpectedArrivalDate

`func (o *ProductPickingShippingDetail) HasExpectedArrivalDate() bool`

HasExpectedArrivalDate returns a boolean if a field has been set.

### GetShipmentDate

`func (o *ProductPickingShippingDetail) GetShipmentDate() time.Time`

GetShipmentDate returns the ShipmentDate field if non-nil, zero value otherwise.

### GetShipmentDateOk

`func (o *ProductPickingShippingDetail) GetShipmentDateOk() (*time.Time, bool)`

GetShipmentDateOk returns a tuple with the ShipmentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentDate

`func (o *ProductPickingShippingDetail) SetShipmentDate(v time.Time)`

SetShipmentDate sets ShipmentDate field to given value.

### HasShipmentDate

`func (o *ProductPickingShippingDetail) HasShipmentDate() bool`

HasShipmentDate returns a boolean if a field has been set.

### GetInfo

`func (o *ProductPickingShippingDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProductPickingShippingDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProductPickingShippingDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProductPickingShippingDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


