# CorporateStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**LevelCount** | Pointer to **NullableString** |  | [optional] 
**Level1Name** | Pointer to **string** |  Max length: 20; | [optional] 
**Level2Name** | Pointer to **string** |  Max length: 20; | [optional] 
**Level3Name** | Pointer to **string** |  Max length: 20; | [optional] 
**Level4Name** | Pointer to **string** |  Max length: 20; | [optional] 
**Level5Name** | Pointer to **string** |  Max length: 20; | [optional] 
**FiscalYearStart** | **NullableString** |  | 
**LocationCaption** | **string** |  Max length: 50; | 
**GroupCaption** | **string** |  Max length: 50; | 
**BaseCurrency** | [**CurrencyReference**](CurrencyReference.md) |  | 
**President** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ChiefOperatingOfficer** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Controller** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Dispatcher** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ServiceManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**DutyManager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCorporateStructure

`func NewCorporateStructure(fiscalYearStart NullableString, locationCaption string, groupCaption string, baseCurrency CurrencyReference, ) *CorporateStructure`

NewCorporateStructure instantiates a new CorporateStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorporateStructureWithDefaults

`func NewCorporateStructureWithDefaults() *CorporateStructure`

NewCorporateStructureWithDefaults instantiates a new CorporateStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CorporateStructure) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CorporateStructure) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CorporateStructure) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CorporateStructure) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevelCount

`func (o *CorporateStructure) GetLevelCount() string`

GetLevelCount returns the LevelCount field if non-nil, zero value otherwise.

### GetLevelCountOk

`func (o *CorporateStructure) GetLevelCountOk() (*string, bool)`

GetLevelCountOk returns a tuple with the LevelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelCount

`func (o *CorporateStructure) SetLevelCount(v string)`

SetLevelCount sets LevelCount field to given value.

### HasLevelCount

`func (o *CorporateStructure) HasLevelCount() bool`

HasLevelCount returns a boolean if a field has been set.

### SetLevelCountNil

`func (o *CorporateStructure) SetLevelCountNil(b bool)`

 SetLevelCountNil sets the value for LevelCount to be an explicit nil

### UnsetLevelCount
`func (o *CorporateStructure) UnsetLevelCount()`

UnsetLevelCount ensures that no value is present for LevelCount, not even an explicit nil
### GetLevel1Name

`func (o *CorporateStructure) GetLevel1Name() string`

GetLevel1Name returns the Level1Name field if non-nil, zero value otherwise.

### GetLevel1NameOk

`func (o *CorporateStructure) GetLevel1NameOk() (*string, bool)`

GetLevel1NameOk returns a tuple with the Level1Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel1Name

`func (o *CorporateStructure) SetLevel1Name(v string)`

SetLevel1Name sets Level1Name field to given value.

### HasLevel1Name

`func (o *CorporateStructure) HasLevel1Name() bool`

HasLevel1Name returns a boolean if a field has been set.

### GetLevel2Name

`func (o *CorporateStructure) GetLevel2Name() string`

GetLevel2Name returns the Level2Name field if non-nil, zero value otherwise.

### GetLevel2NameOk

`func (o *CorporateStructure) GetLevel2NameOk() (*string, bool)`

GetLevel2NameOk returns a tuple with the Level2Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel2Name

`func (o *CorporateStructure) SetLevel2Name(v string)`

SetLevel2Name sets Level2Name field to given value.

### HasLevel2Name

`func (o *CorporateStructure) HasLevel2Name() bool`

HasLevel2Name returns a boolean if a field has been set.

### GetLevel3Name

`func (o *CorporateStructure) GetLevel3Name() string`

GetLevel3Name returns the Level3Name field if non-nil, zero value otherwise.

### GetLevel3NameOk

`func (o *CorporateStructure) GetLevel3NameOk() (*string, bool)`

GetLevel3NameOk returns a tuple with the Level3Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel3Name

`func (o *CorporateStructure) SetLevel3Name(v string)`

SetLevel3Name sets Level3Name field to given value.

### HasLevel3Name

`func (o *CorporateStructure) HasLevel3Name() bool`

HasLevel3Name returns a boolean if a field has been set.

### GetLevel4Name

`func (o *CorporateStructure) GetLevel4Name() string`

GetLevel4Name returns the Level4Name field if non-nil, zero value otherwise.

### GetLevel4NameOk

`func (o *CorporateStructure) GetLevel4NameOk() (*string, bool)`

GetLevel4NameOk returns a tuple with the Level4Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel4Name

`func (o *CorporateStructure) SetLevel4Name(v string)`

SetLevel4Name sets Level4Name field to given value.

### HasLevel4Name

`func (o *CorporateStructure) HasLevel4Name() bool`

HasLevel4Name returns a boolean if a field has been set.

### GetLevel5Name

`func (o *CorporateStructure) GetLevel5Name() string`

GetLevel5Name returns the Level5Name field if non-nil, zero value otherwise.

### GetLevel5NameOk

`func (o *CorporateStructure) GetLevel5NameOk() (*string, bool)`

GetLevel5NameOk returns a tuple with the Level5Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel5Name

`func (o *CorporateStructure) SetLevel5Name(v string)`

SetLevel5Name sets Level5Name field to given value.

### HasLevel5Name

`func (o *CorporateStructure) HasLevel5Name() bool`

HasLevel5Name returns a boolean if a field has been set.

### GetFiscalYearStart

`func (o *CorporateStructure) GetFiscalYearStart() string`

GetFiscalYearStart returns the FiscalYearStart field if non-nil, zero value otherwise.

### GetFiscalYearStartOk

`func (o *CorporateStructure) GetFiscalYearStartOk() (*string, bool)`

