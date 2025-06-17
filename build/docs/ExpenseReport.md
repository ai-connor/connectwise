# ExpenseReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Period** | Pointer to **NullableInt32** |  | [optional] 
**DateStart** | Pointer to **string** |  | [optional] 
**DateEnd** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Total** | Pointer to **NullableFloat64** |  | [optional] 
**DueDate** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewExpenseReport

`func NewExpenseReport() *ExpenseReport`

NewExpenseReport instantiates a new ExpenseReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExpenseReportWithDefaults

`func NewExpenseReportWithDefaults() *ExpenseReport`

NewExpenseReportWithDefaults instantiates a new ExpenseReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExpenseReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExpenseReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExpenseReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExpenseReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMember

`func (o *ExpenseReport) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *ExpenseReport) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *ExpenseReport) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *ExpenseReport) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetYear

`func (o *ExpenseReport) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *ExpenseReport) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *ExpenseReport) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *ExpenseReport) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *ExpenseReport) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *ExpenseReport) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetPeriod

`func (o *ExpenseReport) GetPeriod() int32`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ExpenseReport) GetPeriodOk() (*int32, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ExpenseReport) SetPeriod(v int32)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *ExpenseReport) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### SetPeriodNil

`func (o *ExpenseReport) SetPeriodNil(b bool)`

 SetPeriodNil sets the value for Period to be an explicit nil

### UnsetPeriod
`func (o *ExpenseReport) UnsetPeriod()`

UnsetPeriod ensures that no value is present for Period, not even an explicit nil
### GetDateStart

`func (o *ExpenseReport) GetDateStart() string`

GetDateStart returns the DateStart field if non-nil, zero value otherwise.

### GetDateStartOk

`func (o *ExpenseReport) GetDateStartOk() (*string, bool)`

GetDateStartOk returns a tuple with the DateStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateStart

`func (o *ExpenseReport) SetDateStart(v string)`

SetDateStart sets DateStart field to given value.

### HasDateStart

`func (o *ExpenseReport) HasDateStart() bool`

HasDateStart returns a boolean if a field has been set.

### GetDateEnd

`func (o *ExpenseReport) GetDateEnd() string`

GetDateEnd returns the DateEnd field if non-nil, zero value otherwise.

### GetDateEndOk

`func (o *ExpenseReport) GetDateEndOk() (*string, bool)`

GetDateEndOk returns a tuple with the DateEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateEnd

`func (o *ExpenseReport) SetDateEnd(v string)`

SetDateEnd sets DateEnd field to given value.

### HasDateEnd

`func (o *ExpenseReport) HasDateEnd() bool`

HasDateEnd returns a boolean if a field has been set.

### GetStatus

`func (o *ExpenseReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExpenseReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExpenseReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExpenseReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ExpenseReport) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ExpenseReport) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetTotal

`func (o *ExpenseReport) GetTotal() float64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExpenseReport) GetTotalOk() (*float64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExpenseReport) SetTotal(v float64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ExpenseReport) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *ExpenseReport) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *ExpenseReport) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil
### GetDueDate

`func (o *ExpenseReport) GetDueDate() string`

GetDueDate returns the DueDate field if non-nil, zero value otherwise.

### GetDueDateOk

`func (o *ExpenseReport) GetDueDateOk() (*string, bool)`

GetDueDateOk returns a tuple with the DueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueDate

`func (o *ExpenseReport) SetDueDate(v string)`

SetDueDate sets DueDate field to given value.

### HasDueDate

`func (o *ExpenseReport) HasDueDate() bool`

HasDueDate returns a boolean if a field has been set.

### GetInfo

`func (o *ExpenseReport) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ExpenseReport) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ExpenseReport) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ExpenseReport) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


