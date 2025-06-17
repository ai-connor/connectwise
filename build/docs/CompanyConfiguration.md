# CompanyConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 100; | 
**Type** | [**ConfigurationTypeReference**](ConfigurationTypeReference.md) |  | 
**Status** | Pointer to [**ConfigurationStatusReference**](ConfigurationStatusReference.md) |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**BusinessUnitId** | Pointer to **NullableInt32** |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**DeviceIdentifier** | Pointer to **string** |  Max length: 100; | [optional] 
**SerialNumber** | Pointer to **string** |  Max length: 250; | [optional] 
**ModelNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**TagNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**PurchaseDate** | Pointer to **time.Time** |  | [optional] 
**InstallationDate** | Pointer to **time.Time** |  | [optional] 
**InstalledBy** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**WarrantyExpirationDate** | Pointer to **time.Time** |  | [optional] 
**VendorNotes** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**MacAddress** | Pointer to **string** |  Max length: 25; | [optional] 
**LastLoginName** | Pointer to **string** |  Max length: 100; | [optional] 
**BillFlag** | Pointer to **NullableBool** |  | [optional] 
**BackupSuccesses** | Pointer to **NullableInt32** |  | [optional] 
**BackupIncomplete** | Pointer to **NullableInt32** |  | [optional] 
**BackupFailed** | Pointer to **NullableInt32** |  | [optional] 
**BackupRestores** | Pointer to **NullableInt32** |  | [optional] 
**LastBackupDate** | Pointer to **time.Time** |  | [optional] 
**BackupServerName** | Pointer to **string** |  Max length: 50; | [optional] 
**BackupBillableSpaceGb** | Pointer to **NullableFloat64** |  | [optional] 
**BackupProtectedDeviceList** | Pointer to **string** |  | [optional] 
**BackupYear** | Pointer to **NullableInt32** |  | [optional] 
**BackupMonth** | Pointer to **NullableInt32** |  | [optional] 
**IpAddress** | Pointer to **string** |  Max length: 50; | [optional] 
**DefaultGateway** | Pointer to **string** |  Max length: 50; | [optional] 
**OsType** | Pointer to **string** |  Max length: 250; | [optional] 
**OsInfo** | Pointer to **string** |  Max length: 250; | [optional] 
**CpuSpeed** | Pointer to **string** |  Max length: 100; | [optional] 
**Ram** | Pointer to **string** |  Max length: 25; | [optional] 
**LocalHardDrives** | Pointer to **string** |  | [optional] 
**ParentConfigurationId** | Pointer to **NullableInt32** |  | [optional] 
**Vendor** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Manufacturer** | Pointer to [**ManufacturerReference**](ManufacturerReference.md) |  | [optional] 
**Questions** | Pointer to [**[]ConfigurationQuestion**](ConfigurationQuestion.md) |  | [optional] 
**ActiveFlag** | Pointer to **NullableBool** |  | [optional] 
**ManagementLink** | Pointer to **string** |  Max length: 1000; | [optional] 
**RemoteLink** | Pointer to **string** |  Max length: 1000; | [optional] 
**Sla** | Pointer to [**SLAReference**](SLAReference.md) |  | [optional] 
**MobileGuid** | Pointer to **NullableString** |  | [optional] 
**DisplayVendorFlag** | Pointer to **NullableBool** |  | [optional] 
**CompanyLocationId** | Pointer to **NullableInt32** |  | [optional] 
**ShowRemoteFlag** | Pointer to **NullableBool** |  | [optional] 
**ShowAutomateFlag** | Pointer to **NullableBool** |  | [optional] 
**NeedsRenewalFlag** | Pointer to **NullableBool** |  | [optional] 
**ManufacturerPartNumber** | Pointer to **string** |  Max length: 50; | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCompanyConfiguration

`func NewCompanyConfiguration(name string, type_ ConfigurationTypeReference, company CompanyReference, ) *CompanyConfiguration`

NewCompanyConfiguration instantiates a new CompanyConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyConfigurationWithDefaults

`func NewCompanyConfigurationWithDefaults() *CompanyConfiguration`

NewCompanyConfigurationWithDefaults instantiates a new CompanyConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyConfiguration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyConfiguration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyConfiguration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CompanyConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyConfiguration) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CompanyConfiguration) GetType() ConfigurationTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyConfiguration) GetTypeOk() (*ConfigurationTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyConfiguration) SetType(v ConfigurationTypeReference)`

SetType sets Type field to given value.


### GetStatus

`func (o *CompanyConfiguration) GetStatus() ConfigurationStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyConfiguration) GetStatusOk() (*ConfigurationStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyConfiguration) SetStatus(v ConfigurationStatusReference)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyConfiguration) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyConfiguration) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyConfiguration) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetContact

`func (o *CompanyConfiguration) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CompanyConfiguration) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CompanyConfiguration) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CompanyConfiguration) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetSite

`func (o *CompanyConfiguration) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *CompanyConfiguration) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *CompanyConfiguration) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *CompanyConfiguration) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetLocationId

