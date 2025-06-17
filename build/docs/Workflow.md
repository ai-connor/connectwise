# Workflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**TableType** | [**WorkflowTableTypeReference**](WorkflowTableTypeReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**ActivateFlag** | Pointer to **NullableBool** | Batches can not be turned on until after the workflow is created and it has atleast one event associated with it | [optional] 
**BatchInterval** | Pointer to **NullableInt32** |  | [optional] 
**BatchFrequencyUnit** | Pointer to **NullableString** | If not specified, defaults to Minutes. Months is not supported as month length varies | [optional] 
**BatchLastRan** | Pointer to **time.Time** |  | [optional] 
**BatchSchedule** | Pointer to **NullableString** | If activateFlag is true, batchSchedule is required | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflow

`func NewWorkflow(name string, tableType WorkflowTableTypeReference, ) *Workflow`

NewWorkflow instantiates a new Workflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowWithDefaults

`func NewWorkflowWithDefaults() *Workflow`

NewWorkflowWithDefaults instantiates a new Workflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Workflow) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Workflow) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Workflow) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Workflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Workflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Workflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Workflow) SetName(v string)`

SetName sets Name field to given value.


### GetTableType

`func (o *Workflow) GetTableType() WorkflowTableTypeReference`

GetTableType returns the TableType field if non-nil, zero value otherwise.

### GetTableTypeOk

`func (o *Workflow) GetTableTypeOk() (*WorkflowTableTypeReference, bool)`

GetTableTypeOk returns a tuple with the TableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableType

`func (o *Workflow) SetTableType(v WorkflowTableTypeReference)`

SetTableType sets TableType field to given value.


### GetLocation

`func (o *Workflow) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Workflow) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Workflow) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Workflow) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *Workflow) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Workflow) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Workflow) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Workflow) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetActivateFlag

`func (o *Workflow) GetActivateFlag() bool`

GetActivateFlag returns the ActivateFlag field if non-nil, zero value otherwise.

### GetActivateFlagOk

`func (o *Workflow) GetActivateFlagOk() (*bool, bool)`

GetActivateFlagOk returns a tuple with the ActivateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateFlag

`func (o *Workflow) SetActivateFlag(v bool)`

SetActivateFlag sets ActivateFlag field to given value.

### HasActivateFlag

`func (o *Workflow) HasActivateFlag() bool`

HasActivateFlag returns a boolean if a field has been set.

### SetActivateFlagNil

`func (o *Workflow) SetActivateFlagNil(b bool)`

 SetActivateFlagNil sets the value for ActivateFlag to be an explicit nil

### UnsetActivateFlag
`func (o *Workflow) UnsetActivateFlag()`

UnsetActivateFlag ensures that no value is present for ActivateFlag, not even an explicit nil
### GetBatchInterval

`func (o *Workflow) GetBatchInterval() int32`

GetBatchInterval returns the BatchInterval field if non-nil, zero value otherwise.

### GetBatchIntervalOk

`func (o *Workflow) GetBatchIntervalOk() (*int32, bool)`

GetBatchIntervalOk returns a tuple with the BatchInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchInterval

`func (o *Workflow) SetBatchInterval(v int32)`

SetBatchInterval sets BatchInterval field to given value.

### HasBatchInterval

`func (o *Workflow) HasBatchInterval() bool`

HasBatchInterval returns a boolean if a field has been set.

### SetBatchIntervalNil

`func (o *Workflow) SetBatchIntervalNil(b bool)`

 SetBatchIntervalNil sets the value for BatchInterval to be an explicit nil

### UnsetBatchInterval
`func (o *Workflow) UnsetBatchInterval()`

UnsetBatchInterval ensures that no value is present for BatchInterval, not even an explicit nil
### GetBatchFrequencyUnit

`func (o *Workflow) GetBatchFrequencyUnit() string`

GetBatchFrequencyUnit returns the BatchFrequencyUnit field if non-nil, zero value otherwise.

### GetBatchFrequencyUnitOk

`func (o *Workflow) GetBatchFrequencyUnitOk() (*string, bool)`

GetBatchFrequencyUnitOk returns a tuple with the BatchFrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchFrequencyUnit

`func (o *Workflow) SetBatchFrequencyUnit(v string)`

SetBatchFrequencyUnit sets BatchFrequencyUnit field to given value.

### HasBatchFrequencyUnit

`func (o *Workflow) HasBatchFrequencyUnit() bool`

HasBatchFrequencyUnit returns a boolean if a field has been set.

### SetBatchFrequencyUnitNil

`func (o *Workflow) SetBatchFrequencyUnitNil(b bool)`

 SetBatchFrequencyUnitNil sets the value for BatchFrequencyUnit to be an explicit nil

### UnsetBatchFrequencyUnit
`func (o *Workflow) UnsetBatchFrequencyUnit()`

UnsetBatchFrequencyUnit ensures that no value is present for BatchFrequencyUnit, not even an explicit nil
### GetBatchLastRan

`func (o *Workflow) GetBatchLastRan() time.Time`

GetBatchLastRan returns the BatchLastRan field if non-nil, zero value otherwise.

### GetBatchLastRanOk

`func (o *Workflow) GetBatchLastRanOk() (*time.Time, bool)`

GetBatchLastRanOk returns a tuple with the BatchLastRan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchLastRan

`func (o *Workflow) SetBatchLastRan(v time.Time)`

SetBatchLastRan sets BatchLastRan field to given value.

### HasBatchLastRan

`func (o *Workflow) HasBatchLastRan() bool`

HasBatchLastRan returns a boolean if a field has been set.

### GetBatchSchedule

`func (o *Workflow) GetBatchSchedule() string`

GetBatchSchedule returns the BatchSchedule field if non-nil, zero value otherwise.

### GetBatchScheduleOk

`func (o *Workflow) GetBatchScheduleOk() (*string, bool)`

GetBatchScheduleOk returns a tuple with the BatchSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSchedule

`func (o *Workflow) SetBatchSchedule(v string)`

SetBatchSchedule sets BatchSchedule field to given value.

### HasBatchSchedule

`func (o *Workflow) HasBatchSchedule() bool`

HasBatchSchedule returns a boolean if a field has been set.

### SetBatchScheduleNil

`func (o *Workflow) SetBatchScheduleNil(b bool)`

 SetBatchScheduleNil sets the value for BatchSchedule to be an explicit nil

### UnsetBatchSchedule
`func (o *Workflow) UnsetBatchSchedule()`

UnsetBatchSchedule ensures that no value is present for BatchSchedule, not even an explicit nil
### GetBoard

`func (o *Workflow) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *Workflow) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *Workflow) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *Workflow) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *Workflow) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *Workflow) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *Workflow) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *Workflow) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetInfo

`func (o *Workflow) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Workflow) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Workflow) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Workflow) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