GetFiscalYearStartOk returns a tuple with the FiscalYearStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiscalYearStart

`func (o *CorporateStructure) SetFiscalYearStart(v string)`

SetFiscalYearStart sets FiscalYearStart field to given value.


### SetFiscalYearStartNil

`func (o *CorporateStructure) SetFiscalYearStartNil(b bool)`

 SetFiscalYearStartNil sets the value for FiscalYearStart to be an explicit nil

### UnsetFiscalYearStart
`func (o *CorporateStructure) UnsetFiscalYearStart()`

UnsetFiscalYearStart ensures that no value is present for FiscalYearStart, not even an explicit nil
### GetLocationCaption

`func (o *CorporateStructure) GetLocationCaption() string`

GetLocationCaption returns the LocationCaption field if non-nil, zero value otherwise.

### GetLocationCaptionOk

`func (o *CorporateStructure) GetLocationCaptionOk() (*string, bool)`

GetLocationCaptionOk returns a tuple with the LocationCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationCaption

`func (o *CorporateStructure) SetLocationCaption(v string)`

SetLocationCaption sets LocationCaption field to given value.


### GetGroupCaption

`func (o *CorporateStructure) GetGroupCaption() string`

GetGroupCaption returns the GroupCaption field if non-nil, zero value otherwise.

### GetGroupCaptionOk

`func (o *CorporateStructure) GetGroupCaptionOk() (*string, bool)`

GetGroupCaptionOk returns a tuple with the GroupCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCaption

`func (o *CorporateStructure) SetGroupCaption(v string)`

SetGroupCaption sets GroupCaption field to given value.


### GetBaseCurrency

`func (o *CorporateStructure) GetBaseCurrency() CurrencyReference`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *CorporateStructure) GetBaseCurrencyOk() (*CurrencyReference, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *CorporateStructure) SetBaseCurrency(v CurrencyReference)`

SetBaseCurrency sets BaseCurrency field to given value.


### GetPresident

`func (o *CorporateStructure) GetPresident() MemberReference`

GetPresident returns the President field if non-nil, zero value otherwise.

### GetPresidentOk

`func (o *CorporateStructure) GetPresidentOk() (*MemberReference, bool)`

GetPresidentOk returns a tuple with the President field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresident

`func (o *CorporateStructure) SetPresident(v MemberReference)`

SetPresident sets President field to given value.

### HasPresident

`func (o *CorporateStructure) HasPresident() bool`

HasPresident returns a boolean if a field has been set.

### GetChiefOperatingOfficer

`func (o *CorporateStructure) GetChiefOperatingOfficer() MemberReference`

GetChiefOperatingOfficer returns the ChiefOperatingOfficer field if non-nil, zero value otherwise.

### GetChiefOperatingOfficerOk

`func (o *CorporateStructure) GetChiefOperatingOfficerOk() (*MemberReference, bool)`

GetChiefOperatingOfficerOk returns a tuple with the ChiefOperatingOfficer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChiefOperatingOfficer

`func (o *CorporateStructure) SetChiefOperatingOfficer(v MemberReference)`

SetChiefOperatingOfficer sets ChiefOperatingOfficer field to given value.

### HasChiefOperatingOfficer

`func (o *CorporateStructure) HasChiefOperatingOfficer() bool`

HasChiefOperatingOfficer returns a boolean if a field has been set.

### GetController

`func (o *CorporateStructure) GetController() MemberReference`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *CorporateStructure) GetControllerOk() (*MemberReference, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *CorporateStructure) SetController(v MemberReference)`

SetController sets Controller field to given value.

### HasController

`func (o *CorporateStructure) HasController() bool`

HasController returns a boolean if a field has been set.

### GetDispatcher

`func (o *CorporateStructure) GetDispatcher() MemberReference`

GetDispatcher returns the Dispatcher field if non-nil, zero value otherwise.

### GetDispatcherOk

`func (o *CorporateStructure) GetDispatcherOk() (*MemberReference, bool)`

GetDispatcherOk returns a tuple with the Dispatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatcher

`func (o *CorporateStructure) SetDispatcher(v MemberReference)`

SetDispatcher sets Dispatcher field to given value.

### HasDispatcher

`func (o *CorporateStructure) HasDispatcher() bool`

HasDispatcher returns a boolean if a field has been set.

### GetServiceManager

`func (o *CorporateStructure) GetServiceManager() MemberReference`

GetServiceManager returns the ServiceManager field if non-nil, zero value otherwise.

### GetServiceManagerOk

`func (o *CorporateStructure) GetServiceManagerOk() (*MemberReference, bool)`

GetServiceManagerOk returns a tuple with the ServiceManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceManager

`func (o *CorporateStructure) SetServiceManager(v MemberReference)`

SetServiceManager sets ServiceManager field to given value.

### HasServiceManager

`func (o *CorporateStructure) HasServiceManager() bool`

HasServiceManager returns a boolean if a field has been set.

### GetDutyManager

`func (o *CorporateStructure) GetDutyManager() MemberReference`

GetDutyManager returns the DutyManager field if non-nil, zero value otherwise.

### GetDutyManagerOk

`func (o *CorporateStructure) GetDutyManagerOk() (*MemberReference, bool)`

GetDutyManagerOk returns a tuple with the DutyManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutyManager

`func (o *CorporateStructure) SetDutyManager(v MemberReference)`

SetDutyManager sets DutyManager field to given value.

### HasDutyManager

`func (o *CorporateStructure) HasDutyManager() bool`

HasDutyManager returns a boolean if a field has been set.

### GetInfo

`func (o *CorporateStructure) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CorporateStructure) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CorporateStructure) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CorporateStructure) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


