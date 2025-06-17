# CompanyMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToCompanyId** | **int32** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**PrimaryAddress** | Pointer to **NullableString** |  | [optional] 
**PrimaryContact** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Fax** | Pointer to **NullableString** |  | [optional] 
**Website** | Pointer to **NullableString** |  | [optional] 
**Market** | Pointer to **NullableString** |  | [optional] 
**Territory** | Pointer to **NullableString** |  | [optional] 
**Revenue** | Pointer to **NullableString** |  | [optional] 
**RevenueYear** | Pointer to **NullableString** |  | [optional] 
**NumberOfEmployees** | Pointer to **NullableString** |  | [optional] 
**SicCode** | Pointer to **NullableString** |  | [optional] 
**DateAcquired** | Pointer to **NullableString** |  | [optional] 
**TimeZone** | Pointer to **NullableString** |  | [optional] 
**SourceList** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField1** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField2** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField3** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField4** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField5** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField6** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField7** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField8** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField9** | Pointer to **NullableString** |  | [optional] 
**UserDefinedField10** | Pointer to **NullableString** |  | [optional] 
**BillingAddress** | Pointer to **NullableString** |  | [optional] 
**BillingContact** | Pointer to **NullableString** |  | [optional] 
**TaxCode** | Pointer to **NullableString** |  | [optional] 
**AccountNumber** | Pointer to **NullableString** |  | [optional] 
**BillingTerms** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Sites** | Pointer to **NullableString** |  | [optional] 
**Activities** | Pointer to **NullableString** |  | [optional] 
**Opportunities** | Pointer to **NullableString** |  | [optional] 
**Services** | Pointer to **NullableString** |  | [optional] 
**Projects** | Pointer to **NullableString** |  | [optional] 
**Contacts** | Pointer to **NullableString** |  | [optional] 
**Documents** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCompanyMerge

`func NewCompanyMerge(toCompanyId int32, ) *CompanyMerge`

NewCompanyMerge instantiates a new CompanyMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyMergeWithDefaults

`func NewCompanyMergeWithDefaults() *CompanyMerge`

NewCompanyMergeWithDefaults instantiates a new CompanyMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToCompanyId

`func (o *CompanyMerge) GetToCompanyId() int32`

GetToCompanyId returns the ToCompanyId field if non-nil, zero value otherwise.

### GetToCompanyIdOk

`func (o *CompanyMerge) GetToCompanyIdOk() (*int32, bool)`

GetToCompanyIdOk returns a tuple with the ToCompanyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCompanyId

`func (o *CompanyMerge) SetToCompanyId(v int32)`

SetToCompanyId sets ToCompanyId field to given value.


### GetName

`func (o *CompanyMerge) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CompanyMerge) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CompanyMerge) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CompanyMerge) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CompanyMerge) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CompanyMerge) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIdentifier

`func (o *CompanyMerge) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CompanyMerge) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CompanyMerge) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CompanyMerge) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *CompanyMerge) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *CompanyMerge) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetStatus

