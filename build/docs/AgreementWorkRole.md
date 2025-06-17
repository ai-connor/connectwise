# AgreementWorkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**LocationId** | Pointer to **NullableInt32** |  | [optional] 
**Location** | Pointer to [**OwnerLevelReference**](OwnerLevelReference.md) |  | [optional] 
**RateType** | **NullableString** |  | 
**Rate** | Pointer to **NullableFloat64** |  | [optional] 
**LimitTo** | Pointer to **NullableFloat64** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Agreement** | Pointer to [**AgreementReference**](AgreementReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementWorkRole

`func NewAgreementWorkRole(rateType NullableString, ) *AgreementWorkRole`

NewAgreementWorkRole instantiates a new AgreementWorkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWorkRoleWithDefaults

`func NewAgreementWorkRoleWithDefaults() *AgreementWorkRole`

NewAgreementWorkRoleWithDefaults instantiates a new AgreementWorkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementWorkRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementWorkRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementWorkRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementWorkRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkRole

`func (o *AgreementWorkRole) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementWorkRole) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementWorkRole) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *AgreementWorkRole) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetLocationId

`func (o *AgreementWorkRole) GetLocationId() int32`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *AgreementWorkRole) GetLocationIdOk() (*int32, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *AgreementWorkRole) SetLocationId(v int32)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *AgreementWorkRole) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### SetLocationIdNil

`func (o *AgreementWorkRole) SetLocationIdNil(b bool)`

 SetLocationIdNil sets the value for LocationId to be an explicit nil

### UnsetLocationId
`func (o *AgreementWorkRole) UnsetLocationId()`

UnsetLocationId ensures that no value is present for LocationId, not even an explicit nil
### GetLocation

`func (o *AgreementWorkRole) GetLocation() OwnerLevelReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AgreementWorkRole) GetLocationOk() (*OwnerLevelReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AgreementWorkRole) SetLocation(v OwnerLevelReference)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AgreementWorkRole) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRateType

`func (o *AgreementWorkRole) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *AgreementWorkRole) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *AgreementWorkRole) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *AgreementWorkRole) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *AgreementWorkRole) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetRate

`func (o *AgreementWorkRole) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AgreementWorkRole) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AgreementWorkRole) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AgreementWorkRole) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *AgreementWorkRole) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *AgreementWorkRole) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetLimitTo

`func (o *AgreementWorkRole) GetLimitTo() float64`

GetLimitTo returns the LimitTo field if non-nil, zero value otherwise.

### GetLimitToOk

`func (o *AgreementWorkRole) GetLimitToOk() (*float64, bool)`

GetLimitToOk returns a tuple with the LimitTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTo

`func (o *AgreementWorkRole) SetLimitTo(v float64)`

SetLimitTo sets LimitTo field to given value.

### HasLimitTo

`func (o *AgreementWorkRole) HasLimitTo() bool`

HasLimitTo returns a boolean if a field has been set.

### SetLimitToNil

`func (o *AgreementWorkRole) SetLimitToNil(b bool)`

 SetLimitToNil sets the value for LimitTo to be an explicit nil

### UnsetLimitTo
`func (o *AgreementWorkRole) UnsetLimitTo()`

UnsetLimitTo ensures that no value is present for LimitTo, not even an explicit nil
### GetEffectiveDate

`func (o *AgreementWorkRole) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AgreementWorkRole) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AgreementWorkRole) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AgreementWorkRole) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEndingDate

`func (o *AgreementWorkRole) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *AgreementWorkRole) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *AgreementWorkRole) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *AgreementWorkRole) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetAgreementId

`func (o *AgreementWorkRole) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementWorkRole) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementWorkRole) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementWorkRole) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementWorkRole) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementWorkRole) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetAgreement

`func (o *AgreementWorkRole) GetAgreement() AgreementReference`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *AgreementWorkRole) GetAgreementOk() (*AgreementReference, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *AgreementWorkRole) SetAgreement(v AgreementReference)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *AgreementWorkRole) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.

### GetInfo

`func (o *AgreementWorkRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementWorkRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementWorkRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementWorkRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


