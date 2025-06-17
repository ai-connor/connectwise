# AgreementTypeWorkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**WorkType** | Pointer to [**WorkTypeReference**](WorkTypeReference.md) |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**Rate** | Pointer to **NullableFloat64** |  | [optional] 
**RateType** | **NullableString** |  | 
**BillTime** | **NullableString** |  | 
**HoursMin** | Pointer to **NullableFloat64** |  | [optional] 
**HoursMax** | Pointer to **NullableFloat64** |  | [optional] 
**RoundBillHours** | Pointer to **NullableFloat64** |  | [optional] 
**OverageRate** | Pointer to **NullableFloat64** |  | [optional] 
**OverageRateType** | **NullableString** |  | 
**LimitTo** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeWorkType

`func NewAgreementTypeWorkType(rateType NullableString, billTime NullableString, overageRateType NullableString, ) *AgreementTypeWorkType`

NewAgreementTypeWorkType instantiates a new AgreementTypeWorkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWorkTypeWithDefaults

`func NewAgreementTypeWorkTypeWithDefaults() *AgreementTypeWorkType`

NewAgreementTypeWorkTypeWithDefaults instantiates a new AgreementTypeWorkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeWorkType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeWorkType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeWorkType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeWorkType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeWorkType) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeWorkType) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeWorkType) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeWorkType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkType

`func (o *AgreementTypeWorkType) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *AgreementTypeWorkType) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *AgreementTypeWorkType) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.

### HasWorkType

`func (o *AgreementTypeWorkType) HasWorkType() bool`

HasWorkType returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AgreementTypeWorkType) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AgreementTypeWorkType) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AgreementTypeWorkType) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AgreementTypeWorkType) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEndingDate

`func (o *AgreementTypeWorkType) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *AgreementTypeWorkType) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *AgreementTypeWorkType) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *AgreementTypeWorkType) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetRate

`func (o *AgreementTypeWorkType) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AgreementTypeWorkType) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AgreementTypeWorkType) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AgreementTypeWorkType) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *AgreementTypeWorkType) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *AgreementTypeWorkType) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetRateType

`func (o *AgreementTypeWorkType) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *AgreementTypeWorkType) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *AgreementTypeWorkType) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *AgreementTypeWorkType) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *AgreementTypeWorkType) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetBillTime

`func (o *AgreementTypeWorkType) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *AgreementTypeWorkType) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *AgreementTypeWorkType) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.


### SetBillTimeNil

`func (o *AgreementTypeWorkType) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *AgreementTypeWorkType) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetHoursMin

`func (o *AgreementTypeWorkType) GetHoursMin() float64`

GetHoursMin returns the HoursMin field if non-nil, zero value otherwise.

### GetHoursMinOk

`func (o *AgreementTypeWorkType) GetHoursMinOk() (*float64, bool)`

GetHoursMinOk returns a tuple with the HoursMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMin

`func (o *AgreementTypeWorkType) SetHoursMin(v float64)`

SetHoursMin sets HoursMin field to given value.

### HasHoursMin

`func (o *AgreementTypeWorkType) HasHoursMin() bool`

HasHoursMin returns a boolean if a field has been set.

### SetHoursMinNil

`func (o *AgreementTypeWorkType) SetHoursMinNil(b bool)`

 SetHoursMinNil sets the value for HoursMin to be an explicit nil

### UnsetHoursMin
`func (o *AgreementTypeWorkType) UnsetHoursMin()`

UnsetHoursMin ensures that no value is present for HoursMin, not even an explicit nil
### GetHoursMax

`func (o *AgreementTypeWorkType) GetHoursMax() float64`

GetHoursMax returns the HoursMax field if non-nil, zero value otherwise.

### GetHoursMaxOk

`func (o *AgreementTypeWorkType) GetHoursMaxOk() (*float64, bool)`

GetHoursMaxOk returns a tuple with the HoursMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMax

`func (o *AgreementTypeWorkType) SetHoursMax(v float64)`

SetHoursMax sets HoursMax field to given value.

### HasHoursMax

`func (o *AgreementTypeWorkType) HasHoursMax() bool`

HasHoursMax returns a boolean if a field has been set.

### SetHoursMaxNil

`func (o *AgreementTypeWorkType) SetHoursMaxNil(b bool)`

 SetHoursMaxNil sets the value for HoursMax to be an explicit nil