`func (o *CompanyConfiguration) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *CompanyConfiguration) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *CompanyConfiguration) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *CompanyConfiguration) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *CompanyConfiguration) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *CompanyConfiguration) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *CompanyConfiguration) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CompanyConfiguration) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CompanyConfiguration) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CompanyConfiguration) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBusinessUnitId

`func (o *CompanyConfiguration) GetBusinessUnitId() int32`

GetBusinessUnitId returns the BusinessUnitId field if non-nil, zero value otherwise.

### GetBusinessUnitIdOk

`func (o *CompanyConfiguration) GetBusinessUnitIdOk() (*int32, bool)`

GetBusinessUnitIdOk returns a tuple with the BusinessUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnitId

`func (o *CompanyConfiguration) SetBusinessUnitId(v int32)`

SetBusinessUnitId sets BusinessUnitId field to given value.

### HasBusinessUnitId

`func (o *CompanyConfiguration) HasBusinessUnitId() bool`

HasBusinessUnitId returns a boolean if a field has been set.

### SetBusinessUnitIdNil

`func (o *CompanyConfiguration) SetBusinessUnitIdNil(b bool)`

 SetBusinessUnitIdNil sets the value for BusinessUnitId to be an explicit nil

### UnsetBusinessUnitId
`func (o *CompanyConfiguration) UnsetBusinessUnitId()`

UnsetBusinessUnitId ensures that no value is present for BusinessUnitId, not even an explicit nil
### GetDepartment

`func (o *CompanyConfiguration) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CompanyConfiguration) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CompanyConfiguration) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CompanyConfiguration) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDeviceIdentifier

`func (o *CompanyConfiguration) GetDeviceIdentifier() string`

GetDeviceIdentifier returns the DeviceIdentifier field if non-nil, zero value otherwise.

### GetDeviceIdentifierOk

`func (o *CompanyConfiguration) GetDeviceIdentifierOk() (*string, bool)`

GetDeviceIdentifierOk returns a tuple with the DeviceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIdentifier

`func (o *CompanyConfiguration) SetDeviceIdentifier(v string)`

SetDeviceIdentifier sets DeviceIdentifier field to given value.

### HasDeviceIdentifier

`func (o *CompanyConfiguration) HasDeviceIdentifier() bool`

HasDeviceIdentifier returns a boolean if a field has been set.

### GetSerialNumber

`func (o *CompanyConfiguration) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CompanyConfiguration) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CompanyConfiguration) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CompanyConfiguration) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetModelNumber

`func (o *CompanyConfiguration) GetModelNumber() string`

GetModelNumber returns the ModelNumber field if non-nil, zero value otherwise.

### GetModelNumberOk

`func (o *CompanyConfiguration) GetModelNumberOk() (*string, bool)`

GetModelNumberOk returns a tuple with the ModelNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelNumber

`func (o *CompanyConfiguration) SetModelNumber(v string)`

SetModelNumber sets ModelNumber field to given value.

### HasModelNumber

`func (o *CompanyConfiguration) HasModelNumber() bool`

HasModelNumber returns a boolean if a field has been set.

### GetTagNumber

`func (o *CompanyConfiguration) GetTagNumber() string`

GetTagNumber returns the TagNumber field if non-nil, zero value otherwise.

### GetTagNumberOk

`func (o *CompanyConfiguration) GetTagNumberOk() (*string, bool)`

GetTagNumberOk returns a tuple with the TagNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagNumber

`func (o *CompanyConfiguration) SetTagNumber(v string)`

SetTagNumber sets TagNumber field to given value.

### HasTagNumber

`func (o *CompanyConfiguration) HasTagNumber() bool`

HasTagNumber returns a boolean if a field has been set.

### GetPurchaseDate

`func (o *CompanyConfiguration) GetPurchaseDate() time.Time`

GetPurchaseDate returns the PurchaseDate field if non-nil, zero value otherwise.

### GetPurchaseDateOk

`func (o *CompanyConfiguration) GetPurchaseDateOk() (*time.Time, bool)`

GetPurchaseDateOk returns a tuple with the PurchaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseDate

`func (o *CompanyConfiguration) SetPurchaseDate(v time.Time)`

SetPurchaseDate sets PurchaseDate field to given value.

### HasPurchaseDate

`func (o *CompanyConfiguration) HasPurchaseDate() bool`

HasPurchaseDate returns a boolean if a field has been set.

### GetInstallationDate

`func (o *CompanyConfiguration) GetInstallationDate() time.Time`

GetInstallationDate returns the InstallationDate field if non-nil, zero value otherwise.

### GetInstallationDateOk

`func (o *CompanyConfiguration) GetInstallationDateOk() (*time.Time, bool)`

GetInstallationDateOk returns a tuple with the InstallationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationDate

`func (o *CompanyConfiguration) SetInstallationDate(v time.Time)`

SetInstallationDate sets InstallationDate field to given value.

### HasInstallationDate

`func (o *CompanyConfiguration) HasInstallationDate() bool`

HasInstallationDate returns a boolean if a field has been set.

### GetInstalledBy

`func (o *CompanyConfiguration) GetInstalledBy() MemberReference`

GetInstalledBy returns the InstalledBy field if non-nil, zero value otherwise.

### GetInstalledByOk

`func (o *CompanyConfiguration) GetInstalledByOk() (*MemberReference, bool)`

GetInstalledByOk returns a tuple with the InstalledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledBy

`func (o *CompanyConfiguration) SetInstalledBy(v MemberReference)`

SetInstalledBy sets InstalledBy field to given value.

### HasInstalledBy

`func (o *CompanyConfiguration) HasInstalledBy() bool`

HasInstalledBy returns a boolean if a field has been set.

### GetWarrantyExpirationDate

`func (o *CompanyConfiguration) GetWarrantyExpirationDate() time.Time`

GetWarrantyExpirationDate returns the WarrantyExpirationDate field if non-nil, zero value otherwise.

### GetWarrantyExpirationDateOk

`func (o *CompanyConfiguration) GetWarrantyExpirationDateOk() (*time.Time, bool)`

GetWarrantyExpirationDateOk returns a tuple with the WarrantyExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarrantyExpirationDate

`func (o *CompanyConfiguration) SetWarrantyExpirationDate(v time.Time)`

SetWarrantyExpirationDate sets WarrantyExpirationDate field to given value.

### HasWarrantyExpirationDate

`func (o *CompanyConfiguration) HasWarrantyExpirationDate() bool`

HasWarrantyExpirationDate returns a boolean if a field has been set.

### GetVendorNotes

`func (o *CompanyConfiguration) GetVendorNotes() string`

GetVendorNotes returns the VendorNotes field if non-nil, zero value otherwise.

### GetVendorNotesOk

`func (o *CompanyConfiguration) GetVendorNotesOk() (*string, bool)`

GetVendorNotesOk returns a tuple with the VendorNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorNotes

`func (o *CompanyConfiguration) SetVendorNotes(v string)`

SetVendorNotes sets VendorNotes field to given value.

### HasVendorNotes

`func (o *CompanyConfiguration) HasVendorNotes() bool`

HasVendorNotes returns a boolean if a field has been set.

### GetNotes

`func (o *CompanyConfiguration) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CompanyConfiguration) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CompanyConfiguration) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CompanyConfiguration) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetMacAddress

`func (o *CompanyConfiguration) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *CompanyConfiguration) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *CompanyConfiguration) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *CompanyConfiguration) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetLastLoginName

`func (o *CompanyConfiguration) GetLastLoginName() string`

GetLastLoginName returns the LastLoginName field if non-nil, zero value otherwise.

### GetLastLoginNameOk

`func (o *CompanyConfiguration) GetLastLoginNameOk() (*string, bool)`

GetLastLoginNameOk returns a tuple with the LastLoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginName

`func (o *CompanyConfiguration) SetLastLoginName(v string)`

SetLastLoginName sets LastLoginName field to given value.

### HasLastLoginName

`func (o *CompanyConfiguration) HasLastLoginName() bool`

HasLastLoginName returns a boolean if a field has been set.

### GetBillFlag

`func (o *CompanyConfiguration) GetBillFlag() bool`

GetBillFlag returns the BillFlag field if non-nil, zero value otherwise.

### GetBillFlagOk

`func (o *CompanyConfiguration) GetBillFlagOk() (*bool, bool)`

GetBillFlagOk returns a tuple with the BillFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillFlag

`func (o *CompanyConfiguration) SetBillFlag(v bool)`

SetBillFlag sets BillFlag field to given value.

### HasBillFlag

`func (o *CompanyConfiguration) HasBillFlag() bool`

HasBillFlag returns a boolean if a field has been set.

### SetBillFlagNil

`func (o *CompanyConfiguration) SetBillFlagNil(b bool)`

 SetBillFlagNil sets the value for BillFlag to be an explicit nil

### UnsetBillFlag
`func (o *CompanyConfiguration) UnsetBillFlag()`

UnsetBillFlag ensures that no value is present for BillFlag, not even an explicit nil
### GetBackupSuccesses

`func (o *CompanyConfiguration) GetBackupSuccesses() int32`

GetBackupSuccesses returns the BackupSuccesses field if non-nil, zero value otherwise.

### GetBackupSuccessesOk

`func (o *CompanyConfiguration) GetBackupSuccessesOk() (*int32, bool)`

GetBackupSuccessesOk returns a tuple with the BackupSuccesses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSuccesses

`func (o *CompanyConfiguration) SetBackupSuccesses(v int32)`

SetBackupSuccesses sets BackupSuccesses field to given value.

### HasBackupSuccesses

`func (o *CompanyConfiguration) HasBackupSuccesses() bool`

HasBackupSuccesses returns a boolean if a field has been set.

### SetBackupSuccessesNil

`func (o *CompanyConfiguration) SetBackupSuccessesNil(b bool)`

 SetBackupSuccessesNil sets the value for BackupSuccesses to be an explicit nil

### UnsetBackupSuccesses
`func (o *CompanyConfiguration) UnsetBackupSuccesses()`

UnsetBackupSuccesses ensures that no value is present for BackupSuccesses, not even an explicit nil
### GetBackupIncomplete

`func (o *CompanyConfiguration) GetBackupIncomplete() int32`

GetBackupIncomplete returns the BackupIncomplete field if non-nil, zero value otherwise.

### GetBackupIncompleteOk

`func (o *CompanyConfiguration) GetBackupIncompleteOk() (*int32, bool)`

GetBackupIncompleteOk returns a tuple with the BackupIncomplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupIncomplete

`func (o *CompanyConfiguration) SetBackupIncomplete(v int32)`

SetBackupIncomplete sets BackupIncomplete field to given value.

### HasBackupIncomplete

`func (o *CompanyConfiguration) HasBackupIncomplete() bool`

HasBackupIncomplete returns a boolean if a field has been set.

### SetBackupIncompleteNil

`func (o *CompanyConfiguration) SetBackupIncompleteNil(b bool)`

 SetBackupIncompleteNil sets the value for BackupIncomplete to be an explicit nil

### UnsetBackupIncomplete
`func (o *CompanyConfiguration) UnsetBackupIncomplete()`

UnsetBackupIncomplete ensures that no value is present for BackupIncomplete, not even an explicit nil
### GetBackupFailed

`func (o *CompanyConfiguration) GetBackupFailed() int32`

GetBackupFailed returns the BackupFailed field if non-nil, zero value otherwise.

### GetBackupFailedOk

`func (o *CompanyConfiguration) GetBackupFailedOk() (*int32, bool)`

GetBackupFailedOk returns a tuple with the BackupFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupFailed

`func (o *CompanyConfiguration) SetBackupFailed(v int32)`

SetBackupFailed sets BackupFailed field to given value.

### HasBackupFailed

`func (o *CompanyConfiguration) HasBackupFailed() bool`

HasBackupFailed returns a boolean if a field has been set.

### SetBackupFailedNil

`func (o *CompanyConfiguration) SetBackupFailedNil(b bool)`

 SetBackupFailedNil sets the value for BackupFailed to be an explicit nil

### UnsetBackupFailed
`func (o *CompanyConfiguration) UnsetBackupFailed()`

UnsetBackupFailed ensures that no value is present for BackupFailed, not even an explicit nil
### GetBackupRestores

`func (o *CompanyConfiguration) GetBackupRestores() int32`

GetBackupRestores returns the BackupRestores field if non-nil, zero value otherwise.

### GetBackupRestoresOk

`func (o *CompanyConfiguration) GetBackupRestoresOk() (*int32, bool)`

GetBackupRestoresOk returns a tuple with the BackupRestores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupRestores

`func (o *CompanyConfiguration) SetBackupRestores(v int32)`

SetBackupRestores sets BackupRestores field to given value.

### HasBackupRestores

`func (o *CompanyConfiguration) HasBackupRestores() bool`

HasBackupRestores returns a boolean if a field has been set.

### SetBackupRestoresNil

`func (o *CompanyConfiguration) SetBackupRestoresNil(b bool)`

 SetBackupRestoresNil sets the value for BackupRestores to be an explicit nil

### UnsetBackupRestores
`func (o *CompanyConfiguration) UnsetBackupRestores()`

UnsetBackupRestores ensures that no value is present for BackupRestores, not even an explicit nil
### GetLastBackupDate

`func (o *CompanyConfiguration) GetLastBackupDate() time.Time`

GetLastBackupDate returns the LastBackupDate field if non-nil, zero value otherwise.

### GetLastBackupDateOk

`func (o *CompanyConfiguration) GetLastBackupDateOk() (*time.Time, bool)`

GetLastBackupDateOk returns a tuple with the LastBackupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBackupDate

`func (o *CompanyConfiguration) SetLastBackupDate(v time.Time)`

SetLastBackupDate sets LastBackupDate field to given value.

### HasLastBackupDate

`func (o *CompanyConfiguration) HasLastBackupDate() bool`

HasLastBackupDate returns a boolean if a field has been set.

### GetBackupServerName

`func (o *CompanyConfiguration) GetBackupServerName() string`

GetBackupServerName returns the BackupServerName field if non-nil, zero value otherwise.

### GetBackupServerNameOk

`func (o *CompanyConfiguration) GetBackupServerNameOk() (*string, bool)`

GetBackupServerNameOk returns a tuple with the BackupServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupServerName

`func (o *CompanyConfiguration) SetBackupServerName(v string)`

SetBackupServerName sets BackupServerName field to given value.

### HasBackupServerName

`func (o *CompanyConfiguration) HasBackupServerName() bool`

HasBackupServerName returns a boolean if a field has been set.

### GetBackupBillableSpaceGb

`func (o *CompanyConfiguration) GetBackupBillableSpaceGb() float64`

GetBackupBillableSpaceGb returns the BackupBillableSpaceGb field if non-nil, zero value otherwise.

### GetBackupBillableSpaceGbOk

`func (o *CompanyConfiguration) GetBackupBillableSpaceGbOk() (*float64, bool)`

GetBackupBillableSpaceGbOk returns a tuple with the BackupBillableSpaceGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupBillableSpaceGb

`func (o *CompanyConfiguration) SetBackupBillableSpaceGb(v float64)`

SetBackupBillableSpaceGb sets BackupBillableSpaceGb field to given value.

### HasBackupBillableSpaceGb

`func (o *CompanyConfiguration) HasBackupBillableSpaceGb() bool`

HasBackupBillableSpaceGb returns a boolean if a field has been set.

### SetBackupBillableSpaceGbNil

`func (o *CompanyConfiguration) SetBackupBillableSpaceGbNil(b bool)`

 SetBackupBillableSpaceGbNil sets the value for BackupBillableSpaceGb to be an explicit nil

### UnsetBackupBillableSpaceGb
`func (o *CompanyConfiguration) UnsetBackupBillableSpaceGb()`

UnsetBackupBillableSpaceGb ensures that no value is present for BackupBillableSpaceGb, not even an explicit nil
### GetBackupProtectedDeviceList

`func (o *CompanyConfiguration) GetBackupProtectedDeviceList() string`

GetBackupProtectedDeviceList returns the BackupProtectedDeviceList field if non-nil, zero value otherwise.

### GetBackupProtectedDeviceListOk

`func (o *CompanyConfiguration) GetBackupProtectedDeviceListOk() (*string, bool)`

GetBackupProtectedDeviceListOk returns a tuple with the BackupProtectedDeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupProtectedDeviceList

`func (o *CompanyConfiguration) SetBackupProtectedDeviceList(v string)`

SetBackupProtectedDeviceList sets BackupProtectedDeviceList field to given value.

### HasBackupProtectedDeviceList

`func (o *CompanyConfiguration) HasBackupProtectedDeviceList() bool`

HasBackupProtectedDeviceList returns a boolean if a field has been set.

### GetBackupYear

`func (o *CompanyConfiguration) GetBackupYear() int32`

GetBackupYear returns the BackupYear field if non-nil, zero value otherwise.

### GetBackupYearOk

`func (o *CompanyConfiguration) GetBackupYearOk() (*int32, bool)`

GetBackupYearOk returns a tuple with the BackupYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupYear

`func (o *CompanyConfiguration) SetBackupYear(v int32)`

SetBackupYear sets BackupYear field to given value.

### HasBackupYear

`func (o *CompanyConfiguration) HasBackupYear() bool`

HasBackupYear returns a boolean if a field has been set.

### SetBackupYearNil

`func (o *CompanyConfiguration) SetBackupYearNil(b bool)`

 SetBackupYearNil sets the value for BackupYear to be an explicit nil

### UnsetBackupYear
`func (o *CompanyConfiguration) UnsetBackupYear()`

UnsetBackupYear ensures that no value is present for BackupYear, not even an explicit nil
### GetBackupMonth

`func (o *CompanyConfiguration) GetBackupMonth() int32`

GetBackupMonth returns the BackupMonth field if non-nil, zero value otherwise.

### GetBackupMonthOk

`func (o *CompanyConfiguration) GetBackupMonthOk() (*int32, bool)`

GetBackupMonthOk returns a tuple with the BackupMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupMonth

`func (o *CompanyConfiguration) SetBackupMonth(v int32)`

SetBackupMonth sets BackupMonth field to given value.

### HasBackupMonth

`func (o *CompanyConfiguration) HasBackupMonth() bool`

HasBackupMonth returns a boolean if a field has been set.

### SetBackupMonthNil

`func (o *CompanyConfiguration) SetBackupMonthNil(b bool)`

 SetBackupMonthNil sets the value for BackupMonth to be an explicit nil

### UnsetBackupMonth
`func (o *CompanyConfiguration) UnsetBackupMonth()`

UnsetBackupMonth ensures that no value is present for BackupMonth, not even an explicit nil
### GetIpAddress

`func (o *CompanyConfiguration) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CompanyConfiguration) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CompanyConfiguration) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CompanyConfiguration) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetDefaultGateway

`func (o *CompanyConfiguration) GetDefaultGateway() string`

GetDefaultGateway returns the DefaultGateway field if non-nil, zero value otherwise.

### GetDefaultGatewayOk

`func (o *CompanyConfiguration) GetDefaultGatewayOk() (*string, bool)`

GetDefaultGatewayOk returns a tuple with the DefaultGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGateway

`func (o *CompanyConfiguration) SetDefaultGateway(v string)`

SetDefaultGateway sets DefaultGateway field to given value.

### HasDefaultGateway

`func (o *CompanyConfiguration) HasDefaultGateway() bool`

HasDefaultGateway returns a boolean if a field has been set.

### GetOsType

`func (o *CompanyConfiguration) GetOsType() string`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *CompanyConfiguration) GetOsTypeOk() (*string, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *CompanyConfiguration) SetOsType(v string)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *CompanyConfiguration) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetOsInfo

`func (o *CompanyConfiguration) GetOsInfo() string`

GetOsInfo returns the OsInfo field if non-nil, zero value otherwise.

### GetOsInfoOk

`func (o *CompanyConfiguration) GetOsInfoOk() (*string, bool)`

GetOsInfoOk returns a tuple with the OsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsInfo

`func (o *CompanyConfiguration) SetOsInfo(v string)`

