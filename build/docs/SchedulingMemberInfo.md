# SchedulingMemberInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**MiddleInitial** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**FullName** | Pointer to **string** |  | [optional] 
**DefaultEmail** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSchedulingMemberInfo

`func NewSchedulingMemberInfo() *SchedulingMemberInfo`

NewSchedulingMemberInfo instantiates a new SchedulingMemberInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulingMemberInfoWithDefaults

`func NewSchedulingMemberInfoWithDefaults() *SchedulingMemberInfo`

NewSchedulingMemberInfoWithDefaults instantiates a new SchedulingMemberInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SchedulingMemberInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchedulingMemberInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchedulingMemberInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SchedulingMemberInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *SchedulingMemberInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SchedulingMemberInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SchedulingMemberInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SchedulingMemberInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetFirstName

`func (o *SchedulingMemberInfo) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SchedulingMemberInfo) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SchedulingMemberInfo) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SchedulingMemberInfo) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetMiddleInitial

`func (o *SchedulingMemberInfo) GetMiddleInitial() string`

GetMiddleInitial returns the MiddleInitial field if non-nil, zero value otherwise.

### GetMiddleInitialOk

`func (o *SchedulingMemberInfo) GetMiddleInitialOk() (*string, bool)`

GetMiddleInitialOk returns a tuple with the MiddleInitial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleInitial

`func (o *SchedulingMemberInfo) SetMiddleInitial(v string)`

SetMiddleInitial sets MiddleInitial field to given value.

### HasMiddleInitial

`func (o *SchedulingMemberInfo) HasMiddleInitial() bool`

HasMiddleInitial returns a boolean if a field has been set.

### GetLastName

`func (o *SchedulingMemberInfo) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SchedulingMemberInfo) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SchedulingMemberInfo) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SchedulingMemberInfo) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetFullName

`func (o *SchedulingMemberInfo) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *SchedulingMemberInfo) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *SchedulingMemberInfo) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *SchedulingMemberInfo) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDefaultEmail

`func (o *SchedulingMemberInfo) GetDefaultEmail() string`

GetDefaultEmail returns the DefaultEmail field if non-nil, zero value otherwise.

### GetDefaultEmailOk

`func (o *SchedulingMemberInfo) GetDefaultEmailOk() (*string, bool)`

GetDefaultEmailOk returns a tuple with the DefaultEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEmail

`func (o *SchedulingMemberInfo) SetDefaultEmail(v string)`

SetDefaultEmail sets DefaultEmail field to given value.

### HasDefaultEmail

`func (o *SchedulingMemberInfo) HasDefaultEmail() bool`

HasDefaultEmail returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *SchedulingMemberInfo) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *SchedulingMemberInfo) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *SchedulingMemberInfo) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *SchedulingMemberInfo) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *SchedulingMemberInfo) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *SchedulingMemberInfo) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInfo

`func (o *SchedulingMemberInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SchedulingMemberInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SchedulingMemberInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SchedulingMemberInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


