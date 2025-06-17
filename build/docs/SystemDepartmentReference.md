# SystemDepartmentReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewSystemDepartmentReference

`func NewSystemDepartmentReference() *SystemDepartmentReference`

NewSystemDepartmentReference instantiates a new SystemDepartmentReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemDepartmentReferenceWithDefaults

`func NewSystemDepartmentReferenceWithDefaults() *SystemDepartmentReference`

NewSystemDepartmentReferenceWithDefaults instantiates a new SystemDepartmentReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemDepartmentReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemDepartmentReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemDepartmentReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SystemDepartmentReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SystemDepartmentReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SystemDepartmentReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *SystemDepartmentReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SystemDepartmentReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SystemDepartmentReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *SystemDepartmentReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *SystemDepartmentReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemDepartmentReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemDepartmentReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemDepartmentReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetInfo

`func (o *SystemDepartmentReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *SystemDepartmentReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *SystemDepartmentReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *SystemDepartmentReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