### UnsetHoursMax
`func (o *AgreementTypeWorkType) UnsetHoursMax()`

UnsetHoursMax ensures that no value is present for HoursMax, not even an explicit nil
### GetRoundBillHours

`func (o *AgreementTypeWorkType) GetRoundBillHours() float64`

GetRoundBillHours returns the RoundBillHours field if non-nil, zero value otherwise.

### GetRoundBillHoursOk

`func (o *AgreementTypeWorkType) GetRoundBillHoursOk() (*float64, bool)`

GetRoundBillHoursOk returns a tuple with the RoundBillHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundBillHours

`func (o *AgreementTypeWorkType) SetRoundBillHours(v float64)`

SetRoundBillHours sets RoundBillHours field to given value.

### HasRoundBillHours

`func (o *AgreementTypeWorkType) HasRoundBillHours() bool`

HasRoundBillHours returns a boolean if a field has been set.

### SetRoundBillHoursNil

`func (o *AgreementTypeWorkType) SetRoundBillHoursNil(b bool)`

 SetRoundBillHoursNil sets the value for RoundBillHours to be an explicit nil

### UnsetRoundBillHours
`func (o *AgreementTypeWorkType) UnsetRoundBillHours()`

UnsetRoundBillHours ensures that no value is present for RoundBillHours, not even an explicit nil
### GetOverageRate

`func (o *AgreementTypeWorkType) GetOverageRate() float64`

GetOverageRate returns the OverageRate field if non-nil, zero value otherwise.

### GetOverageRateOk

`func (o *AgreementTypeWorkType) GetOverageRateOk() (*float64, bool)`

GetOverageRateOk returns a tuple with the OverageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageRate

`func (o *AgreementTypeWorkType) SetOverageRate(v float64)`

SetOverageRate sets OverageRate field to given value.

### HasOverageRate

`func (o *AgreementTypeWorkType) HasOverageRate() bool`

HasOverageRate returns a boolean if a field has been set.

### SetOverageRateNil

`func (o *AgreementTypeWorkType) SetOverageRateNil(b bool)`

 SetOverageRateNil sets the value for OverageRate to be an explicit nil

### UnsetOverageRate
`func (o *AgreementTypeWorkType) UnsetOverageRate()`

UnsetOverageRate ensures that no value is present for OverageRate, not even an explicit nil
### GetOverageRateType

`func (o *AgreementTypeWorkType) GetOverageRateType() string`

GetOverageRateType returns the OverageRateType field if non-nil, zero value otherwise.

### GetOverageRateTypeOk

`func (o *AgreementTypeWorkType) GetOverageRateTypeOk() (*string, bool)`

GetOverageRateTypeOk returns a tuple with the OverageRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverageRateType

`func (o *AgreementTypeWorkType) SetOverageRateType(v string)`

SetOverageRateType sets OverageRateType field to given value.


### SetOverageRateTypeNil

`func (o *AgreementTypeWorkType) SetOverageRateTypeNil(b bool)`

 SetOverageRateTypeNil sets the value for OverageRateType to be an explicit nil

### UnsetOverageRateType
`func (o *AgreementTypeWorkType) UnsetOverageRateType()`

UnsetOverageRateType ensures that no value is present for OverageRateType, not even an explicit nil
### GetLimitTo

`func (o *AgreementTypeWorkType) GetLimitTo() float64`

GetLimitTo returns the LimitTo field if non-nil, zero value otherwise.

### GetLimitToOk

`func (o *AgreementTypeWorkType) GetLimitToOk() (*float64, bool)`

GetLimitToOk returns a tuple with the LimitTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTo

`func (o *AgreementTypeWorkType) SetLimitTo(v float64)`

SetLimitTo sets LimitTo field to given value.

### HasLimitTo

`func (o *AgreementTypeWorkType) HasLimitTo() bool`

HasLimitTo returns a boolean if a field has been set.

### SetLimitToNil

`func (o *AgreementTypeWorkType) SetLimitToNil(b bool)`

 SetLimitToNil sets the value for LimitTo to be an explicit nil

### UnsetLimitTo
`func (o *AgreementTypeWorkType) UnsetLimitTo()`

UnsetLimitTo ensures that no value is present for LimitTo, not even an explicit nil
### GetInfo

`func (o *AgreementTypeWorkType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeWorkType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeWorkType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeWorkType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


