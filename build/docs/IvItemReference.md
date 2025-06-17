# IvItemReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**SerializedFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewIvItemReference

`func NewIvItemReference() *IvItemReference`

NewIvItemReference instantiates a new IvItemReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIvItemReferenceWithDefaults

`func NewIvItemReferenceWithDefaults() *IvItemReference`

NewIvItemReferenceWithDefaults instantiates a new IvItemReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IvItemReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IvItemReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IvItemReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *IvItemReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *IvItemReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *IvItemReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *IvItemReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *IvItemReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *IvItemReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *IvItemReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetSerializedFlag

`func (o *IvItemReference) GetSerializedFlag() bool`

GetSerializedFlag returns the SerializedFlag field if non-nil, zero value otherwise.

### GetSerializedFlagOk

`func (o *IvItemReference) GetSerializedFlagOk() (*bool, bool)`

GetSerializedFlagOk returns a tuple with the SerializedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerializedFlag

`func (o *IvItemReference) SetSerializedFlag(v bool)`

SetSerializedFlag sets SerializedFlag field to given value.

### HasSerializedFlag

`func (o *IvItemReference) HasSerializedFlag() bool`

HasSerializedFlag returns a boolean if a field has been set.

### SetSerializedFlagNil

`func (o *IvItemReference) SetSerializedFlagNil(b bool)`

 SetSerializedFlagNil sets the value for SerializedFlag to be an explicit nil

### UnsetSerializedFlag
`func (o *IvItemReference) UnsetSerializedFlag()`

UnsetSerializedFlag ensures that no value is present for SerializedFlag, not even an explicit nil
### GetInfo

`func (o *IvItemReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *IvItemReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *IvItemReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *IvItemReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


