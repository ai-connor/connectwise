# CompanyFinance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**BillOverrideFlag** | Pointer to **NullableBool** |  | [optional] 
**BillSrFlag** | Pointer to **NullableBool** |  | [optional] 
**BillCompleteSrFlag** | Pointer to **NullableBool** |  | [optional] 
**BillUnapprovedSrFlag** | Pointer to **NullableBool** |  | [optional] 
**BillRestrictPmFlag** | Pointer to **NullableBool** |  | [optional] 
**BillCompletePmFlag** | Pointer to **NullableBool** |  | [optional] 
**BillUnapprovedPmFlag** | Pointer to **NullableBool** |  | [optional] 
**EmailTemplate** | Pointer to [**EmailTemplateReference**](EmailTemplateReference.md) |  | [optional] 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 
**CustomFields** | Pointer to [**[]CustomFieldValue**](CustomFieldValue.md) |  | [optional] 

## Methods

### NewCompanyFinance

`func NewCompanyFinance() *CompanyFinance`

NewCompanyFinance instantiates a new CompanyFinance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyFinanceWithDefaults

`func NewCompanyFinanceWithDefaults() *CompanyFinance`

NewCompanyFinanceWithDefaults instantiates a new CompanyFinance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyFinance) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyFinance) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyFinance) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyFinance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetBillOverrideFlag

`func (o *CompanyFinance) GetBillOverrideFlag() bool`

GetBillOverrideFlag returns the BillOverrideFlag field if non-nil, zero value otherwise.

### GetBillOverrideFlagOk

`func (o *CompanyFinance) GetBillOverrideFlagOk() (*bool, bool)`

GetBillOverrideFlagOk returns a tuple with the BillOverrideFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOverrideFlag

`func (o *CompanyFinance) SetBillOverrideFlag(v bool)`

SetBillOverrideFlag sets BillOverrideFlag field to given value.

### HasBillOverrideFlag

`func (o *CompanyFinance) HasBillOverrideFlag() bool`

HasBillOverrideFlag returns a boolean if a field has been set.

### SetBillOverrideFlagNil

`func (o *CompanyFinance) SetBillOverrideFlagNil(b bool)`

 SetBillOverrideFlagNil sets the value for BillOverrideFlag to be an explicit nil

### UnsetBillOverrideFlag
`func (o *CompanyFinance) UnsetBillOverrideFlag()`

UnsetBillOverrideFlag ensures that no value is present for BillOverrideFlag, not even an explicit nil
### GetBillSrFlag

`func (o *CompanyFinance) GetBillSrFlag() bool`

GetBillSrFlag returns the BillSrFlag field if non-nil, zero value otherwise.

### GetBillSrFlagOk

`func (o *CompanyFinance) GetBillSrFlagOk() (*bool, bool)`

GetBillSrFlagOk returns a tuple with the BillSrFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillSrFlag

`func (o *CompanyFinance) SetBillSrFlag(v bool)`

SetBillSrFlag sets BillSrFlag field to given value.

### HasBillSrFlag

`func (o *CompanyFinance) HasBillSrFlag() bool`

HasBillSrFlag returns a boolean if a field has been set.

### SetBillSrFlagNil

`func (o *CompanyFinance) SetBillSrFlagNil(b bool)`

 SetBillSrFlagNil sets the value for BillSrFlag to be an explicit nil

### UnsetBillSrFlag
`func (o *CompanyFinance) UnsetBillSrFlag()`

UnsetBillSrFlag ensures that no value is present for BillSrFlag, not even an explicit nil
### GetBillCompleteSrFlag

`func (o *CompanyFinance) GetBillCompleteSrFlag() bool`

GetBillCompleteSrFlag returns the BillCompleteSrFlag field if non-nil, zero value otherwise.

### GetBillCompleteSrFlagOk

`func (o *CompanyFinance) GetBillCompleteSrFlagOk() (*bool, bool)`

GetBillCompleteSrFlagOk returns a tuple with the BillCompleteSrFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCompleteSrFlag

`func (o *CompanyFinance) SetBillCompleteSrFlag(v bool)`

SetBillCompleteSrFlag sets BillCompleteSrFlag field to given value.

### HasBillCompleteSrFlag

`func (o *CompanyFinance) HasBillCompleteSrFlag() bool`

HasBillCompleteSrFlag returns a boolean if a field has been set.

### SetBillCompleteSrFlagNil

`func (o *CompanyFinance) SetBillCompleteSrFlagNil(b bool)`

 SetBillCompleteSrFlagNil sets the value for BillCompleteSrFlag to be an explicit nil

### UnsetBillCompleteSrFlag
`func (o *CompanyFinance) UnsetBillCompleteSrFlag()`

UnsetBillCompleteSrFlag ensures that no value is present for BillCompleteSrFlag, not even an explicit nil
### GetBillUnapprovedSrFlag

`func (o *CompanyFinance) GetBillUnapprovedSrFlag() bool`

GetBillUnapprovedSrFlag returns the BillUnapprovedSrFlag field if non-nil, zero value otherwise.

### GetBillUnapprovedSrFlagOk

`func (o *CompanyFinance) GetBillUnapprovedSrFlagOk() (*bool, bool)`

GetBillUnapprovedSrFlagOk returns a tuple with the BillUnapprovedSrFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedSrFlag

`func (o *CompanyFinance) SetBillUnapprovedSrFlag(v bool)`

SetBillUnapprovedSrFlag sets BillUnapprovedSrFlag field to given value.

### HasBillUnapprovedSrFlag

`func (o *CompanyFinance) HasBillUnapprovedSrFlag() bool`

