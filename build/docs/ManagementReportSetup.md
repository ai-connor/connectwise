# ManagementReportSetup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ScheduledReportDisabledFlag** | **bool** |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewManagementReportSetup

`func NewManagementReportSetup(scheduledReportDisabledFlag bool, ) *ManagementReportSetup`

NewManagementReportSetup instantiates a new ManagementReportSetup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagementReportSetupWithDefaults

`func NewManagementReportSetupWithDefaults() *ManagementReportSetup`

NewManagementReportSetupWithDefaults instantiates a new ManagementReportSetup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ManagementReportSetup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagementReportSetup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagementReportSetup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ManagementReportSetup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduledReportDisabledFlag

`func (o *ManagementReportSetup) GetScheduledReportDisabledFlag() bool`

GetScheduledReportDisabledFlag returns the ScheduledReportDisabledFlag field if non-nil, zero value otherwise.

### GetScheduledReportDisabledFlagOk

`func (o *ManagementReportSetup) GetScheduledReportDisabledFlagOk() (*bool, bool)`

GetScheduledReportDisabledFlagOk returns a tuple with the ScheduledReportDisabledFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledReportDisabledFlag

`func (o *ManagementReportSetup) SetScheduledReportDisabledFlag(v bool)`

SetScheduledReportDisabledFlag sets ScheduledReportDisabledFlag field to given value.


### GetInfo

`func (o *ManagementReportSetup) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ManagementReportSetup) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ManagementReportSetup) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ManagementReportSetup) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


