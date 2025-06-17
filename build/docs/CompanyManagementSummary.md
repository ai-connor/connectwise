# CompanyManagementSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ManagementSolution** | Pointer to [**ManagementSolutionReference**](ManagementSolutionReference.md) |  | [optional] 
**GroupIdentifier** | **string** |  Max length: 100; | 
**DeviceType** | Pointer to **NullableString** | Gets or sets deviceType is required if the managementSolution is Legacy. | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**SnmpMachines** | Pointer to **NullableInt32** |  | [optional] 
**TotalWorkstations** | Pointer to **NullableInt32** |  | [optional] 
**TotalServers** | Pointer to **NullableInt32** |  | [optional] 
**TotalWindowsServers** | Pointer to **NullableInt32** |  | [optional] 
**TotalWindowsWorkstations** | Pointer to **NullableInt32** |  | [optional] 
**TotalManagedMachines** | Pointer to **NullableInt32** |  | [optional] 
**ServersOffline** | Pointer to **NullableInt32** |  | [optional] 
**ServersDiskSpaceLow** | Pointer to **NullableInt32** |  | [optional] 
**FailedBackupJobs** | Pointer to **NullableInt32** |  | [optional] 
**TotalNotifications** | Pointer to **NullableInt32** |  | [optional] 
**SuccessfulBackupJobs** | Pointer to **NullableInt32** |  | [optional] 
**ServerAvailability** | Pointer to **NullableInt32** |  | [optional] 
**VirusesRemoved** | Pointer to **NullableInt32** |  | [optional] 
**SpywareItemsRemoved** | Pointer to **NullableInt32** |  | [optional] 
**WindowsPatchesInstalled** | Pointer to **NullableInt32** |  | [optional] 
**DiskCleanups** | Pointer to **NullableInt32** |  | [optional] 
**DiskDefragmentations** | Pointer to **NullableInt32** |  | [optional] 
**FullyPatchedMachines** | Pointer to **NullableInt32** |  | [optional] 
**MissingOneTwoPatchesMachines** | Pointer to **NullableInt32** |  | [optional] 
**MissingThreeFivePatchesMachines** | Pointer to **NullableInt32** |  | [optional] 
**MissingMoreFivePatchesMachines** | Pointer to **NullableInt32** |  | [optional] 
**MissingUnscannedPatchesMachines** | Pointer to **NullableInt32** |  | [optional] 
**AlertsGenerated** | Pointer to **string** |  | [optional] 
**InternetConnectivity** | Pointer to **NullableFloat64** |  | [optional] 
**DiskSpaceCleanedMb** | Pointer to **NullableInt32** |  | [optional] 
**MissingSecurityPatches** | Pointer to **string** |  | [optional] 
**CpuUtilization** | Pointer to **NullableFloat64** |  | [optional] 
**MemoryUtilization** | Pointer to **NullableFloat64** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyManagementSummary

`func NewCompanyManagementSummary(groupIdentifier string, ) *CompanyManagementSummary`

NewCompanyManagementSummary instantiates a new CompanyManagementSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyManagementSummaryWithDefaults

`func NewCompanyManagementSummaryWithDefaults() *CompanyManagementSummary`

NewCompanyManagementSummaryWithDefaults instantiates a new CompanyManagementSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyManagementSummary) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyManagementSummary) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyManagementSummary) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyManagementSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagementSolution

`func (o *CompanyManagementSummary) GetManagementSolution() ManagementSolutionReference`

GetManagementSolution returns the ManagementSolution field if non-nil, zero value otherwise.

### GetManagementSolutionOk

`func (o *CompanyManagementSummary) GetManagementSolutionOk() (*ManagementSolutionReference, bool)`

GetManagementSolutionOk returns a tuple with the ManagementSolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementSolution

`func (o *CompanyManagementSummary) SetManagementSolution(v ManagementSolutionReference)`

SetManagementSolution sets ManagementSolution field to given value.

### HasManagementSolution

`func (o *CompanyManagementSummary) HasManagementSolution() bool`

HasManagementSolution returns a boolean if a field has been set.

### GetGroupIdentifier

`func (o *CompanyManagementSummary) GetGroupIdentifier() string`

GetGroupIdentifier returns the GroupIdentifier field if non-nil, zero value otherwise.

### GetGroupIdentifierOk

`func (o *CompanyManagementSummary) GetGroupIdentifierOk() (*string, bool)`

GetGroupIdentifierOk returns a tuple with the GroupIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIdentifier

`func (o *CompanyManagementSummary) SetGroupIdentifier(v string)`

SetGroupIdentifier sets GroupIdentifier field to given value.


### GetDeviceType

`func (o *CompanyManagementSummary) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CompanyManagementSummary) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CompanyManagementSummary) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CompanyManagementSummary) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### SetDeviceTypeNil

`func (o *CompanyManagementSummary) SetDeviceTypeNil(b bool)`

 SetDeviceTypeNil sets the value for DeviceType to be an explicit nil

### UnsetDeviceType
`func (o *CompanyManagementSummary) UnsetDeviceType()`

UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
### GetAgreement

`func (o *CompanyManagementSummary) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *CompanyManagementSummary) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *CompanyManagementSummary) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *CompanyManagementSummary) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetSnmpMachines

