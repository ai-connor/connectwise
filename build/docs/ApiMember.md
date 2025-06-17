# ApiMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Identifier** | **string** |  Max length: 15; | 
**Name** | Pointer to **string** |  Max length: 30; Required On Updates; | [optional] 
**EmailAddress** | Pointer to **string** |  Max length: 250; | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**InactiveDate** | Pointer to **time.Time** |  | [optional] 
**TimeZone** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**SecurityRole** | Pointer to [**SecurityRoleReference**](SecurityRoleReference.md) |  | [optional] 
**StructureLevel** | Pointer to [**StructureReference**](StructureReference.md) |  | [optional] 
**SecurityLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**DefaultDepartment** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**SalesDefaultLocation** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**ServiceDefaultBoard** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ExcludedServiceBoardIds** | Pointer to **[]int32** |  | [optional] 
**BlockPriceFlag** | Pointer to **NullableBool** |  | [optional] 
**BlockCostFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApiMember

`func NewApiMember(identifier string, ) *ApiMember`

NewApiMember instantiates a new ApiMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMemberWithDefaults

`func NewApiMemberWithDefaults() *ApiMember`

NewApiMemberWithDefaults instantiates a new ApiMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiMember) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiMember) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiMember) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ApiMember) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ApiMember) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ApiMember) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ApiMember) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetName

`func (o *ApiMember) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiMember) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiMember) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiMember) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ApiMember) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ApiMember) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ApiMember) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ApiMember) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetInactiveFlag

`func (o *ApiMember) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *ApiMember) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *ApiMember) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *ApiMember) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *ApiMember) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *ApiMember) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetInactiveDate

`func (o *ApiMember) GetInactiveDate() time.Time`

GetInactiveDate returns the InactiveDate field if non-nil, zero value otherwise.

### GetInactiveDateOk

`func (o *ApiMember) GetInactiveDateOk() (*time.Time, bool)`

GetInactiveDateOk returns a tuple with the InactiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveDate

`func (o *ApiMember) SetInactiveDate(v time.Time)`

SetInactiveDate sets InactiveDate field to given value.

### HasInactiveDate

`func (o *ApiMember) HasInactiveDate() bool`

HasInactiveDate returns a boolean if a field has been set.

### GetTimeZone

`func (o *ApiMember) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *ApiMember) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *ApiMember) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *ApiMember) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetSecurityRole

`func (o *ApiMember) GetSecurityRole() SecurityRoleReference`

GetSecurityRole returns the SecurityRole field if non-nil, zero value otherwise.

### GetSecurityRoleOk

`func (o *ApiMember) GetSecurityRoleOk() (*SecurityRoleReference, bool)`

GetSecurityRoleOk returns a tuple with the SecurityRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRole

`func (o *ApiMember) SetSecurityRole(v SecurityRoleReference)`

SetSecurityRole sets SecurityRole field to given value.

### HasSecurityRole

`func (o *ApiMember) HasSecurityRole() bool`

HasSecurityRole returns a boolean if a field has been set.

### GetStructureLevel

`func (o *ApiMember) GetStructureLevel() StructureReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *ApiMember) GetStructureLevelOk() (*StructureReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *ApiMember) SetStructureLevel(v StructureReference)`

SetStructureLevel sets StructureLevel field to given value.

### HasStructureLevel

`func (o *ApiMember) HasStructureLevel() bool`

HasStructureLevel returns a boolean if a field has been set.

### GetSecurityLocation

`func (o *ApiMember) GetSecurityLocation() SystemLocationReference`

GetSecurityLocation returns the SecurityLocation field if non-nil, zero value otherwise.

### GetSecurityLocationOk

`func (o *ApiMember) GetSecurityLocationOk() (*SystemLocationReference, bool)`

GetSecurityLocationOk returns a tuple with the SecurityLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLocation

`func (o *ApiMember) SetSecurityLocation(v SystemLocationReference)`

SetSecurityLocation sets SecurityLocation field to given value.

### HasSecurityLocation

`func (o *ApiMember) HasSecurityLocation() bool`

HasSecurityLocation returns a boolean if a field has been set.

### GetDefaultLocation

`func (o *ApiMember) GetDefaultLocation() SystemLocationReference`

GetDefaultLocation returns the DefaultLocation field if non-nil, zero value otherwise.

### GetDefaultLocationOk

`func (o *ApiMember) GetDefaultLocationOk() (*SystemLocationReference, bool)`

GetDefaultLocationOk returns a tuple with the DefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocation

`func (o *ApiMember) SetDefaultLocation(v SystemLocationReference)`

SetDefaultLocation sets DefaultLocation field to given value.

### HasDefaultLocation

`func (o *ApiMember) HasDefaultLocation() bool`

HasDefaultLocation returns a boolean if a field has been set.

### GetDefaultDepartment

`func (o *ApiMember) GetDefaultDepartment() SystemDepartmentReference`

GetDefaultDepartment returns the DefaultDepartment field if non-nil, zero value otherwise.

### GetDefaultDepartmentOk

`func (o *ApiMember) GetDefaultDepartmentOk() (*SystemDepartmentReference, bool)`

GetDefaultDepartmentOk returns a tuple with the DefaultDepartment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDepartment

`func (o *ApiMember) SetDefaultDepartment(v SystemDepartmentReference)`

SetDefaultDepartment sets DefaultDepartment field to given value.

### HasDefaultDepartment

`func (o *ApiMember) HasDefaultDepartment() bool`

HasDefaultDepartment returns a boolean if a field has been set.

### GetSalesDefaultLocation

`func (o *ApiMember) GetSalesDefaultLocation() SystemLocationReference`

GetSalesDefaultLocation returns the SalesDefaultLocation field if non-nil, zero value otherwise.

### GetSalesDefaultLocationOk

`func (o *ApiMember) GetSalesDefaultLocationOk() (*SystemLocationReference, bool)`

GetSalesDefaultLocationOk returns a tuple with the SalesDefaultLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesDefaultLocation

`func (o *ApiMember) SetSalesDefaultLocation(v SystemLocationReference)`

SetSalesDefaultLocation sets SalesDefaultLocation field to given value.

### HasSalesDefaultLocation

`func (o *ApiMember) HasSalesDefaultLocation() bool`

HasSalesDefaultLocation returns a boolean if a field has been set.

### GetServiceDefaultBoard

`func (o *ApiMember) GetServiceDefaultBoard() BoardReference`

GetServiceDefaultBoard returns the ServiceDefaultBoard field if non-nil, zero value otherwise.

### GetServiceDefaultBoardOk

`func (o *ApiMember) GetServiceDefaultBoardOk() (*BoardReference, bool)`

GetServiceDefaultBoardOk returns a tuple with the ServiceDefaultBoard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceDefaultBoard

`func (o *ApiMember) SetServiceDefaultBoard(v BoardReference)`

SetServiceDefaultBoard sets ServiceDefaultBoard field to given value.

### HasServiceDefaultBoard

`func (o *ApiMember) HasServiceDefaultBoard() bool`

HasServiceDefaultBoard returns a boolean if a field has been set.

### GetNotes

`func (o *ApiMember) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ApiMember) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ApiMember) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ApiMember) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExcludedServiceBoardIds

