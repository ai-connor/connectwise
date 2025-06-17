# LinkResolveUrlInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceId** | **NullableInt32** |  | 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewLinkResolveUrlInfo

`func NewLinkResolveUrlInfo(referenceId NullableInt32, ) *LinkResolveUrlInfo`

NewLinkResolveUrlInfo instantiates a new LinkResolveUrlInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkResolveUrlInfoWithDefaults

`func NewLinkResolveUrlInfoWithDefaults() *LinkResolveUrlInfo`

NewLinkResolveUrlInfoWithDefaults instantiates a new LinkResolveUrlInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceId

`func (o *LinkResolveUrlInfo) GetReferenceId() int32`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *LinkResolveUrlInfo) GetReferenceIdOk() (*int32, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *LinkResolveUrlInfo) SetReferenceId(v int32)`

SetReferenceId sets ReferenceId field to given value.


### SetReferenceIdNil

`func (o *LinkResolveUrlInfo) SetReferenceIdNil(b bool)`

 SetReferenceIdNil sets the value for ReferenceId to be an explicit nil

### UnsetReferenceId
`func (o *LinkResolveUrlInfo) UnsetReferenceId()`

UnsetReferenceId ensures that no value is present for ReferenceId, not even an explicit nil
### GetUrl

`func (o *LinkResolveUrlInfo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkResolveUrlInfo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkResolveUrlInfo) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LinkResolveUrlInfo) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


