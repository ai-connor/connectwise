# ClearPickerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewClearPickerRequest

`func NewClearPickerRequest() *ClearPickerRequest`

NewClearPickerRequest instantiates a new ClearPickerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClearPickerRequestWithDefaults

`func NewClearPickerRequestWithDefaults() *ClearPickerRequest`

NewClearPickerRequestWithDefaults instantiates a new ClearPickerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *ClearPickerRequest) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ClearPickerRequest) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ClearPickerRequest) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ClearPickerRequest) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetType

`func (o *ClearPickerRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClearPickerRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClearPickerRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClearPickerRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *ClearPickerRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ClearPickerRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


