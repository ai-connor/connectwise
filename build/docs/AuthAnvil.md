# AuthAnvil

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ServerLocationUrl** | **string** |  | 
**SiteId** | **NullableInt32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAuthAnvil

`func NewAuthAnvil(serverLocationUrl string, siteId NullableInt32, ) *AuthAnvil`

NewAuthAnvil instantiates a new AuthAnvil object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthAnvilWithDefaults

`func NewAuthAnvilWithDefaults() *AuthAnvil`

NewAuthAnvilWithDefaults instantiates a new AuthAnvil object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthAnvil) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthAnvil) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthAnvil) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthAnvil) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServerLocationUrl

`func (o *AuthAnvil) GetServerLocationUrl() string`

GetServerLocationUrl returns the ServerLocationUrl field if non-nil, zero value otherwise.

### GetServerLocationUrlOk

`func (o *AuthAnvil) GetServerLocationUrlOk() (*string, bool)`

GetServerLocationUrlOk returns a tuple with the ServerLocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLocationUrl

`func (o *AuthAnvil) SetServerLocationUrl(v string)`

SetServerLocationUrl sets ServerLocationUrl field to given value.


### GetSiteId

`func (o *AuthAnvil) GetSiteId() int32`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AuthAnvil) GetSiteIdOk() (*int32, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AuthAnvil) SetSiteId(v int32)`

SetSiteId sets SiteId field to given value.


### SetSiteIdNil

`func (o *AuthAnvil) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *AuthAnvil) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetInfo

`func (o *AuthAnvil) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AuthAnvil) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AuthAnvil) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AuthAnvil) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


