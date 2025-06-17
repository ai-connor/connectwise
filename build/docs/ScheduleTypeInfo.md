# ScheduleTypeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**ChargeCode** | Pointer to [**ChargeCodeReference**](ChargeCodeReference.md) |  | [optional] 
**Where** | Pointer to [**ServiceLocationReference**](ServiceLocationReference.md) |  | [optional] 
**SystemFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewScheduleTypeInfo

`func NewScheduleTypeInfo() *ScheduleTypeInfo`

NewScheduleTypeInfo instantiates a new ScheduleTypeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleTypeInfoWithDefaults

`func NewScheduleTypeInfoWithDefaults() *ScheduleTypeInfo`

NewScheduleTypeInfoWithDefaults instantiates a new ScheduleTypeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScheduleTypeInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScheduleTypeInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScheduleTypeInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScheduleTypeInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScheduleTypeInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduleTypeInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduleTypeInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduleTypeInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIdentifier

`func (o *ScheduleTypeInfo) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ScheduleTypeInfo) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ScheduleTypeInfo) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ScheduleTypeInfo) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetChargeCode

`func (o *ScheduleTypeInfo) GetChargeCode() ChargeCodeReference`

GetChargeCode returns the ChargeCode field if non-nil, zero value otherwise.

### GetChargeCodeOk

`func (o *ScheduleTypeInfo) GetChargeCodeOk() (*ChargeCodeReference, bool)`

GetChargeCodeOk returns a tuple with the ChargeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeCode

`func (o *ScheduleTypeInfo) SetChargeCode(v ChargeCodeReference)`

SetChargeCode sets ChargeCode field to given value.

### HasChargeCode

`func (o *ScheduleTypeInfo) HasChargeCode() bool`

HasChargeCode returns a boolean if a field has been set.

### GetWhere

`func (o *ScheduleTypeInfo) GetWhere() ServiceLocationReference`

GetWhere returns the Where field if non-nil, zero value otherwise.

### GetWhereOk

`func (o *ScheduleTypeInfo) GetWhereOk() (*ServiceLocationReference, bool)`

GetWhereOk returns a tuple with the Where field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhere

`func (o *ScheduleTypeInfo) SetWhere(v ServiceLocationReference)`

SetWhere sets Where field to given value.

### HasWhere

`func (o *ScheduleTypeInfo) HasWhere() bool`

HasWhere returns a boolean if a field has been set.

### GetSystemFlag

`func (o *ScheduleTypeInfo) GetSystemFlag() bool`

GetSystemFlag returns the SystemFlag field if non-nil, zero value otherwise.

### GetSystemFlagOk

`func (o *ScheduleTypeInfo) GetSystemFlagOk() (*bool, bool)`

GetSystemFlagOk returns a tuple with the SystemFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFlag

`func (o *ScheduleTypeInfo) SetSystemFlag(v bool)`

SetSystemFlag sets SystemFlag field to given value.

### HasSystemFlag

`func (o *ScheduleTypeInfo) HasSystemFlag() bool`

HasSystemFlag returns a boolean if a field has been set.

### SetSystemFlagNil

`func (o *ScheduleTypeInfo) SetSystemFlagNil(b bool)`

 SetSystemFlagNil sets the value for SystemFlag to be an explicit nil

### UnsetSystemFlag
`func (o *ScheduleTypeInfo) UnsetSystemFlag()`

UnsetSystemFlag ensures that no value is present for SystemFlag, not even an explicit nil
### GetInfo

`func (o *ScheduleTypeInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ScheduleTypeInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ScheduleTypeInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ScheduleTypeInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


