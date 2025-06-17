# ScheduleStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowAsTentativeFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleStatus

`func NewScheduleStatus(name string, ) *ScheduleStatus`

NewScheduleStatus instantiates a new ScheduleStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleStatusWithDefaults

`func NewScheduleStatusWithDefaults() *ScheduleStatus`

NewScheduleStatusWithDefaults instantiates a new ScheduleStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScheduleStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *ScheduleStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ScheduleStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ScheduleStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ScheduleStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ScheduleStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ScheduleStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetShowAsTentativeFlag

`func (o *ScheduleStatus) GetShowAsTentativeFlag() bool`

GetShowAsTentativeFlag returns the ShowAsTentativeFlag field if non-nil, zero value otherwise.

### GetShowAsTentativeFlagOk

`func (o *ScheduleStatus) GetShowAsTentativeFlagOk() (*bool, bool)`

GetShowAsTentativeFlagOk returns a tuple with the ShowAsTentativeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsTentativeFlag

`func (o *ScheduleStatus) SetShowAsTentativeFlag(v bool)`

SetShowAsTentativeFlag sets ShowAsTentativeFlag field to given value.

### HasShowAsTentativeFlag

`func (o *ScheduleStatus) HasShowAsTentativeFlag() bool`

HasShowAsTentativeFlag returns a boolean if a field has been set.

### SetShowAsTentativeFlagNil

`func (o *ScheduleStatus) SetShowAsTentativeFlagNil(b bool)`

 SetShowAsTentativeFlagNil sets the value for ShowAsTentativeFlag to be an explicit nil

### UnsetShowAsTentativeFlag
`func (o *ScheduleStatus) UnsetShowAsTentativeFlag()`

UnsetShowAsTentativeFlag ensures that no value is present for ShowAsTentativeFlag, not even an explicit nil
### GetInfo

`func (o *ScheduleStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


