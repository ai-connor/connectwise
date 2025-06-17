# ScheduleStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowAsTentativeFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleStatusInfo

`func NewScheduleStatusInfo() *ScheduleStatusInfo`

NewScheduleStatusInfo instantiates a new ScheduleStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleStatusInfoWithDefaults

`func NewScheduleStatusInfoWithDefaults() *ScheduleStatusInfo`

NewScheduleStatusInfoWithDefaults instantiates a new ScheduleStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleStatusInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleStatusInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleStatusInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleStatusInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScheduleStatusInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleStatusInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleStatusInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleStatusInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ScheduleStatusInfo) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ScheduleStatusInfo) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ScheduleStatusInfo) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ScheduleStatusInfo) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ScheduleStatusInfo) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ScheduleStatusInfo) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetShowAsTentativeFlag

`func (o *ScheduleStatusInfo) GetShowAsTentativeFlag() bool`

GetShowAsTentativeFlag returns the ShowAsTentativeFlag field if non-nil, zero value otherwise.

### GetShowAsTentativeFlagOk

`func (o *ScheduleStatusInfo) GetShowAsTentativeFlagOk() (*bool, bool)`

GetShowAsTentativeFlagOk returns a tuple with the ShowAsTentativeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAsTentativeFlag

`func (o *ScheduleStatusInfo) SetShowAsTentativeFlag(v bool)`

SetShowAsTentativeFlag sets ShowAsTentativeFlag field to given value.

### HasShowAsTentativeFlag

`func (o *ScheduleStatusInfo) HasShowAsTentativeFlag() bool`

HasShowAsTentativeFlag returns a boolean if a field has been set.

### SetShowAsTentativeFlagNil

`func (o *ScheduleStatusInfo) SetShowAsTentativeFlagNil(b bool)`

 SetShowAsTentativeFlagNil sets the value for ShowAsTentativeFlag to be an explicit nil

### UnsetShowAsTentativeFlag
`func (o *ScheduleStatusInfo) UnsetShowAsTentativeFlag()`

UnsetShowAsTentativeFlag ensures that no value is present for ShowAsTentativeFlag, not even an explicit nil
### GetInfo

`func (o *ScheduleStatusInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleStatusInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleStatusInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleStatusInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


