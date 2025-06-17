# WorkType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**BillTime** | **NullableString** |  | 
**RateType** | **NullableString** |  | 
**Rate** | **NullableFloat64** |  | 
**HoursMin** | Pointer to **NullableFloat64** |  | [optional] 
**HoursMax** | Pointer to **NullableFloat64** |  | [optional] 
**RoundBillHoursTo** | Pointer to **NullableFloat64** |  | [optional] 
**AccrualType** | Pointer to **NullableString** |  | [optional] 
**InactiveFlag** | Pointer to **NullableBool** |  | [optional] 
**OverallDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**ActivityDefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**UtilizationFlag** | Pointer to **NullableBool** |  | [optional] 
**CostMultiplier** | Pointer to **NullableFloat64** |  | [optional] 
**IntegrationXRef** | Pointer to **string** |  Max length: 50; | [optional] 
**AddAllAgreementExclusions** | Pointer to **NullableBool** | Used only on create to add the work type to all agreement and agreement type exclusion lists | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewWorkType

`func NewWorkType(name string, billTime NullableString, rateType NullableString, rate NullableFloat64, ) *WorkType`

NewWorkType instantiates a new WorkType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkTypeWithDefaults

`func NewWorkTypeWithDefaults() *WorkType`

NewWorkTypeWithDefaults instantiates a new WorkType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *WorkType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WorkType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkType) SetName(v string)`

SetName sets Name field to given value.


### GetBillTime

`func (o *WorkType) GetBillTime() string`

GetBillTime returns the BillTime field if non-nil, zero value otherwise.

### GetBillTimeOk

`func (o *WorkType) GetBillTimeOk() (*string, bool)`

GetBillTimeOk returns a tuple with the BillTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillTime

`func (o *WorkType) SetBillTime(v string)`

SetBillTime sets BillTime field to given value.


### SetBillTimeNil

`func (o *WorkType) SetBillTimeNil(b bool)`

 SetBillTimeNil sets the value for BillTime to be an explicit nil

### UnsetBillTime
`func (o *WorkType) UnsetBillTime()`

UnsetBillTime ensures that no value is present for BillTime, not even an explicit nil
### GetRateType

`func (o *WorkType) GetRateType() string`

GetRateType returns the RateType field if non-nil, zero value otherwise.

### GetRateTypeOk

`func (o *WorkType) GetRateTypeOk() (*string, bool)`

GetRateTypeOk returns a tuple with the RateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateType

`func (o *WorkType) SetRateType(v string)`

SetRateType sets RateType field to given value.


### SetRateTypeNil

`func (o *WorkType) SetRateTypeNil(b bool)`

 SetRateTypeNil sets the value for RateType to be an explicit nil

### UnsetRateType
`func (o *WorkType) UnsetRateType()`

UnsetRateType ensures that no value is present for RateType, not even an explicit nil
### GetRate

`func (o *WorkType) GetRate() float64`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *WorkType) GetRateOk() (*float64, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *WorkType) SetRate(v float64)`

SetRate sets Rate field to given value.


### SetRateNil

`func (o *WorkType) SetRateNil(b bool)`

 SetRateNil sets the value for Rate to be an explicit nil

### UnsetRate
`func (o *WorkType) UnsetRate()`

UnsetRate ensures that no value is present for Rate, not even an explicit nil
### GetHoursMin

`func (o *WorkType) GetHoursMin() float64`

GetHoursMin returns the HoursMin field if non-nil, zero value otherwise.

### GetHoursMinOk

`func (o *WorkType) GetHoursMinOk() (*float64, bool)`

GetHoursMinOk returns a tuple with the HoursMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMin

`func (o *WorkType) SetHoursMin(v float64)`

SetHoursMin sets HoursMin field to given value.

### HasHoursMin

`func (o *WorkType) HasHoursMin() bool`

HasHoursMin returns a boolean if a field has been set.

### SetHoursMinNil

`func (o *WorkType) SetHoursMinNil(b bool)`

 SetHoursMinNil sets the value for HoursMin to be an explicit nil

### UnsetHoursMin
`func (o *WorkType) UnsetHoursMin()`

UnsetHoursMin ensures that no value is present for HoursMin, not even an explicit nil
### GetHoursMax

`func (o *WorkType) GetHoursMax() float64`

GetHoursMax returns the HoursMax field if non-nil, zero value otherwise.

### GetHoursMaxOk

`func (o *WorkType) GetHoursMaxOk() (*float64, bool)`

GetHoursMaxOk returns a tuple with the HoursMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursMax

`func (o *WorkType) SetHoursMax(v float64)`

SetHoursMax sets HoursMax field to given value.

### HasHoursMax

`func (o *WorkType) HasHoursMax() bool`

HasHoursMax returns a boolean if a field has been set.

### SetHoursMaxNil

`func (o *WorkType) SetHoursMaxNil(b bool)`

 SetHoursMaxNil sets the value for HoursMax to be an explicit nil

