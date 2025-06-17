# AgreementWorkTypeExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**WorkType** | [**WorkTypeReference**](WorkTypeReference.md) |  | 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementWorkTypeExclusion

`func NewAgreementWorkTypeExclusion(workType WorkTypeReference, ) *AgreementWorkTypeExclusion`

NewAgreementWorkTypeExclusion instantiates a new AgreementWorkTypeExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementWorkTypeExclusionWithDefaults

`func NewAgreementWorkTypeExclusionWithDefaults() *AgreementWorkTypeExclusion`

NewAgreementWorkTypeExclusionWithDefaults instantiates a new AgreementWorkTypeExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementWorkTypeExclusion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementWorkTypeExclusion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementWorkTypeExclusion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementWorkTypeExclusion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetWorkType

`func (o *AgreementWorkTypeExclusion) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *AgreementWorkTypeExclusion) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *AgreementWorkTypeExclusion) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.


### GetAgreementId

`func (o *AgreementWorkTypeExclusion) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementWorkTypeExclusion) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementWorkTypeExclusion) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementWorkTypeExclusion) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementWorkTypeExclusion) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementWorkTypeExclusion) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetInfo

`func (o *AgreementWorkTypeExclusion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementWorkTypeExclusion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementWorkTypeExclusion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementWorkTypeExclusion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


