# ScheduleType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Identifier** | **string** |  Max length: 1; | 
**ChargeCode** | Pointer to [**ChargeCodeReference**](ChargeCodeReference.md) |  | [optional] 
**Where** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**SystemFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleType

`func NewScheduleType(name string, identifier string, ) *ScheduleType`

NewScheduleType instantiates a new ScheduleType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTypeWithDefaults

`func NewScheduleTypeWithDefaults() *ScheduleType`

NewScheduleTypeWithDefaults instantiates a new ScheduleType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScheduleType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleType) SetName(v string)`

SetName sets Name field to given value.


### GetIdentifier

`func (o *ScheduleType) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ScheduleType) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ScheduleType) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetChargeCode

`func (o *ScheduleType) GetChargeCode() ChargeCodeReference`

GetChargeCode returns the ChargeCode field if non-nil, zero value otherwise.

### GetChargeCodeOk

`func (o *ScheduleType) GetChargeCodeOk() (*ChargeCodeReference, bool)`

GetChargeCodeOk returns a tuple with the ChargeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCode

`func (o *ScheduleType) SetChargeCode(v ChargeCodeReference)`

SetChargeCode sets ChargeCode field to given value.

### HasChargeCode

`func (o *ScheduleType) HasChargeCode() bool`

HasChargeCode returns a boolean if a field has been set.

### GetWhere

`func (o *ScheduleType) GetWhere() ServiceLocationReference`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *ScheduleType) GetWhereOk() (*ServiceLocationReference, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *ScheduleType) SetWhere(v ServiceLocationReference)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *ScheduleType) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetSystemFlag

`func (o *ScheduleType) GetSystemFlag() bool`

GetSystemFlag returns the SystemFlag field if non-nil, zero value otherwise.

### GetSystemFlagOk

`func (o *ScheduleType) GetSystemFlagOk() (*bool, bool)`

GetSystemFlagOk returns a tuple with the SystemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFlag

`func (o *ScheduleType) SetSystemFlag(v bool)`

SetSystemFlag sets SystemFlag field to given value.

### HasSystemFlag

`func (o *ScheduleType) HasSystemFlag() bool`

HasSystemFlag returns a boolean if a field has been set.

### SetSystemFlagNil

`func (o *ScheduleType) SetSystemFlagNil(b bool)`

 SetSystemFlagNil sets the value for SystemFlag to be an explicit nil

### UnsetSystemFlag
`func (o *ScheduleType) UnsetSystemFlag()`

UnsetSystemFlag ensures that no value is present for SystemFlag, not even an explicit nil
### GetInfo

`func (o *ScheduleType) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleType) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleType) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleType) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