`func (o *CompanyMerge) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyMerge) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyMerge) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompanyMerge) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *CompanyMerge) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CompanyMerge) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *CompanyMerge) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompanyMerge) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompanyMerge) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CompanyMerge) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *CompanyMerge) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *CompanyMerge) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetPrimaryAddress

`func (o *CompanyMerge) GetPrimaryAddress() string`

GetPrimaryAddress returns the PrimaryAddress field if non-nil, zero value otherwise.

### GetPrimaryAddressOk

`func (o *CompanyMerge) GetPrimaryAddressOk() (*string, bool)`

GetPrimaryAddressOk returns a tuple with the PrimaryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAddress

`func (o *CompanyMerge) SetPrimaryAddress(v string)`

SetPrimaryAddress sets PrimaryAddress field to given value.

### HasPrimaryAddress

`func (o *CompanyMerge) HasPrimaryAddress() bool`

HasPrimaryAddress returns a boolean if a field has been set.

### SetPrimaryAddressNil

`func (o *CompanyMerge) SetPrimaryAddressNil(b bool)`

 SetPrimaryAddressNil sets the value for PrimaryAddress to be an explicit nil

### UnsetPrimaryAddress
`func (o *CompanyMerge) UnsetPrimaryAddress()`

UnsetPrimaryAddress ensures that no value is present for PrimaryAddress, not even an explicit nil
### GetPrimaryContact

`func (o *CompanyMerge) GetPrimaryContact() string`

GetPrimaryContact returns the PrimaryContact field if non-nil, zero value otherwise.

### GetPrimaryContactOk

`func (o *CompanyMerge) GetPrimaryContactOk() (*string, bool)`

GetPrimaryContactOk returns a tuple with the PrimaryContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryContact

`func (o *CompanyMerge) SetPrimaryContact(v string)`

SetPrimaryContact sets PrimaryContact field to given value.

### HasPrimaryContact

`func (o *CompanyMerge) HasPrimaryContact() bool`

HasPrimaryContact returns a boolean if a field has been set.

### SetPrimaryContactNil

`func (o *CompanyMerge) SetPrimaryContactNil(b bool)`

 SetPrimaryContactNil sets the value for PrimaryContact to be an explicit nil

### UnsetPrimaryContact
`func (o *CompanyMerge) UnsetPrimaryContact()`

UnsetPrimaryContact ensures that no value is present for PrimaryContact, not even an explicit nil
### GetPhone

`func (o *CompanyMerge) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CompanyMerge) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CompanyMerge) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CompanyMerge) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *CompanyMerge) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *CompanyMerge) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetFax

`func (o *CompanyMerge) GetFax() string`

GetFax returns the Fax field if non-nil, zero value otherwise.

### GetFaxOk

`func (o *CompanyMerge) GetFaxOk() (*string, bool)`

GetFaxOk returns a tuple with the Fax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFax

`func (o *CompanyMerge) SetFax(v string)`

SetFax sets Fax field to given value.

### HasFax

`func (o *CompanyMerge) HasFax() bool`

HasFax returns a boolean if a field has been set.

### SetFaxNil

`func (o *CompanyMerge) SetFaxNil(b bool)`

 SetFaxNil sets the value for Fax to be an explicit nil

### UnsetFax
`func (o *CompanyMerge) UnsetFax()`

UnsetFax ensures that no value is present for Fax, not even an explicit nil
### GetWebsite

`func (o *CompanyMerge) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CompanyMerge) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CompanyMerge) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CompanyMerge) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### SetWebsiteNil

`func (o *CompanyMerge) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *CompanyMerge) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil
### GetMarket

`func (o *CompanyMerge) GetMarket() string`

GetMarket returns the Market field if non-nil, zero value otherwise.

### GetMarketOk

`func (o *CompanyMerge) GetMarketOk() (*string, bool)`

GetMarketOk returns a tuple with the Market field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarket

`func (o *CompanyMerge) SetMarket(v string)`

SetMarket sets Market field to given value.

### HasMarket

`func (o *CompanyMerge) HasMarket() bool`

HasMarket returns a boolean if a field has been set.

### SetMarketNil

`func (o *CompanyMerge) SetMarketNil(b bool)`

 SetMarketNil sets the value for Market to be an explicit nil

### UnsetMarket
`func (o *CompanyMerge) UnsetMarket()`

UnsetMarket ensures that no value is present for Market, not even an explicit nil
### GetTerritory

`func (o *CompanyMerge) GetTerritory() string`

GetTerritory returns the Territory field if non-nil, zero value otherwise.

### GetTerritoryOk

`func (o *CompanyMerge) GetTerritoryOk() (*string, bool)`

GetTerritoryOk returns a tuple with the Territory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerritory

`func (o *CompanyMerge) SetTerritory(v string)`

SetTerritory sets Territory field to given value.

### HasTerritory

`func (o *CompanyMerge) HasTerritory() bool`

HasTerritory returns a boolean if a field has been set.

### SetTerritoryNil

`func (o *CompanyMerge) SetTerritoryNil(b bool)`

 SetTerritoryNil sets the value for Territory to be an explicit nil

### UnsetTerritory
`func (o *CompanyMerge) UnsetTerritory()`

UnsetTerritory ensures that no value is present for Territory, not even an explicit nil
### GetRevenue

