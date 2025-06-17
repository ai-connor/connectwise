# HttpRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**Version**](Version.md) |  | [optional] 
**Content** | Pointer to [**HttpContent**](HttpContent.md) |  | [optional] 
**Method** | Pointer to [**HttpMethod**](HttpMethod.md) |  | [optional] 
**RequestUri** | Pointer to **string** |  | [optional] 
**Headers** | Pointer to **[]interface{}** |  | [optional] 
**Properties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 

## Methods

### NewHttpRequestMessage

`func NewHttpRequestMessage() *HttpRequestMessage`

NewHttpRequestMessage instantiates a new HttpRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpRequestMessageWithDefaults

`func NewHttpRequestMessageWithDefaults() *HttpRequestMessage`

NewHttpRequestMessageWithDefaults instantiates a new HttpRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *HttpRequestMessage) GetVersion() Version`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HttpRequestMessage) GetVersionOk() (*Version, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HttpRequestMessage) SetVersion(v Version)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HttpRequestMessage) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetContent

`func (o *HttpRequestMessage) GetContent() HttpContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HttpRequestMessage) GetContentOk() (*HttpContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HttpRequestMessage) SetContent(v HttpContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *HttpRequestMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMethod

`func (o *HttpRequestMessage) GetMethod() HttpMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *HttpRequestMessage) GetMethodOk() (*HttpMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *HttpRequestMessage) SetMethod(v HttpMethod)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *HttpRequestMessage) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetRequestUri

`func (o *HttpRequestMessage) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *HttpRequestMessage) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *HttpRequestMessage) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.

### HasRequestUri

`func (o *HttpRequestMessage) HasRequestUri() bool`

HasRequestUri returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpRequestMessage) GetHeaders() []interface{}`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpRequestMessage) GetHeadersOk() (*[]interface{}, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpRequestMessage) SetHeaders(v []interface{})`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpRequestMessage) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetProperties

`func (o *HttpRequestMessage) GetProperties() map[string]map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *HttpRequestMessage) GetPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *HttpRequestMessage) SetProperties(v map[string]map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *HttpRequestMessage) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


