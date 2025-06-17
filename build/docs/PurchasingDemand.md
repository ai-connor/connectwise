# PurchasingDemand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Warehouse** | Pointer to [**WarehouseReference**](WarehouseReference.md) |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Products** | Pointer to [**[]ProductDemand**](ProductDemand.md) |  | [optional] 
**PurchaseOrder** | Pointer to [**PurchaseOrder**](PurchaseOrder.md) |  | [optional] 

## Methods

### NewPurchasingDemand

`func NewPurchasingDemand() *PurchasingDemand`

NewPurchasingDemand instantiates a new PurchasingDemand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchasingDemandWithDefaults

`func NewPurchasingDemandWithDefaults() *PurchasingDemand`

NewPurchasingDemandWithDefaults instantiates a new PurchasingDemand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWarehouse

`func (o *PurchasingDemand) GetWarehouse() WarehouseReference`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *PurchasingDemand) GetWarehouseOk() (*WarehouseReference, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *PurchasingDemand) SetWarehouse(v WarehouseReference)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *PurchasingDemand) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.

### GetVendor

`func (o *PurchasingDemand) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *PurchasingDemand) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *PurchasingDemand) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *PurchasingDemand) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetProducts

`func (o *PurchasingDemand) GetProducts() []ProductDemand`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *PurchasingDemand) GetProductsOk() (*[]ProductDemand, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *PurchasingDemand) SetProducts(v []ProductDemand)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *PurchasingDemand) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetPurchaseOrder

`func (o *PurchasingDemand) GetPurchaseOrder() PurchaseOrder`

GetPurchaseOrder returns the PurchaseOrder field if non-nil, zero value otherwise.

### GetPurchaseOrderOk

`func (o *PurchasingDemand) GetPurchaseOrderOk() (*PurchaseOrder, bool)`

GetPurchaseOrderOk returns a tuple with the PurchaseOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrder

`func (o *PurchasingDemand) SetPurchaseOrder(v PurchaseOrder)`

SetPurchaseOrder sets PurchaseOrder field to given value.

### HasPurchaseOrder

`func (o *PurchasingDemand) HasPurchaseOrder() bool`

HasPurchaseOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


