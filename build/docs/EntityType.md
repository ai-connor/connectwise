# EntityType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewEntityType

`func NewEntityType() *EntityType`

NewEntityType instantiates a new EntityType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityTypeWithDefaults

`func NewEntityTypeWithDefaults() *EntityType`

NewEntityTypeWithDefaults instantiates a new EntityType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EntityType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntityType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntityType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *EntityType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *EntityType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *EntityType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *EntityType) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


