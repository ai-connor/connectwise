# BundleResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Entities** | Pointer to [**[]IRestIdentifiedItem**](IRestIdentifiedItem.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to [**ErrorResponseMessage**](ErrorResponseMessage.md) |  | [optional] 

## Methods

### NewBundleResult

`func NewBundleResult() *BundleResult`

NewBundleResult instantiates a new BundleResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleResultWithDefaults

`func NewBundleResultWithDefaults() *BundleResult`

NewBundleResultWithDefaults instantiates a new BundleResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *BundleResult) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *BundleResult) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *BundleResult) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *BundleResult) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetResourceType

`func (o *BundleResult) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BundleResult) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BundleResult) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BundleResult) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetEntities

`func (o *BundleResult) GetEntities() []IRestIdentifiedItem`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *BundleResult) GetEntitiesOk() (*[]IRestIdentifiedItem, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *BundleResult) SetEntities(v []IRestIdentifiedItem)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *BundleResult) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetCount

`func (o *BundleResult) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BundleResult) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BundleResult) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BundleResult) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetVersion

`func (o *BundleResult) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BundleResult) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BundleResult) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BundleResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSuccess

`func (o *BundleResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BundleResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BundleResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BundleResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatusCode

`func (o *BundleResult) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BundleResult) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BundleResult) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BundleResult) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetError

`func (o *BundleResult) GetError() ErrorResponseMessage`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BundleResult) GetErrorOk() (*ErrorResponseMessage, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BundleResult) SetError(v ErrorResponseMessage)`

SetError sets Error field to given value.

### HasError

`func (o *BundleResult) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


