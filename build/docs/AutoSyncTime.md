# AutoSyncTime

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SyncTime** | **string** |  | 
**TimeZone** | [**TimeZoneSetupReference**](TimeZoneSetupReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAutoSyncTime

`func NewAutoSyncTime(syncTime string, timeZone TimeZoneSetupReference, ) *AutoSyncTime`

NewAutoSyncTime instantiates a new AutoSyncTime object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSyncTimeWithDefaults

`func NewAutoSyncTimeWithDefaults() *AutoSyncTime`

NewAutoSyncTimeWithDefaults instantiates a new AutoSyncTime object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AutoSyncTime) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutoSyncTime) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutoSyncTime) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AutoSyncTime) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSyncTime

`func (o *AutoSyncTime) GetSyncTime() string`

GetSyncTime returns the SyncTime field if non-nil, zero value otherwise.

### GetSyncTimeOk

`func (o *AutoSyncTime) GetSyncTimeOk() (*string, bool)`

GetSyncTimeOk returns a tuple with the SyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTime

`func (o *AutoSyncTime) SetSyncTime(v string)`

SetSyncTime sets SyncTime field to given value.


### GetTimeZone

`func (o *AutoSyncTime) GetTimeZone() TimeZoneSetupReference`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *AutoSyncTime) GetTimeZoneOk() (*TimeZoneSetupReference, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *AutoSyncTime) SetTimeZone(v TimeZoneSetupReference)`

SetTimeZone sets TimeZone field to given value.


### GetInfo

`func (o *AutoSyncTime) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AutoSyncTime) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AutoSyncTime) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AutoSyncTime) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