`func (o *CompanyManagementSummary) GetSnmpMachines() int32`

GetSnmpMachines returns the SnmpMachines field if non-nil, zero value otherwise.

### GetSnmpMachinesOk

`func (o *CompanyManagementSummary) GetSnmpMachinesOk() (*int32, bool)`

GetSnmpMachinesOk returns a tuple with the SnmpMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpMachines

`func (o *CompanyManagementSummary) SetSnmpMachines(v int32)`

SetSnmpMachines sets SnmpMachines field to given value.

### HasSnmpMachines

`func (o *CompanyManagementSummary) HasSnmpMachines() bool`

HasSnmpMachines returns a boolean if a field has been set.

### SetSnmpMachinesNil

`func (o *CompanyManagementSummary) SetSnmpMachinesNil(b bool)`

 SetSnmpMachinesNil sets the value for SnmpMachines to be an explicit nil

### UnsetSnmpMachines
`func (o *CompanyManagementSummary) UnsetSnmpMachines()`

UnsetSnmpMachines ensures that no value is present for SnmpMachines, not even an explicit nil
### GetTotalWorkstations

`func (o *CompanyManagementSummary) GetTotalWorkstations() int32`

GetTotalWorkstations returns the TotalWorkstations field if non-nil, zero value otherwise.

### GetTotalWorkstationsOk

`func (o *CompanyManagementSummary) GetTotalWorkstationsOk() (*int32, bool)`

GetTotalWorkstationsOk returns a tuple with the TotalWorkstations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWorkstations

`func (o *CompanyManagementSummary) SetTotalWorkstations(v int32)`

SetTotalWorkstations sets TotalWorkstations field to given value.

### HasTotalWorkstations

`func (o *CompanyManagementSummary) HasTotalWorkstations() bool`

HasTotalWorkstations returns a boolean if a field has been set.

### SetTotalWorkstationsNil

`func (o *CompanyManagementSummary) SetTotalWorkstationsNil(b bool)`

 SetTotalWorkstationsNil sets the value for TotalWorkstations to be an explicit nil

### UnsetTotalWorkstations
`func (o *CompanyManagementSummary) UnsetTotalWorkstations()`

UnsetTotalWorkstations ensures that no value is present for TotalWorkstations, not even an explicit nil
### GetTotalServers

`func (o *CompanyManagementSummary) GetTotalServers() int32`

GetTotalServers returns the TotalServers field if non-nil, zero value otherwise.

### GetTotalServersOk

`func (o *CompanyManagementSummary) GetTotalServersOk() (*int32, bool)`

GetTotalServersOk returns a tuple with the TotalServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalServers

`func (o *CompanyManagementSummary) SetTotalServers(v int32)`

SetTotalServers sets TotalServers field to given value.

### HasTotalServers

`func (o *CompanyManagementSummary) HasTotalServers() bool`

HasTotalServers returns a boolean if a field has been set.

### SetTotalServersNil

`func (o *CompanyManagementSummary) SetTotalServersNil(b bool)`

 SetTotalServersNil sets the value for TotalServers to be an explicit nil

### UnsetTotalServers
`func (o *CompanyManagementSummary) UnsetTotalServers()`

UnsetTotalServers ensures that no value is present for TotalServers, not even an explicit nil
### GetTotalWindowsServers

`func (o *CompanyManagementSummary) GetTotalWindowsServers() int32`

GetTotalWindowsServers returns the TotalWindowsServers field if non-nil, zero value otherwise.

### GetTotalWindowsServersOk

`func (o *CompanyManagementSummary) GetTotalWindowsServersOk() (*int32, bool)`

GetTotalWindowsServersOk returns a tuple with the TotalWindowsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWindowsServers

`func (o *CompanyManagementSummary) SetTotalWindowsServers(v int32)`

SetTotalWindowsServers sets TotalWindowsServers field to given value.

### HasTotalWindowsServers

`func (o *CompanyManagementSummary) HasTotalWindowsServers() bool`

HasTotalWindowsServers returns a boolean if a field has been set.

### SetTotalWindowsServersNil

`func (o *CompanyManagementSummary) SetTotalWindowsServersNil(b bool)`

 SetTotalWindowsServersNil sets the value for TotalWindowsServers to be an explicit nil

### UnsetTotalWindowsServers
`func (o *CompanyManagementSummary) UnsetTotalWindowsServers()`

UnsetTotalWindowsServers ensures that no value is present for TotalWindowsServers, not even an explicit nil
### GetTotalWindowsWorkstations

`func (o *CompanyManagementSummary) GetTotalWindowsWorkstations() int32`

GetTotalWindowsWorkstations returns the TotalWindowsWorkstations field if non-nil, zero value otherwise.

### GetTotalWindowsWorkstationsOk

