# CustomFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Caption** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**EntryMethod** | Pointer to **NullableString** |  | [optional] 
**NumberOfDecimals** | Pointer to **NullableInt32** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectWiseId** | Pointer to **string** |  | [optional] 

## Methods

### NewCustomFieldValue

`func NewCustomFieldValue() *CustomFieldValue`

NewCustomFieldValue instantiates a new CustomFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldValueWithDefaults

`func NewCustomFieldValueWithDefaults() *CustomFieldValue`

NewCustomFieldValueWithDefaults instantiates a new CustomFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomFieldValue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomFieldValue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomFieldValue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomFieldValue) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CustomFieldValue) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CustomFieldValue) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetCaption

`func (o *CustomFieldValue) GetCaption() string`

GetCaption returns the Caption field if non-nil, zero value otherwise.

### GetCaptionOk

`func (o *CustomFieldValue) GetCaptionOk() (*string, bool)`

GetCaptionOk returns a tuple with the Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaption

`func (o *CustomFieldValue) SetCaption(v string)`

SetCaption sets Caption field to given value.

### HasCaption

`func (o *CustomFieldValue) HasCaption() bool`

HasCaption returns a boolean if a field has been set.

### GetType

`func (o *CustomFieldValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFieldValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFieldValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CustomFieldValue) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CustomFieldValue) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CustomFieldValue) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetEntryMethod

`func (o *CustomFieldValue) GetEntryMethod() string`

GetEntryMethod returns the EntryMethod field if non-nil, zero value otherwise.

### GetEntryMethodOk

`func (o *CustomFieldValue) GetEntryMethodOk() (*string, bool)`

GetEntryMethodOk returns a tuple with the EntryMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryMethod

`func (o *CustomFieldValue) SetEntryMethod(v string)`

SetEntryMethod sets EntryMethod field to given value.

### HasEntryMethod

`func (o *CustomFieldValue) HasEntryMethod() bool`

HasEntryMethod returns a boolean if a field has been set.

### SetEntryMethodNil

`func (o *CustomFieldValue) SetEntryMethodNil(b bool)`

 SetEntryMethodNil sets the value for EntryMethod to be an explicit nil

### UnsetEntryMethod
`func (o *CustomFieldValue) UnsetEntryMethod()`

UnsetEntryMethod ensures that no value is present for EntryMethod, not even an explicit nil
### GetNumberOfDecimals

`func (o *CustomFieldValue) GetNumberOfDecimals() int32`

GetNumberOfDecimals returns the NumberOfDecimals field if non-nil, zero value otherwise.

### GetNumberOfDecimalsOk

`func (o *CustomFieldValue) GetNumberOfDecimalsOk() (*int32, bool)`

GetNumberOfDecimalsOk returns a tuple with the NumberOfDecimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDecimals

`func (o *CustomFieldValue) SetNumberOfDecimals(v int32)`

SetNumberOfDecimals sets NumberOfDecimals field to given value.

### HasNumberOfDecimals

`func (o *CustomFieldValue) HasNumberOfDecimals() bool`

HasNumberOfDecimals returns a boolean if a field has been set.

### SetNumberOfDecimalsNil

`func (o *CustomFieldValue) SetNumberOfDecimalsNil(b bool)`

 SetNumberOfDecimalsNil sets the value for NumberOfDecimals to be an explicit nil

### UnsetNumberOfDecimals
`func (o *CustomFieldValue) UnsetNumberOfDecimals()`

UnsetNumberOfDecimals ensures that no value is present for NumberOfDecimals, not even an explicit nil
### GetValue

`func (o *CustomFieldValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CustomFieldValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CustomFieldValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *CustomFieldValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetConnectWiseId

`func (o *CustomFieldValue) GetConnectWiseId() string`

GetConnectWiseId returns the ConnectWiseId field if non-nil, zero value otherwise.

### GetConnectWiseIdOk

`func (o *CustomFieldValue) GetConnectWiseIdOk() (*string, bool)`

GetConnectWiseIdOk returns a tuple with the ConnectWiseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseId

`func (o *CustomFieldValue) SetConnectWiseId(v string)`

SetConnectWiseId sets ConnectWiseId field to given value.

### HasConnectWiseId

`func (o *CustomFieldValue) HasConnectWiseId() bool`

HasConnectWiseId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