`func (o *CompanyMerge) GetRevenue() string`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *CompanyMerge) GetRevenueOk() (*string, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *CompanyMerge) SetRevenue(v string)`

SetRevenue sets Revenue field to given value.

### HasRevenue

`func (o *CompanyMerge) HasRevenue() bool`

HasRevenue returns a boolean if a field has been set.

### SetRevenueNil

`func (o *CompanyMerge) SetRevenueNil(b bool)`

 SetRevenueNil sets the value for Revenue to be an explicit nil

### UnsetRevenue
`func (o *CompanyMerge) UnsetRevenue()`

UnsetRevenue ensures that no value is present for Revenue, not even an explicit nil
### GetRevenueYear

`func (o *CompanyMerge) GetRevenueYear() string`

GetRevenueYear returns the RevenueYear field if non-nil, zero value otherwise.

### GetRevenueYearOk

`func (o *CompanyMerge) GetRevenueYearOk() (*string, bool)`

GetRevenueYearOk returns a tuple with the RevenueYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueYear

`func (o *CompanyMerge) SetRevenueYear(v string)`

SetRevenueYear sets RevenueYear field to given value.

### HasRevenueYear

`func (o *CompanyMerge) HasRevenueYear() bool`

HasRevenueYear returns a boolean if a field has been set.

### SetRevenueYearNil

`func (o *CompanyMerge) SetRevenueYearNil(b bool)`

 SetRevenueYearNil sets the value for RevenueYear to be an explicit nil

### UnsetRevenueYear
`func (o *CompanyMerge) UnsetRevenueYear()`

UnsetRevenueYear ensures that no value is present for RevenueYear, not even an explicit nil
### GetNumberOfEmployees

`func (o *CompanyMerge) GetNumberOfEmployees() string`

GetNumberOfEmployees returns the NumberOfEmployees field if non-nil, zero value otherwise.

### GetNumberOfEmployeesOk

`func (o *CompanyMerge) GetNumberOfEmployeesOk() (*string, bool)`

GetNumberOfEmployeesOk returns a tuple with the NumberOfEmployees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfEmployees

`func (o *CompanyMerge) SetNumberOfEmployees(v string)`

SetNumberOfEmployees sets NumberOfEmployees field to given value.

### HasNumberOfEmployees

`func (o *CompanyMerge) HasNumberOfEmployees() bool`

HasNumberOfEmployees returns a boolean if a field has been set.

### SetNumberOfEmployeesNil

`func (o *CompanyMerge) SetNumberOfEmployeesNil(b bool)`

 SetNumberOfEmployeesNil sets the value for NumberOfEmployees to be an explicit nil

### UnsetNumberOfEmployees
`func (o *CompanyMerge) UnsetNumberOfEmployees()`

UnsetNumberOfEmployees ensures that no value is present for NumberOfEmployees, not even an explicit nil
### GetSicCode

`func (o *CompanyMerge) GetSicCode() string`

GetSicCode returns the SicCode field if non-nil, zero value otherwise.

### GetSicCodeOk

`func (o *CompanyMerge) GetSicCodeOk() (*string, bool)`

GetSicCodeOk returns a tuple with the SicCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSicCode

`func (o *CompanyMerge) SetSicCode(v string)`

SetSicCode sets SicCode field to given value.

### HasSicCode

`func (o *CompanyMerge) HasSicCode() bool`

HasSicCode returns a boolean if a field has been set.

### SetSicCodeNil

`func (o *CompanyMerge) SetSicCodeNil(b bool)`

 SetSicCodeNil sets the value for SicCode to be an explicit nil

### UnsetSicCode
`func (o *CompanyMerge) UnsetSicCode()`

UnsetSicCode ensures that no value is present for SicCode, not even an explicit nil
### GetDateAcquired

`func (o *CompanyMerge) GetDateAcquired() string`

GetDateAcquired returns the DateAcquired field if non-nil, zero value otherwise.

### GetDateAcquiredOk

`func (o *CompanyMerge) GetDateAcquiredOk() (*string, bool)`

GetDateAcquiredOk returns a tuple with the DateAcquired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateAcquired

`func (o *CompanyMerge) SetDateAcquired(v string)`

SetDateAcquired sets DateAcquired field to given value.

### HasDateAcquired

`func (o *CompanyMerge) HasDateAcquired() bool`

HasDateAcquired returns a boolean if a field has been set.

### SetDateAcquiredNil

`func (o *CompanyMerge) SetDateAcquiredNil(b bool)`

 SetDateAcquiredNil sets the value for DateAcquired to be an explicit nil

### UnsetDateAcquired
`func (o *CompanyMerge) UnsetDateAcquired()`

UnsetDateAcquired ensures that no value is present for DateAcquired, not even an explicit nil
### GetTimeZone

`func (o *CompanyMerge) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CompanyMerge) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CompanyMerge) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *CompanyMerge) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### SetTimeZoneNil

