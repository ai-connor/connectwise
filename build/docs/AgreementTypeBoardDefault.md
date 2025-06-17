# AgreementTypeBoardDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to [**AgreementTypeReference**](AgreementTypeReference.md) |  | [optional] 
**Location** | [**SystemLocationReference**](SystemLocationReference.md) |  | 
**Department** | Pointer to [**SystemDepartmentReference**](SystemDepartmentReference.md) |  | [optional] 
**Board** | Pointer to [**BoardReference**](BoardReference.md) |  | [optional] 
**ServiceType** | Pointer to [**ServiceTypeReference**](ServiceTypeReference.md) |  | [optional] 
**DefaultFlag** | Pointer to **NullableBool** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewAgreementTypeBoardDefault

`func NewAgreementTypeBoardDefault(location SystemLocationReference, ) *AgreementTypeBoardDefault`

NewAgreementTypeBoardDefault instantiates a new AgreementTypeBoardDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementTypeBoardDefaultWithDefaults

`func NewAgreementTypeBoardDefaultWithDefaults() *AgreementTypeBoardDefault`

NewAgreementTypeBoardDefaultWithDefaults instantiates a new AgreementTypeBoardDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementTypeBoardDefault) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementTypeBoardDefault) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementTypeBoardDefault) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementTypeBoardDefault) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AgreementTypeBoardDefault) GetType() AgreementTypeReference`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgreementTypeBoardDefault) GetTypeOk() (*AgreementTypeReference, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgreementTypeBoardDefault) SetType(v AgreementTypeReference)`

SetType sets Type field to given value.

### HasType

`func (o *AgreementTypeBoardDefault) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLocation

`func (o *AgreementTypeBoardDefault) GetLocation() SystemLocationReference`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AgreementTypeBoardDefault) GetLocationOk() (*SystemLocationReference, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AgreementTypeBoardDefault) SetLocation(v SystemLocationReference)`

SetLocation sets Location field to given value.


### GetDepartment

`func (o *AgreementTypeBoardDefault) GetDepartment() SystemDepartmentReference`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *AgreementTypeBoardDefault) GetDepartmentOk() (*SystemDepartmentReference, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *AgreementTypeBoardDefault) SetDepartment(v SystemDepartmentReference)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *AgreementTypeBoardDefault) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetBoard

`func (o *AgreementTypeBoardDefault) GetBoard() BoardReference`

GetBoard returns the Board field if non-nil, zero value otherwise.

### GetBoardOk

`func (o *AgreementTypeBoardDefault) GetBoardOk() (*BoardReference, bool)`

GetBoardOk returns a tuple with the Board field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoard

`func (o *AgreementTypeBoardDefault) SetBoard(v BoardReference)`

SetBoard sets Board field to given value.

### HasBoard

`func (o *AgreementTypeBoardDefault) HasBoard() bool`

HasBoard returns a boolean if a field has been set.

### GetServiceType

`func (o *AgreementTypeBoardDefault) GetServiceType() ServiceTypeReference`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *AgreementTypeBoardDefault) GetServiceTypeOk() (*ServiceTypeReference, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *AgreementTypeBoardDefault) SetServiceType(v ServiceTypeReference)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *AgreementTypeBoardDefault) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetDefaultFlag

`func (o *AgreementTypeBoardDefault) GetDefaultFlag() bool`

GetDefaultFlag returns the DefaultFlag field if non-nil, zero value otherwise.

### GetDefaultFlagOk

`func (o *AgreementTypeBoardDefault) GetDefaultFlagOk() (*bool, bool)`

GetDefaultFlagOk returns a tuple with the DefaultFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFlag

`func (o *AgreementTypeBoardDefault) SetDefaultFlag(v bool)`

SetDefaultFlag sets DefaultFlag field to given value.

### HasDefaultFlag

`func (o *AgreementTypeBoardDefault) HasDefaultFlag() bool`

HasDefaultFlag returns a boolean if a field has been set.

### SetDefaultFlagNil

`func (o *AgreementTypeBoardDefault) SetDefaultFlagNil(b bool)`

 SetDefaultFlagNil sets the value for DefaultFlag to be an explicit nil

### UnsetDefaultFlag
`func (o *AgreementTypeBoardDefault) UnsetDefaultFlag()`

UnsetDefaultFlag ensures that no value is present for DefaultFlag, not even an explicit nil
### GetInfo

`func (o *AgreementTypeBoardDefault) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementTypeBoardDefault) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementTypeBoardDefault) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementTypeBoardDefault) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


