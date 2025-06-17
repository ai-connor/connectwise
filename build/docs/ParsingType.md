# ParsingType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParseRule** | Pointer to **string** |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewParsingType

`func NewParsingType() *ParsingType`

NewParsingType instantiates a new ParsingType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParsingTypeWithDefaults

`func NewParsingTypeWithDefaults() *ParsingType`

NewParsingTypeWithDefaults instantiates a new ParsingType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ParsingType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ParsingType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ParsingType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ParsingType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ParsingType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ParsingType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ParsingType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ParsingType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParseRule

`func (o *ParsingType) GetParseRule() string`

GetParseRule returns the ParseRule field if non-nil, zero value otherwise.

### GetParseRuleOk

`func (o *ParsingType) GetParseRuleOk() (*string, bool)`

GetParseRuleOk returns a tuple with the ParseRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseRule

`func (o *ParsingType) SetParseRule(v string)`

SetParseRule sets ParseRule field to given value.

### HasParseRule

`func (o *ParsingType) HasParseRule() bool`

HasParseRule returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *ParsingType) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *ParsingType) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *ParsingType) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *ParsingType) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *ParsingType) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *ParsingType) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *ParsingType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ParsingType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ParsingType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ParsingType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