`func (o *CompanyMerge) SetTimeZoneNil(b bool)`

 SetTimeZoneNil sets the value for TimeZone to be an explicit nil

### UnsetTimeZone
`func (o *CompanyMerge) UnsetTimeZone()`

UnsetTimeZone ensures that no value is present for TimeZone, not even an explicit nil
### GetSourceList

`func (o *CompanyMerge) GetSourceList() string`

GetSourceList returns the SourceList field if non-nil, zero value otherwise.

### GetSourceListOk

`func (o *CompanyMerge) GetSourceListOk() (*string, bool)`

GetSourceListOk returns a tuple with the SourceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceList

`func (o *CompanyMerge) SetSourceList(v string)`

SetSourceList sets SourceList field to given value.

### HasSourceList

`func (o *CompanyMerge) HasSourceList() bool`

HasSourceList returns a boolean if a field has been set.

### SetSourceListNil

`func (o *CompanyMerge) SetSourceListNil(b bool)`

 SetSourceListNil sets the value for SourceList to be an explicit nil

### UnsetSourceList
`func (o *CompanyMerge) UnsetSourceList()`

UnsetSourceList ensures that no value is present for SourceList, not even an explicit nil
### GetUserDefinedField1

`func (o *CompanyMerge) GetUserDefinedField1() string`

GetUserDefinedField1 returns the UserDefinedField1 field if non-nil, zero value otherwise.

### GetUserDefinedField1Ok

`func (o *CompanyMerge) GetUserDefinedField1Ok() (*string, bool)`

GetUserDefinedField1Ok returns a tuple with the UserDefinedField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField1

`func (o *CompanyMerge) SetUserDefinedField1(v string)`

SetUserDefinedField1 sets UserDefinedField1 field to given value.

### HasUserDefinedField1

`func (o *CompanyMerge) HasUserDefinedField1() bool`

HasUserDefinedField1 returns a boolean if a field has been set.

### SetUserDefinedField1Nil

`func (o *CompanyMerge) SetUserDefinedField1Nil(b bool)`

 SetUserDefinedField1Nil sets the value for UserDefinedField1 to be an explicit nil

### UnsetUserDefinedField1
`func (o *CompanyMerge) UnsetUserDefinedField1()`

UnsetUserDefinedField1 ensures that no value is present for UserDefinedField1, not even an explicit nil
### GetUserDefinedField2

`func (o *CompanyMerge) GetUserDefinedField2() string`

GetUserDefinedField2 returns the UserDefinedField2 field if non-nil, zero value otherwise.

### GetUserDefinedField2Ok

`func (o *CompanyMerge) GetUserDefinedField2Ok() (*string, bool)`

GetUserDefinedField2Ok returns a tuple with the UserDefinedField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField2

`func (o *CompanyMerge) SetUserDefinedField2(v string)`

SetUserDefinedField2 sets UserDefinedField2 field to given value.

### HasUserDefinedField2

`func (o *CompanyMerge) HasUserDefinedField2() bool`

HasUserDefinedField2 returns a boolean if a field has been set.

### SetUserDefinedField2Nil

`func (o *CompanyMerge) SetUserDefinedField2Nil(b bool)`

 SetUserDefinedField2Nil sets the value for UserDefinedField2 to be an explicit nil

### UnsetUserDefinedField2
`func (o *CompanyMerge) UnsetUserDefinedField2()`

UnsetUserDefinedField2 ensures that no value is present for UserDefinedField2, not even an explicit nil
### GetUserDefinedField3

`func (o *CompanyMerge) GetUserDefinedField3() string`

GetUserDefinedField3 returns the UserDefinedField3 field if non-nil, zero value otherwise.

### GetUserDefinedField3Ok

`func (o *CompanyMerge) GetUserDefinedField3Ok() (*string, bool)`

GetUserDefinedField3Ok returns a tuple with the UserDefinedField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField3

`func (o *CompanyMerge) SetUserDefinedField3(v string)`

