# ActivityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Points** | Pointer to **NullableInt32** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailFlag** | Pointer to **NullableBool** |  | [optional] 
**MemoFlag** | Pointer to **NullableBool** |  | [optional] 
**HistoryFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewActivityType

`func NewActivityType(name string, ) *ActivityType`

NewActivityType instantiates a new ActivityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityTypeWithDefaults

`func NewActivityTypeWithDefaults() *ActivityType`

NewActivityTypeWithDefaults instantiates a new ActivityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityType) SetName(v string)`

SetName sets Name field to given value.


### GetPoints

`func (o *ActivityType) GetPoints() int32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *ActivityType) GetPointsOk() (*int32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *ActivityType) SetPoints(v int32)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *ActivityType) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### SetPointsNil

`func (o *ActivityType) SetPointsNil(b bool)`

 SetPointsNil sets the value for Points to be an explicit nil

### UnsetPoints
`func (o *ActivityType) UnsetPoints()`

UnsetPoints ensures that no value is present for Points, not even an explicit nil
### GetDefaultFlag

`func (o *ActivityType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ActivityType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ActivityType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ActivityType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ActivityType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ActivityType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *ActivityType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ActivityType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ActivityType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ActivityType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ActivityType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ActivityType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetEmailFlag

`func (o *ActivityType) GetEmailFlag() bool`

GetEmailFlag returns the EmailFlag field if non-nil, zero value otherwise.

### GetEmailFlagOk

`func (o *ActivityType) GetEmailFlagOk() (*bool, bool)`

GetEmailFlagOk returns a tuple with the EmailFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailFlag

`func (o *ActivityType) SetEmailFlag(v bool)`

SetEmailFlag sets EmailFlag field to given value.

### HasEmailFlag

`func (o *ActivityType) HasEmailFlag() bool`

HasEmailFlag returns a boolean if a field has been set.

### SetEmailFlagNil

`func (o *ActivityType) SetEmailFlagNil(b bool)`

 SetEmailFlagNil sets the value for EmailFlag to be an explicit nil

### UnsetEmailFlag
`func (o *ActivityType) UnsetEmailFlag()`

UnsetEmailFlag ensures that no value is present for EmailFlag, not even an explicit nil
### GetMemoFlag

`func (o *ActivityType) GetMemoFlag() bool`

GetMemoFlag returns the MemoFlag field if non-nil, zero value otherwise.

### GetMemoFlagOk

`func (o *ActivityType) GetMemoFlagOk() (*bool, bool)`

GetMemoFlagOk returns a tuple with the MemoFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoFlag

`func (o *ActivityType) SetMemoFlag(v bool)`

SetMemoFlag sets MemoFlag field to given value.

### HasMemoFlag

`func (o *ActivityType) HasMemoFlag() bool`

HasMemoFlag returns a boolean if a field has been set.

### SetMemoFlagNil

`func (o *ActivityType) SetMemoFlagNil(b bool)`

 SetMemoFlagNil sets the value for MemoFlag to be an explicit nil

### UnsetMemoFlag
`func (o *ActivityType) UnsetMemoFlag()`

UnsetMemoFlag ensures that no value is present for MemoFlag, not even an explicit nil
### GetHistoryFlag

`func (o *ActivityType) GetHistoryFlag() bool`

GetHistoryFlag returns the HistoryFlag field if non-nil, zero value otherwise.

### GetHistoryFlagOk

`func (o *ActivityType) GetHistoryFlagOk() (*bool, bool)`

GetHistoryFlagOk returns a tuple with the HistoryFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistoryFlag

`func (o *ActivityType) SetHistoryFlag(v bool)`

SetHistoryFlag sets HistoryFlag field to given value.

### HasHistoryFlag

`func (o *ActivityType) HasHistoryFlag() bool`

HasHistoryFlag returns a boolean if a field has been set.

### SetHistoryFlagNil

`func (o *ActivityType) SetHistoryFlagNil(b bool)`

 SetHistoryFlagNil sets the value for HistoryFlag to be an explicit nil

### UnsetHistoryFlag
`func (o *ActivityType) UnsetHistoryFlag()`

UnsetHistoryFlag ensures that no value is present for HistoryFlag, not even an explicit nil
### GetInfo

`func (o *ActivityType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ActivityType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ActivityType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ActivityType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


