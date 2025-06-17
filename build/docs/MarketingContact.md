# MarketingContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **NullableInt32** |  | [optional] 
**Note** | Pointer to **string** |  Max length: 50; | [optional] 
**UnsubscribeFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMarketingContact

`func NewMarketingContact() *MarketingContact`

NewMarketingContact instantiates a new MarketingContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketingContactWithDefaults

`func NewMarketingContactWithDefaults() *MarketingContact`

NewMarketingContactWithDefaults instantiates a new MarketingContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MarketingContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MarketingContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MarketingContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MarketingContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGroupId

`func (o *MarketingContact) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *MarketingContact) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *MarketingContact) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *MarketingContact) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### SetGroupIdNil

`func (o *MarketingContact) SetGroupIdNil(b bool)`

 SetGroupIdNil sets the value for GroupId to be an explicit nil

### UnsetGroupId
`func (o *MarketingContact) UnsetGroupId()`

UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
### GetNote

`func (o *MarketingContact) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *MarketingContact) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *MarketingContact) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *MarketingContact) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetUnsubscribeFlag

`func (o *MarketingContact) GetUnsubscribeFlag() bool`

GetUnsubscribeFlag returns the UnsubscribeFlag field if non-nil, zero value otherwise.

### GetUnsubscribeFlagOk

`func (o *MarketingContact) GetUnsubscribeFlagOk() (*bool, bool)`

GetUnsubscribeFlagOk returns a tuple with the UnsubscribeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsubscribeFlag

`func (o *MarketingContact) SetUnsubscribeFlag(v bool)`

SetUnsubscribeFlag sets UnsubscribeFlag field to given value.

### HasUnsubscribeFlag

`func (o *MarketingContact) HasUnsubscribeFlag() bool`

HasUnsubscribeFlag returns a boolean if a field has been set.

### SetUnsubscribeFlagNil

`func (o *MarketingContact) SetUnsubscribeFlagNil(b bool)`

 SetUnsubscribeFlagNil sets the value for UnsubscribeFlag to be an explicit nil

### UnsetUnsubscribeFlag
`func (o *MarketingContact) UnsetUnsubscribeFlag()`

UnsetUnsubscribeFlag ensures that no value is present for UnsubscribeFlag, not even an explicit nil
### GetInfo

`func (o *MarketingContact) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MarketingContact) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MarketingContact) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MarketingContact) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