SetUserDefinedField3 sets UserDefinedField3 field to given value.

### HasUserDefinedField3

`func (o *CompanyMerge) HasUserDefinedField3() bool`

HasUserDefinedField3 returns a boolean if a field has been set.

### SetUserDefinedField3Nil

`func (o *CompanyMerge) SetUserDefinedField3Nil(b bool)`

 SetUserDefinedField3Nil sets the value for UserDefinedField3 to be an explicit nil

### UnsetUserDefinedField3
`func (o *CompanyMerge) UnsetUserDefinedField3()`

UnsetUserDefinedField3 ensures that no value is present for UserDefinedField3, not even an explicit nil
### GetUserDefinedField4

`func (o *CompanyMerge) GetUserDefinedField4() string`

GetUserDefinedField4 returns the UserDefinedField4 field if non-nil, zero value otherwise.

### GetUserDefinedField4Ok

`func (o *CompanyMerge) GetUserDefinedField4Ok() (*string, bool)`

GetUserDefinedField4Ok returns a tuple with the UserDefinedField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField4

`func (o *CompanyMerge) SetUserDefinedField4(v string)`

SetUserDefinedField4 sets UserDefinedField4 field to given value.

### HasUserDefinedField4

`func (o *CompanyMerge) HasUserDefinedField4() bool`

HasUserDefinedField4 returns a boolean if a field has been set.

### SetUserDefinedField4Nil

`func (o *CompanyMerge) SetUserDefinedField4Nil(b bool)`

 SetUserDefinedField4Nil sets the value for UserDefinedField4 to be an explicit nil

### UnsetUserDefinedField4
`func (o *CompanyMerge) UnsetUserDefinedField4()`

UnsetUserDefinedField4 ensures that no value is present for UserDefinedField4, not even an explicit nil
### GetUserDefinedField5

`func (o *CompanyMerge) GetUserDefinedField5() string`

GetUserDefinedField5 returns the UserDefinedField5 field if non-nil, zero value otherwise.

### GetUserDefinedField5Ok

`func (o *CompanyMerge) GetUserDefinedField5Ok() (*string, bool)`

GetUserDefinedField5Ok returns a tuple with the UserDefinedField5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField5

`func (o *CompanyMerge) SetUserDefinedField5(v string)`

SetUserDefinedField5 sets UserDefinedField5 field to given value.

### HasUserDefinedField5

`func (o *CompanyMerge) HasUserDefinedField5() bool`

HasUserDefinedField5 returns a boolean if a field has been set.

### SetUserDefinedField5Nil

`func (o *CompanyMerge) SetUserDefinedField5Nil(b bool)`

 SetUserDefinedField5Nil sets the value for UserDefinedField5 to be an explicit nil

### UnsetUserDefinedField5
`func (o *CompanyMerge) UnsetUserDefinedField5()`

UnsetUserDefinedField5 ensures that no value is present for UserDefinedField5, not even an explicit nil
### GetUserDefinedField6

`func (o *CompanyMerge) GetUserDefinedField6() string`

GetUserDefinedField6 returns the UserDefinedField6 field if non-nil, zero value otherwise.

### GetUserDefinedField6Ok

`func (o *CompanyMerge) GetUserDefinedField6Ok() (*string, bool)`

GetUserDefinedField6Ok returns a tuple with the UserDefinedField6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField6

`func (o *CompanyMerge) SetUserDefinedField6(v string)`

SetUserDefinedField6 sets UserDefinedField6 field to given value.

### HasUserDefinedField6

`func (o *CompanyMerge) HasUserDefinedField6() bool`

HasUserDefinedField6 returns a boolean if a field has been set.

### SetUserDefinedField6Nil

`func (o *CompanyMerge) SetUserDefinedField6Nil(b bool)`

 SetUserDefinedField6Nil sets the value for UserDefinedField6 to be an explicit nil

### UnsetUserDefinedField6
`func (o *CompanyMerge) UnsetUserDefinedField6()`

UnsetUserDefinedField6 ensures that no value is present for UserDefinedField6, not even an explicit nil
### GetUserDefinedField7

`func (o *CompanyMerge) GetUserDefinedField7() string`

GetUserDefinedField7 returns the UserDefinedField7 field if non-nil, zero value otherwise.

