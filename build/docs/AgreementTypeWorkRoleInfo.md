# AgreementTypeWorkRoleInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**WorkRole** | Pointer to [**WorkRoleReference**](WorkRoleReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeWorkRoleInfo

`func NewAgreementTypeWorkRoleInfo() *AgreementTypeWorkRoleInfo`

NewAgreementTypeWorkRoleInfo instantiates a new AgreementTypeWorkRoleInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeWorkRoleInfoWithDefaults

`func NewAgreementTypeWorkRoleInfoWithDefaults() *AgreementTypeWorkRoleInfo`

NewAgreementTypeWorkRoleInfoWithDefaults instantiates a new AgreementTypeWorkRoleInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeWorkRoleInfo) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeWorkRoleInfo) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeWorkRoleInfo) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeWorkRoleInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeWorkRoleInfo) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeWorkRoleInfo) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeWorkRoleInfo) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeWorkRoleInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWorkRole

`func (o *AgreementTypeWorkRoleInfo) GetWorkRole() WorkRoleReference`

GetWorkRole returns the WorkRole field if non-nil, zero value otherwise.

### GetWorkRoleOk

`func (o *AgreementTypeWorkRoleInfo) GetWorkRoleOk() (*WorkRoleReference, bool)`

GetWorkRoleOk returns a tuple with the WorkRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkRole

`func (o *AgreementTypeWorkRoleInfo) SetWorkRole(v WorkRoleReference)`

SetWorkRole sets WorkRole field to given value.

### HasWorkRole

`func (o *AgreementTypeWorkRoleInfo) HasWorkRole() bool`

HasWorkRole returns a boolean if a field has been set.

### GetInfo

`func (o *AgreementTypeWorkRoleInfo) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeWorkRoleInfo) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeWorkRoleInfo) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeWorkRoleInfo) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


