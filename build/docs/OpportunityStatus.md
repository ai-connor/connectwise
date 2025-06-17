# OpportunityStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**WonFlag** | Pointer to **NullableBool** |  | [optional] 
**LostFlag** | Pointer to **NullableBool** |  | [optional] 
**ClosedFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**EnteredBy** | Pointer to **string** |  | [optional] 
**DateEntered** | Pointer to **time.Time** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOpportunityStatus

`func NewOpportunityStatus(name string, ) *OpportunityStatus`

NewOpportunityStatus instantiates a new OpportunityStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpportunityStatusWithDefaults

`func NewOpportunityStatusWithDefaults() *OpportunityStatus`

NewOpportunityStatusWithDefaults instantiates a new OpportunityStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpportunityStatus) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpportunityStatus) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpportunityStatus) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OpportunityStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OpportunityStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpportunityStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpportunityStatus) SetName(v string)`

SetName sets Name field to given value.


### GetWonFlag

`func (o *OpportunityStatus) GetWonFlag() bool`

GetWonFlag returns the WonFlag field if non-nil, zero value otherwise.

### GetWonFlagOk

`func (o *OpportunityStatus) GetWonFlagOk() (*bool, bool)`

GetWonFlagOk returns a tuple with the WonFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWonFlag

`func (o *OpportunityStatus) SetWonFlag(v bool)`

SetWonFlag sets WonFlag field to given value.

### HasWonFlag

`func (o *OpportunityStatus) HasWonFlag() bool`

HasWonFlag returns a boolean if a field has been set.

### SetWonFlagNil

`func (o *OpportunityStatus) SetWonFlagNil(b bool)`

 SetWonFlagNil sets the value for WonFlag to be an explicit nil

### UnsetWonFlag
`func (o *OpportunityStatus) UnsetWonFlag()`

UnsetWonFlag ensures that no value is present for WonFlag, not even an explicit nil
### GetLostFlag

`func (o *OpportunityStatus) GetLostFlag() bool`

GetLostFlag returns the LostFlag field if non-nil, zero value otherwise.

### GetLostFlagOk

`func (o *OpportunityStatus) GetLostFlagOk() (*bool, bool)`

GetLostFlagOk returns a tuple with the LostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLostFlag

`func (o *OpportunityStatus) SetLostFlag(v bool)`

SetLostFlag sets LostFlag field to given value.

### HasLostFlag

`func (o *OpportunityStatus) HasLostFlag() bool`

HasLostFlag returns a boolean if a field has been set.

### SetLostFlagNil

`func (o *OpportunityStatus) SetLostFlagNil(b bool)`

 SetLostFlagNil sets the value for LostFlag to be an explicit nil

### UnsetLostFlag
`func (o *OpportunityStatus) UnsetLostFlag()`

UnsetLostFlag ensures that no value is present for LostFlag, not even an explicit nil
### GetClosedFlag

`func (o *OpportunityStatus) GetClosedFlag() bool`

GetClosedFlag returns the ClosedFlag field if non-nil, zero value otherwise.

### GetClosedFlagOk

`func (o *OpportunityStatus) GetClosedFlagOk() (*bool, bool)`

GetClosedFlagOk returns a tuple with the ClosedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedFlag

`func (o *OpportunityStatus) SetClosedFlag(v bool)`

SetClosedFlag sets ClosedFlag field to given value.

### HasClosedFlag

`func (o *OpportunityStatus) HasClosedFlag() bool`

HasClosedFlag returns a boolean if a field has been set.

### SetClosedFlagNil

`func (o *OpportunityStatus) SetClosedFlagNil(b bool)`

 SetClosedFlagNil sets the value for ClosedFlag to be an explicit nil

### UnsetClosedFlag
`func (o *OpportunityStatus) UnsetClosedFlag()`

UnsetClosedFlag ensures that no value is present for ClosedFlag, not even an explicit nil
### GetInactiveFlag

`func (o *OpportunityStatus) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *OpportunityStatus) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *OpportunityStatus) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *OpportunityStatus) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *OpportunityStatus) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *OpportunityStatus) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetDefaultFlag

`func (o *OpportunityStatus) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *OpportunityStatus) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *OpportunityStatus) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *OpportunityStatus) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *OpportunityStatus) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *OpportunityStatus) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetEnteredBy

`func (o *OpportunityStatus) GetEnteredBy() string`

GetEnteredBy returns the EnteredBy field if non-nil, zero value otherwise.

### GetEnteredByOk

`func (o *OpportunityStatus) GetEnteredByOk() (*string, bool)`

GetEnteredByOk returns a tuple with the EnteredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredBy

`func (o *OpportunityStatus) SetEnteredBy(v string)`

SetEnteredBy sets EnteredBy field to given value.

### HasEnteredBy

`func (o *OpportunityStatus) HasEnteredBy() bool`

HasEnteredBy returns a boolean if a field has been set.

### GetDateEntered

`func (o *OpportunityStatus) GetDateEntered() time.Time`

GetDateEntered returns the DateEntered field if non-nil, zero value otherwise.

### GetDateEnteredOk

`func (o *OpportunityStatus) GetDateEnteredOk() (*time.Time, bool)`

GetDateEnteredOk returns a tuple with the DateEntered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEntered

`func (o *OpportunityStatus) SetDateEntered(v time.Time)`

SetDateEntered sets DateEntered field to given value.

### HasDateEntered

`func (o *OpportunityStatus) HasDateEntered() bool`

HasDateEntered returns a boolean if a field has been set.

### GetInfo

`func (o *OpportunityStatus) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *OpportunityStatus) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *OpportunityStatus) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *OpportunityStatus) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


