# ExpenseTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**AmountCaption** | Pointer to **string** |  | [optional] 
**MileageFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseTypeInfo

`func NewExpenseTypeInfo() *ExpenseTypeInfo`

NewExpenseTypeInfo instantiates a new ExpenseTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseTypeInfoWithDefaults

`func NewExpenseTypeInfoWithDefaults() *ExpenseTypeInfo`

NewExpenseTypeInfoWithDefaults instantiates a new ExpenseTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ExpenseTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExpenseTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExpenseTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExpenseTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *ExpenseTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ExpenseTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ExpenseTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ExpenseTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ExpenseTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ExpenseTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetAmountCaption

`func (o *ExpenseTypeInfo) GetAmountCaption() string`

GetAmountCaption returns the AmountCaption field if non-nil, zero value otherwise.

### GetAmountCaptionOk

`func (o *ExpenseTypeInfo) GetAmountCaptionOk() (*string, bool)`

GetAmountCaptionOk returns a tuple with the AmountCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCaption

`func (o *ExpenseTypeInfo) SetAmountCaption(v string)`

SetAmountCaption sets AmountCaption field to given value.

### HasAmountCaption

`func (o *ExpenseTypeInfo) HasAmountCaption() bool`

HasAmountCaption returns a boolean if a field has been set.

### GetMileageFlag

`func (o *ExpenseTypeInfo) GetMileageFlag() bool`

GetMileageFlag returns the MileageFlag field if non-nil, zero value otherwise.

### GetMileageFlagOk

`func (o *ExpenseTypeInfo) GetMileageFlagOk() (*bool, bool)`

GetMileageFlagOk returns a tuple with the MileageFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMileageFlag

`func (o *ExpenseTypeInfo) SetMileageFlag(v bool)`

SetMileageFlag sets MileageFlag field to given value.

### HasMileageFlag

`func (o *ExpenseTypeInfo) HasMileageFlag() bool`

HasMileageFlag returns a boolean if a field has been set.

### SetMileageFlagNil

`func (o *ExpenseTypeInfo) SetMileageFlagNil(b bool)`

 SetMileageFlagNil sets the value for MileageFlag to be an explicit nil

### UnsetMileageFlag
`func (o *ExpenseTypeInfo) UnsetMileageFlag()`

UnsetMileageFlag ensures that no value is present for MileageFlag, not even an explicit nil
### GetInfo

`func (o *ExpenseTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


