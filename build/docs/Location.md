# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**OwnerLevelId** | Pointer to **NullableInt32** |  | [optional] 
**StructureLevel** | [**CorporateStructureLevelReference**](CorporateStructureLevelReference.md) |  | 
**Name** | **string** |  Max length: 50; | 
**Manager** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**ReportsTo** | Pointer to [**SystemLocationReference**](SystemLocationReference.md) |  | [optional] 
**SalesRep** | Pointer to **string** |  Max length: 50; | [optional] 
**TimeZoneSetup** | Pointer to [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | [optional] 
**Calendar** | Pointer to [**CalendarReference**](CalendarReference.md) |  | [optional] 
**OverrideAddressLine1** | Pointer to **string** |  Max length: 50; | [optional] 
**OverrideAddressLine2** | Pointer to **string** |  Max length: 50; | [optional] 
**OverrideCity** | Pointer to **string** |  Max length: 50; | [optional] 
**OverrideState** | Pointer to **string** |  Max length: 50; | [optional] 
**OverrideZip** | Pointer to **string** |  Max length: 12; | [optional] 
**OverrideCountry** | Pointer to [**CountryReference**](CountryReference.md) |  | [optional] 
**OverridePhoneNumber** | Pointer to **string** |  Max length: 15; | [optional] 
**OverrideFaxNumber** | Pointer to **string** |  Max length: 15; | [optional] 
**OwaUrl** | Pointer to **string** |  Max length: 100; | [optional] 
**PayrollXref** | Pointer to **string** |  Max length: 10; | [optional] 
**LocationFlag** | Pointer to **NullableBool** |  | [optional] 
**ClientFlag** | Pointer to **NullableBool** |  | [optional] 
**WorkRoleIds** | Pointer to **[]int32** |  | [optional] 
**DepartmentIds** | Pointer to **[]int32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewLocation

`func NewLocation(structureLevel CorporateStructureLevelReference, name string, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Location) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Location) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Location) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Location) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwnerLevelId

`func (o *Location) GetOwnerLevelId() int32`

GetOwnerLevelId returns the OwnerLevelId field if non-nil, zero value otherwise.

### GetOwnerLevelIdOk

`func (o *Location) GetOwnerLevelIdOk() (*int32, bool)`

GetOwnerLevelIdOk returns a tuple with the OwnerLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerLevelId

`func (o *Location) SetOwnerLevelId(v int32)`

SetOwnerLevelId sets OwnerLevelId field to given value.

### HasOwnerLevelId

`func (o *Location) HasOwnerLevelId() bool`

HasOwnerLevelId returns a boolean if a field has been set.

### SetOwnerLevelIdNil

`func (o *Location) SetOwnerLevelIdNil(b bool)`

 SetOwnerLevelIdNil sets the value for OwnerLevelId to be an explicit nil

### UnsetOwnerLevelId
`func (o *Location) UnsetOwnerLevelId()`

UnsetOwnerLevelId ensures that no value is present for OwnerLevelId, not even an explicit nil
### GetStructureLevel

`func (o *Location) GetStructureLevel() CorporateStructureLevelReference`

GetStructureLevel returns the StructureLevel field if non-nil, zero value otherwise.

### GetStructureLevelOk

`func (o *Location) GetStructureLevelOk() (*CorporateStructureLevelReference, bool)`

GetStructureLevelOk returns a tuple with the StructureLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructureLevel

`func (o *Location) SetStructureLevel(v CorporateStructureLevelReference)`

SetStructureLevel sets StructureLevel field to given value.


### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.


### GetManager

`func (o *Location) GetManager() MemberReference`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *Location) GetManagerOk() (*MemberReference, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *Location) SetManager(v MemberReference)`

SetManager sets Manager field to given value.

### HasManager

`func (o *Location) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetReportsTo

`func (o *Location) GetReportsTo() SystemLocationReference`

GetReportsTo returns the ReportsTo field if non-nil, zero value otherwise.

### GetReportsToOk

`func (o *Location) GetReportsToOk() (*SystemLocationReference, bool)`

GetReportsToOk returns a tuple with the ReportsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportsTo

`func (o *Location) SetReportsTo(v SystemLocationReference)`

SetReportsTo sets ReportsTo field to given value.

### HasReportsTo

`func (o *Location) HasReportsTo() bool`

HasReportsTo returns a boolean if a field has been set.

### GetSalesRep

`func (o *Location) GetSalesRep() string`

GetSalesRep returns the SalesRep field if non-nil, zero value otherwise.

### GetSalesRepOk

`func (o *Location) GetSalesRepOk() (*string, bool)`

GetSalesRepOk returns a tuple with the SalesRep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesRep

`func (o *Location) SetSalesRep(v string)`

SetSalesRep sets SalesRep field to given value.

### HasSalesRep

`func (o *Location) HasSalesRep() bool`

HasSalesRep returns a boolean if a field has been set.

### GetTimeZoneSetup

`func (o *Location) GetTimeZoneSetup() TimeZoneSetupReference`

GetTimeZoneSetup returns the TimeZoneSetup field if non-nil, zero value otherwise.

### GetTimeZoneSetupOk

`func (o *Location) GetTimeZoneSetupOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneSetupOk returns a tuple with the TimeZoneSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZoneSetup

`func (o *Location) SetTimeZoneSetup(v TimeZoneSetupReference)`

SetTimeZoneSetup sets TimeZoneSetup field to given value.

### HasTimeZoneSetup

`func (o *Location) HasTimeZoneSetup() bool`

HasTimeZoneSetup returns a boolean if a field has been set.

### GetCalendar

`func (o *Location) GetCalendar() CalendarReference`

GetCalendar returns the Calendar field if non-nil, zero value otherwise.

### GetCalendarOk

`func (o *Location) GetCalendarOk() (*CalendarReference, bool)`

GetCalendarOk returns a tuple with the Calendar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalendar

`func (o *Location) SetCalendar(v CalendarReference)`

SetCalendar sets Calendar field to given value.

### HasCalendar

`func (o *Location) HasCalendar() bool`

HasCalendar returns a boolean if a field has been set.

### GetOverrideAddressLine1

`func (o *Location) GetOverrideAddressLine1() string`

GetOverrideAddressLine1 returns the OverrideAddressLine1 field if non-nil, zero value otherwise.

### GetOverrideAddressLine1Ok

`func (o *Location) GetOverrideAddressLine1Ok() (*string, bool)`

GetOverrideAddressLine1Ok returns a tuple with the OverrideAddressLine1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAddressLine1

`func (o *Location) SetOverrideAddressLine1(v string)`

SetOverrideAddressLine1 sets OverrideAddressLine1 field to given value.

### HasOverrideAddressLine1

`func (o *Location) HasOverrideAddressLine1() bool`

HasOverrideAddressLine1 returns a boolean if a field has been set.

### GetOverrideAddressLine2

`func (o *Location) GetOverrideAddressLine2() string`

GetOverrideAddressLine2 returns the OverrideAddressLine2 field if non-nil, zero value otherwise.

### GetOverrideAddressLine2Ok

`func (o *Location) GetOverrideAddressLine2Ok() (*string, bool)`

GetOverrideAddressLine2Ok returns a tuple with the OverrideAddressLine2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideAddressLine2

`func (o *Location) SetOverrideAddressLine2(v string)`

SetOverrideAddressLine2 sets OverrideAddressLine2 field to given value.

### HasOverrideAddressLine2

`func (o *Location) HasOverrideAddressLine2() bool`

HasOverrideAddressLine2 returns a boolean if a field has been set.

### GetOverrideCity

`func (o *Location) GetOverrideCity() string`

GetOverrideCity returns the OverrideCity field if non-nil, zero value otherwise.

### GetOverrideCityOk

`func (o *Location) GetOverrideCityOk() (*string, bool)`

GetOverrideCityOk returns a tuple with the OverrideCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCity

`func (o *Location) SetOverrideCity(v string)`

SetOverrideCity sets OverrideCity field to given value.

### HasOverrideCity

`func (o *Location) HasOverrideCity() bool`

HasOverrideCity returns a boolean if a field has been set.

### GetOverrideState

`func (o *Location) GetOverrideState() string`

GetOverrideState returns the OverrideState field if non-nil, zero value otherwise.

### GetOverrideStateOk

`func (o *Location) GetOverrideStateOk() (*string, bool)`

GetOverrideStateOk returns a tuple with the OverrideState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideState

`func (o *Location) SetOverrideState(v string)`

SetOverrideState sets OverrideState field to given value.

### HasOverrideState

`func (o *Location) HasOverrideState() bool`

HasOverrideState returns a boolean if a field has been set.

### GetOverrideZip

`func (o *Location) GetOverrideZip() string`

GetOverrideZip returns the OverrideZip field if non-nil, zero value otherwise.

### GetOverrideZipOk

`func (o *Location) GetOverrideZipOk() (*string, bool)`

GetOverrideZipOk returns a tuple with the OverrideZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideZip

`func (o *Location) SetOverrideZip(v string)`

SetOverrideZip sets OverrideZip field to given value.

### HasOverrideZip

`func (o *Location) HasOverrideZip() bool`

HasOverrideZip returns a boolean if a field has been set.

### GetOverrideCountry

`func (o *Location) GetOverrideCountry() CountryReference`

GetOverrideCountry returns the OverrideCountry field if non-nil, zero value otherwise.

### GetOverrideCountryOk

`func (o *Location) GetOverrideCountryOk() (*CountryReference, bool)`

GetOverrideCountryOk returns a tuple with the OverrideCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideCountry

`func (o *Location) SetOverrideCountry(v CountryReference)`

SetOverrideCountry sets OverrideCountry field to given value.

### HasOverrideCountry

`func (o *Location) HasOverrideCountry() bool`

HasOverrideCountry returns a boolean if a field has been set.

### GetOverridePhoneNumber

`func (o *Location) GetOverridePhoneNumber() string`

GetOverridePhoneNumber returns the OverridePhoneNumber field if non-nil, zero value otherwise.

### GetOverridePhoneNumberOk

`func (o *Location) GetOverridePhoneNumberOk() (*string, bool)`

GetOverridePhoneNumberOk returns a tuple with the OverridePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridePhoneNumber

`func (o *Location) SetOverridePhoneNumber(v string)`

SetOverridePhoneNumber sets OverridePhoneNumber field to given value.

### HasOverridePhoneNumber

`func (o *Location) HasOverridePhoneNumber() bool`

HasOverridePhoneNumber returns a boolean if a field has been set.

### GetOverrideFaxNumber

`func (o *Location) GetOverrideFaxNumber() string`

GetOverrideFaxNumber returns the OverrideFaxNumber field if non-nil, zero value otherwise.

### GetOverrideFaxNumberOk

`func (o *Location) GetOverrideFaxNumberOk() (*string, bool)`

GetOverrideFaxNumberOk returns a tuple with the OverrideFaxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideFaxNumber

`func (o *Location) SetOverrideFaxNumber(v string)`

SetOverrideFaxNumber sets OverrideFaxNumber field to given value.

### HasOverrideFaxNumber

`func (o *Location) HasOverrideFaxNumber() bool`

HasOverrideFaxNumber returns a boolean if a field has been set.

### GetOwaUrl

`func (o *Location) GetOwaUrl() string`

GetOwaUrl returns the OwaUrl field if non-nil, zero value otherwise.

### GetOwaUrlOk

`func (o *Location) GetOwaUrlOk() (*string, bool)`

GetOwaUrlOk returns a tuple with the OwaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwaUrl

`func (o *Location) SetOwaUrl(v string)`

SetOwaUrl sets OwaUrl field to given value.

### HasOwaUrl

`func (o *Location) HasOwaUrl() bool`

HasOwaUrl returns a boolean if a field has been set.

### GetPayrollXref

`func (o *Location) GetPayrollXref() string`

GetPayrollXref returns the PayrollXref field if non-nil, zero value otherwise.

### GetPayrollXrefOk

`func (o *Location) GetPayrollXrefOk() (*string, bool)`

GetPayrollXrefOk returns a tuple with the PayrollXref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayrollXref

`func (o *Location) SetPayrollXref(v string)`

SetPayrollXref sets PayrollXref field to given value.

### HasPayrollXref

`func (o *Location) HasPayrollXref() bool`

HasPayrollXref returns a boolean if a field has been set.

### GetLocationFlag

`func (o *Location) GetLocationFlag() bool`

GetLocationFlag returns the LocationFlag field if non-nil, zero value otherwise.

### GetLocationFlagOk

`func (o *Location) GetLocationFlagOk() (*bool, bool)`

GetLocationFlagOk returns a tuple with the LocationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationFlag

`func (o *Location) SetLocationFlag(v bool)`

SetLocationFlag sets LocationFlag field to given value.

### HasLocationFlag

`func (o *Location) HasLocationFlag() bool`

HasLocationFlag returns a boolean if a field has been set.

### SetLocationFlagNil

`func (o *Location) SetLocationFlagNil(b bool)`

 SetLocationFlagNil sets the value for LocationFlag to be an explicit nil

### UnsetLocationFlag
`func (o *Location) UnsetLocationFlag()`

UnsetLocationFlag ensures that no value is present for LocationFlag, not even an explicit nil
### GetClientFlag

`func (o *Location) GetClientFlag() bool`

GetClientFlag returns the ClientFlag field if non-nil, zero value otherwise.

### GetClientFlagOk

`func (o *Location) GetClientFlagOk() (*bool, bool)`

GetClientFlagOk returns a tuple with the ClientFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientFlag

`func (o *Location) SetClientFlag(v bool)`

SetClientFlag sets ClientFlag field to given value.

### HasClientFlag

`func (o *Location) HasClientFlag() bool`

HasClientFlag returns a boolean if a field has been set.

### SetClientFlagNil

`func (o *Location) SetClientFlagNil(b bool)`

 SetClientFlagNil sets the value for ClientFlag to be an explicit nil

### UnsetClientFlag
`func (o *Location) UnsetClientFlag()`

UnsetClientFlag ensures that no value is present for ClientFlag, not even an explicit nil
### GetWorkRoleIds

`func (o *Location) GetWorkRoleIds() []int32`

GetWorkRoleIds returns the WorkRoleIds field if non-nil, zero value otherwise.

### GetWorkRoleIdsOk

`func (o *Location) GetWorkRoleIdsOk() (*[]int32, bool)`

GetWorkRoleIdsOk returns a tuple with the WorkRoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRoleIds

`func (o *Location) SetWorkRoleIds(v []int32)`

SetWorkRoleIds sets WorkRoleIds field to given value.

### HasWorkRoleIds

`func (o *Location) HasWorkRoleIds() bool`

HasWorkRoleIds returns a boolean if a field has been set.

### GetDepartmentIds

`func (o *Location) GetDepartmentIds() []int32`

GetDepartmentIds returns the DepartmentIds field if non-nil, zero value otherwise.

### GetDepartmentIdsOk

`func (o *Location) GetDepartmentIdsOk() (*[]int32, bool)`

GetDepartmentIdsOk returns a tuple with the DepartmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartmentIds

`func (o *Location) SetDepartmentIds(v []int32)`

SetDepartmentIds sets DepartmentIds field to given value.

### HasDepartmentIds

`func (o *Location) HasDepartmentIds() bool`

HasDepartmentIds returns a boolean if a field has been set.

### GetInfo

`func (o *Location) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Location) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Location) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Location) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