`func (o *ApiMember) GetExcludedServiceBoardIds() []int32`

GetExcludedServiceBoardIds returns the ExcludedServiceBoardIds field if non-nil, zero value otherwise.

### GetExcludedServiceBoardIdsOk

`func (o *ApiMember) GetExcludedServiceBoardIdsOk() (*[]int32, bool)`

GetExcludedServiceBoardIdsOk returns a tuple with the ExcludedServiceBoardIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedServiceBoardIds

`func (o *ApiMember) SetExcludedServiceBoardIds(v []int32)`

SetExcludedServiceBoardIds sets ExcludedServiceBoardIds field to given value.

### HasExcludedServiceBoardIds

`func (o *ApiMember) HasExcludedServiceBoardIds() bool`

HasExcludedServiceBoardIds returns a boolean if a field has been set.

### GetBlockPriceFlag

`func (o *ApiMember) GetBlockPriceFlag() bool`

GetBlockPriceFlag returns the BlockPriceFlag field if non-nil, zero value otherwise.

### GetBlockPriceFlagOk

`func (o *ApiMember) GetBlockPriceFlagOk() (*bool, bool)`

GetBlockPriceFlagOk returns a tuple with the BlockPriceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockPriceFlag

`func (o *ApiMember) SetBlockPriceFlag(v bool)`

SetBlockPriceFlag sets BlockPriceFlag field to given value.

### HasBlockPriceFlag

`func (o *ApiMember) HasBlockPriceFlag() bool`

HasBlockPriceFlag returns a boolean if a field has been set.

### SetBlockPriceFlagNil

`func (o *ApiMember) SetBlockPriceFlagNil(b bool)`

 SetBlockPriceFlagNil sets the value for BlockPriceFlag to be an explicit nil

### UnsetBlockPriceFlag
`func (o *ApiMember) UnsetBlockPriceFlag()`

UnsetBlockPriceFlag ensures that no value is present for BlockPriceFlag, not even an explicit nil
### GetBlockCostFlag

`func (o *ApiMember) GetBlockCostFlag() bool`

GetBlockCostFlag returns the BlockCostFlag field if non-nil, zero value otherwise.

### GetBlockCostFlagOk

`func (o *ApiMember) GetBlockCostFlagOk() (*bool, bool)`

GetBlockCostFlagOk returns a tuple with the BlockCostFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCostFlag

`func (o *ApiMember) SetBlockCostFlag(v bool)`

SetBlockCostFlag sets BlockCostFlag field to given value.

### HasBlockCostFlag

`func (o *ApiMember) HasBlockCostFlag() bool`

HasBlockCostFlag returns a boolean if a field has been set.

### SetBlockCostFlagNil

`func (o *ApiMember) SetBlockCostFlagNil(b bool)`

 SetBlockCostFlagNil sets the value for BlockCostFlag to be an explicit nil

### UnsetBlockCostFlag
`func (o *ApiMember) UnsetBlockCostFlag()`

UnsetBlockCostFlag ensures that no value is present for BlockCostFlag, not even an explicit nil
### GetInfo

`func (o *ApiMember) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ApiMember) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ApiMember) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ApiMember) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


