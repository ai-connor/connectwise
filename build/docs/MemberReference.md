# MemberReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**DailyCapacity** | Pointer to **NullableFloat64** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberReference

`func NewMemberReference() *MemberReference`

NewMemberReference instantiates a new MemberReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberReferenceWithDefaults

`func NewMemberReferenceWithDefaults() *MemberReference`

NewMemberReferenceWithDefaults instantiates a new MemberReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberReference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberReference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberReference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MemberReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MemberReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *MemberReference) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *MemberReference) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *MemberReference) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *MemberReference) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetName

`func (o *MemberReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberReference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDailyCapacity

`func (o *MemberReference) GetDailyCapacity() float64`

GetDailyCapacity returns the DailyCapacity field if non-nil, zero value otherwise.

### GetDailyCapacityOk

`func (o *MemberReference) GetDailyCapacityOk() (*float64, bool)`

GetDailyCapacityOk returns a tuple with the DailyCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyCapacity

`func (o *MemberReference) SetDailyCapacity(v float64)`

SetDailyCapacity sets DailyCapacity field to given value.

### HasDailyCapacity

`func (o *MemberReference) HasDailyCapacity() bool`

HasDailyCapacity returns a boolean if a field has been set.

### SetDailyCapacityNil

`func (o *MemberReference) SetDailyCapacityNil(b bool)`

 SetDailyCapacityNil sets the value for DailyCapacity to be an explicit nil

### UnsetDailyCapacity
`func (o *MemberReference) UnsetDailyCapacity()`

UnsetDailyCapacity ensures that no value is present for DailyCapacity, not even an explicit nil
### GetInfo

`func (o *MemberReference) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberReference) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberReference) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberReference) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


