# Crm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CompanyListCount** | Pointer to **NullableInt32** |  | [optional] 
**LockProbabilityFlag** | Pointer to **NullableBool** |  | [optional] 
**AccountManagerRole** | [**TeamRoleReference**](TeamRoleReference.md) |  | 
**TechnicalContactRole** | [**TeamRoleReference**](TeamRoleReference.md) |  | 
**SalesRepRole** | [**TeamRoleReference**](TeamRoleReference.md) |  | 
**CompanyIdGenerationFlag** | Pointer to **NullableBool** |  | [optional] 
**ExcludeSpacesFlag** | Pointer to **NullableBool** |  | [optional] 
**Field1Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field2Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field3Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field4Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field5Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field6Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field7Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field8Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field9Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**Field10Caption** | Pointer to **string** |  Max length: 25; | [optional] 
**PrimaryRepCaption** | Pointer to **string** |  Max length: 50; | [optional] 
**SecondaryRepCaption** | Pointer to **string** |  Max length: 50; | [optional] 
**Other1Caption** | Pointer to **string** |  Max length: 50; | [optional] 
**Other2Caption** | Pointer to **string** |  Max length: 50; | [optional] 
**DefaultYear** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCrm

`func NewCrm(accountManagerRole TeamRoleReference, technicalContactRole TeamRoleReference, salesRepRole TeamRoleReference, ) *Crm`

NewCrm instantiates a new Crm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCrmWithDefaults

`func NewCrmWithDefaults() *Crm`

NewCrmWithDefaults instantiates a new Crm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Crm) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Crm) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Crm) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Crm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompanyListCount

`func (o *Crm) GetCompanyListCount() int32`

GetCompanyListCount returns the CompanyListCount field if non-nil, zero value otherwise.

### GetCompanyListCountOk

`func (o *Crm) GetCompanyListCountOk() (*int32, bool)`

GetCompanyListCountOk returns a tuple with the CompanyListCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyListCount

`func (o *Crm) SetCompanyListCount(v int32)`

SetCompanyListCount sets CompanyListCount field to given value.

### HasCompanyListCount

`func (o *Crm) HasCompanyListCount() bool`

HasCompanyListCount returns a boolean if a field has been set.

### SetCompanyListCountNil

`func (o *Crm) SetCompanyListCountNil(b bool)`

 SetCompanyListCountNil sets the value for CompanyListCount to be an explicit nil

### UnsetCompanyListCount
`func (o *Crm) UnsetCompanyListCount()`

UnsetCompanyListCount ensures that no value is present for CompanyListCount, not even an explicit nil
### GetLockProbabilityFlag

`func (o *Crm) GetLockProbabilityFlag() bool`

GetLockProbabilityFlag returns the LockProbabilityFlag field if non-nil, zero value otherwise.

### GetLockProbabilityFlagOk

`func (o *Crm) GetLockProbabilityFlagOk() (*bool, bool)`

GetLockProbabilityFlagOk returns a tuple with the LockProbabilityFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockProbabilityFlag

`func (o *Crm) SetLockProbabilityFlag(v bool)`

SetLockProbabilityFlag sets LockProbabilityFlag field to given value.

### HasLockProbabilityFlag

`func (o *Crm) HasLockProbabilityFlag() bool`

HasLockProbabilityFlag returns a boolean if a field has been set.

### SetLockProbabilityFlagNil

`func (o *Crm) SetLockProbabilityFlagNil(b bool)`

 SetLockProbabilityFlagNil sets the value for LockProbabilityFlag to be an explicit nil

### UnsetLockProbabilityFlag
`func (o *Crm) UnsetLockProbabilityFlag()`

UnsetLockProbabilityFlag ensures that no value is present for LockProbabilityFlag, not even an explicit nil
### GetAccountManagerRole

`func (o *Crm) GetAccountManagerRole() TeamRoleReference`

GetAccountManagerRole returns the AccountManagerRole field if non-nil, zero value otherwise.

### GetAccountManagerRoleOk

`func (o *Crm) GetAccountManagerRoleOk() (*TeamRoleReference, bool)`

GetAccountManagerRoleOk returns a tuple with the AccountManagerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountManagerRole

`func (o *Crm) SetAccountManagerRole(v TeamRoleReference)`

SetAccountManagerRole sets AccountManagerRole field to given value.


### GetTechnicalContactRole

