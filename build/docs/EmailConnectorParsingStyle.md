# EmailConnectorParsingStyle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ParsingType** | [**EmailConnectorParsingTypeReference**](EmailConnectorParsingTypeReference.md) |  | 
**ParseRule** | **string** |  Max length: 500; | 
**Priority** | **NullableInt32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewEmailConnectorParsingStyle

`func NewEmailConnectorParsingStyle(parsingType EmailConnectorParsingTypeReference, parseRule string, priority NullableInt32, ) *EmailConnectorParsingStyle`

NewEmailConnectorParsingStyle instantiates a new EmailConnectorParsingStyle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailConnectorParsingStyleWithDefaults

`func NewEmailConnectorParsingStyleWithDefaults() *EmailConnectorParsingStyle`

NewEmailConnectorParsingStyleWithDefaults instantiates a new EmailConnectorParsingStyle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailConnectorParsingStyle) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailConnectorParsingStyle) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailConnectorParsingStyle) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailConnectorParsingStyle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParsingType

`func (o *EmailConnectorParsingStyle) GetParsingType() EmailConnectorParsingTypeReference`

GetParsingType returns the ParsingType field if non-nil, zero value otherwise.

### GetParsingTypeOk

`func (o *EmailConnectorParsingStyle) GetParsingTypeOk() (*EmailConnectorParsingTypeReference, bool)`

GetParsingTypeOk returns a tuple with the ParsingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingType

`func (o *EmailConnectorParsingStyle) SetParsingType(v EmailConnectorParsingTypeReference)`

SetParsingType sets ParsingType field to given value.


### GetParseRule

`func (o *EmailConnectorParsingStyle) GetParseRule() string`

GetParseRule returns the ParseRule field if non-nil, zero value otherwise.

### GetParseRuleOk

`func (o *EmailConnectorParsingStyle) GetParseRuleOk() (*string, bool)`

GetParseRuleOk returns a tuple with the ParseRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseRule

`func (o *EmailConnectorParsingStyle) SetParseRule(v string)`

SetParseRule sets ParseRule field to given value.


### GetPriority

`func (o *EmailConnectorParsingStyle) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *EmailConnectorParsingStyle) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *EmailConnectorParsingStyle) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *EmailConnectorParsingStyle) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *EmailConnectorParsingStyle) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil
### GetInfo

`func (o *EmailConnectorParsingStyle) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *EmailConnectorParsingStyle) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *EmailConnectorParsingStyle) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *EmailConnectorParsingStyle) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


