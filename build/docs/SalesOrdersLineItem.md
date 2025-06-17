# SalesOrdersLineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PurchaseOrderNumber** | Pointer to **string** |  Max length: 100; | [optional] 
**SalesOrder** | [**SalesOrderReference**](SalesOrderReference.md) |  | 
**BillStatus** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **int32** |  | [optional] 
**QuantityCancelled** | Pointer to **int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSalesOrdersLineItem

`func NewSalesOrdersLineItem(salesOrder SalesOrderReference, ) *SalesOrdersLineItem`

NewSalesOrdersLineItem instantiates a new SalesOrdersLineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalesOrdersLineItemWithDefaults

`func NewSalesOrdersLineItemWithDefaults() *SalesOrdersLineItem`

NewSalesOrdersLineItemWithDefaults instantiates a new SalesOrdersLineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SalesOrdersLineItem) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SalesOrdersLineItem) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SalesOrdersLineItem) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SalesOrdersLineItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurchaseOrderNumber

`func (o *SalesOrdersLineItem) GetPurchaseOrderNumber() string`

GetPurchaseOrderNumber returns the PurchaseOrderNumber field if non-nil, zero value otherwise.

### GetPurchaseOrderNumberOk

`func (o *SalesOrdersLineItem) GetPurchaseOrderNumberOk() (*string, bool)`

GetPurchaseOrderNumberOk returns a tuple with the PurchaseOrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderNumber

`func (o *SalesOrdersLineItem) SetPurchaseOrderNumber(v string)`

SetPurchaseOrderNumber sets PurchaseOrderNumber field to given value.

### HasPurchaseOrderNumber

`func (o *SalesOrdersLineItem) HasPurchaseOrderNumber() bool`

HasPurchaseOrderNumber returns a boolean if a field has been set.

### GetSalesOrder

`func (o *SalesOrdersLineItem) GetSalesOrder() SalesOrderReference`

GetSalesOrder returns the SalesOrder field if non-nil, zero value otherwise.

### GetSalesOrderOk

`func (o *SalesOrdersLineItem) GetSalesOrderOk() (*SalesOrderReference, bool)`

GetSalesOrderOk returns a tuple with the SalesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesOrder

`func (o *SalesOrdersLineItem) SetSalesOrder(v SalesOrderReference)`

SetSalesOrder sets SalesOrder field to given value.


### GetBillStatus

`func (o *SalesOrdersLineItem) GetBillStatus() string`

GetBillStatus returns the BillStatus field if non-nil, zero value otherwise.

### GetBillStatusOk

`func (o *SalesOrdersLineItem) GetBillStatusOk() (*string, bool)`

GetBillStatusOk returns a tuple with the BillStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillStatus

`func (o *SalesOrdersLineItem) SetBillStatus(v string)`

SetBillStatus sets BillStatus field to given value.

### HasBillStatus

`func (o *SalesOrdersLineItem) HasBillStatus() bool`

HasBillStatus returns a boolean if a field has been set.

### GetQuantity

`func (o *SalesOrdersLineItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SalesOrdersLineItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SalesOrdersLineItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SalesOrdersLineItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityCancelled

`func (o *SalesOrdersLineItem) GetQuantityCancelled() int32`

GetQuantityCancelled returns the QuantityCancelled field if non-nil, zero value otherwise.

### GetQuantityCancelledOk

`func (o *SalesOrdersLineItem) GetQuantityCancelledOk() (*int32, bool)`

GetQuantityCancelledOk returns a tuple with the QuantityCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityCancelled

`func (o *SalesOrdersLineItem) SetQuantityCancelled(v int32)`

SetQuantityCancelled sets QuantityCancelled field to given value.

### HasQuantityCancelled

`func (o *SalesOrdersLineItem) HasQuantityCancelled() bool`

HasQuantityCancelled returns a boolean if a field has been set.

### GetInfo

`func (o *SalesOrdersLineItem) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SalesOrdersLineItem) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SalesOrdersLineItem) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SalesOrdersLineItem) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