`func (o *Crm) GetTechnicalContactRole() TeamRoleReference`

GetTechnicalContactRole returns the TechnicalContactRole field if non-nil, zero value otherwise.

### GetTechnicalContactRoleOk

`func (o *Crm) GetTechnicalContactRoleOk() (*TeamRoleReference, bool)`

GetTechnicalContactRoleOk returns a tuple with the TechnicalContactRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalContactRole

`func (o *Crm) SetTechnicalContactRole(v TeamRoleReference)`

SetTechnicalContactRole sets TechnicalContactRole field to given value.


### GetSalesRepRole

`func (o *Crm) GetSalesRepRole() TeamRoleReference`

GetSalesRepRole returns the SalesRepRole field if non-nil, zero value otherwise.

### GetSalesRepRoleOk

`func (o *Crm) GetSalesRepRoleOk() (*TeamRoleReference, bool)`

GetSalesRepRoleOk returns a tuple with the SalesRepRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRepRole

`func (o *Crm) SetSalesRepRole(v TeamRoleReference)`

SetSalesRepRole sets SalesRepRole field to given value.


### GetCompanyIdGenerationFlag

`func (o *Crm) GetCompanyIdGenerationFlag() bool`

GetCompanyIdGenerationFlag returns the CompanyIdGenerationFlag field if non-nil, zero value otherwise.

### GetCompanyIdGenerationFlagOk

`func (o *Crm) GetCompanyIdGenerationFlagOk() (*bool, bool)`

GetCompanyIdGenerationFlagOk returns a tuple with the CompanyIdGenerationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyIdGenerationFlag

`func (o *Crm) SetCompanyIdGenerationFlag(v bool)`

SetCompanyIdGenerationFlag sets CompanyIdGenerationFlag field to given value.

### HasCompanyIdGenerationFlag

`func (o *Crm) HasCompanyIdGenerationFlag() bool`

HasCompanyIdGenerationFlag returns a boolean if a field has been set.

### SetCompanyIdGenerationFlagNil

`func (o *Crm) SetCompanyIdGenerationFlagNil(b bool)`

 SetCompanyIdGenerationFlagNil sets the value for CompanyIdGenerationFlag to be an explicit nil

### UnsetCompanyIdGenerationFlag
`func (o *Crm) UnsetCompanyIdGenerationFlag()`

UnsetCompanyIdGenerationFlag ensures that no value is present for CompanyIdGenerationFlag, not even an explicit nil
### GetExcludeSpacesFlag

`func (o *Crm) GetExcludeSpacesFlag() bool`

GetExcludeSpacesFlag returns the ExcludeSpacesFlag field if non-nil, zero value otherwise.

### GetExcludeSpacesFlagOk

`func (o *Crm) GetExcludeSpacesFlagOk() (*bool, bool)`

GetExcludeSpacesFlagOk returns a tuple with the ExcludeSpacesFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSpacesFlag

`func (o *Crm) SetExcludeSpacesFlag(v bool)`

SetExcludeSpacesFlag sets ExcludeSpacesFlag field to given value.

### HasExcludeSpacesFlag

`func (o *Crm) HasExcludeSpacesFlag() bool`

HasExcludeSpacesFlag returns a boolean if a field has been set.

### SetExcludeSpacesFlagNil

`func (o *Crm) SetExcludeSpacesFlagNil(b bool)`

 SetExcludeSpacesFlagNil sets the value for ExcludeSpacesFlag to be an explicit nil

### UnsetExcludeSpacesFlag
`func (o *Crm) UnsetExcludeSpacesFlag()`

UnsetExcludeSpacesFlag ensures that no value is present for ExcludeSpacesFlag, not even an explicit nil
### GetField1Caption

`func (o *Crm) GetField1Caption() string`

GetField1Caption returns the Field1Caption field if non-nil, zero value otherwise.

### GetField1CaptionOk

`func (o *Crm) GetField1CaptionOk() (*string, bool)`

GetField1CaptionOk returns a tuple with the Field1Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1Caption

`func (o *Crm) SetField1Caption(v string)`

SetField1Caption sets Field1Caption field to given value.

### HasField1Caption

`func (o *Crm) HasField1Caption() bool`

HasField1Caption returns a boolean if a field has been set.

### GetField2Caption

`func (o *Crm) GetField2Caption() string`

