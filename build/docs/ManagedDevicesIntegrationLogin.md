# ManagedDevicesIntegrationLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ManagedDevicesIntegration** | Pointer to [**ManagedDevicesIntegrationReference**](ManagedDevicesIntegrationReference.md) |  | [optional] 
**Username** | **string** |  Max length: 50; | 
**Password** | Pointer to **string** |  Max length: 50; | [optional] 
**Member** | [**MemberReference**](MemberReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagedDevicesIntegrationLogin

`func NewManagedDevicesIntegrationLogin(username string, member MemberReference, ) *ManagedDevicesIntegrationLogin`

NewManagedDevicesIntegrationLogin instantiates a new ManagedDevicesIntegrationLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDevicesIntegrationLoginWithDefaults

`func NewManagedDevicesIntegrationLoginWithDefaults() *ManagedDevicesIntegrationLogin`

NewManagedDevicesIntegrationLoginWithDefaults instantiates a new ManagedDevicesIntegrationLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedDevicesIntegrationLogin) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedDevicesIntegrationLogin) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedDevicesIntegrationLogin) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedDevicesIntegrationLogin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationLogin) GetManagedDevicesIntegration() ManagedDevicesIntegrationReference`

GetManagedDevicesIntegration returns the ManagedDevicesIntegration field if non-nil, zero value otherwise.

### GetManagedDevicesIntegrationOk

`func (o *ManagedDevicesIntegrationLogin) GetManagedDevicesIntegrationOk() (*ManagedDevicesIntegrationReference, bool)`

GetManagedDevicesIntegrationOk returns a tuple with the ManagedDevicesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationLogin) SetManagedDevicesIntegration(v ManagedDevicesIntegrationReference)`

SetManagedDevicesIntegration sets ManagedDevicesIntegration field to given value.

### HasManagedDevicesIntegration

`func (o *ManagedDevicesIntegrationLogin) HasManagedDevicesIntegration() bool`

HasManagedDevicesIntegration returns a boolean if a field has been set.

### GetUsername

`func (o *ManagedDevicesIntegrationLogin) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManagedDevicesIntegrationLogin) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManagedDevicesIntegrationLogin) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ManagedDevicesIntegrationLogin) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManagedDevicesIntegrationLogin) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManagedDevicesIntegrationLogin) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManagedDevicesIntegrationLogin) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMember

`func (o *ManagedDevicesIntegrationLogin) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ManagedDevicesIntegrationLogin) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ManagedDevicesIntegrationLogin) SetMember(v MemberReference)`

SetMember sets Member field to given value.


### GetInfo

`func (o *ManagedDevicesIntegrationLogin) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagedDevicesIntegrationLogin) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagedDevicesIntegrationLogin) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagedDevicesIntegrationLogin) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


