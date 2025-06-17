# Management

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**RunTime** | Pointer to **time.Time** |  | [optional] 
**AddedConfigurationStatus** | [**ConfigurationStatusReference**](ConfigurationStatusReference.md) |  | 
**DeletedConfigurationStatus** | [**ConfigurationStatusReference**](ConfigurationStatusReference.md) |  | 
**IntegratorLogin** | [**IntegratorLoginReference**](IntegratorLoginReference.md) |  | 
**ScheduleExecutiveSummaryReportFlag** | **NullableBool** |  | 
**ExecutiveSummaryReportScheduleDay** | Pointer to **NullableInt32** | Gets or sets             this is only required when scheduleExecutiveSummaryReportFlag &#x3D; true. | [optional] 
**ExecutiveSummaryReportScheduleHour** | Pointer to **NullableInt32** | Gets or sets             this is only required when scheduleExecutiveSummaryReportFlag &#x3D; true. Input should be in 24 hour format, ie 2pm is 14. | [optional] 
**ExecutiveSummaryReportScheduleMinute** | Pointer to **NullableInt32** | Gets or sets             this is only required when scheduleExecutiveSummaryReportFlag &#x3D; true. | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagement

`func NewManagement(addedConfigurationStatus ConfigurationStatusReference, deletedConfigurationStatus ConfigurationStatusReference, integratorLogin IntegratorLoginReference, scheduleExecutiveSummaryReportFlag NullableBool, ) *Management`

NewManagement instantiates a new Management object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementWithDefaults

`func NewManagementWithDefaults() *Management`

NewManagementWithDefaults instantiates a new Management object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Management) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Management) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Management) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Management) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRunTime

`func (o *Management) GetRunTime() time.Time`

GetRunTime returns the RunTime field if non-nil, zero value otherwise.

### GetRunTimeOk

`func (o *Management) GetRunTimeOk() (*time.Time, bool)`

GetRunTimeOk returns a tuple with the RunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunTime

`func (o *Management) SetRunTime(v time.Time)`

SetRunTime sets RunTime field to given value.

### HasRunTime

`func (o *Management) HasRunTime() bool`

HasRunTime returns a boolean if a field has been set.

### GetAddedConfigurationStatus

`func (o *Management) GetAddedConfigurationStatus() ConfigurationStatusReference`

GetAddedConfigurationStatus returns the AddedConfigurationStatus field if non-nil, zero value otherwise.

### GetAddedConfigurationStatusOk

`func (o *Management) GetAddedConfigurationStatusOk() (*ConfigurationStatusReference, bool)`

GetAddedConfigurationStatusOk returns a tuple with the AddedConfigurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedConfigurationStatus

`func (o *Management) SetAddedConfigurationStatus(v ConfigurationStatusReference)`

SetAddedConfigurationStatus sets AddedConfigurationStatus field to given value.


### GetDeletedConfigurationStatus

`func (o *Management) GetDeletedConfigurationStatus() ConfigurationStatusReference`

GetDeletedConfigurationStatus returns the DeletedConfigurationStatus field if non-nil, zero value otherwise.

### GetDeletedConfigurationStatusOk

`func (o *Management) GetDeletedConfigurationStatusOk() (*ConfigurationStatusReference, bool)`

GetDeletedConfigurationStatusOk returns a tuple with the DeletedConfigurationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedConfigurationStatus

`func (o *Management) SetDeletedConfigurationStatus(v ConfigurationStatusReference)`

SetDeletedConfigurationStatus sets DeletedConfigurationStatus field to given value.


### GetIntegratorLogin

`func (o *Management) GetIntegratorLogin() IntegratorLoginReference`

GetIntegratorLogin returns the IntegratorLogin field if non-nil, zero value otherwise.

### GetIntegratorLoginOk

`func (o *Management) GetIntegratorLoginOk() (*IntegratorLoginReference, bool)`

GetIntegratorLoginOk returns a tuple with the IntegratorLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegratorLogin

`func (o *Management) SetIntegratorLogin(v IntegratorLoginReference)`

SetIntegratorLogin sets IntegratorLogin field to given value.


### GetScheduleExecutiveSummaryReportFlag

`func (o *Management) GetScheduleExecutiveSummaryReportFlag() bool`

GetScheduleExecutiveSummaryReportFlag returns the ScheduleExecutiveSummaryReportFlag field if non-nil, zero value otherwise.

### GetScheduleExecutiveSummaryReportFlagOk

`func (o *Management) GetScheduleExecutiveSummaryReportFlagOk() (*bool, bool)`

GetScheduleExecutiveSummaryReportFlagOk returns a tuple with the ScheduleExecutiveSummaryReportFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleExecutiveSummaryReportFlag

`func (o *Management) SetScheduleExecutiveSummaryReportFlag(v bool)`

SetScheduleExecutiveSummaryReportFlag sets ScheduleExecutiveSummaryReportFlag field to given value.


### SetScheduleExecutiveSummaryReportFlagNil

`func (o *Management) SetScheduleExecutiveSummaryReportFlagNil(b bool)`

 SetScheduleExecutiveSummaryReportFlagNil sets the value for ScheduleExecutiveSummaryReportFlag to be an explicit nil