### GetUserDefinedField7Ok

`func (o *CompanyMerge) GetUserDefinedField7Ok() (*string, bool)`

GetUserDefinedField7Ok returns a tuple with the UserDefinedField7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField7

`func (o *CompanyMerge) SetUserDefinedField7(v string)`

SetUserDefinedField7 sets UserDefinedField7 field to given value.

### HasUserDefinedField7

`func (o *CompanyMerge) HasUserDefinedField7() bool`

HasUserDefinedField7 returns a boolean if a field has been set.

### SetUserDefinedField7Nil

`func (o *CompanyMerge) SetUserDefinedField7Nil(b bool)`

 SetUserDefinedField7Nil sets the value for UserDefinedField7 to be an explicit nil

### UnsetUserDefinedField7
`func (o *CompanyMerge) UnsetUserDefinedField7()`

UnsetUserDefinedField7 ensures that no value is present for UserDefinedField7, not even an explicit nil
### GetUserDefinedField8

`func (o *CompanyMerge) GetUserDefinedField8() string`

GetUserDefinedField8 returns the UserDefinedField8 field if non-nil, zero value otherwise.

### GetUserDefinedField8Ok

`func (o *CompanyMerge) GetUserDefinedField8Ok() (*string, bool)`

GetUserDefinedField8Ok returns a tuple with the UserDefinedField8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField8

`func (o *CompanyMerge) SetUserDefinedField8(v string)`

SetUserDefinedField8 sets UserDefinedField8 field to given value.

### HasUserDefinedField8

`func (o *CompanyMerge) HasUserDefinedField8() bool`

HasUserDefinedField8 returns a boolean if a field has been set.

### SetUserDefinedField8Nil

`func (o *CompanyMerge) SetUserDefinedField8Nil(b bool)`

 SetUserDefinedField8Nil sets the value for UserDefinedField8 to be an explicit nil

### UnsetUserDefinedField8
`func (o *CompanyMerge) UnsetUserDefinedField8()`

UnsetUserDefinedField8 ensures that no value is present for UserDefinedField8, not even an explicit nil
### GetUserDefinedField9

`func (o *CompanyMerge) GetUserDefinedField9() string`

GetUserDefinedField9 returns the UserDefinedField9 field if non-nil, zero value otherwise.

### GetUserDefinedField9Ok

`func (o *CompanyMerge) GetUserDefinedField9Ok() (*string, bool)`

GetUserDefinedField9Ok returns a tuple with the UserDefinedField9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField9

`func (o *CompanyMerge) SetUserDefinedField9(v string)`

SetUserDefinedField9 sets UserDefinedField9 field to given value.

### HasUserDefinedField9

`func (o *CompanyMerge) HasUserDefinedField9() bool`

HasUserDefinedField9 returns a boolean if a field has been set.

### SetUserDefinedField9Nil

`func (o *CompanyMerge) SetUserDefinedField9Nil(b bool)`

 SetUserDefinedField9Nil sets the value for UserDefinedField9 to be an explicit nil

### UnsetUserDefinedField9
`func (o *CompanyMerge) UnsetUserDefinedField9()`

UnsetUserDefinedField9 ensures that no value is present for UserDefinedField9, not even an explicit nil
### GetUserDefinedField10

`func (o *CompanyMerge) GetUserDefinedField10() string`

GetUserDefinedField10 returns the UserDefinedField10 field if non-nil, zero value otherwise.

### GetUserDefinedField10Ok

`func (o *CompanyMerge) GetUserDefinedField10Ok() (*string, bool)`

GetUserDefinedField10Ok returns a tuple with the UserDefinedField10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedField10

`func (o *CompanyMerge) SetUserDefinedField10(v string)`

SetUserDefinedField10 sets UserDefinedField10 field to given value.

### HasUserDefinedField10

`func (o *CompanyMerge) HasUserDefinedField10() bool`

HasUserDefinedField10 returns a boolean if a field has been set.

### SetUserDefinedField10Nil

`func (o *CompanyMerge) SetUserDefinedField10Nil(b bool)`

 SetUserDefinedField10Nil sets the value for UserDefinedField10 to be an explicit nil

### UnsetUserDefinedField10
`func (o *CompanyMerge) UnsetUserDefinedField10()`

