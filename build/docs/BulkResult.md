# BulkResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payload** | Pointer to [**[]ResultInfo**](ResultInfo.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewBulkResult

`func NewBulkResult() *BulkResult`

NewBulkResult instantiates a new BulkResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkResultWithDefaults

`func NewBulkResultWithDefaults() *BulkResult`

NewBulkResultWithDefaults instantiates a new BulkResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayload

`func (o *BulkResult) GetPayload() []ResultInfo`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *BulkResult) GetPayloadOk() (*[]ResultInfo, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *BulkResult) SetPayload(v []ResultInfo)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *BulkResult) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetInfo

`func (o *BulkResult) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BulkResult) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BulkResult) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BulkResult) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


