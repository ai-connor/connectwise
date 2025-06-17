# BoardSubTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**TypeAssociationIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardSubTypeInfo

`func NewBoardSubTypeInfo() *BoardSubTypeInfo`

NewBoardSubTypeInfo instantiates a new BoardSubTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardSubTypeInfoWithDefaults

`func NewBoardSubTypeInfoWithDefaults() *BoardSubTypeInfo`

NewBoardSubTypeInfoWithDefaults instantiates a new BoardSubTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardSubTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardSubTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardSubTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardSubTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BoardSubTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoardSubTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoardSubTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BoardSubTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *BoardSubTypeInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *BoardSubTypeInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *BoardSubTypeInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *BoardSubTypeInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *BoardSubTypeInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *BoardSubTypeInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetTypeAssociationIds

`func (o *BoardSubTypeInfo) GetTypeAssociationIds() []int32`

GetTypeAssociationIds returns the TypeAssociationIds field if non-nil, zero value otherwise.

### GetTypeAssociationIdsOk

`func (o *BoardSubTypeInfo) GetTypeAssociationIdsOk() (*[]int32, bool)`

GetTypeAssociationIdsOk returns a tuple with the TypeAssociationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeAssociationIds

`func (o *BoardSubTypeInfo) SetTypeAssociationIds(v []int32)`

SetTypeAssociationIds sets TypeAssociationIds field to given value.

### HasTypeAssociationIds

`func (o *BoardSubTypeInfo) HasTypeAssociationIds() bool`

HasTypeAssociationIds returns a boolean if a field has been set.

### GetInfo

`func (o *BoardSubTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardSubTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardSubTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardSubTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


