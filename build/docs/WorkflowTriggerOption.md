# WorkflowTriggerOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**CustomField** | Pointer to [**UserDefinedFieldReference**](UserDefinedFieldReference.md) |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**GrandParentId** | Pointer to **NullableInt32** |  | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**GrandParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowTriggerOption

`func NewWorkflowTriggerOption() *WorkflowTriggerOption`

NewWorkflowTriggerOption instantiates a new WorkflowTriggerOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTriggerOptionWithDefaults

`func NewWorkflowTriggerOptionWithDefaults() *WorkflowTriggerOption`

NewWorkflowTriggerOptionWithDefaults instantiates a new WorkflowTriggerOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *WorkflowTriggerOption) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WorkflowTriggerOption) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WorkflowTriggerOption) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *WorkflowTriggerOption) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTriggerOption) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTriggerOption) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTriggerOption) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTriggerOption) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCustomField

`func (o *WorkflowTriggerOption) GetCustomField() UserDefinedFieldReference`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *WorkflowTriggerOption) GetCustomFieldOk() (*UserDefinedFieldReference, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *WorkflowTriggerOption) SetCustomField(v UserDefinedFieldReference)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *WorkflowTriggerOption) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *WorkflowTriggerOption) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowTriggerOption) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowTriggerOption) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowTriggerOption) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowTriggerOption) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowTriggerOption) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowTriggerOption) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowTriggerOption) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowTriggerOption) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowTriggerOption) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGrandParentId

`func (o *WorkflowTriggerOption) GetGrandParentId() int32`

GetGrandParentId returns the GrandParentId field if non-nil, zero value otherwise.

### GetGrandParentIdOk

`func (o *WorkflowTriggerOption) GetGrandParentIdOk() (*int32, bool)`

GetGrandParentIdOk returns a tuple with the GrandParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentId

`func (o *WorkflowTriggerOption) SetGrandParentId(v int32)`

SetGrandParentId sets GrandParentId field to given value.

### HasGrandParentId

`func (o *WorkflowTriggerOption) HasGrandParentId() bool`

HasGrandParentId returns a boolean if a field has been set.

### SetGrandParentIdNil

`func (o *WorkflowTriggerOption) SetGrandParentIdNil(b bool)`

 SetGrandParentIdNil sets the value for GrandParentId to be an explicit nil

### UnsetGrandParentId
`func (o *WorkflowTriggerOption) UnsetGrandParentId()`

UnsetGrandParentId ensures that no value is present for GrandParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowTriggerOption) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowTriggerOption) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowTriggerOption) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowTriggerOption) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetGrandParentConnectWiseId

`func (o *WorkflowTriggerOption) GetGrandParentConnectWiseId() string`

GetGrandParentConnectWiseId returns the GrandParentConnectWiseId field if non-nil, zero value otherwise.

### GetGrandParentConnectWiseIdOk

`func (o *WorkflowTriggerOption) GetGrandParentConnectWiseIdOk() (*string, bool)`

GetGrandParentConnectWiseIdOk returns a tuple with the GrandParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentConnectWiseId

`func (o *WorkflowTriggerOption) SetGrandParentConnectWiseId(v string)`

SetGrandParentConnectWiseId sets GrandParentConnectWiseId field to given value.

### HasGrandParentConnectWiseId

`func (o *WorkflowTriggerOption) HasGrandParentConnectWiseId() bool`

HasGrandParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowTriggerOption) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowTriggerOption) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowTriggerOption) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowTriggerOption) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


