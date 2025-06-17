# BoardExcludedMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**MemberId** | Pointer to **int32** |  | [optional] 
**BoardId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBoardExcludedMember

`func NewBoardExcludedMember() *BoardExcludedMember`

NewBoardExcludedMember instantiates a new BoardExcludedMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardExcludedMemberWithDefaults

`func NewBoardExcludedMemberWithDefaults() *BoardExcludedMember`

NewBoardExcludedMemberWithDefaults instantiates a new BoardExcludedMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BoardExcludedMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardExcludedMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardExcludedMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BoardExcludedMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMemberId

`func (o *BoardExcludedMember) GetMemberId() int32`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *BoardExcludedMember) GetMemberIdOk() (*int32, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *BoardExcludedMember) SetMemberId(v int32)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *BoardExcludedMember) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### GetBoardId

`func (o *BoardExcludedMember) GetBoardId() int32`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *BoardExcludedMember) GetBoardIdOk() (*int32, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *BoardExcludedMember) SetBoardId(v int32)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *BoardExcludedMember) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### SetBoardIdNil

`func (o *BoardExcludedMember) SetBoardIdNil(b bool)`

 SetBoardIdNil sets the value for BoardId to be an explicit nil

### UnsetBoardId
`func (o *BoardExcludedMember) UnsetBoardId()`

UnsetBoardId ensures that no value is present for BoardId, not even an explicit nil
### GetInfo

`func (o *BoardExcludedMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BoardExcludedMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BoardExcludedMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BoardExcludedMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