UnsetUserDefinedField10 ensures that no value is present for UserDefinedField10, not even an explicit nil
### GetBillingAddress

`func (o *CompanyMerge) GetBillingAddress() string`

GetBillingAddress returns the BillingAddress field if non-nil, zero value otherwise.

### GetBillingAddressOk

`func (o *CompanyMerge) GetBillingAddressOk() (*string, bool)`

GetBillingAddressOk returns a tuple with the BillingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAddress

`func (o *CompanyMerge) SetBillingAddress(v string)`

SetBillingAddress sets BillingAddress field to given value.

### HasBillingAddress

`func (o *CompanyMerge) HasBillingAddress() bool`

HasBillingAddress returns a boolean if a field has been set.

### SetBillingAddressNil

`func (o *CompanyMerge) SetBillingAddressNil(b bool)`

 SetBillingAddressNil sets the value for BillingAddress to be an explicit nil

### UnsetBillingAddress
`func (o *CompanyMerge) UnsetBillingAddress()`

UnsetBillingAddress ensures that no value is present for BillingAddress, not even an explicit nil
### GetBillingContact

`func (o *CompanyMerge) GetBillingContact() string`

GetBillingContact returns the BillingContact field if non-nil, zero value otherwise.

### GetBillingContactOk

`func (o *CompanyMerge) GetBillingContactOk() (*string, bool)`

GetBillingContactOk returns a tuple with the BillingContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingContact

`func (o *CompanyMerge) SetBillingContact(v string)`

SetBillingContact sets BillingContact field to given value.

### HasBillingContact

`func (o *CompanyMerge) HasBillingContact() bool`

HasBillingContact returns a boolean if a field has been set.

### SetBillingContactNil

`func (o *CompanyMerge) SetBillingContactNil(b bool)`

 SetBillingContactNil sets the value for BillingContact to be an explicit nil

### UnsetBillingContact
`func (o *CompanyMerge) UnsetBillingContact()`

UnsetBillingContact ensures that no value is present for BillingContact, not even an explicit nil
### GetTaxCode

`func (o *CompanyMerge) GetTaxCode() string`

GetTaxCode returns the TaxCode field if non-nil, zero value otherwise.

### GetTaxCodeOk

`func (o *CompanyMerge) GetTaxCodeOk() (*string, bool)`

GetTaxCodeOk returns a tuple with the TaxCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCode

`func (o *CompanyMerge) SetTaxCode(v string)`

SetTaxCode sets TaxCode field to given value.

### HasTaxCode

`func (o *CompanyMerge) HasTaxCode() bool`

HasTaxCode returns a boolean if a field has been set.

### SetTaxCodeNil

`func (o *CompanyMerge) SetTaxCodeNil(b bool)`

 SetTaxCodeNil sets the value for TaxCode to be an explicit nil

### UnsetTaxCode
`func (o *CompanyMerge) UnsetTaxCode()`

UnsetTaxCode ensures that no value is present for TaxCode, not even an explicit nil
### GetAccountNumber

`func (o *CompanyMerge) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *CompanyMerge) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *CompanyMerge) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *CompanyMerge) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### SetAccountNumberNil

`func (o *CompanyMerge) SetAccountNumberNil(b bool)`

 SetAccountNumberNil sets the value for AccountNumber to be an explicit nil

### UnsetAccountNumber
`func (o *CompanyMerge) UnsetAccountNumber()`

UnsetAccountNumber ensures that no value is present for AccountNumber, not even an explicit nil
### GetBillingTerms

`func (o *CompanyMerge) GetBillingTerms() string`

GetBillingTerms returns the BillingTerms field if non-nil, zero value otherwise.

### GetBillingTermsOk

`func (o *CompanyMerge) GetBillingTermsOk() (*string, bool)`

GetBillingTermsOk returns a tuple with the BillingTerms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingTerms

`func (o *CompanyMerge) SetBillingTerms(v string)`

SetBillingTerms sets BillingTerms field to given value.

### HasBillingTerms

`func (o *CompanyMerge) HasBillingTerms() bool`

HasBillingTerms returns a boolean if a field has been set.

### SetBillingTermsNil

`func (o *CompanyMerge) SetBillingTermsNil(b bool)`

 SetBillingTermsNil sets the value for BillingTerms to be an explicit nil

### UnsetBillingTerms
`func (o *CompanyMerge) UnsetBillingTerms()`

UnsetBillingTerms ensures that no value is present for BillingTerms, not even an explicit nil
### GetNotes

`func (o *CompanyMerge) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CompanyMerge) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CompanyMerge) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CompanyMerge) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CompanyMerge) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CompanyMerge) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSites

