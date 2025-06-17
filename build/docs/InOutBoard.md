# InOutBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**InOutType** | [**InOutTypeReference**](InOutTypeReference.md) |  | 
**AdditionalInfo** | Pointer to **string** |  Max length: 100; | [optional] 
**DateBack** | **time.Time** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewInOutBoard

`func NewInOutBoard(member MemberReference, inOutType InOutTypeReference, dateBack time.Time, ) *InOutBoard`

NewInOutBoard instantiates a new InOutBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInOutBoardWithDefaults

`func NewInOutBoardWithDefaults() *InOutBoard`

NewInOutBoardWithDefaults instantiates a new InOutBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InOutBoard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InOutBoard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InOutBoard) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InOutBoard) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *InOutBoard) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *InOutBoard) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *InOutBoard) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetInOutType

`func (o *InOutBoard) GetInOutType() InOutTypeReference`

GetInOutType returns the InOutType field if non-nil, zero value otherwise.

### GetInOutTypeOk

`func (o *InOutBoard) GetInOutTypeOk() (*InOutTypeReference, bool)`

GetInOutTypeOk returns a tuple with the InOutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInOutType

`func (o *InOutBoard) SetInOutType(v InOutTypeReference)`

SetInOutType sets InOutType field to given value.


### GetAdditionalInfo

`func (o *InOutBoard) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *InOutBoard) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *InOutBoard) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *InOutBoard) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetDateBack

`func (o *InOutBoard) GetDateBack() time.Time`

GetDateBack returns the DateBack field if non-nil, zero value otherwise.

### GetDateBackOk

`func (o *InOutBoard) GetDateBackOk() (*time.Time, bool)`

GetDateBackOk returns a tuple with the DateBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateBack

`func (o *InOutBoard) SetDateBack(v time.Time)`

SetDateBack sets DateBack field to given value.


### GetInfo

`func (o *InOutBoard) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *InOutBoard) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *InOutBoard) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *InOutBoard) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


