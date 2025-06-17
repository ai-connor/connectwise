# TimePeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**TimePeriodSetup** | Pointer to [**TimePeriodSetupReference**](TimePeriodSetupReference.md) |  | [optional] 
**Period** | Pointer to **NullableInt32** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**DeadlineDate** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimePeriod

`func NewTimePeriod() *TimePeriod`

NewTimePeriod instantiates a new TimePeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimePeriodWithDefaults

`func NewTimePeriodWithDefaults() *TimePeriod`

NewTimePeriodWithDefaults instantiates a new TimePeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimePeriod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimePeriod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimePeriod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimePeriod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimePeriodSetup

`func (o *TimePeriod) GetTimePeriodSetup() TimePeriodSetupReference`

GetTimePeriodSetup returns the TimePeriodSetup field if non-nil, zero value otherwise.

### GetTimePeriodSetupOk

`func (o *TimePeriod) GetTimePeriodSetupOk() (*TimePeriodSetupReference, bool)`

GetTimePeriodSetupOk returns a tuple with the TimePeriodSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodSetup

`func (o *TimePeriod) SetTimePeriodSetup(v TimePeriodSetupReference)`

SetTimePeriodSetup sets TimePeriodSetup field to given value.

### HasTimePeriodSetup

`func (o *TimePeriod) HasTimePeriodSetup() bool`

HasTimePeriodSetup returns a boolean if a field has been set.

### GetPeriod

`func (o *TimePeriod) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *TimePeriod) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *TimePeriod) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *TimePeriod) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *TimePeriod) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *TimePeriod) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetStartDate

`func (o *TimePeriod) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *TimePeriod) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *TimePeriod) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *TimePeriod) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *TimePeriod) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *TimePeriod) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *TimePeriod) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *TimePeriod) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDeadlineDate

`func (o *TimePeriod) GetDeadlineDate() string`

GetDeadlineDate returns the DeadlineDate field if non-nil, zero value otherwise.

### GetDeadlineDateOk

`func (o *TimePeriod) GetDeadlineDateOk() (*string, bool)`

GetDeadlineDateOk returns a tuple with the DeadlineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadlineDate

`func (o *TimePeriod) SetDeadlineDate(v string)`

SetDeadlineDate sets DeadlineDate field to given value.

### HasDeadlineDate

`func (o *TimePeriod) HasDeadlineDate() bool`

HasDeadlineDate returns a boolean if a field has been set.

### GetInfo

`func (o *TimePeriod) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimePeriod) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimePeriod) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimePeriod) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