GetField2Caption returns the Field2Caption field if non-nil, zero value otherwise.

### GetField2CaptionOk

`func (o *Crm) GetField2CaptionOk() (*string, bool)`

GetField2CaptionOk returns a tuple with the Field2Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2Caption

`func (o *Crm) SetField2Caption(v string)`

SetField2Caption sets Field2Caption field to given value.

### HasField2Caption

`func (o *Crm) HasField2Caption() bool`

HasField2Caption returns a boolean if a field has been set.

### GetField3Caption

`func (o *Crm) GetField3Caption() string`

GetField3Caption returns the Field3Caption field if non-nil, zero value otherwise.

### GetField3CaptionOk

`func (o *Crm) GetField3CaptionOk() (*string, bool)`

GetField3CaptionOk returns a tuple with the Field3Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3Caption

`func (o *Crm) SetField3Caption(v string)`

SetField3Caption sets Field3Caption field to given value.

### HasField3Caption

`func (o *Crm) HasField3Caption() bool`

HasField3Caption returns a boolean if a field has been set.

### GetField4Caption

`func (o *Crm) GetField4Caption() string`

GetField4Caption returns the Field4Caption field if non-nil, zero value otherwise.

### GetField4CaptionOk

`func (o *Crm) GetField4CaptionOk() (*string, bool)`

GetField4CaptionOk returns a tuple with the Field4Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4Caption

`func (o *Crm) SetField4Caption(v string)`

SetField4Caption sets Field4Caption field to given value.

### HasField4Caption

`func (o *Crm) HasField4Caption() bool`

HasField4Caption returns a boolean if a field has been set.

### GetField5Caption

`func (o *Crm) GetField5Caption() string`

GetField5Caption returns the Field5Caption field if non-nil, zero value otherwise.

### GetField5CaptionOk

`func (o *Crm) GetField5CaptionOk() (*string, bool)`

GetField5CaptionOk returns a tuple with the Field5Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField5Caption

`func (o *Crm) SetField5Caption(v string)`

SetField5Caption sets Field5Caption field to given value.

### HasField5Caption

`func (o *Crm) HasField5Caption() bool`

HasField5Caption returns a boolean if a field has been set.

### GetField6Caption

`func (o *Crm) GetField6Caption() string`

GetField6Caption returns the Field6Caption field if non-nil, zero value otherwise.

### GetField6CaptionOk

`func (o *Crm) GetField6CaptionOk() (*string, bool)`

GetField6CaptionOk returns a tuple with the Field6Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField6Caption

`func (o *Crm) SetField6Caption(v string)`

SetField6Caption sets Field6Caption field to given value.

### HasField6Caption

`func (o *Crm) HasField6Caption() bool`

HasField6Caption returns a boolean if a field has been set.

### GetField7Caption

`func (o *Crm) GetField7Caption() string`

GetField7Caption returns the Field7Caption field if non-nil, zero value otherwise.

### GetField7CaptionOk

`func (o *Crm) GetField7CaptionOk() (*string, bool)`

GetField7CaptionOk returns a tuple with the Field7Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField7Caption

`func (o *Crm) SetField7Caption(v string)`

SetField7Caption sets Field7Caption field to given value.

### HasField7Caption

`func (o *Crm) HasField7Caption() bool`

HasField7Caption returns a boolean if a field has been set.

### GetField8Caption

`func (o *Crm) GetField8Caption() string`

GetField8Caption returns the Field8Caption field if non-nil, zero value otherwise.

### GetField8CaptionOk

`func (o *Crm) GetField8CaptionOk() (*string, bool)`

GetField8CaptionOk returns a tuple with the Field8Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField8Caption

`func (o *Crm) SetField8Caption(v string)`

SetField8Caption sets Field8Caption field to given value.

### HasField8Caption

`func (o *Crm) HasField8Caption() bool`

HasField8Caption returns a boolean if a field has been set.

### GetField9Caption

`func (o *Crm) GetField9Caption() string`

GetField9Caption returns the Field9Caption field if non-nil, zero value otherwise.

### GetField9CaptionOk

`func (o *Crm) GetField9CaptionOk() (*string, bool)`

GetField9CaptionOk returns a tuple with the Field9Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField9Caption

`func (o *Crm) SetField9Caption(v string)`

SetField9Caption sets Field9Caption field to given value.

### HasField9Caption

