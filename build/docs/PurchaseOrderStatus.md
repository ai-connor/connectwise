# PurchaseOrderStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**EmailTemplate** | Pointer to [**PurchaseOrderStatusEmailTemplateReference**](PurchaseOrderStatusEmailTemplateReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPurchaseOrderStatus

`func NewPurchaseOrderStatus(name string, ) *PurchaseOrderStatus`

NewPurchaseOrderStatus instantiates a new PurchaseOrderStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseOrderStatusWithDefaults

`func NewPurchaseOrderStatusWithDefaults() *PurchaseOrderStatus`

NewPurchaseOrderStatusWithDefaults instantiates a new PurchaseOrderStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PurchaseOrderStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PurchaseOrderStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PurchaseOrderStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PurchaseOrderStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PurchaseOrderStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PurchaseOrderStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PurchaseOrderStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *PurchaseOrderStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *PurchaseOrderStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *PurchaseOrderStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *PurchaseOrderStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *PurchaseOrderStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *PurchaseOrderStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetClosedFlag

`func (o *PurchaseOrderStatus) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *PurchaseOrderStatus) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *PurchaseOrderStatus) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *PurchaseOrderStatus) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *PurchaseOrderStatus) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *PurchaseOrderStatus) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetInactiveFlag

`func (o *PurchaseOrderStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *PurchaseOrderStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *PurchaseOrderStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *PurchaseOrderStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *PurchaseOrderStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *PurchaseOrderStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultClosedFlag

`func (o *PurchaseOrderStatus) GetDefaultClosedFlag() bool`

GetDefaultClosedFlag returns the DefaultClosedFlag field if non-nil, zero value otherwise.

### GetDefaultClosedFlagOk

`func (o *PurchaseOrderStatus) GetDefaultClosedFlagOk() (*bool, bool)`

GetDefaultClosedFlagOk returns a tuple with the DefaultClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultClosedFlag

`func (o *PurchaseOrderStatus) SetDefaultClosedFlag(v bool)`

SetDefaultClosedFlag sets DefaultClosedFlag field to given value.

### HasDefaultClosedFlag

`func (o *PurchaseOrderStatus) HasDefaultClosedFlag() bool`

HasDefaultClosedFlag returns a boolean if a field has been set.

### SetDefaultClosedFlagNil

`func (o *PurchaseOrderStatus) SetDefaultClosedFlagNil(b bool)`

 SetDefaultClosedFlagNil sets the value for DefaultClosedFlag to be an explicit nil

### UnsetDefaultClosedFlag
`func (o *PurchaseOrderStatus) UnsetDefaultClosedFlag()`

UnsetDefaultClosedFlag ensures that no value is present for DefaultClosedFlag, not even an explicit nil
### GetSortOrder

`func (o *PurchaseOrderStatus) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *PurchaseOrderStatus) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *PurchaseOrderStatus) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *PurchaseOrderStatus) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *PurchaseOrderStatus) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *PurchaseOrderStatus) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetEmailTemplate

`func (o *PurchaseOrderStatus) GetEmailTemplate() PurchaseOrderStatusEmailTemplateReference`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *PurchaseOrderStatus) GetEmailTemplateOk() (*PurchaseOrderStatusEmailTemplateReference, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *PurchaseOrderStatus) SetEmailTemplate(v PurchaseOrderStatusEmailTemplateReference)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *PurchaseOrderStatus) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetInfo

`func (o *PurchaseOrderStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PurchaseOrderStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PurchaseOrderStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PurchaseOrderStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


