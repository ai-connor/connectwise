# ManagedDeviceAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ManagedDevicesIntegration** | Pointer to [**ManagedDevicesIntegrationReference**](ManagedDevicesIntegrationReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagedDeviceAccount

`func NewManagedDeviceAccount() *ManagedDeviceAccount`

NewManagedDeviceAccount instantiates a new ManagedDeviceAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDeviceAccountWithDefaults

`func NewManagedDeviceAccountWithDefaults() *ManagedDeviceAccount`

NewManagedDeviceAccountWithDefaults instantiates a new ManagedDeviceAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedDeviceAccount) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedDeviceAccount) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedDeviceAccount) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedDeviceAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ManagedDeviceAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ManagedDeviceAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ManagedDeviceAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ManagedDeviceAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ManagedDeviceAccount) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ManagedDeviceAccount) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ManagedDeviceAccount) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ManagedDeviceAccount) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetManagedDevicesIntegration

`func (o *ManagedDeviceAccount) GetManagedDevicesIntegration() ManagedDevicesIntegrationReference`

GetManagedDevicesIntegration returns the ManagedDevicesIntegration field if non-nil, zero value otherwise.

### GetManagedDevicesIntegrationOk

`func (o *ManagedDeviceAccount) GetManagedDevicesIntegrationOk() (*ManagedDevicesIntegrationReference, bool)`

GetManagedDevicesIntegrationOk returns a tuple with the ManagedDevicesIntegration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedDevicesIntegration

`func (o *ManagedDeviceAccount) SetManagedDevicesIntegration(v ManagedDevicesIntegrationReference)`

SetManagedDevicesIntegration sets ManagedDevicesIntegration field to given value.

### HasManagedDevicesIntegration

`func (o *ManagedDeviceAccount) HasManagedDevicesIntegration() bool`

HasManagedDevicesIntegration returns a boolean if a field has been set.

### GetInfo

`func (o *ManagedDeviceAccount) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagedDeviceAccount) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagedDeviceAccount) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagedDeviceAccount) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


