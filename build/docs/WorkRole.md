# WorkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**HourlyRate** | Pointer to **NullableFloat64** |  | [optional] 
**IntegrationXref** | Pointer to **string** |  Max length: 50; | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**AddAllLocations** | Pointer to **NullableBool** |  | [optional] 
**RemoveAllLocations** | Pointer to **NullableBool** |  | [optional] 
**AddAllAgreementExclusions** | Pointer to **NullableBool** | Used only on create to add the work role to all agreement and agreement type exclusion lists | [optional] 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkRole

`func NewWorkRole(name string, ) *WorkRole`

NewWorkRole instantiates a new WorkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkRoleWithDefaults

`func NewWorkRoleWithDefaults() *WorkRole`

NewWorkRoleWithDefaults instantiates a new WorkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkRole) SetName(v string)`

SetName sets Name field to given value.


### GetHourlyRate

`func (o *WorkRole) GetHourlyRate() float64`

GetHourlyRate returns the HourlyRate field if non-nil, zero value otherwise.

### GetHourlyRateOk

`func (o *WorkRole) GetHourlyRateOk() (*float64, bool)`

GetHourlyRateOk returns a tuple with the HourlyRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourlyRate

`func (o *WorkRole) SetHourlyRate(v float64)`

SetHourlyRate sets HourlyRate field to given value.

### HasHourlyRate

`func (o *WorkRole) HasHourlyRate() bool`

HasHourlyRate returns a boolean if a field has been set.

### SetHourlyRateNil

`func (o *WorkRole) SetHourlyRateNil(b bool)`

 SetHourlyRateNil sets the value for HourlyRate to be an explicit nil

### UnsetHourlyRate
`func (o *WorkRole) UnsetHourlyRate()`

UnsetHourlyRate ensures that no value is present for HourlyRate, not even an explicit nil
### GetIntegrationXref

`func (o *WorkRole) GetIntegrationXref() string`

GetIntegrationXref returns the IntegrationXref field if non-nil, zero value otherwise.

### GetIntegrationXrefOk

`func (o *WorkRole) GetIntegrationXrefOk() (*string, bool)`

GetIntegrationXrefOk returns a tuple with the IntegrationXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXref

`func (o *WorkRole) SetIntegrationXref(v string)`

SetIntegrationXref sets IntegrationXref field to given value.

### HasIntegrationXref

`func (o *WorkRole) HasIntegrationXref() bool`

HasIntegrationXref returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *WorkRole) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WorkRole) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WorkRole) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WorkRole) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WorkRole) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WorkRole) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetAddAllLocations

`func (o *WorkRole) GetAddAllLocations() bool`

GetAddAllLocations returns the AddAllLocations field if non-nil, zero value otherwise.

### GetAddAllLocationsOk

`func (o *WorkRole) GetAddAllLocationsOk() (*bool, bool)`

GetAddAllLocationsOk returns a tuple with the AddAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllLocations

`func (o *WorkRole) SetAddAllLocations(v bool)`

SetAddAllLocations sets AddAllLocations field to given value.

### HasAddAllLocations

`func (o *WorkRole) HasAddAllLocations() bool`

HasAddAllLocations returns a boolean if a field has been set.

### SetAddAllLocationsNil

`func (o *WorkRole) SetAddAllLocationsNil(b bool)`

 SetAddAllLocationsNil sets the value for AddAllLocations to be an explicit nil

### UnsetAddAllLocations
`func (o *WorkRole) UnsetAddAllLocations()`

UnsetAddAllLocations ensures that no value is present for AddAllLocations, not even an explicit nil
### GetRemoveAllLocations

`func (o *WorkRole) GetRemoveAllLocations() bool`

GetRemoveAllLocations returns the RemoveAllLocations field if non-nil, zero value otherwise.

### GetRemoveAllLocationsOk

`func (o *WorkRole) GetRemoveAllLocationsOk() (*bool, bool)`

GetRemoveAllLocationsOk returns a tuple with the RemoveAllLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAllLocations

`func (o *WorkRole) SetRemoveAllLocations(v bool)`

SetRemoveAllLocations sets RemoveAllLocations field to given value.

### HasRemoveAllLocations

`func (o *WorkRole) HasRemoveAllLocations() bool`

HasRemoveAllLocations returns a boolean if a field has been set.

### SetRemoveAllLocationsNil

`func (o *WorkRole) SetRemoveAllLocationsNil(b bool)`

 SetRemoveAllLocationsNil sets the value for RemoveAllLocations to be an explicit nil

### UnsetRemoveAllLocations
`func (o *WorkRole) UnsetRemoveAllLocations()`

UnsetRemoveAllLocations ensures that no value is present for RemoveAllLocations, not even an explicit nil
### GetAddAllAgreementExclusions

`func (o *WorkRole) GetAddAllAgreementExclusions() bool`

GetAddAllAgreementExclusions returns the AddAllAgreementExclusions field if non-nil, zero value otherwise.

### GetAddAllAgreementExclusionsOk

`func (o *WorkRole) GetAddAllAgreementExclusionsOk() (*bool, bool)`

GetAddAllAgreementExclusionsOk returns a tuple with the AddAllAgreementExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllAgreementExclusions

`func (o *WorkRole) SetAddAllAgreementExclusions(v bool)`

SetAddAllAgreementExclusions sets AddAllAgreementExclusions field to given value.

### HasAddAllAgreementExclusions

`func (o *WorkRole) HasAddAllAgreementExclusions() bool`

HasAddAllAgreementExclusions returns a boolean if a field has been set.

### SetAddAllAgreementExclusionsNil

`func (o *WorkRole) SetAddAllAgreementExclusionsNil(b bool)`

 SetAddAllAgreementExclusionsNil sets the value for AddAllAgreementExclusions to be an explicit nil

### UnsetAddAllAgreementExclusions
`func (o *WorkRole) UnsetAddAllAgreementExclusions()`

UnsetAddAllAgreementExclusions ensures that no value is present for AddAllAgreementExclusions, not even an explicit nil
### GetLocationIds

`func (o *WorkRole) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *WorkRole) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *WorkRole) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *WorkRole) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetInfo

`func (o *WorkRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


