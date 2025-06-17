# FormSubmitted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CampaignId** | Pointer to **NullableInt32** |  | [optional] 
**ContactId** | **NullableInt32** |  | 
**DateSubmitted** | Pointer to **time.Time** |  | [optional] 
**Url** | **string** |  Max length: 2083; | 
**QueryString** | Pointer to **string** |  | [optional] 
**PageType** | Pointer to **string** |  | [optional] 
**PageSubType** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewFormSubmitted

`func NewFormSubmitted(contactId NullableInt32, url string, ) *FormSubmitted`

NewFormSubmitted instantiates a new FormSubmitted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmittedWithDefaults

`func NewFormSubmittedWithDefaults() *FormSubmitted`

NewFormSubmittedWithDefaults instantiates a new FormSubmitted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FormSubmitted) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FormSubmitted) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FormSubmitted) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FormSubmitted) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCampaignId

`func (o *FormSubmitted) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *FormSubmitted) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *FormSubmitted) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *FormSubmitted) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *FormSubmitted) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *FormSubmitted) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetContactId

`func (o *FormSubmitted) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *FormSubmitted) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *FormSubmitted) SetContactId(v int32)`

SetContactId sets ContactId field to given value.


### SetContactIdNil

`func (o *FormSubmitted) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *FormSubmitted) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetDateSubmitted

`func (o *FormSubmitted) GetDateSubmitted() time.Time`

GetDateSubmitted returns the DateSubmitted field if non-nil, zero value otherwise.

### GetDateSubmittedOk

`func (o *FormSubmitted) GetDateSubmittedOk() (*time.Time, bool)`

GetDateSubmittedOk returns a tuple with the DateSubmitted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateSubmitted

`func (o *FormSubmitted) SetDateSubmitted(v time.Time)`

SetDateSubmitted sets DateSubmitted field to given value.

### HasDateSubmitted

`func (o *FormSubmitted) HasDateSubmitted() bool`

HasDateSubmitted returns a boolean if a field has been set.

### GetUrl

`func (o *FormSubmitted) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FormSubmitted) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FormSubmitted) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetQueryString

`func (o *FormSubmitted) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *FormSubmitted) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *FormSubmitted) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *FormSubmitted) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetPageType

`func (o *FormSubmitted) GetPageType() string`

GetPageType returns the PageType field if non-nil, zero value otherwise.

### GetPageTypeOk

`func (o *FormSubmitted) GetPageTypeOk() (*string, bool)`

GetPageTypeOk returns a tuple with the PageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageType

`func (o *FormSubmitted) SetPageType(v string)`

SetPageType sets PageType field to given value.

### HasPageType

`func (o *FormSubmitted) HasPageType() bool`

HasPageType returns a boolean if a field has been set.

### GetPageSubType

`func (o *FormSubmitted) GetPageSubType() string`

GetPageSubType returns the PageSubType field if non-nil, zero value otherwise.

### GetPageSubTypeOk

`func (o *FormSubmitted) GetPageSubTypeOk() (*string, bool)`

GetPageSubTypeOk returns a tuple with the PageSubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSubType

`func (o *FormSubmitted) SetPageSubType(v string)`

SetPageSubType sets PageSubType field to given value.

### HasPageSubType

`func (o *FormSubmitted) HasPageSubType() bool`

HasPageSubType returns a boolean if a field has been set.

### GetTopic

`func (o *FormSubmitted) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *FormSubmitted) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *FormSubmitted) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *FormSubmitted) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetVersion

`func (o *FormSubmitted) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FormSubmitted) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FormSubmitted) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FormSubmitted) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *FormSubmitted) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FormSubmitted) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FormSubmitted) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FormSubmitted) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


