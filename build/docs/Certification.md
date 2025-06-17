# Certification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  Max length: 50; | 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCertification

`func NewCertification(name string, company CompanyReference, ) *Certification`

NewCertification instantiates a new Certification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificationWithDefaults

`func NewCertificationWithDefaults() *Certification`

NewCertificationWithDefaults instantiates a new Certification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Certification) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certification) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certification) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Certification) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Certification) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Certification) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Certification) SetName(v string)`

SetName sets Name field to given value.


### GetCompany

`func (o *Certification) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Certification) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Certification) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetInfo

`func (o *Certification) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *Certification) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *Certification) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *Certification) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


