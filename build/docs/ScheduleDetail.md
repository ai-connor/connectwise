# ScheduleDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScheduleEntry** | Pointer to [**ScheduleEntryReference**](ScheduleEntryReference.md) |  | [optional] 
**DateStart** | Pointer to **string** |  | [optional] 
**DateEnd** | Pointer to **string** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleDetail

`func NewScheduleDetail() *ScheduleDetail`

NewScheduleDetail instantiates a new ScheduleDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleDetailWithDefaults

`func NewScheduleDetailWithDefaults() *ScheduleDetail`

NewScheduleDetailWithDefaults instantiates a new ScheduleDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleDetail) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleDetail) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleDetail) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduleEntry

`func (o *ScheduleDetail) GetScheduleEntry() ScheduleEntryReference`

GetScheduleEntry returns the ScheduleEntry field if non-nil, zero value otherwise.

### GetScheduleEntryOk

`func (o *ScheduleDetail) GetScheduleEntryOk() (*ScheduleEntryReference, bool)`

GetScheduleEntryOk returns a tuple with the ScheduleEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEntry

`func (o *ScheduleDetail) SetScheduleEntry(v ScheduleEntryReference)`

SetScheduleEntry sets ScheduleEntry field to given value.

### HasScheduleEntry

`func (o *ScheduleDetail) HasScheduleEntry() bool`

HasScheduleEntry returns a boolean if a field has been set.

### GetDateStart

`func (o *ScheduleDetail) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *ScheduleDetail) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *ScheduleDetail) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *ScheduleDetail) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *ScheduleDetail) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *ScheduleDetail) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *ScheduleDetail) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *ScheduleDetail) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetMember

`func (o *ScheduleDetail) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ScheduleDetail) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ScheduleDetail) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ScheduleDetail) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *ScheduleDetail) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleDetail) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleDetail) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleDetail) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


