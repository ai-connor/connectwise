# WorkflowActionAutomateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Value** | Pointer to **string** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyActions_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowActionAutomateParameter

`func NewWorkflowActionAutomateParameter(name string, ) *WorkflowActionAutomateParameter`

NewWorkflowActionAutomateParameter instantiates a new WorkflowActionAutomateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowActionAutomateParameterWithDefaults

`func NewWorkflowActionAutomateParameterWithDefaults() *WorkflowActionAutomateParameter`

NewWorkflowActionAutomateParameterWithDefaults instantiates a new WorkflowActionAutomateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowActionAutomateParameter) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowActionAutomateParameter) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowActionAutomateParameter) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowActionAutomateParameter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowActionAutomateParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowActionAutomateParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowActionAutomateParameter) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *WorkflowActionAutomateParameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowActionAutomateParameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowActionAutomateParameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowActionAutomateParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *WorkflowActionAutomateParameter) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowActionAutomateParameter) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowActionAutomateParameter) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowActionAutomateParameter) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowActionAutomateParameter) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowActionAutomateParameter) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowActionAutomateParameter) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowActionAutomateParameter) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowActionAutomateParameter) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowActionAutomateParameter) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowActionAutomateParameter) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowActionAutomateParameter) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowActionAutomateParameter) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowActionAutomateParameter) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowActionAutomateParameter) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowActionAutomateParameter) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowActionAutomateParameter) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowActionAutomateParameter) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


