# AgreementWorkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**Location** | Pointer to [**OwnerLevelReference**](OwnerLevelReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**RateType** | **NullableString** |  | 
**BillTime** | **NullableString** |  | 
**Rate** | Pointer to **NullableFloat64** |  | [optional] 
**HoursMax** | Pointer to **NullableFloat64** |  | [optional] 
**HoursMin** | Pointer to **NullableFloat64** |  | [optional] 
**RoundBillHours** | Pointer to **NullableFloat64** |  | [optional] 
**OverageRate** | Pointer to **NullableFloat64** |  | [optional] 
**OverageRateType** | Pointer to **NullableString** |  | [optional] 
**AgreementLimit** | Pointer to **NullableFloat64** |  | [optional] 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementWorkType

`func NewAgreementWorkType(rateType NullableString, billTime NullableString, ) *AgreementWorkType`

NewAgreementWorkType instantiates a new AgreementWorkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWorkTypeWithDefaults

`func NewAgreementWorkTypeWithDefaults() *AgreementWorkType`

NewAgreementWorkTypeWithDefaults instantiates a new AgreementWorkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementWorkType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementWorkType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementWorkType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementWorkType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkType

`func (o *AgreementWorkType) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *AgreementWorkType) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *AgreementWorkType) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *AgreementWorkType) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetLocation

`func (o *AgreementWorkType) GetLocation() OwnerLevelReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AgreementWorkType) GetLocationOk() (*OwnerLevelReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AgreementWorkType) SetLocation(v OwnerLevelReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AgreementWorkType) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLocationId

`func (o *AgreementWorkType) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *AgreementWorkType) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *AgreementWorkType) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *AgreementWorkType) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *AgreementWorkType) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *AgreementWorkType) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetRateType

`func (o *AgreementWorkType) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *AgreementWorkType) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *AgreementWorkType) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *AgreementWorkType) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *AgreementWorkType) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetBillTime

`func (o *AgreementWorkType) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *AgreementWorkType) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *AgreementWorkType) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.


### SetBillTimeNil

`func (o *AgreementWorkType) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *AgreementWorkType) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetRate

