# ResultInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**OriginalIndex** | Pointer to **int32** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Data** | Pointer to [**IRestIdentifiedItem**](IRestIdentifiedItem.md) |  | [optional] 
**Error** | Pointer to [**ErrorResponseMessage**](ErrorResponseMessage.md) |  | [optional] 

## Methods

### NewResultInfo

`func NewResultInfo() *ResultInfo`

NewResultInfo instantiates a new ResultInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultInfoWithDefaults

`func NewResultInfoWithDefaults() *ResultInfo`

NewResultInfoWithDefaults instantiates a new ResultInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ResultInfo) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ResultInfo) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ResultInfo) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ResultInfo) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetOriginalIndex

`func (o *ResultInfo) GetOriginalIndex() int32`

GetOriginalIndex returns the OriginalIndex field if non-nil, zero value otherwise.

### GetOriginalIndexOk

`func (o *ResultInfo) GetOriginalIndexOk() (*int32, bool)`

GetOriginalIndexOk returns a tuple with the OriginalIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalIndex

`func (o *ResultInfo) SetOriginalIndex(v int32)`

SetOriginalIndex sets OriginalIndex field to given value.

### HasOriginalIndex

`func (o *ResultInfo) HasOriginalIndex() bool`

HasOriginalIndex returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResultInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResultInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResultInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResultInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetData

`func (o *ResultInfo) GetData() IRestIdentifiedItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResultInfo) GetDataOk() (*IRestIdentifiedItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResultInfo) SetData(v IRestIdentifiedItem)`

SetData sets Data field to given value.

### HasData

`func (o *ResultInfo) HasData() bool`

HasData returns a boolean if a field has been set.

### GetError

`func (o *ResultInfo) GetError() ErrorResponseMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ResultInfo) GetErrorOk() (*ErrorResponseMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ResultInfo) SetError(v ErrorResponseMessage)`

SetError sets Error field to given value.

### HasError

`func (o *ResultInfo) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


