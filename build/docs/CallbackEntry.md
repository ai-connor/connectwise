# CallbackEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  Max length: 100; | [optional] 
**Url** | **string** |  Required Reference | 
**ObjectId** | **NullableInt32** |  Required Reference | 
**Type** | **string** |  Required Reference | 
**Level** | **string** |  Required Reference | 
**MemberId** | Pointer to **NullableInt32** |  | [optional] 
**PayloadVersion** | Pointer to **string** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**IsSoapCallbackFlag** | Pointer to **NullableBool** |  | [optional] 
**IsSelfSuppressedFlag** | Pointer to **NullableBool** |  | [optional] 
**ConnectWiseID** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCallbackEntry

`func NewCallbackEntry(url string, objectId NullableInt32, type_ string, level string, ) *CallbackEntry`

NewCallbackEntry instantiates a new CallbackEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallbackEntryWithDefaults

`func NewCallbackEntryWithDefaults() *CallbackEntry`

NewCallbackEntryWithDefaults instantiates a new CallbackEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallbackEntry) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallbackEntry) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallbackEntry) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CallbackEntry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *CallbackEntry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CallbackEntry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CallbackEntry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CallbackEntry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *CallbackEntry) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CallbackEntry) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CallbackEntry) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetObjectId

`func (o *CallbackEntry) GetObjectId() int32`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CallbackEntry) GetObjectIdOk() (*int32, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CallbackEntry) SetObjectId(v int32)`

SetObjectId sets ObjectId field to given value.


### SetObjectIdNil

`func (o *CallbackEntry) SetObjectIdNil(b bool)`

 SetObjectIdNil sets the value for ObjectId to be an explicit nil

### UnsetObjectId
`func (o *CallbackEntry) UnsetObjectId()`

UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
### GetType

`func (o *CallbackEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CallbackEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CallbackEntry) SetType(v string)`

SetType sets Type field to given value.


### GetLevel

`func (o *CallbackEntry) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CallbackEntry) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CallbackEntry) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetMemberId

`func (o *CallbackEntry) GetMemberId() int32`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *CallbackEntry) GetMemberIdOk() (*int32, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *CallbackEntry) SetMemberId(v int32)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *CallbackEntry) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

### SetMemberIdNil

`func (o *CallbackEntry) SetMemberIdNil(b bool)`

 SetMemberIdNil sets the value for MemberId to be an explicit nil

### UnsetMemberId
`func (o *CallbackEntry) UnsetMemberId()`

UnsetMemberId ensures that no value is present for MemberId, not even an explicit nil
### GetPayloadVersion

`func (o *CallbackEntry) GetPayloadVersion() string`

GetPayloadVersion returns the PayloadVersion field if non-nil, zero value otherwise.

### GetPayloadVersionOk

`func (o *CallbackEntry) GetPayloadVersionOk() (*string, bool)`

GetPayloadVersionOk returns a tuple with the PayloadVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadVersion

`func (o *CallbackEntry) SetPayloadVersion(v string)`

SetPayloadVersion sets PayloadVersion field to given value.

### HasPayloadVersion

`func (o *CallbackEntry) HasPayloadVersion() bool`

HasPayloadVersion returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *CallbackEntry) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *CallbackEntry) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *CallbackEntry) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *CallbackEntry) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *CallbackEntry) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *CallbackEntry) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetIsSoapCallbackFlag

`func (o *CallbackEntry) GetIsSoapCallbackFlag() bool`

GetIsSoapCallbackFlag returns the IsSoapCallbackFlag field if non-nil, zero value otherwise.

### GetIsSoapCallbackFlagOk

`func (o *CallbackEntry) GetIsSoapCallbackFlagOk() (*bool, bool)`

GetIsSoapCallbackFlagOk returns a tuple with the IsSoapCallbackFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSoapCallbackFlag

`func (o *CallbackEntry) SetIsSoapCallbackFlag(v bool)`

SetIsSoapCallbackFlag sets IsSoapCallbackFlag field to given value.

### HasIsSoapCallbackFlag

`func (o *CallbackEntry) HasIsSoapCallbackFlag() bool`

HasIsSoapCallbackFlag returns a boolean if a field has been set.

### SetIsSoapCallbackFlagNil

`func (o *CallbackEntry) SetIsSoapCallbackFlagNil(b bool)`

 SetIsSoapCallbackFlagNil sets the value for IsSoapCallbackFlag to be an explicit nil

### UnsetIsSoapCallbackFlag
`func (o *CallbackEntry) UnsetIsSoapCallbackFlag()`

UnsetIsSoapCallbackFlag ensures that no value is present for IsSoapCallbackFlag, not even an explicit nil
### GetIsSelfSuppressedFlag

`func (o *CallbackEntry) GetIsSelfSuppressedFlag() bool`

GetIsSelfSuppressedFlag returns the IsSelfSuppressedFlag field if non-nil, zero value otherwise.

### GetIsSelfSuppressedFlagOk

`func (o *CallbackEntry) GetIsSelfSuppressedFlagOk() (*bool, bool)`

GetIsSelfSuppressedFlagOk returns a tuple with the IsSelfSuppressedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSelfSuppressedFlag

`func (o *CallbackEntry) SetIsSelfSuppressedFlag(v bool)`

SetIsSelfSuppressedFlag sets IsSelfSuppressedFlag field to given value.

### HasIsSelfSuppressedFlag

`func (o *CallbackEntry) HasIsSelfSuppressedFlag() bool`

HasIsSelfSuppressedFlag returns a boolean if a field has been set.

### SetIsSelfSuppressedFlagNil

`func (o *CallbackEntry) SetIsSelfSuppressedFlagNil(b bool)`

 SetIsSelfSuppressedFlagNil sets the value for IsSelfSuppressedFlag to be an explicit nil

### UnsetIsSelfSuppressedFlag
`func (o *CallbackEntry) UnsetIsSelfSuppressedFlag()`

UnsetIsSelfSuppressedFlag ensures that no value is present for IsSelfSuppressedFlag, not even an explicit nil
### GetConnectWiseID

`func (o *CallbackEntry) GetConnectWiseID() string`

GetConnectWiseID returns the ConnectWiseID field if non-nil, zero value otherwise.

### GetConnectWiseIDOk

`func (o *CallbackEntry) GetConnectWiseIDOk() (*string, bool)`

GetConnectWiseIDOk returns a tuple with the ConnectWiseID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectWiseID

`func (o *CallbackEntry) SetConnectWiseID(v string)`

SetConnectWiseID sets ConnectWiseID field to given value.

### HasConnectWiseID

`func (o *CallbackEntry) HasConnectWiseID() bool`

HasConnectWiseID returns a boolean if a field has been set.

### GetInfo

`func (o *CallbackEntry) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CallbackEntry) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CallbackEntry) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CallbackEntry) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


