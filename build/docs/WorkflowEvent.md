# WorkflowEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EventCondition** | **string** |  | 
**FrequencyUnit** | Pointer to **NullableString** | Required when exectionTimes is set to MultipleTimes or Continuously | [optional] 
**FrequencyOfExecution** | Pointer to **NullableInt32** | Required when exectionTimes is set to MultipleTimes or Continuously | [optional] 
**MaxNumberOfExecution** | Pointer to **NullableInt32** | Required when exectionTimes is set to MultipleTimes | [optional] 
**ExecutionTime** | Pointer to **NullableString** | Defaults to Once when not specified | [optional] 
**DateTestedUTC** | Pointer to **time.Time** |  | [optional] 
**TestRecordsMatched** | Pointer to **NullableInt32** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyHeader_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowEvent

`func NewWorkflowEvent(eventCondition string, ) *WorkflowEvent`

NewWorkflowEvent instantiates a new WorkflowEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEventWithDefaults

`func NewWorkflowEventWithDefaults() *WorkflowEvent`

NewWorkflowEventWithDefaults instantiates a new WorkflowEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowEvent) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEventCondition

`func (o *WorkflowEvent) GetEventCondition() string`

GetEventCondition returns the EventCondition field if non-nil, zero value otherwise.

### GetEventConditionOk

`func (o *WorkflowEvent) GetEventConditionOk() (*string, bool)`

GetEventConditionOk returns a tuple with the EventCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCondition

`func (o *WorkflowEvent) SetEventCondition(v string)`

SetEventCondition sets EventCondition field to given value.


### GetFrequencyUnit

`func (o *WorkflowEvent) GetFrequencyUnit() string`

GetFrequencyUnit returns the FrequencyUnit field if non-nil, zero value otherwise.

### GetFrequencyUnitOk

`func (o *WorkflowEvent) GetFrequencyUnitOk() (*string, bool)`

GetFrequencyUnitOk returns a tuple with the FrequencyUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyUnit

`func (o *WorkflowEvent) SetFrequencyUnit(v string)`

SetFrequencyUnit sets FrequencyUnit field to given value.

### HasFrequencyUnit

`func (o *WorkflowEvent) HasFrequencyUnit() bool`

HasFrequencyUnit returns a boolean if a field has been set.

### SetFrequencyUnitNil

`func (o *WorkflowEvent) SetFrequencyUnitNil(b bool)`

 SetFrequencyUnitNil sets the value for FrequencyUnit to be an explicit nil

### UnsetFrequencyUnit
`func (o *WorkflowEvent) UnsetFrequencyUnit()`

UnsetFrequencyUnit ensures that no value is present for FrequencyUnit, not even an explicit nil
### GetFrequencyOfExecution

`func (o *WorkflowEvent) GetFrequencyOfExecution() int32`

GetFrequencyOfExecution returns the FrequencyOfExecution field if non-nil, zero value otherwise.

### GetFrequencyOfExecutionOk

`func (o *WorkflowEvent) GetFrequencyOfExecutionOk() (*int32, bool)`

GetFrequencyOfExecutionOk returns a tuple with the FrequencyOfExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyOfExecution

`func (o *WorkflowEvent) SetFrequencyOfExecution(v int32)`

SetFrequencyOfExecution sets FrequencyOfExecution field to given value.

### HasFrequencyOfExecution

`func (o *WorkflowEvent) HasFrequencyOfExecution() bool`

HasFrequencyOfExecution returns a boolean if a field has been set.

### SetFrequencyOfExecutionNil

`func (o *WorkflowEvent) SetFrequencyOfExecutionNil(b bool)`

 SetFrequencyOfExecutionNil sets the value for FrequencyOfExecution to be an explicit nil

### UnsetFrequencyOfExecution
`func (o *WorkflowEvent) UnsetFrequencyOfExecution()`

UnsetFrequencyOfExecution ensures that no value is present for FrequencyOfExecution, not even an explicit nil
### GetMaxNumberOfExecution

`func (o *WorkflowEvent) GetMaxNumberOfExecution() int32`

GetMaxNumberOfExecution returns the MaxNumberOfExecution field if non-nil, zero value otherwise.

### GetMaxNumberOfExecutionOk

`func (o *WorkflowEvent) GetMaxNumberOfExecutionOk() (*int32, bool)`

