# CompanyCustomNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CustomNote** | **string** |  Max length: 1500; | 
**Status** | [**CompanyStatusReference**](CompanyStatusReference.md) |  | 
**Company** | Pointer to [**CompanyReference**](CompanyReference.md) |  | [optional] 
**Info** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCompanyCustomNote

`func NewCompanyCustomNote(customNote string, status CompanyStatusReference, ) *CompanyCustomNote`

NewCompanyCustomNote instantiates a new CompanyCustomNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompanyCustomNoteWithDefaults

`func NewCompanyCustomNoteWithDefaults() *CompanyCustomNote`

NewCompanyCustomNoteWithDefaults instantiates a new CompanyCustomNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CompanyCustomNote) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CompanyCustomNote) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CompanyCustomNote) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CompanyCustomNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomNote

`func (o *CompanyCustomNote) GetCustomNote() string`

GetCustomNote returns the CustomNote field if non-nil, zero value otherwise.

### GetCustomNoteOk

`func (o *CompanyCustomNote) GetCustomNoteOk() (*string, bool)`

GetCustomNoteOk returns a tuple with the CustomNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomNote

`func (o *CompanyCustomNote) SetCustomNote(v string)`

SetCustomNote sets CustomNote field to given value.


### GetStatus

`func (o *CompanyCustomNote) GetStatus() CompanyStatusReference`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompanyCustomNote) GetStatusOk() (*CompanyStatusReference, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompanyCustomNote) SetStatus(v CompanyStatusReference)`

SetStatus sets Status field to given value.


### GetCompany

`func (o *CompanyCustomNote) GetCompany() CompanyReference`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CompanyCustomNote) GetCompanyOk() (*CompanyReference, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CompanyCustomNote) SetCompany(v CompanyReference)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CompanyCustomNote) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetInfo

`func (o *CompanyCustomNote) GetInfo() map[string]string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *CompanyCustomNote) GetInfoOk() (*map[string]string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *CompanyCustomNote) SetInfo(v map[string]string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *CompanyCustomNote) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


