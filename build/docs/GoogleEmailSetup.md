# GoogleEmailSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 200; | 
**Username** | **string** |  Max length: 100; | 
**InboxFolder** | **string** |  Max length: 40; | 
**ProcessedFolder** | **string** |  Max length: 40; | 
**FailedFolder** | **string** |  Max length: 40; | 
**ClientId** | Pointer to **string** |  Max length: 200; | [optional] 
**PrivateKey** | Pointer to **string** |  Max length: 4000; | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailConnector** | Pointer to [**EmailConnectorReference**](EmailConnectorReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewGoogleEmailSetup

`func NewGoogleEmailSetup(name string, username string, inboxFolder string, processedFolder string, failedFolder string, ) *GoogleEmailSetup`

NewGoogleEmailSetup instantiates a new GoogleEmailSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleEmailSetupWithDefaults

`func NewGoogleEmailSetupWithDefaults() *GoogleEmailSetup`

NewGoogleEmailSetupWithDefaults instantiates a new GoogleEmailSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GoogleEmailSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GoogleEmailSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GoogleEmailSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GoogleEmailSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GoogleEmailSetup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GoogleEmailSetup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GoogleEmailSetup) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *GoogleEmailSetup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GoogleEmailSetup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GoogleEmailSetup) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetInboxFolder

`func (o *GoogleEmailSetup) GetInboxFolder() string`

GetInboxFolder returns the InboxFolder field if non-nil, zero value otherwise.

### GetInboxFolderOk

`func (o *GoogleEmailSetup) GetInboxFolderOk() (*string, bool)`

GetInboxFolderOk returns a tuple with the InboxFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxFolder

`func (o *GoogleEmailSetup) SetInboxFolder(v string)`

SetInboxFolder sets InboxFolder field to given value.


### GetProcessedFolder

`func (o *GoogleEmailSetup) GetProcessedFolder() string`

GetProcessedFolder returns the ProcessedFolder field if non-nil, zero value otherwise.

### GetProcessedFolderOk

`func (o *GoogleEmailSetup) GetProcessedFolderOk() (*string, bool)`

GetProcessedFolderOk returns a tuple with the ProcessedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedFolder

`func (o *GoogleEmailSetup) SetProcessedFolder(v string)`

SetProcessedFolder sets ProcessedFolder field to given value.


### GetFailedFolder

`func (o *GoogleEmailSetup) GetFailedFolder() string`

GetFailedFolder returns the FailedFolder field if non-nil, zero value otherwise.

### GetFailedFolderOk

`func (o *GoogleEmailSetup) GetFailedFolderOk() (*string, bool)`

GetFailedFolderOk returns a tuple with the FailedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFolder

`func (o *GoogleEmailSetup) SetFailedFolder(v string)`

SetFailedFolder sets FailedFolder field to given value.


### GetClientId

`func (o *GoogleEmailSetup) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *GoogleEmailSetup) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *GoogleEmailSetup) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *GoogleEmailSetup) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetPrivateKey

`func (o *GoogleEmailSetup) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GoogleEmailSetup) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GoogleEmailSetup) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GoogleEmailSetup) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *GoogleEmailSetup) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *GoogleEmailSetup) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *GoogleEmailSetup) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *GoogleEmailSetup) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *GoogleEmailSetup) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *GoogleEmailSetup) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetEmailConnector

`func (o *GoogleEmailSetup) GetEmailConnector() EmailConnectorReference`

GetEmailConnector returns the EmailConnector field if non-nil, zero value otherwise.

### GetEmailConnectorOk

`func (o *GoogleEmailSetup) GetEmailConnectorOk() (*EmailConnectorReference, bool)`

GetEmailConnectorOk returns a tuple with the EmailConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnector

`func (o *GoogleEmailSetup) SetEmailConnector(v EmailConnectorReference)`

SetEmailConnector sets EmailConnector field to given value.

### HasEmailConnector

`func (o *GoogleEmailSetup) HasEmailConnector() bool`

HasEmailConnector returns a boolean if a field has been set.

### GetInfo

`func (o *GoogleEmailSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *GoogleEmailSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *GoogleEmailSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *GoogleEmailSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


