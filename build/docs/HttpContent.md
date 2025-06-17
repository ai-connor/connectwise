# HttpContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewHttpContent

`func NewHttpContent() *HttpContent`

NewHttpContent instantiates a new HttpContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpContentWithDefaults

`func NewHttpContentWithDefaults() *HttpContent`

NewHttpContentWithDefaults instantiates a new HttpContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *HttpContent) GetHeaders() []interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpContent) GetHeadersOk() (*[]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpContent) SetHeaders(v []interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpContent) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