### UnsetHoursMax
`func (o *WorkType) UnsetHoursMax()`

UnsetHoursMax ensures that no value is present for HoursMax, not even an explicit nil
### GetRoundBillHoursTo

`func (o *WorkType) GetRoundBillHoursTo() float64`

GetRoundBillHoursTo returns the RoundBillHoursTo field if non-nil, zero value otherwise.

### GetRoundBillHoursToOk

`func (o *WorkType) GetRoundBillHoursToOk() (*float64, bool)`

GetRoundBillHoursToOk returns a tuple with the RoundBillHoursTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundBillHoursTo

`func (o *WorkType) SetRoundBillHoursTo(v float64)`

SetRoundBillHoursTo sets RoundBillHoursTo field to given value.

### HasRoundBillHoursTo

`func (o *WorkType) HasRoundBillHoursTo() bool`

HasRoundBillHoursTo returns a boolean if a field has been set.

### SetRoundBillHoursToNil

`func (o *WorkType) SetRoundBillHoursToNil(b bool)`

 SetRoundBillHoursToNil sets the value for RoundBillHoursTo to be an explicit nil

### UnsetRoundBillHoursTo
`func (o *WorkType) UnsetRoundBillHoursTo()`

UnsetRoundBillHoursTo ensures that no value is present for RoundBillHoursTo, not even an explicit nil
### GetAccrualType

`func (o *WorkType) GetAccrualType() string`

GetAccrualType returns the AccrualType field if non-nil, zero value otherwise.

### GetAccrualTypeOk

`func (o *WorkType) GetAccrualTypeOk() (*string, bool)`

GetAccrualTypeOk returns a tuple with the AccrualType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualType

`func (o *WorkType) SetAccrualType(v string)`

SetAccrualType sets AccrualType field to given value.

### HasAccrualType

`func (o *WorkType) HasAccrualType() bool`

HasAccrualType returns a boolean if a field has been set.

### SetAccrualTypeNil

`func (o *WorkType) SetAccrualTypeNil(b bool)`

 SetAccrualTypeNil sets the value for AccrualType to be an explicit nil

### UnsetAccrualType
`func (o *WorkType) UnsetAccrualType()`

UnsetAccrualType ensures that no value is present for AccrualType, not even an explicit nil
### GetInactiveFlag

`func (o *WorkType) GetInactiveFlag() bool`

GetInactiveFlag returns the InactiveFlag field if non-nil, zero value otherwise.

### GetInactiveFlagOk

`func (o *WorkType) GetInactiveFlagOk() (*bool, bool)`

GetInactiveFlagOk returns a tuple with the InactiveFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactiveFlag

`func (o *WorkType) SetInactiveFlag(v bool)`

SetInactiveFlag sets InactiveFlag field to given value.

### HasInactiveFlag

`func (o *WorkType) HasInactiveFlag() bool`

HasInactiveFlag returns a boolean if a field has been set.

### SetInactiveFlagNil

`func (o *WorkType) SetInactiveFlagNil(b bool)`

 SetInactiveFlagNil sets the value for InactiveFlag to be an explicit nil

### UnsetInactiveFlag
`func (o *WorkType) UnsetInactiveFlag()`

UnsetInactiveFlag ensures that no value is present for InactiveFlag, not even an explicit nil
### GetOverallDefaultFlag

`func (o *WorkType) GetOverallDefaultFlag() bool`

GetOverallDefaultFlag returns the OverallDefaultFlag field if non-nil, zero value otherwise.

### GetOverallDefaultFlagOk

`func (o *WorkType) GetOverallDefaultFlagOk() (*bool, bool)`

GetOverallDefaultFlagOk returns a tuple with the OverallDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverallDefaultFlag

`func (o *WorkType) SetOverallDefaultFlag(v bool)`

SetOverallDefaultFlag sets OverallDefaultFlag field to given value.

### HasOverallDefaultFlag

`func (o *WorkType) HasOverallDefaultFlag() bool`

HasOverallDefaultFlag returns a boolean if a field has been set.

### SetOverallDefaultFlagNil

`func (o *WorkType) SetOverallDefaultFlagNil(b bool)`

 SetOverallDefaultFlagNil sets the value for OverallDefaultFlag to be an explicit nil

### UnsetOverallDefaultFlag
`func (o *WorkType) UnsetOverallDefaultFlag()`

UnsetOverallDefaultFlag ensures that no value is present for OverallDefaultFlag, not even an explicit nil
### GetActivityDefaultFlag

`func (o *WorkType) GetActivityDefaultFlag() bool`

GetActivityDefaultFlag returns the ActivityDefaultFlag field if non-nil, zero value otherwise.

### GetActivityDefaultFlagOk

`func (o *WorkType) GetActivityDefaultFlagOk() (*bool, bool)`

GetActivityDefaultFlagOk returns a tuple with the ActivityDefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityDefaultFlag

`func (o *WorkType) SetActivityDefaultFlag(v bool)`

SetActivityDefaultFlag sets ActivityDefaultFlag field to given value.

### HasActivityDefaultFlag

`func (o *WorkType) HasActivityDefaultFlag() bool`

HasActivityDefaultFlag returns a boolean if a field has been set.

### SetActivityDefaultFlagNil

`func (o *WorkType) SetActivityDefaultFlagNil(b bool)`

 SetActivityDefaultFlagNil sets the value for ActivityDefaultFlag to be an explicit nil

### UnsetActivityDefaultFlag
`func (o *WorkType) UnsetActivityDefaultFlag()`

UnsetActivityDefaultFlag ensures that no value is present for ActivityDefaultFlag, not even an explicit nil
### GetUtilizationFlag

`func (o *WorkType) GetUtilizationFlag() bool`

GetUtilizationFlag returns the UtilizationFlag field if non-nil, zero value otherwise.

### GetUtilizationFlagOk

`func (o *WorkType) GetUtilizationFlagOk() (*bool, bool)`

GetUtilizationFlagOk returns a tuple with the UtilizationFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationFlag

`func (o *WorkType) SetUtilizationFlag(v bool)`

SetUtilizationFlag sets UtilizationFlag field to given value.

### HasUtilizationFlag

`func (o *WorkType) HasUtilizationFlag() bool`

HasUtilizationFlag returns a boolean if a field has been set.

### SetUtilizationFlagNil

`func (o *WorkType) SetUtilizationFlagNil(b bool)`

 SetUtilizationFlagNil sets the value for UtilizationFlag to be an explicit nil

### UnsetUtilizationFlag
`func (o *WorkType) UnsetUtilizationFlag()`

UnsetUtilizationFlag ensures that no value is present for UtilizationFlag, not even an explicit nil
### GetCostMultiplier

`func (o *WorkType) GetCostMultiplier() float64`

GetCostMultiplier returns the CostMultiplier field if non-nil, zero value otherwise.

### GetCostMultiplierOk

`func (o *WorkType) GetCostMultiplierOk() (*float64, bool)`

GetCostMultiplierOk returns a tuple with the CostMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostMultiplier

`func (o *WorkType) SetCostMultiplier(v float64)`

SetCostMultiplier sets CostMultiplier field to given value.

### HasCostMultiplier

`func (o *WorkType) HasCostMultiplier() bool`

HasCostMultiplier returns a boolean if a field has been set.

### SetCostMultiplierNil

`func (o *WorkType) SetCostMultiplierNil(b bool)`

 SetCostMultiplierNil sets the value for CostMultiplier to be an explicit nil

### UnsetCostMultiplier
`func (o *WorkType) UnsetCostMultiplier()`

UnsetCostMultiplier ensures that no value is present for CostMultiplier, not even an explicit nil
### GetIntegrationXRef

`func (o *WorkType) GetIntegrationXRef() string`

GetIntegrationXRef returns the IntegrationXRef field if non-nil, zero value otherwise.

### GetIntegrationXRefOk

`func (o *WorkType) GetIntegrationXRefOk() (*string, bool)`

GetIntegrationXRefOk returns a tuple with the IntegrationXRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationXRef

`func (o *WorkType) SetIntegrationXRef(v string)`

SetIntegrationXRef sets IntegrationXRef field to given value.

### HasIntegrationXRef

`func (o *WorkType) HasIntegrationXRef() bool`

HasIntegrationXRef returns a boolean if a field has been set.

### GetAddAllAgreementExclusions

`func (o *WorkType) GetAddAllAgreementExclusions() bool`

GetAddAllAgreementExclusions returns the AddAllAgreementExclusions field if non-nil, zero value otherwise.

### GetAddAllAgreementExclusionsOk

`func (o *WorkType) GetAddAllAgreementExclusionsOk() (*bool, bool)`

GetAddAllAgreementExclusionsOk returns a tuple with the AddAllAgreementExclusions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddAllAgreementExclusions

`func (o *WorkType) SetAddAllAgreementExclusions(v bool)`

SetAddAllAgreementExclusions sets AddAllAgreementExclusions field to given value.

### HasAddAllAgreementExclusions

`func (o *WorkType) HasAddAllAgreementExclusions() bool`

HasAddAllAgreementExclusions returns a boolean if a field has been set.

### SetAddAllAgreementExclusionsNil

`func (o *WorkType) SetAddAllAgreementExclusionsNil(b bool)`

 SetAddAllAgreementExclusionsNil sets the value for AddAllAgreementExclusions to be an explicit nil

### UnsetAddAllAgreementExclusions
`func (o *WorkType) UnsetAddAllAgreementExclusions()`

UnsetAddAllAgreementExclusions ensures that no value is present for AddAllAgreementExclusions, not even an explicit nil
### GetInfo

`func (o *WorkType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *WorkType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *WorkType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *WorkType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


