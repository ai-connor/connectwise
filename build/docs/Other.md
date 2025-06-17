# Other

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DefaultLdap** | Pointer to [**LdapConfigurationReference**](LdapConfigurationReference.md) |  | [optional] 
**DefaultFromAddress** | **string** |  Max length: 50; | 
**PortalUrlOverride** | **string** |  Max length: 100; | 
**SiteUrl** | **string** |  Max length: 100; | 
**LogoPath** | Pointer to **string** |  Max length: 200; | [optional] 
**ContactSync** | Pointer to **NullableString** |  | [optional] 
**ServerTimeZone** | [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | 
**DefaultCalendar** | [**CalendarReference**](CalendarReference.md) |  | 
**DefaultAddressFormat** | [**AddressFormatReference**](AddressFormatReference.md) |  | 
**UseSslFlag** | Pointer to **NullableBool** |  | [optional] 
**SyncLeadsFlag** | Pointer to **NullableBool** |  | [optional] 
**IncludePortalLinkFlag** | Pointer to **NullableBool** |  | [optional] 
**UseExpandedFormatTimeFlag** | Pointer to **NullableBool** |  | [optional] 
**UseExpandedFormatActivityFlag** | Pointer to **NullableBool** |  | [optional] 
**DisableZAdminLoginFlag** | Pointer to **NullableBool** |  | [optional] 
**Locale** | [**LocaleReference**](LocaleReference.md) |  | 
**UpdateMemberTimeZonesFlag** | Pointer to **NullableBool** | If true, all Members time zone will also be set to serverTimeZone. Otherwise, only My Company time zone will be updated. | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewOther

`func NewOther(defaultFromAddress string, portalUrlOverride string, siteUrl string, serverTimeZone TimeZoneSetupReference, defaultCalendar CalendarReference, defaultAddressFormat AddressFormatReference, locale LocaleReference, ) *Other`

NewOther instantiates a new Other object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherWithDefaults

`func NewOtherWithDefaults() *Other`

NewOtherWithDefaults instantiates a new Other object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Other) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Other) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Other) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Other) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultLdap

`func (o *Other) GetDefaultLdap() LdapConfigurationReference`

GetDefaultLdap returns the DefaultLdap field if non-nil, zero value otherwise.

### GetDefaultLdapOk

`func (o *Other) GetDefaultLdapOk() (*LdapConfigurationReference, bool)`

GetDefaultLdapOk returns a tuple with the DefaultLdap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLdap

`func (o *Other) SetDefaultLdap(v LdapConfigurationReference)`

SetDefaultLdap sets DefaultLdap field to given value.

### HasDefaultLdap

`func (o *Other) HasDefaultLdap() bool`

HasDefaultLdap returns a boolean if a field has been set.

### GetDefaultFromAddress

`func (o *Other) GetDefaultFromAddress() string`

GetDefaultFromAddress returns the DefaultFromAddress field if non-nil, zero value otherwise.

### GetDefaultFromAddressOk

`func (o *Other) GetDefaultFromAddressOk() (*string, bool)`

GetDefaultFromAddressOk returns a tuple with the DefaultFromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFromAddress

`func (o *Other) SetDefaultFromAddress(v string)`

SetDefaultFromAddress sets DefaultFromAddress field to given value.


### GetPortalUrlOverride

`func (o *Other) GetPortalUrlOverride() string`

GetPortalUrlOverride returns the PortalUrlOverride field if non-nil, zero value otherwise.

### GetPortalUrlOverrideOk

`func (o *Other) GetPortalUrlOverrideOk() (*string, bool)`

GetPortalUrlOverrideOk returns a tuple with the PortalUrlOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrlOverride

`func (o *Other) SetPortalUrlOverride(v string)`

SetPortalUrlOverride sets PortalUrlOverride field to given value.


### GetSiteUrl

