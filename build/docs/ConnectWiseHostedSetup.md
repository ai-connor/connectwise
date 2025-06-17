# ConnectWiseHostedSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScreenId** | **NullableInt32** | Can be obtained via ConnectWiseHostedApiScreen report. | 
**Description** | **string** |  Max length: 50; | 
**Url** | **string** |  Max length: 1024; | 
**Type** | **NullableString** |  | 
**ClientId** | Pointer to **string** | Only required if not already set. Max length: 36; | [optional] 
**Origin** | Pointer to **string** |  Max length: 100; | [optional] 
**PodHeight** | Pointer to **NullableInt32** |  | [optional] 
**ToolbarButtonDialogHeight** | Pointer to **NullableInt32** |  | [optional] 
**ToolbarButtonDialogWidth** | Pointer to **NullableInt32** |  | [optional] 
**ToolbarButtonText** | Pointer to **string** | Only required for ToolbarButtons. Max length: 50; | [optional] 
**ToolbarButtonToolTip** | Pointer to **string** |  Max length: 50; | [optional] 
**ToolbarButtonIconDocumentId** | Pointer to **NullableInt32** |  | [optional] 
**DisabledFlag** | Pointer to **NullableBool** |  | [optional] 
**LocationIds** | Pointer to **[]int32** |  | [optional] 
**LocationsEnabledFlag** | Pointer to **NullableBool** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewConnectWiseHostedSetup

`func NewConnectWiseHostedSetup(screenId NullableInt32, description string, url string, type_ NullableString, ) *ConnectWiseHostedSetup`

NewConnectWiseHostedSetup instantiates a new ConnectWiseHostedSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectWiseHostedSetupWithDefaults

`func NewConnectWiseHostedSetupWithDefaults() *ConnectWiseHostedSetup`

NewConnectWiseHostedSetupWithDefaults instantiates a new ConnectWiseHostedSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectWiseHostedSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectWiseHostedSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectWiseHostedSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectWiseHostedSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScreenId

`func (o *ConnectWiseHostedSetup) GetScreenId() int32`

GetScreenId returns the ScreenId field if non-nil, zero value otherwise.

### GetScreenIdOk

`func (o *ConnectWiseHostedSetup) GetScreenIdOk() (*int32, bool)`

GetScreenIdOk returns a tuple with the ScreenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenId

`func (o *ConnectWiseHostedSetup) SetScreenId(v int32)`

SetScreenId sets ScreenId field to given value.


### SetScreenIdNil

`func (o *ConnectWiseHostedSetup) SetScreenIdNil(b bool)`

 SetScreenIdNil sets the value for ScreenId to be an explicit nil

### UnsetScreenId
`func (o *ConnectWiseHostedSetup) UnsetScreenId()`

UnsetScreenId ensures that no value is present for ScreenId, not even an explicit nil
### GetDescription

`func (o *ConnectWiseHostedSetup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectWiseHostedSetup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectWiseHostedSetup) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUrl

`func (o *ConnectWiseHostedSetup) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConnectWiseHostedSetup) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConnectWiseHostedSetup) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetType

`func (o *ConnectWiseHostedSetup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectWiseHostedSetup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectWiseHostedSetup) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *ConnectWiseHostedSetup) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *ConnectWiseHostedSetup) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetClientId

`func (o *ConnectWiseHostedSetup) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ConnectWiseHostedSetup) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ConnectWiseHostedSetup) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ConnectWiseHostedSetup) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetOrigin

