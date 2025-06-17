# MemberSsoSettingsReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**SsoUserId** | Pointer to **string** |  | [optional] 
**UserName** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberSsoSettingsReference

`func NewMemberSsoSettingsReference() *MemberSsoSettingsReference`

NewMemberSsoSettingsReference instantiates a new MemberSsoSettingsReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSsoSettingsReferenceWithDefaults

`func NewMemberSsoSettingsReferenceWithDefaults() *MemberSsoSettingsReference`

NewMemberSsoSettingsReferenceWithDefaults instantiates a new MemberSsoSettingsReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberSsoSettingsReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberSsoSettingsReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberSsoSettingsReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberSsoSettingsReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MemberSsoSettingsReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MemberSsoSettingsReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetSsoUserId

`func (o *MemberSsoSettingsReference) GetSsoUserId() string`

GetSsoUserId returns the SsoUserId field if non-nil, zero value otherwise.

### GetSsoUserIdOk

`func (o *MemberSsoSettingsReference) GetSsoUserIdOk() (*string, bool)`

GetSsoUserIdOk returns a tuple with the SsoUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUserId

`func (o *MemberSsoSettingsReference) SetSsoUserId(v string)`

SetSsoUserId sets SsoUserId field to given value.

### HasSsoUserId

`func (o *MemberSsoSettingsReference) HasSsoUserId() bool`

HasSsoUserId returns a boolean if a field has been set.

### GetUserName

`func (o *MemberSsoSettingsReference) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *MemberSsoSettingsReference) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *MemberSsoSettingsReference) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *MemberSsoSettingsReference) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetEmail

`func (o *MemberSsoSettingsReference) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MemberSsoSettingsReference) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MemberSsoSettingsReference) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MemberSsoSettingsReference) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetInfo

`func (o *MemberSsoSettingsReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberSsoSettingsReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberSsoSettingsReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberSsoSettingsReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


