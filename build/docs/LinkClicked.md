# LinkClicked

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CampaignId** | Pointer to **NullableInt32** |  | [optional] 
**ContactId** | **NullableInt32** |  | 
**DateClicked** | Pointer to **time.Time** |  | [optional] 
**Url** | **string** |  Max length: 2083; | 
**QueryString** | Pointer to **string** |  | [optional] 

## Methods

### NewLinkClicked

`func NewLinkClicked(contactId NullableInt32, url string, ) *LinkClicked`

NewLinkClicked instantiates a new LinkClicked object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkClickedWithDefaults

`func NewLinkClickedWithDefaults() *LinkClicked`

NewLinkClickedWithDefaults instantiates a new LinkClicked object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkClicked) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkClicked) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkClicked) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LinkClicked) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCampaignId

`func (o *LinkClicked) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *LinkClicked) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *LinkClicked) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *LinkClicked) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *LinkClicked) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *LinkClicked) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetContactId

`func (o *LinkClicked) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *LinkClicked) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *LinkClicked) SetContactId(v int32)`

SetContactId sets ContactId field to given value.


### SetContactIdNil

`func (o *LinkClicked) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *LinkClicked) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetDateClicked

`func (o *LinkClicked) GetDateClicked() time.Time`

GetDateClicked returns the DateClicked field if non-nil, zero value otherwise.

### GetDateClickedOk

`func (o *LinkClicked) GetDateClickedOk() (*time.Time, bool)`

GetDateClickedOk returns a tuple with the DateClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateClicked

`func (o *LinkClicked) SetDateClicked(v time.Time)`

SetDateClicked sets DateClicked field to given value.

### HasDateClicked

`func (o *LinkClicked) HasDateClicked() bool`

HasDateClicked returns a boolean if a field has been set.

### GetUrl

`func (o *LinkClicked) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LinkClicked) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LinkClicked) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetQueryString

`func (o *LinkClicked) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LinkClicked) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LinkClicked) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *LinkClicked) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