`func (o *CompanyManagementSummary) GetTotalWindowsWorkstationsOk() (*int32, bool)`

GetTotalWindowsWorkstationsOk returns a tuple with the TotalWindowsWorkstations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWindowsWorkstations

`func (o *CompanyManagementSummary) SetTotalWindowsWorkstations(v int32)`

SetTotalWindowsWorkstations sets TotalWindowsWorkstations field to given value.

### HasTotalWindowsWorkstations

`func (o *CompanyManagementSummary) HasTotalWindowsWorkstations() bool`

HasTotalWindowsWorkstations returns a boolean if a field has been set.

### SetTotalWindowsWorkstationsNil

`func (o *CompanyManagementSummary) SetTotalWindowsWorkstationsNil(b bool)`

 SetTotalWindowsWorkstationsNil sets the value for TotalWindowsWorkstations to be an explicit nil

### UnsetTotalWindowsWorkstations
`func (o *CompanyManagementSummary) UnsetTotalWindowsWorkstations()`

UnsetTotalWindowsWorkstations ensures that no value is present for TotalWindowsWorkstations, not even an explicit nil
### GetTotalManagedMachines

`func (o *CompanyManagementSummary) GetTotalManagedMachines() int32`

GetTotalManagedMachines returns the TotalManagedMachines field if non-nil, zero value otherwise.

### GetTotalManagedMachinesOk

`func (o *CompanyManagementSummary) GetTotalManagedMachinesOk() (*int32, bool)`

GetTotalManagedMachinesOk returns a tuple with the TotalManagedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalManagedMachines

`func (o *CompanyManagementSummary) SetTotalManagedMachines(v int32)`

SetTotalManagedMachines sets TotalManagedMachines field to given value.

### HasTotalManagedMachines

`func (o *CompanyManagementSummary) HasTotalManagedMachines() bool`

HasTotalManagedMachines returns a boolean if a field has been set.

### SetTotalManagedMachinesNil

`func (o *CompanyManagementSummary) SetTotalManagedMachinesNil(b bool)`

 SetTotalManagedMachinesNil sets the value for TotalManagedMachines to be an explicit nil

### UnsetTotalManagedMachines
`func (o *CompanyManagementSummary) UnsetTotalManagedMachines()`

UnsetTotalManagedMachines ensures that no value is present for TotalManagedMachines, not even an explicit nil
### GetServersOffline

`func (o *CompanyManagementSummary) GetServersOffline() int32`

GetServersOffline returns the ServersOffline field if non-nil, zero value otherwise.

### GetServersOfflineOk

`func (o *CompanyManagementSummary) GetServersOfflineOk() (*int32, bool)`

GetServersOfflineOk returns a tuple with the ServersOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersOffline

`func (o *CompanyManagementSummary) SetServersOffline(v int32)`

SetServersOffline sets ServersOffline field to given value.

### HasServersOffline

`func (o *CompanyManagementSummary) HasServersOffline() bool`

HasServersOffline returns a boolean if a field has been set.

### SetServersOfflineNil

`func (o *CompanyManagementSummary) SetServersOfflineNil(b bool)`

 SetServersOfflineNil sets the value for ServersOffline to be an explicit nil

### UnsetServersOffline
`func (o *CompanyManagementSummary) UnsetServersOffline()`

UnsetServersOffline ensures that no value is present for ServersOffline, not even an explicit nil
### GetServersDiskSpaceLow

`func (o *CompanyManagementSummary) GetServersDiskSpaceLow() int32`

GetServersDiskSpaceLow returns the ServersDiskSpaceLow field if non-nil, zero value otherwise.

### GetServersDiskSpaceLowOk

`func (o *CompanyManagementSummary) GetServersDiskSpaceLowOk() (*int32, bool)`

GetServersDiskSpaceLowOk returns a tuple with the ServersDiskSpaceLow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersDiskSpaceLow

`func (o *CompanyManagementSummary) SetServersDiskSpaceLow(v int32)`

SetServersDiskSpaceLow sets ServersDiskSpaceLow field to given value.

### HasServersDiskSpaceLow

`func (o *CompanyManagementSummary) HasServersDiskSpaceLow() bool`

HasServersDiskSpaceLow returns a boolean if a field has been set.

### SetServersDiskSpaceLowNil

`func (o *CompanyManagementSummary) SetServersDiskSpaceLowNil(b bool)`

 SetServersDiskSpaceLowNil sets the value for ServersDiskSpaceLow to be an explicit nil

### UnsetServersDiskSpaceLow
`func (o *CompanyManagementSummary) UnsetServersDiskSpaceLow()`

UnsetServersDiskSpaceLow ensures that no value is present for ServersDiskSpaceLow, not even an explicit nil
### GetFailedBackupJobs

`func (o *CompanyManagementSummary) GetFailedBackupJobs() int32`

GetFailedBackupJobs returns the FailedBackupJobs field if non-nil, zero value otherwise.

### GetFailedBackupJobsOk

