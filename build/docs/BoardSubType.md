# BoardSubType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**TypeAssociationIds** | Pointer to **[]int32** |  | [optional] 
**AddAllTypesFlag** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllTypesFlag** | Pointer to **NullableBool** |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardSubType

`func NewBoardSubType(name string, ) *BoardSubType`

NewBoardSubType instantiates a new BoardSubType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardSubTypeWithDefaults

`func NewBoardSubTypeWithDefaults() *BoardSubType`

NewBoardSubTypeWithDefaults instantiates a new BoardSubType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardSubType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardSubType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardSubType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardSubType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardSubType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardSubType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardSubType) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *BoardSubType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BoardSubType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BoardSubType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BoardSubType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BoardSubType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BoardSubType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetTypeAssociationIds

`func (o *BoardSubType) GetTypeAssociationIds() []int32`

GetTypeAssociationIds returns the TypeAssociationIds field if non-nil, zero value otherwise.

### GetTypeAssociationIdsOk

`func (o *BoardSubType) GetTypeAssociationIdsOk() (*[]int32, bool)`

GetTypeAssociationIdsOk returns a tuple with the TypeAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAssociationIds

`func (o *BoardSubType) SetTypeAssociationIds(v []int32)`

SetTypeAssociationIds sets TypeAssociationIds field to given value.

### HasTypeAssociationIds

`func (o *BoardSubType) HasTypeAssociationIds() bool`

HasTypeAssociationIds returns a boolean if a field has been set.

### GetAddAllTypesFlag

`func (o *BoardSubType) GetAddAllTypesFlag() bool`

GetAddAllTypesFlag returns the AddAllTypesFlag field if non-nil, zero value otherwise.

### GetAddAllTypesFlagOk

`func (o *BoardSubType) GetAddAllTypesFlagOk() (*bool, bool)`

GetAddAllTypesFlagOk returns a tuple with the AddAllTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllTypesFlag

`func (o *BoardSubType) SetAddAllTypesFlag(v bool)`

SetAddAllTypesFlag sets AddAllTypesFlag field to given value.

### HasAddAllTypesFlag

`func (o *BoardSubType) HasAddAllTypesFlag() bool`

HasAddAllTypesFlag returns a boolean if a field has been set.

### SetAddAllTypesFlagNil

`func (o *BoardSubType) SetAddAllTypesFlagNil(b bool)`

 SetAddAllTypesFlagNil sets the value for AddAllTypesFlag to be an explicit nil

### UnsetAddAllTypesFlag
`func (o *BoardSubType) UnsetAddAllTypesFlag()`

UnsetAddAllTypesFlag ensures that no value is present for AddAllTypesFlag, not even an explicit nil
### GetRemoveAllTypesFlag

`func (o *BoardSubType) GetRemoveAllTypesFlag() bool`

GetRemoveAllTypesFlag returns the RemoveAllTypesFlag field if non-nil, zero value otherwise.

### GetRemoveAllTypesFlagOk

`func (o *BoardSubType) GetRemoveAllTypesFlagOk() (*bool, bool)`

GetRemoveAllTypesFlagOk returns a tuple with the RemoveAllTypesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllTypesFlag

`func (o *BoardSubType) SetRemoveAllTypesFlag(v bool)`

SetRemoveAllTypesFlag sets RemoveAllTypesFlag field to given value.

### HasRemoveAllTypesFlag

`func (o *BoardSubType) HasRemoveAllTypesFlag() bool`

HasRemoveAllTypesFlag returns a boolean if a field has been set.

### SetRemoveAllTypesFlagNil

`func (o *BoardSubType) SetRemoveAllTypesFlagNil(b bool)`

 SetRemoveAllTypesFlagNil sets the value for RemoveAllTypesFlag to be an explicit nil

### UnsetRemoveAllTypesFlag
`func (o *BoardSubType) UnsetRemoveAllTypesFlag()`

UnsetRemoveAllTypesFlag ensures that no value is present for RemoveAllTypesFlag, not even an explicit nil
### GetBoard

`func (o *BoardSubType) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *BoardSubType) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *BoardSubType) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *BoardSubType) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetInfo

`func (o *BoardSubType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardSubType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardSubType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardSubType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


