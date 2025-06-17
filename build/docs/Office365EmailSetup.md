# Office365EmailSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 200; | 
**Username** | Pointer to **string** |  Max length: 100; | [optional] 
**InboxFolder** | **string** |  Max length: 40; | 
**ProcessedFolder** | **string** |  Max length: 40; | 
**FailedFolder** | **string** |  Max length: 40; | 
**TenantId** | Pointer to **string** |  Max length: 36; | [optional] 
**ClientId** | Pointer to **string** |  Max length: 36; | [optional] 
**ClientSecret** | Pointer to **string** |  Max length: 4000; | [optional] 
**AuthorizedFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**Source** | Pointer to **int32** |  | [optional] 
**UseExistingTenantFlag** | Pointer to **NullableBool** |  | [optional] 
**ExistingTenant** | Pointer to [**ExistingTenantReference**](ExistingTenantReference.md) |  | [optional] 
**EmailConnector** | Pointer to [**EmailConnectorReference**](EmailConnectorReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOffice365EmailSetup

`func NewOffice365EmailSetup(name string, inboxFolder string, processedFolder string, failedFolder string, ) *Office365EmailSetup`

NewOffice365EmailSetup instantiates a new Office365EmailSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOffice365EmailSetupWithDefaults

`func NewOffice365EmailSetupWithDefaults() *Office365EmailSetup`

NewOffice365EmailSetupWithDefaults instantiates a new Office365EmailSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Office365EmailSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Office365EmailSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Office365EmailSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Office365EmailSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Office365EmailSetup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Office365EmailSetup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Office365EmailSetup) SetName(v string)`

SetName sets Name field to given value.


### GetUsername

`func (o *Office365EmailSetup) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Office365EmailSetup) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Office365EmailSetup) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Office365EmailSetup) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetInboxFolder

`func (o *Office365EmailSetup) GetInboxFolder() string`

GetInboxFolder returns the InboxFolder field if non-nil, zero value otherwise.

### GetInboxFolderOk

`func (o *Office365EmailSetup) GetInboxFolderOk() (*string, bool)`

GetInboxFolderOk returns a tuple with the InboxFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboxFolder

`func (o *Office365EmailSetup) SetInboxFolder(v string)`

SetInboxFolder sets InboxFolder field to given value.


### GetProcessedFolder

`func (o *Office365EmailSetup) GetProcessedFolder() string`

GetProcessedFolder returns the ProcessedFolder field if non-nil, zero value otherwise.

### GetProcessedFolderOk

`func (o *Office365EmailSetup) GetProcessedFolderOk() (*string, bool)`

GetProcessedFolderOk returns a tuple with the ProcessedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedFolder

`func (o *Office365EmailSetup) SetProcessedFolder(v string)`

SetProcessedFolder sets ProcessedFolder field to given value.


### GetFailedFolder

`func (o *Office365EmailSetup) GetFailedFolder() string`

GetFailedFolder returns the FailedFolder field if non-nil, zero value otherwise.

### GetFailedFolderOk

`func (o *Office365EmailSetup) GetFailedFolderOk() (*string, bool)`

GetFailedFolderOk returns a tuple with the FailedFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedFolder

`func (o *Office365EmailSetup) SetFailedFolder(v string)`

SetFailedFolder sets FailedFolder field to given value.


### GetTenantId

