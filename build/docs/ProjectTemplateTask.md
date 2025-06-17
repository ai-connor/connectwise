# ProjectTemplateTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TicketId** | Pointer to **NullableInt32** |  | [optional] 
**Sequence** | Pointer to **NullableInt32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ConnectWiseId** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **NullableInt32** |  | [optional] 
**GrandParentId** | Pointer to **NullableInt32** |  | [optional] 
**ParentConnectWiseId** | Pointer to **string** |  | [optional] 
**GrandParentConnectWiseId** | Pointer to **string** |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**Code** | Pointer to [**ServiceCodeReference**](ServiceCodeReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewProjectTemplateTask

`func NewProjectTemplateTask() *ProjectTemplateTask`

NewProjectTemplateTask instantiates a new ProjectTemplateTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectTemplateTaskWithDefaults

`func NewProjectTemplateTaskWithDefaults() *ProjectTemplateTask`

NewProjectTemplateTaskWithDefaults instantiates a new ProjectTemplateTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectTemplateTask) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectTemplateTask) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectTemplateTask) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectTemplateTask) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTicketId

`func (o *ProjectTemplateTask) GetTicketId() int32`

GetTicketId returns the TicketId field if non-nil, zero value otherwise.

### GetTicketIdOk

`func (o *ProjectTemplateTask) GetTicketIdOk() (*int32, bool)`

GetTicketIdOk returns a tuple with the TicketId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketId

`func (o *ProjectTemplateTask) SetTicketId(v int32)`

SetTicketId sets TicketId field to given value.

### HasTicketId

`func (o *ProjectTemplateTask) HasTicketId() bool`

HasTicketId returns a boolean if a field has been set.

### SetTicketIdNil

`func (o *ProjectTemplateTask) SetTicketIdNil(b bool)`

 SetTicketIdNil sets the value for TicketId to be an explicit nil

### UnsetTicketId
`func (o *ProjectTemplateTask) UnsetTicketId()`

UnsetTicketId ensures that no value is present for TicketId, not even an explicit nil
### GetSequence

`func (o *ProjectTemplateTask) GetSequence() int32`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *ProjectTemplateTask) GetSequenceOk() (*int32, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *ProjectTemplateTask) SetSequence(v int32)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *ProjectTemplateTask) HasSequence() bool`

HasSequence returns a boolean if a field has been set.

### SetSequenceNil

`func (o *ProjectTemplateTask) SetSequenceNil(b bool)`

 SetSequenceNil sets the value for Sequence to be an explicit nil

### UnsetSequence
`func (o *ProjectTemplateTask) UnsetSequence()`

UnsetSequence ensures that no value is present for Sequence, not even an explicit nil
### GetDescription

`func (o *ProjectTemplateTask) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectTemplateTask) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectTemplateTask) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectTemplateTask) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConnectWiseId

`func (o *ProjectTemplateTask) GetConnectWiseId() string`

GetConnectWiseId returns the ConnectWiseId field if non-nil, zero value otherwise.

### GetConnectWiseIdOk

`func (o *ProjectTemplateTask) GetConnectWiseIdOk() (*string, bool)`

GetConnectWiseIdOk returns a tuple with the ConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseId

`func (o *ProjectTemplateTask) SetConnectWiseId(v string)`

SetConnectWiseId sets ConnectWiseId field to given value.

### HasConnectWiseId

`func (o *ProjectTemplateTask) HasConnectWiseId() bool`

HasConnectWiseId returns a boolean if a field has been set.

### GetParentId

`func (o *ProjectTemplateTask) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProjectTemplateTask) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProjectTemplateTask) SetParentId(v int32)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProjectTemplateTask) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *ProjectTemplateTask) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *ProjectTemplateTask) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetGrandParentId

`func (o *ProjectTemplateTask) GetGrandParentId() int32`

GetGrandParentId returns the GrandParentId field if non-nil, zero value otherwise.

### GetGrandParentIdOk

`func (o *ProjectTemplateTask) GetGrandParentIdOk() (*int32, bool)`

GetGrandParentIdOk returns a tuple with the GrandParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentId

`func (o *ProjectTemplateTask) SetGrandParentId(v int32)`

SetGrandParentId sets GrandParentId field to given value.

### HasGrandParentId

`func (o *ProjectTemplateTask) HasGrandParentId() bool`

HasGrandParentId returns a boolean if a field has been set.

### SetGrandParentIdNil

`func (o *ProjectTemplateTask) SetGrandParentIdNil(b bool)`

 SetGrandParentIdNil sets the value for GrandParentId to be an explicit nil

### UnsetGrandParentId
`func (o *ProjectTemplateTask) UnsetGrandParentId()`

UnsetGrandParentId ensures that no value is present for GrandParentId, not even an explicit nil
### GetParentConnectWiseId

`func (o *ProjectTemplateTask) GetParentConnectWiseId() string`

GetParentConnectWiseId returns the ParentConnectWiseId field if non-nil, zero value otherwise.

### GetParentConnectWiseIdOk

`func (o *ProjectTemplateTask) GetParentConnectWiseIdOk() (*string, bool)`

GetParentConnectWiseIdOk returns a tuple with the ParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConnectWiseId

`func (o *ProjectTemplateTask) SetParentConnectWiseId(v string)`

SetParentConnectWiseId sets ParentConnectWiseId field to given value.

### HasParentConnectWiseId

`func (o *ProjectTemplateTask) HasParentConnectWiseId() bool`

HasParentConnectWiseId returns a boolean if a field has been set.

### GetGrandParentConnectWiseId

`func (o *ProjectTemplateTask) GetGrandParentConnectWiseId() string`

GetGrandParentConnectWiseId returns the GrandParentConnectWiseId field if non-nil, zero value otherwise.

### GetGrandParentConnectWiseIdOk

`func (o *ProjectTemplateTask) GetGrandParentConnectWiseIdOk() (*string, bool)`

GetGrandParentConnectWiseIdOk returns a tuple with the GrandParentConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandParentConnectWiseId

`func (o *ProjectTemplateTask) SetGrandParentConnectWiseId(v string)`

SetGrandParentConnectWiseId sets GrandParentConnectWiseId field to given value.

### HasGrandParentConnectWiseId

`func (o *ProjectTemplateTask) HasGrandParentConnectWiseId() bool`

HasGrandParentConnectWiseId returns a boolean if a field has been set.

### GetSummary

`func (o *ProjectTemplateTask) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ProjectTemplateTask) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ProjectTemplateTask) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ProjectTemplateTask) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetCode

`func (o *ProjectTemplateTask) GetCode() ServiceCodeReference`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ProjectTemplateTask) GetCodeOk() (*ServiceCodeReference, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ProjectTemplateTask) SetCode(v ServiceCodeReference)`

SetCode sets Code field to given value.

### HasCode

`func (o *ProjectTemplateTask) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInfo

`func (o *ProjectTemplateTask) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProjectTemplateTask) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProjectTemplateTask) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ProjectTemplateTask) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