`func (o *Other) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *Other) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *Other) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### GetLogoPath

`func (o *Other) GetLogoPath() string`

GetLogoPath returns the LogoPath field if non-nil, zero value otherwise.

### GetLogoPathOk

`func (o *Other) GetLogoPathOk() (*string, bool)`

GetLogoPathOk returns a tuple with the LogoPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPath

`func (o *Other) SetLogoPath(v string)`

SetLogoPath sets LogoPath field to given value.

### HasLogoPath

`func (o *Other) HasLogoPath() bool`

HasLogoPath returns a boolean if a field has been set.

### GetContactSync

`func (o *Other) GetContactSync() string`

GetContactSync returns the ContactSync field if non-nil, zero value otherwise.

### GetContactSyncOk

`func (o *Other) GetContactSyncOk() (*string, bool)`

GetContactSyncOk returns a tuple with the ContactSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactSync

`func (o *Other) SetContactSync(v string)`

SetContactSync sets ContactSync field to given value.

### HasContactSync

`func (o *Other) HasContactSync() bool`

HasContactSync returns a boolean if a field has been set.

### SetContactSyncNil

`func (o *Other) SetContactSyncNil(b bool)`

 SetContactSyncNil sets the value for ContactSync to be an explicit nil

### UnsetContactSync
`func (o *Other) UnsetContactSync()`

UnsetContactSync ensures that no value is present for ContactSync, not even an explicit nil
### GetServerTimeZone

`func (o *Other) GetServerTimeZone() TimeZoneSetupReference`

GetServerTimeZone returns the ServerTimeZone field if non-nil, zero value otherwise.

### GetServerTimeZoneOk

`func (o *Other) GetServerTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetServerTimeZoneOk returns a tuple with the ServerTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTimeZone

`func (o *Other) SetServerTimeZone(v TimeZoneSetupReference)`

SetServerTimeZone sets ServerTimeZone field to given value.


### GetDefaultCalendar

`func (o *Other) GetDefaultCalendar() CalendarReference`

GetDefaultCalendar returns the DefaultCalendar field if non-nil, zero value otherwise.

### GetDefaultCalendarOk

`func (o *Other) GetDefaultCalendarOk() (*CalendarReference, bool)`

GetDefaultCalendarOk returns a tuple with the DefaultCalendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCalendar

`func (o *Other) SetDefaultCalendar(v CalendarReference)`

SetDefaultCalendar sets DefaultCalendar field to given value.


### GetDefaultAddressFormat

`func (o *Other) GetDefaultAddressFormat() AddressFormatReference`

GetDefaultAddressFormat returns the DefaultAddressFormat field if non-nil, zero value otherwise.

### GetDefaultAddressFormatOk

`func (o *Other) GetDefaultAddressFormatOk() (*AddressFormatReference, bool)`

GetDefaultAddressFormatOk returns a tuple with the DefaultAddressFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAddressFormat

`func (o *Other) SetDefaultAddressFormat(v AddressFormatReference)`

SetDefaultAddressFormat sets DefaultAddressFormat field to given value.


### GetUseSslFlag

`func (o *Other) GetUseSslFlag() bool`

GetUseSslFlag returns the UseSslFlag field if non-nil, zero value otherwise.

### GetUseSslFlagOk

`func (o *Other) GetUseSslFlagOk() (*bool, bool)`

GetUseSslFlagOk returns a tuple with the UseSslFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSslFlag

`func (o *Other) SetUseSslFlag(v bool)`

SetUseSslFlag sets UseSslFlag field to given value.

### HasUseSslFlag

`func (o *Other) HasUseSslFlag() bool`

HasUseSslFlag returns a boolean if a field has been set.

### SetUseSslFlagNil

`func (o *Other) SetUseSslFlagNil(b bool)`

 SetUseSslFlagNil sets the value for UseSslFlag to be an explicit nil

### UnsetUseSslFlag
`func (o *Other) UnsetUseSslFlag()`

UnsetUseSslFlag ensures that no value is present for UseSslFlag, not even an explicit nil
### GetSyncLeadsFlag

`func (o *Other) GetSyncLeadsFlag() bool`

GetSyncLeadsFlag returns the SyncLeadsFlag field if non-nil, zero value otherwise.

### GetSyncLeadsFlagOk

`func (o *Other) GetSyncLeadsFlagOk() (*bool, bool)`

