# MemberPersona

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**JobRolePercentage** | Pointer to **NullableInt32** |  | [optional] 
**Name** | **string** |  Max length: 20; | 
**PersonaId** | **int32** |  | 
**Member** | Pointer to [**MemberReference**](MemberReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMemberPersona

`func NewMemberPersona(name string, personaId int32, ) *MemberPersona`

NewMemberPersona instantiates a new MemberPersona object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberPersonaWithDefaults

`func NewMemberPersonaWithDefaults() *MemberPersona`

NewMemberPersonaWithDefaults instantiates a new MemberPersona object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MemberPersona) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MemberPersona) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MemberPersona) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *MemberPersona) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobRolePercentage

`func (o *MemberPersona) GetJobRolePercentage() int32`

GetJobRolePercentage returns the JobRolePercentage field if non-nil, zero value otherwise.

### GetJobRolePercentageOk

`func (o *MemberPersona) GetJobRolePercentageOk() (*int32, bool)`

GetJobRolePercentageOk returns a tuple with the JobRolePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobRolePercentage

`func (o *MemberPersona) SetJobRolePercentage(v int32)`

SetJobRolePercentage sets JobRolePercentage field to given value.

### HasJobRolePercentage

`func (o *MemberPersona) HasJobRolePercentage() bool`

HasJobRolePercentage returns a boolean if a field has been set.

### SetJobRolePercentageNil

`func (o *MemberPersona) SetJobRolePercentageNil(b bool)`

 SetJobRolePercentageNil sets the value for JobRolePercentage to be an explicit nil

### UnsetJobRolePercentage
`func (o *MemberPersona) UnsetJobRolePercentage()`

UnsetJobRolePercentage ensures that no value is present for JobRolePercentage, not even an explicit nil
### GetName

`func (o *MemberPersona) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberPersona) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberPersona) SetName(v string)`

SetName sets Name field to given value.


### GetPersonaId

`func (o *MemberPersona) GetPersonaId() int32`

GetPersonaId returns the PersonaId field if non-nil, zero value otherwise.

### GetPersonaIdOk

`func (o *MemberPersona) GetPersonaIdOk() (*int32, bool)`

GetPersonaIdOk returns a tuple with the PersonaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonaId

`func (o *MemberPersona) SetPersonaId(v int32)`

SetPersonaId sets PersonaId field to given value.


### GetMember

`func (o *MemberPersona) GetMember() MemberReference`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *MemberPersona) GetMemberOk() (*MemberReference, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *MemberPersona) SetMember(v MemberReference)`

SetMember sets Member field to given value.

### HasMember

`func (o *MemberPersona) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetInfo

`func (o *MemberPersona) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *MemberPersona) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *MemberPersona) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *MemberPersona) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