`func (o *CompanyManagementSummary) GetFailedBackupJobsOk() (*int32, bool)`

GetFailedBackupJobsOk returns a tuple with the FailedBackupJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedBackupJobs

`func (o *CompanyManagementSummary) SetFailedBackupJobs(v int32)`

SetFailedBackupJobs sets FailedBackupJobs field to given value.

### HasFailedBackupJobs

`func (o *CompanyManagementSummary) HasFailedBackupJobs() bool`

HasFailedBackupJobs returns a boolean if a field has been set.

### SetFailedBackupJobsNil

`func (o *CompanyManagementSummary) SetFailedBackupJobsNil(b bool)`

 SetFailedBackupJobsNil sets the value for FailedBackupJobs to be an explicit nil

### UnsetFailedBackupJobs
`func (o *CompanyManagementSummary) UnsetFailedBackupJobs()`

UnsetFailedBackupJobs ensures that no value is present for FailedBackupJobs, not even an explicit nil
### GetTotalNotifications

`func (o *CompanyManagementSummary) GetTotalNotifications() int32`

GetTotalNotifications returns the TotalNotifications field if non-nil, zero value otherwise.

### GetTotalNotificationsOk

`func (o *CompanyManagementSummary) GetTotalNotificationsOk() (*int32, bool)`

GetTotalNotificationsOk returns a tuple with the TotalNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNotifications

`func (o *CompanyManagementSummary) SetTotalNotifications(v int32)`

SetTotalNotifications sets TotalNotifications field to given value.

### HasTotalNotifications

`func (o *CompanyManagementSummary) HasTotalNotifications() bool`

HasTotalNotifications returns a boolean if a field has been set.

### SetTotalNotificationsNil

`func (o *CompanyManagementSummary) SetTotalNotificationsNil(b bool)`

 SetTotalNotificationsNil sets the value for TotalNotifications to be an explicit nil

### UnsetTotalNotifications
`func (o *CompanyManagementSummary) UnsetTotalNotifications()`

UnsetTotalNotifications ensures that no value is present for TotalNotifications, not even an explicit nil
### GetSuccessfulBackupJobs

`func (o *CompanyManagementSummary) GetSuccessfulBackupJobs() int32`

GetSuccessfulBackupJobs returns the SuccessfulBackupJobs field if non-nil, zero value otherwise.

### GetSuccessfulBackupJobsOk

`func (o *CompanyManagementSummary) GetSuccessfulBackupJobsOk() (*int32, bool)`

GetSuccessfulBackupJobsOk returns a tuple with the SuccessfulBackupJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulBackupJobs

`func (o *CompanyManagementSummary) SetSuccessfulBackupJobs(v int32)`

SetSuccessfulBackupJobs sets SuccessfulBackupJobs field to given value.

### HasSuccessfulBackupJobs

`func (o *CompanyManagementSummary) HasSuccessfulBackupJobs() bool`

HasSuccessfulBackupJobs returns a boolean if a field has been set.

### SetSuccessfulBackupJobsNil

`func (o *CompanyManagementSummary) SetSuccessfulBackupJobsNil(b bool)`

 SetSuccessfulBackupJobsNil sets the value for SuccessfulBackupJobs to be an explicit nil

### UnsetSuccessfulBackupJobs
`func (o *CompanyManagementSummary) UnsetSuccessfulBackupJobs()`

UnsetSuccessfulBackupJobs ensures that no value is present for SuccessfulBackupJobs, not even an explicit nil
### GetServerAvailability

`func (o *CompanyManagementSummary) GetServerAvailability() int32`

GetServerAvailability returns the ServerAvailability field if non-nil, zero value otherwise.

### GetServerAvailabilityOk

`func (o *CompanyManagementSummary) GetServerAvailabilityOk() (*int32, bool)`

GetServerAvailabilityOk returns a tuple with the ServerAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAvailability

`func (o *CompanyManagementSummary) SetServerAvailability(v int32)`

SetServerAvailability sets ServerAvailability field to given value.

### HasServerAvailability

`func (o *CompanyManagementSummary) HasServerAvailability() bool`

HasServerAvailability returns a boolean if a field has been set.

### SetServerAvailabilityNil

`func (o *CompanyManagementSummary) SetServerAvailabilityNil(b bool)`

 SetServerAvailabilityNil sets the value for ServerAvailability to be an explicit nil

### UnsetServerAvailability
`func (o *CompanyManagementSummary) UnsetServerAvailability()`

UnsetServerAvailability ensures that no value is present for ServerAvailability, not even an explicit nil
### GetVirusesRemoved

`func (o *CompanyManagementSummary) GetVirusesRemoved() int32`

GetVirusesRemoved returns the VirusesRemoved field if non-nil, zero value otherwise.

### GetVirusesRemovedOk

`func (o *CompanyManagementSummary) GetVirusesRemovedOk() (*int32, bool)`

GetVirusesRemovedOk returns a tuple with the VirusesRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirusesRemoved

`func (o *CompanyManagementSummary) SetVirusesRemoved(v int32)`

