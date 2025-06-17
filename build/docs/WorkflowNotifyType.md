# WorkflowNotifyType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**IsSetupFlag** | Pointer to **NullableBool** | If the current action is available because it is already set up. Pertains to integrations such as Automate | [optional] 
**ExternalFlag** | Pointer to **NullableBool** | If the current action effects external objects e.g. integrations or sending an email | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** | WF_NotifyHeader_RecID | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkflowNotifyType

`func NewWorkflowNotifyType() *WorkflowNotifyType`

NewWorkflowNotifyType instantiates a new WorkflowNotifyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowNotifyTypeWithDefaults

`func NewWorkflowNotifyTypeWithDefaults() *WorkflowNotifyType`

NewWorkflowNotifyTypeWithDefaults instantiates a new WorkflowNotifyType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowNotifyType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowNotifyType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowNotifyType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowNotifyType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *WorkflowNotifyType) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *WorkflowNotifyType) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *WorkflowNotifyType) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *WorkflowNotifyType) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *WorkflowNotifyType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowNotifyType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowNotifyType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowNotifyType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsSetupFlag

`func (o *WorkflowNotifyType) GetIsSetupFlag() bool`

GetIsSetupFlag returns the IsSetupFlag field if non-nil, zero value otherwise.

### GetIsSetupFlagOk

`func (o *WorkflowNotifyType) GetIsSetupFlagOk() (*bool, bool)`

GetIsSetupFlagOk returns a tuple with the IsSetupFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSetupFlag

`func (o *WorkflowNotifyType) SetIsSetupFlag(v bool)`

SetIsSetupFlag sets IsSetupFlag field to given value.

### HasIsSetupFlag

`func (o *WorkflowNotifyType) HasIsSetupFlag() bool`

HasIsSetupFlag returns a boolean if a field has been set.

### SetIsSetupFlagNil

`func (o *WorkflowNotifyType) SetIsSetupFlagNil(b bool)`

 SetIsSetupFlagNil sets the value for IsSetupFlag to be an explicit nil

### UnsetIsSetupFlag
`func (o *WorkflowNotifyType) UnsetIsSetupFlag()`

UnsetIsSetupFlag ensures that no value is present for IsSetupFlag, not even an explicit nil
### GetExternalFlag

`func (o *WorkflowNotifyType) GetExternalFlag() bool`

GetExternalFlag returns the ExternalFlag field if non-nil, zero value otherwise.

### GetExternalFlagOk

`func (o *WorkflowNotifyType) GetExternalFlagOk() (*bool, bool)`

GetExternalFlagOk returns a tuple with the ExternalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalFlag

`func (o *WorkflowNotifyType) SetExternalFlag(v bool)`

SetExternalFlag sets ExternalFlag field to given value.

### HasExternalFlag

`func (o *WorkflowNotifyType) HasExternalFlag() bool`

HasExternalFlag returns a boolean if a field has been set.

### SetExternalFlagNil

`func (o *WorkflowNotifyType) SetExternalFlagNil(b bool)`

 SetExternalFlagNil sets the value for ExternalFlag to be an explicit nil

### UnsetExternalFlag
`func (o *WorkflowNotifyType) UnsetExternalFlag()`

UnsetExternalFlag ensures that no value is present for ExternalFlag, not even an explicit nil
### GetConnectWiseID

`func (o *WorkflowNotifyType) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *WorkflowNotifyType) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *WorkflowNotifyType) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *WorkflowNotifyType) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetParentId

`func (o *WorkflowNotifyType) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *WorkflowNotifyType) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *WorkflowNotifyType) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *WorkflowNotifyType) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *WorkflowNotifyType) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *WorkflowNotifyType) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *WorkflowNotifyType) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *WorkflowNotifyType) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *WorkflowNotifyType) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *WorkflowNotifyType) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetInfo

`func (o *WorkflowNotifyType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkflowNotifyType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkflowNotifyType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkflowNotifyType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


