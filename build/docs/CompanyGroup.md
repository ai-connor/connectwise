# CompanyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  Required On Updates; | [optional] 
**Group** | [**GroupReference**](GroupReference.md) |  | 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**DefaultContactFlag** | Pointer to **NullableBool** |  | [optional] 
**AllContactsFlag** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllContactsFlag** | Pointer to **NullableBool** |  | [optional] 
**UnsubscribeFlag** | Pointer to **NullableBool** |  | [optional] 
**ContactIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyGroup

`func NewCompanyGroup(group GroupReference, ) *CompanyGroup`

NewCompanyGroup instantiates a new CompanyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyGroupWithDefaults

`func NewCompanyGroupWithDefaults() *CompanyGroup`

NewCompanyGroupWithDefaults instantiates a new CompanyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroup

`func (o *CompanyGroup) GetGroup() GroupReference`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CompanyGroup) GetGroupOk() (*GroupReference, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CompanyGroup) SetGroup(v GroupReference)`

SetGroup sets Group field to given value.


### GetCompany

`func (o *CompanyGroup) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyGroup) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyGroup) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyGroup) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetDefaultContactFlag

`func (o *CompanyGroup) GetDefaultContactFlag() bool`

GetDefaultContactFlag returns the DefaultContactFlag field if non-nil, zero value otherwise.

### GetDefaultContactFlagOk

`func (o *CompanyGroup) GetDefaultContactFlagOk() (*bool, bool)`

GetDefaultContactFlagOk returns a tuple with the DefaultContactFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContactFlag

`func (o *CompanyGroup) SetDefaultContactFlag(v bool)`

SetDefaultContactFlag sets DefaultContactFlag field to given value.

### HasDefaultContactFlag

`func (o *CompanyGroup) HasDefaultContactFlag() bool`

HasDefaultContactFlag returns a boolean if a field has been set.

### SetDefaultContactFlagNil

`func (o *CompanyGroup) SetDefaultContactFlagNil(b bool)`

 SetDefaultContactFlagNil sets the value for DefaultContactFlag to be an explicit nil

### UnsetDefaultContactFlag
`func (o *CompanyGroup) UnsetDefaultContactFlag()`

UnsetDefaultContactFlag ensures that no value is present for DefaultContactFlag, not even an explicit nil
### GetAllContactsFlag

`func (o *CompanyGroup) GetAllContactsFlag() bool`

GetAllContactsFlag returns the AllContactsFlag field if non-nil, zero value otherwise.

### GetAllContactsFlagOk

`func (o *CompanyGroup) GetAllContactsFlagOk() (*bool, bool)`

GetAllContactsFlagOk returns a tuple with the AllContactsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllContactsFlag

`func (o *CompanyGroup) SetAllContactsFlag(v bool)`

SetAllContactsFlag sets AllContactsFlag field to given value.

### HasAllContactsFlag

`func (o *CompanyGroup) HasAllContactsFlag() bool`

HasAllContactsFlag returns a boolean if a field has been set.

### SetAllContactsFlagNil

`func (o *CompanyGroup) SetAllContactsFlagNil(b bool)`

 SetAllContactsFlagNil sets the value for AllContactsFlag to be an explicit nil

### UnsetAllContactsFlag
`func (o *CompanyGroup) UnsetAllContactsFlag()`

UnsetAllContactsFlag ensures that no value is present for AllContactsFlag, not even an explicit nil
### GetRemoveAllContactsFlag

`func (o *CompanyGroup) GetRemoveAllContactsFlag() bool`

GetRemoveAllContactsFlag returns the RemoveAllContactsFlag field if non-nil, zero value otherwise.

### GetRemoveAllContactsFlagOk

`func (o *CompanyGroup) GetRemoveAllContactsFlagOk() (*bool, bool)`

GetRemoveAllContactsFlagOk returns a tuple with the RemoveAllContactsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllContactsFlag

`func (o *CompanyGroup) SetRemoveAllContactsFlag(v bool)`

SetRemoveAllContactsFlag sets RemoveAllContactsFlag field to given value.

### HasRemoveAllContactsFlag

`func (o *CompanyGroup) HasRemoveAllContactsFlag() bool`

HasRemoveAllContactsFlag returns a boolean if a field has been set.

### SetRemoveAllContactsFlagNil

`func (o *CompanyGroup) SetRemoveAllContactsFlagNil(b bool)`

 SetRemoveAllContactsFlagNil sets the value for RemoveAllContactsFlag to be an explicit nil

### UnsetRemoveAllContactsFlag
`func (o *CompanyGroup) UnsetRemoveAllContactsFlag()`

UnsetRemoveAllContactsFlag ensures that no value is present for RemoveAllContactsFlag, not even an explicit nil
### GetUnsubscribeFlag

`func (o *CompanyGroup) GetUnsubscribeFlag() bool`

GetUnsubscribeFlag returns the UnsubscribeFlag field if non-nil, zero value otherwise.

### GetUnsubscribeFlagOk

`func (o *CompanyGroup) GetUnsubscribeFlagOk() (*bool, bool)`

GetUnsubscribeFlagOk returns a tuple with the UnsubscribeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeFlag

`func (o *CompanyGroup) SetUnsubscribeFlag(v bool)`

SetUnsubscribeFlag sets UnsubscribeFlag field to given value.

### HasUnsubscribeFlag

`func (o *CompanyGroup) HasUnsubscribeFlag() bool`

HasUnsubscribeFlag returns a boolean if a field has been set.

### SetUnsubscribeFlagNil

`func (o *CompanyGroup) SetUnsubscribeFlagNil(b bool)`

 SetUnsubscribeFlagNil sets the value for UnsubscribeFlag to be an explicit nil

### UnsetUnsubscribeFlag
`func (o *CompanyGroup) UnsetUnsubscribeFlag()`

UnsetUnsubscribeFlag ensures that no value is present for UnsubscribeFlag, not even an explicit nil
### GetContactIds

`func (o *CompanyGroup) GetContactIds() []int32`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *CompanyGroup) GetContactIdsOk() (*[]int32, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *CompanyGroup) SetContactIds(v []int32)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *CompanyGroup) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyGroup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyGroup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyGroup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyGroup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


