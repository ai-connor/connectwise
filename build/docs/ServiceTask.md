# ServiceTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TicketId** | Pointer to **NullableInt32** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 
**Schedule** | Pointer to [**ScheduleEntryReference**](ScheduleEntryReference.md) |  | [optional] 
**Code** | Pointer to [**ServiceCodeReference**](ServiceCodeReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Resolution** | Pointer to **string** |  | [optional] 
**ChildScheduleAction** | Pointer to **NullableString** |  | [optional] 
**ChildTicketId** | Pointer to **NullableInt32** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceTask

`func NewServiceTask() *ServiceTask`

NewServiceTask instantiates a new ServiceTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTaskWithDefaults

`func NewServiceTaskWithDefaults() *ServiceTask`

NewServiceTaskWithDefaults instantiates a new ServiceTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTicketId

`func (o *ServiceTask) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *ServiceTask) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *ServiceTask) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *ServiceTask) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### SetTicketIdNil

`func (o *ServiceTask) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *ServiceTask) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetNotes

`func (o *ServiceTask) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ServiceTask) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ServiceTask) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ServiceTask) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetClosedFlag

`func (o *ServiceTask) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *ServiceTask) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *ServiceTask) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *ServiceTask) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *ServiceTask) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *ServiceTask) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetPriority

`func (o *ServiceTask) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ServiceTask) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ServiceTask) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ServiceTask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *ServiceTask) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *ServiceTask) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetSchedule

`func (o *ServiceTask) GetSchedule() ScheduleEntryReference`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ServiceTask) GetScheduleOk() (*ScheduleEntryReference, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ServiceTask) SetSchedule(v ScheduleEntryReference)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ServiceTask) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCode

`func (o *ServiceTask) GetCode() ServiceCodeReference`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceTask) GetCodeOk() (*ServiceCodeReference, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceTask) SetCode(v ServiceCodeReference)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMember

`func (o *ServiceTask) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ServiceTask) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ServiceTask) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ServiceTask) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetResolution

`func (o *ServiceTask) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *ServiceTask) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *ServiceTask) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *ServiceTask) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetChildScheduleAction

`func (o *ServiceTask) GetChildScheduleAction() string`

GetChildScheduleAction returns the ChildScheduleAction field if non-nil, zero value otherwise.

### GetChildScheduleActionOk

`func (o *ServiceTask) GetChildScheduleActionOk() (*string, bool)`

GetChildScheduleActionOk returns a tuple with the ChildScheduleAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildScheduleAction

`func (o *ServiceTask) SetChildScheduleAction(v string)`

SetChildScheduleAction sets ChildScheduleAction field to given value.

### HasChildScheduleAction

`func (o *ServiceTask) HasChildScheduleAction() bool`

HasChildScheduleAction returns a boolean if a field has been set.

### SetChildScheduleActionNil

`func (o *ServiceTask) SetChildScheduleActionNil(b bool)`

 SetChildScheduleActionNil sets the value for ChildScheduleAction to be an explicit nil

### UnsetChildScheduleAction
`func (o *ServiceTask) UnsetChildScheduleAction()`

UnsetChildScheduleAction ensures that no value is present for ChildScheduleAction, not even an explicit nil
### GetChildTicketId

`func (o *ServiceTask) GetChildTicketId() int32`

GetChildTicketId returns the ChildTicketId field if non-nil, zero value otherwise.

### GetChildTicketIdOk

`func (o *ServiceTask) GetChildTicketIdOk() (*int32, bool)`

GetChildTicketIdOk returns a tuple with the ChildTicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildTicketId

`func (o *ServiceTask) SetChildTicketId(v int32)`

SetChildTicketId sets ChildTicketId field to given value.

### HasChildTicketId

`func (o *ServiceTask) HasChildTicketId() bool`

HasChildTicketId returns a boolean if a field has been set.

### SetChildTicketIdNil

`func (o *ServiceTask) SetChildTicketIdNil(b bool)`

 SetChildTicketIdNil sets the value for ChildTicketId to be an explicit nil

### UnsetChildTicketId
`func (o *ServiceTask) UnsetChildTicketId()`

UnsetChildTicketId ensures that no value is present for ChildTicketId, not even an explicit nil
### GetSummary

`func (o *ServiceTask) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ServiceTask) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ServiceTask) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ServiceTask) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetInfo

`func (o *ServiceTask) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceTask) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceTask) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceTask) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