SetVirusesRemoved sets VirusesRemoved field to given value.

### HasVirusesRemoved

`func (o *CompanyManagementSummary) HasVirusesRemoved() bool`

HasVirusesRemoved returns a boolean if a field has been set.

### SetVirusesRemovedNil

`func (o *CompanyManagementSummary) SetVirusesRemovedNil(b bool)`

 SetVirusesRemovedNil sets the value for VirusesRemoved to be an explicit nil

### UnsetVirusesRemoved
`func (o *CompanyManagementSummary) UnsetVirusesRemoved()`

UnsetVirusesRemoved ensures that no value is present for VirusesRemoved, not even an explicit nil
### GetSpywareItemsRemoved

`func (o *CompanyManagementSummary) GetSpywareItemsRemoved() int32`

GetSpywareItemsRemoved returns the SpywareItemsRemoved field if non-nil, zero value otherwise.

### GetSpywareItemsRemovedOk

`func (o *CompanyManagementSummary) GetSpywareItemsRemovedOk() (*int32, bool)`

GetSpywareItemsRemovedOk returns a tuple with the SpywareItemsRemoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpywareItemsRemoved

`func (o *CompanyManagementSummary) SetSpywareItemsRemoved(v int32)`

SetSpywareItemsRemoved sets SpywareItemsRemoved field to given value.

### HasSpywareItemsRemoved

`func (o *CompanyManagementSummary) HasSpywareItemsRemoved() bool`

HasSpywareItemsRemoved returns a boolean if a field has been set.

### SetSpywareItemsRemovedNil

`func (o *CompanyManagementSummary) SetSpywareItemsRemovedNil(b bool)`

 SetSpywareItemsRemovedNil sets the value for SpywareItemsRemoved to be an explicit nil

### UnsetSpywareItemsRemoved
`func (o *CompanyManagementSummary) UnsetSpywareItemsRemoved()`

UnsetSpywareItemsRemoved ensures that no value is present for SpywareItemsRemoved, not even an explicit nil
### GetWindowsPatchesInstalled

`func (o *CompanyManagementSummary) GetWindowsPatchesInstalled() int32`

GetWindowsPatchesInstalled returns the WindowsPatchesInstalled field if non-nil, zero value otherwise.

### GetWindowsPatchesInstalledOk

`func (o *CompanyManagementSummary) GetWindowsPatchesInstalledOk() (*int32, bool)`

GetWindowsPatchesInstalledOk returns a tuple with the WindowsPatchesInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPatchesInstalled

`func (o *CompanyManagementSummary) SetWindowsPatchesInstalled(v int32)`

SetWindowsPatchesInstalled sets WindowsPatchesInstalled field to given value.

### HasWindowsPatchesInstalled

`func (o *CompanyManagementSummary) HasWindowsPatchesInstalled() bool`

HasWindowsPatchesInstalled returns a boolean if a field has been set.

### SetWindowsPatchesInstalledNil

`func (o *CompanyManagementSummary) SetWindowsPatchesInstalledNil(b bool)`

 SetWindowsPatchesInstalledNil sets the value for WindowsPatchesInstalled to be an explicit nil

### UnsetWindowsPatchesInstalled
`func (o *CompanyManagementSummary) UnsetWindowsPatchesInstalled()`

UnsetWindowsPatchesInstalled ensures that no value is present for WindowsPatchesInstalled, not even an explicit nil
### GetDiskCleanups

`func (o *CompanyManagementSummary) GetDiskCleanups() int32`

GetDiskCleanups returns the DiskCleanups field if non-nil, zero value otherwise.

### GetDiskCleanupsOk

`func (o *CompanyManagementSummary) GetDiskCleanupsOk() (*int32, bool)`

GetDiskCleanupsOk returns a tuple with the DiskCleanups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskCleanups

`func (o *CompanyManagementSummary) SetDiskCleanups(v int32)`

SetDiskCleanups sets DiskCleanups field to given value.

### HasDiskCleanups

`func (o *CompanyManagementSummary) HasDiskCleanups() bool`

HasDiskCleanups returns a boolean if a field has been set.

### SetDiskCleanupsNil

`func (o *CompanyManagementSummary) SetDiskCleanupsNil(b bool)`

 SetDiskCleanupsNil sets the value for DiskCleanups to be an explicit nil

### UnsetDiskCleanups
`func (o *CompanyManagementSummary) UnsetDiskCleanups()`

UnsetDiskCleanups ensures that no value is present for DiskCleanups, not even an explicit nil
### GetDiskDefragmentations

`func (o *CompanyManagementSummary) GetDiskDefragmentations() int32`

GetDiskDefragmentations returns the DiskDefragmentations field if non-nil, zero value otherwise.

### GetDiskDefragmentationsOk

`func (o *CompanyManagementSummary) GetDiskDefragmentationsOk() (*int32, bool)`

GetDiskDefragmentationsOk returns a tuple with the DiskDefragmentations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDefragmentations

`func (o *CompanyManagementSummary) SetDiskDefragmentations(v int32)`

