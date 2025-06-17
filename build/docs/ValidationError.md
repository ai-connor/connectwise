# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 

## Methods

### NewValidationError

`func NewValidationError() *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ValidationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ValidationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResource

`func (o *ValidationError) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ValidationError) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ValidationError) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ValidationError) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetField

`func (o *ValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ValidationError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ValidationError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetDetails

`func (o *ValidationError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ValidationError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ValidationError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ValidationError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


