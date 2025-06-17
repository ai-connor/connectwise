# CompanyStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**NotifyFlag** | Pointer to **NullableBool** |  | [optional] 
**DisallowSavingFlag** | Pointer to **NullableBool** |  | [optional] 
**NotificationMessage** | Pointer to **string** |  Max length: 500; | [optional] 
**CustomNoteFlag** | Pointer to **NullableBool** |  | [optional] 
**CancelOpenTracksFlag** | Pointer to **NullableBool** |  | [optional] 
**Track** | Pointer to [**TrackReference**](TrackReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyStatus

`func NewCompanyStatus(name string, ) *CompanyStatus`

NewCompanyStatus instantiates a new CompanyStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyStatusWithDefaults

`func NewCompanyStatusWithDefaults() *CompanyStatus`

NewCompanyStatusWithDefaults instantiates a new CompanyStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanyStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyStatus) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultFlag

`func (o *CompanyStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *CompanyStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *CompanyStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *CompanyStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *CompanyStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *CompanyStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInactiveFlag

`func (o *CompanyStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CompanyStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CompanyStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CompanyStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CompanyStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CompanyStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetNotifyFlag

`func (o *CompanyStatus) GetNotifyFlag() bool`

GetNotifyFlag returns the NotifyFlag field if non-nil, zero value otherwise.

### GetNotifyFlagOk

`func (o *CompanyStatus) GetNotifyFlagOk() (*bool, bool)`

GetNotifyFlagOk returns a tuple with the NotifyFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyFlag

`func (o *CompanyStatus) SetNotifyFlag(v bool)`

SetNotifyFlag sets NotifyFlag field to given value.

### HasNotifyFlag

`func (o *CompanyStatus) HasNotifyFlag() bool`

HasNotifyFlag returns a boolean if a field has been set.

### SetNotifyFlagNil

`func (o *CompanyStatus) SetNotifyFlagNil(b bool)`

 SetNotifyFlagNil sets the value for NotifyFlag to be an explicit nil

### UnsetNotifyFlag
`func (o *CompanyStatus) UnsetNotifyFlag()`

UnsetNotifyFlag ensures that no value is present for NotifyFlag, not even an explicit nil
### GetDisallowSavingFlag

`func (o *CompanyStatus) GetDisallowSavingFlag() bool`

GetDisallowSavingFlag returns the DisallowSavingFlag field if non-nil, zero value otherwise.

### GetDisallowSavingFlagOk

`func (o *CompanyStatus) GetDisallowSavingFlagOk() (*bool, bool)`

GetDisallowSavingFlagOk returns a tuple with the DisallowSavingFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisallowSavingFlag

`func (o *CompanyStatus) SetDisallowSavingFlag(v bool)`

SetDisallowSavingFlag sets DisallowSavingFlag field to given value.

### HasDisallowSavingFlag

`func (o *CompanyStatus) HasDisallowSavingFlag() bool`

HasDisallowSavingFlag returns a boolean if a field has been set.

### SetDisallowSavingFlagNil

`func (o *CompanyStatus) SetDisallowSavingFlagNil(b bool)`

 SetDisallowSavingFlagNil sets the value for DisallowSavingFlag to be an explicit nil

### UnsetDisallowSavingFlag
`func (o *CompanyStatus) UnsetDisallowSavingFlag()`

UnsetDisallowSavingFlag ensures that no value is present for DisallowSavingFlag, not even an explicit nil
### GetNotificationMessage

`func (o *CompanyStatus) GetNotificationMessage() string`

GetNotificationMessage returns the NotificationMessage field if non-nil, zero value otherwise.

### GetNotificationMessageOk

`func (o *CompanyStatus) GetNotificationMessageOk() (*string, bool)`

GetNotificationMessageOk returns a tuple with the NotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessage

`func (o *CompanyStatus) SetNotificationMessage(v string)`

SetNotificationMessage sets NotificationMessage field to given value.

### HasNotificationMessage

`func (o *CompanyStatus) HasNotificationMessage() bool`

HasNotificationMessage returns a boolean if a field has been set.

### GetCustomNoteFlag

`func (o *CompanyStatus) GetCustomNoteFlag() bool`

GetCustomNoteFlag returns the CustomNoteFlag field if non-nil, zero value otherwise.

### GetCustomNoteFlagOk

`func (o *CompanyStatus) GetCustomNoteFlagOk() (*bool, bool)`

GetCustomNoteFlagOk returns a tuple with the CustomNoteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNoteFlag

`func (o *CompanyStatus) SetCustomNoteFlag(v bool)`

SetCustomNoteFlag sets CustomNoteFlag field to given value.

### HasCustomNoteFlag

`func (o *CompanyStatus) HasCustomNoteFlag() bool`

HasCustomNoteFlag returns a boolean if a field has been set.

### SetCustomNoteFlagNil

`func (o *CompanyStatus) SetCustomNoteFlagNil(b bool)`

 SetCustomNoteFlagNil sets the value for CustomNoteFlag to be an explicit nil

### UnsetCustomNoteFlag
`func (o *CompanyStatus) UnsetCustomNoteFlag()`

UnsetCustomNoteFlag ensures that no value is present for CustomNoteFlag, not even an explicit nil
### GetCancelOpenTracksFlag

`func (o *CompanyStatus) GetCancelOpenTracksFlag() bool`

GetCancelOpenTracksFlag returns the CancelOpenTracksFlag field if non-nil, zero value otherwise.

### GetCancelOpenTracksFlagOk

`func (o *CompanyStatus) GetCancelOpenTracksFlagOk() (*bool, bool)`

GetCancelOpenTracksFlagOk returns a tuple with the CancelOpenTracksFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelOpenTracksFlag

`func (o *CompanyStatus) SetCancelOpenTracksFlag(v bool)`

SetCancelOpenTracksFlag sets CancelOpenTracksFlag field to given value.

### HasCancelOpenTracksFlag

`func (o *CompanyStatus) HasCancelOpenTracksFlag() bool`

HasCancelOpenTracksFlag returns a boolean if a field has been set.

### SetCancelOpenTracksFlagNil

`func (o *CompanyStatus) SetCancelOpenTracksFlagNil(b bool)`

 SetCancelOpenTracksFlagNil sets the value for CancelOpenTracksFlag to be an explicit nil

### UnsetCancelOpenTracksFlag
`func (o *CompanyStatus) UnsetCancelOpenTracksFlag()`

UnsetCancelOpenTracksFlag ensures that no value is present for CancelOpenTracksFlag, not even an explicit nil
### GetTrack

`func (o *CompanyStatus) GetTrack() TrackReference`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *CompanyStatus) GetTrackOk() (*TrackReference, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *CompanyStatus) SetTrack(v TrackReference)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *CompanyStatus) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


