# AgreementTypeWorkTypeExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**WorkType** | [**WorkTypeReference**](WorkTypeReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeWorkTypeExclusion

`func NewAgreementTypeWorkTypeExclusion(workType WorkTypeReference, ) *AgreementTypeWorkTypeExclusion`

NewAgreementTypeWorkTypeExclusion instantiates a new AgreementTypeWorkTypeExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWorkTypeExclusionWithDefaults

`func NewAgreementTypeWorkTypeExclusionWithDefaults() *AgreementTypeWorkTypeExclusion`

NewAgreementTypeWorkTypeExclusionWithDefaults instantiates a new AgreementTypeWorkTypeExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeWorkTypeExclusion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeWorkTypeExclusion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeWorkTypeExclusion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeWorkTypeExclusion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeWorkTypeExclusion) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeWorkTypeExclusion) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeWorkTypeExclusion) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeWorkTypeExclusion) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkType

`func (o *AgreementTypeWorkTypeExclusion) GetWorkType() WorkTypeReference`

GetWorkType returns the WorkType field if non-nil, zero value otherwise.

### GetWorkTypeOk

`func (o *AgreementTypeWorkTypeExclusion) GetWorkTypeOk() (*WorkTypeReference, bool)`

GetWorkTypeOk returns a tuple with the WorkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkType

`func (o *AgreementTypeWorkTypeExclusion) SetWorkType(v WorkTypeReference)`

SetWorkType sets WorkType field to given value.


### GetInfo

`func (o *AgreementTypeWorkTypeExclusion) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeWorkTypeExclusion) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeWorkTypeExclusion) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeWorkTypeExclusion) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


