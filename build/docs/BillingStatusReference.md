# BillingStatusReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsClosed** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBillingStatusReference

`func NewBillingStatusReference() *BillingStatusReference`

NewBillingStatusReference instantiates a new BillingStatusReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingStatusReferenceWithDefaults

`func NewBillingStatusReferenceWithDefaults() *BillingStatusReference`

NewBillingStatusReferenceWithDefaults instantiates a new BillingStatusReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BillingStatusReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingStatusReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingStatusReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BillingStatusReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BillingStatusReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BillingStatusReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BillingStatusReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BillingStatusReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BillingStatusReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BillingStatusReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsClosed

`func (o *BillingStatusReference) GetIsClosed() bool`

GetIsClosed returns the IsClosed field if non-nil, zero value otherwise.

### GetIsClosedOk

`func (o *BillingStatusReference) GetIsClosedOk() (*bool, bool)`

GetIsClosedOk returns a tuple with the IsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosed

`func (o *BillingStatusReference) SetIsClosed(v bool)`

SetIsClosed sets IsClosed field to given value.

### HasIsClosed

`func (o *BillingStatusReference) HasIsClosed() bool`

HasIsClosed returns a boolean if a field has been set.

### SetIsClosedNil

`func (o *BillingStatusReference) SetIsClosedNil(b bool)`

 SetIsClosedNil sets the value for IsClosed to be an explicit nil

### UnsetIsClosed
`func (o *BillingStatusReference) UnsetIsClosed()`

UnsetIsClosed ensures that no value is present for IsClosed, not even an explicit nil
### GetInfo

`func (o *BillingStatusReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BillingStatusReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BillingStatusReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BillingStatusReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


