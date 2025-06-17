# BundleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SequenceNumber** | Pointer to **int32** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**ApiRequest** | Pointer to [**ApiRequest**](ApiRequest.md) |  | [optional] 

## Methods

### NewBundleRequest

`func NewBundleRequest() *BundleRequest`

NewBundleRequest instantiates a new BundleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleRequestWithDefaults

`func NewBundleRequestWithDefaults() *BundleRequest`

NewBundleRequestWithDefaults instantiates a new BundleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSequenceNumber

`func (o *BundleRequest) GetSequenceNumber() int32`

GetSequenceNumber returns the SequenceNumber field if non-nil, zero value otherwise.

### GetSequenceNumberOk

`func (o *BundleRequest) GetSequenceNumberOk() (*int32, bool)`

GetSequenceNumberOk returns a tuple with the SequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceNumber

`func (o *BundleRequest) SetSequenceNumber(v int32)`

SetSequenceNumber sets SequenceNumber field to given value.

### HasSequenceNumber

`func (o *BundleRequest) HasSequenceNumber() bool`

HasSequenceNumber returns a boolean if a field has been set.

### GetResourceType

`func (o *BundleRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *BundleRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *BundleRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *BundleRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetVersion

`func (o *BundleRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *BundleRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *BundleRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *BundleRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetApiRequest

`func (o *BundleRequest) GetApiRequest() ApiRequest`

GetApiRequest returns the ApiRequest field if non-nil, zero value otherwise.

### GetApiRequestOk

`func (o *BundleRequest) GetApiRequestOk() (*ApiRequest, bool)`

GetApiRequestOk returns a tuple with the ApiRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRequest

`func (o *BundleRequest) SetApiRequest(v ApiRequest)`

SetApiRequest sets ApiRequest field to given value.

### HasApiRequest

`func (o *BundleRequest) HasApiRequest() bool`

HasApiRequest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


