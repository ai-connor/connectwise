# TicketSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 80; | 
**VendorType** | **NullableString** |  | 
**IntegratorLogin** | [**IntegratorLoginReference**](IntegratorLoginReference.md) |  | 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Url** | **string** |  | 
**UserName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Psg** | Pointer to **string** |  | [optional] 
**ProblemDescriptionFlag** | Pointer to **NullableBool** |  | [optional] 
**InternalAnalysisFlag** | Pointer to **NullableBool** |  | [optional] 
**ResolutionFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTicketSync

`func NewTicketSync(name string, vendorType NullableString, integratorLogin IntegratorLoginReference, company CompanyReference, url string, ) *TicketSync`

NewTicketSync instantiates a new TicketSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTicketSyncWithDefaults

`func NewTicketSyncWithDefaults() *TicketSync`

NewTicketSyncWithDefaults instantiates a new TicketSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TicketSync) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TicketSync) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TicketSync) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TicketSync) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TicketSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TicketSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TicketSync) SetName(v string)`

SetName sets Name field to given value.


### GetVendorType

`func (o *TicketSync) GetVendorType() string`

GetVendorType returns the VendorType field if non-nil, zero value otherwise.

### GetVendorTypeOk

`func (o *TicketSync) GetVendorTypeOk() (*string, bool)`

GetVendorTypeOk returns a tuple with the VendorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorType

`func (o *TicketSync) SetVendorType(v string)`

SetVendorType sets VendorType field to given value.


### SetVendorTypeNil

`func (o *TicketSync) SetVendorTypeNil(b bool)`

 SetVendorTypeNil sets the value for VendorType to be an explicit nil

### UnsetVendorType
`func (o *TicketSync) UnsetVendorType()`

UnsetVendorType ensures that no value is present for VendorType, not even an explicit nil
### GetIntegratorLogin

`func (o *TicketSync) GetIntegratorLogin() IntegratorLoginReference`

GetIntegratorLogin returns the IntegratorLogin field if non-nil, zero value otherwise.

### GetIntegratorLoginOk

`func (o *TicketSync) GetIntegratorLoginOk() (*IntegratorLoginReference, bool)`

GetIntegratorLoginOk returns a tuple with the IntegratorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorLogin

`func (o *TicketSync) SetIntegratorLogin(v IntegratorLoginReference)`

SetIntegratorLogin sets IntegratorLogin field to given value.


### GetCompany

`func (o *TicketSync) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TicketSync) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TicketSync) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetUrl

`func (o *TicketSync) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TicketSync) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TicketSync) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUserName

`func (o *TicketSync) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *TicketSync) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *TicketSync) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *TicketSync) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPassword

`func (o *TicketSync) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TicketSync) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TicketSync) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TicketSync) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPsg

`func (o *TicketSync) GetPsg() string`

GetPsg returns the Psg field if non-nil, zero value otherwise.

### GetPsgOk

`func (o *TicketSync) GetPsgOk() (*string, bool)`

GetPsgOk returns a tuple with the Psg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsg

`func (o *TicketSync) SetPsg(v string)`

SetPsg sets Psg field to given value.

### HasPsg

`func (o *TicketSync) HasPsg() bool`

HasPsg returns a boolean if a field has been set.

### GetProblemDescriptionFlag

`func (o *TicketSync) GetProblemDescriptionFlag() bool`

GetProblemDescriptionFlag returns the ProblemDescriptionFlag field if non-nil, zero value otherwise.

### GetProblemDescriptionFlagOk

`func (o *TicketSync) GetProblemDescriptionFlagOk() (*bool, bool)`

GetProblemDescriptionFlagOk returns a tuple with the ProblemDescriptionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemDescriptionFlag

`func (o *TicketSync) SetProblemDescriptionFlag(v bool)`

SetProblemDescriptionFlag sets ProblemDescriptionFlag field to given value.

### HasProblemDescriptionFlag

`func (o *TicketSync) HasProblemDescriptionFlag() bool`

HasProblemDescriptionFlag returns a boolean if a field has been set.

### SetProblemDescriptionFlagNil

`func (o *TicketSync) SetProblemDescriptionFlagNil(b bool)`

 SetProblemDescriptionFlagNil sets the value for ProblemDescriptionFlag to be an explicit nil

### UnsetProblemDescriptionFlag
`func (o *TicketSync) UnsetProblemDescriptionFlag()`

UnsetProblemDescriptionFlag ensures that no value is present for ProblemDescriptionFlag, not even an explicit nil
### GetInternalAnalysisFlag

`func (o *TicketSync) GetInternalAnalysisFlag() bool`

GetInternalAnalysisFlag returns the InternalAnalysisFlag field if non-nil, zero value otherwise.

### GetInternalAnalysisFlagOk

`func (o *TicketSync) GetInternalAnalysisFlagOk() (*bool, bool)`

GetInternalAnalysisFlagOk returns a tuple with the InternalAnalysisFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAnalysisFlag

`func (o *TicketSync) SetInternalAnalysisFlag(v bool)`

SetInternalAnalysisFlag sets InternalAnalysisFlag field to given value.

### HasInternalAnalysisFlag

`func (o *TicketSync) HasInternalAnalysisFlag() bool`

HasInternalAnalysisFlag returns a boolean if a field has been set.

### SetInternalAnalysisFlagNil

`func (o *TicketSync) SetInternalAnalysisFlagNil(b bool)`

 SetInternalAnalysisFlagNil sets the value for InternalAnalysisFlag to be an explicit nil

### UnsetInternalAnalysisFlag
`func (o *TicketSync) UnsetInternalAnalysisFlag()`

UnsetInternalAnalysisFlag ensures that no value is present for InternalAnalysisFlag, not even an explicit nil
### GetResolutionFlag

`func (o *TicketSync) GetResolutionFlag() bool`

GetResolutionFlag returns the ResolutionFlag field if non-nil, zero value otherwise.

### GetResolutionFlagOk

`func (o *TicketSync) GetResolutionFlagOk() (*bool, bool)`

GetResolutionFlagOk returns a tuple with the ResolutionFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionFlag

`func (o *TicketSync) SetResolutionFlag(v bool)`

SetResolutionFlag sets ResolutionFlag field to given value.

### HasResolutionFlag

`func (o *TicketSync) HasResolutionFlag() bool`

HasResolutionFlag returns a boolean if a field has been set.

### SetResolutionFlagNil

`func (o *TicketSync) SetResolutionFlagNil(b bool)`

 SetResolutionFlagNil sets the value for ResolutionFlag to be an explicit nil

### UnsetResolutionFlag
`func (o *TicketSync) UnsetResolutionFlag()`

UnsetResolutionFlag ensures that no value is present for ResolutionFlag, not even an explicit nil
### GetInfo

`func (o *TicketSync) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TicketSync) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TicketSync) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TicketSync) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


