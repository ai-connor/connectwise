# AgreementTypeWorkRoleExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**WorkRole** | [**WorkRoleReference**](WorkRoleReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeWorkRoleExclusion

`func NewAgreementTypeWorkRoleExclusion(workRole WorkRoleReference, ) *AgreementTypeWorkRoleExclusion`

NewAgreementTypeWorkRoleExclusion instantiates a new AgreementTypeWorkRoleExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWorkRoleExclusionWithDefaults

`func NewAgreementTypeWorkRoleExclusionWithDefaults() *AgreementTypeWorkRoleExclusion`

NewAgreementTypeWorkRoleExclusionWithDefaults instantiates a new AgreementTypeWorkRoleExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeWorkRoleExclusion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeWorkRoleExclusion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeWorkRoleExclusion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeWorkRoleExclusion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeWorkRoleExclusion) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeWorkRoleExclusion) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeWorkRoleExclusion) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeWorkRoleExclusion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkRole

`func (o *AgreementTypeWorkRoleExclusion) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementTypeWorkRoleExclusion) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementTypeWorkRoleExclusion) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.


### GetInfo

`func (o *AgreementTypeWorkRoleExclusion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeWorkRoleExclusion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeWorkRoleExclusion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeWorkRoleExclusion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


