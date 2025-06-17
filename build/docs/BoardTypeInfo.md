# BoardTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardTypeInfo

`func NewBoardTypeInfo() *BoardTypeInfo`

NewBoardTypeInfo instantiates a new BoardTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardTypeInfoWithDefaults

`func NewBoardTypeInfoWithDefaults() *BoardTypeInfo`

NewBoardTypeInfoWithDefaults instantiates a new BoardTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBoard

`func (o *BoardTypeInfo) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardTypeInfo) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardTypeInfo) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardTypeInfo) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *BoardTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BoardTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BoardTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BoardTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BoardTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BoardTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultFlag

`func (o *BoardTypeInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *BoardTypeInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *BoardTypeInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *BoardTypeInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *BoardTypeInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *BoardTypeInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *BoardTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


