# WorkTypeReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UtilizationFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkTypeReference

`func NewWorkTypeReference() *WorkTypeReference`

NewWorkTypeReference instantiates a new WorkTypeReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkTypeReferenceWithDefaults

`func NewWorkTypeReferenceWithDefaults() *WorkTypeReference`

NewWorkTypeReferenceWithDefaults instantiates a new WorkTypeReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkTypeReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkTypeReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkTypeReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkTypeReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkTypeReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkTypeReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *WorkTypeReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkTypeReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkTypeReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkTypeReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUtilizationFlag

`func (o *WorkTypeReference) GetUtilizationFlag() bool`

GetUtilizationFlag returns the UtilizationFlag field if non-nil, zero value otherwise.

### GetUtilizationFlagOk

`func (o *WorkTypeReference) GetUtilizationFlagOk() (*bool, bool)`

GetUtilizationFlagOk returns a tuple with the UtilizationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationFlag

`func (o *WorkTypeReference) SetUtilizationFlag(v bool)`

SetUtilizationFlag sets UtilizationFlag field to given value.

### HasUtilizationFlag

`func (o *WorkTypeReference) HasUtilizationFlag() bool`

HasUtilizationFlag returns a boolean if a field has been set.

### SetUtilizationFlagNil

`func (o *WorkTypeReference) SetUtilizationFlagNil(b bool)`

 SetUtilizationFlagNil sets the value for UtilizationFlag to be an explicit nil

### UnsetUtilizationFlag
`func (o *WorkTypeReference) UnsetUtilizationFlag()`

UnsetUtilizationFlag ensures that no value is present for UtilizationFlag, not even an explicit nil
### GetInfo

`func (o *WorkTypeReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkTypeReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkTypeReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkTypeReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


