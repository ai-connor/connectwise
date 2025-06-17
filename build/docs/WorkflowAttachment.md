# WorkflowAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyHeader_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowAttachment

`func NewWorkflowAttachment() *WorkflowAttachment`

NewWorkflowAttachment instantiates a new WorkflowAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowAttachmentWithDefaults

`func NewWorkflowAttachmentWithDefaults() *WorkflowAttachment`

NewWorkflowAttachmentWithDefaults instantiates a new WorkflowAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowAttachment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowAttachment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowAttachment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowAttachment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkflowAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConnectWiseID

`func (o *WorkflowAttachment) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowAttachment) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowAttachment) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowAttachment) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowAttachment) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowAttachment) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowAttachment) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowAttachment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowAttachment) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowAttachment) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowAttachment) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowAttachment) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowAttachment) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowAttachment) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowAttachment) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowAttachment) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowAttachment) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowAttachment) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