SetDiskDefragmentations sets DiskDefragmentations field to given value.

### HasDiskDefragmentations

`func (o *CompanyManagementSummary) HasDiskDefragmentations() bool`

HasDiskDefragmentations returns a boolean if a field has been set.

### SetDiskDefragmentationsNil

`func (o *CompanyManagementSummary) SetDiskDefragmentationsNil(b bool)`

 SetDiskDefragmentationsNil sets the value for DiskDefragmentations to be an explicit nil

### UnsetDiskDefragmentations
`func (o *CompanyManagementSummary) UnsetDiskDefragmentations()`

UnsetDiskDefragmentations ensures that no value is present for DiskDefragmentations, not even an explicit nil
### GetFullyPatchedMachines

`func (o *CompanyManagementSummary) GetFullyPatchedMachines() int32`

GetFullyPatchedMachines returns the FullyPatchedMachines field if non-nil, zero value otherwise.

### GetFullyPatchedMachinesOk

`func (o *CompanyManagementSummary) GetFullyPatchedMachinesOk() (*int32, bool)`

GetFullyPatchedMachinesOk returns a tuple with the FullyPatchedMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyPatchedMachines

`func (o *CompanyManagementSummary) SetFullyPatchedMachines(v int32)`

SetFullyPatchedMachines sets FullyPatchedMachines field to given value.

### HasFullyPatchedMachines

`func (o *CompanyManagementSummary) HasFullyPatchedMachines() bool`

HasFullyPatchedMachines returns a boolean if a field has been set.

### SetFullyPatchedMachinesNil

`func (o *CompanyManagementSummary) SetFullyPatchedMachinesNil(b bool)`

 SetFullyPatchedMachinesNil sets the value for FullyPatchedMachines to be an explicit nil

### UnsetFullyPatchedMachines
`func (o *CompanyManagementSummary) UnsetFullyPatchedMachines()`

UnsetFullyPatchedMachines ensures that no value is present for FullyPatchedMachines, not even an explicit nil
### GetMissingOneTwoPatchesMachines

`func (o *CompanyManagementSummary) GetMissingOneTwoPatchesMachines() int32`

GetMissingOneTwoPatchesMachines returns the MissingOneTwoPatchesMachines field if non-nil, zero value otherwise.

### GetMissingOneTwoPatchesMachinesOk

`func (o *CompanyManagementSummary) GetMissingOneTwoPatchesMachinesOk() (*int32, bool)`

GetMissingOneTwoPatchesMachinesOk returns a tuple with the MissingOneTwoPatchesMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingOneTwoPatchesMachines

`func (o *CompanyManagementSummary) SetMissingOneTwoPatchesMachines(v int32)`

SetMissingOneTwoPatchesMachines sets MissingOneTwoPatchesMachines field to given value.

### HasMissingOneTwoPatchesMachines

`func (o *CompanyManagementSummary) HasMissingOneTwoPatchesMachines() bool`

HasMissingOneTwoPatchesMachines returns a boolean if a field has been set.

### SetMissingOneTwoPatchesMachinesNil

`func (o *CompanyManagementSummary) SetMissingOneTwoPatchesMachinesNil(b bool)`

 SetMissingOneTwoPatchesMachinesNil sets the value for MissingOneTwoPatchesMachines to be an explicit nil

### UnsetMissingOneTwoPatchesMachines
`func (o *CompanyManagementSummary) UnsetMissingOneTwoPatchesMachines()`

UnsetMissingOneTwoPatchesMachines ensures that no value is present for MissingOneTwoPatchesMachines, not even an explicit nil
### GetMissingThreeFivePatchesMachines

`func (o *CompanyManagementSummary) GetMissingThreeFivePatchesMachines() int32`

GetMissingThreeFivePatchesMachines returns the MissingThreeFivePatchesMachines field if non-nil, zero value otherwise.

### GetMissingThreeFivePatchesMachinesOk

`func (o *CompanyManagementSummary) GetMissingThreeFivePatchesMachinesOk() (*int32, bool)`

GetMissingThreeFivePatchesMachinesOk returns a tuple with the MissingThreeFivePatchesMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingThreeFivePatchesMachines

`func (o *CompanyManagementSummary) SetMissingThreeFivePatchesMachines(v int32)`

SetMissingThreeFivePatchesMachines sets MissingThreeFivePatchesMachines field to given value.

### HasMissingThreeFivePatchesMachines

`func (o *CompanyManagementSummary) HasMissingThreeFivePatchesMachines() bool`

HasMissingThreeFivePatchesMachines returns a boolean if a field has been set.

### SetMissingThreeFivePatchesMachinesNil

`func (o *CompanyManagementSummary) SetMissingThreeFivePatchesMachinesNil(b bool)`

 SetMissingThreeFivePatchesMachinesNil sets the value for MissingThreeFivePatchesMachines to be an explicit nil

### UnsetMissingThreeFivePatchesMachines
`func (o *CompanyManagementSummary) UnsetMissingThreeFivePatchesMachines()`