GetMaxNumberOfExecutionOk returns a tuple with the MaxNumberOfExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfExecution

`func (o *WorkflowEvent) SetMaxNumberOfExecution(v int32)`

SetMaxNumberOfExecution sets MaxNumberOfExecution field to given value.

### HasMaxNumberOfExecution

`func (o *WorkflowEvent) HasMaxNumberOfExecution() bool`

HasMaxNumberOfExecution returns a boolean if a field has been set.

### SetMaxNumberOfExecutionNil

`func (o *WorkflowEvent) SetMaxNumberOfExecutionNil(b bool)`

 SetMaxNumberOfExecutionNil sets the value for MaxNumberOfExecution to be an explicit nil

### UnsetMaxNumberOfExecution
`func (o *WorkflowEvent) UnsetMaxNumberOfExecution()`

UnsetMaxNumberOfExecution ensures that no value is present for MaxNumberOfExecution, not even an explicit nil
### GetExecutionTime

`func (o *WorkflowEvent) GetExecutionTime() string`

GetExecutionTime returns the ExecutionTime field if non-nil, zero value otherwise.

### GetExecutionTimeOk

`func (o *WorkflowEvent) GetExecutionTimeOk() (*string, bool)`

GetExecutionTimeOk returns a tuple with the ExecutionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTime

`func (o *WorkflowEvent) SetExecutionTime(v string)`

SetExecutionTime sets ExecutionTime field to given value.

### HasExecutionTime

`func (o *WorkflowEvent) HasExecutionTime() bool`

HasExecutionTime returns a boolean if a field has been set.

### SetExecutionTimeNil

`func (o *WorkflowEvent) SetExecutionTimeNil(b bool)`

 SetExecutionTimeNil sets the value for ExecutionTime to be an explicit nil

### UnsetExecutionTime
`func (o *WorkflowEvent) UnsetExecutionTime()`

UnsetExecutionTime ensures that no value is present for ExecutionTime, not even an explicit nil
### GetDateTestedUTC

`func (o *WorkflowEvent) GetDateTestedUTC() time.Time`

GetDateTestedUTC returns the DateTestedUTC field if non-nil, zero value otherwise.

### GetDateTestedUTCOk

`func (o *WorkflowEvent) GetDateTestedUTCOk() (*time.Time, bool)`

GetDateTestedUTCOk returns a tuple with the DateTestedUTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTestedUTC

`func (o *WorkflowEvent) SetDateTestedUTC(v time.Time)`

SetDateTestedUTC sets DateTestedUTC field to given value.

### HasDateTestedUTC

`func (o *WorkflowEvent) HasDateTestedUTC() bool`

HasDateTestedUTC returns a boolean if a field has been set.

### GetTestRecordsMatched

`func (o *WorkflowEvent) GetTestRecordsMatched() int32`

GetTestRecordsMatched returns the TestRecordsMatched field if non-nil, zero value otherwise.

### GetTestRecordsMatchedOk

`func (o *WorkflowEvent) GetTestRecordsMatchedOk() (*int32, bool)`

GetTestRecordsMatchedOk returns a tuple with the TestRecordsMatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRecordsMatched

`func (o *WorkflowEvent) SetTestRecordsMatched(v int32)`

SetTestRecordsMatched sets TestRecordsMatched field to given value.

### HasTestRecordsMatched

`func (o *WorkflowEvent) HasTestRecordsMatched() bool`

HasTestRecordsMatched returns a boolean if a field has been set.

### SetTestRecordsMatchedNil

`func (o *WorkflowEvent) SetTestRecordsMatchedNil(b bool)`

 SetTestRecordsMatchedNil sets the value for TestRecordsMatched to be an explicit nil

### UnsetTestRecordsMatched
`func (o *WorkflowEvent) UnsetTestRecordsMatched()`

UnsetTestRecordsMatched ensures that no value is present for TestRecordsMatched, not even an explicit nil
### GetConnectWiseID

`func (o *WorkflowEvent) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowEvent) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowEvent) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowEvent) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowEvent) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowEvent) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowEvent) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowEvent) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowEvent) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowEvent) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowEvent) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowEvent) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowEvent) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowEvent) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowEvent) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowEvent) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowEvent) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowEvent) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


