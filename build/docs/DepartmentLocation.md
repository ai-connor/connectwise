# DepartmentLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**DepartmentManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Dispatch** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ServiceManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**DutyManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**LdapConfig** | Pointer to [**LdapConfigurationReference**](LdapConfigurationReference.md) |  | [optional] 
**AddAllLocations** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllLocations** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDepartmentLocation

`func NewDepartmentLocation(location SystemLocationReference, ) *DepartmentLocation`

NewDepartmentLocation instantiates a new DepartmentLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentLocationWithDefaults

`func NewDepartmentLocationWithDefaults() *DepartmentLocation`

NewDepartmentLocationWithDefaults instantiates a new DepartmentLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DepartmentLocation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepartmentLocation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepartmentLocation) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *DepartmentLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *DepartmentLocation) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DepartmentLocation) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DepartmentLocation) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetDepartment

`func (o *DepartmentLocation) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *DepartmentLocation) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *DepartmentLocation) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *DepartmentLocation) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDepartmentManager

`func (o *DepartmentLocation) GetDepartmentManager() MemberReference`

GetDepartmentManager returns the DepartmentManager field if non-nil, zero value otherwise.

### GetDepartmentManagerOk

`func (o *DepartmentLocation) GetDepartmentManagerOk() (*MemberReference, bool)`

GetDepartmentManagerOk returns a tuple with the DepartmentManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentManager

`func (o *DepartmentLocation) SetDepartmentManager(v MemberReference)`

SetDepartmentManager sets DepartmentManager field to given value.

### HasDepartmentManager

`func (o *DepartmentLocation) HasDepartmentManager() bool`

HasDepartmentManager returns a boolean if a field has been set.

### GetDispatch

`func (o *DepartmentLocation) GetDispatch() MemberReference`

GetDispatch returns the Dispatch field if non-nil, zero value otherwise.

### GetDispatchOk

`func (o *DepartmentLocation) GetDispatchOk() (*MemberReference, bool)`

GetDispatchOk returns a tuple with the Dispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatch

`func (o *DepartmentLocation) SetDispatch(v MemberReference)`

SetDispatch sets Dispatch field to given value.

### HasDispatch

`func (o *DepartmentLocation) HasDispatch() bool`

HasDispatch returns a boolean if a field has been set.

### GetServiceManager

`func (o *DepartmentLocation) GetServiceManager() MemberReference`

GetServiceManager returns the ServiceManager field if non-nil, zero value otherwise.

### GetServiceManagerOk

`func (o *DepartmentLocation) GetServiceManagerOk() (*MemberReference, bool)`

GetServiceManagerOk returns a tuple with the ServiceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManager

`func (o *DepartmentLocation) SetServiceManager(v MemberReference)`

SetServiceManager sets ServiceManager field to given value.

### HasServiceManager

`func (o *DepartmentLocation) HasServiceManager() bool`

HasServiceManager returns a boolean if a field has been set.

### GetDutyManager

`func (o *DepartmentLocation) GetDutyManager() MemberReference`

GetDutyManager returns the DutyManager field if non-nil, zero value otherwise.

### GetDutyManagerOk

`func (o *DepartmentLocation) GetDutyManagerOk() (*MemberReference, bool)`

GetDutyManagerOk returns a tuple with the DutyManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyManager

`func (o *DepartmentLocation) SetDutyManager(v MemberReference)`

SetDutyManager sets DutyManager field to given value.

### HasDutyManager

`func (o *DepartmentLocation) HasDutyManager() bool`

HasDutyManager returns a boolean if a field has been set.

### GetLdapConfig

`func (o *DepartmentLocation) GetLdapConfig() LdapConfigurationReference`

GetLdapConfig returns the LdapConfig field if non-nil, zero value otherwise.

### GetLdapConfigOk

`func (o *DepartmentLocation) GetLdapConfigOk() (*LdapConfigurationReference, bool)`

GetLdapConfigOk returns a tuple with the LdapConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapConfig

`func (o *DepartmentLocation) SetLdapConfig(v LdapConfigurationReference)`

SetLdapConfig sets LdapConfig field to given value.

### HasLdapConfig

`func (o *DepartmentLocation) HasLdapConfig() bool`

HasLdapConfig returns a boolean if a field has been set.

### GetAddAllLocations

`func (o *DepartmentLocation) GetAddAllLocations() bool`

GetAddAllLocations returns the AddAllLocations field if non-nil, zero value otherwise.

### GetAddAllLocationsOk

`func (o *DepartmentLocation) GetAddAllLocationsOk() (*bool, bool)`

GetAddAllLocationsOk returns a tuple with the AddAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllLocations

`func (o *DepartmentLocation) SetAddAllLocations(v bool)`

SetAddAllLocations sets AddAllLocations field to given value.

### HasAddAllLocations

`func (o *DepartmentLocation) HasAddAllLocations() bool`

HasAddAllLocations returns a boolean if a field has been set.

### SetAddAllLocationsNil

`func (o *DepartmentLocation) SetAddAllLocationsNil(b bool)`

 SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil

### UnsetAddAllLocations
`func (o *DepartmentLocation) UnsetAddAllLocations()`

UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
### GetRemoveAllLocations

`func (o *DepartmentLocation) GetRemoveAllLocations() bool`

GetRemoveAllLocations returns the RemoveAllLocations field if non-nil, zero value otherwise.

### GetRemoveAllLocationsOk

`func (o *DepartmentLocation) GetRemoveAllLocationsOk() (*bool, bool)`

GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllLocations

`func (o *DepartmentLocation) SetRemoveAllLocations(v bool)`

SetRemoveAllLocations sets RemoveAllLocations field to given value.

### HasRemoveAllLocations

`func (o *DepartmentLocation) HasRemoveAllLocations() bool`

HasRemoveAllLocations returns a boolean if a field has been set.

### SetRemoveAllLocationsNil

`func (o *DepartmentLocation) SetRemoveAllLocationsNil(b bool)`

 SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil

### UnsetRemoveAllLocations
`func (o *DepartmentLocation) UnsetRemoveAllLocations()`

UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
### GetInfo

`func (o *DepartmentLocation) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *DepartmentLocation) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *DepartmentLocation) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *DepartmentLocation) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