`func (o *Crm) HasField9Caption() bool`

HasField9Caption returns a boolean if a field has been set.

### GetField10Caption

`func (o *Crm) GetField10Caption() string`

GetField10Caption returns the Field10Caption field if non-nil, zero value otherwise.

### GetField10CaptionOk

`func (o *Crm) GetField10CaptionOk() (*string, bool)`

GetField10CaptionOk returns a tuple with the Field10Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField10Caption

`func (o *Crm) SetField10Caption(v string)`

SetField10Caption sets Field10Caption field to given value.

### HasField10Caption

`func (o *Crm) HasField10Caption() bool`

HasField10Caption returns a boolean if a field has been set.

### GetPrimaryRepCaption

`func (o *Crm) GetPrimaryRepCaption() string`

GetPrimaryRepCaption returns the PrimaryRepCaption field if non-nil, zero value otherwise.

### GetPrimaryRepCaptionOk

`func (o *Crm) GetPrimaryRepCaptionOk() (*string, bool)`

GetPrimaryRepCaptionOk returns a tuple with the PrimaryRepCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRepCaption

`func (o *Crm) SetPrimaryRepCaption(v string)`

SetPrimaryRepCaption sets PrimaryRepCaption field to given value.

### HasPrimaryRepCaption

`func (o *Crm) HasPrimaryRepCaption() bool`

HasPrimaryRepCaption returns a boolean if a field has been set.

### GetSecondaryRepCaption

`func (o *Crm) GetSecondaryRepCaption() string`

GetSecondaryRepCaption returns the SecondaryRepCaption field if non-nil, zero value otherwise.

### GetSecondaryRepCaptionOk

`func (o *Crm) GetSecondaryRepCaptionOk() (*string, bool)`

GetSecondaryRepCaptionOk returns a tuple with the SecondaryRepCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRepCaption

`func (o *Crm) SetSecondaryRepCaption(v string)`

SetSecondaryRepCaption sets SecondaryRepCaption field to given value.

### HasSecondaryRepCaption

`func (o *Crm) HasSecondaryRepCaption() bool`

HasSecondaryRepCaption returns a boolean if a field has been set.

### GetOther1Caption

`func (o *Crm) GetOther1Caption() string`

GetOther1Caption returns the Other1Caption field if non-nil, zero value otherwise.

### GetOther1CaptionOk

`func (o *Crm) GetOther1CaptionOk() (*string, bool)`

GetOther1CaptionOk returns a tuple with the Other1Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther1Caption

`func (o *Crm) SetOther1Caption(v string)`

SetOther1Caption sets Other1Caption field to given value.

### HasOther1Caption

`func (o *Crm) HasOther1Caption() bool`

HasOther1Caption returns a boolean if a field has been set.

### GetOther2Caption

`func (o *Crm) GetOther2Caption() string`

GetOther2Caption returns the Other2Caption field if non-nil, zero value otherwise.

### GetOther2CaptionOk

`func (o *Crm) GetOther2CaptionOk() (*string, bool)`

GetOther2CaptionOk returns a tuple with the Other2Caption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOther2Caption

`func (o *Crm) SetOther2Caption(v string)`

SetOther2Caption sets Other2Caption field to given value.

### HasOther2Caption

`func (o *Crm) HasOther2Caption() bool`

HasOther2Caption returns a boolean if a field has been set.

### GetDefaultYear

`func (o *Crm) GetDefaultYear() bool`

GetDefaultYear returns the DefaultYear field if non-nil, zero value otherwise.

### GetDefaultYearOk

`func (o *Crm) GetDefaultYearOk() (*bool, bool)`

GetDefaultYearOk returns a tuple with the DefaultYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultYear

`func (o *Crm) SetDefaultYear(v bool)`

SetDefaultYear sets DefaultYear field to given value.

### HasDefaultYear

`func (o *Crm) HasDefaultYear() bool`

HasDefaultYear returns a boolean if a field has been set.

### SetDefaultYearNil

`func (o *Crm) SetDefaultYearNil(b bool)`

 SetDefaultYearNil sets the value for DefaultYear to be an explicit nil

### UnsetDefaultYear
`func (o *Crm) UnsetDefaultYear()`

UnsetDefaultYear ensures that no value is present for DefaultYear, not even an explicit nil
### GetInfo

`func (o *Crm) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Crm) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Crm) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Crm) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


