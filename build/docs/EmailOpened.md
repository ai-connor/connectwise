# EmailOpened

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CampaignId** | Pointer to **NullableInt32** |  | [optional] 
**ContactId** | **NullableInt32** |  | 
**DateOpened** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEmailOpened

`func NewEmailOpened(contactId NullableInt32, ) *EmailOpened`

NewEmailOpened instantiates a new EmailOpened object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailOpenedWithDefaults

`func NewEmailOpenedWithDefaults() *EmailOpened`

NewEmailOpenedWithDefaults instantiates a new EmailOpened object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailOpened) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailOpened) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailOpened) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EmailOpened) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCampaignId

`func (o *EmailOpened) GetCampaignId() int32`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *EmailOpened) GetCampaignIdOk() (*int32, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *EmailOpened) SetCampaignId(v int32)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *EmailOpened) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *EmailOpened) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *EmailOpened) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetContactId

`func (o *EmailOpened) GetContactId() int32`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *EmailOpened) GetContactIdOk() (*int32, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *EmailOpened) SetContactId(v int32)`

SetContactId sets ContactId field to given value.


### SetContactIdNil

`func (o *EmailOpened) SetContactIdNil(b bool)`

 SetContactIdNil sets the value for ContactId to be an explicit nil

### UnsetContactId
`func (o *EmailOpened) UnsetContactId()`

UnsetContactId ensures that no value is present for ContactId, not even an explicit nil
### GetDateOpened

`func (o *EmailOpened) GetDateOpened() time.Time`

GetDateOpened returns the DateOpened field if non-nil, zero value otherwise.

### GetDateOpenedOk

`func (o *EmailOpened) GetDateOpenedOk() (*time.Time, bool)`

GetDateOpenedOk returns a tuple with the DateOpened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOpened

`func (o *EmailOpened) SetDateOpened(v time.Time)`

SetDateOpened sets DateOpened field to given value.

### HasDateOpened

`func (o *EmailOpened) HasDateOpened() bool`

HasDateOpened returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


