# AgreementSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Company** | [**CompanyReference**](CompanyReference.md) |  | 
**Site** | Pointer to [**SiteReference**](SiteReference.md) |  | [optional] 
**AgreementId** | Pointer to **NullableInt32** |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewAgreementSite

`func NewAgreementSite(company CompanyReference, ) *AgreementSite`

NewAgreementSite instantiates a new AgreementSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgreementSiteWithDefaults

`func NewAgreementSiteWithDefaults() *AgreementSite`

NewAgreementSiteWithDefaults instantiates a new AgreementSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgreementSite) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgreementSite) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgreementSite) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgreementSite) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompany

`func (o *AgreementSite) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *AgreementSite) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *AgreementSite) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.


### GetSite

`func (o *AgreementSite) GetSite() SiteReference`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AgreementSite) GetSiteOk() (*SiteReference, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AgreementSite) SetSite(v SiteReference)`

SetSite sets Site field to given value.

### HasSite

`func (o *AgreementSite) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetAgreementId

`func (o *AgreementSite) GetAgreementId() int32`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *AgreementSite) GetAgreementIdOk() (*int32, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *AgreementSite) SetAgreementId(v int32)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *AgreementSite) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### SetAgreementIdNil

`func (o *AgreementSite) SetAgreementIdNil(b bool)`

 SetAgreementIdNil sets the value for AgreementId to be an explicit nil

### UnsetAgreementId
`func (o *AgreementSite) UnsetAgreementId()`

UnsetAgreementId ensures that no value is present for AgreementId, not even an explicit nil
### GetInfo

`func (o *AgreementSite) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *AgreementSite) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *AgreementSite) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *AgreementSite) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *AgreementSite) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *AgreementSite) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *AgreementSite) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *AgreementSite) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


