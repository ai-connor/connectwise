# ManagedDevicesIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 30; | 
**Solution** | **string** |  Max length: 30; | 
**PortalUrl** | Pointer to **string** |  Max length: 200; | [optional] 
**LoginBy** | **NullableString** |  | 
**GlobalLoginUsername** | Pointer to **string** | Gets or sets             this is only required when globalLoginFlag &#x3D; true. Max length: 50; | [optional] 
**GlobalLoginPassword** | Pointer to **string** | Gets or sets             this is only required when globalLoginFlag &#x3D; true. Max length: 50; | [optional] 
**DefaultBillingLevel** | **NullableString** |  | 
**ManagementItSetupType** | Pointer to **string** |  | [optional] 
**DefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**IntegratorLogin** | Pointer to [**IntegratorLoginReference**](IntegratorLoginReference.md) |  | [optional] 
**MatchOnSerialNumberFlag** | Pointer to **NullableBool** |  | [optional] 
**DisableNewCrossReferencesFlag** | Pointer to **NullableBool** |  | [optional] 
**ConfigBillCustomerFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagedDevicesIntegration

`func NewManagedDevicesIntegration(name string, solution string, loginBy NullableString, defaultBillingLevel NullableString, ) *ManagedDevicesIntegration`

NewManagedDevicesIntegration instantiates a new ManagedDevicesIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedDevicesIntegrationWithDefaults

`func NewManagedDevicesIntegrationWithDefaults() *ManagedDevicesIntegration`