HasBillUnapprovedSrFlag returns a boolean if a field has been set.

### SetBillUnapprovedSrFlagNil

`func (o *CompanyFinance) SetBillUnapprovedSrFlagNil(b bool)`

 SetBillUnapprovedSrFlagNil sets the value for BillUnapprovedSrFlag to be an explicit nil

### UnsetBillUnapprovedSrFlag
`func (o *CompanyFinance) UnsetBillUnapprovedSrFlag()`

UnsetBillUnapprovedSrFlag ensures that no value is present for BillUnapprovedSrFlag, not even an explicit nil
### GetBillRestrictPmFlag

`func (o *CompanyFinance) GetBillRestrictPmFlag() bool`

GetBillRestrictPmFlag returns the BillRestrictPmFlag field if non-nil, zero value otherwise.

### GetBillRestrictPmFlagOk

`func (o *CompanyFinance) GetBillRestrictPmFlagOk() (*bool, bool)`

GetBillRestrictPmFlagOk returns a tuple with the BillRestrictPmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillRestrictPmFlag

`func (o *CompanyFinance) SetBillRestrictPmFlag(v bool)`

SetBillRestrictPmFlag sets BillRestrictPmFlag field to given value.

### HasBillRestrictPmFlag

`func (o *CompanyFinance) HasBillRestrictPmFlag() bool`

HasBillRestrictPmFlag returns a boolean if a field has been set.

### SetBillRestrictPmFlagNil

`func (o *CompanyFinance) SetBillRestrictPmFlagNil(b bool)`

 SetBillRestrictPmFlagNil sets the value for BillRestrictPmFlag to be an explicit nil

### UnsetBillRestrictPmFlag
`func (o *CompanyFinance) UnsetBillRestrictPmFlag()`

UnsetBillRestrictPmFlag ensures that no value is present for BillRestrictPmFlag, not even an explicit nil
### GetBillCompletePmFlag

`func (o *CompanyFinance) GetBillCompletePmFlag() bool`

GetBillCompletePmFlag returns the BillCompletePmFlag field if non-nil, zero value otherwise.

### GetBillCompletePmFlagOk

`func (o *CompanyFinance) GetBillCompletePmFlagOk() (*bool, bool)`

GetBillCompletePmFlagOk returns a tuple with the BillCompletePmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillCompletePmFlag

`func (o *CompanyFinance) SetBillCompletePmFlag(v bool)`

SetBillCompletePmFlag sets BillCompletePmFlag field to given value.

### HasBillCompletePmFlag

`func (o *CompanyFinance) HasBillCompletePmFlag() bool`

HasBillCompletePmFlag returns a boolean if a field has been set.

### SetBillCompletePmFlagNil

`func (o *CompanyFinance) SetBillCompletePmFlagNil(b bool)`

 SetBillCompletePmFlagNil sets the value for BillCompletePmFlag to be an explicit nil

### UnsetBillCompletePmFlag
`func (o *CompanyFinance) UnsetBillCompletePmFlag()`

UnsetBillCompletePmFlag ensures that no value is present for BillCompletePmFlag, not even an explicit nil
### GetBillUnapprovedPmFlag

`func (o *CompanyFinance) GetBillUnapprovedPmFlag() bool`

GetBillUnapprovedPmFlag returns the BillUnapprovedPmFlag field if non-nil, zero value otherwise.

### GetBillUnapprovedPmFlagOk

`func (o *CompanyFinance) GetBillUnapprovedPmFlagOk() (*bool, bool)`

GetBillUnapprovedPmFlagOk returns a tuple with the BillUnapprovedPmFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillUnapprovedPmFlag

`func (o *CompanyFinance) SetBillUnapprovedPmFlag(v bool)`

SetBillUnapprovedPmFlag sets BillUnapprovedPmFlag field to given value.

### HasBillUnapprovedPmFlag

`func (o *CompanyFinance) HasBillUnapprovedPmFlag() bool`

HasBillUnapprovedPmFlag returns a boolean if a field has been set.

### SetBillUnapprovedPmFlagNil

`func (o *CompanyFinance) SetBillUnapprovedPmFlagNil(b bool)`

 SetBillUnapprovedPmFlagNil sets the value for BillUnapprovedPmFlag to be an explicit nil

### UnsetBillUnapprovedPmFlag
`func (o *CompanyFinance) UnsetBillUnapprovedPmFlag()`

UnsetBillUnapprovedPmFlag ensures that no value is present for BillUnapprovedPmFlag, not even an explicit nil
### GetEmailTemplate

`func (o *CompanyFinance) GetEmailTemplate() EmailTemplateReference`

GetEmailTemplate returns the EmailTemplate field if non-nil, zero value otherwise.

### GetEmailTemplateOk

`func (o *CompanyFinance) GetEmailTemplateOk() (*EmailTemplateReference, bool)`

GetEmailTemplateOk returns a tuple with the EmailTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailTemplate

`func (o *CompanyFinance) SetEmailTemplate(v EmailTemplateReference)`

SetEmailTemplate sets EmailTemplate field to given value.

### HasEmailTemplate

`func (o *CompanyFinance) HasEmailTemplate() bool`

HasEmailTemplate returns a boolean if a field has been set.

### GetCompany

`func (o *CompanyFinance) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyFinance) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyFinance) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyFinance) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyFinance) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyFinance) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyFinance) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyFinance) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetCustomFields

`func (o *CompanyFinance) GetCustomFields() []CustomFieldValue`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CompanyFinance) GetCustomFieldsOk() (*[]CustomFieldValue, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CompanyFinance) SetCustomFields(v []CustomFieldValue)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CompanyFinance) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