`func (o *ConnectWiseHostedSetup) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ConnectWiseHostedSetup) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ConnectWiseHostedSetup) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ConnectWiseHostedSetup) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPodHeight

`func (o *ConnectWiseHostedSetup) GetPodHeight() int32`

GetPodHeight returns the PodHeight field if non-nil, zero value otherwise.

### GetPodHeightOk

`func (o *ConnectWiseHostedSetup) GetPodHeightOk() (*int32, bool)`

GetPodHeightOk returns a tuple with the PodHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodHeight

`func (o *ConnectWiseHostedSetup) SetPodHeight(v int32)`

SetPodHeight sets PodHeight field to given value.

### HasPodHeight

`func (o *ConnectWiseHostedSetup) HasPodHeight() bool`

HasPodHeight returns a boolean if a field has been set.

### SetPodHeightNil

`func (o *ConnectWiseHostedSetup) SetPodHeightNil(b bool)`

 SetPodHeightNil sets the value for PodHeight to be an explicit nil

### UnsetPodHeight
`func (o *ConnectWiseHostedSetup) UnsetPodHeight()`

UnsetPodHeight ensures that no value is present for PodHeight, not even an explicit nil
### GetToolbarButtonDialogHeight

`func (o *ConnectWiseHostedSetup) GetToolbarButtonDialogHeight() int32`

GetToolbarButtonDialogHeight returns the ToolbarButtonDialogHeight field if non-nil, zero value otherwise.

### GetToolbarButtonDialogHeightOk

`func (o *ConnectWiseHostedSetup) GetToolbarButtonDialogHeightOk() (*int32, bool)`

GetToolbarButtonDialogHeightOk returns a tuple with the ToolbarButtonDialogHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarButtonDialogHeight

`func (o *ConnectWiseHostedSetup) SetToolbarButtonDialogHeight(v int32)`

SetToolbarButtonDialogHeight sets ToolbarButtonDialogHeight field to given value.

### HasToolbarButtonDialogHeight

`func (o *ConnectWiseHostedSetup) HasToolbarButtonDialogHeight() bool`

HasToolbarButtonDialogHeight returns a boolean if a field has been set.

### SetToolbarButtonDialogHeightNil

`func (o *ConnectWiseHostedSetup) SetToolbarButtonDialogHeightNil(b bool)`

 SetToolbarButtonDialogHeightNil sets the value for ToolbarButtonDialogHeight to be an explicit nil

### UnsetToolbarButtonDialogHeight
`func (o *ConnectWiseHostedSetup) UnsetToolbarButtonDialogHeight()`

UnsetToolbarButtonDialogHeight ensures that no value is present for ToolbarButtonDialogHeight, not even an explicit nil
### GetToolbarButtonDialogWidth

`func (o *ConnectWiseHostedSetup) GetToolbarButtonDialogWidth() int32`

GetToolbarButtonDialogWidth returns the ToolbarButtonDialogWidth field if non-nil, zero value otherwise.

### GetToolbarButtonDialogWidthOk

`func (o *ConnectWiseHostedSetup) GetToolbarButtonDialogWidthOk() (*int32, bool)`

GetToolbarButtonDialogWidthOk returns a tuple with the ToolbarButtonDialogWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarButtonDialogWidth

`func (o *ConnectWiseHostedSetup) SetToolbarButtonDialogWidth(v int32)`

SetToolbarButtonDialogWidth sets ToolbarButtonDialogWidth field to given value.

### HasToolbarButtonDialogWidth

`func (o *ConnectWiseHostedSetup) HasToolbarButtonDialogWidth() bool`

HasToolbarButtonDialogWidth returns a boolean if a field has been set.

### SetToolbarButtonDialogWidthNil

`func (o *ConnectWiseHostedSetup) SetToolbarButtonDialogWidthNil(b bool)`

 SetToolbarButtonDialogWidthNil sets the value for ToolbarButtonDialogWidth to be an explicit nil

### UnsetToolbarButtonDialogWidth
`func (o *ConnectWiseHostedSetup) UnsetToolbarButtonDialogWidth()`

UnsetToolbarButtonDialogWidth ensures that no value is present for ToolbarButtonDialogWidth, not even an explicit nil
### GetToolbarButtonText

`func (o *ConnectWiseHostedSetup) GetToolbarButtonText() string`

GetToolbarButtonText returns the ToolbarButtonText field if non-nil, zero value otherwise.

### GetToolbarButtonTextOk

`func (o *ConnectWiseHostedSetup) GetToolbarButtonTextOk() (*string, bool)`

GetToolbarButtonTextOk returns a tuple with the ToolbarButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarButtonText

`func (o *ConnectWiseHostedSetup) SetToolbarButtonText(v string)`

SetToolbarButtonText sets ToolbarButtonText field to given value.

### HasToolbarButtonText

`func (o *ConnectWiseHostedSetup) HasToolbarButtonText() bool`

HasToolbarButtonText returns a boolean if a field has been set.

### GetToolbarButtonToolTip

`func (o *ConnectWiseHostedSetup) GetToolbarButtonToolTip() string`

GetToolbarButtonToolTip returns the ToolbarButtonToolTip field if non-nil, zero value otherwise.

### GetToolbarButtonToolTipOk

`func (o *ConnectWiseHostedSetup) GetToolbarButtonToolTipOk() (*string, bool)`

GetToolbarButtonToolTipOk returns a tuple with the ToolbarButtonToolTip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarButtonToolTip

`func (o *ConnectWiseHostedSetup) SetToolbarButtonToolTip(v string)`

SetToolbarButtonToolTip sets ToolbarButtonToolTip field to given value.

### HasToolbarButtonToolTip

`func (o *ConnectWiseHostedSetup) HasToolbarButtonToolTip() bool`

HasToolbarButtonToolTip returns a boolean if a field has been set.

### GetToolbarButtonIconDocumentId

`func (o *ConnectWiseHostedSetup) GetToolbarButtonIconDocumentId() int32`

GetToolbarButtonIconDocumentId returns the ToolbarButtonIconDocumentId field if non-nil, zero value otherwise.

### GetToolbarButtonIconDocumentIdOk

`func (o *ConnectWiseHostedSetup) GetToolbarButtonIconDocumentIdOk() (*int32, bool)`

GetToolbarButtonIconDocumentIdOk returns a tuple with the ToolbarButtonIconDocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolbarButtonIconDocumentId

`func (o *ConnectWiseHostedSetup) SetToolbarButtonIconDocumentId(v int32)`

SetToolbarButtonIconDocumentId sets ToolbarButtonIconDocumentId field to given value.

### HasToolbarButtonIconDocumentId

`func (o *ConnectWiseHostedSetup) HasToolbarButtonIconDocumentId() bool`

HasToolbarButtonIconDocumentId returns a boolean if a field has been set.

### SetToolbarButtonIconDocumentIdNil

`func (o *ConnectWiseHostedSetup) SetToolbarButtonIconDocumentIdNil(b bool)`

 SetToolbarButtonIconDocumentIdNil sets the value for ToolbarButtonIconDocumentId to be an explicit nil

### UnsetToolbarButtonIconDocumentId
`func (o *ConnectWiseHostedSetup) UnsetToolbarButtonIconDocumentId()`

UnsetToolbarButtonIconDocumentId ensures that no value is present for ToolbarButtonIconDocumentId, not even an explicit nil
### GetDisabledFlag

`func (o *ConnectWiseHostedSetup) GetDisabledFlag() bool`

GetDisabledFlag returns the DisabledFlag field if non-nil, zero value otherwise.

### GetDisabledFlagOk

`func (o *ConnectWiseHostedSetup) GetDisabledFlagOk() (*bool, bool)`

GetDisabledFlagOk returns a tuple with the DisabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledFlag

`func (o *ConnectWiseHostedSetup) SetDisabledFlag(v bool)`

SetDisabledFlag sets DisabledFlag field to given value.

### HasDisabledFlag

`func (o *ConnectWiseHostedSetup) HasDisabledFlag() bool`

HasDisabledFlag returns a boolean if a field has been set.

### SetDisabledFlagNil

`func (o *ConnectWiseHostedSetup) SetDisabledFlagNil(b bool)`

 SetDisabledFlagNil sets the value for DisabledFlag to be an explicit nil

### UnsetDisabledFlag
`func (o *ConnectWiseHostedSetup) UnsetDisabledFlag()`

UnsetDisabledFlag ensures that no value is present for DisabledFlag, not even an explicit nil
### GetLocationIds

`func (o *ConnectWiseHostedSetup) GetLocationIds() []int32`

GetLocationIds returns the LocationIds field if non-nil, zero value otherwise.

### GetLocationIdsOk

`func (o *ConnectWiseHostedSetup) GetLocationIdsOk() (*[]int32, bool)`

GetLocationIdsOk returns a tuple with the LocationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationIds

`func (o *ConnectWiseHostedSetup) SetLocationIds(v []int32)`

SetLocationIds sets LocationIds field to given value.

### HasLocationIds

`func (o *ConnectWiseHostedSetup) HasLocationIds() bool`

HasLocationIds returns a boolean if a field has been set.

### GetLocationsEnabledFlag

`func (o *ConnectWiseHostedSetup) GetLocationsEnabledFlag() bool`

GetLocationsEnabledFlag returns the LocationsEnabledFlag field if non-nil, zero value otherwise.

### GetLocationsEnabledFlagOk

`func (o *ConnectWiseHostedSetup) GetLocationsEnabledFlagOk() (*bool, bool)`

GetLocationsEnabledFlagOk returns a tuple with the LocationsEnabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationsEnabledFlag

`func (o *ConnectWiseHostedSetup) SetLocationsEnabledFlag(v bool)`

SetLocationsEnabledFlag sets LocationsEnabledFlag field to given value.

### HasLocationsEnabledFlag

`func (o *ConnectWiseHostedSetup) HasLocationsEnabledFlag() bool`

HasLocationsEnabledFlag returns a boolean if a field has been set.

### SetLocationsEnabledFlagNil

`func (o *ConnectWiseHostedSetup) SetLocationsEnabledFlagNil(b bool)`

 SetLocationsEnabledFlagNil sets the value for LocationsEnabledFlag to be an explicit nil

### UnsetLocationsEnabledFlag
`func (o *ConnectWiseHostedSetup) UnsetLocationsEnabledFlag()`

UnsetLocationsEnabledFlag ensures that no value is present for LocationsEnabledFlag, not even an explicit nil
### GetCreatedBy

`func (o *ConnectWiseHostedSetup) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ConnectWiseHostedSetup) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ConnectWiseHostedSetup) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ConnectWiseHostedSetup) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ConnectWiseHostedSetup) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ConnectWiseHostedSetup) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ConnectWiseHostedSetup) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ConnectWiseHostedSetup) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetInfo

`func (o *ConnectWiseHostedSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ConnectWiseHostedSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ConnectWiseHostedSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ConnectWiseHostedSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