NewManagedDevicesIntegrationWithDefaults instantiates a new ManagedDevicesIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagedDevicesIntegration) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedDevicesIntegration) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedDevicesIntegration) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagedDevicesIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ManagedDevicesIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedDevicesIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedDevicesIntegration) SetName(v string)`

SetName sets Name field to given value.


### GetSolution

`func (o *ManagedDevicesIntegration) GetSolution() string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *ManagedDevicesIntegration) GetSolutionOk() (*string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *ManagedDevicesIntegration) SetSolution(v string)`

SetSolution sets Solution field to given value.


### GetPortalUrl

`func (o *ManagedDevicesIntegration) GetPortalUrl() string`

GetPortalUrl returns the PortalUrl field if non-nil, zero value otherwise.

### GetPortalUrlOk

`func (o *ManagedDevicesIntegration) GetPortalUrlOk() (*string, bool)`

GetPortalUrlOk returns a tuple with the PortalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalUrl

`func (o *ManagedDevicesIntegration) SetPortalUrl(v string)`

SetPortalUrl sets PortalUrl field to given value.

### HasPortalUrl

`func (o *ManagedDevicesIntegration) HasPortalUrl() bool`

HasPortalUrl returns a boolean if a field has been set.

### GetLoginBy

`func (o *ManagedDevicesIntegration) GetLoginBy() string`

GetLoginBy returns the LoginBy field if non-nil, zero value otherwise.

### GetLoginByOk

`func (o *ManagedDevicesIntegration) GetLoginByOk() (*string, bool)`

GetLoginByOk returns a tuple with the LoginBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBy

`func (o *ManagedDevicesIntegration) SetLoginBy(v string)`

SetLoginBy sets LoginBy field to given value.


### SetLoginByNil

`func (o *ManagedDevicesIntegration) SetLoginByNil(b bool)`

 SetLoginByNil sets the value for LoginBy to be an explicit nil

### UnsetLoginBy
`func (o *ManagedDevicesIntegration) UnsetLoginBy()`

UnsetLoginBy ensures that no value is present for LoginBy, not even an explicit nil
### GetGlobalLoginUsername

`func (o *ManagedDevicesIntegration) GetGlobalLoginUsername() string`

GetGlobalLoginUsername returns the GlobalLoginUsername field if non-nil, zero value otherwise.

### GetGlobalLoginUsernameOk

`func (o *ManagedDevicesIntegration) GetGlobalLoginUsernameOk() (*string, bool)`

GetGlobalLoginUsernameOk returns a tuple with the GlobalLoginUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLoginUsername

`func (o *ManagedDevicesIntegration) SetGlobalLoginUsername(v string)`

SetGlobalLoginUsername sets GlobalLoginUsername field to given value.

### HasGlobalLoginUsername

`func (o *ManagedDevicesIntegration) HasGlobalLoginUsername() bool`

HasGlobalLoginUsername returns a boolean if a field has been set.

### GetGlobalLoginPassword

`func (o *ManagedDevicesIntegration) GetGlobalLoginPassword() string`

GetGlobalLoginPassword returns the GlobalLoginPassword field if non-nil, zero value otherwise.

### GetGlobalLoginPasswordOk

`func (o *ManagedDevicesIntegration) GetGlobalLoginPasswordOk() (*string, bool)`

GetGlobalLoginPasswordOk returns a tuple with the GlobalLoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalLoginPassword

`func (o *ManagedDevicesIntegration) SetGlobalLoginPassword(v string)`

SetGlobalLoginPassword sets GlobalLoginPassword field to given value.

### HasGlobalLoginPassword

`func (o *ManagedDevicesIntegration) HasGlobalLoginPassword() bool`

HasGlobalLoginPassword returns a boolean if a field has been set.

### GetDefaultBillingLevel

`func (o *ManagedDevicesIntegration) GetDefaultBillingLevel() string`

GetDefaultBillingLevel returns the DefaultBillingLevel field if non-nil, zero value otherwise.

### GetDefaultBillingLevelOk

`func (o *ManagedDevicesIntegration) GetDefaultBillingLevelOk() (*string, bool)`

GetDefaultBillingLevelOk returns a tuple with the DefaultBillingLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBillingLevel

`func (o *ManagedDevicesIntegration) SetDefaultBillingLevel(v string)`

SetDefaultBillingLevel sets DefaultBillingLevel field to given value.


### SetDefaultBillingLevelNil

`func (o *ManagedDevicesIntegration) SetDefaultBillingLevelNil(b bool)`

 SetDefaultBillingLevelNil sets the value for DefaultBillingLevel to be an explicit nil

### UnsetDefaultBillingLevel
`func (o *ManagedDevicesIntegration) UnsetDefaultBillingLevel()`

UnsetDefaultBillingLevel ensures that no value is present for DefaultBillingLevel, not even an explicit nil
### GetManagementItSetupType

`func (o *ManagedDevicesIntegration) GetManagementItSetupType() string`

GetManagementItSetupType returns the ManagementItSetupType field if non-nil, zero value otherwise.

### GetManagementItSetupTypeOk

`func (o *ManagedDevicesIntegration) GetManagementItSetupTypeOk() (*string, bool)`

GetManagementItSetupTypeOk returns a tuple with the ManagementItSetupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementItSetupType

`func (o *ManagedDevicesIntegration) SetManagementItSetupType(v string)`

SetManagementItSetupType sets ManagementItSetupType field to given value.

### HasManagementItSetupType

`func (o *ManagedDevicesIntegration) HasManagementItSetupType() bool`

HasManagementItSetupType returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *ManagedDevicesIntegration) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *ManagedDevicesIntegration) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *ManagedDevicesIntegration) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *ManagedDevicesIntegration) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *ManagedDevicesIntegration) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *ManagedDevicesIntegration) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *ManagedDevicesIntegration) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *ManagedDevicesIntegration) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetIntegratorLogin

`func (o *ManagedDevicesIntegration) GetIntegratorLogin() IntegratorLoginReference`

GetIntegratorLogin returns the IntegratorLogin field if non-nil, zero value otherwise.

### GetIntegratorLoginOk

`func (o *ManagedDevicesIntegration) GetIntegratorLoginOk() (*IntegratorLoginReference, bool)`

GetIntegratorLoginOk returns a tuple with the IntegratorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorLogin

`func (o *ManagedDevicesIntegration) SetIntegratorLogin(v IntegratorLoginReference)`

SetIntegratorLogin sets IntegratorLogin field to given value.

### HasIntegratorLogin

`func (o *ManagedDevicesIntegration) HasIntegratorLogin() bool`

HasIntegratorLogin returns a boolean if a field has been set.

### GetMatchOnSerialNumberFlag

`func (o *ManagedDevicesIntegration) GetMatchOnSerialNumberFlag() bool`

GetMatchOnSerialNumberFlag returns the MatchOnSerialNumberFlag field if non-nil, zero value otherwise.

### GetMatchOnSerialNumberFlagOk

`func (o *ManagedDevicesIntegration) GetMatchOnSerialNumberFlagOk() (*bool, bool)`

GetMatchOnSerialNumberFlagOk returns a tuple with the MatchOnSerialNumberFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchOnSerialNumberFlag

`func (o *ManagedDevicesIntegration) SetMatchOnSerialNumberFlag(v bool)`

SetMatchOnSerialNumberFlag sets MatchOnSerialNumberFlag field to given value.

### HasMatchOnSerialNumberFlag

`func (o *ManagedDevicesIntegration) HasMatchOnSerialNumberFlag() bool`

HasMatchOnSerialNumberFlag returns a boolean if a field has been set.

### SetMatchOnSerialNumberFlagNil

`func (o *ManagedDevicesIntegration) SetMatchOnSerialNumberFlagNil(b bool)`

 SetMatchOnSerialNumberFlagNil sets the value for MatchOnSerialNumberFlag to be an explicit nil

### UnsetMatchOnSerialNumberFlag
`func (o *ManagedDevicesIntegration) UnsetMatchOnSerialNumberFlag()`

UnsetMatchOnSerialNumberFlag ensures that no value is present for MatchOnSerialNumberFlag, not even an explicit nil
### GetDisableNewCrossReferencesFlag

`func (o *ManagedDevicesIntegration) GetDisableNewCrossReferencesFlag() bool`

GetDisableNewCrossReferencesFlag returns the DisableNewCrossReferencesFlag field if non-nil, zero value otherwise.

### GetDisableNewCrossReferencesFlagOk

`func (o *ManagedDevicesIntegration) GetDisableNewCrossReferencesFlagOk() (*bool, bool)`

GetDisableNewCrossReferencesFlagOk returns a tuple with the DisableNewCrossReferencesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNewCrossReferencesFlag

`func (o *ManagedDevicesIntegration) SetDisableNewCrossReferencesFlag(v bool)`

SetDisableNewCrossReferencesFlag sets DisableNewCrossReferencesFlag field to given value.

### HasDisableNewCrossReferencesFlag

`func (o *ManagedDevicesIntegration) HasDisableNewCrossReferencesFlag() bool`

HasDisableNewCrossReferencesFlag returns a boolean if a field has been set.

### SetDisableNewCrossReferencesFlagNil

`func (o *ManagedDevicesIntegration) SetDisableNewCrossReferencesFlagNil(b bool)`

 SetDisableNewCrossReferencesFlagNil sets the value for DisableNewCrossReferencesFlag to be an explicit nil

### UnsetDisableNewCrossReferencesFlag
`func (o *ManagedDevicesIntegration) UnsetDisableNewCrossReferencesFlag()`

UnsetDisableNewCrossReferencesFlag ensures that no value is present for DisableNewCrossReferencesFlag, not even an explicit nil
### GetConfigBillCustomerFlag

`func (o *ManagedDevicesIntegration) GetConfigBillCustomerFlag() bool`

GetConfigBillCustomerFlag returns the ConfigBillCustomerFlag field if non-nil, zero value otherwise.

### GetConfigBillCustomerFlagOk

`func (o *ManagedDevicesIntegration) GetConfigBillCustomerFlagOk() (*bool, bool)`

GetConfigBillCustomerFlagOk returns a tuple with the ConfigBillCustomerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBillCustomerFlag

`func (o *ManagedDevicesIntegration) SetConfigBillCustomerFlag(v bool)`

SetConfigBillCustomerFlag sets ConfigBillCustomerFlag field to given value.

### HasConfigBillCustomerFlag

`func (o *ManagedDevicesIntegration) HasConfigBillCustomerFlag() bool`

HasConfigBillCustomerFlag returns a boolean if a field has been set.

### SetConfigBillCustomerFlagNil

`func (o *ManagedDevicesIntegration) SetConfigBillCustomerFlagNil(b bool)`

 SetConfigBillCustomerFlagNil sets the value for ConfigBillCustomerFlag to be an explicit nil

### UnsetConfigBillCustomerFlag
`func (o *ManagedDevicesIntegration) UnsetConfigBillCustomerFlag()`

UnsetConfigBillCustomerFlag ensures that no value is present for ConfigBillCustomerFlag, not even an explicit nil
### GetInfo

`func (o *ManagedDevicesIntegration) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagedDevicesIntegration) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagedDevicesIntegration) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagedDevicesIntegration) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