SetOsInfo sets OsInfo field to given value.

### HasOsInfo

`func (o *CompanyConfiguration) HasOsInfo() bool`

HasOsInfo returns a boolean if a field has been set.

### GetCpuSpeed

`func (o *CompanyConfiguration) GetCpuSpeed() string`

GetCpuSpeed returns the CpuSpeed field if non-nil, zero value otherwise.

### GetCpuSpeedOk

`func (o *CompanyConfiguration) GetCpuSpeedOk() (*string, bool)`

GetCpuSpeedOk returns a tuple with the CpuSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSpeed

`func (o *CompanyConfiguration) SetCpuSpeed(v string)`

SetCpuSpeed sets CpuSpeed field to given value.

### HasCpuSpeed

`func (o *CompanyConfiguration) HasCpuSpeed() bool`

HasCpuSpeed returns a boolean if a field has been set.

### GetRam

`func (o *CompanyConfiguration) GetRam() string`

GetRam returns the Ram field if non-nil, zero value otherwise.

### GetRamOk

`func (o *CompanyConfiguration) GetRamOk() (*string, bool)`

GetRamOk returns a tuple with the Ram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRam

`func (o *CompanyConfiguration) SetRam(v string)`

SetRam sets Ram field to given value.

### HasRam

`func (o *CompanyConfiguration) HasRam() bool`

