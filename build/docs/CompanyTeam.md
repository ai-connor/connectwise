# CompanyTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**TeamRole** | [**TeamRoleReference**](TeamRoleReference.md) |  | 
**Location** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Contact** | Pointer to [**ContactReference**](ContactReference.md) |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**AccountManagerFlag** | Pointer to **NullableBool** |  | [optional] 
**TechFlag** | Pointer to **NullableBool** |  | [optional] 
**SalesFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyTeam

`func NewCompanyTeam(teamRole TeamRoleReference, ) *CompanyTeam`

NewCompanyTeam instantiates a new CompanyTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyTeamWithDefaults

`func NewCompanyTeamWithDefaults() *CompanyTeam`

NewCompanyTeamWithDefaults instantiates a new CompanyTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyTeam) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyTeam) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyTeam) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyTeam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyTeam) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyTeam) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyTeam) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyTeam) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTeamRole

`func (o *CompanyTeam) GetTeamRole() TeamRoleReference`

GetTeamRole returns the TeamRole field if non-nil, zero value otherwise.

### GetTeamRoleOk

`func (o *CompanyTeam) GetTeamRoleOk() (*TeamRoleReference, bool)`

GetTeamRoleOk returns a tuple with the TeamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamRole

`func (o *CompanyTeam) SetTeamRole(v TeamRoleReference)`

SetTeamRole sets TeamRole field to given value.


### GetLocation

`func (o *CompanyTeam) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CompanyTeam) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CompanyTeam) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CompanyTeam) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetDepartment

`func (o *CompanyTeam) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CompanyTeam) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CompanyTeam) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CompanyTeam) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetContact

`func (o *CompanyTeam) GetContact() ContactReference`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CompanyTeam) GetContactOk() (*ContactReference, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CompanyTeam) SetContact(v ContactReference)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CompanyTeam) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetMember

`func (o *CompanyTeam) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *CompanyTeam) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *CompanyTeam) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *CompanyTeam) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetAccountManagerFlag

`func (o *CompanyTeam) GetAccountManagerFlag() bool`

GetAccountManagerFlag returns the AccountManagerFlag field if non-nil, zero value otherwise.

### GetAccountManagerFlagOk

`func (o *CompanyTeam) GetAccountManagerFlagOk() (*bool, bool)`

GetAccountManagerFlagOk returns a tuple with the AccountManagerFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagerFlag

`func (o *CompanyTeam) SetAccountManagerFlag(v bool)`

SetAccountManagerFlag sets AccountManagerFlag field to given value.

### HasAccountManagerFlag

`func (o *CompanyTeam) HasAccountManagerFlag() bool`

HasAccountManagerFlag returns a boolean if a field has been set.

### SetAccountManagerFlagNil

`func (o *CompanyTeam) SetAccountManagerFlagNil(b bool)`

 SetAccountManagerFlagNil sets the value for AccountManagerFlag to be an explicit nil

### UnsetAccountManagerFlag
`func (o *CompanyTeam) UnsetAccountManagerFlag()`

UnsetAccountManagerFlag ensures that no value is present for AccountManagerFlag, not even an explicit nil
### GetTechFlag

`func (o *CompanyTeam) GetTechFlag() bool`

GetTechFlag returns the TechFlag field if non-nil, zero value otherwise.

### GetTechFlagOk

`func (o *CompanyTeam) GetTechFlagOk() (*bool, bool)`

GetTechFlagOk returns a tuple with the TechFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechFlag

`func (o *CompanyTeam) SetTechFlag(v bool)`

SetTechFlag sets TechFlag field to given value.

### HasTechFlag

`func (o *CompanyTeam) HasTechFlag() bool`

HasTechFlag returns a boolean if a field has been set.

### SetTechFlagNil

`func (o *CompanyTeam) SetTechFlagNil(b bool)`

 SetTechFlagNil sets the value for TechFlag to be an explicit nil

### UnsetTechFlag
`func (o *CompanyTeam) UnsetTechFlag()`

UnsetTechFlag ensures that no value is present for TechFlag, not even an explicit nil
### GetSalesFlag

`func (o *CompanyTeam) GetSalesFlag() bool`

GetSalesFlag returns the SalesFlag field if non-nil, zero value otherwise.

### GetSalesFlagOk

`func (o *CompanyTeam) GetSalesFlagOk() (*bool, bool)`

GetSalesFlagOk returns a tuple with the SalesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesFlag

`func (o *CompanyTeam) SetSalesFlag(v bool)`

SetSalesFlag sets SalesFlag field to given value.

### HasSalesFlag

`func (o *CompanyTeam) HasSalesFlag() bool`

HasSalesFlag returns a boolean if a field has been set.

### SetSalesFlagNil

`func (o *CompanyTeam) SetSalesFlagNil(b bool)`

 SetSalesFlagNil sets the value for SalesFlag to be an explicit nil

### UnsetSalesFlag
`func (o *CompanyTeam) UnsetSalesFlag()`

UnsetSalesFlag ensures that no value is present for SalesFlag, not even an explicit nil
### GetInfo

`func (o *CompanyTeam) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyTeam) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyTeam) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyTeam) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


