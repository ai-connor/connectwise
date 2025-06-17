# AgreementBatchSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**NextRunDate** | **time.Time** |  | 
**DaysInAdvance** | **NullableInt32** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementBatchSetup

`func NewAgreementBatchSetup(nextRunDate time.Time, daysInAdvance NullableInt32, ) *AgreementBatchSetup`

NewAgreementBatchSetup instantiates a new AgreementBatchSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementBatchSetupWithDefaults

`func NewAgreementBatchSetupWithDefaults() *AgreementBatchSetup`

NewAgreementBatchSetupWithDefaults instantiates a new AgreementBatchSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementBatchSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementBatchSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementBatchSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementBatchSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNextRunDate

`func (o *AgreementBatchSetup) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *AgreementBatchSetup) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *AgreementBatchSetup) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.


### GetDaysInAdvance

`func (o *AgreementBatchSetup) GetDaysInAdvance() int32`

GetDaysInAdvance returns the DaysInAdvance field if non-nil, zero value otherwise.

### GetDaysInAdvanceOk

`func (o *AgreementBatchSetup) GetDaysInAdvanceOk() (*int32, bool)`

GetDaysInAdvanceOk returns a tuple with the DaysInAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysInAdvance

`func (o *AgreementBatchSetup) SetDaysInAdvance(v int32)`

SetDaysInAdvance sets DaysInAdvance field to given value.


### SetDaysInAdvanceNil

`func (o *AgreementBatchSetup) SetDaysInAdvanceNil(b bool)`

 SetDaysInAdvanceNil sets the value for DaysInAdvance to be an explicit nil

### UnsetDaysInAdvance
`func (o *AgreementBatchSetup) UnsetDaysInAdvance()`

UnsetDaysInAdvance ensures that no value is present for DaysInAdvance, not even an explicit nil
### GetInfo

`func (o *AgreementBatchSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementBatchSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementBatchSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementBatchSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


