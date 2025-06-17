# ChangeOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**PurchaseHeaderRecId** | **int32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewChangeOrder

`func NewChangeOrder(purchaseHeaderRecId int32, ) *ChangeOrder`

NewChangeOrder instantiates a new ChangeOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeOrderWithDefaults

`func NewChangeOrderWithDefaults() *ChangeOrder`

NewChangeOrderWithDefaults instantiates a new ChangeOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChangeOrder) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChangeOrder) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChangeOrder) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ChangeOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurchaseHeaderRecId

`func (o *ChangeOrder) GetPurchaseHeaderRecId() int32`

GetPurchaseHeaderRecId returns the PurchaseHeaderRecId field if non-nil, zero value otherwise.

### GetPurchaseHeaderRecIdOk

`func (o *ChangeOrder) GetPurchaseHeaderRecIdOk() (*int32, bool)`

GetPurchaseHeaderRecIdOk returns a tuple with the PurchaseHeaderRecId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseHeaderRecId

`func (o *ChangeOrder) SetPurchaseHeaderRecId(v int32)`

SetPurchaseHeaderRecId sets PurchaseHeaderRecId field to given value.


### GetInfo

`func (o *ChangeOrder) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ChangeOrder) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ChangeOrder) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ChangeOrder) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