GetSyncLeadsFlagOk returns a tuple with the SyncLeadsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncLeadsFlag

`func (o *Other) SetSyncLeadsFlag(v bool)`

SetSyncLeadsFlag sets SyncLeadsFlag field to given value.

### HasSyncLeadsFlag

`func (o *Other) HasSyncLeadsFlag() bool`

HasSyncLeadsFlag returns a boolean if a field has been set.

### SetSyncLeadsFlagNil

`func (o *Other) SetSyncLeadsFlagNil(b bool)`

 SetSyncLeadsFlagNil sets the value for SyncLeadsFlag to be an explicit nil

### UnsetSyncLeadsFlag
`func (o *Other) UnsetSyncLeadsFlag()`

UnsetSyncLeadsFlag ensures that no value is present for SyncLeadsFlag, not even an explicit nil
### GetIncludePortalLinkFlag

`func (o *Other) GetIncludePortalLinkFlag() bool`

GetIncludePortalLinkFlag returns the IncludePortalLinkFlag field if non-nil, zero value otherwise.

### GetIncludePortalLinkFlagOk

`func (o *Other) GetIncludePortalLinkFlagOk() (*bool, bool)`

GetIncludePortalLinkFlagOk returns a tuple with the IncludePortalLinkFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePortalLinkFlag

`func (o *Other) SetIncludePortalLinkFlag(v bool)`

SetIncludePortalLinkFlag sets IncludePortalLinkFlag field to given value.

### HasIncludePortalLinkFlag

`func (o *Other) HasIncludePortalLinkFlag() bool`

HasIncludePortalLinkFlag returns a boolean if a field has been set.

### SetIncludePortalLinkFlagNil

`func (o *Other) SetIncludePortalLinkFlagNil(b bool)`

 SetIncludePortalLinkFlagNil sets the value for IncludePortalLinkFlag to be an explicit nil

### UnsetIncludePortalLinkFlag
`func (o *Other) UnsetIncludePortalLinkFlag()`

UnsetIncludePortalLinkFlag ensures that no value is present for IncludePortalLinkFlag, not even an explicit nil
### GetUseExpandedFormatTimeFlag

`func (o *Other) GetUseExpandedFormatTimeFlag() bool`

GetUseExpandedFormatTimeFlag returns the UseExpandedFormatTimeFlag field if non-nil, zero value otherwise.

### GetUseExpandedFormatTimeFlagOk

`func (o *Other) GetUseExpandedFormatTimeFlagOk() (*bool, bool)`

GetUseExpandedFormatTimeFlagOk returns a tuple with the UseExpandedFormatTimeFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExpandedFormatTimeFlag

`func (o *Other) SetUseExpandedFormatTimeFlag(v bool)`

SetUseExpandedFormatTimeFlag sets UseExpandedFormatTimeFlag field to given value.

### HasUseExpandedFormatTimeFlag

`func (o *Other) HasUseExpandedFormatTimeFlag() bool`

HasUseExpandedFormatTimeFlag returns a boolean if a field has been set.

### SetUseExpandedFormatTimeFlagNil

`func (o *Other) SetUseExpandedFormatTimeFlagNil(b bool)`

 SetUseExpandedFormatTimeFlagNil sets the value for UseExpandedFormatTimeFlag to be an explicit nil

### UnsetUseExpandedFormatTimeFlag
`func (o *Other) UnsetUseExpandedFormatTimeFlag()`

UnsetUseExpandedFormatTimeFlag ensures that no value is present for UseExpandedFormatTimeFlag, not even an explicit nil
### GetUseExpandedFormatActivityFlag

`func (o *Other) GetUseExpandedFormatActivityFlag() bool`

GetUseExpandedFormatActivityFlag returns the UseExpandedFormatActivityFlag field if non-nil, zero value otherwise.

### GetUseExpandedFormatActivityFlagOk

`func (o *Other) GetUseExpandedFormatActivityFlagOk() (*bool, bool)`

GetUseExpandedFormatActivityFlagOk returns a tuple with the UseExpandedFormatActivityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExpandedFormatActivityFlag

