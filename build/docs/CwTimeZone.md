# CwTimeZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Offset** | Pointer to **float64** | The hours offset (+/-). | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**DaylightSavingsFlag** | Pointer to **NullableBool** | Determined based on system library value for specified timeZone.             Not able to be used in query params at this time. | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCwTimeZone

`func NewCwTimeZone() *CwTimeZone`

NewCwTimeZone instantiates a new CwTimeZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCwTimeZoneWithDefaults

`func NewCwTimeZoneWithDefaults() *CwTimeZone`

NewCwTimeZoneWithDefaults instantiates a new CwTimeZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CwTimeZone) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CwTimeZone) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CwTimeZone) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CwTimeZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CwTimeZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CwTimeZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CwTimeZone) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CwTimeZone) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOffset

`func (o *CwTimeZone) GetOffset() float64`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *CwTimeZone) GetOffsetOk() (*float64, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *CwTimeZone) SetOffset(v float64)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *CwTimeZone) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetStartDate

`func (o *CwTimeZone) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *CwTimeZone) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *CwTimeZone) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *CwTimeZone) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *CwTimeZone) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CwTimeZone) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CwTimeZone) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CwTimeZone) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDaylightSavingsFlag

`func (o *CwTimeZone) GetDaylightSavingsFlag() bool`

GetDaylightSavingsFlag returns the DaylightSavingsFlag field if non-nil, zero value otherwise.

### GetDaylightSavingsFlagOk

`func (o *CwTimeZone) GetDaylightSavingsFlagOk() (*bool, bool)`

GetDaylightSavingsFlagOk returns a tuple with the DaylightSavingsFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaylightSavingsFlag

`func (o *CwTimeZone) SetDaylightSavingsFlag(v bool)`

SetDaylightSavingsFlag sets DaylightSavingsFlag field to given value.

### HasDaylightSavingsFlag

`func (o *CwTimeZone) HasDaylightSavingsFlag() bool`

HasDaylightSavingsFlag returns a boolean if a field has been set.

### SetDaylightSavingsFlagNil

`func (o *CwTimeZone) SetDaylightSavingsFlagNil(b bool)`

 SetDaylightSavingsFlagNil sets the value for DaylightSavingsFlag to be an explicit nil

### UnsetDaylightSavingsFlag
`func (o *CwTimeZone) UnsetDaylightSavingsFlag()`

UnsetDaylightSavingsFlag ensures that no value is present for DaylightSavingsFlag, not even an explicit nil
### GetInfo

`func (o *CwTimeZone) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CwTimeZone) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CwTimeZone) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CwTimeZone) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


