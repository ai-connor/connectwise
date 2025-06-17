# SLA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 25; | 
**BasedOn** | **NullableString** |  | 
**CustomCalendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ApplicationOrder** | Pointer to **NullableInt32** |  | [optional] 
**HiImpactHiUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**HiImpactMedUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**HiImpactLowUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**MedImpactHiUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**MedImpactMedUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**MedImpactLowUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**LowImpactHiUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**LowImpactMedUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**LowImpactLowUrgency** | Pointer to [**PriorityReference**](PriorityReference.md) |  | [optional] 
**RespondHours** | Pointer to **NullableFloat64** |  | [optional] 
**RespondPercent** | Pointer to **NullableInt32** |  | [optional] 
**PlanWithin** | Pointer to **NullableFloat64** |  | [optional] 
**PlanWithinPercent** | Pointer to **NullableInt32** |  | [optional] 
**ResolutionHours** | Pointer to **NullableFloat64** |  | [optional] 
**ResolutionPercent** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSLA

`func NewSLA(name string, basedOn NullableString, ) *SLA`

NewSLA instantiates a new SLA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLAWithDefaults

`func NewSLAWithDefaults() *SLA`

NewSLAWithDefaults instantiates a new SLA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SLA) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLA) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLA) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SLA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SLA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLA) SetName(v string)`

SetName sets Name field to given value.


### GetBasedOn

`func (o *SLA) GetBasedOn() string`

GetBasedOn returns the BasedOn field if non-nil, zero value otherwise.

### GetBasedOnOk

`func (o *SLA) GetBasedOnOk() (*string, bool)`

GetBasedOnOk returns a tuple with the BasedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasedOn

`func (o *SLA) SetBasedOn(v string)`

SetBasedOn sets BasedOn field to given value.


### SetBasedOnNil

`func (o *SLA) SetBasedOnNil(b bool)`

 SetBasedOnNil sets the value for BasedOn to be an explicit nil

### UnsetBasedOn
`func (o *SLA) UnsetBasedOn()`

UnsetBasedOn ensures that no value is present for BasedOn, not even an explicit nil
### GetCustomCalendar

`func (o *SLA) GetCustomCalendar() CalendarReference`

GetCustomCalendar returns the CustomCalendar field if non-nil, zero value otherwise.

### GetCustomCalendarOk

`func (o *SLA) GetCustomCalendarOk() (*CalendarReference, bool)`

GetCustomCalendarOk returns a tuple with the CustomCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCalendar

`func (o *SLA) SetCustomCalendar(v CalendarReference)`

SetCustomCalendar sets CustomCalendar field to given value.

### HasCustomCalendar

`func (o *SLA) HasCustomCalendar() bool`

HasCustomCalendar returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *SLA) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *SLA) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *SLA) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *SLA) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *SLA) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *SLA) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetApplicationOrder

`func (o *SLA) GetApplicationOrder() int32`

GetApplicationOrder returns the ApplicationOrder field if non-nil, zero value otherwise.

### GetApplicationOrderOk

`func (o *SLA) GetApplicationOrderOk() (*int32, bool)`

GetApplicationOrderOk returns a tuple with the ApplicationOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationOrder

`func (o *SLA) SetApplicationOrder(v int32)`

SetApplicationOrder sets ApplicationOrder field to given value.

### HasApplicationOrder

`func (o *SLA) HasApplicationOrder() bool`

HasApplicationOrder returns a boolean if a field has been set.

### SetApplicationOrderNil

`func (o *SLA) SetApplicationOrderNil(b bool)`

 SetApplicationOrderNil sets the value for ApplicationOrder to be an explicit nil

### UnsetApplicationOrder
`func (o *SLA) UnsetApplicationOrder()`

UnsetApplicationOrder ensures that no value is present for ApplicationOrder, not even an explicit nil
### GetHiImpactHiUrgency

`func (o *SLA) GetHiImpactHiUrgency() PriorityReference`

GetHiImpactHiUrgency returns the HiImpactHiUrgency field if non-nil, zero value otherwise.

### GetHiImpactHiUrgencyOk

`func (o *SLA) GetHiImpactHiUrgencyOk() (*PriorityReference, bool)`

GetHiImpactHiUrgencyOk returns a tuple with the HiImpactHiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiImpactHiUrgency

`func (o *SLA) SetHiImpactHiUrgency(v PriorityReference)`

SetHiImpactHiUrgency sets HiImpactHiUrgency field to given value.

### HasHiImpactHiUrgency

`func (o *SLA) HasHiImpactHiUrgency() bool`

HasHiImpactHiUrgency returns a boolean if a field has been set.

### GetHiImpactMedUrgency

`func (o *SLA) GetHiImpactMedUrgency() PriorityReference`

GetHiImpactMedUrgency returns the HiImpactMedUrgency field if non-nil, zero value otherwise.

### GetHiImpactMedUrgencyOk

`func (o *SLA) GetHiImpactMedUrgencyOk() (*PriorityReference, bool)`

GetHiImpactMedUrgencyOk returns a tuple with the HiImpactMedUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiImpactMedUrgency

`func (o *SLA) SetHiImpactMedUrgency(v PriorityReference)`

SetHiImpactMedUrgency sets HiImpactMedUrgency field to given value.

### HasHiImpactMedUrgency

`func (o *SLA) HasHiImpactMedUrgency() bool`

HasHiImpactMedUrgency returns a boolean if a field has been set.

### GetHiImpactLowUrgency

`func (o *SLA) GetHiImpactLowUrgency() PriorityReference`

GetHiImpactLowUrgency returns the HiImpactLowUrgency field if non-nil, zero value otherwise.

### GetHiImpactLowUrgencyOk

`func (o *SLA) GetHiImpactLowUrgencyOk() (*PriorityReference, bool)`

GetHiImpactLowUrgencyOk returns a tuple with the HiImpactLowUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiImpactLowUrgency

`func (o *SLA) SetHiImpactLowUrgency(v PriorityReference)`

SetHiImpactLowUrgency sets HiImpactLowUrgency field to given value.

### HasHiImpactLowUrgency

`func (o *SLA) HasHiImpactLowUrgency() bool`

HasHiImpactLowUrgency returns a boolean if a field has been set.

### GetMedImpactHiUrgency

`func (o *SLA) GetMedImpactHiUrgency() PriorityReference`

GetMedImpactHiUrgency returns the MedImpactHiUrgency field if non-nil, zero value otherwise.

### GetMedImpactHiUrgencyOk

`func (o *SLA) GetMedImpactHiUrgencyOk() (*PriorityReference, bool)`

GetMedImpactHiUrgencyOk returns a tuple with the MedImpactHiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedImpactHiUrgency

`func (o *SLA) SetMedImpactHiUrgency(v PriorityReference)`

SetMedImpactHiUrgency sets MedImpactHiUrgency field to given value.

### HasMedImpactHiUrgency

`func (o *SLA) HasMedImpactHiUrgency() bool`

HasMedImpactHiUrgency returns a boolean if a field has been set.

### GetMedImpactMedUrgency

`func (o *SLA) GetMedImpactMedUrgency() PriorityReference`

GetMedImpactMedUrgency returns the MedImpactMedUrgency field if non-nil, zero value otherwise.

### GetMedImpactMedUrgencyOk

`func (o *SLA) GetMedImpactMedUrgencyOk() (*PriorityReference, bool)`

GetMedImpactMedUrgencyOk returns a tuple with the MedImpactMedUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedImpactMedUrgency

`func (o *SLA) SetMedImpactMedUrgency(v PriorityReference)`

SetMedImpactMedUrgency sets MedImpactMedUrgency field to given value.

### HasMedImpactMedUrgency

`func (o *SLA) HasMedImpactMedUrgency() bool`

HasMedImpactMedUrgency returns a boolean if a field has been set.

### GetMedImpactLowUrgency

`func (o *SLA) GetMedImpactLowUrgency() PriorityReference`

GetMedImpactLowUrgency returns the MedImpactLowUrgency field if non-nil, zero value otherwise.

### GetMedImpactLowUrgencyOk

`func (o *SLA) GetMedImpactLowUrgencyOk() (*PriorityReference, bool)`

GetMedImpactLowUrgencyOk returns a tuple with the MedImpactLowUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedImpactLowUrgency

`func (o *SLA) SetMedImpactLowUrgency(v PriorityReference)`

SetMedImpactLowUrgency sets MedImpactLowUrgency field to given value.

### HasMedImpactLowUrgency

`func (o *SLA) HasMedImpactLowUrgency() bool`

HasMedImpactLowUrgency returns a boolean if a field has been set.

### GetLowImpactHiUrgency

`func (o *SLA) GetLowImpactHiUrgency() PriorityReference`

GetLowImpactHiUrgency returns the LowImpactHiUrgency field if non-nil, zero value otherwise.

### GetLowImpactHiUrgencyOk

`func (o *SLA) GetLowImpactHiUrgencyOk() (*PriorityReference, bool)`

GetLowImpactHiUrgencyOk returns a tuple with the LowImpactHiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowImpactHiUrgency

`func (o *SLA) SetLowImpactHiUrgency(v PriorityReference)`

SetLowImpactHiUrgency sets LowImpactHiUrgency field to given value.

### HasLowImpactHiUrgency

`func (o *SLA) HasLowImpactHiUrgency() bool`

HasLowImpactHiUrgency returns a boolean if a field has been set.

### GetLowImpactMedUrgency

`func (o *SLA) GetLowImpactMedUrgency() PriorityReference`

GetLowImpactMedUrgency returns the LowImpactMedUrgency field if non-nil, zero value otherwise.

### GetLowImpactMedUrgencyOk

`func (o *SLA) GetLowImpactMedUrgencyOk() (*PriorityReference, bool)`

GetLowImpactMedUrgencyOk returns a tuple with the LowImpactMedUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowImpactMedUrgency

`func (o *SLA) SetLowImpactMedUrgency(v PriorityReference)`

SetLowImpactMedUrgency sets LowImpactMedUrgency field to given value.

### HasLowImpactMedUrgency

`func (o *SLA) HasLowImpactMedUrgency() bool`

HasLowImpactMedUrgency returns a boolean if a field has been set.

### GetLowImpactLowUrgency

`func (o *SLA) GetLowImpactLowUrgency() PriorityReference`

GetLowImpactLowUrgency returns the LowImpactLowUrgency field if non-nil, zero value otherwise.

### GetLowImpactLowUrgencyOk

`func (o *SLA) GetLowImpactLowUrgencyOk() (*PriorityReference, bool)`

GetLowImpactLowUrgencyOk returns a tuple with the LowImpactLowUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowImpactLowUrgency

`func (o *SLA) SetLowImpactLowUrgency(v PriorityReference)`

SetLowImpactLowUrgency sets LowImpactLowUrgency field to given value.

### HasLowImpactLowUrgency

`func (o *SLA) HasLowImpactLowUrgency() bool`

HasLowImpactLowUrgency returns a boolean if a field has been set.

### GetRespondHours

`func (o *SLA) GetRespondHours() float64`

GetRespondHours returns the RespondHours field if non-nil, zero value otherwise.

### GetRespondHoursOk

`func (o *SLA) GetRespondHoursOk() (*float64, bool)`

GetRespondHoursOk returns a tuple with the RespondHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondHours

`func (o *SLA) SetRespondHours(v float64)`

SetRespondHours sets RespondHours field to given value.

### HasRespondHours

`func (o *SLA) HasRespondHours() bool`

HasRespondHours returns a boolean if a field has been set.

### SetRespondHoursNil

`func (o *SLA) SetRespondHoursNil(b bool)`

 SetRespondHoursNil sets the value for RespondHours to be an explicit nil

### UnsetRespondHours
`func (o *SLA) UnsetRespondHours()`

UnsetRespondHours ensures that no value is present for RespondHours, not even an explicit nil
### GetRespondPercent

`func (o *SLA) GetRespondPercent() int32`

GetRespondPercent returns the RespondPercent field if non-nil, zero value otherwise.

### GetRespondPercentOk

`func (o *SLA) GetRespondPercentOk() (*int32, bool)`

GetRespondPercentOk returns a tuple with the RespondPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondPercent

`func (o *SLA) SetRespondPercent(v int32)`

SetRespondPercent sets RespondPercent field to given value.

### HasRespondPercent

`func (o *SLA) HasRespondPercent() bool`

HasRespondPercent returns a boolean if a field has been set.

### SetRespondPercentNil

`func (o *SLA) SetRespondPercentNil(b bool)`

 SetRespondPercentNil sets the value for RespondPercent to be an explicit nil

### UnsetRespondPercent
`func (o *SLA) UnsetRespondPercent()`

UnsetRespondPercent ensures that no value is present for RespondPercent, not even an explicit nil
### GetPlanWithin

`func (o *SLA) GetPlanWithin() float64`

GetPlanWithin returns the PlanWithin field if non-nil, zero value otherwise.

### GetPlanWithinOk

`func (o *SLA) GetPlanWithinOk() (*float64, bool)`

GetPlanWithinOk returns a tuple with the PlanWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanWithin

`func (o *SLA) SetPlanWithin(v float64)`

SetPlanWithin sets PlanWithin field to given value.

### HasPlanWithin

`func (o *SLA) HasPlanWithin() bool`

HasPlanWithin returns a boolean if a field has been set.

### SetPlanWithinNil

`func (o *SLA) SetPlanWithinNil(b bool)`

 SetPlanWithinNil sets the value for PlanWithin to be an explicit nil

### UnsetPlanWithin
`func (o *SLA) UnsetPlanWithin()`

UnsetPlanWithin ensures that no value is present for PlanWithin, not even an explicit nil
### GetPlanWithinPercent

`func (o *SLA) GetPlanWithinPercent() int32`

GetPlanWithinPercent returns the PlanWithinPercent field if non-nil, zero value otherwise.

### GetPlanWithinPercentOk

`func (o *SLA) GetPlanWithinPercentOk() (*int32, bool)`

GetPlanWithinPercentOk returns a tuple with the PlanWithinPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanWithinPercent

`func (o *SLA) SetPlanWithinPercent(v int32)`

SetPlanWithinPercent sets PlanWithinPercent field to given value.

### HasPlanWithinPercent

`func (o *SLA) HasPlanWithinPercent() bool`

HasPlanWithinPercent returns a boolean if a field has been set.

### SetPlanWithinPercentNil

`func (o *SLA) SetPlanWithinPercentNil(b bool)`

 SetPlanWithinPercentNil sets the value for PlanWithinPercent to be an explicit nil

### UnsetPlanWithinPercent
`func (o *SLA) UnsetPlanWithinPercent()`

UnsetPlanWithinPercent ensures that no value is present for PlanWithinPercent, not even an explicit nil
### GetResolutionHours

`func (o *SLA) GetResolutionHours() float64`

GetResolutionHours returns the ResolutionHours field if non-nil, zero value otherwise.

### GetResolutionHoursOk

`func (o *SLA) GetResolutionHoursOk() (*float64, bool)`

GetResolutionHoursOk returns a tuple with the ResolutionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionHours

`func (o *SLA) SetResolutionHours(v float64)`

SetResolutionHours sets ResolutionHours field to given value.

### HasResolutionHours

`func (o *SLA) HasResolutionHours() bool`

HasResolutionHours returns a boolean if a field has been set.

### SetResolutionHoursNil

`func (o *SLA) SetResolutionHoursNil(b bool)`

 SetResolutionHoursNil sets the value for ResolutionHours to be an explicit nil

### UnsetResolutionHours
`func (o *SLA) UnsetResolutionHours()`

UnsetResolutionHours ensures that no value is present for ResolutionHours, not even an explicit nil
### GetResolutionPercent

`func (o *SLA) GetResolutionPercent() int32`

GetResolutionPercent returns the ResolutionPercent field if non-nil, zero value otherwise.

### GetResolutionPercentOk

`func (o *SLA) GetResolutionPercentOk() (*int32, bool)`

GetResolutionPercentOk returns a tuple with the ResolutionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPercent

`func (o *SLA) SetResolutionPercent(v int32)`

SetResolutionPercent sets ResolutionPercent field to given value.

### HasResolutionPercent

`func (o *SLA) HasResolutionPercent() bool`

HasResolutionPercent returns a boolean if a field has been set.

### SetResolutionPercentNil

`func (o *SLA) SetResolutionPercentNil(b bool)`

 SetResolutionPercentNil sets the value for ResolutionPercent to be an explicit nil

### UnsetResolutionPercent
`func (o *SLA) UnsetResolutionPercent()`

UnsetResolutionPercent ensures that no value is present for ResolutionPercent, not even an explicit nil
### GetInfo

`func (o *SLA) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SLA) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SLA) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SLA) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


