# ServiceSignoffCustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SequenceNumber** | **NullableFloat64** |  | 
**DisplaySection** | **NullableString** |  | 
**UserDefinedField** | [**UserDefinedFieldReference**](UserDefinedFieldReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewServiceSignoffCustomField

`func NewServiceSignoffCustomField(sequenceNumber NullableFloat64, displaySection NullableString, userDefinedField UserDefinedFieldReference, ) *ServiceSignoffCustomField`

NewServiceSignoffCustomField instantiates a new ServiceSignoffCustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceSignoffCustomFieldWithDefaults

`func NewServiceSignoffCustomFieldWithDefaults() *ServiceSignoffCustomField`

NewServiceSignoffCustomFieldWithDefaults instantiates a new ServiceSignoffCustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceSignoffCustomField) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceSignoffCustomField) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceSignoffCustomField) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceSignoffCustomField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSequenceNumber

`func (o *ServiceSignoffCustomField) GetSequenceNumber() float64`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *ServiceSignoffCustomField) GetSequenceNumberOk() (*float64, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *ServiceSignoffCustomField) SetSequenceNumber(v float64)`

SetSequenceNumber sets SequenceNumber field to given value.


### SetSequenceNumberNil

`func (o *ServiceSignoffCustomField) SetSequenceNumberNil(b bool)`

 SetSequenceNumberNil sets the value for SequenceNumber to be an explicit nil

### UnsetSequenceNumber
`func (o *ServiceSignoffCustomField) UnsetSequenceNumber()`

UnsetSequenceNumber ensures that no value is present for SequenceNumber, not even an explicit nil
### GetDisplaySection

`func (o *ServiceSignoffCustomField) GetDisplaySection() string`

GetDisplaySection returns the DisplaySection field if non-nil, zero value otherwise.

### GetDisplaySectionOk

`func (o *ServiceSignoffCustomField) GetDisplaySectionOk() (*string, bool)`

GetDisplaySectionOk returns a tuple with the DisplaySection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaySection

`func (o *ServiceSignoffCustomField) SetDisplaySection(v string)`

SetDisplaySection sets DisplaySection field to given value.


### SetDisplaySectionNil

`func (o *ServiceSignoffCustomField) SetDisplaySectionNil(b bool)`

 SetDisplaySectionNil sets the value for DisplaySection to be an explicit nil

### UnsetDisplaySection
`func (o *ServiceSignoffCustomField) UnsetDisplaySection()`

UnsetDisplaySection ensures that no value is present for DisplaySection, not even an explicit nil
### GetUserDefinedField

`func (o *ServiceSignoffCustomField) GetUserDefinedField() UserDefinedFieldReference`

GetUserDefinedField returns the UserDefinedField field if non-nil, zero value otherwise.

### GetUserDefinedFieldOk

`func (o *ServiceSignoffCustomField) GetUserDefinedFieldOk() (*UserDefinedFieldReference, bool)`

GetUserDefinedFieldOk returns a tuple with the UserDefinedField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField

`func (o *ServiceSignoffCustomField) SetUserDefinedField(v UserDefinedFieldReference)`

SetUserDefinedField sets UserDefinedField field to given value.


### GetInfo

`func (o *ServiceSignoffCustomField) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ServiceSignoffCustomField) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ServiceSignoffCustomField) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ServiceSignoffCustomField) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


