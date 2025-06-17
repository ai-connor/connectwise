# Track

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**NotifyActionIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTrack

`func NewTrack(name string, ) *Track`

NewTrack instantiates a new Track object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackWithDefaults

`func NewTrackWithDefaults() *Track`

NewTrackWithDefaults instantiates a new Track object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Track) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Track) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Track) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Track) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Track) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Track) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Track) SetName(v string)`

SetName sets Name field to given value.


### GetInactiveFlag

`func (o *Track) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Track) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Track) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Track) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Track) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Track) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetNotifyActionIds

`func (o *Track) GetNotifyActionIds() []int32`

GetNotifyActionIds returns the NotifyActionIds field if non-nil, zero value otherwise.

### GetNotifyActionIdsOk

`func (o *Track) GetNotifyActionIdsOk() (*[]int32, bool)`

GetNotifyActionIdsOk returns a tuple with the NotifyActionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyActionIds

`func (o *Track) SetNotifyActionIds(v []int32)`

SetNotifyActionIds sets NotifyActionIds field to given value.

### HasNotifyActionIds

`func (o *Track) HasNotifyActionIds() bool`

HasNotifyActionIds returns a boolean if a field has been set.

### GetInfo

`func (o *Track) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Track) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Track) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Track) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