HasRam returns a boolean if a field has been set.

### GetLocalHardDrives

`func (o *CompanyConfiguration) GetLocalHardDrives() string`

GetLocalHardDrives returns the LocalHardDrives field if non-nil, zero value otherwise.

### GetLocalHardDrivesOk

`func (o *CompanyConfiguration) GetLocalHardDrivesOk() (*string, bool)`

GetLocalHardDrivesOk returns a tuple with the LocalHardDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalHardDrives

`func (o *CompanyConfiguration) SetLocalHardDrives(v string)`

SetLocalHardDrives sets LocalHardDrives field to given value.

### HasLocalHardDrives

`func (o *CompanyConfiguration) HasLocalHardDrives() bool`

HasLocalHardDrives returns a boolean if a field has been set.

### GetParentConfigurationId

`func (o *CompanyConfiguration) GetParentConfigurationId() int32`

GetParentConfigurationId returns the ParentConfigurationId field if non-nil, zero value otherwise.

### GetParentConfigurationIdOk

`func (o *CompanyConfiguration) GetParentConfigurationIdOk() (*int32, bool)`

GetParentConfigurationIdOk returns a tuple with the ParentConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentConfigurationId

`func (o *CompanyConfiguration) SetParentConfigurationId(v int32)`

SetParentConfigurationId sets ParentConfigurationId field to given value.

### HasParentConfigurationId

`func (o *CompanyConfiguration) HasParentConfigurationId() bool`

HasParentConfigurationId returns a boolean if a field has been set.

### SetParentConfigurationIdNil

`func (o *CompanyConfiguration) SetParentConfigurationIdNil(b bool)`

 SetParentConfigurationIdNil sets the value for ParentConfigurationId to be an explicit nil

### UnsetParentConfigurationId
`func (o *CompanyConfiguration) UnsetParentConfigurationId()`

UnsetParentConfigurationId ensures that no value is present for ParentConfigurationId, not even an explicit nil
### GetVendor

`func (o *CompanyConfiguration) GetVendor() CompanyReference`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CompanyConfiguration) GetVendorOk() (*CompanyReference, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CompanyConfiguration) SetVendor(v CompanyReference)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *CompanyConfiguration) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetManufacturer

`func (o *CompanyConfiguration) GetManufacturer() ManufacturerReference`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CompanyConfiguration) GetManufacturerOk() (*ManufacturerReference, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CompanyConfiguration) SetManufacturer(v ManufacturerReference)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CompanyConfiguration) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetQuestions

