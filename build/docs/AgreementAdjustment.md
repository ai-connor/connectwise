# AgreementAdjustment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Amount** | Pointer to **NullableFloat64** |  | [optional] 
**Description** | Pointer to **string** |  Max length: 1000; | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewAgreementAdjustment

`func NewAgreementAdjustment() *AgreementAdjustment`

NewAgreementAdjustment instantiates a new AgreementAdjustment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementAdjustmentWithDefaults

`func NewAgreementAdjustmentWithDefaults() *AgreementAdjustment`

NewAgreementAdjustmentWithDefaults instantiates a new AgreementAdjustment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementAdjustment) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementAdjustment) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementAdjustment) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementAdjustment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAmount

`func (o *AgreementAdjustment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *AgreementAdjustment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *AgreementAdjustment) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *AgreementAdjustment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *AgreementAdjustment) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *AgreementAdjustment) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetDescription

`func (o *AgreementAdjustment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AgreementAdjustment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AgreementAdjustment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AgreementAdjustment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEffectiveDate

`func (o *AgreementAdjustment) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *AgreementAdjustment) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *AgreementAdjustment) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *AgreementAdjustment) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetAgreementId

`func (o *AgreementAdjustment) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementAdjustment) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementAdjustment) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementAdjustment) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementAdjustment) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementAdjustment) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetInfo

`func (o *AgreementAdjustment) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementAdjustment) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementAdjustment) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementAdjustment) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *AgreementAdjustment) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AgreementAdjustment) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AgreementAdjustment) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AgreementAdjustment) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


