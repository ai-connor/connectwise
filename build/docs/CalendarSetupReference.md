# CalendarSetupReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCalendarSetupReference

`func NewCalendarSetupReference() *CalendarSetupReference`

NewCalendarSetupReference instantiates a new CalendarSetupReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalendarSetupReferenceWithDefaults

`func NewCalendarSetupReferenceWithDefaults() *CalendarSetupReference`

NewCalendarSetupReferenceWithDefaults instantiates a new CalendarSetupReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CalendarSetupReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CalendarSetupReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CalendarSetupReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CalendarSetupReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CalendarSetupReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CalendarSetupReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetInfo

`func (o *CalendarSetupReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CalendarSetupReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CalendarSetupReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CalendarSetupReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