`func (o *CompanyMerge) GetSites() string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CompanyMerge) GetSitesOk() (*string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CompanyMerge) SetSites(v string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CompanyMerge) HasSites() bool`

HasSites returns a boolean if a field has been set.

### SetSitesNil

`func (o *CompanyMerge) SetSitesNil(b bool)`

 SetSitesNil sets the value for Sites to be an explicit nil

### UnsetSites
`func (o *CompanyMerge) UnsetSites()`

UnsetSites ensures that no value is present for Sites, not even an explicit nil
### GetActivities

`func (o *CompanyMerge) GetActivities() string`

GetActivities returns the Activities field if non-nil, zero value otherwise.

### GetActivitiesOk

`func (o *CompanyMerge) GetActivitiesOk() (*string, bool)`

GetActivitiesOk returns a tuple with the Activities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivities

`func (o *CompanyMerge) SetActivities(v string)`

SetActivities sets Activities field to given value.

### HasActivities

`func (o *CompanyMerge) HasActivities() bool`

HasActivities returns a boolean if a field has been set.

### SetActivitiesNil

`func (o *CompanyMerge) SetActivitiesNil(b bool)`

 SetActivitiesNil sets the value for Activities to be an explicit nil

### UnsetActivities
`func (o *CompanyMerge) UnsetActivities()`

UnsetActivities ensures that no value is present for Activities, not even an explicit nil
### GetOpportunities

`func (o *CompanyMerge) GetOpportunities() string`

GetOpportunities returns the Opportunities field if non-nil, zero value otherwise.

### GetOpportunitiesOk

`func (o *CompanyMerge) GetOpportunitiesOk() (*string, bool)`

GetOpportunitiesOk returns a tuple with the Opportunities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpportunities

`func (o *CompanyMerge) SetOpportunities(v string)`

SetOpportunities sets Opportunities field to given value.

### HasOpportunities

`func (o *CompanyMerge) HasOpportunities() bool`

HasOpportunities returns a boolean if a field has been set.

### SetOpportunitiesNil

`func (o *CompanyMerge) SetOpportunitiesNil(b bool)`

 SetOpportunitiesNil sets the value for Opportunities to be an explicit nil

### UnsetOpportunities
`func (o *CompanyMerge) UnsetOpportunities()`

UnsetOpportunities ensures that no value is present for Opportunities, not even an explicit nil
### GetServices

`func (o *CompanyMerge) GetServices() string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *CompanyMerge) GetServicesOk() (*string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *CompanyMerge) SetServices(v string)`

SetServices sets Services field to given value.

### HasServices

`func (o *CompanyMerge) HasServices() bool`

HasServices returns a boolean if a field has been set.

### SetServicesNil

`func (o *CompanyMerge) SetServicesNil(b bool)`

 SetServicesNil sets the value for Services to be an explicit nil

### UnsetServices
`func (o *CompanyMerge) UnsetServices()`

UnsetServices ensures that no value is present for Services, not even an explicit nil
### GetProjects

`func (o *CompanyMerge) GetProjects() string`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *CompanyMerge) GetProjectsOk() (*string, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *CompanyMerge) SetProjects(v string)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *CompanyMerge) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### SetProjectsNil

`func (o *CompanyMerge) SetProjectsNil(b bool)`

 SetProjectsNil sets the value for Projects to be an explicit nil

### UnsetProjects
`func (o *CompanyMerge) UnsetProjects()`

UnsetProjects ensures that no value is present for Projects, not even an explicit nil
### GetContacts

`func (o *CompanyMerge) GetContacts() string`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *CompanyMerge) GetContactsOk() (*string, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *CompanyMerge) SetContacts(v string)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *CompanyMerge) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *CompanyMerge) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *CompanyMerge) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetDocuments

`func (o *CompanyMerge) GetDocuments() string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *CompanyMerge) GetDocumentsOk() (*string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *CompanyMerge) SetDocuments(v string)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *CompanyMerge) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *CompanyMerge) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *CompanyMerge) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


