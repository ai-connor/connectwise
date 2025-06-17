# PurchaseOrderStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPurchaseOrderStatusInfo

`func NewPurchaseOrderStatusInfo() *PurchaseOrderStatusInfo`

NewPurchaseOrderStatusInfo instantiates a new PurchaseOrderStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderStatusInfoWithDefaults

`func NewPurchaseOrderStatusInfoWithDefaults() *PurchaseOrderStatusInfo`

NewPurchaseOrderStatusInfoWithDefaults instantiates a new PurchaseOrderStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderStatusInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderStatusInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderStatusInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderStatusInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PurchaseOrderStatusInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurchaseOrderStatusInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurchaseOrderStatusInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PurchaseOrderStatusInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *PurchaseOrderStatusInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PurchaseOrderStatusInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PurchaseOrderStatusInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PurchaseOrderStatusInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PurchaseOrderStatusInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PurchaseOrderStatusInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetClosedFlag

`func (o *PurchaseOrderStatusInfo) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PurchaseOrderStatusInfo) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PurchaseOrderStatusInfo) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PurchaseOrderStatusInfo) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PurchaseOrderStatusInfo) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PurchaseOrderStatusInfo) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetInactiveFlag

`func (o *PurchaseOrderStatusInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *PurchaseOrderStatusInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *PurchaseOrderStatusInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *PurchaseOrderStatusInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *PurchaseOrderStatusInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *PurchaseOrderStatusInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultClosedFlag

`func (o *PurchaseOrderStatusInfo) GetDefaultClosedFlag() bool`

GetDefaultClosedFlag returns the DefaultClosedFlag field if non-nil, zero value otherwise.

### GetDefaultClosedFlagOk

`func (o *PurchaseOrderStatusInfo) GetDefaultClosedFlagOk() (*bool, bool)`

GetDefaultClosedFlagOk returns a tuple with the DefaultClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClosedFlag

`func (o *PurchaseOrderStatusInfo) SetDefaultClosedFlag(v bool)`

SetDefaultClosedFlag sets DefaultClosedFlag field to given value.

### HasDefaultClosedFlag

`func (o *PurchaseOrderStatusInfo) HasDefaultClosedFlag() bool`

HasDefaultClosedFlag returns a boolean if a field has been set.

### SetDefaultClosedFlagNil

`func (o *PurchaseOrderStatusInfo) SetDefaultClosedFlagNil(b bool)`

 SetDefaultClosedFlagNil sets the value for DefaultClosedFlag to be an explicit nil

### UnsetDefaultClosedFlag
`func (o *PurchaseOrderStatusInfo) UnsetDefaultClosedFlag()`

UnsetDefaultClosedFlag ensures that no value is present for DefaultClosedFlag, not even an explicit nil
### GetSortOrder

`func (o *PurchaseOrderStatusInfo) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *PurchaseOrderStatusInfo) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *PurchaseOrderStatusInfo) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *PurchaseOrderStatusInfo) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *PurchaseOrderStatusInfo) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *PurchaseOrderStatusInfo) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetInfo

`func (o *PurchaseOrderStatusInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderStatusInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderStatusInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderStatusInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


