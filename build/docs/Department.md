# Department

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 15; | 
**Name** | **string** |  Max length: 50; | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDepartment

`func NewDepartment(identifier string, name string, ) *Department`

NewDepartment instantiates a new Department object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentWithDefaults

`func NewDepartmentWithDefaults() *Department`

NewDepartmentWithDefaults instantiates a new Department object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Department) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Department) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Department) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Department) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *Department) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Department) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Department) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *Department) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Department) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Department) SetName(v string)`

SetName sets Name field to given value.


### GetInfo

`func (o *Department) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Department) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Department) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Department) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