`func (o *Office365EmailSetup) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Office365EmailSetup) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Office365EmailSetup) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Office365EmailSetup) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *Office365EmailSetup) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Office365EmailSetup) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Office365EmailSetup) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Office365EmailSetup) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *Office365EmailSetup) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Office365EmailSetup) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Office365EmailSetup) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Office365EmailSetup) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAuthorizedFlag

`func (o *Office365EmailSetup) GetAuthorizedFlag() bool`

GetAuthorizedFlag returns the AuthorizedFlag field if non-nil, zero value otherwise.

### GetAuthorizedFlagOk

`func (o *Office365EmailSetup) GetAuthorizedFlagOk() (*bool, bool)`

GetAuthorizedFlagOk returns a tuple with the AuthorizedFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedFlag

`func (o *Office365EmailSetup) SetAuthorizedFlag(v bool)`

SetAuthorizedFlag sets AuthorizedFlag field to given value.

### HasAuthorizedFlag

`func (o *Office365EmailSetup) HasAuthorizedFlag() bool`

HasAuthorizedFlag returns a boolean if a field has been set.

### SetAuthorizedFlagNil

`func (o *Office365EmailSetup) SetAuthorizedFlagNil(b bool)`

 SetAuthorizedFlagNil sets the value for AuthorizedFlag to be an explicit nil

### UnsetAuthorizedFlag
`func (o *Office365EmailSetup) UnsetAuthorizedFlag()`

UnsetAuthorizedFlag ensures that no value is present for AuthorizedFlag, not even an explicit nil
### GetInactiveFlag

`func (o *Office365EmailSetup) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *Office365EmailSetup) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *Office365EmailSetup) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *Office365EmailSetup) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *Office365EmailSetup) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *Office365EmailSetup) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetSource

`func (o *Office365EmailSetup) GetSource() int32`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Office365EmailSetup) GetSourceOk() (*int32, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Office365EmailSetup) SetSource(v int32)`

SetSource sets Source field to given value.

### HasSource

`func (o *Office365EmailSetup) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUseExistingTenantFlag

`func (o *Office365EmailSetup) GetUseExistingTenantFlag() bool`

GetUseExistingTenantFlag returns the UseExistingTenantFlag field if non-nil, zero value otherwise.

### GetUseExistingTenantFlagOk

`func (o *Office365EmailSetup) GetUseExistingTenantFlagOk() (*bool, bool)`

GetUseExistingTenantFlagOk returns a tuple with the UseExistingTenantFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExistingTenantFlag

`func (o *Office365EmailSetup) SetUseExistingTenantFlag(v bool)`

SetUseExistingTenantFlag sets UseExistingTenantFlag field to given value.

### HasUseExistingTenantFlag

`func (o *Office365EmailSetup) HasUseExistingTenantFlag() bool`

HasUseExistingTenantFlag returns a boolean if a field has been set.

### SetUseExistingTenantFlagNil

`func (o *Office365EmailSetup) SetUseExistingTenantFlagNil(b bool)`

 SetUseExistingTenantFlagNil sets the value for UseExistingTenantFlag to be an explicit nil

### UnsetUseExistingTenantFlag
`func (o *Office365EmailSetup) UnsetUseExistingTenantFlag()`

UnsetUseExistingTenantFlag ensures that no value is present for UseExistingTenantFlag, not even an explicit nil
### GetExistingTenant

`func (o *Office365EmailSetup) GetExistingTenant() ExistingTenantReference`

GetExistingTenant returns the ExistingTenant field if non-nil, zero value otherwise.

### GetExistingTenantOk

`func (o *Office365EmailSetup) GetExistingTenantOk() (*ExistingTenantReference, bool)`

GetExistingTenantOk returns a tuple with the ExistingTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingTenant

`func (o *Office365EmailSetup) SetExistingTenant(v ExistingTenantReference)`

SetExistingTenant sets ExistingTenant field to given value.

### HasExistingTenant

`func (o *Office365EmailSetup) HasExistingTenant() bool`

HasExistingTenant returns a boolean if a field has been set.

### GetEmailConnector

`func (o *Office365EmailSetup) GetEmailConnector() EmailConnectorReference`

GetEmailConnector returns the EmailConnector field if non-nil, zero value otherwise.

### GetEmailConnectorOk

`func (o *Office365EmailSetup) GetEmailConnectorOk() (*EmailConnectorReference, bool)`

GetEmailConnectorOk returns a tuple with the EmailConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailConnector

`func (o *Office365EmailSetup) SetEmailConnector(v EmailConnectorReference)`

SetEmailConnector sets EmailConnector field to given value.

### HasEmailConnector

`func (o *Office365EmailSetup) HasEmailConnector() bool`

HasEmailConnector returns a boolean if a field has been set.

### GetInfo

`func (o *Office365EmailSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Office365EmailSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Office365EmailSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Office365EmailSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


