# PhaseStatusReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPhaseStatusReference

`func NewPhaseStatusReference() *PhaseStatusReference`

NewPhaseStatusReference instantiates a new PhaseStatusReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhaseStatusReferenceWithDefaults

`func NewPhaseStatusReferenceWithDefaults() *PhaseStatusReference`

NewPhaseStatusReferenceWithDefaults instantiates a new PhaseStatusReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhaseStatusReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhaseStatusReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhaseStatusReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PhaseStatusReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *PhaseStatusReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *PhaseStatusReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *PhaseStatusReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhaseStatusReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhaseStatusReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhaseStatusReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *PhaseStatusReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *PhaseStatusReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *PhaseStatusReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *PhaseStatusReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


