# BillingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**SortOrder** | Pointer to **NullableInt32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**SentFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingStatus

`func NewBillingStatus(name string, ) *BillingStatus`

NewBillingStatus instantiates a new BillingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingStatusWithDefaults

`func NewBillingStatusWithDefaults() *BillingStatus`

NewBillingStatusWithDefaults instantiates a new BillingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BillingStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingStatus) SetName(v string)`

SetName sets Name field to given value.


### GetSortOrder

`func (o *BillingStatus) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *BillingStatus) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *BillingStatus) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *BillingStatus) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### SetSortOrderNil

`func (o *BillingStatus) SetSortOrderNil(b bool)`

 SetSortOrderNil sets the value for SortOrder to be an explicit nil

### UnsetSortOrder
`func (o *BillingStatus) UnsetSortOrder()`

UnsetSortOrder ensures that no value is present for SortOrder, not even an explicit nil
### GetDefaultFlag

`func (o *BillingStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BillingStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BillingStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BillingStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BillingStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BillingStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetClosedFlag

`func (o *BillingStatus) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *BillingStatus) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *BillingStatus) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *BillingStatus) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *BillingStatus) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *BillingStatus) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetInactiveFlag

`func (o *BillingStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BillingStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BillingStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BillingStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BillingStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BillingStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSentFlag

`func (o *BillingStatus) GetSentFlag() bool`

GetSentFlag returns the SentFlag field if non-nil, zero value otherwise.

### GetSentFlagOk

`func (o *BillingStatus) GetSentFlagOk() (*bool, bool)`

GetSentFlagOk returns a tuple with the SentFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentFlag

`func (o *BillingStatus) SetSentFlag(v bool)`

SetSentFlag sets SentFlag field to given value.

### HasSentFlag

`func (o *BillingStatus) HasSentFlag() bool`

HasSentFlag returns a boolean if a field has been set.

### SetSentFlagNil

`func (o *BillingStatus) SetSentFlagNil(b bool)`

 SetSentFlagNil sets the value for SentFlag to be an explicit nil

### UnsetSentFlag
`func (o *BillingStatus) UnsetSentFlag()`

UnsetSentFlag ensures that no value is present for SentFlag, not even an explicit nil
### GetInfo

`func (o *BillingStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


