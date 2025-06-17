# ContactTrack

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TrackId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**ActionTaken** | Pointer to **NullableInt32** |  | [optional] 
**ActionRemaining** | Pointer to **NullableInt32** |  | [optional] 
**StartedBy** | Pointer to **string** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewContactTrack

`func NewContactTrack() *ContactTrack`

NewContactTrack instantiates a new ContactTrack object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactTrackWithDefaults

`func NewContactTrackWithDefaults() *ContactTrack`

NewContactTrackWithDefaults instantiates a new ContactTrack object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContactTrack) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContactTrack) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContactTrack) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ContactTrack) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTrackId

`func (o *ContactTrack) GetTrackId() int32`

GetTrackId returns the TrackId field if non-nil, zero value otherwise.

### GetTrackIdOk

`func (o *ContactTrack) GetTrackIdOk() (*int32, bool)`

GetTrackIdOk returns a tuple with the TrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackId

`func (o *ContactTrack) SetTrackId(v int32)`

SetTrackId sets TrackId field to given value.

### HasTrackId

`func (o *ContactTrack) HasTrackId() bool`

HasTrackId returns a boolean if a field has been set.

### GetName

`func (o *ContactTrack) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContactTrack) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContactTrack) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContactTrack) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartDate

`func (o *ContactTrack) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ContactTrack) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ContactTrack) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ContactTrack) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ContactTrack) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ContactTrack) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ContactTrack) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ContactTrack) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetActionTaken

`func (o *ContactTrack) GetActionTaken() int32`

GetActionTaken returns the ActionTaken field if non-nil, zero value otherwise.

### GetActionTakenOk

`func (o *ContactTrack) GetActionTakenOk() (*int32, bool)`

GetActionTakenOk returns a tuple with the ActionTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTaken

`func (o *ContactTrack) SetActionTaken(v int32)`

SetActionTaken sets ActionTaken field to given value.

### HasActionTaken

`func (o *ContactTrack) HasActionTaken() bool`

HasActionTaken returns a boolean if a field has been set.

### SetActionTakenNil

`func (o *ContactTrack) SetActionTakenNil(b bool)`

 SetActionTakenNil sets the value for ActionTaken to be an explicit nil

### UnsetActionTaken
`func (o *ContactTrack) UnsetActionTaken()`

UnsetActionTaken ensures that no value is present for ActionTaken, not even an explicit nil
### GetActionRemaining

`func (o *ContactTrack) GetActionRemaining() int32`

GetActionRemaining returns the ActionRemaining field if non-nil, zero value otherwise.

### GetActionRemainingOk

`func (o *ContactTrack) GetActionRemainingOk() (*int32, bool)`

GetActionRemainingOk returns a tuple with the ActionRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionRemaining

`func (o *ContactTrack) SetActionRemaining(v int32)`

SetActionRemaining sets ActionRemaining field to given value.

### HasActionRemaining

`func (o *ContactTrack) HasActionRemaining() bool`

HasActionRemaining returns a boolean if a field has been set.

### SetActionRemainingNil

`func (o *ContactTrack) SetActionRemainingNil(b bool)`

 SetActionRemainingNil sets the value for ActionRemaining to be an explicit nil

### UnsetActionRemaining
`func (o *ContactTrack) UnsetActionRemaining()`

UnsetActionRemaining ensures that no value is present for ActionRemaining, not even an explicit nil
### GetStartedBy

`func (o *ContactTrack) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *ContactTrack) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *ContactTrack) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.

### HasStartedBy

`func (o *ContactTrack) HasStartedBy() bool`

HasStartedBy returns a boolean if a field has been set.

### GetCompany

`func (o *ContactTrack) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ContactTrack) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ContactTrack) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ContactTrack) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetContact

`func (o *ContactTrack) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ContactTrack) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ContactTrack) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ContactTrack) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetInfo

`func (o *ContactTrack) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ContactTrack) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ContactTrack) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ContactTrack) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


