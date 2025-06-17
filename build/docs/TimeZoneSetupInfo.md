# TimeZoneSetupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewTimeZoneSetupInfo

`func NewTimeZoneSetupInfo() *TimeZoneSetupInfo`

NewTimeZoneSetupInfo instantiates a new TimeZoneSetupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeZoneSetupInfoWithDefaults

`func NewTimeZoneSetupInfoWithDefaults() *TimeZoneSetupInfo`

NewTimeZoneSetupInfoWithDefaults instantiates a new TimeZoneSetupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimeZoneSetupInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimeZoneSetupInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimeZoneSetupInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TimeZoneSetupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TimeZoneSetupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimeZoneSetupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimeZoneSetupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TimeZoneSetupInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *TimeZoneSetupInfo) GetOffset() float64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TimeZoneSetupInfo) GetOffsetOk() (*float64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TimeZoneSetupInfo) SetOffset(v float64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TimeZoneSetupInfo) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### SetOffsetNil

`func (o *TimeZoneSetupInfo) SetOffsetNil(b bool)`

 SetOffsetNil sets the value for Offset to be an explicit nil

### UnsetOffset
`func (o *TimeZoneSetupInfo) UnsetOffset()`

UnsetOffset ensures that no value is present for Offset, not even an explicit nil
### GetInfo

`func (o *TimeZoneSetupInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *TimeZoneSetupInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *TimeZoneSetupInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *TimeZoneSetupInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


