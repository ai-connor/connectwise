# WorkflowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**HasOptionsFlag** | Pointer to **NullableBool** |  | [optional] 
**HasOperatorFlag** | Pointer to **NullableBool** |  | [optional] 
**CustomField** | Pointer to [**UserDefinedFieldReference**](UserDefinedFieldReference.md) |  | [optional] 
**ExpectedType** | Pointer to **string** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowTrigger

`func NewWorkflowTrigger() *WorkflowTrigger`

NewWorkflowTrigger instantiates a new WorkflowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTriggerWithDefaults

`func NewWorkflowTriggerWithDefaults() *WorkflowTrigger`

NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTrigger) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTrigger) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTrigger) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasOptionsFlag

`func (o *WorkflowTrigger) GetHasOptionsFlag() bool`

GetHasOptionsFlag returns the HasOptionsFlag field if non-nil, zero value otherwise.

### GetHasOptionsFlagOk

`func (o *WorkflowTrigger) GetHasOptionsFlagOk() (*bool, bool)`

GetHasOptionsFlagOk returns a tuple with the HasOptionsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOptionsFlag

`func (o *WorkflowTrigger) SetHasOptionsFlag(v bool)`

SetHasOptionsFlag sets HasOptionsFlag field to given value.

### HasHasOptionsFlag

`func (o *WorkflowTrigger) HasHasOptionsFlag() bool`

HasHasOptionsFlag returns a boolean if a field has been set.

### SetHasOptionsFlagNil

`func (o *WorkflowTrigger) SetHasOptionsFlagNil(b bool)`

 SetHasOptionsFlagNil sets the value for HasOptionsFlag to be an explicit nil

### UnsetHasOptionsFlag
`func (o *WorkflowTrigger) UnsetHasOptionsFlag()`

UnsetHasOptionsFlag ensures that no value is present for HasOptionsFlag, not even an explicit nil
### GetHasOperatorFlag

`func (o *WorkflowTrigger) GetHasOperatorFlag() bool`

GetHasOperatorFlag returns the HasOperatorFlag field if non-nil, zero value otherwise.

### GetHasOperatorFlagOk

`func (o *WorkflowTrigger) GetHasOperatorFlagOk() (*bool, bool)`

GetHasOperatorFlagOk returns a tuple with the HasOperatorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOperatorFlag

`func (o *WorkflowTrigger) SetHasOperatorFlag(v bool)`

SetHasOperatorFlag sets HasOperatorFlag field to given value.

### HasHasOperatorFlag

`func (o *WorkflowTrigger) HasHasOperatorFlag() bool`

HasHasOperatorFlag returns a boolean if a field has been set.

### SetHasOperatorFlagNil

`func (o *WorkflowTrigger) SetHasOperatorFlagNil(b bool)`

 SetHasOperatorFlagNil sets the value for HasOperatorFlag to be an explicit nil

### UnsetHasOperatorFlag
`func (o *WorkflowTrigger) UnsetHasOperatorFlag()`

UnsetHasOperatorFlag ensures that no value is present for HasOperatorFlag, not even an explicit nil
### GetCustomField

`func (o *WorkflowTrigger) GetCustomField() UserDefinedFieldReference`

GetCustomField returns the CustomField field if non-nil, zero value otherwise.

### GetCustomFieldOk

`func (o *WorkflowTrigger) GetCustomFieldOk() (*UserDefinedFieldReference, bool)`

GetCustomFieldOk returns a tuple with the CustomField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField

`func (o *WorkflowTrigger) SetCustomField(v UserDefinedFieldReference)`

SetCustomField sets CustomField field to given value.

### HasCustomField

`func (o *WorkflowTrigger) HasCustomField() bool`

HasCustomField returns a boolean if a field has been set.

### GetExpectedType

`func (o *WorkflowTrigger) GetExpectedType() string`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *WorkflowTrigger) GetExpectedTypeOk() (*string, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *WorkflowTrigger) SetExpectedType(v string)`

SetExpectedType sets ExpectedType field to given value.

### HasExpectedType

`func (o *WorkflowTrigger) HasExpectedType() bool`

HasExpectedType returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *WorkflowTrigger) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowTrigger) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowTrigger) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowTrigger) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowTrigger) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowTrigger) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowTrigger) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowTrigger) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowTrigger) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowTrigger) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowTrigger) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowTrigger) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowTrigger) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowTrigger) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowTrigger) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowTrigger) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowTrigger) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowTrigger) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