### UnsetScheduleExecutiveSummaryReportFlag
`func (o *Management) UnsetScheduleExecutiveSummaryReportFlag()`

UnsetScheduleExecutiveSummaryReportFlag ensures that no value is present for ScheduleExecutiveSummaryReportFlag, not even an explicit nil
### GetExecutiveSummaryReportScheduleDay

`func (o *Management) GetExecutiveSummaryReportScheduleDay() int32`

GetExecutiveSummaryReportScheduleDay returns the ExecutiveSummaryReportScheduleDay field if non-nil, zero value otherwise.

### GetExecutiveSummaryReportScheduleDayOk

`func (o *Management) GetExecutiveSummaryReportScheduleDayOk() (*int32, bool)`

GetExecutiveSummaryReportScheduleDayOk returns a tuple with the ExecutiveSummaryReportScheduleDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummaryReportScheduleDay

`func (o *Management) SetExecutiveSummaryReportScheduleDay(v int32)`

SetExecutiveSummaryReportScheduleDay sets ExecutiveSummaryReportScheduleDay field to given value.

### HasExecutiveSummaryReportScheduleDay

`func (o *Management) HasExecutiveSummaryReportScheduleDay() bool`

HasExecutiveSummaryReportScheduleDay returns a boolean if a field has been set.

### SetExecutiveSummaryReportScheduleDayNil

`func (o *Management) SetExecutiveSummaryReportScheduleDayNil(b bool)`

 SetExecutiveSummaryReportScheduleDayNil sets the value for ExecutiveSummaryReportScheduleDay to be an explicit nil

### UnsetExecutiveSummaryReportScheduleDay
`func (o *Management) UnsetExecutiveSummaryReportScheduleDay()`

UnsetExecutiveSummaryReportScheduleDay ensures that no value is present for ExecutiveSummaryReportScheduleDay, not even an explicit nil
### GetExecutiveSummaryReportScheduleHour

`func (o *Management) GetExecutiveSummaryReportScheduleHour() int32`

GetExecutiveSummaryReportScheduleHour returns the ExecutiveSummaryReportScheduleHour field if non-nil, zero value otherwise.

### GetExecutiveSummaryReportScheduleHourOk

`func (o *Management) GetExecutiveSummaryReportScheduleHourOk() (*int32, bool)`

GetExecutiveSummaryReportScheduleHourOk returns a tuple with the ExecutiveSummaryReportScheduleHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummaryReportScheduleHour

`func (o *Management) SetExecutiveSummaryReportScheduleHour(v int32)`

SetExecutiveSummaryReportScheduleHour sets ExecutiveSummaryReportScheduleHour field to given value.

### HasExecutiveSummaryReportScheduleHour

`func (o *Management) HasExecutiveSummaryReportScheduleHour() bool`

HasExecutiveSummaryReportScheduleHour returns a boolean if a field has been set.

### SetExecutiveSummaryReportScheduleHourNil

`func (o *Management) SetExecutiveSummaryReportScheduleHourNil(b bool)`

 SetExecutiveSummaryReportScheduleHourNil sets the value for ExecutiveSummaryReportScheduleHour to be an explicit nil

### UnsetExecutiveSummaryReportScheduleHour
`func (o *Management) UnsetExecutiveSummaryReportScheduleHour()`

UnsetExecutiveSummaryReportScheduleHour ensures that no value is present for ExecutiveSummaryReportScheduleHour, not even an explicit nil
### GetExecutiveSummaryReportScheduleMinute

`func (o *Management) GetExecutiveSummaryReportScheduleMinute() int32`

GetExecutiveSummaryReportScheduleMinute returns the ExecutiveSummaryReportScheduleMinute field if non-nil, zero value otherwise.

### GetExecutiveSummaryReportScheduleMinuteOk

`func (o *Management) GetExecutiveSummaryReportScheduleMinuteOk() (*int32, bool)`

GetExecutiveSummaryReportScheduleMinuteOk returns a tuple with the ExecutiveSummaryReportScheduleMinute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutiveSummaryReportScheduleMinute

`func (o *Management) SetExecutiveSummaryReportScheduleMinute(v int32)`

SetExecutiveSummaryReportScheduleMinute sets ExecutiveSummaryReportScheduleMinute field to given value.

### HasExecutiveSummaryReportScheduleMinute

`func (o *Management) HasExecutiveSummaryReportScheduleMinute() bool`

HasExecutiveSummaryReportScheduleMinute returns a boolean if a field has been set.

### SetExecutiveSummaryReportScheduleMinuteNil

`func (o *Management) SetExecutiveSummaryReportScheduleMinuteNil(b bool)`

 SetExecutiveSummaryReportScheduleMinuteNil sets the value for ExecutiveSummaryReportScheduleMinute to be an explicit nil

### UnsetExecutiveSummaryReportScheduleMinute
`func (o *Management) UnsetExecutiveSummaryReportScheduleMinute()`

UnsetExecutiveSummaryReportScheduleMinute ensures that no value is present for ExecutiveSummaryReportScheduleMinute, not even an explicit nil
### GetInfo

`func (o *Management) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Management) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Management) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Management) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