`func (o *AgreementWorkType) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AgreementWorkType) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AgreementWorkType) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AgreementWorkType) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *AgreementWorkType) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *AgreementWorkType) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetHoursMax

`func (o *AgreementWorkType) GetHoursMax() float64`

GetHoursMax returns the HoursMax field if non-nil, zero value otherwise.

### GetHoursMaxOk

`func (o *AgreementWorkType) GetHoursMaxOk() (*float64, bool)`

GetHoursMaxOk returns a tuple with the HoursMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMax

`func (o *AgreementWorkType) SetHoursMax(v float64)`

SetHoursMax sets HoursMax field to given value.

### HasHoursMax

`func (o *AgreementWorkType) HasHoursMax() bool`

HasHoursMax returns a boolean if a field has been set.

### SetHoursMaxNil

`func (o *AgreementWorkType) SetHoursMaxNil(b bool)`

 SetHoursMaxNil sets the value for HoursMax to be an explicit nil

### UnsetHoursMax
`func (o *AgreementWorkType) UnsetHoursMax()`

UnsetHoursMax ensures that no value is present for HoursMax, not even an explicit nil
### GetHoursMin

`func (o *AgreementWorkType) GetHoursMin() float64`

GetHoursMin returns the HoursMin field if non-nil, zero value otherwise.

### GetHoursMinOk

`func (o *AgreementWorkType) GetHoursMinOk() (*float64, bool)`

GetHoursMinOk returns a tuple with the HoursMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMin

`func (o *AgreementWorkType) SetHoursMin(v float64)`

SetHoursMin sets HoursMin field to given value.

### HasHoursMin

`func (o *AgreementWorkType) HasHoursMin() bool`

HasHoursMin returns a boolean if a field has been set.

### SetHoursMinNil

`func (o *AgreementWorkType) SetHoursMinNil(b bool)`

 SetHoursMinNil sets the value for HoursMin to be an explicit nil

### UnsetHoursMin
`func (o *AgreementWorkType) UnsetHoursMin()`

UnsetHoursMin ensures that no value is present for HoursMin, not even an explicit nil
### GetRoundBillHours

`func (o *AgreementWorkType) GetRoundBillHours() float64`

GetRoundBillHours returns the RoundBillHours field if non-nil, zero value otherwise.

### GetRoundBillHoursOk

`func (o *AgreementWorkType) GetRoundBillHoursOk() (*float64, bool)`

GetRoundBillHoursOk returns a tuple with the RoundBillHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundBillHours

`func (o *AgreementWorkType) SetRoundBillHours(v float64)`

SetRoundBillHours sets RoundBillHours field to given value.

### HasRoundBillHours

`func (o *AgreementWorkType) HasRoundBillHours() bool`

HasRoundBillHours returns a boolean if a field has been set.

### SetRoundBillHoursNil

`func (o *AgreementWorkType) SetRoundBillHoursNil(b bool)`

 SetRoundBillHoursNil sets the value for RoundBillHours to be an explicit nil

### UnsetRoundBillHours
`func (o *AgreementWorkType) UnsetRoundBillHours()`

UnsetRoundBillHours ensures that no value is present for RoundBillHours, not even an explicit nil
### GetOverageRate

`func (o *AgreementWorkType) GetOverageRate() float64`

GetOverageRate returns the OverageRate field if non-nil, zero value otherwise.

### GetOverageRateOk

`func (o *AgreementWorkType) GetOverageRateOk() (*float64, bool)`

GetOverageRateOk returns a tuple with the OverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageRate

`func (o *AgreementWorkType) SetOverageRate(v float64)`

SetOverageRate sets OverageRate field to given value.

### HasOverageRate

`func (o *AgreementWorkType) HasOverageRate() bool`

HasOverageRate returns a boolean if a field has been set.

### SetOverageRateNil

`func (o *AgreementWorkType) SetOverageRateNil(b bool)`

 SetOverageRateNil sets the value for OverageRate to be an explicit nil

### UnsetOverageRate
`func (o *AgreementWorkType) UnsetOverageRate()`

UnsetOverageRate ensures that no value is present for OverageRate, not even an explicit nil
### GetOverageRateType

`func (o *AgreementWorkType) GetOverageRateType() string`

GetOverageRateType returns the OverageRateType field if non-nil, zero value otherwise.

### GetOverageRateTypeOk

`func (o *AgreementWorkType) GetOverageRateTypeOk() (*string, bool)`

GetOverageRateTypeOk returns a tuple with the OverageRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageRateType

`func (o *AgreementWorkType) SetOverageRateType(v string)`

SetOverageRateType sets OverageRateType field to given value.

### HasOverageRateType

`func (o *AgreementWorkType) HasOverageRateType() bool`

HasOverageRateType returns a boolean if a field has been set.

### SetOverageRateTypeNil

`func (o *AgreementWorkType) SetOverageRateTypeNil(b bool)`

 SetOverageRateTypeNil sets the value for OverageRateType to be an explicit nil

### UnsetOverageRateType
`func (o *AgreementWorkType) UnsetOverageRateType()`

UnsetOverageRateType ensures that no value is present for OverageRateType, not even an explicit nil
### GetAgreementLimit

`func (o *AgreementWorkType) GetAgreementLimit() float64`

GetAgreementLimit returns the AgreementLimit field if non-nil, zero value otherwise.

### GetAgreementLimitOk

`func (o *AgreementWorkType) GetAgreementLimitOk() (*float64, bool)`

GetAgreementLimitOk returns a tuple with the AgreementLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementLimit

`func (o *AgreementWorkType) SetAgreementLimit(v float64)`

SetAgreementLimit sets AgreementLimit field to given value.

### HasAgreementLimit

`func (o *AgreementWorkType) HasAgreementLimit() bool`

HasAgreementLimit returns a boolean if a field has been set.

### SetAgreementLimitNil

`func (o *AgreementWorkType) SetAgreementLimitNil(b bool)`

 SetAgreementLimitNil sets the value for AgreementLimit to be an explicit nil

### UnsetAgreementLimit
`func (o *AgreementWorkType) UnsetAgreementLimit()`

UnsetAgreementLimit ensures that no value is present for AgreementLimit, not even an explicit nil
### GetSite

`func (o *AgreementWorkType) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AgreementWorkType) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AgreementWorkType) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *AgreementWorkType) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AgreementWorkType) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AgreementWorkType) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AgreementWorkType) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AgreementWorkType) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEndingDate

`func (o *AgreementWorkType) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *AgreementWorkType) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *AgreementWorkType) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *AgreementWorkType) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetAgreement

`func (o *AgreementWorkType) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *AgreementWorkType) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *AgreementWorkType) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *AgreementWorkType) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetAgreementId

`func (o *AgreementWorkType) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementWorkType) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementWorkType) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementWorkType) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementWorkType) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementWorkType) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetCompany

`func (o *AgreementWorkType) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AgreementWorkType) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AgreementWorkType) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *AgreementWorkType) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *AgreementWorkType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementWorkType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementWorkType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementWorkType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


