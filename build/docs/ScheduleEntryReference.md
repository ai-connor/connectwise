# ScheduleEntryReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleEntryReference

`func NewScheduleEntryReference() *ScheduleEntryReference`

NewScheduleEntryReference instantiates a new ScheduleEntryReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleEntryReferenceWithDefaults

`func NewScheduleEntryReferenceWithDefaults() *ScheduleEntryReference`

NewScheduleEntryReferenceWithDefaults instantiates a new ScheduleEntryReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleEntryReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleEntryReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleEntryReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleEntryReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *ScheduleEntryReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ScheduleEntryReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDescription

`func (o *ScheduleEntryReference) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ScheduleEntryReference) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ScheduleEntryReference) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ScheduleEntryReference) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInfo

`func (o *ScheduleEntryReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleEntryReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleEntryReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleEntryReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