`func (o *CompanyConfiguration) GetQuestions() []ConfigurationQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CompanyConfiguration) GetQuestionsOk() (*[]ConfigurationQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CompanyConfiguration) SetQuestions(v []ConfigurationQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *CompanyConfiguration) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetActiveFlag

`func (o *CompanyConfiguration) GetActiveFlag() bool`

GetActiveFlag returns the ActiveFlag field if non-nil, zero value otherwise.

### GetActiveFlagOk

`func (o *CompanyConfiguration) GetActiveFlagOk() (*bool, bool)`

GetActiveFlagOk returns a tuple with the ActiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFlag

`func (o *CompanyConfiguration) SetActiveFlag(v bool)`

SetActiveFlag sets ActiveFlag field to given value.

### HasActiveFlag

`func (o *CompanyConfiguration) HasActiveFlag() bool`

HasActiveFlag returns a boolean if a field has been set.

### SetActiveFlagNil

`func (o *CompanyConfiguration) SetActiveFlagNil(b bool)`

 SetActiveFlagNil sets the value for ActiveFlag to be an explicit nil

### UnsetActiveFlag
`func (o *CompanyConfiguration) UnsetActiveFlag()`

UnsetActiveFlag ensures that no value is present for ActiveFlag, not even an explicit nil
### GetManagementLink

`func (o *CompanyConfiguration) GetManagementLink() string`

GetManagementLink returns the ManagementLink field if non-nil, zero value otherwise.

### GetManagementLinkOk

`func (o *CompanyConfiguration) GetManagementLinkOk() (*string, bool)`

GetManagementLinkOk returns a tuple with the ManagementLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementLink

`func (o *CompanyConfiguration) SetManagementLink(v string)`

SetManagementLink sets ManagementLink field to given value.

### HasManagementLink

`func (o *CompanyConfiguration) HasManagementLink() bool`

HasManagementLink returns a boolean if a field has been set.

### GetRemoteLink

`func (o *CompanyConfiguration) GetRemoteLink() string`

GetRemoteLink returns the RemoteLink field if non-nil, zero value otherwise.

### GetRemoteLinkOk

`func (o *CompanyConfiguration) GetRemoteLinkOk() (*string, bool)`

GetRemoteLinkOk returns a tuple with the RemoteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteLink

`func (o *CompanyConfiguration) SetRemoteLink(v string)`

SetRemoteLink sets RemoteLink field to given value.

### HasRemoteLink

`func (o *CompanyConfiguration) HasRemoteLink() bool`

HasRemoteLink returns a boolean if a field has been set.

### GetSla

`func (o *CompanyConfiguration) GetSla() SLAReference`

GetSla returns the Sla field if non-nil, zero value otherwise.

### GetSlaOk

`func (o *CompanyConfiguration) GetSlaOk() (*SLAReference, bool)`

GetSlaOk returns a tuple with the Sla field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSla

`func (o *CompanyConfiguration) SetSla(v SLAReference)`

SetSla sets Sla field to given value.

### HasSla

`func (o *CompanyConfiguration) HasSla() bool`

HasSla returns a boolean if a field has been set.

### GetMobileGuid

`func (o *CompanyConfiguration) GetMobileGuid() string`

GetMobileGuid returns the MobileGuid field if non-nil, zero value otherwise.

### GetMobileGuidOk

`func (o *CompanyConfiguration) GetMobileGuidOk() (*string, bool)`

GetMobileGuidOk returns a tuple with the MobileGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileGuid

`func (o *CompanyConfiguration) SetMobileGuid(v string)`

SetMobileGuid sets MobileGuid field to given value.

### HasMobileGuid

`func (o *CompanyConfiguration) HasMobileGuid() bool`

HasMobileGuid returns a boolean if a field has been set.

### SetMobileGuidNil

`func (o *CompanyConfiguration) SetMobileGuidNil(b bool)`

 SetMobileGuidNil sets the value for MobileGuid to be an explicit nil

### UnsetMobileGuid
`func (o *CompanyConfiguration) UnsetMobileGuid()`

UnsetMobileGuid ensures that no value is present for MobileGuid, not even an explicit nil
### GetDisplayVendorFlag

`func (o *CompanyConfiguration) GetDisplayVendorFlag() bool`

GetDisplayVendorFlag returns the DisplayVendorFlag field if non-nil, zero value otherwise.

### GetDisplayVendorFlagOk

`func (o *CompanyConfiguration) GetDisplayVendorFlagOk() (*bool, bool)`

GetDisplayVendorFlagOk returns a tuple with the DisplayVendorFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVendorFlag

`func (o *CompanyConfiguration) SetDisplayVendorFlag(v bool)`

SetDisplayVendorFlag sets DisplayVendorFlag field to given value.

### HasDisplayVendorFlag

`func (o *CompanyConfiguration) HasDisplayVendorFlag() bool`

HasDisplayVendorFlag returns a boolean if a field has been set.

### SetDisplayVendorFlagNil

`func (o *CompanyConfiguration) SetDisplayVendorFlagNil(b bool)`

 SetDisplayVendorFlagNil sets the value for DisplayVendorFlag to be an explicit nil

### UnsetDisplayVendorFlag
`func (o *CompanyConfiguration) UnsetDisplayVendorFlag()`

UnsetDisplayVendorFlag ensures that no value is present for DisplayVendorFlag, not even an explicit nil
### GetCompanyLocationId

`func (o *CompanyConfiguration) GetCompanyLocationId() int32`

GetCompanyLocationId returns the CompanyLocationId field if non-nil, zero value otherwise.

### GetCompanyLocationIdOk

`func (o *CompanyConfiguration) GetCompanyLocationIdOk() (*int32, bool)`

GetCompanyLocationIdOk returns a tuple with the CompanyLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyLocationId

`func (o *CompanyConfiguration) SetCompanyLocationId(v int32)`

SetCompanyLocationId sets CompanyLocationId field to given value.

### HasCompanyLocationId

`func (o *CompanyConfiguration) HasCompanyLocationId() bool`

HasCompanyLocationId returns a boolean if a field has been set.

### SetCompanyLocationIdNil

`func (o *CompanyConfiguration) SetCompanyLocationIdNil(b bool)`

 SetCompanyLocationIdNil sets the value for CompanyLocationId to be an explicit nil

### UnsetCompanyLocationId
`func (o *CompanyConfiguration) UnsetCompanyLocationId()`

UnsetCompanyLocationId ensures that no value is present for CompanyLocationId, not even an explicit nil
### GetShowRemoteFlag

`func (o *CompanyConfiguration) GetShowRemoteFlag() bool`

GetShowRemoteFlag returns the ShowRemoteFlag field if non-nil, zero value otherwise.

### GetShowRemoteFlagOk

`func (o *CompanyConfiguration) GetShowRemoteFlagOk() (*bool, bool)`

GetShowRemoteFlagOk returns a tuple with the ShowRemoteFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowRemoteFlag

`func (o *CompanyConfiguration) SetShowRemoteFlag(v bool)`

SetShowRemoteFlag sets ShowRemoteFlag field to given value.

### HasShowRemoteFlag

`func (o *CompanyConfiguration) HasShowRemoteFlag() bool`

HasShowRemoteFlag returns a boolean if a field has been set.

### SetShowRemoteFlagNil

`func (o *CompanyConfiguration) SetShowRemoteFlagNil(b bool)`

 SetShowRemoteFlagNil sets the value for ShowRemoteFlag to be an explicit nil

### UnsetShowRemoteFlag
`func (o *CompanyConfiguration) UnsetShowRemoteFlag()`

UnsetShowRemoteFlag ensures that no value is present for ShowRemoteFlag, not even an explicit nil
### GetShowAutomateFlag

`func (o *CompanyConfiguration) GetShowAutomateFlag() bool`

GetShowAutomateFlag returns the ShowAutomateFlag field if non-nil, zero value otherwise.

### GetShowAutomateFlagOk

`func (o *CompanyConfiguration) GetShowAutomateFlagOk() (*bool, bool)`

GetShowAutomateFlagOk returns a tuple with the ShowAutomateFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutomateFlag

`func (o *CompanyConfiguration) SetShowAutomateFlag(v bool)`

SetShowAutomateFlag sets ShowAutomateFlag field to given value.

### HasShowAutomateFlag

`func (o *CompanyConfiguration) HasShowAutomateFlag() bool`

HasShowAutomateFlag returns a boolean if a field has been set.

### SetShowAutomateFlagNil

`func (o *CompanyConfiguration) SetShowAutomateFlagNil(b bool)`

 SetShowAutomateFlagNil sets the value for ShowAutomateFlag to be an explicit nil

### UnsetShowAutomateFlag
`func (o *CompanyConfiguration) UnsetShowAutomateFlag()`

UnsetShowAutomateFlag ensures that no value is present for ShowAutomateFlag, not even an explicit nil
### GetNeedsRenewalFlag

`func (o *CompanyConfiguration) GetNeedsRenewalFlag() bool`

GetNeedsRenewalFlag returns the NeedsRenewalFlag field if non-nil, zero value otherwise.

### GetNeedsRenewalFlagOk

`func (o *CompanyConfiguration) GetNeedsRenewalFlagOk() (*bool, bool)`

GetNeedsRenewalFlagOk returns a tuple with the NeedsRenewalFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRenewalFlag

`func (o *CompanyConfiguration) SetNeedsRenewalFlag(v bool)`

SetNeedsRenewalFlag sets NeedsRenewalFlag field to given value.

### HasNeedsRenewalFlag

`func (o *CompanyConfiguration) HasNeedsRenewalFlag() bool`

HasNeedsRenewalFlag returns a boolean if a field has been set.

### SetNeedsRenewalFlagNil

`func (o *CompanyConfiguration) SetNeedsRenewalFlagNil(b bool)`

 SetNeedsRenewalFlagNil sets the value for NeedsRenewalFlag to be an explicit nil

### UnsetNeedsRenewalFlag
`func (o *CompanyConfiguration) UnsetNeedsRenewalFlag()`

UnsetNeedsRenewalFlag ensures that no value is present for NeedsRenewalFlag, not even an explicit nil
### GetManufacturerPartNumber

`func (o *CompanyConfiguration) GetManufacturerPartNumber() string`

GetManufacturerPartNumber returns the ManufacturerPartNumber field if non-nil, zero value otherwise.

### GetManufacturerPartNumberOk

`func (o *CompanyConfiguration) GetManufacturerPartNumberOk() (*string, bool)`

GetManufacturerPartNumberOk returns a tuple with the ManufacturerPartNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturerPartNumber

`func (o *CompanyConfiguration) SetManufacturerPartNumber(v string)`

SetManufacturerPartNumber sets ManufacturerPartNumber field to given value.

### HasManufacturerPartNumber

`func (o *CompanyConfiguration) HasManufacturerPartNumber() bool`

HasManufacturerPartNumber returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyConfiguration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyConfiguration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyConfiguration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyConfiguration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *CompanyConfiguration) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CompanyConfiguration) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CompanyConfiguration) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CompanyConfiguration) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


