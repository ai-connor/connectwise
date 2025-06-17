# ErrorResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]ValidationError**](ValidationError.md) |  | [optional] 

## Methods

### NewErrorResponseMessage

`func NewErrorResponseMessage() *ErrorResponseMessage`

NewErrorResponseMessage instantiates a new ErrorResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseMessageWithDefaults

`func NewErrorResponseMessageWithDefaults() *ErrorResponseMessage`

NewErrorResponseMessageWithDefaults instantiates a new ErrorResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseMessage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseMessage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseMessage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponseMessage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseMessage) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ErrorResponseMessage) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorResponseMessage) GetErrors() []ValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorResponseMessage) GetErrorsOk() (*[]ValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorResponseMessage) SetErrors(v []ValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorResponseMessage) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


