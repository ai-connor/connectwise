# SLAPriority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Priority** | [**PriorityReference**](PriorityReference.md) |  | 
**RespondHours** | Pointer to **NullableFloat64** |  | [optional] 
**RespondPercent** | Pointer to **NullableInt32** |  | [optional] 
**PlanWithin** | Pointer to **NullableFloat64** |  | [optional] 
**PlanWithinPercent** | Pointer to **NullableInt32** |  | [optional] 
**ResolutionHours** | Pointer to **NullableFloat64** |  | [optional] 
**ResolutionPercent** | Pointer to **NullableInt32** |  | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSLAPriority

`func NewSLAPriority(priority PriorityReference, ) *SLAPriority`

NewSLAPriority instantiates a new SLAPriority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLAPriorityWithDefaults

`func NewSLAPriorityWithDefaults() *SLAPriority`

NewSLAPriorityWithDefaults instantiates a new SLAPriority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SLAPriority) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLAPriority) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLAPriority) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SLAPriority) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SLAPriority) GetPriority() PriorityReference`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SLAPriority) GetPriorityOk() (*PriorityReference, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SLAPriority) SetPriority(v PriorityReference)`

SetPriority sets Priority field to given value.


### GetRespondHours

`func (o *SLAPriority) GetRespondHours() float64`

GetRespondHours returns the RespondHours field if non-nil, zero value otherwise.

### GetRespondHoursOk

`func (o *SLAPriority) GetRespondHoursOk() (*float64, bool)`

GetRespondHoursOk returns a tuple with the RespondHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondHours

`func (o *SLAPriority) SetRespondHours(v float64)`

SetRespondHours sets RespondHours field to given value.

### HasRespondHours

`func (o *SLAPriority) HasRespondHours() bool`

HasRespondHours returns a boolean if a field has been set.

### SetRespondHoursNil

`func (o *SLAPriority) SetRespondHoursNil(b bool)`

 SetRespondHoursNil sets the value for RespondHours to be an explicit nil

### UnsetRespondHours
`func (o *SLAPriority) UnsetRespondHours()`

UnsetRespondHours ensures that no value is present for RespondHours, not even an explicit nil
### GetRespondPercent

`func (o *SLAPriority) GetRespondPercent() int32`

GetRespondPercent returns the RespondPercent field if non-nil, zero value otherwise.

### GetRespondPercentOk

`func (o *SLAPriority) GetRespondPercentOk() (*int32, bool)`

GetRespondPercentOk returns a tuple with the RespondPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespondPercent

`func (o *SLAPriority) SetRespondPercent(v int32)`

SetRespondPercent sets RespondPercent field to given value.

### HasRespondPercent

`func (o *SLAPriority) HasRespondPercent() bool`

HasRespondPercent returns a boolean if a field has been set.

### SetRespondPercentNil

`func (o *SLAPriority) SetRespondPercentNil(b bool)`

 SetRespondPercentNil sets the value for RespondPercent to be an explicit nil

### UnsetRespondPercent
`func (o *SLAPriority) UnsetRespondPercent()`

UnsetRespondPercent ensures that no value is present for RespondPercent, not even an explicit nil
### GetPlanWithin

`func (o *SLAPriority) GetPlanWithin() float64`

GetPlanWithin returns the PlanWithin field if non-nil, zero value otherwise.

### GetPlanWithinOk

`func (o *SLAPriority) GetPlanWithinOk() (*float64, bool)`

GetPlanWithinOk returns a tuple with the PlanWithin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanWithin

`func (o *SLAPriority) SetPlanWithin(v float64)`

SetPlanWithin sets PlanWithin field to given value.

### HasPlanWithin

`func (o *SLAPriority) HasPlanWithin() bool`

HasPlanWithin returns a boolean if a field has been set.

### SetPlanWithinNil

`func (o *SLAPriority) SetPlanWithinNil(b bool)`

 SetPlanWithinNil sets the value for PlanWithin to be an explicit nil

### UnsetPlanWithin
`func (o *SLAPriority) UnsetPlanWithin()`

UnsetPlanWithin ensures that no value is present for PlanWithin, not even an explicit nil
### GetPlanWithinPercent

`func (o *SLAPriority) GetPlanWithinPercent() int32`

GetPlanWithinPercent returns the PlanWithinPercent field if non-nil, zero value otherwise.

### GetPlanWithinPercentOk

`func (o *SLAPriority) GetPlanWithinPercentOk() (*int32, bool)`

GetPlanWithinPercentOk returns a tuple with the PlanWithinPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanWithinPercent

`func (o *SLAPriority) SetPlanWithinPercent(v int32)`

SetPlanWithinPercent sets PlanWithinPercent field to given value.

### HasPlanWithinPercent

`func (o *SLAPriority) HasPlanWithinPercent() bool`

HasPlanWithinPercent returns a boolean if a field has been set.

### SetPlanWithinPercentNil

`func (o *SLAPriority) SetPlanWithinPercentNil(b bool)`

 SetPlanWithinPercentNil sets the value for PlanWithinPercent to be an explicit nil

### UnsetPlanWithinPercent
`func (o *SLAPriority) UnsetPlanWithinPercent()`

UnsetPlanWithinPercent ensures that no value is present for PlanWithinPercent, not even an explicit nil
### GetResolutionHours

`func (o *SLAPriority) GetResolutionHours() float64`

GetResolutionHours returns the ResolutionHours field if non-nil, zero value otherwise.

### GetResolutionHoursOk

`func (o *SLAPriority) GetResolutionHoursOk() (*float64, bool)`

GetResolutionHoursOk returns a tuple with the ResolutionHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionHours

`func (o *SLAPriority) SetResolutionHours(v float64)`

SetResolutionHours sets ResolutionHours field to given value.

### HasResolutionHours

`func (o *SLAPriority) HasResolutionHours() bool`

HasResolutionHours returns a boolean if a field has been set.

### SetResolutionHoursNil

`func (o *SLAPriority) SetResolutionHoursNil(b bool)`

 SetResolutionHoursNil sets the value for ResolutionHours to be an explicit nil

### UnsetResolutionHours
`func (o *SLAPriority) UnsetResolutionHours()`

UnsetResolutionHours ensures that no value is present for ResolutionHours, not even an explicit nil
### GetResolutionPercent

`func (o *SLAPriority) GetResolutionPercent() int32`

GetResolutionPercent returns the ResolutionPercent field if non-nil, zero value otherwise.

### GetResolutionPercentOk

`func (o *SLAPriority) GetResolutionPercentOk() (*int32, bool)`

GetResolutionPercentOk returns a tuple with the ResolutionPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionPercent

`func (o *SLAPriority) SetResolutionPercent(v int32)`

SetResolutionPercent sets ResolutionPercent field to given value.

### HasResolutionPercent

`func (o *SLAPriority) HasResolutionPercent() bool`

HasResolutionPercent returns a boolean if a field has been set.

### SetResolutionPercentNil

`func (o *SLAPriority) SetResolutionPercentNil(b bool)`

 SetResolutionPercentNil sets the value for ResolutionPercent to be an explicit nil

### UnsetResolutionPercent
`func (o *SLAPriority) UnsetResolutionPercent()`

UnsetResolutionPercent ensures that no value is present for ResolutionPercent, not even an explicit nil
### GetSla

`func (o *SLAPriority) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *SLAPriority) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *SLAPriority) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *SLAPriority) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetInfo

`func (o *SLAPriority) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SLAPriority) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SLAPriority) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SLAPriority) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


