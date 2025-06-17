# AgreementWorkRoleExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WorkRole** | [**WorkRoleReference**](WorkRoleReference.md) |  | 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementWorkRoleExclusion

`func NewAgreementWorkRoleExclusion(workRole WorkRoleReference, ) *AgreementWorkRoleExclusion`

NewAgreementWorkRoleExclusion instantiates a new AgreementWorkRoleExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWorkRoleExclusionWithDefaults

`func NewAgreementWorkRoleExclusionWithDefaults() *AgreementWorkRoleExclusion`

NewAgreementWorkRoleExclusionWithDefaults instantiates a new AgreementWorkRoleExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementWorkRoleExclusion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementWorkRoleExclusion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementWorkRoleExclusion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementWorkRoleExclusion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkRole

`func (o *AgreementWorkRoleExclusion) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementWorkRoleExclusion) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementWorkRoleExclusion) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.


### GetAgreementId

`func (o *AgreementWorkRoleExclusion) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementWorkRoleExclusion) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementWorkRoleExclusion) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementWorkRoleExclusion) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementWorkRoleExclusion) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementWorkRoleExclusion) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetInfo

`func (o *AgreementWorkRoleExclusion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementWorkRoleExclusion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementWorkRoleExclusion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementWorkRoleExclusion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


