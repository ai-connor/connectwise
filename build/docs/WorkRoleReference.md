# WorkRoleReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkRoleReference

`func NewWorkRoleReference() *WorkRoleReference`

NewWorkRoleReference instantiates a new WorkRoleReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkRoleReferenceWithDefaults

`func NewWorkRoleReferenceWithDefaults() *WorkRoleReference`

NewWorkRoleReferenceWithDefaults instantiates a new WorkRoleReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkRoleReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkRoleReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkRoleReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkRoleReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkRoleReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkRoleReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *WorkRoleReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkRoleReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkRoleReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkRoleReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *WorkRoleReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkRoleReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkRoleReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkRoleReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