UnsetMissingThreeFivePatchesMachines ensures that no value is present for MissingThreeFivePatchesMachines, not even an explicit nil
### GetMissingMoreFivePatchesMachines

`func (o *CompanyManagementSummary) GetMissingMoreFivePatchesMachines() int32`

GetMissingMoreFivePatchesMachines returns the MissingMoreFivePatchesMachines field if non-nil, zero value otherwise.

### GetMissingMoreFivePatchesMachinesOk

`func (o *CompanyManagementSummary) GetMissingMoreFivePatchesMachinesOk() (*int32, bool)`

GetMissingMoreFivePatchesMachinesOk returns a tuple with the MissingMoreFivePatchesMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingMoreFivePatchesMachines

`func (o *CompanyManagementSummary) SetMissingMoreFivePatchesMachines(v int32)`

SetMissingMoreFivePatchesMachines sets MissingMoreFivePatchesMachines field to given value.

### HasMissingMoreFivePatchesMachines

`func (o *CompanyManagementSummary) HasMissingMoreFivePatchesMachines() bool`

HasMissingMoreFivePatchesMachines returns a boolean if a field has been set.

### SetMissingMoreFivePatchesMachinesNil

`func (o *CompanyManagementSummary) SetMissingMoreFivePatchesMachinesNil(b bool)`

 SetMissingMoreFivePatchesMachinesNil sets the value for MissingMoreFivePatchesMachines to be an explicit nil

### UnsetMissingMoreFivePatchesMachines
`func (o *CompanyManagementSummary) UnsetMissingMoreFivePatchesMachines()`

UnsetMissingMoreFivePatchesMachines ensures that no value is present for MissingMoreFivePatchesMachines, not even an explicit nil
### GetMissingUnscannedPatchesMachines

`func (o *CompanyManagementSummary) GetMissingUnscannedPatchesMachines() int32`

GetMissingUnscannedPatchesMachines returns the MissingUnscannedPatchesMachines field if non-nil, zero value otherwise.

### GetMissingUnscannedPatchesMachinesOk

`func (o *CompanyManagementSummary) GetMissingUnscannedPatchesMachinesOk() (*int32, bool)`

GetMissingUnscannedPatchesMachinesOk returns a tuple with the MissingUnscannedPatchesMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingUnscannedPatchesMachines

`func (o *CompanyManagementSummary) SetMissingUnscannedPatchesMachines(v int32)`

SetMissingUnscannedPatchesMachines sets MissingUnscannedPatchesMachines field to given value.

### HasMissingUnscannedPatchesMachines

`func (o *CompanyManagementSummary) HasMissingUnscannedPatchesMachines() bool`

HasMissingUnscannedPatchesMachines returns a boolean if a field has been set.

### SetMissingUnscannedPatchesMachinesNil

`func (o *CompanyManagementSummary) SetMissingUnscannedPatchesMachinesNil(b bool)`

 SetMissingUnscannedPatchesMachinesNil sets the value for MissingUnscannedPatchesMachines to be an explicit nil

### UnsetMissingUnscannedPatchesMachines
`func (o *CompanyManagementSummary) UnsetMissingUnscannedPatchesMachines()`

UnsetMissingUnscannedPatchesMachines ensures that no value is present for MissingUnscannedPatchesMachines, not even an explicit nil
### GetAlertsGenerated

`func (o *CompanyManagementSummary) GetAlertsGenerated() string`

GetAlertsGenerated returns the AlertsGenerated field if non-nil, zero value otherwise.

### GetAlertsGeneratedOk

`func (o *CompanyManagementSummary) GetAlertsGeneratedOk() (*string, bool)`

GetAlertsGeneratedOk returns a tuple with the AlertsGenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertsGenerated

`func (o *CompanyManagementSummary) SetAlertsGenerated(v string)`

SetAlertsGenerated sets AlertsGenerated field to given value.

### HasAlertsGenerated

`func (o *CompanyManagementSummary) HasAlertsGenerated() bool`

HasAlertsGenerated returns a boolean if a field has been set.

### GetInternetConnectivity

`func (o *CompanyManagementSummary) GetInternetConnectivity() float64`

GetInternetConnectivity returns the InternetConnectivity field if non-nil, zero value otherwise.

### GetInternetConnectivityOk

`func (o *CompanyManagementSummary) GetInternetConnectivityOk() (*float64, bool)`

GetInternetConnectivityOk returns a tuple with the InternetConnectivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternetConnectivity

`func (o *CompanyManagementSummary) SetInternetConnectivity(v float64)`

SetInternetConnectivity sets InternetConnectivity field to given value.

### HasInternetConnectivity

`func (o *CompanyManagementSummary) HasInternetConnectivity() bool`

HasInternetConnectivity returns a boolean if a field has been set.

### SetInternetConnectivityNil

`func (o *CompanyManagementSummary) SetInternetConnectivityNil(b bool)`

 SetInternetConnectivityNil sets the value for InternetConnectivity to be an explicit nil

