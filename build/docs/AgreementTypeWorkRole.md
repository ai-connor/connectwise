# AgreementTypeWorkRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**EndingDate** | Pointer to **time.Time** |  | [optional] 
**Rate** | Pointer to **NullableFloat64** |  | [optional] 
**RateType** | **NullableString** |  | 
**LimitTo** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeWorkRole

`func NewAgreementTypeWorkRole(rateType NullableString, ) *AgreementTypeWorkRole`

NewAgreementTypeWorkRole instantiates a new AgreementTypeWorkRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWorkRoleWithDefaults

`func NewAgreementTypeWorkRoleWithDefaults() *AgreementTypeWorkRole`

NewAgreementTypeWorkRoleWithDefaults instantiates a new AgreementTypeWorkRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeWorkRole) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeWorkRole) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeWorkRole) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeWorkRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeWorkRole) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeWorkRole) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeWorkRole) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeWorkRole) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkRole

`func (o *AgreementTypeWorkRole) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementTypeWorkRole) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementTypeWorkRole) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *AgreementTypeWorkRole) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AgreementTypeWorkRole) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AgreementTypeWorkRole) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AgreementTypeWorkRole) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AgreementTypeWorkRole) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetEndingDate

`func (o *AgreementTypeWorkRole) GetEndingDate() time.Time`

GetEndingDate returns the EndingDate field if non-nil, zero value otherwise.

### GetEndingDateOk

`func (o *AgreementTypeWorkRole) GetEndingDateOk() (*time.Time, bool)`

GetEndingDateOk returns a tuple with the EndingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingDate

`func (o *AgreementTypeWorkRole) SetEndingDate(v time.Time)`

SetEndingDate sets EndingDate field to given value.

### HasEndingDate

`func (o *AgreementTypeWorkRole) HasEndingDate() bool`

HasEndingDate returns a boolean if a field has been set.

### GetRate

`func (o *AgreementTypeWorkRole) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *AgreementTypeWorkRole) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *AgreementTypeWorkRole) SetRate(v float64)`

SetRate sets Rate field to given value.

### HasRate

`func (o *AgreementTypeWorkRole) HasRate() bool`

HasRate returns a boolean if a field has been set.

### SetRateNil

`func (o *AgreementTypeWorkRole) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *AgreementTypeWorkRole) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetRateType

`func (o *AgreementTypeWorkRole) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *AgreementTypeWorkRole) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *AgreementTypeWorkRole) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *AgreementTypeWorkRole) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *AgreementTypeWorkRole) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetLimitTo

`func (o *AgreementTypeWorkRole) GetLimitTo() float64`

GetLimitTo returns the LimitTo field if non-nil, zero value otherwise.

### GetLimitToOk

`func (o *AgreementTypeWorkRole) GetLimitToOk() (*float64, bool)`

GetLimitToOk returns a tuple with the LimitTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTo

`func (o *AgreementTypeWorkRole) SetLimitTo(v float64)`

SetLimitTo sets LimitTo field to given value.

### HasLimitTo

`func (o *AgreementTypeWorkRole) HasLimitTo() bool`

HasLimitTo returns a boolean if a field has been set.

### SetLimitToNil

`func (o *AgreementTypeWorkRole) SetLimitToNil(b bool)`

 SetLimitToNil sets the value for LimitTo to be an explicit nil

### UnsetLimitTo
`func (o *AgreementTypeWorkRole) UnsetLimitTo()`

UnsetLimitTo ensures that no value is present for LimitTo, not even an explicit nil
### GetInfo

`func (o *AgreementTypeWorkRole) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeWorkRole) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeWorkRole) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeWorkRole) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


