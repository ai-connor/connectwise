# HttpResponseMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Content** | Pointer to [**HttpContent**](HttpContent.md) |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**ReasonPhrase** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **[]interface{}** |  | [optional] 
**RequestMessage** | Pointer to [**HttpRequestMessage**](HttpRequestMessage.md) |  | [optional] 
**IsSuccessStatusCode** | Pointer to **bool** |  | [optional] 

## Methods

### NewHttpResponseMessage

`func NewHttpResponseMessage() *HttpResponseMessage`

NewHttpResponseMessage instantiates a new HttpResponseMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpResponseMessageWithDefaults

`func NewHttpResponseMessageWithDefaults() *HttpResponseMessage`

NewHttpResponseMessageWithDefaults instantiates a new HttpResponseMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HttpResponseMessage) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HttpResponseMessage) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HttpResponseMessage) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HttpResponseMessage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetContent

`func (o *HttpResponseMessage) GetContent() HttpContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HttpResponseMessage) GetContentOk() (*HttpContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HttpResponseMessage) SetContent(v HttpContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *HttpResponseMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStatusCode

`func (o *HttpResponseMessage) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HttpResponseMessage) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HttpResponseMessage) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HttpResponseMessage) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetReasonPhrase

`func (o *HttpResponseMessage) GetReasonPhrase() string`

GetReasonPhrase returns the ReasonPhrase field if non-nil, zero value otherwise.

### GetReasonPhraseOk

`func (o *HttpResponseMessage) GetReasonPhraseOk() (*string, bool)`

GetReasonPhraseOk returns a tuple with the ReasonPhrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonPhrase

`func (o *HttpResponseMessage) SetReasonPhrase(v string)`

SetReasonPhrase sets ReasonPhrase field to given value.

### HasReasonPhrase

`func (o *HttpResponseMessage) HasReasonPhrase() bool`

HasReasonPhrase returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpResponseMessage) GetHeaders() []interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpResponseMessage) GetHeadersOk() (*[]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpResponseMessage) SetHeaders(v []interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpResponseMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRequestMessage

`func (o *HttpResponseMessage) GetRequestMessage() HttpRequestMessage`

GetRequestMessage returns the RequestMessage field if non-nil, zero value otherwise.

### GetRequestMessageOk

`func (o *HttpResponseMessage) GetRequestMessageOk() (*HttpRequestMessage, bool)`

GetRequestMessageOk returns a tuple with the RequestMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMessage

`func (o *HttpResponseMessage) SetRequestMessage(v HttpRequestMessage)`

SetRequestMessage sets RequestMessage field to given value.

### HasRequestMessage

`func (o *HttpResponseMessage) HasRequestMessage() bool`

HasRequestMessage returns a boolean if a field has been set.

### GetIsSuccessStatusCode

`func (o *HttpResponseMessage) GetIsSuccessStatusCode() bool`

GetIsSuccessStatusCode returns the IsSuccessStatusCode field if non-nil, zero value otherwise.

### GetIsSuccessStatusCodeOk

`func (o *HttpResponseMessage) GetIsSuccessStatusCodeOk() (*bool, bool)`

GetIsSuccessStatusCodeOk returns a tuple with the IsSuccessStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSuccessStatusCode

`func (o *HttpResponseMessage) SetIsSuccessStatusCode(v bool)`

SetIsSuccessStatusCode sets IsSuccessStatusCode field to given value.

### HasIsSuccessStatusCode

`func (o *HttpResponseMessage) HasIsSuccessStatusCode() bool`

HasIsSuccessStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