### UnsetInternetConnectivity
`func (o *CompanyManagementSummary) UnsetInternetConnectivity()`

UnsetInternetConnectivity ensures that no value is present for InternetConnectivity, not even an explicit nil
### GetDiskSpaceCleanedMb

`func (o *CompanyManagementSummary) GetDiskSpaceCleanedMb() int32`

GetDiskSpaceCleanedMb returns the DiskSpaceCleanedMb field if non-nil, zero value otherwise.

### GetDiskSpaceCleanedMbOk

`func (o *CompanyManagementSummary) GetDiskSpaceCleanedMbOk() (*int32, bool)`

GetDiskSpaceCleanedMbOk returns a tuple with the DiskSpaceCleanedMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSpaceCleanedMb

`func (o *CompanyManagementSummary) SetDiskSpaceCleanedMb(v int32)`

SetDiskSpaceCleanedMb sets DiskSpaceCleanedMb field to given value.

### HasDiskSpaceCleanedMb

`func (o *CompanyManagementSummary) HasDiskSpaceCleanedMb() bool`

HasDiskSpaceCleanedMb returns a boolean if a field has been set.

### SetDiskSpaceCleanedMbNil

`func (o *CompanyManagementSummary) SetDiskSpaceCleanedMbNil(b bool)`

 SetDiskSpaceCleanedMbNil sets the value for DiskSpaceCleanedMb to be an explicit nil

### UnsetDiskSpaceCleanedMb
`func (o *CompanyManagementSummary) UnsetDiskSpaceCleanedMb()`

UnsetDiskSpaceCleanedMb ensures that no value is present for DiskSpaceCleanedMb, not even an explicit nil
### GetMissingSecurityPatches

`func (o *CompanyManagementSummary) GetMissingSecurityPatches() string`

GetMissingSecurityPatches returns the MissingSecurityPatches field if non-nil, zero value otherwise.

### GetMissingSecurityPatchesOk

`func (o *CompanyManagementSummary) GetMissingSecurityPatchesOk() (*string, bool)`

GetMissingSecurityPatchesOk returns a tuple with the MissingSecurityPatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingSecurityPatches

`func (o *CompanyManagementSummary) SetMissingSecurityPatches(v string)`

SetMissingSecurityPatches sets MissingSecurityPatches field to given value.

### HasMissingSecurityPatches

`func (o *CompanyManagementSummary) HasMissingSecurityPatches() bool`

HasMissingSecurityPatches returns a boolean if a field has been set.

### GetCpuUtilization

`func (o *CompanyManagementSummary) GetCpuUtilization() float64`

GetCpuUtilization returns the CpuUtilization field if non-nil, zero value otherwise.

### GetCpuUtilizationOk

`func (o *CompanyManagementSummary) GetCpuUtilizationOk() (*float64, bool)`

GetCpuUtilizationOk returns a tuple with the CpuUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtilization

`func (o *CompanyManagementSummary) SetCpuUtilization(v float64)`

SetCpuUtilization sets CpuUtilization field to given value.

### HasCpuUtilization

`func (o *CompanyManagementSummary) HasCpuUtilization() bool`

HasCpuUtilization returns a boolean if a field has been set.

### SetCpuUtilizationNil

`func (o *CompanyManagementSummary) SetCpuUtilizationNil(b bool)`

 SetCpuUtilizationNil sets the value for CpuUtilization to be an explicit nil

### UnsetCpuUtilization
`func (o *CompanyManagementSummary) UnsetCpuUtilization()`

UnsetCpuUtilization ensures that no value is present for CpuUtilization, not even an explicit nil
### GetMemoryUtilization

`func (o *CompanyManagementSummary) GetMemoryUtilization() float64`

GetMemoryUtilization returns the MemoryUtilization field if non-nil, zero value otherwise.

### GetMemoryUtilizationOk

`func (o *CompanyManagementSummary) GetMemoryUtilizationOk() (*float64, bool)`

GetMemoryUtilizationOk returns a tuple with the MemoryUtilization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtilization

`func (o *CompanyManagementSummary) SetMemoryUtilization(v float64)`

SetMemoryUtilization sets MemoryUtilization field to given value.

### HasMemoryUtilization

`func (o *CompanyManagementSummary) HasMemoryUtilization() bool`

HasMemoryUtilization returns a boolean if a field has been set.

### SetMemoryUtilizationNil

`func (o *CompanyManagementSummary) SetMemoryUtilizationNil(b bool)`

 SetMemoryUtilizationNil sets the value for MemoryUtilization to be an explicit nil

### UnsetMemoryUtilization
`func (o *CompanyManagementSummary) UnsetMemoryUtilization()`

UnsetMemoryUtilization ensures that no value is present for MemoryUtilization, not even an explicit nil
### GetCompany

`func (o *CompanyManagementSummary) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyManagementSummary) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyManagementSummary) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyManagementSummary) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyManagementSummary) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyManagementSummary) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyManagementSummary) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyManagementSummary) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