`func (o *Other) SetUseExpandedFormatActivityFlag(v bool)`

SetUseExpandedFormatActivityFlag sets UseExpandedFormatActivityFlag field to given value.

### HasUseExpandedFormatActivityFlag

`func (o *Other) HasUseExpandedFormatActivityFlag() bool`

HasUseExpandedFormatActivityFlag returns a boolean if a field has been set.

### SetUseExpandedFormatActivityFlagNil

`func (o *Other) SetUseExpandedFormatActivityFlagNil(b bool)`

 SetUseExpandedFormatActivityFlagNil sets the value for UseExpandedFormatActivityFlag to be an explicit nil

### UnsetUseExpandedFormatActivityFlag
`func (o *Other) UnsetUseExpandedFormatActivityFlag()`

UnsetUseExpandedFormatActivityFlag ensures that no value is present for UseExpandedFormatActivityFlag, not even an explicit nil
### GetDisableZAdminLoginFlag

`func (o *Other) GetDisableZAdminLoginFlag() bool`

GetDisableZAdminLoginFlag returns the DisableZAdminLoginFlag field if non-nil, zero value otherwise.

### GetDisableZAdminLoginFlagOk

`func (o *Other) GetDisableZAdminLoginFlagOk() (*bool, bool)`

GetDisableZAdminLoginFlagOk returns a tuple with the DisableZAdminLoginFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableZAdminLoginFlag

`func (o *Other) SetDisableZAdminLoginFlag(v bool)`

SetDisableZAdminLoginFlag sets DisableZAdminLoginFlag field to given value.

### HasDisableZAdminLoginFlag

`func (o *Other) HasDisableZAdminLoginFlag() bool`

HasDisableZAdminLoginFlag returns a boolean if a field has been set.

### SetDisableZAdminLoginFlagNil

`func (o *Other) SetDisableZAdminLoginFlagNil(b bool)`

 SetDisableZAdminLoginFlagNil sets the value for DisableZAdminLoginFlag to be an explicit nil

### UnsetDisableZAdminLoginFlag
`func (o *Other) UnsetDisableZAdminLoginFlag()`

UnsetDisableZAdminLoginFlag ensures that no value is present for DisableZAdminLoginFlag, not even an explicit nil
### GetLocale

`func (o *Other) GetLocale() LocaleReference`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *Other) GetLocaleOk() (*LocaleReference, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *Other) SetLocale(v LocaleReference)`

SetLocale sets Locale field to given value.


### GetUpdateMemberTimeZonesFlag

`func (o *Other) GetUpdateMemberTimeZonesFlag() bool`

GetUpdateMemberTimeZonesFlag returns the UpdateMemberTimeZonesFlag field if non-nil, zero value otherwise.

### GetUpdateMemberTimeZonesFlagOk

`func (o *Other) GetUpdateMemberTimeZonesFlagOk() (*bool, bool)`

GetUpdateMemberTimeZonesFlagOk returns a tuple with the UpdateMemberTimeZonesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMemberTimeZonesFlag

`func (o *Other) SetUpdateMemberTimeZonesFlag(v bool)`

SetUpdateMemberTimeZonesFlag sets UpdateMemberTimeZonesFlag field to given value.

### HasUpdateMemberTimeZonesFlag

`func (o *Other) HasUpdateMemberTimeZonesFlag() bool`

HasUpdateMemberTimeZonesFlag returns a boolean if a field has been set.

### SetUpdateMemberTimeZonesFlagNil

`func (o *Other) SetUpdateMemberTimeZonesFlagNil(b bool)`

 SetUpdateMemberTimeZonesFlagNil sets the value for UpdateMemberTimeZonesFlag to be an explicit nil

### UnsetUpdateMemberTimeZonesFlag
`func (o *Other) UnsetUpdateMemberTimeZonesFlag()`

UnsetUpdateMemberTimeZonesFlag ensures that no value is present for UpdateMemberTimeZonesFlag, not even an explicit nil
### GetInfo

`func (o *Other) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Other) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Other) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Other) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


